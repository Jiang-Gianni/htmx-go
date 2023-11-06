package api

import (
	"context"
	"database/sql"
	"io/fs"

	database "github.com/Jiang-Gianni/htmx-go/db"
	"github.com/Jiang-Gianni/htmx-go/mylogger"
)

type Api struct {
	bgCtx    context.Context
	db       *sql.DB
	query    *database.Queries
	log      *mylogger.Logger
	assetsFs fs.FS
}

func New(ctx context.Context, db *sql.DB, log *mylogger.Logger, assetsFs fs.FS) *Api {
	return &Api{
		bgCtx:    ctx,
		db:       db,
		query:    database.New(db),
		log:      log,
		assetsFs: assetsFs,
	}
}
