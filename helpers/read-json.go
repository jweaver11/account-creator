package helpers

import (
	"log"
)

func ReadJson() {
	// Replacd with json data received from front end
	jsonData := []byte(`{
		"Username": "username",
		"Password": "password"
	}`)
	log.Println(jsonData)

	/*
		// Unmarshall data from jsonData into LoginData
		err := json.Unmarshal(jsonData, &models.LoginData)
		if err != nil {
			log.Println("Error unmarshalling JSON:", err)
			return
		}

		log.Print("Submitted data now in backend: ")
		log.Println()
		if err := json.NewDecoder(r.Body).Decode(&loginData); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		fmt.Printf("Received text: %s\n", loginData)

		// Respond to client
		w.WriteHeader(http.StatusOK)
	*/

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
