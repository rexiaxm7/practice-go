package repositories

import (
	"github.com/rexiaxm7/practice-go/database"
	"github.com/rexiaxm7/practice-go/models"
)

func GetUsers() ([]models.User, error) {
	var users []models.User
	result := database.Database.Find(&users)
	if result.Error != nil {
		return nil, result.Error
	}

	return users, nil
}

func CreateUser(name string) error {
	user := models.User{
		Name: name,
	}

	result := database.Database.Create(&user)
	if result.Error == nil {
		return result.Error
	}

	return nil
}

func GetUser(id int) (models.User, error) {
	user := new(models.User)
	result := database.Database.First(user, id)
	if result.Error != nil {
		return models.User{}, result.Error
	}

	return *user, nil
}

func UpdateUser(id int, name string) error {
	user := models.User{
		Id: id,
	}

	result := database.Database.Model(user).Update("name", name)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func DeleteUser(id int) error {
	user := models.User{
		Id: id,
	}

	result := database.Database.Delete(user)
	if result.Error != nil {
		return result.Error
	}

	return nil
}
