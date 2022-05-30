package main

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"

	"github.com/wralith/wrabox/pkg/models"
)

func (a *app) home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		a.notFound(w)
		return
	}

	s, err := a.snippets.Latest()
	if err != nil {
		a.serverError(w, err)
		return
	}

	data := &templateData{Snippets: s}

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

	err = ts.Execute(w, data) // Different err, SCOPE
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

	s, err := a.snippets.Get(id)
	if err == models.ErrNoRecord {
		a.notFound(w)
		return
	}
	if err != nil {
		a.serverError(w, err)
		return
	}

	// Contained in template data
	data := &templateData{Snippet: s}

	files := []string{
		"./web/template/show.page.html",
		"./web/template/base.layout.html",
		"./web/template/footer.partial.html",
	}

	ts, err := template.ParseFiles(files...)
	if err != nil {
		a.serverError(w, err)
		return
	}

	err = ts.Execute(w, data)
	if err != nil {
		a.serverError(w, err)
	}
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
