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

func TestRequest(t *testing.T) {
	var handler http.HandlerFunc = func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, r.Method)
		fmt.Fprintln(w, r.RequestURI)
		fmt.Fprintln(w, r.Body)
	}

	server := http.Server{
		Addr:    "localhost:8000",
		Handler: handler,
	}
	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}

}
