package repositories

import (
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jinzhu/gorm"
	"github.com/stretchr/testify/assert"

	"github.com/rexiaxm7/practice-go/models"
)

func TestGetUsers(t *testing.T) {

	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatal(err)
	}

	gormDB, err := gorm.Open("mysql", db)
	if err != nil {
		return
	}

	mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `user`")).
		WillReturnRows(sqlmock.NewRows([]string{"id", "name"}).
			AddRow(1, "test"))

	userRepository := NewUserRepository(gormDB)
	results, err := userRepository.GetUsers()
	assert.Equal(t, nil, err)
	assert.Equal(t, []models.User{{Id: 1, Name: "test"}}, results)
}

func TestCreateUser(t *testing.T) {
}

func TestGetUser(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatal(err)
	}

	gormDB, err := gorm.Open("mysql", db)
	if err != nil {
		return
	}

	mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM user WHERE (user.id = ?) ORDER BY user.id ASC LIMIT 1")).
		WithArgs(1).
		WillReturnRows(sqlmock.NewRows([]string{"id", "name"}).
			AddRow(1, "test"))

	userRepository := NewUserRepository(gormDB)
	results, err := userRepository.GetUser(1)
	assert.Equal(t, nil, err)
	assert.Equal(t, models.User{Id: 1, Name: "test"}, results)
}

func TestUpdateUser(t *testing.T) {
}

func TestDeleteUser(t *testing.T) {
}
