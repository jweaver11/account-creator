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
	// This is temporary for front-end
	fmt.Println("Please put in your username and password: ")

	// Calls inputlogin function to accept user input
	// This handles all logic and will write to database if login
	// Passes checks
	helpers.InputLogin()

	/*router := gin.Default()

	fmt.Println(router)*/
}
