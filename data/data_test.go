package data

import (
	"fmt"
	"os"
	"testing"

	"github.com/adrdev10/movie-deliver/util"
)

func TestData(t *testing.T) {
	movieData := MovieData{}
	api, err := util.GetAPIKey()
	if err != nil {
		t.Error("Failed getting API key")
	}
	url := "http://www.omdbapi.com/?apikey=" + api + "&" + "s=action"
	err = movieData.FetchMovies(url)
	if err != nil {
		t.Errorf("FetchMovies failed. Expected %v, got %v. Culprit: %v", "Data", movieData, err)
	} else {
		t.Logf("FetchMovies success. Expected %v, got %v. Culprit: %v", "Data", movieData, err)
	}

}

func TestPopFunc(t *testing.T) {
	movieData := MovieData{}
	api, err := util.GetAPIKey()
	if err != nil {
		t.Error("Failed getting API key")
	}
	url := "http://www.omdbapi.com/?apikey=" + api + "&" + "s=action"
	err = movieData.FetchMovies(url)
	if err != nil {
		t.Errorf("FetchMovies failed. Expected %v, got %v. Culprit: %v", "Data", movieData, err)
	} else {
		fmt.Printf("Length: %v\n", len(movieData.Movies))
		sm := movieData.PopMovie()
		fmt.Printf("Poppped movie: %v, movie list length %v\n", sm, len(movieData.Movies))
		t.Logf("Popped movie was %v", sm)
	}

}

func TestMovieInfo(t *testing.T) {
	movieData := MovieData{}
	api, ok := os.LookupEnv("MOVIE_API")
	if !ok {
		t.Error("Could not get env key")
	}
	url := "http://www.omdbapi.com/?apikey=" + api + "&" + "s=action"
	err := movieData.FetchMovies(url)
	if err != nil {
		t.Errorf("FetchMovies failed. Expected %v, got %v. Culprit: %v", "Data", movieData, err)
	} else {
		for len(movieData.Movies) != 0 {
			sm := movieData.PopMovie()
			idSm, err := sm.GetImdbID()
			if err != nil {
				t.Errorf("GetImdbID failed. Expected %v, got %v", "#1234", err)
			} else {
				movieInfo, err := GetMovieInfo(idSm)
				if err != nil {
					t.Errorf("GetMovieInfo failed. %v", err)
				}
				t.Logf("Passed. Movie info obtained: %v", movieInfo)
			}
		}
	}
}
