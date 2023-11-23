package main

import (
	"context"
	"database/sql"
	"embed"
	"io/fs"
	"log"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/Jiang-Gianni/htmx-go/api"
	"github.com/joho/godotenv"
	_ "github.com/mattn/go-sqlite3"
)

//go:embed all:assets
var assetsFs embed.FS

// These variables will be set by -ldflags during the compilation and printed out in main
// go build -ldflags="-X main.environment=${ENV} -X main.gitCommit=${GIT_COMMIT}" main.go
var environment = "DEV"
var gitCommit = "gitCommit"

func main() {
	if err := run(); err != nil {
		log.Println("error :", err)
		os.Exit(1)
	}
}

func run() error {

	// Background Context
	ctx := context.Background()

	// Database Configurations
	godotenv.Load()
	dbDriver := "sqlite3"
	dbConnection := os.Getenv("DB_CONNECTION")

	// Server Configurations
	srv := &http.Server{
		Addr:         ":3000",
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 5 * time.Second,
		IdleTimeout:  5 * time.Second,
	}

	// Logger
	logWriter := os.Stdout

	// logWriter := api.OpenObserveWriter{}
	logHandler := slog.NewJSONHandler(logWriter, &slog.HandlerOptions{
		Level:     slog.LevelDebug,
		AddSource: true,
		ReplaceAttr: func(groups []string, a slog.Attr) slog.Attr {
			if a.Key == slog.TimeKey {
				a.Value = slog.TimeValue(time.Now().Local().Truncate(time.Millisecond))
			}
			return a
		},
	})
	mySlog := slog.New(logHandler)

	mySlog.Info("START", "environment", environment, "gitCommit", gitCommit, "port", srv.Addr)

	// Assets folder to be served
	fsys, err := fs.Sub(assetsFs, "assets")
	if err != nil {
		return err
	}

	// Database
	db, err := sql.Open(dbDriver, dbConnection)
	if err != nil {
		return err
	}

	// Api initialization
	myApi := api.New(ctx, db, mySlog, fsys)
	srv.Handler = myApi.MountHandlers()

	// Error channel to listen to errors
	errs := make(chan error)

	// Start of server
	go func() {
		mySlog.Info("LISTENING", "port", srv.Addr)
		errs <- srv.ListenAndServe()
	}()

	// Graceful Shutdown
	go func() {
		errs <- GracefulShutdown(ctx, srv)
	}()

	// Return when first error occurs
	return <-errs
}

func GracefulShutdown(ctx context.Context, server *http.Server) error {
	sigChn := make(chan os.Signal, 1)
	signal.Notify(sigChn, os.Interrupt, syscall.SIGTERM)
	<-sigChn
	timeout := time.Duration(5*time.Second) * time.Second
	ctx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()
	return server.Shutdown(ctx)
}
