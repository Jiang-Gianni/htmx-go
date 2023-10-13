package main

import (
	"embed"
	"io/fs"
	"log"

	"github.com/Jiang-Gianni/htmx-go/server"
)

//go:embed all:assets
var assetsFs embed.FS

func main() {
	fsys, err := fs.Sub(assetsFs, "assets")
	if err != nil {
		log.Fatal(err)
	}
	server.RegisterAndRun(fsys)
}
