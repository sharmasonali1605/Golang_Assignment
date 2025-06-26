package main

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/google/uuid"
	"github.com/sharmasonali1605/Golang_Assignment/blogpb"
)

type BlogServer struct {
	blogpb.UnimplementedBlogServiceServer
	mu    sync.Mutex
	posts map[string]*blogpb.Post
}

// Constructor for the BlogServer
func NewBlogServer() *BlogServer {
	return &BlogServer{
		posts: make(map[string]*blogpb.Post),
	}
}

func (s *BlogServer) CreatePost(ctx context.Context, req *blogpb.CreatePostRequest) (*blogpb.CreatePostResponse, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	id := uuid.New().String()
	post := req.GetPost()
	post.PostId = id

	if post.PublicationDate == "" {
		post.PublicationDate = time.Now().Format(time.RFC3339)
	}

	s.posts[id] = post

	return &blogpb.CreatePostResponse{Post: post}, nil
}

func (s *BlogServer) ReadPost(ctx context.Context, req *blogpb.ReadPostRequest) (*blogpb.ReadPostResponse, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	post, exists := s.posts[req.GetPostId()]
	if !exists {
		return nil, fmt.Errorf("post not found")
	}

	return &blogpb.ReadPostResponse{Post: post}, nil
}

func (s *BlogServer) UpdatePost(ctx context.Context, req *blogpb.UpdatePostRequest) (*blogpb.UpdatePostResponse, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	post := req.GetPost()
	if post == nil || post.PostId == "" {
		return nil, fmt.Errorf("invalid update request")
	}

	existing, ok := s.posts[post.PostId]
	if !ok {
		return nil, fmt.Errorf("post not found")
	}

	post.PublicationDate = existing.PublicationDate
	s.posts[post.PostId] = post

	return &blogpb.UpdatePostResponse{Post: post}, nil
}

func (s *BlogServer) DeletePost(ctx context.Context, req *blogpb.DeletePostRequest) (*blogpb.DeletePostResponse, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	if _, ok := s.posts[req.GetPostId()]; !ok {
		return nil, fmt.Errorf("post not found")
	}

	delete(s.posts, req.GetPostId())
	return &blogpb.DeletePostResponse{Message: "Post deleted"}, nil
}
