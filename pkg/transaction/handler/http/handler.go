package http

import (
	"api-project/internal/domain/interfaces"

	"github.com/gofiber/fiber/v2"
)

type handler struct {
	usecase interfaces.TransactionUsecase
}

func New(app fiber.Router, TransactionUsecase interfaces.TransactionUsecase) {
	handler := &handler{
		usecase: TransactionUsecase,
	}

	transactionHandler := app.Group("/transaction/")
	transactionHandler.Patch("deposit/", (handler.Deposit()))
	transactionHandler.Patch("withdraw/", handler.Withdraw())
	transactionHandler.Get("list/", handler.GetBalance())
	transactionHandler.Patch("transfer/", handler.Transfer())

}
