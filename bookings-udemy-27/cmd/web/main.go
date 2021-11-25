package main

import (
	"GO/trevor/bookings-udemy-27/pkg/handlers"
	"fmt"
	"net/http"
	// "GO/trevor/bookings-udemy-27/pkg/"
)

const portNumber = ":3000"

// main is the main function
func main() {

	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)
	// http.HandleFunc("/", handlers.Repo.Home)
	// http.HandleFunc("/about", handlers.Repo.About)
	tmp := fmt.Sprintf("Staring application on port %s", portNumber)
	fmt.Println(tmp)
	_ = http.ListenAndServe(portNumber, nil)
}
