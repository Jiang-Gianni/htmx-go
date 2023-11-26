package api

import (
	"context"
	"database/sql"
	"io/fs"
	"log/slog"

	"github.com/Jiang-Gianni/htmx-go/db"
)

type Api struct {
	bgCtx    context.Context
	db       *sql.DB
	query    *db.Queries
	log      *slog.Logger
	assetsFs fs.FS
}

func New(ctx context.Context, database *sql.DB, log *slog.Logger, assetsFs fs.FS) *Api {
	return &Api{
		bgCtx:    ctx,
		db:       database,
		query:    db.New(database),
		log:      log,
		assetsFs: assetsFs,
	}
}
