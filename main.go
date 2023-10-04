package main

import (
	"ac/helpers"
	"ac/models"
	"fmt"
)

func main() {

	// Prompts user to input username and password
	fmt.Println("Please put in your username and password: ")

	// Calls inputlogin function to accept user input
	helpers.InputLogin()

	helpers.SaveLogin()

	// Prints login to see if it was saved to database or not
	fmt.Printf("Username: %s \nPassword: %s\n", models.Login.Username, models.Login.Password)

	/*router := gin.Default()

	fmt.Println(router)*/

}
