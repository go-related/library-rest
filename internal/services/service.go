package services

import (
	"context"
	"github.com/go-related/library-rest/internal/models"
	"github.com/go-related/library-rest/internal/persistance"
)

type Service interface {
	CreateAuthor(ctx context.Context, data models.Author) (*models.Author, error)
	UpdateAuthor(ctx context.Context, data models.Author) error
	DeleteAuthor(ctx context.Context, Id uint) error
	GetAllAuthors(ctx context.Context) ([]*models.Author, error)
	GetAuthorById(ctx context.Context, Id uint) (*models.Author, error)
}

type service struct {
	Db persistance.BooksDB
}

func NewService(db persistance.BooksDB) (*service, error) {
	return &service{Db: db}, nil
}
