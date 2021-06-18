package route

import (
	"github.com/labstack/echo/v4"
	"github.com/rexiaxm7/practice-go/api"
)

func Initialize(e *echo.Echo) {

	userGroup := e.Group("/users")
	{
		userGroup.GET("", api.GetUsers)
		userGroup.POST("", api.PostUser)
		userGroup.GET("/:id", api.GetUser)
		userGroup.PATCH("/:id", api.PatchUser)
		userGroup.DELETE("/:id", api.DeleteUser)
	}
}
