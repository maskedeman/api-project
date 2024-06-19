package presenters

import (
	"time"

	"github.com/gofiber/fiber"
	"github.com/golang-jwt/jwt"
)

// Presenter struct for storing custom claims data generated during token generation
type JwtCustomClaims struct {
	ID       uint   `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	jwt.StandardClaims
}

type LoginRequest struct {
	Email    string `json:"email" validate:"required"`
	Phone    uint   `json:"phone" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type AccessTokenRequest struct {
	Refresh string `json:"refresh" validate:"required"`
}

type NewAccessTokenResponse struct {
	Access string `json:"access"`
}

type LoginResponse struct {
	Refresh string            `json:"refresh"`
	Access  string            `json:"access"`
	User    UserLoginResponse `json:"user"`
}

type UserCreateRequest struct {
	ID             uint   `json:"id"`
	Name           string `json:"name"`
	Email          string `json:"email" validate:"required,email" gorm:"unique"`
	Phone          uint   `json:"phone" validate:"required" gorm:"unique" min:"10" max:"10"`
	Password       string `json:"password" validate:"required" min:"8" `
	InitialDeposit uint   `json:"initial_deposit" validate:"required"`
}

type UserLoginResponse struct {
	ID    uint   `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
	Phone string `json:"phone"`
}

// Presenter struct for storing User response data from database query
type UserResponse struct {
	CreatedAt time.Time `json:"createdAt"`
	ID        uint      `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Phone     string    `json:"phone"`
	RoleID    int       `json:"roleID"`
	Verified  bool      `json:"verified"`
	Image     string    `json:"image"`
	Password  string    `json:"-"`
}

/*
LoginSuccessResponse creates a success response for login operations.

Parameters:
  - data: The LoginResponse containing data related to the successful login.

Returns:
  - map: A pointer to a fiber.Map representing the success response.
*/
func LoginSuccessResponse(data LoginResponse) *fiber.Map {
	return &fiber.Map{
		"success": true,
		"data":    data,
	}
}

/*
LogoutSuccessResponse creates a success response for logout operations.

Returns:
  - map: A pointer to a fiber.Map representing the success response.
*/
func LogoutSuccessResponse() *fiber.Map {
	return &fiber.Map{
		"success": true,
	}
}

/*
AuthErrorResponse creates an error response for authentication-related operations.

Parameters:
  - errMap: A map containing error messages related to authentication operations.

Returns:
  - map: A pointer to a fiber.Map representing the error response.
*/
func AuthErrorResponse(errMap map[string]string) *fiber.Map {
	return &fiber.Map{
		"success": false,
		"errors":  errMap,
	}
}

/*
NewAccessTokenSuccessResponse creates a success response map for the "GenerateAccessTokenFromRefreshToken"
endpoint.

Parameters:
  - data: A NewAccessTokenResponse presenter struct containing the newly generated access token.

Returns:
  - map: A Fiber Map representing the success response with the provided data.
*/
func NewAccessTokenSuccessResponse(data NewAccessTokenResponse) *fiber.Map {
	return &fiber.Map{
		"success": true,
		"data":    data,
	}
}
