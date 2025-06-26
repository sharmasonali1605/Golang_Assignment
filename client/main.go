package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/sharmasonali1605/Golang_Assignment/blogpb"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:8072", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Could not connect: %v", err)
	}
	defer conn.Close()

	client := blogpb.NewBlogServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Create Post
	post := &blogpb.Post{
		Title:   "My First Post",
		Content: "Hello from gRPC!",
		Author:  "Sonali",
		Tags:    []string{"go", "grpc"},
	}

	createRes, err := client.CreatePost(ctx, &blogpb.CreatePostRequest{Post: post})
	if err != nil {
		log.Fatalf("CreatePost failed: %v", err)
	}
	createdPost := createRes.GetPost()
	fmt.Println("CREATED POST:", createdPost)
	fmt.Println("******************************")

	// Read Post
	readRes, err := client.ReadPost(ctx, &blogpb.ReadPostRequest{PostId: createdPost.PostId})
	if err != nil {
		log.Fatalf("ReadPost failed: %v", err)
	}
	fmt.Println("READ POST:", readRes.GetPost())
	fmt.Println("******************************")

	// Update Post
	updatedPost := &blogpb.Post{
		PostId:  createdPost.PostId,
		Title:   "My Updated Post",
		Content: "This post has been updated!",
		Author:  "Sonali Sharma",
		Tags:    []string{"golang", "grpc", "update"},
	}

	updateRes, err := client.UpdatePost(ctx, &blogpb.UpdatePostRequest{Post: updatedPost})
	if err != nil {
		log.Fatalf("UpdatePost failed: %v", err)
	}
	fmt.Println("UPDATED POST:", updateRes.GetPost())
	fmt.Println("******************************")

	// Delete Post
	deleteRes, err := client.DeletePost(ctx, &blogpb.DeletePostRequest{PostId: createdPost.PostId})
	if err != nil {
		log.Fatalf("DeletePost failed: %v", err)
	}
	fmt.Println("DELETE RESPONSE:", deleteRes.GetMessage())
	fmt.Println("******************************")
}
