package controllers

import (
	"net/http"
	"strconv"

	"github.com/hikayat13/alterra-agcm/day-2/submission/helper"
	"github.com/hikayat13/alterra-agcm/day-2/submission/lib/database"
	"github.com/hikayat13/alterra-agcm/day-2/submission/models"
	"github.com/labstack/echo/v4"
)

func GetUserController(c echo.Context) error {
	users, err := database.GetUsers(c)
	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.APIResponse("Bad request", http.StatusBadRequest, "error", nil))
	}
	res := helper.APIResponse("Get all users", http.StatusOK, "success", users)

	return c.JSON(http.StatusOK, res)
}

func GetUserByIdController(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	users, err := database.GetUserById(c, id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.APIResponse(err.Error(), http.StatusBadRequest, "error", nil))
	}
	res := helper.APIResponse("Get user by id", http.StatusOK, "success", users)

	return c.JSON(http.StatusOK, res)
}

func CreateUserController(c echo.Context) error {
	user := models.User{}
	if err := c.Bind(user); err != nil {
		return err
	}
	err := database.CreateUser(c, user)
	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.APIResponse("Bad request", http.StatusBadRequest, "error", nil))
	}
	res := helper.APIResponse("Create user success", http.StatusOK, "success", nil)

	return c.JSON(http.StatusOK, res)
}

func UpdateUserController(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	user := new(models.User)
	if err := c.Bind(user); err != nil {
		return err
	}
	users, err := database.UpdateUser(c, user, id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.APIResponse("Bad request", http.StatusBadRequest, "error", nil))
	}
	res := helper.APIResponse("Update user success", http.StatusOK, "success", users)

	return c.JSON(http.StatusOK, res)
}

func DeleteUserIdController(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	err := database.DeleteUserById(c, id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.APIResponse(err.Error(), http.StatusBadRequest, "error", nil))
	}
	res := helper.APIResponse("Delete user success", http.StatusOK, "success", nil)

	return c.JSON(http.StatusOK, res)
}
