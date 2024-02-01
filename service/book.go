package service

import (
	"github.com/raflyfarhandika/bookstore/app"
	"github.com/raflyfarhandika/bookstore/models"
	"github.com/raflyfarhandika/bookstore/repository"
)

type BookService struct {
	Repository *repository.BookRepo
}

func NewBookService(Repository *repository.BookRepo) *BookService {
	return &BookService{Repository}
}

func (s *BookService) GetAllBook() app.Response {
	result, err := s.Repository.FindAll()

	if err != nil {
		return app.Response{
			Error: true,
			Data: map[string]interface{}{
				"message": "There is something error",
			},
		}
	}

	if len(result) == 0 {
		return app.Response{
			Error: true,
			Data: map[string]interface{}{
				"message": "No data records",
			},
		}
	}

	return app.Response{
		Error: false,
		Data:  result,
	}
}

func (s *BookService) GetBookById(id uint) app.Response {
	result, err := s.Repository.FindOne(id)

	if err != nil {
		return app.Response{
			Error: true,
			Data: map[string]interface{}{
				"message": "Data Not Found!",
			},
		}
	}

	return app.Response{
		Error: false,
		Data:  result,
	}
}

func (s *BookService) CreateBook(book []models.Book) app.Response {
	result, err := s.Repository.Create()

	if err != nil {
		return app.Response{
			Error: true,
			Data: map[string]interface{}{
				"message": "Error to create data",
			},
		}
	}

	return app.Response{
		Error: false,
		Data:  result,
	}
}

func (s *BookService) DeleteBook(id uint) app.Response {
	book, err := s.Repository.FindOne(id)

	if err != nil {
		return app.Response{
			Error: true,
			Data: map[string]interface{}{
				"message": "Error to delete data",
			},
		}
	}

	err = s.Repository.Delete(book)

	if err != nil {
		return app.Response{
			Error: true,
			Data: map[string]interface{}{
				"message": "Error to delete data",
			},
		}
	}

	return app.Response{
		Error: false,
	}
}

func (s *BookService) UpdateBook(id uint, newBook []models.Book) app.Response {
	oldBook, err := s.Repository.FindOne(id)

	if err != nil {
		return app.Response{
			Error: true,
			Data: map[string]interface{}{
				"message": "Data Not Found!",
			},
		}
	}

	result, err := s.Repository.Update(oldBook, newBook)

	if err != nil {
		return app.Response{
			Error: true,
			Data: map[string]interface{}{
				"message": "Error update book data",
			},
		}
	}

	return app.Response{
		Error: false,
		Data:  result,
	}
}
