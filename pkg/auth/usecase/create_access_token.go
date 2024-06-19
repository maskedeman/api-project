package usecase

import (
	"time"

	"api-project/internal/config"
	"api-project/internal/domain/presenters"

	"github.com/golang-jwt/jwt"
	"github.com/spf13/viper"
)

/*
CreateAccessToken generates an access token for the specified user ID, signs it using the HMAC
SHA-256 algorithm with the access secret from the configuration, and sets an expiration time based
on the configured access token expiry duration. The generated access token is then stored in Redis
with a corresponding key.
Parameters:
  - userID: The user ID for which the refresh token is generated.

Returns:
  - string: The generated access token.
  - error: An error message, if any.
*/
func (uCase *usecase) CreateAccessToken(userID uint, username string) (string, error) {
	var err error

	// Retrieve access token secret and expiry duration from configuration file
	access_secret := viper.GetString("jwt.access_secret")
	access_expiry := viper.GetInt("jwt.access_expiry_hour")

	// Set expiration time to a distant future if access token never expires
	var expiresAt int64
	if access_expiry > 0 {
		expiresAt = time.Now().Add(time.Hour * time.Duration(access_expiry)).Unix()
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
	access_token, err := jwtToken.SignedString([]byte(access_secret))
	if err != nil {
		return "", err
	}

	// Store the generated access token in Redis
	access_key := "access_" + username
	config.SetToRedis(uCase.redisClient, access_key, access_token)

	return access_token, err
}
