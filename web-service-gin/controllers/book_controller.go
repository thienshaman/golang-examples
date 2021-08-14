package controllers

import (
	"net/http"

	"example.com/web-service-gin/repositories"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)

func GetAllBook(c *gin.Context) {
	books, err := repositories.FindAllBook()
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, bson.D{{"error: ", err.Error()}})
	} else {
		c.IndentedJSON(http.StatusOK, books)
	}

}

func GetBook(c *gin.Context) {
	id := c.Query("id")
	book, err := repositories.FindBookById(id)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, bson.D{{"error: ", err.Error()}})
	} else {
		c.IndentedJSON(http.StatusOK, book)

	}

}
