package main

import (
	"net/http"

	"github.com/justinas/alice"
)

func (a *app) routes() http.Handler {

	mux := http.NewServeMux() // If not declared -> DefaultServerMux

	// Routes
	mux.HandleFunc("/", a.home)
	mux.HandleFunc("/snippet", a.showSnippet)
	mux.HandleFunc("/snippet/create", a.createSnippet)

	// Static Files
	fileServer := http.FileServer(http.Dir("./web/static/"))
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	// Middleware chain
	mw := alice.New(a.recoverPanic, a.logRequest, secureHeaders)

	// middleware -> serveMux -> ... ↩
	return mw.Then(mux)
}
