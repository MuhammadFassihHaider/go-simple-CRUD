package main

import (
	"encoding/json"
	"log"
	"math/rand"
	"net/http"
	"time"

	"fmt"

	"github.com/gorilla/mux"
	// "math/rand"

	"strconv"
)

type AddMovieBody struct {
	ISBN     string    `json:"isbn"`
	Title    string    `json:"title"`
	Director *Director `json:"director"`
}

type Movie struct {
	Id       string    `json:"id"`
	ISBN     string    `json:"isbn"`
	Title    string    `json:"title"`
	Director *Director `json:"director"`
}

type Director struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

var movies []Movie

type Error struct {
	Message string
}

type DeleteMovieByIdResponse struct {
	Movies []Movie
	Movie  Movie
	Index  int
}

type AddMovieResponse struct {
	Movies []Movie
}

func catchErrors() {
	if err := recover(); err != nil {
		log.Println(err)
	}
}

func getMovies(w http.ResponseWriter, r *http.Request) {
	defer catchErrors()
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(movies)
}

func getMovieById(w http.ResponseWriter, r *http.Request) {
	defer catchErrors()
	w.Header().Set("Content-Type", "application/json")
	id := mux.Vars(r)["id"]

	var movie Movie
	for _, m := range movies {
		if m.Id == id {
			fmt.Println("if m.Id == id", m)
			movie = m
			break
		}
	}

	if movie == (Movie{}) {
		fmt.Println("error")
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(Error{Message: "Not found!"})

	} else {
		fmt.Println("response")
		json.NewEncoder(w).Encode(movie)
	}
}

func deleteMovieById(w http.ResponseWriter, r *http.Request) {
	defer catchErrors()

	w.Header().Set("Content-Type", "application/json")

	id := mux.Vars(r)["id"]
	idAsInteger, err := strconv.ParseUint(id, 10, 32)

	if err != nil || idAsInteger <= 0 {
		w.WriteHeader(400)
		json.NewEncoder(w).Encode(Error{Message: "Invalid Id!"})
		return
	}
	var filteredMovies []Movie
	index := -1
	for i, m := range movies {
		if m.Id == id {
			index = i
		} else {
			filteredMovies = append(filteredMovies, m)
		}
	}

	if index == -1 {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(Error{Message: "Not found!"})
		return
	} else {
		movie := movies[index]
		movies = filteredMovies
		json.NewEncoder(w).Encode(DeleteMovieByIdResponse{
			Movies: filteredMovies,
			Movie:  movie,
			Index:  index,
		})
	}
}

func getRandomId() string {
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
	randomId := r1.Intn(100)
	randomIdAsString := strconv.Itoa(randomId)
	return randomIdAsString
}

func addMovie(w http.ResponseWriter, r *http.Request) {
	defer catchErrors()
	w.Header().Set("Content-Type:", "application/json")

	var movie AddMovieBody

	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()
	err := decoder.Decode(&movie)
	if err != nil {
		w.WriteHeader(400)
		json.NewEncoder(w).Encode(Error{Message: "Invalid arguments!"})
		return
	}

	isbn := movie.ISBN
	title := movie.Title
	directorFirstName := movie.Director.Firstname
	directorLastname := movie.Director.Lastname

	if title == "" || directorFirstName == "" || directorLastname == "" {
		w.WriteHeader(400)
		json.NewEncoder(w).Encode(Error{Message: "Invalid arguments!"})
		return
	} else {
		Id := getRandomId()
		movies = append(movies, Movie{
			Id:    Id,
			ISBN:  isbn,
			Title: title,
			Director: &Director{
				Firstname: directorFirstName,
				Lastname:  directorLastname,
			},
		})
		json.NewEncoder(w).Encode(AddMovieResponse{Movies: movies})
	}

}

func main() {
	router := mux.NewRouter()

	movies = addDummyData(movies)

	router.HandleFunc("/movies", getMovies).Methods("GET")
	router.HandleFunc("/movies/{id}", getMovieById).Methods("GET")
	router.HandleFunc("/movies/{id}", deleteMovieById).Methods("DELETE")
	router.HandleFunc("/movies", addMovie).Methods("POST")

	fmt.Println("Starting server on port 8000")
	log.Fatal(http.ListenAndServe(":8000", router))
}
