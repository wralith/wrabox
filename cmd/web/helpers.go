package main

import (
	"bytes"
	"fmt"
	"net/http"
	"runtime/debug"
	"time"
)

func (a *app) serverError(w http.ResponseWriter, err error) {
	trace := fmt.Sprintf("%s\n%s", err.Error(), debug.Stack())
	// Error happens where this function called not here, so depth = 2
	a.errorLog.Output(2, trace)

	http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
}

func (a *app) clientError(w http.ResponseWriter, status int) {
	http.Error(w, http.StatusText(status), status)
}

func (a *app) notFound(w http.ResponseWriter) {
	a.clientError(w, http.StatusNotFound)
}

func (a *app) render(w http.ResponseWriter, r *http.Request, name string, td *templateData) {
	ts, ok := a.templateCache[name]
	if !ok {
		a.serverError(w, fmt.Errorf("there is %s template", name))
		return
	}

	buf := new(bytes.Buffer)

	err := ts.Execute(buf, a.addDefaultData(td, r))
	if err != nil {
		a.serverError(w, err)
		return
	}

	buf.WriteTo(w)
}

func (a *app) addDefaultData(td *templateData, r *http.Request) *templateData {
	if td == nil {
		td = &templateData{}
	}
	td.CurrentYear = time.Now().Year()
	return td

}
