package route

import (
	"github.com/labstack/echo/v4"
	"github.com/rexiaxm7/practice-go/controllers"
)

func Initialize(e *echo.Echo) {

	userGroup := e.Group("/users")
	{
		userGroup.GET("", controllers.GetUsers)
		userGroup.POST("", controllers.PostUser)
		userGroup.GET("/:id", controllers.GetUser)
		userGroup.PATCH("/:id", controllers.PatchUser)
		userGroup.DELETE("/:id", controllers.DeleteUser)
	}
}
