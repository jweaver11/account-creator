package migrations

import (
	"ac/models"
	"context"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func WriteData(client *mongo.Client) {

	// Goes into this database or creates it, with the collection name
	LoginCollection := client.Database("account").Collection("login")

	// Inserts our kvp with the name username and the password value
	res, err := LoginCollection.InsertOne(context.Background(),
		bson.M{
			"Username": models.Login.Username,
			"Password": models.Login.Password})
	if err != nil {
		log.Fatal(err)
	}
	// Logs our username written ID
	id := res.InsertedID
	log.Printf("Username successfully written to database. ID: %d\n", id)
}
