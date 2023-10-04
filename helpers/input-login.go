/*
Helper file that accepts username and password and calls the check...
helper functions. Adds them to database if they pass checks
*/
package helpers

import (
	"ac/models"
	"fmt"
)

// Calls the input Username and Password functions
func InputLogin() {
	InputUsername()
	InputPassword()
}

// Accepts our username and calls the check username function
func InputUsername() {
	// Sets our username variable and reads in user input
	var un string
	fmt.Scanln(&un)

	// Check that username is not taken and meets criteria
	// If it does, it saves username to login struct
	// Otherwise it returns the error
	if CheckUsername(un) {
		models.Login.Username = un
	} else {
		fmt.Println(models.Login.Err)
	}
}

// Accepts our password and calls the check password function
func InputPassword() {

	// Sets our password variable and reads in user input
	var pw string
	fmt.Scanln(&pw)

	// Check if password meets criteria
	// If it does, it saves password to login struct
	// Otherwise it returns the error
	if CheckPassword(pw) {
		models.Login.Password = pw
	} else {
		fmt.Println(models.Login.Err)
	}
}
