package usecase

import (
	"time"

	"api-project/internal/config"
	"api-project/internal/domain/presenters"
)

/*
GenerateAccessFromRefreshUsecase generates a new access token based on the provided refresh token.
It parses the refresh token to extract claims,and validates the expiration time of the refresh token.
If the refresh token is valid, it checks its existence in Redis, generates a new access token, and
returns the response containing the access token.
Parameters:
  - user: An AccessTokenRequest presenters struct containing the refresh token.

Returns:
  - response: A NewAccessTokenResponse presenters struct containing the newly generated access token.
  - errMap:   A map containing error details, if any occurred during the usecase process.
*/
func (uCase *usecase) GenerateAccessFromRefreshUsecase(user *presenters.AccessTokenRequest) (presenters.NewAccessTokenResponse, map[string]string) {
	// Initialize an error map to store potential errors during the usecase process
	errMap := make(map[string]string)

	var access_token string
	var response presenters.NewAccessTokenResponse

	// Parse the refresh token to extract claims and validate its expiration time
	claims, err := uCase.ParseToken(user.Refresh, "refresh")
	if err != nil {
		errMap["error"] = err.Error()
		return response, errMap
	}

	// Check if the refresh token is expired
	if claims.ExpiresAt < time.Now().Unix() {
		errMap["error"] = "invalid or expired refresh token"
		return response, errMap
	}

	// Check the existence of the refresh token in Redis
	refresh_key := "refresh_" + claims.Username
	_, err = config.GetFromRedis(uCase.redisClient, refresh_key)
	if err != nil {
		errMap["error"] = "invalid or expired refresh token"
		return response, errMap
	}

	// Generate a new access token using the user ID from the refresh token claims
	access_token, err = uCase.CreateAccessToken(claims.ID, claims.Username)
	if err != nil {
		errMap["error"] = err.Error()
		return response, errMap
	}

	// Set the generated access token in the response
	response.Access = access_token

	return response, errMap
}
