package http

import (
	"api-project/internal/domain/presenters"

	"github.com/gofiber/fiber/v2"
)

// LogoutHandler returns a Fiber.Handler function for handling logout requests.
func (handler *handler) LogoutHandler() fiber.Handler {

	// Initialize an empty map to store field names as key and their translated error messages as value
	errMap := make(map[string]string)

	// Define a Fiber.Handler function which will serve as the handler for processing logout requests
	return func(c *fiber.Ctx) error {

		// Initiate the logout usecase
		errMap = handler.usecase.LogoutUsecase(c)
		if len(errMap) != 0 {
			return c.JSON(presenters.AuthErrorResponse(errMap))
		}
		return c.JSON(presenters.LogoutSuccessResponse())
	}
}
