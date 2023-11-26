package logger

import (
	"net/http"
	"testing"
)

func TestOpenObserveWriter_Write(t *testing.T) {
	type fields struct {
		URL      string
		User     string
		Password string
		Client   *http.Client
	}
	type args struct {
		p []byte
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantN   int
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			o := OpenObserveWriter{
				URL:      tt.fields.URL,
				User:     tt.fields.User,
				Password: tt.fields.Password,
				Client:   tt.fields.Client,
			}
			gotN, err := o.Write(tt.args.p)
			if (err != nil) != tt.wantErr {
				t.Errorf("OpenObserveWriter.Write(%v) error = %v, wantErr %v", tt.args.p, err, tt.wantErr)
				return
			}
			if gotN != tt.wantN {
				t.Errorf("OpenObserveWriter.Write(%v) = %v, want %v", tt.args.p, gotN, tt.wantN)
			}
		})
	}
}
