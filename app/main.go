package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/adrdev10/movie-deliver/data"
	"github.com/adrdev10/movie-deliver/util"
)

var movieData *data.MovieData

func handleRequest(w http.ResponseWriter, r *http.Request) {
	keys, ok := r.URL.Query()["genre"]
	if !ok || len(keys[0]) < 1 {
		log.Println("Url Param 'genre' is missing")
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"message": "Url Param 'genre' is missing"}`))
		return
	}
	api, err := util.GetAPIKey()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"message": "Could not get API key"}`))
	}
	url := "http://www.omdbapi.com/?apikey=" + api + "&" + "s=" + keys[0]
	err = movieData.FetchMovies(url)
	if err != nil {
		w.Write([]byte(`{"message": "Error when fetching movies"}`))
		return
	}
	json.NewEncoder(w).Encode(movieData)
}

func main() {
	movieData = &data.MovieData{}
	http.HandleFunc("/", handleRequest)
	fmt.Println("Server starting on port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
