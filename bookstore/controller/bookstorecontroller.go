package controller

import "go.mongodb.org/mongo-driver/mongo"

// Package-level variable to store the books collection
var BookCollection *mongo.Collection

// InitCollections initializes MongoDB collections
func InitCollections(client *mongo.Client, database *mongo.Database) {

	BookCollection = database.Collection("books")
}
