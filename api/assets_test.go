package api

import (
	"context"
	"database/sql"
	"io"
	"io/fs"
	"log/slog"
	"net/http"
	"reflect"
	"testing"

	"github.com/Jiang-Gianni/htmx-go/db"
)

func Test_gzipResponseWriter_Write(t *testing.T) {
	type fields struct {
		Writer         io.Writer
		ResponseWriter http.ResponseWriter
	}
	type args struct {
		b []byte
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    int
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			w := gzipResponseWriter{
				Writer:         tt.fields.Writer,
				ResponseWriter: tt.fields.ResponseWriter,
			}
			got, err := w.Write(tt.args.b)
			if (err != nil) != tt.wantErr {
				t.Errorf("gzipResponseWriter.Write(%v) error = %v, wantErr %v", tt.args.b, err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("gzipResponseWriter.Write(%v) = %v, want %v", tt.args.b, got, tt.want)
			}
		})
	}
}

func TestApi_getAssets(t *testing.T) {
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
		want   http.HandlerFunc
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
			if got := a.getAssets(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Api.getAssets() = %v, want %v", got, tt.want)
			}
		})
	}
}
