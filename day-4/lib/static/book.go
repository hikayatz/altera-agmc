package static

import (
	"errors"
	"fmt"

	"github.com/hikayat13/alterra-agcm/day-2/submission/models"
)

var Books = []models.Book{}

func init() {

	for i := 0; i < 10; i++ {
		book := models.Book{
			ID:     i,
			Title:  fmt.Sprintf("Title %d", i),
			Author: fmt.Sprintf("Author %d", i),
		}
		Books = append(Books, book)
	}
}

func GetBooks() (interface{}, error) {
	return Books, nil
}

func GetBookByIndex(index int) (interface{}, error) {
	if index > len(Books) {
		return nil, errors.New("index ditemukan")
	}
	return Books[index], nil
}

func AddBook(book *models.Book) error {
	newBook := models.Book{}
	newBook.ID = len(Books) + 1
	newBook.Title = book.Title
	newBook.Author = book.Author
	newBook.Isbn = book.Isbn
	Books = append(Books, newBook)

	return nil
}

func DeleteBook(index int) error {
	if index > len(Books) {
		return errors.New("index ditemukan")
	}
	Books = removeIndex(Books, index)
	return nil
}

func removeIndex(s []models.Book, index int) []models.Book {
	return append(s[:index], s[index+1:]...)
}
