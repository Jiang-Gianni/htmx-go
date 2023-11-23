package api

import (
	"context"
	"database/sql"
	"io/fs"
	"log/slog"

	database "github.com/Jiang-Gianni/htmx-go/db"
)

type Api struct {
	bgCtx    context.Context
	db       *sql.DB
	query    *database.Queries
	log      *slog.Logger
	assetsFs fs.FS
}

func New(ctx context.Context, db *sql.DB, log *slog.Logger, assetsFs fs.FS) *Api {
	return &Api{
		bgCtx:    ctx,
		db:       db,
		query:    database.New(db),
		log:      log,
		assetsFs: assetsFs,
	}
}
