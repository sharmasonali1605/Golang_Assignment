package service

import (
	"time"

	"github.com/google/uuid"
	"github.com/sharmasonali1605/Golang_Assignment/blogpb"
	"github.com/sharmasonali1605/Golang_Assignment/repository"
)

type BlogService struct {
	repo repository.BlogRepository
}

func NewBlogService(repo repository.BlogRepository) *BlogService {
	return &BlogService{repo: repo}
}

func (s *BlogService) CreatePost(post *blogpb.Post) (*blogpb.Post, error) {
	post.PostId = uuid.New().String()
	post.PublicationDate = time.Now().Format(time.RFC3339)
	return s.repo.Create(post)
}

func (s *BlogService) ReadPost(id string) (*blogpb.Post, error) {
	return s.repo.Read(id)
}

func (s *BlogService) UpdatePost(post *blogpb.Post) (*blogpb.Post, error) {
	return s.repo.Update(post)
}

func (s *BlogService) DeletePost(id string) error {
	return s.repo.Delete(id)
}

func (s *BlogService) ListPosts() ([]*blogpb.Post, error) {
	return s.repo.List()
}
