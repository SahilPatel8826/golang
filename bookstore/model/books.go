package model

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// Book represents a book document in MongoDB
type Book struct {
	ID       primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Title    string             `json:"title" bson:"title"`
	Quantity int                `json:"quantity" bson:"quantity"`
	Price    int64              `json:"price" bson:"price"`
}

var BookCollection *mongo.Collection

// InitCollections initializes MongoDB collections
func InitCollections(client *mongo.Client, dbName string) {
	db := client.Database(dbName)
	BookCollection = db.Collection("books")
}

// GetAllBooks fetches all books
func GetAllBooks() ([]Book, error) {
	var books []Book
	ctx := context.TODO()

	cursor, err := BookCollection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var book Book
		if err := cursor.Decode(&book); err != nil {
			return nil, err
		}
		books = append(books, book)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	return books, nil
}

// CreateBook inserts one book
func CreateBook(book Book) (Book, error) {
	book.ID = primitive.NewObjectID()
	_, err := BookCollection.InsertOne(context.TODO(), book)
	if err != nil {
		return Book{}, err
	}
	return book, nil
}

// CreateManyBooks inserts multiple books
func CreateManyBooks(books []Book) ([]interface{}, error) {
	var docs []interface{}
	for _, b := range books {
		if b.ID.IsZero() {
			b.ID = primitive.NewObjectID()
		}
		docs = append(docs, b)
	}

	result, err := BookCollection.InsertMany(context.TODO(), docs)
	if err != nil {
		return nil, err
	}
	return result.InsertedIDs, nil
}

// DeleteOneBook deletes one book by ID
func DeleteOneBook(id string) (int64, error) {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return 0, fmt.Errorf("invalid book ID: %v", err)
	}

	result, err := BookCollection.DeleteOne(context.TODO(), bson.M{"_id": objID})
	if err != nil {
		return 0, err
	}
	return result.DeletedCount, nil
}

// GetOneBook fetches a book by ID
func GetOneBook(id string) (*Book, error) {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, fmt.Errorf("invalid book ID: %v", err)
	}

	var book Book
	err = BookCollection.FindOne(context.TODO(), bson.M{"_id": objID}).Decode(&book)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, fmt.Errorf("book not found")
		}
		return nil, err
	}

	return &book, nil
}
func DeleteAllBooks() (int64, error) {
	result, err := BookCollection.DeleteMany(context.TODO(), bson.M{})
	if err != nil {
		fmt.Errorf("unable to delete")
	}
	return result.DeletedCount, err
}

func UpdateOne(id string, book Book) (*Book, error) {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, fmt.Errorf("invalid book ID: %v", err)
	}

	filter := bson.M{"_id": objID}
	update := bson.M{
		"$set": bson.M{
			"title":    book.Title,
			"quantity": book.Quantity,
			"price":    book.Price,
		},
	}
	result, err := BookCollection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		return nil, err
	}

	if result.MatchedCount == 0 {
		return nil, fmt.Errorf("book not found")
	}

	// (Optional) Fetch the updated book from DB to return the latest data
	var updatedBook Book
	err = BookCollection.FindOne(context.TODO(), filter).Decode(&updatedBook)
	if err != nil {
		return nil, err
	}

	return &updatedBook, nil
}
