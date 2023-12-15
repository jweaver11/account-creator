package models

// Creates our login model
type LoginStruct struct {
	Username string // String for username
	Password string // String for password
}

// Sets our login variable as the login struct
var Login LoginStruct

// Error message if any check fails
var Error string

// json struct of our login struct
type LoginData struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
