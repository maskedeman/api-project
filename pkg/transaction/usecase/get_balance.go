package usecase

import (
	"api-project/internal/domain/presenters"
)

func (uCase *usecase) GetBalance(userID uint) (*presenters.BalanceResponse, map[string]string) {
	var err error
	errMap := make(map[string]string)

	// Check if the user exists
	usr, err := uCase.authRepo.GetUserByID(userID)
	if err != nil {
		errMap["userID"] = err.Error()
		return nil, errMap
	}

	// Delegate the update of data
	amount, err := uCase.repo.GetAmountByUserID(userID)
	if err != nil {
		errMap["error"] = err.Error()
		return nil, errMap
	}
	userData := make(map[string]interface{})
	userData["id"] = usr.ID
	userData["name"] = usr.Name
	userData["phone"] = usr.Phone

	return &presenters.BalanceResponse{
		User:           userData,
		CurrentBalance: amount,
	}, errMap
}
