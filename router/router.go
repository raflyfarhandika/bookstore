package router

import (
	"github.com/gin-gonic/gin"
	"github.com/raflyfarhandika/bookstore/controller"
	"github.com/raflyfarhandika/bookstore/database"
	"github.com/raflyfarhandika/bookstore/repository"
	"github.com/raflyfarhandika/bookstore/service"
)

func StartApp() *gin.Engine {
	r := gin.Default()

	db := database.GetDB()
	repo := repository.NewBookRepo(db)
	service := service.NewBookService(repo)
	bookController := controller.NewBookController(service)

	api := r.Group("/api/v1")

	bookRoute := api.Group("/book")
	{
		bookRoute.GET("/", bookController.GetAllBook)
		bookRoute.GET("/:id", bookController.GetBookById)
		bookRoute.POST("/", bookController.InputNewBook)
		bookRoute.PUT("/:id", bookController.UpdateBook)
		bookRoute.DELETE("/:id", bookController.DeleteBook)
	}

	return r

}
