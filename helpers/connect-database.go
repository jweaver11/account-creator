package helpers

import (
	"ac/migrations"
	"context"
	"log"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func SaveLogin() {
	// Function will connect to database and then write to it
	ConnectDB()
}

// Connects to our database and initializes our Json file
func ConnectDB() {
	// Context handles information to functions that handle timeouts, cancellations, etc.
	// Sets our context to 10 seconds of failure, then backs out of program
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	serverAPI := options.ServerAPI(options.ServerAPIVersion1)

	// Adds file path from git ignore to keep database secure
	filePath := ".ignore/mongodb-connect.txt"
	// Reads connection string from file
	fileData, fail := os.ReadFile(filePath)
	if fail != nil {
		log.Println("Error reading mongodb-connect.txt file")
	}
	// From []byte to string
	connectionString := string(fileData)

	// Sets our options
	opts := options.Client().ApplyURI(connectionString).SetServerAPIOptions(serverAPI)

	// Sets our mongo client that handles connections and
	// connects to our desired cluster
	client, err := mongo.Connect(ctx, opts)
	if err != nil {
		log.Fatal(err)
		panic(err)
	}
	defer client.Disconnect(ctx)

	// Writes to our database using the just made connections
	migrations.WriteData(client)
}
