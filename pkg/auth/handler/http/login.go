package http

import (
	"net/http"

	"api-project/internal/domain/presenters"

	"api-project/utils"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

// LoginHandler returns a Fiber.Handler function for handling login requests.
func (handler *handler) LoginHandler() fiber.Handler {
	// Initialize an empty map to store field names as key and their translated error messages as value
	errMap := make(map[string]string)

	// Define a Fiber.Handler function which will serve as the handler for processing login requests
	return func(c *fiber.Ctx) error {
		var requestBody presenters.LoginRequest

		// Parse the JSON request body into an presenters.LoginRequest struct.
		err := c.BodyParser(&requestBody)
		if err != nil {
			errMap["error"] = err.Error()
			c.Status(http.StatusBadRequest)
			return c.JSON(presenters.AuthErrorResponse(errMap))
		}

		/* Validate the parsed request body. If failed, translate the validation errors into user-friendly
		messages, add them to the error map */
		validate, trans := utils.InitTranslator()
		err = validate.Struct(requestBody)
		if err != nil {
			validationErrors := err.(validator.ValidationErrors)
			errMap = utils.TranslateError(validationErrors, trans)
			c.Status(http.StatusBadRequest)
			return c.JSON(presenters.AuthErrorResponse(errMap))
		}

		// Initiate the login usecase with the validated request
		response, errMap := handler.usecase.LoginUsecase(c, &requestBody)
		if len(errMap) != 0 {
			if _, ok := errMap["403"]; ok {
				errMap2 := make(map[string]string)
				errMap2["error"] = errMap["403"]
				return c.Status(http.StatusForbidden).JSON(presenters.AuthErrorResponse(errMap2))
			}

			return c.Status(http.StatusBadRequest).JSON(presenters.AuthErrorResponse(errMap))
		}

		return c.JSON(presenters.LoginSuccessResponse(response))
	}
}
