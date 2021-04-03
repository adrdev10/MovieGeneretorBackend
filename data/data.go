package data

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"

	http "net/http"

	"github.com/adrdev10/movie-deliver/movie"
)

type MovieDataInterface interface {
	FetchMovies(url string) (movie.Movies, error)
}

//MovieData represents a
type MovieData struct {
	Movies []*movie.Movie `json:"Search"`
}

func (MovieData) FetchMovies(url string) (*MovieData, error) {

	mc, client := &MovieData{}, http.Client{}
	resp, err := client.Get(url)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	err = json.NewDecoder(resp.Body).Decode(mc)
	if err != nil {
		return nil, errors.New("error: could not decode data into the structure")
	}

	mcJson, err := json.Marshal(mc)
	if err != nil {
		return nil, errors.New("error: could not decode the body into the structure")
	}
	fmt.Println(string(mcJson))

	return mc, nil
}

//PopMovie removes the last movie in the array
func (mo *MovieData) PopMovie() *movie.Movie {
	popped := mo.Movies[len(mo.Movies)-1]
	mo.Movies = mo.Movies[:len(mo.Movies)-1]
	return popped
}

//GetAllMovieNames collects and return all the names from the movie collection
func (mo *MovieData) GetAllMovieNames() ([]string, error) {
	moviesNames := []string{}
	for _, movie := range mo.Movies {
		name, err := movie.GetTitle()
		if err != nil {
			return nil, err
		}
		moviesNames = append(moviesNames, name)
	}
	return moviesNames, nil
}

func (mo *MovieData) filterOnlyMovies() *MovieData {
	for i, movie := range mo.Movies {
		//type can only be series or movie
		if movie.Type == "series" {
			mo.Movies[i], mo.Movies[len(mo.Movies)] = mo.Movies[len(mo.Movies)-1], nil
		}
	}
	return mo
}

func GetMovieInfo(movieID string) (*movie.MovieInfo, error) {
	client := http.Client{}
	api, ok := os.LookupEnv("MOVIE_API")
	if !ok {
		return nil, errors.New("ENV key not found")
	}
	url := "http://www.omdbapi.com/?apikey=" + api + "&" + "i=" + movieID
	resp, err := client.Get(url)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	mc := &movie.MovieInfo{}

	err = json.NewDecoder(resp.Body).Decode(mc)
	if err != nil {
		return nil, errors.New("error: could not decode data into the structure")
	}

	mcJson, err := json.Marshal(mc)
	if err != nil {
		return nil, errors.New("error: could not decode the body into the structure")
	}
	fmt.Println(string(mcJson))

	return mc, nil

}
