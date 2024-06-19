package presenters

import (
	"github.com/gofiber/fiber/v2"
)

/*
ListResponse is a common data structure used for API responses that represent a list of items. It
includes information about the success status, the current page number, the total number of pages,
and the actual data that is being returned. The data field is of type interface{} to accommodate
various types of data structures.
*/
type ListResponse struct {
	Success bool        `json:"success"`
	Data    interface{} `json:"data"`
}

// EmptyResponse represents a response structure used for endpoints that return empty data arrays.
type EmptyResponse struct {
	Success bool        `json:"success"`
	Data    interface{} `json:"data"`
}

/*
SuccessResponse creates a success response in the form of a Fiber Map, commonly used for indicating
the success of an operation.
*/
func SuccessResponse() *fiber.Map {
	return &fiber.Map{
		"success": true,
	}
}

/*
ErrorResponse creates an error response in the form of a Fiber Map, which includes a map of errors
and sets the success flag to false, often used to communicate errors in responses.
*/
func ErrorResponse(err map[string]string) *fiber.Map {
	return &fiber.Map{
		"success": false,
		"errors":  err,
	}
}
