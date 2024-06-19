package usecase

import (
	"fmt"
	"strings"

	"api-project/internal/domain/presenters"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
	"github.com/spf13/viper"
)

/*
ExtractToken extracts the JWT token from the "Authorization" header in the Fiber context. It retrieves
the header value, which is expected to be in the format "Bearer <token>", and returns the extracted
token string. If the header is not present or is not in the expected format, an empty string is returned.

Parameters:
  - c: Fiber context representing the incoming HTTP request.

Returns:
  - token: The extracted JWT token string from the "Authorization" header.
*/
func ExtractToken(c *fiber.Ctx) string {
	// Retrieve the "Authorization" header from the Fiber context
	bearer_token := c.Get("Authorization")

	/* Split the header value to extract the token part and check if the header is in the expected
	format "Bearer <token>" and return the extracted token. */
	strArr := strings.Split(bearer_token, " ")
	if len(strArr) == 2 {
		return strArr[1]
	}

	// Return an empty string if the header is not present or not in the expected format
	return ""
}

/*
ParseToken parses a JWT token string and extracts its claims, specifically of type JwtCustomClaims.
If the token is successfully parsed and its claims are valid, the parsed claims are returned;
otherwise, an error is returned.

Parameters:
  - tokenString: The JWT token string to be parsed.
  - tokenType:   The type of token, either "access" or "refresh," to determine the secret key.

Returns:
  - claims: A pointer to JwtCustomClaims containing the parsed claims, if successful.
  - error:  An error, if any occurred during the token parsing or validation process.
*/
func (uCase *usecase) ParseToken(tokenString string, tokenType string) (*presenters.JwtCustomClaims, error) {
	var secret string

	if tokenType == "access" {
		secret = viper.GetString("jwt.access_secret")
	} else {
		secret = viper.GetString("jwt.refresh_secret")
	}

	// Parse the JWT token with the specified claims type and secret key
	token, err := jwt.ParseWithClaims(tokenString, &presenters.JwtCustomClaims{},
		func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			return []byte(secret), nil
		})
	if err != nil {
		return nil, err
	}

	// Check if the token claims are of the expected type and the token is valid
	if claims, ok := token.Claims.(*presenters.JwtCustomClaims); ok && token.Valid {
		return claims, nil
	}

	// Return an error if the token claims are not as expected or the token is not valid
	return nil, err
}

/*
ValidateToken validates the provided JWT token extracted from the Fiber context. It uses the
ParseToken method to parse and extract the claims from the token, and then retrieves the user
information from the repository based on the user ID obtained from the token claims. If the token
is valid and the corresponding user is found in the repository, the user information is returned;
otherwise, an error is returned.

Parameters:
  - c: Fiber context representing the incoming HTTP request.

Returns:
  - user: A pointer to account.UserResponse containing the user information if validation is successful.
  - error: An error, if any occurred during token parsing, user retrieval, or validation.
*/
func (uCase *usecase) ValidateToken(c *fiber.Ctx) (*presenters.UserResponse, error) {
	// Extract the JWT token from the Fiber context
	token := ExtractToken(c)

	// Parse and extract claims from the token using the "access" token type
	claims, err := uCase.ParseToken(token, "access")
	if err != nil {
		return nil, err
	}

	// Retrieve the user information from the repository based on the user ID obtained from the token claims
	user, err := uCase.repo.GetUserByID(claims.ID)
	if err != nil {
		return nil, err
	}

	return user, nil
}
