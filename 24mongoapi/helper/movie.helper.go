package helper

import (
	"context"
	"fmt"
	"log"

	"github.com/papihack/mongoapi/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func InsertOneMovie(collection *mongo.Collection, movie model.Netflix) {
	result, err := collection.InsertOne(context.Background(), movie)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("[*] Inserted 1 movie in DB with id:", result.InsertedID)
}

func UpdateOneMovie(collection *mongo.Collection, movieId string) int64 {
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
	return result.ModifiedCount
}

func DeleteOneMovie(collection *mongo.Collection, movieId string) int64 {
	id, err := primitive.ObjectIDFromHex(movieId)
	if err != nil {
		log.Fatal(err)
	}
	filter := bson.M{"_id": id}
	result, err := collection.DeleteOne(context.Background(), filter)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("[*] Deleted count:", result.DeletedCount)
	return result.DeletedCount
}

func DeleteAllMovie(collection *mongo.Collection) int64 {
	result, err := collection.DeleteMany(context.Background(), bson.D{{}}, nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("[*] Number of movies deleted:", result.DeletedCount)
	return result.DeletedCount
}

func GetAllMovies(collection *mongo.Collection) []primitive.M {
	fmt.Println("[*] Get All Movie")
	cursor, err := collection.Find(context.Background(), bson.D{{}})
	
	if err != nil {
		log.Fatal(err)
	}
	
	var movies []primitive.M

	for cursor.Next(context.Background()) {
		var movie bson.M
		er := cursor.Decode(&movie)
		if er != nil {
			log.Fatal(er)
		}
		movies = append(movies, movie)
	}

	return movies
}

func GetOneMovie(collection *mongo.Collection, movieId string) model.Netflix {
	fmt.Println("[*] Get Movie of Id: #", movieId)
	id, err := primitive.ObjectIDFromHex(movieId)
	var movie model.Netflix
	if err != nil {
		log.Fatal(err)
	}
	filter := bson.M{"_id": id}
	result := collection.FindOne(context.Background(), filter)
	_ = result.Decode(&movie)
	fmt.Println("[*] Get Movie of Id:", movieId)
	return movie
}

