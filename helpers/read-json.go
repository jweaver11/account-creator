package helpers

import (
	"ac/models"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
)

func ReadJson() bool {
	// Read the JSON file
	data, err := os.ReadFile("login.json")
	if err != nil {
		log.Println("Error reading login.json:", err)
		return false
	}

	// Unmarshal the JSON data into the struct
	if err := json.Unmarshal(data, &models.Login); err != nil {
		log.Println("Error unmarshaling JSON:", err)
		return false
	} else {
		return true
	}
}

func ReadJsonx2() {
	var loginData loginData
	if err := json.NewDecoder(r.Body).Decode(&loginData); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	fmt.Printf("Received text: %s\n", loginData)

	// Respond to client
	w.WriteHeader(http.StatusOK)
}

/* Handle enpoints to receive data from front end
http.HandleFunc("/submit", func(w http.ResponseWriter, r *http.Request) {

	var loginData loginData
	if err := json.NewDecoder(r.Body).Decode(&loginData); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	fmt.Printf("Received text: %s\n", loginData)

	// Respond to client
	w.WriteHeader(http.StatusOK)
})
*/
