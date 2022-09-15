package controllers

import (
	"github.com/hikayat13/alterra-agcm/day-2/submission/middlewares"
	"github.com/hikayat13/alterra-agcm/day-2/submission/request"
	"github.com/hikayat13/alterra-agcm/day-2/submission/transform"
	"net/http"
	"strconv"

	"github.com/hikayat13/alterra-agcm/day-2/submission/helper"
	"github.com/hikayat13/alterra-agcm/day-2/submission/lib/database"
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
	req := &request.UserCreate{}

	if err := c.Bind(req); err != nil {
		return err
	}
	if err := c.Validate(req); err != nil {
		res := helper.APIResponse(err.Error(), http.StatusBadRequest, "error", nil)
		return c.JSON(http.StatusBadRequest, res)
	}
	// get token
	// token := middlewares.ExtractToken(c)
	// if token == nil {
	//	res := helper.APIResponse("invalid token", http.StatusBadRequest, "error", nil)
	//	return c.JSON(http.StatusBadRequest, res)
	//}
	userExisting, _ := database.FindByEmail(c, req.Email)
	// find existing
	if userExisting != nil {
		res := helper.APIResponse("User sudah ada di database", http.StatusBadRequest, "error", nil)
		return c.JSON(http.StatusBadRequest, res)
	}
	// karena create user tidak perlu autentikasi  maka untuk created_by dummy 1
	//req.UserId = token.UserId
	req.UserId = 1
	err := database.CreateUser(c, req)

	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.APIResponse("Bad request", http.StatusBadRequest, "error", nil))
	}
	res := helper.APIResponse("Create user success", http.StatusOK, "success", nil)

	return c.JSON(http.StatusOK, res)
}

func UpdateUserController(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	req := &request.UserCreate{}
	if err := c.Bind(req); err != nil {
		return err
	}
	if err := c.Validate(req); err != nil {
		res := helper.APIResponse(err.Error(), http.StatusBadRequest, "error", nil)
		return c.JSON(http.StatusBadRequest, res)
	}
	// get token
	token := middlewares.ExtractToken(c)
	if token == nil {
		res := helper.APIResponse("invalid token", http.StatusBadRequest, "error", nil)
		return c.JSON(http.StatusBadRequest, res)
	}

	// cek  record user will delete
	userExists, err := database.GetUserById(c, id)
	if err != nil {
		res := helper.APIResponse(err.Error(), http.StatusBadRequest, "error", nil)
		return c.JSON(http.StatusBadRequest, res)
	}

	// compareate
	if userExists.CreatedBy != token.UserId {
		res := helper.APIResponse("User tidak diperbolehkan menghapus item ID "+c.Param("id")+" ini ", http.StatusBadRequest, "error", nil)
		return c.JSON(http.StatusBadRequest, res)
	}

	users, err := database.UpdateUser(c, req, id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.APIResponse(err.Error(), http.StatusBadRequest, "error", nil))
	}
	res := helper.APIResponse("Update user success", http.StatusOK, "success", users)

	return c.JSON(http.StatusOK, res)
}

func DeleteUserIdController(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	// get token
	token := middlewares.ExtractToken(c)
	if token == nil {
		res := helper.APIResponse("invalid token", http.StatusBadRequest, "error", nil)
		return c.JSON(http.StatusBadRequest, res)
	}

	// cek  record user will delete
	userExists, err := database.GetUserById(c, id)
	if err != nil {
		res := helper.APIResponse(err.Error(), http.StatusBadRequest, "error", nil)
		return c.JSON(http.StatusBadRequest, res)
	}
	if userExists.CreatedBy != token.UserId {
		res := helper.APIResponse("User tidak diperbolehkan menghapus item ID "+c.Param("id")+" ini ", http.StatusBadRequest, "error", nil)
		return c.JSON(http.StatusBadRequest, res)
	}

	err = database.DeleteUserById(c, id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.APIResponse(err.Error(), http.StatusBadRequest, "error", nil))
	}
	res := helper.APIResponse("Delete user success", http.StatusOK, "success", nil)

	return c.JSON(http.StatusOK, res)
}

func LoginUserController(c echo.Context) error {
	req := &request.Login{}
	var err error
	if err = c.Bind(req); err != nil {
		res := helper.APIResponse(err.Error(), http.StatusInternalServerError, "error", nil)
		return c.JSON(http.StatusInternalServerError, res)
	}

	if err = c.Validate(req); err != nil {
		res := helper.APIResponse(err.Error(), http.StatusBadRequest, "error", nil)
		return c.JSON(http.StatusBadRequest, res)
	}
	// login user
	user, err := database.FindUserAuth(c, req)
	if err != nil {
		res := helper.APIResponse(err.Error(), http.StatusBadRequest, "error", nil)
		return c.JSON(http.StatusBadRequest, res)
	}
	usTrans := transform.UserTransform(user)
	usTrans.Token, err = middlewares.GenerateToken(*user)
	res := helper.APIResponse("login berhasil", http.StatusBadRequest, "success", usTrans)
	return c.JSON(http.StatusOK, res)
}
