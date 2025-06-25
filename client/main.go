package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"grpc-blog/blogpb"

	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:8072", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Could not connect: %v", err)
	}
	defer conn.Close()

	client := blogpb.NewBlogServiceClient(conn)

	post := &blogpb.BlogPost{
		Title:   "My First Post",
		Content: "Hello from gRPC!",
		Author:  "Sonali",
		Tags:    []string{"go", "grpc"},
	}

	// Create
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	res, err := client.CreatePost(ctx, &blogpb.CreatePostRequest{Post: post})
	if err != nil {
		log.Fatalf("CreatePost failed: %v", err)
	}

	fmt.Println("Created Post:", res.GetPost())
}
