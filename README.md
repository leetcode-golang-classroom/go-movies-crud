# go-movies-crud

golang crud service with [go mux](github.com/gorilla/mux)

## pre-request

1. install [go mux](github.com/gorilla/mux)

## structure

![](https://i.imgur.com/Sor5tcT.png)

## data model

```golang
package main

type Movie struct {
	ID       string    `json:"id"`
	Isbn     string    `json:"isbn"`
	Title    string    `json:"title"`
	Director *Director `json:"director"`
}

type Director struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

var movies []Movie
```
## handler logic

```golang
package main

import (
	"encoding/json"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func getMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(movies)
}

func deleteMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	// GET request id
	params := mux.Vars(r)
	var movie Movie
	for idx, item := range movies {
		if item.ID == params["id"] {
			movie = item
			movies = append(movies[:idx], movies[idx+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(movie)
}

func getMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	// GET request id
	params := mux.Vars(r)
	for _, item := range movies {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
}

func createMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var movie Movie
	// decode body into movie
	_ = json.NewDecoder(r.Body).Decode(&movie)
	movie.ID = strconv.Itoa(rand.Intn(100000000))
	movies = append(movies, movie)
	json.NewEncoder(w).Encode(movie)
}

func updateMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	// GET request id
	params := mux.Vars(r)
	var movie Movie
	for idx, item := range movies {
		if item.ID == params["id"] {
			movies = append(movies[:idx], movies[idx+1:]...)
			// decode body into movie
			_ = json.NewDecoder(r.Body).Decode(&movie)
			movie.ID = item.ID
			break
		}
	}
	movies = append(movies, movie)
	json.NewEncoder(w).Encode(movie)
}

```

## server listen logic

```golang
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
```