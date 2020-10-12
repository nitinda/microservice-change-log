package security

import (
	"golang.org/x/crypto/bcrypt"
)

// HashPassword will incrypt the password
func HashPassword(password string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
}

// VarrifyPassword validate hashed password and password
func VarrifyPassword(hashedPassword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}
