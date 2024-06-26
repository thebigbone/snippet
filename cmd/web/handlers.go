package main

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"
)

func (app *applicaton) home(w http.ResponseWriter, r *http.Request) {
	files := []string{
		"/home/syscall/snippet/ui/html/base.html",
		"/home/syscall/snippet/ui/html/pages/home.html",
		"/home/syscall/snippet/ui/html/partials/nav.html",
	}

	ts, err := template.ParseFiles(files...)

	if err != nil {
		app.errorLog.Println(err.Error())

		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	err = ts.ExecuteTemplate(w, "base", nil)
	if err != nil {
		app.errorLog.Println(err.Error())
		http.Error(w, "Internal server error", http.StatusInternalServerError)
	}

}
func (app *applicaton) snippetView(w http.ResponseWriter, r *http.Request) {
	get_id := r.URL.Query().Get("id")

	id, err := strconv.Atoi(get_id)
	if err != nil || id < 1 {
		http.NotFound(w, r)
		return
	}

	fmt.Fprintf(w, "display a specific snippet with ID %d...", id)
}
func (app *applicaton) snippetCreate(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.Header().Set("Allow", http.MethodPost)
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	w.Write([]byte("create a new snippet..."))
}
