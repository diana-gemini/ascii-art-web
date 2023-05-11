package main

import (
	f "ascii-art-web/internal"
	"fmt"
	"log"
	"net/http"
	"time"
)

func main() {
	var err error

	mux := http.NewServeMux()

	mux.HandleFunc("/", f.MainPage)
	mux.HandleFunc("/ascii-art", f.AsciiArtPage)

	mux.Handle("/static/", http.StripPrefix("/static", http.FileServer(http.Dir("./ui/static/"))))

	srv := &http.Server{
		Addr:         ":8000",
		Handler:      mux,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 5 * time.Second,
	}

	log.Print("http://localhost" + srv.Addr)

	err = srv.ListenAndServe()
	if err != nil {
		fmt.Println(http.StatusText(http.StatusInternalServerError))
	}
}
