package persistance

import (
	"database/sql/driver"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/go-related/library-rest/internal/models"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestCreateAuthor(t *testing.T) {

	testCases := []struct {
		name           string
		requestItem    RequestItem
		expectedResult interface{}
	}{
		{
			name: "verify_create",
			requestItem: RequestItem{
				query: `INSERT INTO "authors"`,
				args:  []driver.Value{sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), "Jane Doe"},
				rows:  sqlmock.NewRows([]string{"1"}),
				item:  models.Author{PublicName: "Jane Doe"},
			},
			expectedResult: models.Author{
				PublicName: "Jane Doe",
			},
		},
		{
			name: "verify_create_db_error",
			requestItem: RequestItem{
				query:       `INSERT INTO "authors"`,
				args:        []driver.Value{sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), "Jane Doe"},
				item:        models.Author{PublicName: "Jane Doe"},
				returnError: true,
				err:         errors.New("test"),
			},
			expectedResult: models.Author{
				PublicName: "Jane Doe",
			},
		},
	}
	for _, test := range testCases {
		t.Run(test.name, func(t *testing.T) {
			// arrange
			db, err := NewMockDb()
			require.NoError(t, err)
			test.requestItem.SetupQuery(db)

			// act
			result, err := db.Db.CreateAuthor(test.requestItem.item.(models.Author))

			// assert
			if test.requestItem.returnError {
				require.Error(t, err, "error expected but didn't get it.")
			} else {
				require.NoError(t, err, "creating author should not return an error")
				require.NotNil(t, result, "created author should not be nil")
				err = db.SqlMock.ExpectationsWereMet()
				assert.NoError(t, err, "all expectations should have been met")

				expectedResult := test.expectedResult.(models.Author)
				assert.Equal(t, expectedResult.PublicName, result.PublicName)
			}
		})
	}
}
