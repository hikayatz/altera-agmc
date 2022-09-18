package controllers

import (
	"log"
	"net/http"
	"testing"

	"github.com/joho/godotenv"

	"github.com/hikayat13/alterra-agcm/day-2/submission/config"
	"github.com/hikayat13/alterra-agcm/day-2/submission/database/seeder"
	"github.com/hikayat13/alterra-agcm/day-2/submission/mocks"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func init() {
	err := godotenv.Load("../.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	config.InitDB()
}
func TestLoginInvidReqest(t *testing.T) {
	// setup database
	config.InitDB()
	seeder.NewSeeder().DeleteAll()
	seeder.NewSeeder().SeedAll()
	asserts := assert.New(t)

	// setup context
	e := echo.New()
	echoMock := mocks.EchoMock{E: e}
	c, rec := echoMock.RequestMock(http.MethodPost, "/", nil)
	c.SetPath("/login")

	// testing
	if asserts.NoError(LoginUserController(c)) {
		asserts.Equal(400, rec.Code)
	}
}