package main

import (
	"ac/helpers"
	"ac/models"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {

	// Creates a new router (controller)
	router := gin.Default()

	/* router accepts post requests called submit
	router.POST("/submit", func(c *gin.Context) {

		// Logic to pass correct handler
		helpers.ReadJson()
		// display status of request
		c.String(http.StatusOK, "You submitted username and password")
	})
	*/

	// middleware to serve static files
	router.Static("/web", "./web")

	// Runs the server on port 8080
	router.Run(":8080")

	// TEEEMEMMMMMMMMPPPPP
	// Calls inputlogin function to accept un and pw from front end
	helpers.InputLogin()

	// LOGIC WHEN USERNAME AND PASSWORD ARE PASSED TO BACKEND
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
