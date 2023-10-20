package helpers

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func WriteDB(client *mongo.Client) {
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
}
