package golang_web

import (
	"embed"
	"fmt"
	"html/template"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func SimpleHTML(w http.ResponseWriter, r *http.Request) {
	templateText := `<html><body>{{.}}</body></html>`

	// t,err :=template.New("SIMPLE").Parse(templateText)
	// if err !=nil {
	// 	panic(err)
	// }

	t := template.Must(template.New("SIMPLE").Parse(templateText))

	t.ExecuteTemplate(w, "SIMPLE", "Hello Template World")
}

func TestSimpleHTML(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	recorder := httptest.NewRecorder()

	SimpleHTML(recorder, request)

	body, _ := io.ReadAll(recorder.Result().Body)

	fmt.Println(string(body))
}

func SimpleHTMLFile(w http.ResponseWriter, r *http.Request) {
	// t,err :=template.New("SIMPLE").Parse(templateText)
	// if err !=nil {
	// 	panic(err)
	// }

	t := template.Must(template.ParseFiles("./templates/simple.gohtml"))

	t.ExecuteTemplate(w, "simple.gohtml", "Hello Template World") // kalau pake parsefiles, nama templatenya di samakan dengan nama file
}

func TestSimpleHTMLFile(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	recorder := httptest.NewRecorder()

	SimpleHTMLFile(recorder, request)

	body, _ := io.ReadAll(recorder.Result().Body)

	fmt.Println(string(body))
}

func TemplateDirectory(w http.ResponseWriter, r *http.Request) {
	// t,err :=template.New("SIMPLE").Parse(templateText)
	// if err !=nil {
	// 	panic(err)
	// }

	t := template.Must(template.ParseGlob("./templates/*.gohtml")) // untuk meload semua file template pada sebuah direktori
	t.ExecuteTemplate(w, "simple.gohtml", "Hello Template World")  // tinggal di panggil saja nama file nya
}

func TestTemplateDirectory(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	recorder := httptest.NewRecorder()

	TemplateDirectory(recorder, request)

	body, _ := io.ReadAll(recorder.Result().Body)

	fmt.Println(string(body))
}

//go:embed templates/*.gohtml
var templates embed.FS

func TemplateEmbed(w http.ResponseWriter, r *http.Request) {
	// t,err :=template.New("SIMPLE").Parse(templateText)
	// if err !=nil {
	// 	panic(err)
	// }

	// menggunakan embed menjadikan file hanya di simpan satu kali ke binary sistem golangnya dan tidak perlu memanggil filenya lagi
	// sehingga jika filenya dihapus tidak akan terjadi error
	t := template.Must(template.ParseFS(templates, "templates/*gohtml")) // menggunakan file embed tidak perlu ditambah "." untuk penyebutan direktorynya
	t.ExecuteTemplate(w, "simple.gohtml", "Hello Template World")        // tinggal di panggil saja nama file nya
}

func TestTemplateEmbed(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	recorder := httptest.NewRecorder()

	TemplateEmbed(recorder, request)

	body, _ := io.ReadAll(recorder.Result().Body)

	fmt.Println(string(body))
}
