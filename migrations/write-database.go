package migrations

import (
	"ac/models"
	"context"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func WriteDB(client *mongo.Client) {
	// Goes into this database or creates it, with the collection name
	UNcollection := client.Database("login").Collection("username")
	PWcollection := client.Database("login").Collection("password")

	// Inserts our kvp with the name username and the username value
	res, err := UNcollection.InsertOne(context.Background(), bson.M{"Username": models.Login.Username})
	if err != nil {
		log.Fatal(err)
	}

	// Inserts our kvp with the name password and the password value
	res2, err2 := PWcollection.InsertOne(context.Background(), bson.M{"Password": models.Login.Password})
	if err2 != nil {
		log.Fatal(err2)
	}

	// Sets the id of the just inserted data for loggin purposes
	id := res.InsertedID
	id2 := res2.InsertedID
	log.Println(id)
	log.Println(id2)
}
