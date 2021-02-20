package errors

import (
	"errors"
	"net/http"
)

type RestErr struct {
	Message string
	Status int
	Error string
}

func NewError(msg string)error  {
return errors.New(msg)
}

func NewBadRequestError(message string)*RestErr  {
	return &RestErr{
		Message: message,
		Status:    http.StatusBadRequest,
		Error:   "Bad Request",
	}
}
func NewNotFoundError(message string)*RestErr  {
	return &RestErr{
		Message: message,
		Status:    http.StatusNotFound,
		Error:   "Not Found",
	}
}

func NewInternalServerError(message string)*RestErr  {
	return &RestErr{
		Message: message,
		Status:    http.StatusInternalServerError,
		Error:   "Internal Server Error",
	}
}