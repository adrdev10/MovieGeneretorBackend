package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/adrdev10/movie-deliver/data"
)

var movieData *data.MovieData

func getMovieListRequest(w http.ResponseWriter, r *http.Request) {
	keys, ok := r.URL.Query()["genre"]
	if !ok || len(keys[0]) < 1 {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"message": "Url Param 'genre' is missing"}`))
		return
	}
	err := movieData.FetchMovies(keys[0])
	if err != nil {
		w.Write([]byte(`{"message": "Error when fetching movies"}`))
		return
	}
	json.NewEncoder(w).Encode(movieData)
}

func getMovieInfo(w http.ResponseWriter, r *http.Request) {
	if len(movieData.Movies) < 1 {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"message": "No movies found in the list"}`))
		return
	}
	sm := movieData.PopMovie()
	idSm, err := sm.GetImdbID()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"message": "Could not get imdbid from selected movie"}`))
		return
	} else {
		movieInfo, err := data.GetMovieInfo(idSm)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(`{"message": "Movie info could not be reached"}`))
			return
		}
		json.NewEncoder(w).Encode(movieInfo)
	}
}

func main() {
	movieData = &data.MovieData{}
	http.HandleFunc("/", getMovieListRequest)
	http.HandleFunc("/movieInfo", getMovieInfo)
	fmt.Println("Server starting on port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
