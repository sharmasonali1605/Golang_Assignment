package main

import (
	"context"
	"fmt"
	"sync"

	"github.com/sharmasonali1605/Golang_Assignment/blogpb"

	"github.com/google/uuid"
)

type BlogServer struct {
	blogpb.UnimplementedBlogServiceServer
	mu    sync.Mutex
	posts map[string]*blogpb.Post
}

func NewBlogServer() *BlogServer {
	return &BlogServer{
		posts: make(map[string]*blogpb.Post),
	}
}

func (s *BlogServer) CreatePost(ctx context.Context, req *blogpb.CreatePostRequest) (*blogpb.CreatePostResponse, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	id := uuid.New().String()
	req.Post.PostId = id
	s.posts[id] = req.Post

	return &blogpb.CreatePostResponse{Post: req.Post}, nil
}

func (s *BlogServer) ReadPost(ctx context.Context, req *blogpb.ReadPostRequest) (*blogpb.ReadPostResponse, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	post, exists := s.posts[req.PostId]
	if !exists {
		return nil, fmt.Errorf("post not found")
	}

	return &blogpb.ReadPostResponse{Post: post}, nil
}

func (s *BlogServer) UpdatePost(ctx context.Context, req *blogpb.UpdatePostRequest) (*blogpb.UpdatePostResponse, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	post := req.Post
	if _, ok := s.posts[post.PostId]; !ok {
		return nil, fmt.Errorf("post not found")
	}
	s.posts[post.PostId] = post
	return &blogpb.UpdatePostResponse{Post: post}, nil
}

func (s *BlogServer) DeletePost(ctx context.Context, req *blogpb.DeletePostRequest) (*blogpb.DeletePostResponse, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	if _, ok := s.posts[req.PostId]; !ok {
		return nil, fmt.Errorf("post not found")
	}
	delete(s.posts, req.PostId)
	return &blogpb.DeletePostResponse{Message: "Post deleted"}, nil
}
