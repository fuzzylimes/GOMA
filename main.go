package main

import (
	"log"

	"github.com/fuzzylimes/goma/pkg/routes"
	"github.com/fuzzylimes/goma/pkg/templateloader"
)

func main() {
	// Load templates to memory
	templateloader.GetTemplates()

	// Build the router
	srv := routes.BuildRoutes()

	// Serve the router
	log.Fatalln(srv.ListenAndServe())
}
