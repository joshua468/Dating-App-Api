package utils

import (
	"errors"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
)

// GenerateJWT generates a JWT token for the given userID.
func GenerateJWT(userID uint) (string, error) {
	claims := jwt.MapClaims{
		"user_id": userID,
		"exp":     time.Now().Add(time.Hour * 24).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(GetEnv("JWT_SECRET")))
}

// ValidateToken validates the JWT token and returns the claims.
func ValidateToken(token string) (jwt.MapClaims, error) {
	parsedToken, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return []byte(GetEnv("JWT_SECRET")), nil
	})

	if err != nil {
		return nil, err
	}

	// Check if claims can be asserted to jwt.MapClaims
	if claims, ok := parsedToken.Claims.(jwt.MapClaims); ok && parsedToken.Valid {
		return claims, nil
	}
	return nil, errors.New("invalid token")
}

// GetEnv retrieves the value of the environment variable.
func GetEnv(key string) string {
	return os.Getenv(key)
}
