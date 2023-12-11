package models

// Creates our login model
type LoginStruct struct {
	Username string // String for username
	Password string // String for password
}

// Error message if any check fails
var Error string

// Sets our login variable as the login struct
var Login LoginStruct
