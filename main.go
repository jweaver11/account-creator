package main

import (
	"ac/helpers"
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {

	router := gin.Default()

	fmt.Println(router)

	RunProgram()
}

func RunProgram() {

	// Prompts user to input username and password
	// This is temporary for front-end
	fmt.Println("Please put in your username and password: ")

	// Input Login will be unneccessary later and instead will start with ReadJson...
	// Whenever there is a request sent through the front end

	// Calls inputlogin function to accept user input
	// This handles all logic and will write to database if login
	// Passes checks
	helpers.InputLogin()
}
