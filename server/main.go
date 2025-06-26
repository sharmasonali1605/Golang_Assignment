package main

import (
	"fmt"
	"log"
	"net"
	"os"

	"github.com/sharmasonali1605/Golang_Assignment/blogpb"
	"github.com/sharmasonali1605/Golang_Assignment/handler"
	"github.com/sharmasonali1605/Golang_Assignment/repository"
	"github.com/sharmasonali1605/Golang_Assignment/service"
	"google.golang.org/grpc"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8072"
	}

	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", port))
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	repo := repository.NewInMemoryBlogRepository()
	svc := service.NewBlogService(repo)
	grpcServer := grpc.NewServer()

	blogpb.RegisterBlogServiceServer(grpcServer, handler.NewBlogHandler(svc))

	log.Printf("gRPC server started on port %s", port)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
