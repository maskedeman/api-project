package utils

import (
	"github.com/spf13/viper"
	validator "github.com/wagslane/go-password-validator"
	"golang.org/x/crypto/bcrypt"
)

/*
HashPassword generates a bcrypt hash from the given plaintext password.

Parameters:
  - password: The plaintext password to be hashed.

Returns:
  - The hashed password as a string.
  - An error, if any, encountered during the hashing process.
*/
func HashPassword(password string) (string, error) {
	// Generate a bcrypt hash from the plaintext password
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	return string(bytes), err
}

/*
CheckPasswordHash compares a given plain text password with its bcrypt hashed password equivalent.

Parameters:
  - password: The plaintext password to be checked.
  - hash: The hashed password to compare against.

Returns true if the password matches the hash, and false otherwise.
*/
func CheckPasswordHash(password, hash string) bool {
	// Check if the password matches the hash
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))

	return err == nil
}

/*
CheckPasswordStrength evaluates the strength of a given password based on the configured minimum
entropy threshold. It is used for enforcing secure password strength policies.

Parameters:
  - password: The password string to be checked for strength.

Returns an error if the password does not meet the minimum strength requirement.
*/
func CheckPasswordStrength(password string) error {
	// Retrieve the minimum entropy threshold from the Viper configuration
	minEntropyBits := viper.GetInt("security.password_entropy")

	// Check the strength of the password
	err := validator.Validate(password, float64(minEntropyBits))
	return err
}
