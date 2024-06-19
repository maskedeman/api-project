package usecase

import (
	"api-project/internal/domain/models"
)

func (uCase *usecase) Deposit(data models.UserBalance) (*uint, map[string]string) {
	var err error
	var curBalance *uint
	errMap := make(map[string]string)

	_, err = uCase.authRepo.GetUserByID(data.UserID)
	if err != nil {
		errMap["userID"] = err.Error()
		return nil, errMap
	}

	// Delegate the update of data
	curBalance, err = uCase.repo.Deposit(data)
	if err != nil {
		errMap["error"] = err.Error()
		return curBalance, errMap
	}

	return curBalance, errMap
}
