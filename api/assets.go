package api

import (
	"io"
	"net/http"
)

type gzipResponseWriter struct {
	io.Writer
	http.ResponseWriter
}

func (w gzipResponseWriter) Write(b []byte) (int, error) {
	return w.Writer.Write(b)
}

func (a *Api) getAssets() http.HandlerFunc {
	// var (
	// 	gz  *gzip.Writer
	// 	gzr gzipResponseWriter
	// 	err error
	// )
	fs := http.StripPrefix("/assets/", http.FileServer(http.FS(a.assetsFs)))
	return func(w http.ResponseWriter, r *http.Request) {
		// w.Header().Set("Cache-Control", "max-age=2592000")
		w.Header().Set("Cache-Control", "no-store")
		// w.Header().Set("Content-Encoding", "gzip")
		// gz, err = gzip.NewWriterLevel(w, gzip.BestCompression)
		// if err != nil {
		// 	a.log.Error(err.Error())
		// }
		// defer gz.Close()
		// gzr = gzipResponseWriter{Writer: gz, ResponseWriter: w}
		// fs.ServeHTTP(gzr, r)
		fs.ServeHTTP(w, r)
	}
}
