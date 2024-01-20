// main.go
package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Book model
type Book struct {
	ID       string    `json:"id,omitempty" bson:"_id,omitempty"`
	Title    string    `json:"title,omitempty" bson:"title,omitempty"`
	Author   string    `json:"author,omitempty" bson:"author,omitempty"`
	Isbn     string    `json:"isbn,omitempty" bson:"isbn,omitempty"`
	Released time.Time `json:"released,omitempty" bson:"released,omitempty"`
}

var client *mongo.Client

func main() {
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	client, _ = mongo.Connect(context.Background(), clientOptions)

	router := mux.NewRouter()

	router.HandleFunc("/books", getBooks).Methods("GET")
	router.HandleFunc("/book/{id}", getBook).Methods("GET")
	router.HandleFunc("/book", createBook).Methods("POST")
	router.HandleFunc("/book/{id}", updateBook).Methods("PUT")
	router.HandleFunc("/book/{id}", deleteBook).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8000", router))
}

func getBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var books []Book
	collection := client.Database("mydatabase").Collection("books")
	cursor, err := collection.Find(context.Background(), bson.M{})
	if err != nil {
		log.Fatal(err)
	}
	defer cursor.Close(context.Background())

	for cursor.Next(context.Background()) {
		var book Book
		if err := cursor.Decode(&book); err != nil {
			log.Fatal(err)
		}
		books = append(books, book)
	}

	if err := json.NewEncoder(w).Encode(books); err != nil {
		log.Fatal(err)
	}
}

func getBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)
	var book Book

	collection := client.Database("mydatabase").Collection("books")
	err := collection.FindOne(context.Background(), bson.M{"_id": params["id"]}).Decode(&book)
	if err != nil {
		log.Fatal(err)
	}

	if err := json.NewEncoder(w).Encode(book); err != nil {
		log.Fatal(err)
	}
}

func createBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var book Book
	if err := json.NewDecoder(r.Body).Decode(&book); err != nil {
		log.Fatal(err)
	}

	collection := client.Database("mydatabase").Collection("books")
	result, err := collection.InsertOne(context.Background(), book)
	if err != nil {
		log.Fatal(err)
	}

	if err := json.NewEncoder(w).Encode(result.InsertedID); err != nil {
		log.Fatal(err)
	}
}

func updateBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)
	var book Book
	if err := json.NewDecoder(r.Body).Decode(&book); err != nil {
		log.Fatal(err)
	}

	collection := client.Database("mydatabase").Collection("books")
	filter := bson.M{"_id": params["id"]}
	update := bson.M{"$set": book}

	_, err := collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		log.Fatal(err)
	}

	if err := json.NewEncoder(w).Encode(book); err != nil {
		log.Fatal(err)
	}
}

func deleteBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)

	collection := client.Database("mydatabase").Collection("books")
	filter := bson.M{"_id": params["id"]}

	_, err := collection.DeleteOne(context.Background(), filter)
	if err != nil {
		log.Fatal(err)
	}

	w.Write([]byte("Kitap silindi"))
}
