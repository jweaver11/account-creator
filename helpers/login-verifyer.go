package helpers

import (
	"ac/models"
	"log"
)

// Basic Username check function that calls all the checks the username needs
// Returns false and sets the error appropriately if any checks fail
func CheckUsername(un string) bool {
	// Set err to none initially
	models.Login.Err = "No errors"

	// Checks if username fits length
	if CheckLength(un) {

		models.Login.Err = "Username failed length check"
		return false
	}
	return true
}

// Basic Password check function that calls all the checks the password needs
// Returns false and sets the error appropriately if any checks fail
func CheckPassword(pw string) bool {
	// Set err to none initially
	models.Login.Err = "No errors"

	// Checks if password fits length
	if CheckLength(pw) == false {
		models.Login.Err = "Password failed length check"
		return false
	}
	return true
}

// Check if username and password are long enough
func CheckLength(str string) bool {
	if len(str) < 4 {
		log.Println("Username too short")
		models.Login.Err = "Username too short"
		return false
	} else if len(str) > 19 {
		log.Println("Username too long")
		models.Login.Err = "Username too long"
		return false
	} else {
		log.Println("Username passed the length check")
		return true
	}
}
