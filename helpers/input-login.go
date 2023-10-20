/*
Helper file that accepts username and password and calls the check...
helper functions. Adds them to database if they pass checks
*/
package helpers

import (
	"ac/models"
	"fmt"
	"log"
)

// Calls username and password
// If username or password fails check
// then program won't write to database
func InputLogin() {

	// Check if un and pw pass checks
	if InputUsername() {

		if InputPassword() {

			// If both pass checks, connect and save to database
			ConnectDB()
			log.Printf("\nUsername: %s \nPassword: %s\n", models.Login.Username, models.Login.Password)
		} else {

			// Otherwise log the errors
			log.Println(models.Login.Err)
		}
	} else {

		// Otherwise log the errors (only username passed check)
		log.Println(models.Login.Err)
	}
}

// Accepts our username and calls the check username function
func InputUsername() bool {

	// Sets our username variable and reads in user input
	var un string
	fmt.Scanln(&un)

	// Makes sure username passes all checks
	if CheckUsername(un) {

		// Saves un to our struct
		models.Login.Username = un
		return true
	} else {

		// Otherwise returns it did not pass checks
		return false
	}
}

// Accepts our password and calls the check password function
func InputPassword() bool {

	// Sets our password variable and reads in user input
	var pw string
	fmt.Scanln(&pw)

	// Makes sure password passes all chekcs
	if CheckPassword(pw) {

		// Saves pw to our struct if checks passed
		models.Login.Password = pw
		return true
	} else {

		// Otherwise returns that it did not pass checks
		return false
	}
}
