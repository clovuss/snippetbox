package main

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"
)

func (app *application) home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		app.NotfoundError(w)
		return
	}
	files := []string{"./ui/html/home.page.tmpl",
		"./ui/html/base.layout.tmpl",
		"./ui/html/footer.partial.tmpl"}
	temp, err := template.ParseFiles(files...)
	if err != nil {
		app.ServerError(w, err)
		return
	}
	err = temp.Execute(w, nil)

	if err != nil {
		app.ServerError(w, err)
	}
}

func (app *application) showSnippet(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id < 1 {
		app.NotfoundError(w)
		return
	}

	_, err = fmt.Fprintf(w, "Сниппет с id %d", id)
	if err != nil {
		app.ServerError(w, err)
	}

}
func (app *application) createSnippet(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		w.Header().Set("allow", http.MethodPost)

		app.ClientError(w, http.StatusMethodNotAllowed)
		return
	}
	w.Write([]byte("Создаем заметку!\n"))
}
