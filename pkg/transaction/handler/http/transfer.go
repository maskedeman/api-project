package http

import (
	"api-project/internal/domain/presenters"
	"api-project/utils"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

func (h *handler) Transfer() fiber.Handler {
	return func(c *fiber.Ctx) error {
		errMap := make(map[string]string)
		var requestBody presenters.TransferRequest
		var curBalance *presenters.TransferResponse

		requestBody.From = c.Locals("requester").(uint)

		err := c.BodyParser(&requestBody)
		if err != nil {
			errMap["error"] = err.Error()
			return c.Status(http.StatusBadRequest).JSON(presenters.ErrorResponse(errMap))
		}

		validate, trans := utils.InitTranslator()

		// Validate the request validateBody using a validator and return the translated validation errors if present
		err = validate.Struct(requestBody)
		if err != nil {
			validationErrors := err.(validator.ValidationErrors)
			errMap = utils.TranslateError(validationErrors, trans)

			return c.Status(http.StatusBadRequest).JSON(presenters.ErrorResponse(errMap))
		}

		// Invoke request to the transfer usecase to update the amount
		curBalance, errMap = h.usecase.Transfer(requestBody)
		if len(errMap) > 0 {
			return c.Status(http.StatusBadRequest).JSON(presenters.ErrorResponse(errMap))
		} else if curBalance == nil {
			return c.JSON(presenters.EmptyResponse{Data: nil, Success: true})
		}

		response := presenters.ListResponse{
			Success: true,
			Data:    curBalance,
		}

		return c.JSON(response)
	}
}
