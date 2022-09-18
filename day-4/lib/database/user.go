package database

import (
	"errors"

	"github.com/hikayat13/alterra-agcm/day-2/submission/request"

	"github.com/hikayat13/alterra-agcm/day-2/submission/config"
	"github.com/hikayat13/alterra-agcm/day-2/submission/models"
	"github.com/hikayat13/alterra-agcm/day-2/submission/transform"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func GetUsers(ctx echo.Context) (interface{}, error) {
	var users []models.User
	queryRaw := `SELECT id, name, email, photo  FROM users`
	if e := config.DB.Session(&gorm.Session{PrepareStmt: true, QueryFields: true}).WithContext(ctx.Request().Context()).Raw(queryRaw).Scan(&users).Error; e != nil {
		return nil, e
	}
	return users, nil
}

func GetUserById(ctx echo.Context, id int) (*transform.User, error) {
	var user models.User
	if e := config.DB.Session(&gorm.Session{PrepareStmt: true, QueryFields: true}).
		WithContext(ctx.Request().Context()).
		First(&user, id).Error; e != nil {
		return nil, e
	}

	return transform.UserTransform(&user), nil
}

func FindByEmail(ctx echo.Context, email string) (*models.User, error) {
	var user models.User
	var err error
	if err = config.DB.Session(&gorm.Session{PrepareStmt: true, QueryFields: true}).
		WithContext(ctx.Request().Context()).Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}

	return &user, nil
}

func CreateUser(ctx echo.Context, req *request.UserCreate) error {

	user := &models.User{
		Name:      req.Name,
		Email:     req.Email,
		Password:  req.Password,
		CreatedBy: req.UserId,
	}

	tx := config.DB.WithContext(ctx.Request().Context()).Begin()
	if err := tx.Session(&gorm.Session{FullSaveAssociations: true}).Save(user).Error; err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()

	return nil
}

func UpdateUser(ctx echo.Context, req *request.UserCreate, id int) (*transform.User, error) {
	user := models.User{
		Name:     req.Name,
		Email:    req.Email,
		Password: req.Password,
	}
	var result *models.User
	tx := config.DB.WithContext(ctx.Request().Context()).Begin()
	if err := tx.Session(&gorm.Session{FullSaveAssociations: true}).Clauses(clause.Returning{}).Where("id = ?", id).Updates(user).Scan(&result).Error; err != nil {
		tx.Rollback()
		return nil, err
	}
	tx.Commit()

	if result == nil {
		return nil, errors.New("record not found")
	}

	return transform.UserTransform(result), nil
}

func DeleteUserById(ctx echo.Context, id int) error {
	tx := config.DB.Unscoped().WithContext(ctx.Request().Context()).Begin()
	if err := tx.Session(&gorm.Session{FullSaveAssociations: true}).Delete(&models.User{}, id).Error; err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()

	return nil
}

func FindUserAuth(ctx echo.Context, req *request.Login) (*models.User, error) {
	user := &models.User{}
	var err error
	if err = config.DB.WithContext(ctx.Request().Context()).Where("email = ? AND password = ? ", req.Email, req.Password).First(&user).Error; err != nil {
		return nil, err
	}

	return user, nil
}