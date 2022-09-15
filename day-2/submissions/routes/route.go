package routes

import (
	"github.com/hikayat13/alterra-agcm/day-2/submission/controllers"
	"github.com/labstack/echo/v4"
)

func New() *echo.Echo {
	e := echo.New()
	e.GET("users", controllers.GetUserController)
	e.GET("users/:id", controllers.GetUserByIdController)
	e.POST("users", controllers.CreateUserController)
	e.PUT("users/:id", controllers.UpdateUserController)
	e.DELETE("users/:id", controllers.DeleteUserIdController)

	e.GET("books", controllers.GetBooksController)
	e.GET("books/:id", controllers.GetBookByIdController)
	e.POST("books", controllers.CreateBookController)
	e.PUT("books/:id", controllers.UpdateBookController)
	e.DELETE("books/:id", controllers.DeleteBookIdController)

	return e
}
