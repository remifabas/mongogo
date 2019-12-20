package gateway_mongo

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	trainer "go.remifabas/mongogo/entity/trainer"
)

func GetMongoClient(url string) *mongo.Client {
	// Set client options
	clientOptions := options.Client().ApplyURI(url)
	// Connect to MongoDB
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	// Check the connection
	err = client.Ping(context.TODO(), nil)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected to MongoDB!")
	return client
}

func FindInCollection(c *mongo.Collection) string {
	// create a value into which the result can be decoded
	var result trainer.Trainer
	filter := bson.D{{"name", "Ash"}}

	err := c.FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Found a single document: %+v\n", result)

	out, err := json.Marshal(result)
	if err != nil {
		log.Fatal(err)
	}

	return string(out)
}
