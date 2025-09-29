package model

import (
	"bookstore/model"
	"context"

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

func GetBooksCollection()([]model.Book,error) {
	var fetchedbooks []model.Book
	cursor,err:=BookCollection.Find(context.TODO(),bson.M{})
    if err != nil{
		return nil,err
	}
    defer clos
}
