package routes

import (
	"github.com/hikayat13/alterra-agcm/day-2/submission/controllers"
	"github.com/hikayat13/alterra-agcm/day-2/submission/middlewares"
	"github.com/labstack/echo/v4"
)

func New() *echo.Echo {
	e := echo.New()
	e.GET("users", controllers.GetUserController, middlewares.JWTMiddleware())
	e.GET("users/:id", controllers.GetUserByIdController)
	e.POST("users", controllers.CreateUserController)
	e.PUT("users/:id", controllers.UpdateUserController, middlewares.JWTMiddleware())
	e.DELETE("users/:id", controllers.DeleteUserIdController, middlewares.JWTMiddleware())
	e.POST("login", controllers.LoginUserController)

	e.GET("books", controllers.GetBooksController)
	e.GET("books/:id", controllers.GetBookByIdController)
	e.POST("books", controllers.CreateBookController, middlewares.JWTMiddleware())
	e.PUT("books/:id", controllers.UpdateBookController, middlewares.JWTMiddleware())
	e.DELETE("books/:id", controllers.DeleteBookIdController, middlewares.JWTMiddleware())

	return e
}
