package usecase

import (
	"api-project/internal/domain/presenters"
	"api-project/utils"

	"github.com/gofiber/fiber/v2"
)

/*
LoginUsecase performs the logic for user login, including validating credentials, creating access
and refresh tokens, and constructing the appropriate login response.

Parameters:
  - c: Fiber context representing the incoming HTTP request.
  - user: Pointer to an presenters.LoginRequest presenters containing user login information.

Returns:
  - response: An presenters.LoginResponse presener containing access and refresh tokens, along with user details.
  - errMap: A map containing error details, if any occurred during the login process.
*/
func (uCase *usecase) LoginUsecase(c *fiber.Ctx, user *presenters.LoginRequest) (presenters.LoginResponse, map[string]string) {
	errMap := make(map[string]string)
	var response presenters.LoginResponse

	// Retrieve the user by email
	result, err := uCase.repo.GetUserByEmail(user.Email)
	if err != nil {
		errMap["error"] = "invalid credentials"
	} else {
		errMap = make(map[string]string)
	}

	// If no user is found by email, attempt to retrieve a user by phone
	if result == nil {
		result, err = uCase.repo.GetUserByPhone(user.Phone)
		if err != nil {
			errMap["error"] = "invalid credentials"
		} else {
			errMap = make(map[string]string)
		}
	}

	if len(errMap) != 0 {
		return response, errMap
	}

	// errMap = utils.ValidateAccess(c.Get("Origin"), result.RoleID)
	// if len(errMap) != 0 {
	// 	return response, errMap
	// }

	// Check if the provided password matches the stored hashed password
	if !utils.CheckPasswordHash(user.Password, result.Password) {
		errMap := make(map[string]string)
		errMap["error"] = "invalid credentials"
		return response, errMap
	}

	// Generate an access token for the authenticated user
	accessToken, err := uCase.CreateAccessToken(result.ID, result.Email)
	if err != nil {
		errMap["error"] = err.Error()
		return response, errMap
	}

	// Generate a refresh token for the authenticated user
	refreshToken, err := uCase.CreateRefreshToken(result.ID, result.Email)
	if err != nil {
		errMap["error"] = err.Error()
		return response, errMap
	}

	// Construct the login response with access and refresh tokens, along with user details
	response = presenters.LoginResponse{
		Refresh: refreshToken,
		Access:  accessToken,
		User: presenters.UserLoginResponse{
			ID:    result.ID,
			Name:  result.Name,
			Email: result.Email,
			Phone: result.Phone,
		},
	}

	return response, errMap
}
