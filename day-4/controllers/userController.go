package controllers

import (
	"errors"
	"fmt"
	res "github.com/hikayat13/alterra-agcm/day-2/submission/util/response"
	"net/http"
	"strconv"

	"github.com/hikayat13/alterra-agcm/day-2/submission/middlewares"
	"github.com/hikayat13/alterra-agcm/day-2/submission/request"
	"github.com/hikayat13/alterra-agcm/day-2/submission/transform"

	"github.com/hikayat13/alterra-agcm/day-2/submission/lib/database"
	"github.com/labstack/echo/v4"
)

func GetUserController(c echo.Context) error {
	//token := middlewares.ExtractToken(c)
	//if token == nil {
	//	return res.CustomErrorBuilder(http.StatusBadRequest, res.S_INVALID_REQUEST, "Token invalid").Send(c)
	//}
	users, err := database.GetUsers(c)
	if err != nil {
		return res.ErrorBuilder(&res.ErrorConstant.InternalServerError, err).Send(c)
	}
	return res.CustomSuccessBuilder(http.StatusOK, users, "Get All users").Send(c)
}

func GetUserByIdController(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return res.ErrorBuilder(&res.ErrorConstant.ConvertionNotValid, err).Send(c)
	}
	users, err := database.GetUserById(c, id)
	if errors.Is(err, res.S_NOT_FOUND) {
		return res.ErrorBuilder(&res.ErrorConstant.NotFound, err).Send(c)
	}
	if err != nil {
		return res.ErrorBuilder(&res.ErrorConstant.InternalServerError, err).Send(c)
	}
	return res.CustomSuccessBuilder(http.StatusOK, users, "Get user by id ").Send(c)

}

func CreateUserController(c echo.Context) error {
	req := &request.UserCreate{}

	if err := c.Bind(req); err != nil {
		return res.ErrorBuilder(&res.ErrorConstant.UnprocessableEntity, err).Send(c)
	}
	if err := c.Validate(req); err != nil {
		return res.ErrorBuilder(&res.ErrorConstant.BadRequest, err).Send(c)
	}
	// get token
	// token := middlewares.ExtractToken(c)
	// if token == nil {
	//	res := helper.APIResponse("invalid token", http.StatusBadRequest, "error", nil)
	//	return c.JSON(http.StatusBadRequest, res)
	//}
	userExisting, err := database.FindByEmail(c, req.Email)
	if err != nil && !errors.Is(err, res.S_NOT_FOUND) {
		return res.ErrorBuilder(&res.ErrorConstant.InternalServerError, err).Send(c)
	}
	// find existing
	if userExisting != nil {
		return res.CustomErrorBuilder(http.StatusBadRequest, res.S_INVALID_REQUEST, "user sudah digunakan").Send(c)
	}
	// karena create user tidak perlu autentikasi  maka untuk created_by dummy 1
	//req.UserId = token.UserId
	req.UserId = 1
	err = database.CreateUser(c, req)
	if err != nil {
		return res.ErrorBuilder(&res.ErrorConstant.InternalServerError, err).Send(c)
	}
	return res.SuccessBuilder(&res.SuccessConstant.OK, nil).Send(c)
}

func UpdateUserController(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return res.ErrorBuilder(&res.ErrorConstant.ConvertionNotValid, err).Send(c)
	}
	req := &request.UserCreate{}
	if err := c.Bind(req); err != nil {
		return res.ErrorBuilder(&res.ErrorConstant.UnprocessableEntity, err).Send(c)
	}
	if err := c.Validate(req); err != nil {
		return res.ErrorBuilder(&res.ErrorConstant.BadRequest, err).Send(c)
	}
	// get token
	token := middlewares.ExtractToken(c)
	if token == nil {
		return res.ErrorBuilder(&res.ErrorConstant.BadRequest, errors.New("Token invalid")).Send(c)
	}

	// cek  record user will delete
	userExists, err := database.GetUserById(c, id)
	if errors.Is(err, res.S_NOT_FOUND) {
		return res.ErrorBuilder(&res.ErrorConstant.NotFound, err).Send(c)
	}
	if err != nil {
		return res.ErrorBuilder(&res.ErrorConstant.InternalServerError, err).Send(c)
	}

	// compareate
	if userExists.CreatedBy != token.UserId {
		return res.CustomErrorBuilder(http.StatusBadRequest, res.S_INVALID_REQUEST, fmt.Sprintf("userID  yang diperbolehkan => %d hapus data", userExists.CreatedBy)).Send(c)
	}

	user, err := database.UpdateUser(c, req, id)
	if err != nil {
		return res.ErrorBuilder(&res.ErrorConstant.InternalServerError, err).Send(c)
	}
	return res.SuccessBuilder(&res.SuccessConstant.OK, user).Send(c)
}

func DeleteUserIdController(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return res.ErrorBuilder(&res.ErrorConstant.ConvertionNotValid, err).Send(c)
	}
	// get token
	token := middlewares.ExtractToken(c)
	if token == nil {
		return res.ErrorBuilder(&res.ErrorConstant.BadRequest, errors.New("Token invalid")).Send(c)
	}

	// cek  record user will delete
	userExists, err := database.GetUserById(c, id)
	if err != nil {
		return res.ErrorBuilder(&res.ErrorConstant.NotFound, err).Send(c)
	}
	if err != nil {
		return res.ErrorBuilder(&res.ErrorConstant.InternalServerError, err).Send(c)
	}
	if userExists.CreatedBy != token.UserId {
		return res.CustomErrorBuilder(http.StatusBadRequest, res.S_INVALID_REQUEST, fmt.Sprintf("userID  yang diperbolehkan => %d hapus data", userExists.CreatedBy)).Send(c)
	}

	err = database.DeleteUserById(c, id)
	if err != nil {
		return res.ErrorBuilder(&res.ErrorConstant.InternalServerError, err).Send(c)
	}
	return res.SuccessBuilder(&res.SuccessConstant.OK, nil).Send(c)
}

func LoginUserController(c echo.Context) error {
	req := &request.Login{}
	var err error
	if err = c.Bind(req); err != nil {
		return res.ErrorBuilder(&res.ErrorConstant.UnprocessableEntity, err).Send(c)
	}

	if err = c.Validate(req); err != nil {
		return res.ErrorBuilder(&res.ErrorConstant.BadRequest, err).Send(c)
	}
	// login user
	user, err := database.FindUserAuth(c, req)
	if errors.Is(err, res.S_NOT_FOUND) {
		return res.ErrorBuilder(&res.ErrorConstant.EmailOrPasswordIncorrect, err).Send(c)
	}

	if err != nil {
		return res.ErrorBuilder(&res.ErrorConstant.InternalServerError, err).Send(c)
	}
	usTrans := transform.UserTransform(user)
	usTrans.Token, _ = middlewares.GenerateToken(*user)
	return res.SuccessBuilder(&res.SuccessConstant.OK, usTrans).Send(c)

}