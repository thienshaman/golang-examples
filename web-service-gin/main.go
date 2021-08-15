package main

import (
	"example.com/web-service-gin/controllers"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/books", controllers.GetAllBook)
	router.GET("/book", controllers.GetBook)
	router.POST("/addBook", controllers.AddBook)
	router.POST("/updateBook", controllers.UpdateBook)
	router.GET("/deleteBook", controllers.DeleteBook)

	router.Run("localhost:8080")

}
