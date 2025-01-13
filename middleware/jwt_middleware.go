package middleware

import (
	"ContractSIMSPPOB/helper"
	"ContractSIMSPPOB/service"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/julienschmidt/httprouter"
)

// GenerateJWT generates a JWT token
func GenerateJWT(userID int, email string) (string, error) {
	helper.LoadEnv()
	secretKey := helper.GetSecretKey()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": userID,
		"email":   email,
		"exp":     time.Now().Add(24 * time.Hour).Unix(),
	})
	return token.SignedString([]byte(secretKey))
}

func JWTAuth(userService service.UserService, next httprouter.Handle) httprouter.Handle {
	helper.LoadEnv()
	secretKey := helper.GetSecretKey()

	return func(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
		authHeader := request.Header.Get("Authorization")
		if authHeader == "" {
			http.Error(writer, "Authorization header is missing", http.StatusUnauthorized)
			return
		}

		// Check if the authorization header starts with "Bearer "
		if !strings.HasPrefix(authHeader, "Bearer ") {
			http.Error(writer, "Invalid Authorization header format", http.StatusUnauthorized)
			return
		}

		// Remove the "Bearer " prefix and quotes (if any)
		tokenString := strings.Trim(strings.TrimPrefix(authHeader, "Bearer "), "\"")

		// Parse and validate the JWT token
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			// Ensure the token's signing method is HMAC
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			return []byte(secretKey), nil
		})

		// Add log to debug the token claims
		if err != nil {
			http.Error(writer, "Error parsing token: "+err.Error(), http.StatusUnauthorized)
			return
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok || !token.Valid {
			http.Error(writer, "Invalid or expired token", http.StatusUnauthorized)
			return
		}

		// Debug: log token's expiration time and other claims
		expiration := claims["exp"].(float64) // or `int64` based on JWT library version
		expirationTime := time.Unix(int64(expiration), 0)
		fmt.Printf("Token expiration time: %v\n", expirationTime)

		if expirationTime.Before(time.Now()) {
			http.Error(writer, "Token has expired", http.StatusUnauthorized)
			return
		}

		// Continue to the next handler if the token is valid
		next(writer, request, params)
	}
}
