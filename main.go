package main

import (
	"ac/helpers"
	"fmt"
)

func main() {

	// Prompts user to input username and password
	fmt.Println("Please put in your username and password: ")

	// Calls inputlogin function to accept user input
	helpers.InputLogin()

	/*router := gin.Default()

	fmt.Println(router)*/

}
