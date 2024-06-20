package http

import (
	"api-project/internal/domain/presenters"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func (h *handler) GetBalance() fiber.Handler {
	return func(c *fiber.Ctx) error {

		var userID uint

		var balanceResponse *presenters.BalanceResponse

		userID = c.Locals("requester").(uint)

		// Invoke request to the get balance usecase to update the amount
		balanceResponse, errMap := h.usecase.GetBalance(userID)
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
