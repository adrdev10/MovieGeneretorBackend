package movie

import (
	"errors"
)

//Collection of moviess
type Movies struct {
	Search []Movie `json:"Search"`
}

//Movie object represents a movie to be sent to the UI
type Movie struct {
	Title  string `json:"Title"`
	ImdbID string `json:"imdbID"`
	Type   string `json:"Type"`
	Poster string `json:"Poster"`
}

type MovieInfo struct {
	Title    string `json:"Title"`
	Year     string `json:"Year"`
	Director string `json:"Director"`
	Plot     string `json:"Plot"`
	Released string `json:"Released"`
}

//CreateNewMovie creates a new movie object
func CreateNewMovie(title, imdbID, typeMS, poster string) *Movie {
	movie := &Movie{
		Title:  title,
		ImdbID: imdbID,
		Type:   typeMS,
		Poster: poster,
	}
	// movie.generateID()
	return movie
}

//Gettitle returns the title of the movie
func (m *Movie) GetTitle() (string, error) {
	if m.Title != "" {
		return m.Title, nil
	}
	return "", errors.New("title is not found")
}

//GetDirector returns the director of the movie
func (m *Movie) GetImdbID() (string, error) {
	if m.ImdbID != "" {
		return m.ImdbID, nil
	}
	return "", errors.New("ImdbID is not found")
}

//GetDescription returns the description of the movie
func (m *Movie) GetType() (string, error) {
	if m.Type != "" {
		return m.Type, nil
	}
	return "", errors.New("title is not found")
}

func (m Movie) GetPoster() (string, error) {
	if m.Poster != "" {
		return m.Poster, nil
	}
	return "", errors.New("Poster not found")
}
