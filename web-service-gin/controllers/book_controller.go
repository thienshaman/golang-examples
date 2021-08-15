package controllers

import (
	"net/http"

	"example.com/web-service-gin/entities"
	"example.com/web-service-gin/repositories"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetAllBook(c *gin.Context) {
	books, err := repositories.FindAllBook()
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, entities.ErrorResponse{Message: err.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, entities.FindDataSuccessResponse{Total: len(books), Data: books})

}

func GetBook(c *gin.Context) {
	id := c.Query("id")
	book, err := repositories.FindBookById(id)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, entities.ErrorResponse{Message: err.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, book)

}

func AddBook(c *gin.Context) {
	var newBook entities.Book
	err := c.BindJSON(&newBook)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, entities.ErrorResponse{Message: err.Error()})
		return
	}
	bookId, err := repositories.AddBook(newBook)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, entities.ErrorResponse{Message: err.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, bson.D{primitive.E{Key: "id", Value: bookId}})

}

func UpdateBook(c *gin.Context) {
	var book entities.Book
	err := c.BindJSON(&book)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, entities.ErrorResponse{Message: err.Error()})
		return
	}
	updatedCount, err := repositories.UpdateBook(book)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, entities.ErrorResponse{Message: err.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, bson.D{primitive.E{Key: "number_record_updated", Value: updatedCount}})
}

func DeleteBook(c *gin.Context) {
	id := c.Query("id")
	deleteCount, err := repositories.DeleteBook(id)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, entities.ErrorResponse{Message: err.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, bson.D{primitive.E{Key: "number_record_deleted", Value: deleteCount}})
}
