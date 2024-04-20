package persistance

import (
	"context"
	"github.com/go-related/library-rest/internal/models"
	"github.com/pkg/errors"
)

func (b *booksDb) CreateAuthor(ctx context.Context, data models.Author) (*models.Author, error) {
	// cancellation check
	select {
	case <-ctx.Done():
		return nil, ctx.Err()
	default:
	}
	result := b.Db.Model(&models.Author{}).Create(&data)
	if result.Error != nil {
		return nil, errors.Wrap(result.Error, "couldn't create the author")
	}
	return &data, result.Error
}

func (b *booksDb) UpdateAuthor(ctx context.Context, data models.Author) error {
	select {
	case <-ctx.Done():
		return ctx.Err()
	default:
	}
	currentData, err := b.GetAuthorById(ctx, data.ID)
	if err != nil {
		return err
	}
	currentData.PublicName = data.PublicName
	result := b.Db.Save(currentData)
	if result.Error != nil {
		//logrus.WithError(result.Error).WithField("id", data.ID).Error("Error updating author")
		return errors.Wrap(result.Error, "error updating author")
	}
	return result.Error
}

func (b *booksDb) DeleteAuthor(ctx context.Context, Id uint) error {
	select {
	case <-ctx.Done():
		return ctx.Err()
	default:
	}
	currentData, err := b.GetAuthorById(ctx, Id)
	if err != nil {
		return err
	}
	result := b.Db.Delete(&currentData)
	if result.Error != nil {
		//logrus.WithError(result.Error).WithField("id", Id).Error("Error deleting author")
		return errors.Wrap(result.Error, "error deleting author")
	}
	return nil
}

func (b *booksDb) GetAllAuthors(ctx context.Context) ([]*models.Author, error) {
	select {
	case <-ctx.Done():
		return nil, ctx.Err()
	default:
	}
	var outputList []*models.Author
	result := b.Db.Model(&models.Author{}).Find(&outputList)
	if result.Error != nil {
		return nil, errors.Wrap(result.Error, "error loading authors")
	}
	return outputList, result.Error
}

func (b *booksDb) GetAuthorById(ctx context.Context, Id uint) (*models.Author, error) {
	select {
	case <-ctx.Done():
		return nil, ctx.Err()
	default:
	}
	var output models.Author
	result := b.Db.Model(&models.Author{}).First(&output, Id)
	if result.Error != nil {
		//logrus.WithError(result.Error).Error("couldn't load author")
		return nil, errors.Wrap(result.Error, "")
	}
	return &output, result.Error
}
