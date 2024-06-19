package interfaces

import (
	"api-project/internal/domain/presenters"

	"github.com/gofiber/fiber/v2"
)

/*
Usecase represents the authentication usecase interface, defining methods for handling various
authentication-related operations.
*/
type AuthUsecase interface {
	LoginUsecase(c *fiber.Ctx, user *presenters.LoginRequest) (presenters.LoginResponse, map[string]string)
	LogoutUsecase(c *fiber.Ctx) map[string]string
	GenerateAccessFromRefreshUsecase(user *presenters.AccessTokenRequest) (presenters.NewAccessTokenResponse, map[string]string)

	RegisterUser(request presenters.UserCreateRequest) map[string]string

	// ForgotPassword(request presenters.ForgotPasswordRequest) error
	// ResetPassword(request presenters.ResetPasswordRequest) map[string]string
}

/*
Repository represents the authentication repository interface, defining methods for handling various
authentication-related operations.
*/
type AuthRepository interface {
	RegisterUser(user presenters.UserCreateRequest) (*uint, error)

	DeleteUser(id *uint) error

	GetUserByID(id uint) (*presenters.UserResponse, error)
	GetUserByEmail(email string) (*presenters.UserResponse, error)
	GetUserByPhone(phone uint) (*presenters.UserResponse, error)
	// ForgotPassword(request presenters.ForgotPasswordRequest) error
	// ResetPassword(request presenters.ResetPasswordRequest) error
}
