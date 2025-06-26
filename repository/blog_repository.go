package repository

import (
	"errors"
	"sync"

	"github.com/sharmasonali1605/Golang_Assignment/blogpb"
)

type BlogRepository interface {
	Create(post *blogpb.Post) (*blogpb.Post, error)
	Read(id string) (*blogpb.Post, error)
	Update(post *blogpb.Post) (*blogpb.Post, error)
	Delete(id string) error
	List() ([]*blogpb.Post, error)
}

type InMemoryBlogRepository struct {
	mu    sync.Mutex
	posts map[string]*blogpb.Post
}

func NewInMemoryBlogRepository() *InMemoryBlogRepository {
	return &InMemoryBlogRepository{
		posts: make(map[string]*blogpb.Post),
	}
}

func (r *InMemoryBlogRepository) Create(post *blogpb.Post) (*blogpb.Post, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.posts[post.PostId] = post
	return post, nil
}

func (r *InMemoryBlogRepository) Read(id string) (*blogpb.Post, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	post, ok := r.posts[id]
	if !ok {
		return nil, errors.New("post not found")
	}
	return post, nil
}

func (r *InMemoryBlogRepository) Update(post *blogpb.Post) (*blogpb.Post, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	_, ok := r.posts[post.PostId]
	if !ok {
		return nil, errors.New("post not found")
	}
	r.posts[post.PostId] = post
	return post, nil
}

func (r *InMemoryBlogRepository) Delete(id string) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	if _, ok := r.posts[id]; !ok {
		return errors.New("post not found")
	}
	delete(r.posts, id)
	return nil
}

func (r *InMemoryBlogRepository) List() ([]*blogpb.Post, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	var posts []*blogpb.Post
	for _, post := range r.posts {
		posts = append(posts, post)
	}
	return posts, nil
}
