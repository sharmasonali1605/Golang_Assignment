package service

import (
	"testing"

	"github.com/sharmasonali1605/Golang_Assignment/blogpb"
	"github.com/sharmasonali1605/Golang_Assignment/repository"
)

func setup() *BlogService {
	repo := repository.NewInMemoryBlogRepository()
	return NewBlogService(repo)
}

func TestCreatePost(t *testing.T) {
	svc := setup()

	post := &blogpb.Post{
		Title:   "Test Title",
		Content: "Test Content",
		Author:  "Tester",
		Tags:    []string{"go", "test"},
	}

	created, err := svc.CreatePost(post)
	if err != nil {
		t.Fatalf("CreatePost failed: %v", err)
	}

	if created.PostId == "" {
		t.Error("expected PostId to be set")
	}
}

func TestReadPost(t *testing.T) {
	svc := setup()

	post := &blogpb.Post{
		Title:   "Read Title",
		Content: "Read Content",
		Author:  "Reader",
	}
	created, _ := svc.CreatePost(post)

	read, err := svc.ReadPost(created.PostId)
	if err != nil {
		t.Fatalf("ReadPost failed: %v", err)
	}

	if read.Title != post.Title {
		t.Errorf("expected title %s, got %s", post.Title, read.Title)
	}
}

func TestUpdatePost(t *testing.T) {
	svc := setup()

	original, _ := svc.CreatePost(&blogpb.Post{
		Title:   "Old",
		Content: "Old content",
		Author:  "Updater",
	})

	original.Title = "Updated Title"
	updated, err := svc.UpdatePost(original)
	if err != nil {
		t.Fatalf("UpdatePost failed: %v", err)
	}

	if updated.Title != "Updated Title" {
		t.Errorf("expected title to be updated, got %s", updated.Title)
	}
}

func TestDeletePost(t *testing.T) {
	svc := setup()

	post, _ := svc.CreatePost(&blogpb.Post{
		Title: "Delete Me",
	})

	err := svc.DeletePost(post.PostId)
	if err != nil {
		t.Fatalf("DeletePost failed: %v", err)
	}

	_, err = svc.ReadPost(post.PostId)
	if err == nil {
		t.Error("expected error when reading deleted post, got nil")
	}
}
