package golang_web

import (
	"fmt"
	"net/http"
	"testing"
)

func TestServeMux(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hello World")
	})

	mux.HandleFunc("/hi", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hi")
	})

	mux.HandleFunc("/images/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Images ")
	})

	mux.HandleFunc("/images/thumbnail", func(w http.ResponseWriter, r *http.Request) { // jika tidak menambahkan slash, maka akan ditembak ke endpoint /images/ jika url tidak persis dengan /images/thumbnail
		fmt.Fprint(w, "Thumbnail")
	})

	server := http.Server{
		Addr:    "localhost:8000",
		Handler: mux,
	}

	if err := server.ListenAndServe(); err != nil {
		panic(err)
	}
}
