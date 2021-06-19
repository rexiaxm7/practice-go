package controllers

import (
	"github.com/labstack/echo/v4"
	"github.com/rexiaxm7/practice-go/repositories"
	"net/http"
	"strconv"
)

func GetUsers(context echo.Context) error {

	users, err := repositories.GetUsers()
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return context.JSON(http.StatusOK, users)
}

func PostUser(context echo.Context) error {

	type postUserRequest struct {
		Name string `json:"name"`
	}
	request := new(postUserRequest)
	if err := context.Bind(request); err != nil {
		return err
	}

	err := repositories.CreateUser(request.Name)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return context.String(http.StatusOK, "ok")
}

func GetUser(context echo.Context) error {

	id, err := strconv.Atoi(context.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	user, err := repositories.GetUser(id)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return context.JSON(http.StatusOK, user)
}

func PatchUser(context echo.Context) error {

	id, err := strconv.Atoi(context.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	type patchUserRequest struct {
		Name string `json:"name"`
	}
	request := new(patchUserRequest)
	if err := context.Bind(request); err != nil {
		return err
	}

	err = repositories.UpdateUser(id, request.Name)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return context.NoContent(http.StatusOK)
}

func DeleteUser(context echo.Context) error {

	id, err := strconv.Atoi(context.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	err = repositories.DeleteUser(id)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return context.NoContent(http.StatusOK)
}
