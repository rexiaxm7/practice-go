package route

import (
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo/v4"

	"github.com/rexiaxm7/practice-go/controllers"
	"github.com/rexiaxm7/practice-go/repositories"
)

func Initialize(e *echo.Echo, gormDB *gorm.DB) {

	userRepository := repositories.NewUserRepository(gormDB)

	userGroup := e.Group("/users")
	{
		userController := controllers.NewUserController(userRepository)
		userGroup.GET("", userController.GetUsers)
		userGroup.POST("", userController.PostUser)
		userGroup.GET("/:id", userController.GetUser)
		userGroup.PATCH("/:id", userController.PatchUser)
		userGroup.DELETE("/:id", userController.DeleteUser)
	}
}
