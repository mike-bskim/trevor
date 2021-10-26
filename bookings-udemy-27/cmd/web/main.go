package main

import (
	"GO/trevor/bookings-udemy-27/pkg/handlers"
	// "GO/trevor/bookings-udemy-27/pkg/handlers"
	// "bookings-udemy/pkg/handlers"
	"fmt"
	"net/http"
)

const portNumber = ":8080"

// main is the main function
func main() {
	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)

	fmt.Println(fmt.Sprintf("Staring application on port %s", portNumber))
	_ = http.ListenAndServe(portNumber, nil)
}
