package main

import (
	"log"
	"net"

	"github.com/DuvanAlbarracin/movies_movies/pkg/config"
	"github.com/DuvanAlbarracin/movies_movies/pkg/db"
	"github.com/DuvanAlbarracin/movies_movies/pkg/proto/genre"
	"github.com/DuvanAlbarracin/movies_movies/pkg/proto/movie"
	"github.com/DuvanAlbarracin/movies_movies/pkg/services"
	"google.golang.org/grpc"
)

func main() {
	config, err := config.LoadConfig()
	if err != nil {
		log.Fatalln("Failed loading config:", err)
	}

	h := db.Init(config.DBUrl)
	lis, err := net.Listen("tcp", config.Port)
	if err != nil {
		log.Fatalln("Failed to listening:", err)
	}
	log.Println("Movie service on:", config.Port)

	movieServer := services.MovieServer{
		H: h,
	}
	defer movieServer.H.Conn.Close()

	genreServer := services.GenreServer{
		H: h,
	}
	defer genreServer.H.Conn.Close()

	grpcServer := grpc.NewServer()
	movie.RegisterMoviesServiceServer(grpcServer, &movieServer)
	genre.RegisterGenreServiceServer(grpcServer, &genreServer)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalln("Failed to serve:", err)
	}
}
