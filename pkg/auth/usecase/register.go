package usecase

import (
	"fmt"

	"api-project/internal/domain/models"
	"api-project/internal/domain/presenters"
	"api-project/utils"
)

/*
RegisterUser registers a new user by creating a user account with the provided information.

Parameters:
  - user: The user model containing registration information such as email, phone, password, etc.

Returns:
  - errMap:  A map of error messages, where keys represent the field names and values contain
    respective error details. If registration is successful, it returns nil.
*/
func (uCase *usecase) RegisterUser(request presenters.UserCreateRequest) map[string]string {
	var err error
	errMap := make(map[string]string)

	if _, err := uCase.repo.GetUserByEmail(request.Email); err == nil {
		errMap["email"] = fmt.Errorf("user with the email: '%s' already exists", request.Email).Error()
	}

	if _, err := uCase.repo.GetUserByPhone(request.Phone); err == nil {
		errMap["phone"] = fmt.Errorf("user with the phone: '%v' already exists", request.Phone).Error()
	}

	if len(errMap) != 0 {
		return errMap
	}

	// Check the strength of the provided password
	err = utils.CheckPasswordStrength(request.Password)
	if err != nil {
		errMap["password"] = err.Error()
		return errMap
	}

	// Hash the password for secure storage
	request.Password, err = utils.HashPassword(request.Password)
	if err != nil {
		errMap["password"] = err.Error()
		return errMap
	}

	// Call the repository to save the user information
	uID, err := uCase.repo.RegisterUser(request)
	if err != nil {
		errMap["error"] = err.Error()
		return errMap
	}

	// initial deposit
	if _, err = uCase.transactionRepo.Deposit(models.UserBalance{
		UserID: *uID,
		Amount: request.InitialDeposit,
	}); err != nil {

		// If deposit fails, undo user registration
		undoErr := uCase.repo.DeleteUser(uID)

		if undoErr != nil {
			errMap["error"] = fmt.Sprintf("failed to deposit and undo user registration: %v, %v", err, undoErr)
		} else {
			errMap["error"] = fmt.Sprintf("failed to deposit and user registration has been undone: %v", err)
		}

		return errMap
	}

	return nil
}
