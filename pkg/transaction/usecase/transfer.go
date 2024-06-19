package usecase

import (
	"api-project/internal/domain/presenters"
	"time"
)

func (uCase *usecase) Transfer(data presenters.TransferRequest) (*presenters.TransferResponse, map[string]string) {
	var err error
	errMap := make(map[string]string)

	debitedFrom, err := uCase.authRepo.GetUserByID(data.From)
	if err != nil {
		errMap["debitedFrom"] = err.Error()
		return nil, errMap
	}

	debitUserData := make(map[string]interface{})
	debitUserData["id"] = debitedFrom.ID
	debitUserData["name"] = debitedFrom.Name
	debitUserData["phone"] = debitedFrom.Phone

	creditedTo, err := uCase.authRepo.GetUserByID(data.To)
	if err != nil {
		errMap["creditedTo"] = err.Error()
		return nil, errMap
	}

	creditUserData := make(map[string]interface{})
	creditUserData["id"] = creditedTo.ID
	creditUserData["name"] = creditedTo.Name
	creditUserData["phone"] = creditedTo.Phone

	// Delegate the update of data
	_, err = uCase.repo.Transfer(data)
	if err != nil {
		errMap["error"] = err.Error()
		return nil, errMap
	}

	remainingAmount, err := uCase.repo.GetAmountByUserID(data.From)
	if err != nil {
		errMap["error"] = err.Error()
		return nil, errMap
	}

	tLog, err := uCase.repo.GetTransactionLogsByUserID(data.From)
	if err != nil {
		errMap["error"] = err.Error()
		return nil, errMap
	}

	return &presenters.TransferResponse{
		TransactionID:     tLog.ID,
		Time:              time.Now(),
		FromUser:          debitUserData,
		ToUser:            creditUserData,
		TransactionAmount: data.Amount,
		RemainingAmount:   remainingAmount,
		TransactionType:   "Transfer(-)",
	}, errMap
}
