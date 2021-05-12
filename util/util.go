package util

import (
	"errors"
	"os"
)

func GetAPIKey() (string, error) {
	api, ok := os.LookupEnv("MOVIE_API")
	if ok {
		return api, nil
	}
	return "", errors.New("could not find API key")
}
