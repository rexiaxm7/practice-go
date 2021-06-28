package repositories

import (
	"github.com/jinzhu/gorm"

	"github.com/rexiaxm7/practice-go/models"
)

type UserRepository interface {
	GetUsers() ([]models.User, error)
	CreateUser(name string) error
	GetUser(id int) (models.User, error)
	UpdateUser(id int, name string) error
	DeleteUser(id int) error
}

type userRepository struct {
	gormDB *gorm.DB
}

func NewUserRepository(gormDB *gorm.DB) UserRepository {
	return &userRepository{gormDB: gormDB}
}

func (u userRepository) GetUsers() ([]models.User, error) {
	var users []models.User
	result := u.gormDB.Find(&users)
	if result.Error != nil {
		return nil, result.Error
	}

	return users, nil
}

func (u userRepository) CreateUser(name string) error {
	user := models.User{
		Name: name,
	}

	result := u.gormDB.Create(&user)
	if result.Error == nil {
		return result.Error
	}

	return nil
}

func (u userRepository) GetUser(id int) (models.User, error) {
	user := new(models.User)
	result := u.gormDB.First(user, id)
	if result.Error != nil {
		return models.User{}, result.Error
	}

	return *user, nil
}

func (u userRepository) UpdateUser(id int, name string) error {
	user := models.User{
		Id: id,
	}

	result := u.gormDB.Model(user).Update("name", name)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (u userRepository) DeleteUser(id int) error {
	user := models.User{
		Id: id,
	}

	result := u.gormDB.Delete(user)
	if result.Error != nil {
		return result.Error
	}

	return nil
}
