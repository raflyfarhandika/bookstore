package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/raflyfarhandika/bookstore/database"
)

func main() {

	database.InitDB()
	database.MigrateDB()

	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello World!",
		})
	})

	if err := r.Run(":3000"); err != nil {
		log.Fatal(err)
	}
}
