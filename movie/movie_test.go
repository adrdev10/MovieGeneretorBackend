package movie

import (
	"testing"
)

//TestMovie test the creation of a new movie object
func TestMovie(t *testing.T) {
	mo := CreateNewMovie("American Horror Story", "t1844624", "series", "poster")

	if des, err := mo.GetTitle(); err != nil && des != "American Horror Story" {
		t.Errorf("CreateNewMovieFunction() failed, expected %v, got %v", "description", des)
	}
}
