package services

import (
	"context"

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
