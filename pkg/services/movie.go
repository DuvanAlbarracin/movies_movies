package services

import (
	"context"
	"net/http"

	"github.com/DuvanAlbarracin/movies_movies/pkg/db"
	"github.com/DuvanAlbarracin/movies_movies/pkg/models"
	proto "github.com/DuvanAlbarracin/movies_movies/pkg/proto/movie"
	"github.com/DuvanAlbarracin/movies_movies/pkg/utils"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type MovieServer struct {
	H db.Handler
	proto.UnimplementedMoviesServiceServer
}

func (s *MovieServer) Create(ctx context.Context, req *proto.CreateRequest) (*proto.CreateResponse, error) {
	tSanitized := utils.SanitizeString(req.Title)
	movie := models.Movie{
		Title:       &tSanitized,
		ReleaseYear: &req.ReleaseYear,
		DirectorID:  req.DirectorID,
		MusicBy:     req.MusicBy,
		WrittenBy:   req.WrittenBy,
	}

	if err := db.CreateMovie(s.H.Conn, &movie); err != nil {
		st := status.New(codes.Internal, "Cannot create the movie")
		return nil, st.Err()
	}

	return &proto.CreateResponse{
		Status: http.StatusCreated,
	}, nil
}

func (s *MovieServer) Delete(ctx context.Context, req *proto.DeleteRequest) (*proto.DeleteResponse, error) {
	_, err := db.FindMovieById(s.H.Conn, req.Id)
	if err != nil {
		var code codes.Code
		var message string
		if err.Error() == "no rows in result set" {
			code = codes.NotFound
			message = "Movie not found"
		} else {
			code = codes.Internal
			message = "There was an error finding the movie"
		}

		return nil, status.New(code, message).Err()
	}

	if err = db.DeleteMovie(s.H.Conn, req.Id); err != nil {
		st := status.New(codes.Internal, "There was an error deleting the movie")
		return nil, st.Err()
	}

	return &proto.DeleteResponse{
		Status: http.StatusOK,
	}, nil
}

func (s *MovieServer) Modify(ctx context.Context, req *proto.ModifyRequest) (*proto.ModifyResponse, error) {
	var tSanitized string
	reqMovie := models.Movie{
		ReleaseYear: req.ReleaseYear,
		DirectorID:  req.DirectorID,
		MusicBy:     req.MusicBy,
		WrittenBy:   req.WrittenBy,
	}

	if req.Title != nil {
		tSanitized = utils.SanitizeString(*req.Title)
		reqMovie.Title = &tSanitized
	}

	movieMap, args := utils.SetModifyMap(reqMovie)
	if err := db.ModifyMovie(s.H.Conn, req.Id, movieMap, args); err != nil {
		st := status.New(codes.Internal, err.Error())
		return nil, st.Err()
	}

	return &proto.ModifyResponse{
		Status: http.StatusOK,
	}, nil
}

func (s *MovieServer) GetById(ctx context.Context, req *proto.GetByIdRequest) (*proto.GetByIdResponse, error) {
	movie, err := db.FindMovieById(s.H.Conn, req.Id)
	if err != nil {
		var code codes.Code
		var message string
		if err.Error() == "no rows in result set" {
			code = codes.NotFound
			message = "Movie not found"
		} else {
			code = codes.Internal
			message = "There was an error finding the movie"
		}

		return nil, status.New(code, message).Err()
	}

	mp := utils.SetProtoMovie(movie)

	return &proto.GetByIdResponse{
		Movie: mp,
	}, nil
}

func (s *MovieServer) GetAll(ctx context.Context, req *proto.GetAllRequest) (*proto.GetAllResponse, error) {
	var protoMovies []*proto.Movie

	movies, err := db.GetAllMovies(s.H.Conn)
	if err != nil {
		st := status.New(codes.Internal, "There was an error getting all movies: "+err.Error())
		return nil, st.Err()
	}

	if len(movies) < 1 {
		st := status.New(codes.NotFound, "There is no movies")
		return nil, st.Err()
	}

	for _, movie := range movies {
		protoMovie := utils.SetProtoMovie(movie)
		protoMovies = append(protoMovies, protoMovie)
	}

	return &proto.GetAllResponse{
		Movies: protoMovies,
	}, nil
}

func (s *MovieServer) AddGenre(ctx context.Context, req *proto.AddGenreRequest) (*proto.AddGenreResponse, error) {
	if _, err := db.FindMovieById(s.H.Conn, req.MovieId); err != nil {
		return nil, status.New(codes.NotFound, "There is no movie with that id").Err()
	}

	if _, err := db.FindGenreById(s.H.Conn, req.GenreId); err != nil {
		return nil, status.New(codes.NotFound, "There is no genre with that id").Err()
	}

	if err := db.AddGenreToMovie(s.H.Conn, req.GenreId, req.MovieId); err != nil {
		return nil, status.New(codes.Internal, "There was an error creating the join record").Err()
	}

	return &proto.AddGenreResponse{
		Status: http.StatusOK,
	}, nil
}

func (s *MovieServer) RemoveGenre(ctx context.Context, req *proto.RemoveGenreRequest) (*proto.RemoveGenreResponse, error) {
	if _, err := db.FindMovieById(s.H.Conn, req.MovieId); err != nil {
		return nil, status.New(codes.NotFound, "There is no movie with that id").Err()
	}

	if _, err := db.FindGenreById(s.H.Conn, req.GenreId); err != nil {
		return nil, status.New(codes.NotFound, "There is no genre with that id").Err()
	}

	if err := db.RemoveGenreFromMovie(s.H.Conn, req.GenreId, req.MovieId); err != nil {
		return nil, status.New(codes.Internal, "There was an error deleting from join record").Err()
	}

	return &proto.RemoveGenreResponse{
		Status: http.StatusOK,
	}, nil
}
