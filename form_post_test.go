package golang_web

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func FormPost(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		panic(err)
	}

	// r.PostFormValue("first_name") // method ini digunkan untuk mengambil data tanpa parsing menggunakan ParseForm

	firstName := r.PostForm.Get("first_name") // cara manual
	lastName := r.PostForm.Get("last_name")

	fmt.Fprintf(w, "Hello %s %s", firstName, lastName)
}

func TestFormPost(t *testing.T) {
	requestBody := strings.NewReader("first_name=Hendrik&last_name=Array") // untuk data di request body menggunakan form post
	request := httptest.NewRequest(http.MethodPost, "http://localhost:8000", requestBody)
	request.Header.Add("Content-type", "application/x-www-form-urlencoded") // wajib tambahkan header ini jika menggunakan form post untuk mengirim request body

	recorder := httptest.NewRecorder()

	FormPost(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)
	fmt.Println(string(body))
}
