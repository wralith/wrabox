package main

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"
)

func (a *app) home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		a.notFound(w)
		return
	}

	files := []string{
		"./web/template/home.page.html",
		"./web/template/base.layout.html",
		"./web/template/footer.partial.html",
	}

	ts, err := template.ParseFiles(files...)
	if err != nil {
		a.serverError(w, err)
		return
	}

	err = ts.Execute(w, nil) // Different err, SCOPE
	if err != nil {
		a.serverError(w, err)
	}
}

func (a *app) showSnippet(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id < 1 {
		a.notFound(w)
		return
	}
	fmt.Fprintf(w, "A snippet with ID %d...", id)
}

func (a *app) createSnippet(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		w.Header().Set("Allow", "POST")
		a.clientError(w, http.StatusMethodNotAllowed)
		return
	}

	// Dummy Data to add
	// TODO how to mock mysql to write test?
	title := "I am dummy"
	content := "I don't feel like i am motivated enough to...\n populate dummy boy."
	expires := "7"

	id, err := a.snippets.Insert(title, content, expires)
	if err != nil {
		a.serverError(w, err)
		return
	}

	http.Redirect(w, r, fmt.Sprintf("/snippet?id=%d", id), http.StatusSeeOther)

}
