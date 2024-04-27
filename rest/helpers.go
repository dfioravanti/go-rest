package rest

import (
	"encoding/json"
	"net/http"
	"runtime/debug"
)

type HTTPError struct {
	Title  string
	Status int
	Detail string
}

func (app Application) ReturnError(w http.ResponseWriter, err HTTPError) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Header().Set("X-Content-Type-Options", "nosniff")
	w.WriteHeader(err.Status)
	json.NewEncoder(w).Encode(err)
}

func (app Application) ServerError(w http.ResponseWriter, r *http.Request, err error) {
	var (
		method = r.Method
		uri    = r.URL.RequestURI()
		trace  = string(debug.Stack())
	)

	app.Logger.Error(err.Error(), "method", method, "uri", uri, "trace", trace)
	app.ReturnError(w, HTTPError{
		Status: http.StatusInternalServerError,
		Title:  "Internal server error",
		Detail: "Internal server error",
	})
}

func (app Application) notFound(w http.ResponseWriter) {
	app.ReturnError(w, HTTPError{
		Status: http.StatusNotFound,
		Title:  "Not found",
		Detail: "ID not found",
	})
}
