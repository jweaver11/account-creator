package models

// Creates our login model
type LoginStruct struct {
	Username string // String for username
	Password string // String for password
	Err      string // Error message if any check fails
	Status   string // Status if login was saved or not
}

// Sets our login variable as the login struct
var Login LoginStruct
