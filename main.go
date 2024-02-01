package main

import (
	"log"

	"github.com/raflyfarhandika/bookstore/database"
	"github.com/raflyfarhandika/bookstore/router"
)

func main() {

	database.InitDB()
	database.MigrateDB()
	route := router.StartApp()

	// r.GET("/", func(c *gin.Context) {
	// 	c.JSON(200, gin.H{
	// 		"message": "Hello World!",
	// 	})
	// })

	if err := route.Run(":3000"); err != nil {
		log.Fatal(err)
	}
}
