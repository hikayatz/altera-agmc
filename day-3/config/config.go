package config

import (
	"fmt"
	"github.com/go-playground/validator"
	"log"
	"os"
	"time"

	gormLogger "gorm.io/gorm/logger"

	"github.com/hikayat13/alterra-agcm/day-2/submission/models"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {

	newLogger := gormLogger.Default.LogMode(gormLogger.Warn)

	envApp := os.Getenv("APP_ENV")
	if envApp != "production" {
		newLogger = gormLogger.New(
			log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
			gormLogger.Config{
				SlowThreshold:             200 * time.Second, // Slow SQL threshold
				LogLevel:                  gormLogger.Info,   // Log level
				IgnoreRecordNotFoundError: true,              // Ignore ErrRecordNotFound error for logger
				Colorful:                  true,              // Disable color
			},
		)
	}

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", os.Getenv("DB_USERNAME"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_NAME"))

	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{Logger: newLogger})

	if err != nil {
		panic(err)
	}

	InitMigrate()
}

func InitMigrate() {
	e := DB.AutoMigrate(&models.User{})
	if e != nil {
		log.Fatalln(e.Error())
	}
}

func SetupLogger(e *echo.Echo) {
	envApp := os.Getenv("APP_ENV")
	if envApp != "production" {
		e.Use(middleware.Logger())
	}
}

type customValidator struct {
	validator *validator.Validate
}

func (cv *customValidator) Validate(i interface{}) error {
	if err := cv.validator.Struct(i); err != nil {
		// Optionally, you could return the error to give each route more control over the status code
		return err
	}
	return nil
}

func SetupValidator(e *echo.Echo) {
	e.Validator = &customValidator{validator: validator.New()}
}
