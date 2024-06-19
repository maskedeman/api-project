package http

import (
	"net/http"

	"api-project/internal/domain/presenters"

	"github.com/gofiber/fiber/v2"
)

/*
GenerateAccessFromRefreshHandler returns a Fiber handler function for the re-generating access token
from refresh token. It handles incoming requests by parsing the request body, extracting the refresh
token, and using it to generate a new access token.

Parameters:
  - c: Fiber context representing the incoming HTTP request.

Returns:
  - error: An error, if any occurred during the handling of the request.
*/
func (handler *handler) GenerateAccessFromRefreshHandler() fiber.Handler {
	// Initialize an error map to store potential errors during request handling
	errMap := make(map[string]string)

	// Return a Fiber handler function
	return func(c *fiber.Ctx) error {
		var requestBody presenters.AccessTokenRequest

		// Parse the request body to extract the refresh token
		err := c.BodyParser(&requestBody)
		if err != nil {
			errMap["error"] = err.Error()
			return c.Status(http.StatusBadRequest).JSON(presenters.AuthErrorResponse(errMap))
		}

		// Call the usecase method to generate a new access token from the provided refresh token
		data, errMap := handler.usecase.GenerateAccessFromRefreshUsecase(&requestBody)
		if len(errMap) != 0 {
			return c.Status(http.StatusUnauthorized).JSON(presenters.AuthErrorResponse(errMap))
		}

		return c.JSON(presenters.NewAccessTokenSuccessResponse(data))
	}
}
