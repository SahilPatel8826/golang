package model

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// Book represents a book document in MongoDB
type Book struct {
	ID       string `json:"id" bson:"_id,omitempty"`
	Title    string `json:"title" bson:"title"`
	Quantity int    `json:"quantity" bson:"quantity"`
	Price    int64  `json:"price" bson:"price"`
}

// Package-level variable to store the books collection
var BookCollection *mongo.Collection

// InitCollections initializes MongoDB collections
func InitCollections(client *mongo.Client, dbName string) {
	db := client.Database(dbName)
	BookCollection = db.Collection("books")
}

func GetBookCollection() ([]Book, error) {
	var books []Book

	ctx := context.TODO()
	cursor, err := BookCollection.Find(ctx, bson.M{})
	if err != nil {
		log.Println("books not found:", err)
		return nil, err
	}
	defer cursor.Close(ctx)

	// Loop through documents
	for cursor.Next(ctx) {
		var book Book
		if err := cursor.Decode(&book); err != nil {
			log.Println("error decoding book:", err)
			return nil, err
		}
		books = append(books, book)
	}

	// Check for cursor errors
	if err := cursor.Err(); err != nil {
		return nil, err
	}

	return books, nil
}
func CreateBook(book Book) (Book, error) {
	inserted, err := BookCollection.InsertOne(context.TODO(), book)
	if err != nil {
		log.Fatal("not insertred", err)
	}

	fmt.Println("inserted", inserted.InsertedID)
	return book, nil
}
func CreateManyBooks(books []interface{}) ([]interface{}, error) {
	// docs := make([]interface{}, len(books))
	// for i, b := range books {
	// 	docs[i] = b

	result, err := BookCollection.InsertMany(context.TODO(), books)
	if err != nil {
		log.Fatal("not insertred", err)
	}
	return result.InsertedIDs, nil
}
