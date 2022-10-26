package golang_web

import (
	"embed"
	"io/fs"
	"net/http"
	"testing"
)

func TestFileServer(t *testing.T) {
	directory := http.Dir("./resources")     // untuk mengambil dari direktory
	fileServer := http.FileServer(directory) // membuat handle file server untuk meload file

	mux := http.NewServeMux()
	mux.Handle("/static/", http.StripPrefix("/static", fileServer)) // http.StripPrefix, untuk menghapus prefix dari url misal "/static" nya dihapus

	server := http.Server{
		Addr:    "localhost:8000",
		Handler: mux,
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}

//go:embed resources
var resources embed.FS // untuk mengintegrasikan fitur embed dan fileServer

func TestFileServerGoEmbed(t *testing.T) {
	directory, _ := fs.Sub(resources, "resources")    // untuk masuk kedalam directory dari file yang di embed ke file server
	fileServer := http.FileServer(http.FS(directory)) // membuat handle file server untuk meload file

	mux := http.NewServeMux()
	mux.Handle("/static/", http.StripPrefix("/static", fileServer)) // http.StripPrefix, untuk menghapus prefix dari url misal "/static" nya dihapus

	server := http.Server{
		Addr:    "localhost:8000",
		Handler: mux,
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
