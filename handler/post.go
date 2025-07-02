package handler

import (
	"context"

	"github.com/sharmasonali1605/Golang_Assignment/blogpb"
	"github.com/sharmasonali1605/Golang_Assignment/service"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
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
		return nil, status.Errorf(codes.Internal, "failed to create post: %v", err)
	}
	return &blogpb.CreatePostResponse{Post: post}, nil
}

func (h *BlogHandler) ReadPost(ctx context.Context, req *blogpb.ReadPostRequest) (*blogpb.ReadPostResponse, error) {
	post, err := h.svc.ReadPost(req.GetPostId())
	if err != nil {
		return nil, status.Errorf(codes.NotFound, "post not found: %v", err)
	}
	return &blogpb.ReadPostResponse{Post: post}, nil
}

func (h *BlogHandler) UpdatePost(ctx context.Context, req *blogpb.UpdatePostRequest) (*blogpb.UpdatePostResponse, error) {
	post := req.GetPost()
	if post == nil || post.PostId == "" {
		return nil, status.Error(codes.InvalidArgument, "post ID is required for update")
	}

	updatedPost, err := h.svc.UpdatePost(post)
	if err != nil {
		return nil, status.Errorf(codes.NotFound, "update failed: %v", err)
	}
	return &blogpb.UpdatePostResponse{Post: updatedPost}, nil
}

func (h *BlogHandler) DeletePost(ctx context.Context, req *blogpb.DeletePostRequest) (*blogpb.DeletePostResponse, error) {
	err := h.svc.DeletePost(req.GetPostId())
	if err != nil {
		return nil, status.Errorf(codes.NotFound, "delete failed: %v", err)
	}
	return &blogpb.DeletePostResponse{Message: "Post deleted"}, nil
}

func (h *BlogHandler) ListPost(ctx context.Context, req *blogpb.ListPostRequest) (*blogpb.ListPostResponse, error) {
	posts, err := h.svc.ListPosts()
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to list posts: %v", err)
	}

	return &blogpb.ListPostResponse{Post: posts}, nil
}
