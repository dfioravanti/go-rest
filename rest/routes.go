package rest

import "net/http"

// The routes() method returns a servemux containing our application routes.
func (app Application) Routes() *http.ServeMux {
	mux := http.NewServeMux()

	fileServer := http.FileServer(http.Dir("./ui/static/"))
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	mux.HandleFunc("/snippet/view", app.SnippetHandler.View)
	mux.HandleFunc("/snippet/create", app.SnippetHandler.Create)

	return mux
}
