package services

import (
	"context"
	"net/http"

	"github.com/DuvanAlbarracin/movies_movies/pkg/db"
	"github.com/DuvanAlbarracin/movies_movies/pkg/proto"
	"github.com/DuvanAlbarracin/movies_movies/pkg/utils"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) GetGenreById(ctx context.Context, req *proto.GetGenreByIdRequest) (*proto.GetGenreByIdResponse, error) {
	genre, err := db.FindGenreById(s.H.Conn, req.Id)
	if err != nil {
		var code codes.Code
		var message string
		if err.Error() == "no rows in result set" {
			code = codes.NotFound
			message = "Genre not found"
		} else {
			code = codes.Internal
			message = "There was an error finding the genre"
		}

		return nil, status.New(code, message).Err()
	}

	gp := utils.SetProtoGenre(genre)

	return &proto.GetGenreByIdResponse{
		Genre: gp,
	}, nil
}

func (s *Server) GetAllGenres(ctx context.Context, req *proto.GetAllGenresRequest) (*proto.GetAllGenresResponse, error) {
	var protoGenres []*proto.Genre

	genres, err := db.GetAllGenres(s.H.Conn)
	if err != nil {
		st := status.New(codes.Internal, "There was an error getting all genres: "+err.Error())
		return nil, st.Err()
	}

	if len(genres) < 1 {
		st := status.New(codes.NotFound, "There is no genres")
		return nil, st.Err()
	}

	for _, genre := range genres {
		protoGenre := utils.SetProtoGenre(genre)
		protoGenres = append(protoGenres, protoGenre)
	}

	return &proto.GetAllGenresResponse{
		Genres: protoGenres,
	}, nil
}

func (s *Server) AddGenderToMovie(ctx context.Context, req *proto.AddGenderToMovieRequest) (*proto.AddGenderToMovieResponse, error) {
	if _, err := db.FindMovieById(s.H.Conn, req.MovieId); err != nil {
		return nil, status.New(codes.NotFound, "There is no movie with that id").Err()
	}

	if _, err := db.FindGenreById(s.H.Conn, req.GenreId); err != nil {
		return nil, status.New(codes.NotFound, "There is no genre with that id").Err()
	}

	if err := db.AddGenderToMovie(s.H.Conn, req.GenreId, req.MovieId); err != nil {
		return nil, status.New(codes.Internal, "There was an error creating the join record").Err()
	}

	return &proto.AddGenderToMovieResponse{
		Status: http.StatusOK,
	}, nil
}

func (s *Server) RemoveGenderFromMovie(ctx context.Context, req *proto.RemoveGenderFromMovieRequest) (*proto.RemoveGenderFromMovieResponse, error) {
	if _, err := db.FindMovieById(s.H.Conn, req.MovieId); err != nil {
		return nil, status.New(codes.NotFound, "There is no movie with that id").Err()
	}

	if _, err := db.FindGenreById(s.H.Conn, req.GenreId); err != nil {
		return nil, status.New(codes.NotFound, "There is no genre with that id").Err()
	}

	if err := db.RemoveGenderFromMovie(s.H.Conn, req.GenreId, req.MovieId); err != nil {
		return nil, status.New(codes.Internal, "There was an error deleting from join record").Err()
	}

	return &proto.RemoveGenderFromMovieResponse{
		Status: http.StatusOK,
	}, nil
}
