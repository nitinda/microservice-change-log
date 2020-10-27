package security

import "golang.org/x/crypto/bcrypt"

// HashSecret will incrypt the password
func HashSecret(secret string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(secret), bcrypt.DefaultCost)
}

// VarrifySecret validate hashed secret and secret
func VarrifySecret(hashedSecret, secret string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedSecret), []byte(secret))
}
