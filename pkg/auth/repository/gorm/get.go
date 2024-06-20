package gorm

import (
	"errors"
	"fmt"

	"api-project/internal/domain/presenters"

	"gorm.io/gorm"
)

/*
GetUserByID retrieves a user by their unique ID from the database.

Parameters:
  - id: The unique identifier(ID) of the user to be retrieved.

Returns:
  - user: A pointer to an account.UserResponse containing the details of the retrieved user.
  - error: An error, if any occurred during the database query.
*/
func (repo *repository) GetUserByID(id uint) (*presenters.UserResponse, error) {
	var user presenters.UserResponse

	if err := repo.db.Debug().Table("users").Where("id = ?", id).First(&user).Error; err != nil {

		// If the user is not found, return a specific error message.
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("user not found for id: '%d'", id)
		}

		// Return other database-related errors as-is
		return nil, err
	}

	return &user, nil
}

/*
GetUserByEmail retrieves a user by their email from the database.
Parameters:
  - email: The email of the user to be retrieved.

Returns:
  - user: A pointer to an auth.UserResponse containing the details of the retrieved user.
  - error: An error, if any occurred during the database query.
*/
func (repo *repository) GetUserByEmail(email string) (*presenters.UserResponse, error) {
	var user *presenters.UserResponse

	if err := repo.db.Debug().Table("users").Where("email = ?", email).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("user not found for email: '%s'", email)
		}
		return nil, err
	}

	return user, nil
}

/*
GetUserByPhone retrieves a user by their phone from the database.
Parameters:
  - phone: The phone number of the user to be retrieved.

Returns:
  - user: A pointer to an presenters.UserResponse containing the details of the retrieved user.
  - error: An error, if any occurred during the database query.
*/
func (repo *repository) GetUserByPhone(phone uint) (*presenters.UserResponse, error) {
	var user *presenters.UserResponse

	if err := repo.db.Debug().Table("users").Where("phone = ?", phone).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("user not found for id: '%d'", phone)
		}
		return nil, err
	}

	return user, nil
}
