package usecases

import "github.com/adrdev10/movie-deliver/movie"

type movieUseCase interface {
	FetchMovies() []movie.Movie
}
