package controllers

import (
	"bytes"
	"encoding/json"
	"github.com/hikayat13/alterra-agcm/day-2/submission/config"
	"github.com/hikayat13/alterra-agcm/day-2/submission/database/seeder"
	"github.com/hikayat13/alterra-agcm/day-2/submission/middlewares"
	"github.com/hikayat13/alterra-agcm/day-2/submission/mocks"
	"github.com/hikayat13/alterra-agcm/day-2/submission/request"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"log"
	"net/http"
	"testing"
)

func init() {
	err := godotenv.Load("../.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	config.InitDB()
}

func TestAuthHandlerLoginInvalidPayload(t *testing.T) {
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
		body := rec.Body.String()
		asserts.JSONEq(`{"Meta":{"message":"Bad Request","code":400,"status":"invalid_request"}}`, body)
	}
}
func TestAuthHandlerLoginPayloadIncorrect(t *testing.T) {
	// setup database
	config.InitDB()
	seeder.NewSeeder().DeleteAll()
	seeder.NewSeeder().SeedAll()
	asserts := assert.New(t)

	// setup context
	emailAndPassword := request.Login{
		Email:    "hikayat@gmail.com",
		Password: "1234567",
	}

	e := echo.New()
	echoMock := mocks.EchoMock{E: e}
	payload, err := json.Marshal(emailAndPassword)
	if err != nil {
		t.Fatal(err)
	}
	c, rec := echoMock.RequestMock(http.MethodPost, "/", bytes.NewBuffer(payload))
	c.Request().Header.Set("Content-Type", "application/json")
	c.SetPath("/login")

	// testing
	if asserts.NoError(LoginUserController(c)) {
		asserts.Equal(400, rec.Code)
		body := rec.Body.String()
		asserts.JSONEq(`{"Meta":{"message":"Email or password is incorrect","code":400,"status":"invalid_request"}}`, body)
	}
}

func TestAuthHandlerLoginPayloadSuccess(t *testing.T) {
	// setup database
	config.InitDB()
	seeder.NewSeeder().DeleteAll()
	seeder.NewSeeder().SeedAll()
	asserts := assert.New(t)

	// setup context
	emailAndPassword := request.Login{
		Email:    "hikayat@gmail.com",
		Password: "123456",
	}

	e := echo.New()
	echoMock := mocks.EchoMock{E: e}
	payload, err := json.Marshal(emailAndPassword)
	if err != nil {
		t.Fatal(err)
	}
	c, rec := echoMock.RequestMock(http.MethodPost, "/", bytes.NewBuffer(payload))
	c.Request().Header.Set("Content-Type", "application/json")
	c.SetPath("/login")

	// testing
	if asserts.NoError(LoginUserController(c)) {
		asserts.Equal(200, rec.Code)

		body := rec.Body.String()
		asserts.Contains(body, "id")
		asserts.Contains(body, "name")
		asserts.Contains(body, "email")
		asserts.Contains(body, "token")
	}
}

func TestCreateUserAlreadyExist(t *testing.T) {
	// setup database
	config.InitDB()
	seeder.NewSeeder().DeleteAll()
	seeder.NewSeeder().SeedAll()
	asserts := assert.New(t)

	// setup context
	user := request.UserCreate{
		Name:     "hikayat",
		Email:    "hikayat@gmail.com",
		Password: "123456",
	}

	e := echo.New()
	echoMock := mocks.EchoMock{E: e}
	payload, err := json.Marshal(user)
	if err != nil {
		t.Fatal(err)
	}
	c, rec := echoMock.RequestMock(http.MethodPost, "/", bytes.NewBuffer(payload))
	c.Request().Header.Set("Content-Type", "application/json")
	c.SetPath("/users")
	// testing
	if asserts.NoError(CreateUserController(c)) {
		asserts.Equal(400, rec.Code)

		body := rec.Body.String()
		asserts.Contains(body, "user sudah digunakan")
	}
}

func TestGetAllUserSuccess(t *testing.T) {
	// setup database
	config.InitDB()
	seeder.NewSeeder().DeleteAll()
	seeder.NewSeeder().SeedAll()
	asserts := assert.New(t)

	const token = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoyLCJuYW1lIjoiaGlrYXlhdCIsImVtYWlsIjoiaGlrYXlhdEBnbWFpbC5jb20iLCJleHAiOjE2NjM1MjAxMTZ9.eEX_m08vab8YqJ0qhh5EivJRtveUTW7M5e98XcpnzNI"

	e := echo.New()
	echoMock := mocks.EchoMock{E: e}

	c, rec := echoMock.RequestMock(http.MethodPost, "/", nil)
	c.Request().Header.Set("Content-Type", "application/json")
	c.Request().Header.Set("Authorization", "Bearer "+token)
	c.SetPath("/users")

	if asserts.NoError(GetUserController(c)) {
		asserts.Equal(200, rec.Code)

		body := rec.Body.String()
		asserts.Contains(body, "data")
	}

}

func TestGetAllUserInvalid(t *testing.T) {
	// setup database
	config.InitDB()
	seeder.NewSeeder().DeleteAll()
	seeder.NewSeeder().SeedAll()
	asserts := assert.New(t)

	const token = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoyLCJuYW1lIjoiaGlrYXlhdCIsImVtYWlsIjoiaGlrYXlhdEBnbWFpbC5jb20iLCJleHAiOjE2NjM1MjAxMTZ9.eEX_m08vab8YqJ0qhh5EivJRtveUTW7M5e98XcpnzNI"

	e := echo.New()
	echoMock := mocks.EchoMock{E: e}
	e.Use(middlewares.JWTMiddleware())

	c, rec := echoMock.RequestMock(http.MethodGet, "/", nil)
	c.Request().Header.Set("Content-Type", "application/json")
	c.Request().Header.Set(echo.HeaderAuthorization, "Bearer "+token)
	c.SetPath("/users")

	if asserts.NoError(GetUserController(c)) {
		asserts.Equal(200, rec.Code)

		body := rec.Body.String()
		asserts.Contains(body, "data")
	}

}