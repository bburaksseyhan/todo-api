package utils

import (
	"errors"
	"net/http"

	"github.com/labstack/gommon/log"
)

type RestErr struct {
	Message string `json:"message"`
	Status  int    `json:"status"`
	Error   string `json:"error"`
}

func NewError(message string) error {
	return errors.New(message)
}

func BadRequestError(message string, err error) *RestErr {

	log.Error(err)

	return &RestErr{
		Message: message,
		Status:  http.StatusBadRequest,
		Error:   "bad_request",
	}
}

func NotFoundRequestError(message string, err error) *RestErr {
	log.Error(err)

	return &RestErr{
		Message: message,
		Status:  http.StatusNotFound,
		Error:   "not_found",
	}
}

func InternalServerError(message string, err error) *RestErr {

	log.Error(err)

	return &RestErr{
		Message: message,
		Status:  http.StatusInternalServerError,
		Error:   "internal_server",
	}
}
