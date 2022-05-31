package main

import (
	"net/http"

	"github.com/bmizerany/pat"
	"github.com/justinas/alice"
)

func (a *app) routes() http.Handler {

	// mux := http.NewServeMux() // If not declared -> DefaultServerMux
	mux := pat.New()
	// Middleware chain
	standardMw := alice.New(a.recoverPanic, a.logRequest, secureHeaders)

	// Session
	dynamicMw := alice.New(a.session.Enable)

	// Routes
	mux.Get("/", dynamicMw.ThenFunc(http.HandlerFunc(a.home)))
	mux.Get("/snippet/create", dynamicMw.ThenFunc(http.HandlerFunc(a.createSnippetForm)))
	mux.Post("/snippet/create", dynamicMw.ThenFunc(http.HandlerFunc(a.createSnippet)))
	mux.Get("/snippet/:id", dynamicMw.ThenFunc(http.HandlerFunc(a.showSnippet)))

	// Static Files
	fileServer := http.FileServer(http.Dir("./web/static/"))
	mux.Get("/static/", http.StripPrefix("/static", fileServer))

	// secureHeaders -> serveMux -> ... ↩
	return standardMw.Then(mux)
}
