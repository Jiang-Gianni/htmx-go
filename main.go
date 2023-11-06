package main

import (
	"context"
	"database/sql"
	"embed"
	"io/fs"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/Jiang-Gianni/htmx-go/api"
	"github.com/Jiang-Gianni/htmx-go/mylogger"
	"github.com/joho/godotenv"
	_ "github.com/mattn/go-sqlite3"
)

//go:embed all:assets
var assetsFs embed.FS

// These variables will be set by -ldflags during the compilation and printed out in main
// go build -ldflags="-X main.environment=${ENV} -X main.gitBranch=${GIT_BRANCH}" main.go
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
	myLog := mylogger.New(os.Stdout, log.LstdFlags|log.Lshortfile)
	myLog.Info.Println("Environment: ", environment, "arstneio")
	myLog.Info.Println("Git Commit: ", gitCommit)

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
	myApi := api.New(ctx, db, myLog, fsys)
	srv.Handler = myApi.MountHandlers()

	// Error channel to listen to errors
	errs := make(chan error)

	// Start of server
	go func() {
		myLog.Info.Println("Listening on port: ", srv.Addr)
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
