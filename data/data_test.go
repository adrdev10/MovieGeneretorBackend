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
	mc, err := movieData.FetchMovies(url)
	if err != nil {
		t.Errorf("FetchMovies failed. Expected %v, got %v. Culprit: %v", "Data", mc, err)
	} else {
		t.Logf("FetchMovies success. Expected %v, got %v. Culprit: %v", "Data", mc, err)
	}

}

func TestPopFunc(t *testing.T) {
	movieData := MovieData{}
	api, err := util.GetAPIKey()
	if err != nil {
		t.Error("Failed getting API key")
	}
	url := "http://www.omdbapi.com/?apikey=" + api + "&" + "s=action"
	mc, err := movieData.FetchMovies(url)
	if err != nil {
		t.Errorf("FetchMovies failed. Expected %v, got %v. Culprit: %v", "Data", mc, err)
	} else {
		fmt.Printf("Length: %v\n", len(mc.Movies))
		sm := mc.PopMovie()
		fmt.Printf("Poppped movie: %v, movie list length %v\n", sm, len(mc.Movies))
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
	mc, err := movieData.FetchMovies(url)
	if err != nil {
		t.Errorf("FetchMovies failed. Expected %v, got %v. Culprit: %v", "Data", mc, err)
	} else {
		for len(mc.Movies) != 0 {
			sm := mc.PopMovie()
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
