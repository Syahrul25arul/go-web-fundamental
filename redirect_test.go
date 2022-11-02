package golang_web

import (
	"fmt"
	"net/http"
	"testing"
)

func RedirectTo(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello from redirect")
}

func RedirectFrom(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/redirect-to", http.StatusTemporaryRedirect)
}

func RedirectYoutube(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "https://www.youtube.com", http.StatusTemporaryRedirect)
}

func TestRedirect(t *testing.T) {
	mux := http.NewServeMux()

	mux.HandleFunc("/redirect-from", RedirectFrom)
	mux.HandleFunc("/redirect-to", RedirectTo)
	mux.HandleFunc("/redirect-youtube", RedirectYoutube)

	serve := http.Server{
		Addr:    "localhost:8080",
		Handler: mux,
	}

	if err := serve.ListenAndServe(); err != nil {
		panic(err)
	}
}
