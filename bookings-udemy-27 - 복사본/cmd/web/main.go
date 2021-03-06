package main

import (
	"GO/trevor/bookings-udemy-27/pkg/config"
	"GO/trevor/bookings-udemy-27/pkg/handlers"
	"GO/trevor/bookings-udemy-27/pkg/render"
	"fmt"
	"log"
	"net/http"
)

const portNumber = ":8080"

// main is the main function
func main() {
	var app config.AppConfig

	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("cannot create template cache")
	}

	app.TemplateCache = tc
	app.UseCache = false
	log.Println("main > app.UseCache")

	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)

	render.NewTemplates(&app)

	// http.HandleFunc("/", repo.Home)
	// http.HandleFunc("/about", repo.About)
	http.HandleFunc("/", handlers.Repo.Home)
	http.HandleFunc("/about", handlers.Repo.About)

	fmt.Println(fmt.Sprintf("Staring application on port %s", portNumber))
	_ = http.ListenAndServe(portNumber, nil)
}
