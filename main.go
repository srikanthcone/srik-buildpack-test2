package main

import (
	"encoding/json"
	"log"
	"net/http"
	"fmt"
	"github.com/gorilla/mux"
)

type Movie struct {
	Id          string    `json:"id,omitempty"`
	Title       string    `json:"title,omitempty"`
	Description string    `json:"description,omitempty"`
	Genre       string    `json:"genre,omitempty"`
	Rating      string    `json:"rating,omitempty"`
	CastCrew    *CastCrew `json:"castCrew,omitempty"`
}

type CastCrew struct {
	Name string `json:"name,omitempty"`
	Role string `json:"role,omitempty"`
}

var movies []Movie

func defaultHandler(w http.ResponseWriter, r *http.Request) {
	title := "GO Rest API of Movie Resources\n Fetch All Movies : GET /movies\n Fetch Movie By Id : GET /movies/{id}\n Create Movie : POST /movies\n Update Movie : PUT /movies/{id}\n Delete Movie : DELETE /movies\n"
	fmt.Fprintf(w, "Hello from:  "+title+"\n")
}

func getMovieById(w http.ResponseWriter, req *http.Request) {
	w.Header().Add("content-type", "application/json")
	params := mux.Vars(req)
	for _, movie := range movies {
		if movie.Id == params["id"] {
			json.NewEncoder(w).Encode(movie)
			return
		}
	}
	json.NewEncoder(w).Encode(&Movie{})
}

func getMovies(w http.ResponseWriter, req *http.Request) {
	w.Header().Add("content-type", "application/json")
	json.NewEncoder(w).Encode(movies)
}

func addMovies(w http.ResponseWriter, req *http.Request) {
	w.Header().Add("content-type", "application/json")
	var movie Movie
	_ = json.NewDecoder(req.Body).Decode(&movie)
	movies = append(movies, movie)
	json.NewEncoder(w).Encode(movies)
}

func updateMovies(w http.ResponseWriter, req *http.Request) {
	w.Header().Add("content-type", "application/json")
	params := mux.Vars(req)
	for index, movie := range movies {
		if movie.Id == params["id"] {
			_ = json.NewDecoder(req.Body).Decode(&movie)
			movies = append(movies[:index], movie)
			break
		}
	}
	json.NewEncoder(w).Encode(movies)
}

func deleteMovies(w http.ResponseWriter, req *http.Request) {
	w.Header().Add("content-type", "application/json")
	params := mux.Vars(req)
	for index, movie := range movies {
		if movie.Id == params["id"] {
			movies = append(movies[:index], movies[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(movies)
}

func hello(user string) string {
	return fmt.Sprintf("Hello %v", user)
}

func hello2(user string) string {
	return fmt.Sprintf("Sorry %v", user)
}

func main() {
	router := mux.NewRouter()

	//perload some movies
	movies = append(movies, Movie{Id: "1", Title: "Spider Man Far From Home", Description: "Peter Parker's world has changed a lot since the events of Avengers: Endgame (2019)", Genre: "Action", Rating: "8", CastCrew: &CastCrew{Name: "Jon Watts", Role: "Director"}})

	router.HandleFunc("/", defaultHandler).Methods("GET")
	router.HandleFunc("/movies", getMovies).Methods("GET")
	router.HandleFunc("/movies/{id}", getMovieById).Methods("GET")
	router.HandleFunc("/movies", addMovies).Methods("POST")
	router.HandleFunc("/movies/{id}", updateMovies).Methods("PUT")
	router.HandleFunc("/movies/{id}", deleteMovies).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8080", router))
}
