package main

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"unicode/utf8"

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

	a.render(w, r, "create.page.html", nil)
}

func (a *app) createSnippet(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		a.clientError(w, http.StatusBadRequest)
		return
	}

	title := r.PostForm.Get("title")
	content := r.PostForm.Get("content")
	expires := r.PostForm.Get("expires")

	errors := make(map[string]string)

	// Validations
	if strings.TrimSpace(title) == "" {
		errors["title"] = "This field cannot be blank"
		// Number of characters not bytes
	} else if utf8.RuneCountInString(title) > 100 {
		errors["title"] = "This field is too long (maximum is 100 characters)"
	}

	if strings.TrimSpace(content) == "" {
		errors["content"] = "This field cannot be blank"
	}

	if strings.TrimSpace(expires) == "" {
		errors["expires"] = "This field cannot be blank"
	} else if expires != "365" && expires != "7" && expires != "1" {
		errors["expires"] = "This field is invalid"
	}

	if len(errors) > 0 {
		fmt.Fprint(w, errors)
		return
	}

	id, err := a.snippets.Insert(title, content, expires)
	if err != nil {
		a.serverError(w, err)
		return
	}

	a.session.Put(r, "flash", "Snippet created successfully")
	http.Redirect(w, r, fmt.Sprintf("/snippet/%d", id), http.StatusSeeOther)
}

func (a *app) signUpUserForm(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "")
}

func (a *app) signUpUser(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "")
}

func (a *app) loginUserForm(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "")
}

func (a *app) loginUser(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "")
}

func (a *app) logoutUser(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "")
}
