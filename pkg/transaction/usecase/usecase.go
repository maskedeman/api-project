package usecase

import (
	"api-project/internal/domain/interfaces"
)

/*
usecase represents the authentication usecase, which contains the necessary components for handling
authentication-related business logic. It includes an authentication repository for data access and
a Redis client for token management and storage.
*/
type usecase struct {
	repo     interfaces.TransactionRepository
	authRepo interfaces.AuthRepository
}

/*
New initializes and returns a new instance of the authentication usecase. It takes an authentication
repository and a Redis client as parameters. The usecase is responsible for handling business logic
related to authentication.
*/
func New(repo interfaces.TransactionRepository, authRepo interfaces.AuthRepository) *usecase {
	return &usecase{
		repo:     repo,
		authRepo: authRepo,
	}
}
