/* File to write login information to the database */
package helpers

import (
	"ac/models"
	"fmt"
	"log"
	"os"
)

func SaveLogin() {

	// Path for our database later
	databasePath := "test-files/example.txt"

	// Create file right now, but later this part will connect us to the Database
	file, err := os.Create(databasePath)
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer file.Close()

	// Write data to the file.
	username := []byte(models.Login.Username + "\n")
	password := []byte(models.Login.Password + "\n")
	_, err = file.Write(username)
	if err != nil {
		log.Println("Error writing username to file:", err)
		return
	}
	_, err = file.Write(password)
	if err != nil {
		log.Println("Error writing password to file:", err)
		return
	}

	log.Println("Data written to", databasePath)
}
