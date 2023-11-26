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
)

func TestApi_getFirst(t *testing.T) {
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
			if got := a.getFirst(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Api.getFirst() = %v, want %v", got, tt.want)
			}
		})
	}
}
