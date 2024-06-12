package main

import (
	"log"
	"net"

	"github.com/DuvanAlbarracin/movies_movies/pkg/config"
	"github.com/DuvanAlbarracin/movies_movies/pkg/db"
	"github.com/DuvanAlbarracin/movies_movies/pkg/proto"
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
	log.Println("Auth service on:", config.Port)

	server := services.Server{
		H: h,
	}
	defer server.H.Conn.Close()

	grpcServer := grpc.NewServer()
	proto.RegisterMoviesServiceServer(grpcServer, &server)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalln("Failed to serve:", err)
	}
}
