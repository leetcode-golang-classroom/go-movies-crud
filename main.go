package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	movies = append(movies, Movie{
		ID:    "1",
		Isbn:  "418227",
		Title: "Movie One",
		Director: &Director{
			Firstname: "John",
			Lastname:  "Doe",
		},
	})
	movies = append(movies, Movie{
		ID:    "2",
		Isbn:  "45455",
		Title: "Movie Two",
		Director: &Director{
			Firstname: "Steve",
			Lastname:  "Smith",
		},
	})
	router.HandleFunc("/movies", getMovies).Methods("GET")
	router.HandleFunc("/movies/{id}", getMovie).Methods("GET")
	router.HandleFunc("/movies", createMovie).Methods("POST")
	router.HandleFunc("/movies/{id}", updateMovie).Methods("PUT")
	router.HandleFunc("/movies/{id}", deleteMovie).Methods("DELETE")
	fmt.Println("movie-crud-service start at 8000")
	log.Fatal(http.ListenAndServe(":8000", router))
}
