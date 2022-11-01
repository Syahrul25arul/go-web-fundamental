package golang_web

import (
	"fmt"
	"html/template"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

type MyPage struct {
	Name string
}

func (m MyPage) SayHello(name string) string {
	return "Hello " + name + ", My name is " + m.Name
}

func TemplateFunction(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.New("FUNCTION").Parse(`{{.SayHello "Hendrik"}}`))
	t.ExecuteTemplate(w, "FUNCTION", MyPage{Name: "Array"})
}

func TestTemplateFunction(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	recorder := httptest.NewRecorder()

	TemplateFunction(recorder, request)

	body, _ := io.ReadAll(recorder.Result().Body)

	fmt.Println(string(body))
}

func TemplateFunctionGlobal(w http.ResponseWriter, r *http.Request) {
	// len adalah function global, tidak perlu menggunakan titik untuk memanggil function global di template
	t := template.Must(template.New("FUNCTION").Parse(`{{len .Name}}`))
	t.ExecuteTemplate(w, "FUNCTION", MyPage{Name: "Array"})
}

func TestTemplateFunctionGlobal(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	recorder := httptest.NewRecorder()

	TemplateFunctionGlobal(recorder, request)

	body, _ := io.ReadAll(recorder.Result().Body)

	fmt.Println(string(body))
}

func TemplateFunctionCreateGlobal(w http.ResponseWriter, r *http.Request) {
	// membuat function global sendiri harus di registrasikan terlebih dahulu sebelum parsing template
	t := template.New("FUNCTION")
	t = t.Funcs(map[string]interface{}{
		"upper": func(name string) string {
			return strings.ToUpper(name)
		},
	})
	t = template.Must(t.Parse(`{{ upper .Name}}`))
	t.ExecuteTemplate(w, "FUNCTION", MyPage{Name: "rizal"})
}

func TestTemplateFunctionCreateGlobal(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	recorder := httptest.NewRecorder()

	TemplateFunctionCreateGlobal(recorder, request)

	body, _ := io.ReadAll(recorder.Result().Body)

	fmt.Println(string(body))
}

func TemplateFunctionPipeline(w http.ResponseWriter, r *http.Request) {
	// function pipeline adalah cara meneruskan hasil function ke function berikutnya
	t := template.New("FUNCTION")
	t = t.Funcs(map[string]interface{}{
		"sayHello": func(name string) string {
			return "Hello " + name
		},
		"upper": func(value string) string {
			return strings.ToUpper(value)
		},
	})
	t = template.Must(t.Parse(`{{ sayHello .Name | upper }}`))
	t.ExecuteTemplate(w, "FUNCTION", MyPage{Name: "rizal"})
}

func TestTemplateFunctionPipeline(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	recorder := httptest.NewRecorder()

	TemplateFunctionPipeline(recorder, request)

	body, _ := io.ReadAll(recorder.Result().Body)

	fmt.Println(string(body))
}
