package repositories

import (
	"context"

	"example.com/web-service-gin/entities"
	"example.com/web-service-gin/helpers"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func FindAllBook() ([]entities.Book, error) {
	client := helpers.InitConnection()
	defer client.Disconnect(context.TODO())
	db := client.Database("mydb")
	bookCollection := db.Collection("book")
	var books []entities.Book
	cur, err := bookCollection.Find(context.TODO(), bson.D{{}}, options.Find())
	if err != nil {
		return books, err
	}
	defer cur.Close(context.TODO())
	for cur.Next(context.TODO()) {
		var book entities.Book
		err := cur.Decode(&book)
		if err != nil {
			return books, err
		}
		books = append(books, book)
	}

	return books, nil
}

func FindBookById(id string) (entities.Book, error) {
	client := helpers.InitConnection()
	defer client.Disconnect(context.TODO())
	db := client.Database("mydb")
	bookCollection := db.Collection("book")

	var book entities.Book
	objId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return book, err
	}
	err = bookCollection.FindOne(context.TODO(), bson.D{{"_id", objId}}).Decode(&book)
	if err != nil {
		return book, err
	}

	return book, nil
}
