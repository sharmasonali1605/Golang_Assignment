package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/sharmasonali1605/Golang_Assignment/blogpb"
	"google.golang.org/grpc"
)

var client blogpb.BlogServiceClient
var reader *bufio.Reader

func main() {
	conn, err := grpc.Dial("localhost:8072", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Could not connect: %v", err)
	}
	defer conn.Close()

	client = blogpb.NewBlogServiceClient(conn)
	reader = bufio.NewReader(os.Stdin)

	for {
		fmt.Println("\n Choose operation: create | read | update | delete | exit")
		fmt.Print("-> ")
		cmd := readLine()

		switch cmd {
		case "create":
			createPost()
		case "read":
			readPost()
		case "update":
			updatePost()
		case "delete":
			deletePost()
		default:
			fmt.Println("Unknown command.")
		}
	}
}

func createPost() {
	title := prompt("Title")
	content := prompt("Content")
	author := prompt("Author")
	tags := strings.Split(prompt("Tags (comma separated)"), ",")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	post := &blogpb.Post{
		Title:   title,
		Content: content,
		Author:  author,
		Tags:    cleanTags(tags),
	}
	res, err := client.CreatePost(ctx, &blogpb.CreatePostRequest{Post: post})
	if err != nil {
		fmt.Printf("Failed to create post: %v\n", err)
		return
	}
	fmt.Println("Post Created:", res.GetPost())
}

func readPost() {
	postID := prompt("Enter Post ID")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	res, err := client.ReadPost(ctx, &blogpb.ReadPostRequest{PostId: postID})
	if err != nil {
		fmt.Printf("Failed to read post: %v\n", err)
		return
	}
	fmt.Println("Post Details:", res.GetPost())
}

func updatePost() {
	postID := prompt("Post ID to update")
	title := prompt("New Title")
	content := prompt("New Content")
	author := prompt("New Author")
	tags := strings.Split(prompt("New Tags (comma separated)"), ",")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	post := &blogpb.Post{
		PostId:  postID,
		Title:   title,
		Content: content,
		Author:  author,
		Tags:    cleanTags(tags),
	}
	res, err := client.UpdatePost(ctx, &blogpb.UpdatePostRequest{Post: post})
	if err != nil {
		fmt.Printf("Failed to update post: %v\n", err)
		return
	}
	fmt.Println("Post Updated:", res.GetPost())
}

func deletePost() {
	postID := prompt("Enter Post ID to delete")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	res, err := client.DeletePost(ctx, &blogpb.DeletePostRequest{PostId: postID})
	if err != nil {
		fmt.Printf(" Failed to delete post: %v\n", err)
		return
	}
	fmt.Println("Post Deleted:", res.GetMessage())
}

func prompt(label string) string {
	fmt.Print(label + ": ")
	return strings.TrimSpace(readLine())
}

func readLine() string {
	input, _ := reader.ReadString('\n')
	return strings.TrimSpace(input)
}

func cleanTags(tags []string) []string {
	var cleaned []string
	for _, tag := range tags {
		if t := strings.TrimSpace(tag); t != "" {
			cleaned = append(cleaned, t)
		}
	}
	return cleaned
}
