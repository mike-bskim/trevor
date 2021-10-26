package handlers

import (
	// "GO/trevor/bookings-udemy-27/pkg/render"
	// "bookings-udemy/pkg/render"
	"GO/trevor/bookings-udemy-27/pkg/render"
	"net/http"
)

// Home is the handler for the home page
func Home(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "home.page.tmpl")
}

// About is the handler for the about page
func About(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "about.page.tmpl")
}
