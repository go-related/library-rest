package persistance

import (
	"github.com/go-related/library-rest/internal/models"
	"github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type BooksDB interface {
	CreateAuthor(data models.Author) (*models.Author, error)
	UpdateAuthor(data models.Author) error
	DeleteAuthor(Id uint) error
	GetAllAuthors() ([]*models.Author, error)
	GetAuthorById(Id uint) (*models.Author, error)

	//CreateGenre(ctx context.Context, data models.Genre) (*models.Genre, error)
	//UpdateGenre(ctx context.Context, data models.Genre) error
	//DeleteGenre(ctx context.Context, Id uint) error
	//GetAllGenres(ctx context.Context) ([]*models.Genre, error)
	//GetGenresById(ctx context.Context, Id uint) (*models.Genre, error)
	//
	//CreateBook(ctx context.Context, data models.Book) (*models.Book, error)
	//UpdateBook(ctx context.Context, data models.Book) error
	//DeleteBook(ctx context.Context, Id uint) error
	//GetAllBooks(ctx context.Context) ([]*models.Book, error)
	//GetBookById(ctx context.Context, Id uint) (*models.Book, error)
	//
	//CreateSubscriber(ctx context.Context, data models.Subscriber) (*models.Subscriber, error)
	//UpdateSubscriber(ctx context.Context, data models.Subscriber) error
	//DeleteSubscriber(ctx context.Context, Id uint) error
	//GetAllSubscribers(ctx context.Context) ([]*models.Subscriber, error)
	//GetSubscriberById(ctx context.Context, Id uint) (*models.Subscriber, error)
	//
	//Subscribe(ctx context.Context, subscriberID uint, listOfBooks []models.Book, listOfAuthors []models.Author) (*models.Subscribe, error)
	//DeleteSubscribe(ctx context.Context, Id uint) error
	//GetAllSubscribes(ctx context.Context) ([]*models.Subscribe, error)
	//GetSubscribeById(ctx context.Context, Id uint) (*models.Subscribe, error)
	//GetAuthorsSubscribers(ctx context.Context, listOfAuthors []models.Author) ([]*models.Subscriber, error)
}

type booksDb struct {
	Db *gorm.DB
}

func NewBooks(connection string) (BooksDB, error) {
	result := booksDb{}
	db, err := gorm.Open(postgres.Open(connection), &gorm.Config{})
	if err != nil {
		logrus.WithError(err).Error("error connecting to db")
		return &result, err
	}
	err = db.AutoMigrate(&models.Author{}, &models.Genre{}, &models.Subscriber{}, &models.Book{}, &models.Subscribe{})
	if err != nil {
		logrus.WithError(err).Error("couldn't migrate the db")
		return &result, err
	}
	result.Db = db
	return &result, nil
}
