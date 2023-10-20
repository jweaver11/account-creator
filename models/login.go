package models

// Creates our login model
type LoginStruct struct {
	Username string // String for username
	Password string // String for password
	Err      string // Error message if any check fails
}

// Sets our login variable as the login struct
var Login LoginStruct
