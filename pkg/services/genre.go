package services

import (
	"context"

	"github.com/DuvanAlbarracin/movies_movies/pkg/db"
	proto "github.com/DuvanAlbarracin/movies_movies/pkg/proto/genre"
	"github.com/DuvanAlbarracin/movies_movies/pkg/utils"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type GenreServer struct {
	H db.Handler
	proto.UnimplementedGenreServiceServer
}

func (s *GenreServer) GetById(ctx context.Context, req *proto.GetByIdRequest) (*proto.GetByIdResponse, error) {
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

	return &proto.GetByIdResponse{
		Genre: gp,
	}, nil
}

func (s *GenreServer) GetAll(ctx context.Context, req *proto.GetAllRequest) (*proto.GetAllResponse, error) {
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

	return &proto.GetAllResponse{
		Genres: protoGenres,
	}, nil
}
