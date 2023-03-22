package controller

import (
	"context"
	"fmt"
	"log"

	"github.com/papihack/mongoapi/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const connectionString = "mongo://localhost:27017"
const databaseName = "netflix"
const collectionName = "watchlist"

var collection *mongo.Collection

// connection with mongoDB

func init() {
	// client option
	clientOption := options.Client().ApplyURI(connectionString)

	// connect to mongoDB
	client, err := mongo.Connect(context.TODO(), clientOption)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("[*] MongoDB connection success.")

	collection = client.Database(databaseName).Collection(collectionName)

	fmt.Println("[*] Collection reference/instance is ready.")
}

// MongoDB helpers - file

func insertOneMovie(movie model.Netflix) {
	result, err := collection.InsertOne(context.Background(), movie)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("[*] Inserted 1 movie in DB with id:", result.InsertedID)
}

func updateOneMovie(movieId string) {
	id, err := primitive.ObjectIDFromHex(movieId)
	if err != nil {
		log.Fatal(err)
	}
	filter := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"watched": true}}

	result, err := collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("[*] Modified count:", result.ModifiedCount)
}
