package api

import (
	"context"
	"database/sql"
	"io/fs"
	"log/slog"
	"reflect"
	"testing"
)

func TestNew(t *testing.T) {
	type args struct {
		ctx      context.Context
		database *sql.DB
		log      *slog.Logger
		assetsFs fs.FS
	}
	tests := []struct {
		name string
		args args
		want *Api
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if got := New(tt.args.ctx, tt.args.database, tt.args.log, tt.args.assetsFs); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("New(%v, %v, %v, %v) = %v, want %v", tt.args.ctx, tt.args.database, tt.args.log, tt.args.assetsFs, got, tt.want)
			}
		})
	}
}
