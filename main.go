package main

import (
	"ac/helpers"
	"ac/models"
	"fmt"
	"log"

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

	// Calls inputlogin function to accept un and pw from front end
	helpers.InputLogin()

	// If the username and password pass the checks, return true
	if helpers.LoginChecks() {
		// Write the login to database
		helpers.SaveLogin()
		log.Printf("\nUsername: %s \nPassword: %s\n", models.Login.Username, models.Login.Password)

	} else {
		// Otherwise log the errors
		log.Println(models.Error)
	}
}
