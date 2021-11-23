package main

import (
	"fmt"
	"net/http"
	"runtime/debug"
)

func (app *application) ServerError(wr http.ResponseWriter, err error) {
	trace := fmt.Sprintf("Ошибка сервера %s\n%s", err.Error(), debug.Stack())
	app.errorLog.Output(2, trace)
	http.Error(wr, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
}
func (app *application) ClientError(wr http.ResponseWriter, status int) {
	http.Error(wr, http.StatusText(status), status)
}

func (app *application) NotfoundError(wr http.ResponseWriter) {
	app.ClientError(wr, http.StatusNotFound)
}
