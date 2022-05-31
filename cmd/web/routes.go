package main

import (
	"net/http"

	"github.com/bmizerany/pat"
	"github.com/justinas/alice"
)

func (a *app) routes() http.Handler {

	// mux := http.NewServeMux() // If not declared -> DefaultServerMux
	mux := pat.New()

	// Routes
	mux.Get("/", http.HandlerFunc(a.home))
	mux.Get("/snippet/create", http.HandlerFunc(a.createSnippet))
	mux.Post("/snippet/create", http.HandlerFunc(a.createSnippet))
	mux.Get("/snippet/:id", http.HandlerFunc(a.showSnippet))

	// Static Files
	fileServer := http.FileServer(http.Dir("./web/static/"))
	mux.Get("/static/", http.StripPrefix("/static", fileServer))

	// Middleware chain
	mw := alice.New(a.recoverPanic, a.logRequest, secureHeaders)

	// secureHeaders -> serveMux -> ... ↩
	return mw.Then(mux)
}
