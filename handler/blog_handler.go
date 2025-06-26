package handler

import (
	"context"
	"fmt"

	"github.com/sharmasonali1605/Golang_Assignment/blogpb"
	"github.com/sharmasonali1605/Golang_Assignment/service"
)

type BlogHandler struct {
	blogpb.UnimplementedBlogServiceServer
	svc *service.BlogService
}

func NewBlogHandler(svc *service.BlogService) *BlogHandler {
	return &BlogHandler{svc: svc}
}

func (h *BlogHandler) CreatePost(ctx context.Context, req *blogpb.CreatePostRequest) (*blogpb.CreatePostResponse, error) {
	post, err := h.svc.CreatePost(req.GetPost())
	if err != nil {
		return nil, fmt.Errorf("failed to create post: %w", err)
	}
	return &blogpb.CreatePostResponse{Post: post}, nil
}

func (h *BlogHandler) ReadPost(ctx context.Context, req *blogpb.ReadPostRequest) (*blogpb.ReadPostResponse, error) {
	post, err := h.svc.ReadPost(req.GetPostId())
	if err != nil {
		return nil, fmt.Errorf("failed to read post: %w", err)
	}
	return &blogpb.ReadPostResponse{Post: post}, nil
}

func (h *BlogHandler) UpdatePost(ctx context.Context, req *blogpb.UpdatePostRequest) (*blogpb.UpdatePostResponse, error) {
	post, err := h.svc.UpdatePost(req.GetPost())
	if err != nil {
		return nil, fmt.Errorf("failed to update post: %w", err)
	}
	return &blogpb.UpdatePostResponse{Post: post}, nil
}

func (h *BlogHandler) DeletePost(ctx context.Context, req *blogpb.DeletePostRequest) (*blogpb.DeletePostResponse, error) {
	err := h.svc.DeletePost(req.GetPostId())
	if err != nil {
		return nil, fmt.Errorf("failed to delete post: %w", err)
	}
	return &blogpb.DeletePostResponse{Message: "Post deleted"}, nil
}
