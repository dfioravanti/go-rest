package main

import (
	"errors"
	"fmt"
	"net/http"
	"text/template"

	"github.com/dfioravanti/go-rest/internal/services"
)

func (app *Application) home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	// Include the navigation partial in the template files.
	files := []string{
		"./ui/html/base.tmpl",
		"./ui/html/partials/nav.tmpl",
		"./ui/html/pages/home.tmpl",
	}

	ts, err := template.ParseFiles(files...)
	if err != nil {
		app.ServerError(w, r, err)
		return
	}

	err = ts.ExecuteTemplate(w, "base", nil)
	if err != nil {
		app.ServerError(w, r, err)
	}
}

func (app *Application) snippetView(w http.ResponseWriter, r *http.Request) {

	snippet, err := app.service.Get(r.URL.Query().Get("id"))
	if err != nil {
		if errors.Is(err, services.ErrNoSnippetFound) {
			app.notFound(w)
		} else {
			app.ServerError(w, r, err)
		}
		return
	}

	fmt.Fprintf(w, "%+v", snippet)
}

func (app *Application) snippetCreate(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.Header().Set("Allow", http.MethodPost)
		app.ReturnError(w, HTTPError{Status: http.StatusMethodNotAllowed, Title: "Method not allowed", Detail: "Please use POST on this endpoint"})
		return
	}

	w.Write([]byte("Create a new snippet..."))
}
