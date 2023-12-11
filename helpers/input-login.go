/*
Helper file that accepts username and password and calls the check...
helper functions. Adds them to database if they pass checks
*/
package helpers

import (
	"ac/models"
	"fmt"
)

// Passes username and password input and runs checks
// If they pass checks, it writes to database
func InputLogin() {
	// Sets our username variable and reads in user input
	var un string
	fmt.Scanln(&un)
	models.Login.Username = un

	// Sets our password variable and reads in user input
	var pw string
	fmt.Scanln(&pw)
	models.Login.Password = pw

}
