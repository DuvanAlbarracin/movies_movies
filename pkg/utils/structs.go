package utils

import (
	"reflect"

	"github.com/DuvanAlbarracin/movies_movies/pkg/models"
	genreP "github.com/DuvanAlbarracin/movies_movies/pkg/proto/genre"
	movieP "github.com/DuvanAlbarracin/movies_movies/pkg/proto/movie"
)

func SetModifyMap(s any) (map[string]any, []any) {
	resMap := make(map[string]any)
	var args []any
	ref := reflect.ValueOf(s)

	for i := 0; i < ref.NumField(); i++ {
		f := ref.Field(i)
		if f.IsZero() {
			continue
		}

		if f.Kind() == reflect.Pointer {
			f = f.Elem()
		}

		v := ref.Type().Field(i)
		fValue := f.Interface()
		resMap[ToSnakeCase(v.Name)] = fValue
		args = append(args, fValue)
	}

	return resMap, args
}

func SetProtoMovie(movie models.Movie) *movieP.Movie {
	mp := movieP.Movie{
		Id:          movie.Id,
		Title:       movie.Title,
		ReleaseYear: movie.ReleaseYear,
		DirectorID:  movie.DirectorID,
		MusicBy:     movie.MusicBy,
		WrittenBy:   movie.WrittenBy,
	}

	return &mp
}

func SetProtoGenre(genre models.Genre) *genreP.Genre {
	gp := genreP.Genre{
		Id:   genre.Id,
		Name: genre.Name,
	}

	return &gp
}
