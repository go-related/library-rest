package persistance

import (
	"github.com/go-related/library-rest/internal/models"
	"github.com/pkg/errors"
)

func (b *booksDb) CreateAuthor(data models.Author) (*models.Author, error) {
	result := b.Db.Model(&models.Author{}).Create(&data)
	if result.Error != nil {
		return nil, errors.Wrap(result.Error, "couldn't create the author")
	}
	return &data, result.Error
}

func (b *booksDb) UpdateAuthor(data models.Author) error {
	currentData, err := b.GetAuthorById(data.ID)
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

func (b *booksDb) DeleteAuthor(Id uint) error {
	currentData, err := b.GetAuthorById(Id)
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

func (b *booksDb) GetAllAuthors() ([]*models.Author, error) {
	var outputList []*models.Author
	result := b.Db.Model(&models.Author{}).Find(&outputList)
	if result.Error != nil {
		return nil, errors.Wrap(result.Error, "error loading authors")
	}
	return outputList, result.Error
}

func (b *booksDb) GetAuthorById(Id uint) (*models.Author, error) {
	var output models.Author
	result := b.Db.Model(&models.Author{}).First(&output, Id)
	return &output, result.Error
}
