package persistance

import (
	"database/sql/driver"
	"github.com/DATA-DOG/go-sqlmock"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type mockDb struct {
	Db      *booksDb
	SqlMock sqlmock.Sqlmock
}

func NewMockDb() (*mockDb, error) {
	db, mock, err := sqlmock.New()
	if err != nil {
		return nil, err
	}

	gormDB, err := gorm.Open(postgres.New(postgres.Config{
		Conn: db,
	}), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	return &mockDb{
		Db:      &booksDb{Db: gormDB},
		SqlMock: mock,
	}, nil
}

func (mb *mockDb) SetupCreate(query string, args []driver.Value, rows *sqlmock.Rows) {
	mb.SqlMock.ExpectBegin()
	//mock.ExpectExec(`INSERT INTO "authors" (.+) VALUES (.+)`).
	//doesnt work with postgress driver very well
	mb.SqlMock.ExpectQuery(query).
		WithArgs(args...).
		WillReturnRows(rows)
	mb.SqlMock.ExpectCommit()
}

type RequestItem struct {
	query       string
	args        []driver.Value
	rows        *sqlmock.Rows
	item        interface{}
	returnError bool
	err         error
}

func (rq *RequestItem) SetupCreate(db *mockDb) {
	//mock.ExpectExec doesnt work with postgress driver very well
	if rq.returnError {
		db.SqlMock.ExpectBegin()
		db.SqlMock.ExpectQuery(rq.query).
			WithArgs(rq.args...).
			WillReturnError(rq.err)
		db.SqlMock.ExpectRollback()
	} else {
		db.SqlMock.ExpectBegin()
		db.SqlMock.ExpectQuery(rq.query).
			WithArgs(rq.args...).
			WillReturnRows(rq.rows)
		db.SqlMock.ExpectCommit()
	}
}
