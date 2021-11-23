package main

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"
)

func (app *application) home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	files := []string{"./ui/html/home.page.tmpl",
		"./ui/html/base.layout.tmpl",
		"./ui/html/footer.partial.tmpl"}
	temp, err := template.ParseFiles(files...)
	if err != nil {

		http.Error(w, "something wrong", 500)
		app.errorLog.Println(err)
		return
	}
	err = temp.Execute(w, nil)

	if err != nil {
		http.Error(w, "something wrong", 500)
		app.errorLog.Println(err.Error())
	}
}

func (app *application) showSnippet(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id < 1 {
		http.NotFound(w, r)
		return
	}

	_, err = fmt.Fprintf(w, "Сниппет с id %d", id)
	if err != nil {
		fmt.Println(err)
	}

}
func (app *application) createSnippet(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		w.Header().Set("allow", http.MethodPost)

		http.Error(w, "Запрещено", 405)
		return
	}
	w.Write([]byte("Создаем заметку!\n"))
}
