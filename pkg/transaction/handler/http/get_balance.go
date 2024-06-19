package http

import (
	"api-project/internal/domain/presenters"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func (h *handler) GetBalance() fiber.Handler {
	return func(c *fiber.Ctx) error {
		errMap := make(map[string]string)
		var userID uint

		var balanceResponse *presenters.BalanceResponse

		userID = c.Locals("requester").(uint)

		// err := c.BodyParser(userID)
		// if err != nil {
		// 	errMap["error"] = err.Error()
		// 	return c.Status(http.StatusBadRequest).JSON(presenters.ErrorResponse(errMap))
		// }

		// validate, trans := utils.InitTranslator()

		// // Validate the request validateBody using a validator and return the translated validation errors if present
		// err = validate.Struct(userID)
		// if err != nil {
		// 	validationErrors := err.(validator.ValidationErrors)
		// 	errMap = utils.TranslateError(validationErrors, trans)

		// 	return c.Status(http.StatusBadRequest).JSON(presenters.ErrorResponse(errMap))
		// }

		// Invoke request to the get balance usecase to update the amount
		balanceResponse, errMap = h.usecase.GetBalance(userID)
		if len(errMap) > 0 {
			return c.Status(http.StatusBadRequest).JSON(presenters.ErrorResponse(errMap))
		}

		response := presenters.ListResponse{
			Success: true,
			Data:    balanceResponse,
		}

		return c.JSON(response)
	}
}
