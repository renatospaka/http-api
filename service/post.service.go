package service

import (
	"errors"
	"math/rand"

	"github.com/renatospaka/golang-rest-api/entity"
	"github.com/renatospaka/golang-rest-api/repository"
)

type PostService interface {
	Validate(post *entity.Post) error
	Create(post *entity.Post) (*entity.Post, error)
	FindAll() ([]entity.Post, error)
}

type service struct {}

var (
	repo repository.PostRepository 
)

func NewPostService(repository repository.PostRepository) PostService {
	repo = repository
	return &service{}
}

func (*service) Validate(post *entity.Post) error {
	if post == nil {
		return errors.New("The post is empty")
	}

	if post.Title == ""  {
		return errors.New("The post title is empty")
	}

	if post.Text == ""  {
		return errors.New("The post text is empty")
	}

	return nil
}


func (*service) Create(post *entity.Post) (*entity.Post, error) {
	post.ID = rand.Int63()
	_, err := repo.Save(post)
	return post, err
}

func (*service) FindAll() ([]entity.Post, error) {
	return repo.FindAll()
}

