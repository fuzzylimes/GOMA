package routes

import (
	"net/http"
	"time"

	"github.com/fuzzylimes/goma/pkg/handlers"
	"github.com/gorilla/mux"
)

func BuildRoutes() *http.Server {
	router := mux.NewRouter()
	router.HandleFunc("/", handlers.HomeHandler)
	router.HandleFunc("/about", handlers.AboutHandler)

	// handle static files
	router.PathPrefix("/css/").Handler(http.StripPrefix("/css/", http.FileServer(http.Dir("./web/public/css"))))
	router.PathPrefix("/js/").Handler(http.StripPrefix("/js/", http.FileServer(http.Dir("./web/public/js"))))
	router.PathPrefix("/img/").Handler(http.StripPrefix("/img/", http.FileServer(http.Dir("./web/public/img"))))

	srv := &http.Server{
		Addr: ":8000",
		// Good practice to set timeouts to avoid Slowloris attacks.
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
		Handler:      router,
	}

	return srv
}
