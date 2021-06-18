package api

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func GetUsers(context echo.Context) error {
	return context.String(http.StatusOK, "GetUsers")
}

func PostUser(context echo.Context) error {
	return context.String(http.StatusOK, "PostUser")
}

func GetUser(context echo.Context) error {
	return context.String(http.StatusOK, "GetUser")
}

func PatchUser(context echo.Context) error {
	return context.String(http.StatusOK, "PatchUser")
}

func DeleteUser(context echo.Context) error {
	return context.String(http.StatusOK, "DeleteUser")
}
