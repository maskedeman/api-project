package middleware

import (
	"net/http"

	"api-project/internal/domain/presenters"
	repo "api-project/pkg/auth/repository/gorm"
	transactionRepo "api-project/pkg/transaction/repository/gorm"

	"api-project/pkg/auth/usecase"

	"github.com/go-redis/redis"
	"github.com/gofiber/fiber/v2"
	jwtWare "github.com/gofiber/jwt/v2"

	"gorm.io/gorm"
)

/*
Protected is a Fiber middleware function that applies JWT protection to routes. It takes a secret
string as a parameter, which represents the key for validating JWT tokens. It uses the jwtWare
package to configure JWT protection with the provided secret, custom claims, and an error handler
for handling JWT-related errors.
Parameters:
  - secret:       A string representing the key for validating JWT tokens.

Returns:
  - middleware:   A Fiber middleware function for applying JWT protection to routes.
*/
func Protected(secret string) fiber.Handler {
	return jwtWare.New(jwtWare.Config{
		Claims:       &presenters.JwtCustomClaims{},
		SigningKey:   []byte(secret),
		ErrorHandler: jwtError,
	})
}

/*
ValidateJWT is a Fiber middleware function that validates JWT tokens in incoming HTTP requests. It
attempts to validate the JWT token in the request context, and if successful, it sets the validated
user's ID in Fiber locals under the key "requester".
Parameters:
  - db:            GORM database instance for interacting with the database.
  - redisClient:   Redis client instance for additional token validation.
  - config:        JWT configuration settings.

Returns:
  - handler:    A Fiber middleware or handler function to serve HTTP requests.
*/
func ValidateJWT(db *gorm.DB, redisClient *redis.Client, config jwtWare.Config) fiber.Handler {
	errMap := make(map[string]string)

	return func(c *fiber.Ctx) error {
		// Initialize a usecase instance with the provided database and Redis client
		uc := usecase.New(repo.New(db), transactionRepo.New(db), redisClient)

		// Validate the JWT token in the request context
		user, err := uc.ValidateToken(c)
		if err != nil {
			errMap["error"] = err.Error()
			return c.Status(http.StatusForbidden).JSON(presenters.AuthErrorResponse(errMap))
		}

		// Set the validated user's ID in Fiber locals under the key "requesterID".
		c.Locals("requester", user.ID)

		// Continue with the next middleware or route handler
		return c.Next()
	}
}

/*
jwtError generates and returns a JSON response for handling JWT-related errors in a Fiber context.
Parameters:
  - c:   Fiber context representing the incoming HTTP request.
  - err: Error indicating the JWT-related issue.

Returns:
  - error: An error, if any occurred during the response generation.
*/
func jwtError(c *fiber.Ctx, err error) error {
	errMap := make(map[string]string)

	if err.Error() == "Missing or malformed JWT" {
		errMap["error"] = err.Error()
		return c.Status(http.StatusBadRequest).JSON(presenters.AuthErrorResponse(errMap))
	}

	errMap["error"] = err.Error()
	return c.Status(http.StatusUnauthorized).JSON(presenters.AuthErrorResponse(errMap))
}
