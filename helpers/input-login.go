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
			// If both pass checks, save to database
			SaveLogin()
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

	// Check that username is not taken and meets criteria
	// If it does, it saves username to login struct and writes to database
	// Otherwise it returns the error
	if CheckUsername(un) {
		models.Login.Username = un
		return true
	} else {
		models.Login.Status = "Login not saved"
		return false
	}
}

// Accepts our password and calls the check password function
func InputPassword() bool {

	// Sets our password variable and reads in user input
	var pw string
	fmt.Scanln(&pw)

	// Check if password meets criteria
	// If it does, it saves password to login struct and writes to database
	// Otherwise it returns the error
	if CheckPassword(pw) {
		models.Login.Password = pw
		return true
	} else {
		models.Login.Status = "Login not saved"
		return false
	}
}
