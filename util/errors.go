package util

import "net/http"

type ErrorMessage struct {
	HttpStatus int
	Message    string
}


func (e ErrorMessage) ErrorLoadingEnvFile() *ErrorMessage {

	return &ErrorMessage{
		HttpStatus: 	http.StatusInternalServerError,
		Message: 		"Error loading .env file",
	}
}


func (e ErrorMessage) FailedToOpenDB() *ErrorMessage {

	return &ErrorMessage{
		HttpStatus: 	http.StatusInternalServerError,
		Message: 		"Failed to open database!",
	}
}


func (e ErrorMessage) FailedToCreateHashFromPassword() *ErrorMessage {

	return &ErrorMessage{
		HttpStatus: 	http.StatusInternalServerError,
		Message: 		"Failed to create hash from password",
	}
}


func (e ErrorMessage) UserNotFound() *ErrorMessage {

	return &ErrorMessage{
		HttpStatus: 	http.StatusNotFound,
		Message: 		"Invalid username or password!",
	}
}

func (e ErrorMessage) FailedToCreateJWToken() *ErrorMessage {

	return &ErrorMessage{
		HttpStatus: 	http.StatusInternalServerError,
		Message: 		"Failed to create jwt",
	}
}