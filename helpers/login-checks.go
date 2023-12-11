package helpers

import (
	"ac/models"
)

// Runs check for username and password
func LoginChecks() bool {
	// Sets errors to none initially
	models.Error = "No errors"
	// Calls all username and password checks to return either true of false
	if CheckUsername(models.Login.Username) && CheckPassword(models.Login.Password) {
		return true
	} else {
		return false
	}
}

// Basic Username check function that calls all username checks
func CheckUsername(un string) bool {
	// Checks username length
	if CheckUNLength(un) {
		// Do nothing
	} else if CheckIsNotTaken(un) {
		// Do nothing
	} else {
		// Fails a check, return false
		return false
	}
	// All username checks passed, return true
	return true
}

// Basic Password check function that calls all password checks
func CheckPassword(pw string) bool {
	// Checks password length
	if CheckPWLength(pw) {
		// Do nothing
	} else {
		// Fails a check, return false
		return false
	}
	// All password checks passed, return true
	return true
}

// Check username length
func CheckUNLength(un string) bool {
	if len(un) < 4 {
		models.Error = "username too short"
		return false
	} else if len(un) > 19 {
		models.Error = "username too long"
		return false
	} else {
		return true
	}
}

// Checks password length
func CheckPWLength(pw string) bool {
	if len(pw) < 4 {
		models.Error = "password too short"
		return false
	} else if len(pw) > 19 {
		models.Error = "password too long"
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
