package usecase

import (
	"time"

	"api-project/internal/config"
	"api-project/internal/domain/presenters"

	"github.com/golang-jwt/jwt"
	"github.com/spf13/viper"
)

/*
CreateRefreshToken generates a refresh token for the specified user ID, signs it using the HMAC
SHA-256 algorithm with the refresh secret from the configuration, and sets an expiration time based
on the configured refresh token expiry duration. The generated refresh token is then stored in Redis
with a corresponding key.

Parameters:
  - userID: The user ID for which the refresh token is generated.

Returns:
  - string: The generated refresh token.
  - error: An error message, if any.
*/
func (uCase *usecase) CreateRefreshToken(userID uint, username string) (string, error) {
	var err error

	// Retrieve refresh token secret and expiry duration from configuration file
	refresh_secret := viper.GetString("jwt.refresh_secret")
	refresh_expiry := viper.GetInt("jwt.refresh_expiry_hour")

	// Set expiration time to a distant future if access token never expires
	var expiresAt int64
	if refresh_expiry > 0 {
		expiresAt = time.Now().Add(time.Hour * time.Duration(refresh_expiry)).Unix()
	} else {
		// Set expiration time to a distant future (e.g., 100 years)
		expiresAt = time.Now().AddDate(100, 0, 0).Unix()
	}

	// Create a JWT custom claim with user ID, issuance time, and expiration time
	claims := &presenters.JwtCustomClaims{
		ID:       userID,
		Username: username,
		StandardClaims: jwt.StandardClaims{
			IssuedAt:  time.Now().Unix(),
			ExpiresAt: expiresAt,
		},
	}

	// Generate a signed JWT token using HMAC-SHA256 algorithm
	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	refresh_token, err := jwtToken.SignedString([]byte(refresh_secret))

	if err != nil {
		return "", err
	}

	// Store the generated refresh token in Redis
	refresh_key := "refresh_" + username
	config.SetToRedis(uCase.redisClient, refresh_key, refresh_token)

	return refresh_token, err
}
