package controller

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/raflyfarhandika/bookstore/models"
	"github.com/raflyfarhandika/bookstore/service"
)

type BookController struct {
	Service *service.BookService
}

func NewBookController(Service *service.BookService) *BookController {
	return &BookController{Service}
}

func (c *BookController) GetAllBook(ctx *gin.Context) {
	response := c.Service.GetAllBook()

	if response.Error {
		ctx.JSON(http.StatusBadRequest, response.Data)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "Get all data success",
		"data":    response.Data,
	})
}

func (c *BookController) GetBookById(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("bookId"))

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "Something error",
		})
		return
	}

	response := c.Service.GetBookById(uint(id))

	if response.Error {
		ctx.JSON(http.StatusBadRequest, response.Data)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "Get data success",
		"data":    response.Data,
	})
}

func (c *BookController) InputNewBook(ctx *gin.Context) {
	var request models.Book

	err := ctx.ShouldBindJSON(&request)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "Something error",
		})
		return
	}

	response := c.Service.CreateBook(request)

	if response.Error {
		ctx.JSON(http.StatusBadRequest, response.Data)
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"code":    201,
		"message": "Create data success",
		"data":    response.Data,
	})
}

func (c *BookController) DeleteBook(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("bookId"))

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "Something error",
		})
		return
	}

	response := c.Service.DeleteBook(uint(id))

	if response.Error {
		ctx.JSON(http.StatusBadRequest, response.Data)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "Delete data success",
	})
}

func (c *BookController) UpdateBook(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("bookId"))

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "Something error",
		})
		return
	}

	var request []models.Book

	err = ctx.ShouldBindJSON(&request)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "Something error",
		})
		return
	}

	response := c.Service.UpdateBook(uint(id), request)

	if response.Error {
		ctx.JSON(http.StatusBadRequest, response.Data)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "Get data success",
		"data":    response.Data,
	})
}
