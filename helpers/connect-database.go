package helpers

import (
	"context"
	"log"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func ConnectDB() {

	// Context handles information to functions that handle timeouts, cancellations, etc.
	// Sets our context to 10 seconds of failure, then backs out of program
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Adds file path from git ignore to keep database secure
	filePath := ".ignore/mongodb-connect.txt"
	fileData, fail := os.ReadFile(filePath)
	if fail != nil {
		log.Println("Error reading mongodb-connect.txt file")
	}
	URIPath := string(fileData)

	// Sets our client path to the previously set one
	clientOptions := options.Client().ApplyURI(URIPath)

	// Sets our mongo client that handles connections and
	// connects to our desired cluster
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(ctx)

	WriteDB(client)

}
