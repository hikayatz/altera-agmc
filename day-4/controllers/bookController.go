package controllers

import (
	helper "github.com/hikayat13/alterra-agcm/day-2/submission/util/response"
	"net/http"
	"strconv"

	"github.com/hikayat13/alterra-agcm/day-2/submission/lib/static"
	"github.com/hikayat13/alterra-agcm/day-2/submission/models"
	"github.com/labstack/echo/v4"
)

func GetBooksController(c echo.Context) error {
	books, _ := static.GetBooks()
	res := helper.APIResponse("Get all books", http.StatusOK, "success", books)

	return c.JSON(http.StatusOK, res)
}

func CreateBooksController(c echo.Context) error {
	book := &models.Book{}
	if err := c.Bind(book); err != nil {
		return err
	}
	err := static.AddBook(book)
	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.APIResponse(err.Error(), http.StatusBadRequest, "error", nil))
	}
	return c.JSON(http.StatusOK, static.AddBook)
}

func GetBookByIdController(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	book, err := static.GetBookByIndex(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.APIResponse(err.Error(), http.StatusBadRequest, "error", nil))
	}
	res := helper.APIResponse("Get book by idex", http.StatusOK, "success", book)

	return c.JSON(http.StatusOK, res)
}

func CreateBookController(c echo.Context) error {
	book := new(models.Book)
	if err := c.Bind(book); err != nil {
		return err
	}
	err := static.AddBook(book)
	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.APIResponse("Bad request", http.StatusBadRequest, "error", nil))
	}
	res := helper.APIResponse("Add book success", http.StatusOK, "success", nil)

	return c.JSON(http.StatusOK, res)
}

func DeleteBookIdController(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	err := static.DeleteBook(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.APIResponse(err.Error(), http.StatusBadRequest, "error", nil))
	}
	res := helper.APIResponse("Delete book by index success", http.StatusOK, "success", nil)

	return c.JSON(http.StatusOK, res)
}

func UpdateBookController(c echo.Context) error {
	// _, _ := strconv.Atoi(c.Param("id"))

	res := helper.APIResponse("Delete book by index success", http.StatusOK, "success", nil)

	return c.JSON(http.StatusOK, res)
}