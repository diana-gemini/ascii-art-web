package internal

import (
	"net/http"
	"text/template"
)

func MainPage(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		errorHandler(w, r, http.StatusNotFound)
		return
	}

	switch r.Method {
	case http.MethodGet:
		{
			tmpl, err := template.ParseFiles("templates/mainPage.html")
			if err != nil {
				errorHandler(w, r, http.StatusNotFound)
				return
			}

			err = tmpl.Execute(w, nil)
			if err != nil {
				errorHandler(w, r, http.StatusInternalServerError)
				return
			}
		}
	default:
		{
			errorHandler(w, r, http.StatusMethodNotAllowed)
		}
	}
}
