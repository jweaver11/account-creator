package main

import (
	"ac/helpers"
	"fmt"
)

func main() {
	Run()
}

func Run() {

	// Prompts user to input username and password
	fmt.Println("Please put in your username and password: ")

	// Calls inputlogin function to accept user input
	// This function will also call other functions to write to database
	// Or quit the program if requirements not met
	helpers.InputLogin()

	// Initializes our client variable in helpers to
	// The correct client and database
	helpers.ConnectDB()

	/*router := gin.Default()

	fmt.Println(router)*/
}
