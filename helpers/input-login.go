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

// Passes username and password input and runs checks
// If they pass checks, it writes to database
func InputLogin() {
	InputUsername()
	InputPassword()
	// If the username and password pass the checks, return true
	if LoginChecks() {
		// Write the login to database
		SaveLogin()
		log.Printf("\nUsername: %s \nPassword: %s\n", models.Login.Username, models.Login.Password)
	} else {
		// Otherwise log the errors
		log.Println(models.Error)
	}
}

// Accepts our username and calls the check username function
func InputUsername() {
	// Sets our username variable and reads in user input
	var un string
	fmt.Scanln(&un)
	models.Login.Username = un
}

// Accepts our password and calls the check password function
func InputPassword() {
	// Sets our password variable and reads in user input
	var pw string
	fmt.Scanln(&pw)
	models.Login.Password = pw
}
