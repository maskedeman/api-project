package http

import (
	"api-project/internal/domain/models"
	"api-project/internal/domain/presenters"
	"api-project/utils"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

func (handler *handler) Withdraw() fiber.Handler {
	return func(c *fiber.Ctx) error {
		errMap := make(map[string]string)
		var requestBody models.UserBalance
		var curBalance *uint

		requestBody.UserID = c.Locals("requester").(uint)
		// Parse and validate the request json body
		err := c.BodyParser(&requestBody)
		if err != nil {
			errMap["error"] = err.Error()
			return c.Status(http.StatusBadRequest).JSON(presenters.ErrorResponse(errMap))
		}

		// Initialize the translation function
		validate, trans := utils.InitTranslator()

		// Validate the request validateBody using a validator and return the translated validation errors if present
		err = validate.Struct(requestBody)
		if err != nil {
			validationErrors := err.(validator.ValidationErrors)
			errMap = utils.TranslateError(validationErrors, trans)

			return c.Status(http.StatusBadRequest).JSON(presenters.ErrorResponse(errMap))
		}

		// Invoke request to the withdraw  usecase to update the amount
		curBalance, errMap = handler.usecase.Withdraw(requestBody)

		if len(errMap) > 0 {
			return c.Status(http.StatusBadRequest).JSON(presenters.ErrorResponse(errMap))
		} else if curBalance == nil {
			return c.JSON(presenters.EmptyResponse{Data: nil, Success: true})
		}

		response := presenters.ListResponse{
			Success: true,
			Data: map[string]interface{}{
				"currentBalance": curBalance,
			},
		}

		return c.JSON(response)
	}
}
