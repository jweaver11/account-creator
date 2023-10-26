package helpers

import (
	"ac/models"
	"encoding/json"
	"fmt"
	"os"
)

func ReadJson() bool {
	// Read the JSON file
	data, err := os.ReadFile("login.json")
	if err != nil {
		fmt.Println("Error reading login.json:", err)
		return false
	}

	// Unmarshal the JSON data into the struct
	if err := json.Unmarshal(data, &models.Login); err != nil {
		fmt.Println("Error unmarshaling JSON:", err)
		return false
	} else {
		return true
	}
}
