package golang_web

import (
	"fmt"
	"html/template"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TemplateActionIf(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("./templates/if.gohtml"))
	t.ExecuteTemplate(w, "if.gohtml", map[string]interface{}{
		"Title": "Template data map",
	})
}

func TestTemplateActionIf(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	recorder := httptest.NewRecorder()

	TemplateActionIf(recorder, request)

	body, _ := io.ReadAll(recorder.Result().Body)

	fmt.Println(string(body))
}

func TemplateActionOperator(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("./templates/comparator.gohtml"))
	t.ExecuteTemplate(w, "comparator.gohtml", map[string]interface{}{
		"Title":      "Template data map",
		"FinalValue": 40,
	})
}

func TestTemplateActionOperator(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	recorder := httptest.NewRecorder()

	TemplateActionOperator(recorder, request)

	body, _ := io.ReadAll(recorder.Result().Body)

	fmt.Println(string(body))
}

func TemplateActionRange(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("./templates/range.gohtml"))
	t.ExecuteTemplate(w, "range.gohtml", map[string]interface{}{
		"Title": "Test Range",
		"Hobbies": []string{
			"Tai", "Babi",
		},
	})
}

func TestTemplateActionRange(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	recorder := httptest.NewRecorder()

	TemplateActionRange(recorder, request)

	body, _ := io.ReadAll(recorder.Result().Body)

	fmt.Println(string(body))
}

func TemplateActionWith(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("./templates/addres.gohtml"))
	t.ExecuteTemplate(w, "addres.gohtml", map[string]interface{}{
		"Title": "Test Range",
		"Name":  "Hendrik",
		"Address": map[string]string{
			"Street": "Jalan Belum Ada",
			"City":   "Jakarta",
		},
	})
}

func TestTemplateActionWith(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	recorder := httptest.NewRecorder()

	TemplateActionWith(recorder, request)

	body, _ := io.ReadAll(recorder.Result().Body)

	fmt.Println(string(body))
}

func TemplateActionLayout(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseGlob("./templates/*.gohtml"))
	t.ExecuteTemplate(w, "layout", map[string]interface{}{
		"Title": "Test Layout",
		"Hobbies": map[string]string{
			"Name":  "Hendrik",
			"Hobby": "Koco",
		},
		"Hobbies2": []map[string]string{
			{
				"Name":  "Hendrik",
				"Hobby": "Koco",
			},
			{
				"Name":  "Jamal",
				"Hobby": "Baca",
			},
		},
	})
	t.ExecuteTemplate(w, "test", nil)
}

func TestTemplateActionLayout(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	recorder := httptest.NewRecorder()

	TemplateActionLayout(recorder, request)

	body, _ := io.ReadAll(recorder.Result().Body)

	fmt.Println(string(body))
}
