package main

import (
	"log"
	"net/http"

	"github.com/fuzzylimes/goma/pkg/handlers"
	"github.com/fuzzylimes/goma/pkg/templateloader"
)

func main() {
	templateloader.GetTemplates()
	http.HandleFunc("/", handlers.HomeHandler)
	http.HandleFunc("/about", handlers.AboutHandler)

	// handle static files
	http.Handle("/css/", http.StripPrefix("/css/", http.FileServer(http.Dir("./web/public/css"))))
	http.Handle("/js/", http.StripPrefix("/js/", http.FileServer(http.Dir("./web/public/js"))))
	http.Handle("/img/", http.StripPrefix("/img/", http.FileServer(http.Dir("./web/public/img"))))

	// serve files
	log.Fatalln(http.ListenAndServe(":8000", nil))
}
