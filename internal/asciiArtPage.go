package internal

import (
	"net/http"
	"text/template"
)

func AsciiArtPage(w http.ResponseWriter, r *http.Request) {
	var data Ascii
	var err error

	switch r.Method {
	case http.MethodPost:
		{
			data.BannerName = r.FormValue("bannerName")
			data.InputText = r.FormValue("inputText")

			err = CheckInputText(&data)
			if err != nil {
				errorHandler(w, r, http.StatusBadRequest)
				return
			}

			err = ReadBannerFile(&data)
			if err != nil {
				errorHandler(w, r, http.StatusNotFound)
				return
			}

			MakeMap(&data)
			CollectOutputText(&data)

			tmpl, err := template.ParseFiles("templates/mainPage.html")
			if err != nil {
				errorHandler(w, r, http.StatusNotFound)
				return
			}

			err = tmpl.Execute(w, data)
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
