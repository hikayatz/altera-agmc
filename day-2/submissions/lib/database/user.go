package database

import (
	"errors"

	"github.com/hikayat13/alterra-agcm/day-2/submission/config"
	"github.com/hikayat13/alterra-agcm/day-2/submission/models"
	"github.com/hikayat13/alterra-agcm/day-2/submission/transform"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func GetUsers(ctx echo.Context) (interface{}, error) {
	var users []models.User
	if e := config.DB.WithContext(ctx.Request().Context()).Find(&users).Error; e != nil {
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

func CreateUser(ctx echo.Context, user models.User) error {

	tx := config.DB.WithContext(ctx.Request().Context()).Begin()
	if err := tx.Session(&gorm.Session{FullSaveAssociations: true}).Save(user).Error; err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()

	return nil
}

func UpdateUser(ctx echo.Context, user *models.User, id int) (*transform.User, error) {
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
	tx := config.DB.WithContext(ctx.Request().Context()).Begin()
	if err := tx.Session(&gorm.Session{FullSaveAssociations: true}).Delete(&models.User{}, id).Error; err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()

	return nil
}
