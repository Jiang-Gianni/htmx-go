package api

import (
	"context"
	"database/sql"
	"io/fs"
	"log/slog"
	"net/http"
	"reflect"
	"testing"

	"github.com/Jiang-Gianni/htmx-go/db"
	"github.com/go-chi/chi/v5"
)

func TestApi_MountHandlers(t *testing.T) {
	type fields struct {
		bgCtx    context.Context
		db       *sql.DB
		query    *db.Queries
		log      *slog.Logger
		assetsFs fs.FS
	}
	tests := []struct {
		name   string
		fields fields
		want   *chi.Mux
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			a := &Api{
				bgCtx:    tt.fields.bgCtx,
				db:       tt.fields.db,
				query:    tt.fields.query,
				log:      tt.fields.log,
				assetsFs: tt.fields.assetsFs,
			}
			if got := a.MountHandlers(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Api.MountHandlers() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestApi_AllowedMiddeware(t *testing.T) {
	type fields struct {
		bgCtx    context.Context
		db       *sql.DB
		query    *db.Queries
		log      *slog.Logger
		assetsFs fs.FS
	}
	type args struct {
		next http.Handler
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   http.Handler
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			a := &Api{
				bgCtx:    tt.fields.bgCtx,
				db:       tt.fields.db,
				query:    tt.fields.query,
				log:      tt.fields.log,
				assetsFs: tt.fields.assetsFs,
			}
			if got := a.AllowedMiddeware(tt.args.next); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Api.AllowedMiddeware(%v) = %v, want %v", tt.args.next, got, tt.want)
			}
		})
	}
}

func TestApi_Allowed(t *testing.T) {
	type fields struct {
		bgCtx    context.Context
		db       *sql.DB
		query    *db.Queries
		log      *slog.Logger
		assetsFs fs.FS
	}
	type args struct {
		r http.Request
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			a := &Api{
				bgCtx:    tt.fields.bgCtx,
				db:       tt.fields.db,
				query:    tt.fields.query,
				log:      tt.fields.log,
				assetsFs: tt.fields.assetsFs,
			}
			if got := a.Allowed(tt.args.r); got != tt.want {
				t.Errorf("Api.Allowed(%v) = %v, want %v", tt.args.r, got, tt.want)
			}
		})
	}
}
