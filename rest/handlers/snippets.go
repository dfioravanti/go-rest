package handlers

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/dfioravanti/go-rest/services"
)

type SnippetHandler struct {
	snippetService services.SnippetService
}

func NewSnippetHandler(service services.SnippetService) SnippetHandler {
	return SnippetHandler{snippetService: service}
}

func (h SnippetHandler) View(w http.ResponseWriter, r *http.Request) {

	snippet, err := h.snippetService.Get(r.URL.Query().Get("id"))
	if err != nil {
		if errors.Is(err, services.ErrNoSnippetFound) {
			// add error
		} else {
			// add error
		}
		return
	}

	fmt.Fprintf(w, "%+v", snippet)
}

func (app *SnippetHandler) Create(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Create a new snippet..."))
}
