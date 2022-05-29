package main

import "net/http"

func (a *app) routes() *http.ServeMux {

	mux := http.NewServeMux() // If not declared -> DefaultServerMux

	// Routes
	mux.HandleFunc("/", a.home)
	mux.HandleFunc("/snippet", a.showSnippet)
	mux.HandleFunc("/snippet/create", a.createSnippet)

	// Static Files
	fileServer := http.FileServer(http.Dir("./web/static/"))
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	return mux
}
