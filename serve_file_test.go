package golang_web

import (
	_ "embed"
	"fmt"
	"net/http"
	"testing"
)

func ServeFile(w http.ResponseWriter, r *http.Request) {
	// http.ServeFile di gunakan untuk meload file statis yang kita inginkan
	if r.URL.Query().Get("name") != "" {
		http.ServeFile(w, r, "./resources/ok.html")
	} else {
		http.ServeFile(w, r, "./resources/notfound.html")
	}
}

func TestServeFile(t *testing.T) {
	serve := http.Server{
		Addr:    "localhost:8080",
		Handler: http.HandlerFunc(ServeFile),
	}

	err := serve.ListenAndServe()
	if err != nil {
		panic(err)
	}

}

//go:embed resources/ok.html
var resourcesOk string

//go:embed resources/notfound.html
var resourcesNotFound string

func ServeFileEmbed(w http.ResponseWriter, r *http.Request) {
	// jika menggunakan embed maka tidak perlu menggunakan http.ServeFile
	if r.URL.Query().Get("name") != "" {
		fmt.Fprint(w, resourcesOk)
	} else {
		fmt.Fprint(w, resourcesNotFound)
	}
}

func TestServeFileEmbed(t *testing.T) {
	serve := http.Server{
		Addr:    "localhost:8080",
		Handler: http.HandlerFunc(ServeFileEmbed),
	}

	err := serve.ListenAndServe()
	if err != nil {
		panic(err)
	}

}
