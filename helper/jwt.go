package helper

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

// GenerateJWT membuat token JWT berdasarkan user ID dan email
func GenerateJWT(userID int, email string) (string, error) {
	LoadEnv()
	jwtSecret := GetSecretKey()
	secretKey := []byte(jwtSecret)
	claims := jwt.MapClaims{
		"user_id": userID,
		"email":   email,
		"exp":     time.Now().Add(24 * time.Hour).Unix(), // Token berlaku 24 jam
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(secretKey)
}

// ValidateJWT memverifikasi token JWT dan mengembalikan claims jika valid
func ValidateJWT(tokenString string) (jwt.MapClaims, error) {
	LoadEnv()
	jwtSecret := GetSecretKey()
	secretKey := []byte(jwtSecret)
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return secretKey, nil
	})

	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		return claims, nil
	}
	return nil, errors.New("invalid token")
}
