package controllers

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/rexiaxm7/practice-go/repositories"
)

type UserController interface {
	GetUsers(context echo.Context) error
	PostUser(context echo.Context) error
	GetUser(context echo.Context) error
	PatchUser(context echo.Context) error
	DeleteUser(context echo.Context) error
}

type userController struct {
	userRepository repositories.UserRepository
}

func NewUserController(userRepository repositories.UserRepository) UserController {
	return &userController{userRepository: userRepository}
}

func (u userController) GetUsers(context echo.Context) error {
	users, err := u.userRepository.GetUsers()
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return context.JSON(http.StatusOK, users)
}

func (u userController) PostUser(context echo.Context) error {
	type postUserRequest struct {
		Name string `json:"name"`
	}
	request := new(postUserRequest)
	if err := context.Bind(request); err != nil {
		return err
	}

	err := u.userRepository.CreateUser(request.Name)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return context.String(http.StatusOK, "ok")
}

func (u userController) GetUser(context echo.Context) error {
	id, err := strconv.Atoi(context.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	user, err := u.userRepository.GetUser(id)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return context.JSON(http.StatusOK, user)
}

func (u userController) PatchUser(context echo.Context) error {
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

	err = u.userRepository.UpdateUser(id, request.Name)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return context.NoContent(http.StatusOK)
}

func (u userController) DeleteUser(context echo.Context) error {
	id, err := strconv.Atoi(context.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	err = u.userRepository.DeleteUser(id)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return context.NoContent(http.StatusOK)
}
