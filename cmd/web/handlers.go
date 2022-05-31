package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/wralith/wrabox/pkg/models"
)

func (a *app) home(w http.ResponseWriter, r *http.Request) {

	s, err := a.snippets.Latest()
	if err != nil {
		a.serverError(w, err)
		return
	}

	a.render(w, r, "home.page.html", &templateData{
		Snippets: s,
	})

}

func (a *app) showSnippet(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get(":id"))
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

	a.render(w, r, "show.page.html", &templateData{
		Snippet: s,
	})
}

func (a *app) createSnippetForm(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Create snippet"))
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
