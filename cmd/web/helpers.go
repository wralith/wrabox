package main

import (
	"fmt"
	"net/http"
	"runtime/debug"
)

func (a *App) serverError(w http.ResponseWriter, err error) {
	trace := fmt.Sprintf("%s\n%s", err.Error(), debug.Stack())
	// Error happens where this function called not here, so depth = 2
	a.errorLog.Output(2, trace)

	http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
}

func (a *App) clientError(w http.ResponseWriter, status int) {
	http.Error(w, http.StatusText(status), status)
}

func (a *App) notFound(w http.ResponseWriter) {
	a.clientError(w, http.StatusNotFound)
}
