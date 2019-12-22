package main

import (
	"context"
	"fmt"
	"log"

	mongo_client "github.com/remifabas/mongogo/gateways/mongo"
)

// You will be using this Trainer type later in the program
type Trainer struct {
	Name string
	Age  int
	City string
}

func main() {
	// Set client options
	client := mongo_client.GetMongoClient("mongodb://localhost:27017")

	collection := client.Database("test").Collection("trainers")

	ash := Trainer{"Ash", 10, "Pallet Town"}
	insertResult, err := collection.InsertOne(context.TODO(), ash)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Inserted a single document: ", insertResult.InsertedID)

	result := mongo_client.FindInCollection(collection)
	fmt.Println("result \n" + result)
	err = client.Disconnect(context.TODO())

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connection to MongoDB closed.")
}
