package usecase

import (
	"api-project/internal/domain/models"
)

func (uCase *usecase) Withdraw(data models.UserBalance) (*uint, map[string]string) {
	var err error
	var curBalance *uint
	errMap := make(map[string]string)

	// Check if the user exists
	_, err = uCase.authRepo.GetUserByID(data.UserID)
	if err != nil {
		errMap["userID"] = err.Error()
		return nil, errMap
	}

	// Delegate the update of data
	if curBalance, err = uCase.repo.Withdraw(data); err != nil {
		errMap["error"] = err.Error()
		return curBalance, errMap
	}

	return curBalance, errMap
}
