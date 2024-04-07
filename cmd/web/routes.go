package main

import "net/http"

func (app *applicaton) routes() *http.ServeMux {
	mux := http.NewServeMux()

	// routes
	mux.HandleFunc("/home", app.home)
	mux.HandleFunc("/snippet/view", app.snippetView)
	mux.HandleFunc("/snippet/create", app.snippetCreate)

	return mux
}
