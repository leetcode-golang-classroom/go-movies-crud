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
