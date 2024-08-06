package utils

import (
	"reflect"

	"github.com/DuvanAlbarracin/movies_movies/pkg/models"
	"github.com/DuvanAlbarracin/movies_movies/pkg/proto"
)

func SetModifyMap(s any) (map[string]any, []any) {
	resMap := make(map[string]any)
	var args []any
	ref := reflect.ValueOf(s)

	for i := 0; i < ref.NumField(); i++ {
		f := ref.Field(i)
		if !f.IsZero() {
			v := ref.Type().Field(i)
			resMap[v.Name] = f.Interface()
			args = append(args, f.Interface())
		}
	}

	return resMap, args
}

func SetProtoMovie(movie models.Movie) *proto.Movie {
	mp := proto.Movie{
		Id:          movie.Id,
		Title:       movie.Title,
		ReleaseYear: movie.ReleaseYear,
		DirectorID:  movie.DirectorID,
		MusicBy:     movie.MusicBy,
		WrittenBy:   movie.WrittenBy,
	}

	return &mp
}

func SetProtoGenre(genre models.Genre) *proto.Genre {
	mp := proto.Genre{
		Id:   genre.Id,
		Name: genre.Name,
	}

	return &mp
}
