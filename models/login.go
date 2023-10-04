package models

// Creates our login model
type LoginStruct struct {
	Username string
	Password string
	Err      string
}

// Sets our login variable as the login struct
var Login LoginStruct
