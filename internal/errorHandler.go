package internal

import (
	"net/http"
	"text/template"
)

func errorHandler(w http.ResponseWriter, r *http.Request, status int) {
	e := Err{
		StatusCode: status,
		StatusText: http.StatusText(status),
	}
	w.WriteHeader(status)

	tmpl, err := template.ParseFiles("templates/error.html")
	if err != nil {
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	err = tmpl.Execute(w, e)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
}
