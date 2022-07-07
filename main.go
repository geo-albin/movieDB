package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

var movies = []Movie{
	{
		ID:       1,
		Name:     "Movie 1",
		Director: "Director 1",
	},
	{
		ID:       2,
		Name:     "Movie 2",
		Director: "Director 2",
	},
	{
		ID:       3,
		Name:     "Movie 3",
		Director: "Director 3",
	},
	{
		ID:       4,
		Name:     "Movie 4",
		Director: "Director 4",
	},
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", getMovies).Methods("GET")
	r.HandleFunc("/movies", getMovies).Methods("GET")
	r.HandleFunc("/movies/{id}", getMovie).Methods("GET")
	r.HandleFunc("/movies", createMovie).Methods("POST")
	r.HandleFunc("/movies/{id}", updateMovie).Methods("PUT")
	r.HandleFunc("/movies/{id}", deleteMovie).Methods("DELETE")

	log.Println("Starting the server at port 8080")
	http.ListenAndServe(":8080", r)
}

func getMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-Type", "application/json")
	json.NewEncoder(w).Encode(movies)
}

func getMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-Type", "application/json")

	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, "Invalid id", http.StatusBadRequest)
		return
	}

	for i := range movies {
		if movies[i].ID == id {
			json.NewEncoder(w).Encode(movies[i])
			break
		}
	}
}

func deleteMovie(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, "Invalid id", http.StatusBadRequest)
		return
	}

	for i := range movies {
		if movies[i].ID == id {
			movies = append(movies[:i], movies[i+1:]...)
			break
		}
	}

	getMovies(w, r)
}

func createMovie(w http.ResponseWriter, r *http.Request) {
	var movie Movie
	json.NewDecoder(r.Body).Decode(&movie)
	id := 1

	if (len(movies)) != 0 {
		id = movies[len(movies)-1].ID + 1
	}

	movie.ID = id

	movies = append(movies, movie)

	getMovies(w, r)
}

func updateMovie(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, "Invalid id", http.StatusBadRequest)
		return
	}

	for i := range movies {
		if movies[i].ID == id {
			movies = append(movies[:i], movies[i+1:]...)
			break
		}
	}

	createMovie(w, r)
}
