package utils

import (
	"net/http"
)

type RestErr struct {
	Message string `json:"message"`
	Status  int    `json:"status"`
	Error   string `json:"error"`
}

// custom error handling return RestErr with bad_request status and messages
func BadRequestError(message string, err error) *RestErr {
	return &RestErr{
		Message: message,
		Status:  http.StatusBadRequest,
		Error:   "bad_request",
	}
}

// custom error handling return RestErr with not_found status and messages
func NotFoundRequestError(message string, err error) *RestErr {
	return &RestErr{
		Message: message,
		Status:  http.StatusNotFound,
		Error:   "not_found",
	}
}

//mcustom error handling return RestErr with internal_server status and messages
func InternalServerError(message string, err error) *RestErr {
	return &RestErr{
		Message: message,
		Status:  http.StatusInternalServerError,
		Error:   "internal_server",
	}
}
