package gorm

import (
	"api-project/internal/domain/models"
	"api-project/internal/domain/presenters"
)

/*
RegisterUser persists a new user to the database.

Parameters:
  - user: The user model containing registration information related to the user.

Returns:
  - err: An error if any step of the registration process fails; otherwise, returns nil.
*/
func (repo *repository) RegisterUser(user presenters.UserCreateRequest) (*uint, error) {
	var err error

	// Initiates a database transaction for atomic operations
	tsx := repo.db.Begin()

	newUser := &models.User{
		Timestamp: models.Timestamp{},
		Name:      user.Name,
		Email:     user.Email,
		Phone:     user.Phone,
		Password:  user.Password,
	}

	err = repo.db.Create(&newUser).Error
	if err != nil {
		tsx.Rollback()
		return nil, err
	}

	tsx.Commit()

	return &newUser.ID, nil
}
