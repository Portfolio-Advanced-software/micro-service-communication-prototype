package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/api/movies", getAllMovies).Methods("GET")
	router.HandleFunc("/api/movies/{id}", getMovieByID).Methods("GET")
	router.HandleFunc("/api/movies", createMovie).Methods("POST")
	router.HandleFunc("/api/movies/{id}", updateMovie).Methods("PUT")
	router.HandleFunc("/api/movies/{id}", deleteMovie).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8000", router))
}

func getAllMovies(w http.ResponseWriter, r *http.Request) {
	// Hard-coded response
	response := "This is the response for GET all movies"

	// Set the response content type
	w.Header().Set("Content-Type", "text/plain")

	// Write the response
	fmt.Fprint(w, response)
}

func getMovieByID(w http.ResponseWriter, r *http.Request) {
	// Hard-coded response
	response := "This is the response for GET movie by ID"

	// Set the response content type
	w.Header().Set("Content-Type", "text/plain")

	// Write the response
	fmt.Fprint(w, response)
}

func createMovie(w http.ResponseWriter, r *http.Request) {
	// Hard-coded response
	response := "This is the response for creating a movie"

	// Set the response content type
	w.Header().Set("Content-Type", "text/plain")

	// Write the response
	fmt.Fprint(w, response)
}

func updateMovie(w http.ResponseWriter, r *http.Request) {
	// Hard-coded response
	response := "This is the response for updating a movie"

	// Set the response content type
	w.Header().Set("Content-Type", "text/plain")

	// Write the response
	fmt.Fprint(w, response)
}

func deleteMovie(w http.ResponseWriter, r *http.Request) {
	// Hard-coded response
	response := "This is the response for deleting a movie"

	// Set the response content type
	w.Header().Set("Content-Type", "text/plain")

	// Write the response
	fmt.Fprint(w, response)
}
