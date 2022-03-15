package mock

import (
	"database/sql"
	"regexp"
	"testing"

	"22dojo-online/pkg/domain/entity"
	driver "22dojo-online/pkg/driver/mysql"
	"22dojo-online/pkg/driver/mysql/database"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

func DummySQLHandler(db *sql.DB) database.SQLHandler {
	SQLHandler := new(driver.SQLHandlerImpl)
	SQLHandler.Conn = db
	return SQLHandler
}

// go test -v pkg/driver/mysql/database/mock/user_repository_test.go
func TestSelectUserByPrimaryKey(t *testing.T) {
	table := []struct {
		testName string
		id       string
		user     entity.User
		err      error
	}{
		{
			"FIRST TEST CASE: SelectUserByPrimaryKey from pkg/interfaces/database/user_repository.go",
			"78164dcf-6b7c-45e4-862a-2a0f6735a449",
			entity.User{
				ID:        "78164dcf-6b7c-45e4-862a-2a0f6735a449",
				AuthToken: "b187b9e0-08e6-42dd-a9b3-a900b137983c",
				Name:      "whatt",
				HighScore: 1000,
				Coin:      85100,
			},
			nil,
		},
	}

	/*   prepare   */
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Error("sqlmock not work")
	}
	defer db.Close()

	repo := &database.UserRepository{
		SQLHandler: DummySQLHandler(db),
	}

	query := "SELECT `id`, `auth_token`, `name`, `high_score`, `coin` FROM `user` WHERE `id`= ?"

	for _, tt := range table {
		t.Run(tt.testName, func(t *testing.T) {
			b := &tt.user
			rows := sqlmock.NewRows([]string{
				"id", "auth_token", "name", "high_score", "coin",
			}).AddRow(b.ID, b.AuthToken, b.Name, b.HighScore, b.Coin)
			mock.ExpectQuery(regexp.QuoteMeta(query)).WithArgs(tt.id).WillReturnRows(rows)
			got, err := repo.SelectUserByPrimaryKey(b.ID)
			assert.Equal(t, tt.err, err)
			assert.Equal(t, b, got)
		})
	}
}
func TestSelectUserByAuthToken(t *testing.T) {
	table := []struct {
		testName  string
		authToken string
		user      entity.User
		err       error
	}{
		{
			"FIRST TEST CASE: SelectUserByAuthToken from pkg/interfaces/database/user_repository.go",
			"85c005f2-13bf-4542-8eba-4e69b569ee2b",
			entity.User{
				ID:        "281a813c-839f-4bc7-834c-f4bc59389f9a",
				AuthToken: "85c005f2-13bf-4542-8eba-4e69b569ee2b",
				Name:      "a",
				HighScore: 0,
				Coin:      0,
			},
			nil,
		},
	}

	/*   prepare   */
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Error("sqlmock not work")
	}
	defer db.Close()

	repo := &database.UserRepository{
		SQLHandler: DummySQLHandler(db),
	}

	query := "SELECT `id`, `auth_token`, `name`, `high_score`, `coin` FROM `user` WHERE `auth_token`= ?"

	for _, tt := range table {
		t.Run(tt.testName, func(t *testing.T) {
			b := &tt.user
			rows := sqlmock.NewRows([]string{
				"id", "auth_token", "name", "high_score", "coin",
			}).AddRow(b.ID, b.AuthToken, b.Name, b.HighScore, b.Coin)
			mock.ExpectQuery(regexp.QuoteMeta(query)).WithArgs(tt.authToken).WillReturnRows(rows)
			got, err := repo.SelectUserByAuthToken(tt.authToken)

			assert.Equal(t, tt.err, err)
			assert.Equal(t, b, got)
		})
	}
}

func TestUpdateUser(t *testing.T) {
	table := []struct {
		testName     string
		id           string
		nameToChange string
		user         entity.User
		err          error
	}{
		{
			"FIRST TEST CASE: InsertUser from pkg/interfaces/database/user_repository.go",
			"78164dcf-6b7c-45e4-862a-2a0f6735a449",
			"whattt",
			entity.User{
				ID:        "78164dcf-6b7c-45e4-862a-2a0f6735a449",
				AuthToken: "b187b9e0-08e6-42dd-a9b3-a900b137983c",
				Name:      "whatt",
				HighScore: 100,
				Coin:      10000,
			},
			nil,
		},
	}

	/*   prepare   */
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Error("sqlmock not work")
	}
	defer db.Close()

	repo := &database.UserRepository{
		SQLHandler: DummySQLHandler(db),
	}

	query := "UPDATE user SET name = ? WHERE id = ?"

	for _, tt := range table {
		t.Run(tt.testName, func(t *testing.T) {
			mock.ExpectExec(regexp.QuoteMeta(query)).WithArgs(tt.nameToChange, tt.id).WillReturnResult(sqlmock.NewResult(1, 1))
			tt.user.Name = "whattt"
			if err = repo.UpdateUserByPrimaryKey(&tt.user); err != nil {
				t.Errorf("error was not expected while updating stats: %s", err)
			}

			if err := mock.ExpectationsWereMet(); err != nil {
				t.Errorf("there were unfulfilled expectaions: %s", err)
			}
		})
	}
}

func TestInsertUser(t *testing.T) {
	table := []struct {
		testName string
		name     string
		user     entity.User
		err      error
	}{
		{
			"FIRST TEST CASE: InsertUser from pkg/interfaces/database/user_repository.go",
			"whattt",
			entity.User{
				ID:        "78164dcf-6b7c-45e4-862a-2a0f6735a449",
				AuthToken: "b187b9e0-08e6-42dd-a9b3-a900b137983c",
				Name:      "whatt",
				HighScore: 0,
				Coin:      0,
			},
			nil,
		},
	}

	/*   prepare   */
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Error("sqlmock not work")
	}
	defer db.Close()

	repo := &database.UserRepository{
		SQLHandler: DummySQLHandler(db),
	}

	query := "INSERT INTO `user` (`id`, `auth_token`, `name`, `high_score`, `coin`) VALUES (?, ?, ?, ?, ?);"

	for _, tt := range table {
		t.Run(tt.testName, func(t *testing.T) {
			mock.ExpectExec(regexp.QuoteMeta(query)).WithArgs(tt.user.ID, tt.user.AuthToken, tt.user.Name, tt.user.HighScore, tt.user.Coin).WillReturnResult(sqlmock.NewResult(1, 1))
			if err = repo.InsertUser(&tt.user); err != nil {
				t.Errorf("error was not expected while updating stats: %s", err)
			}
			if err := mock.ExpectationsWereMet(); err != nil {
				t.Errorf("there were unfulfilled expectaions: %s", err)
			}
		})
	}
}
