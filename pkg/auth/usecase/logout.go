package usecase

import (
	"api-project/internal/config"

	"github.com/gofiber/fiber/v2"
)

/*
LogoutUsecase handles user logout by invalidating access and refresh tokens stored in Redis. It
extracts the token from the Fiber context, parses the user ID from the token claims, and deletes
the corresponding access and refresh tokens from Redis.

Parameters:
  - c: Fiber context representing the incoming HTTP request.

Returns:
  - errMap: A map containing error details, if any occurred during the logout process.
*/
func (uCase *usecase) LogoutUsecase(c *fiber.Ctx) map[string]string {
	// Initialize an error map to store potential errors during the logout process
	errMap := make(map[string]string)

	// Extract the token from the Fiber context
	token := ExtractToken(c)

	// Parse the token claims to retrieve user ID
	claims, err := uCase.ParseToken(token, "access")
	if err != nil {
		errMap["error"] = err.Error()
		return errMap
	}

	// Delete the access token from Redis using the user ID
	access_key := "access_" + claims.Username
	err = config.DeleteFromRedis(uCase.redisClient, access_key)
	if err != nil {
		errMap["error"] = err.Error()
	}

	// Delete the refresh token from Redis using the user ID
	refresh_key := "refresh_" + claims.Username
	err = config.DeleteFromRedis(uCase.redisClient, refresh_key)
	if err != nil {
		errMap["error"] = err.Error()
	}

	return errMap
}
