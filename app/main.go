package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/adrdev10/movie-deliver/data"
)

var movieData data.MovieData

func handleRequest(w http.ResponseWriter, r *http.Request) {
	keys, ok := r.URL.Query()["genre"]
	if !ok || len(keys[0]) < 1 {
		log.Println("Url Param 'genre' is missing")
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"message": "Url Param 'genre' is missing"}`))
		return
	}
	api, ok := os.LookupEnv("MOVIE_API")
	if !ok {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Println("No API key found")
	}
	url := "http://www.omdbapi.com/?apikey=" + api + "&" + "s=" + keys[0]
	mc, err := movieData.FetchMovies(url)
	if err != nil {
		w.Write([]byte(`{"message": "Error when fetching movies"}`))
		return
	}
	json.NewEncoder(w).Encode(mc)
}

func main() {
	movieData = data.MovieData{}
	http.HandleFunc("/", handleRequest)
	fmt.Println("Server starting on port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
