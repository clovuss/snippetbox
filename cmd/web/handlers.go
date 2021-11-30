package main

import (
	"errors"
	"fmt"
	"github.com/clovuss/snippetbox/pkg/models"
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
	res, err := app.snippets.Get(id)
	if err != nil {
		if errors.Is(err, models.ErrNoRecord) {
			app.NotfoundError(w)
		} else {
			fmt.Println(err)
			app.ServerError(w, err)
		}
		return
	}

	_, err = fmt.Fprint(w, "Сниппет с id ", res)
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

	title := "Пост про улитку"
	text := "Улитка выползла из раковины,\nвытянула рожки,\nи опять подобрала их."
	err := app.snippets.Insert(title, text)
	if err != nil {
		println(err)
	}
	w.Write([]byte("Создаем заметку!\n"))
}
