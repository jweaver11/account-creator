package helpers

import (
	"ac/models"
)

// Basic Username check function that calls all the checks the username needs
// Returns false and sets the error appropriately if any checks fail
func CheckUsername(un string) bool {

	// Set err to none initially
	models.Login.Err = "No errors"

	// Checks if username fits length
	if CheckUNLength(un) {
		if CheckIsNotTaken(un) {
			return true
		} else {
			return false
		}
	} else {
		return false
	}
}

// Basic Password check function that calls all the checks the password needs
// Returns false and sets the error appropriately if any checks fail
func CheckPassword(pw string) bool {

	// Set err to none initially
	models.Login.Err = "No errors"

	// Checks if password fits length
	if CheckPWLength(pw) {
		return true
	} else {
		return false
	}
}

// Check if username and password are long enough
func CheckUNLength(str string) bool {

	if len(str) < 4 {
		models.Login.Err = "username too short"
		return false
	} else if len(str) > 19 {
		models.Login.Err = "username too long"
		return false
	} else {
		return true
	}
}

// Check if username and password are long enough
func CheckPWLength(str string) bool {

	if len(str) < 4 {
		models.Login.Err = "password too short"
		return false
	} else if len(str) > 19 {
		models.Login.Err = "password too long"
		return false
	} else {
		return true
	}
}

// Checks if username is taken in database
func CheckIsNotTaken(un string) bool {

	// Logic here will read the entire collection in database
	// And make sure no username is the same
	// If it is, return false, otherwise return true
	return true
}
