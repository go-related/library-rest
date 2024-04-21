package services

import (
	"context"
	"errors"
	"github.com/go-related/library-rest/internal/models"
	"gorm.io/gorm"
)

func (s *service) CreateAuthor(ctx context.Context, data *models.Author) (*models.Author, error) {
	// cancellation check
	select {
	case <-ctx.Done():
		return nil, ctx.Err()
	default:
	}
	err := s.validateAuthor(data)
	if err != nil {
		return nil, err
	}
	return s.Db.CreateAuthor(*data)
}

func (s *service) UpdateAuthor(ctx context.Context, data *models.Author) error {
	select {
	case <-ctx.Done():
		return ctx.Err()
	default:
	}
	err := s.validateAuthor(data)
	if err != nil {
		return err
	}
	return s.Db.UpdateAuthor(*data)
}

func (s *service) GetAuthorById(ctx context.Context, Id uint) (*models.Author, error) {
	select {
	case <-ctx.Done():
		return nil, ctx.Err()
	default:
	}
	author, err := s.Db.GetAuthorById(Id)
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, NewServiceError("author not found")
	}
	return author, err
}

func (s *service) DeleteAuthor(ctx context.Context, Id uint) error {
	select {
	case <-ctx.Done():
		return ctx.Err()
	default:
	}
	author, _ := s.Db.GetAuthorById(Id)
	if author == nil {
		return NewServiceError("invalid id for the author")
	}
	return s.Db.DeleteAuthor(Id)
}

func (s *service) GetAllAuthors(ctx context.Context) ([]*models.Author, error) {
	select {
	case <-ctx.Done():
		return nil, ctx.Err()
	default:
	}
	return s.Db.GetAllAuthors()
}

func (s *service) validateAuthor(data *models.Author) error {
	if data == nil {
		return NewServiceError("invalid input")
	}

	if data.PublicName == "" {
		return NewServiceError("invalid public_name for the author")
	}

	if data.Model.ID != 0 {
		author, _ := s.Db.GetAuthorById(data.Model.ID)
		if author == nil {
			return NewServiceError("invalid id for the author")
		}
	}
	return nil
}
