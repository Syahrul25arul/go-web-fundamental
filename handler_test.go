package golang_web

import (
	"fmt"
	"net/http"
	"testing"
)

func TestHandler(t *testing.T) {
	var handler http.HandlerFunc = func(w http.ResponseWriter, r *http.Request) { // writer adalah response yang akan diberikan ke client, request adalah request yang diterima dari client
		// logic web
		fmt.Fprint(w, "Hello World")
	}

	server := http.Server{
		Addr:    "localhost:8000",
		Handler: handler,
	}

	if err := server.ListenAndServe(); err != nil {
		panic(err)
	}

}
