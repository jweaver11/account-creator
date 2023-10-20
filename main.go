package main

import (
	"ac/helpers"
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {

	// Context handles information to functions that handle timeouts, cancellations, etc.
	// Sets our context to 10 seconds of failure, then backs out of program
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	clientOptions := options.Client().ApplyURI("mongodb+srv://jweaver:1232@cluster0.tdnjlbj.mongodb.net/?retryWrites=true&w=majority")

	// Sets our mongo client that handles connections and
	// connects to our desired cluster
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(ctx)

	// Goes into this database or creates it, with the collection name
	collection := client.Database("john").Collection("cena")

	// Inserts our kvp with the name goodbye and the value square
	res, err2 := collection.InsertOne(context.Background(), bson.M{"goodbye": "square"})
	if err2 != nil {
		log.Fatal(err2)
	}

	// Sets the id of the just inserted data
	id := res.InsertedID
	log.Println(id)

	// Prompts user to input username and password
	fmt.Println("Please put in your username and password: ")

	// Calls inputlogin function to accept user input
	helpers.InputLogin()

	/*router := gin.Default()

	fmt.Println(router)*/

}
