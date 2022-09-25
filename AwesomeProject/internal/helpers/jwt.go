package helpers

import (
	"github.com/golang-jwt/jwt/v4"
	"go.uber.org/zap"
	"os"
	"time"
)

// Generate JWT token
func GenerateJWT(username, role string) (*string, error) {
	secretKey := os.Getenv("SecretKey")
	var mySigningKey = []byte(secretKey)
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["authorized"] = true
	claims["userName"] = username
	claims["role"] = role
	claims["exp"] = time.Now().Add(time.Minute * 30).Unix()

	tokenString, err := token.SignedString(mySigningKey)
	if err != nil {
		Logger.Error("token_string_error", zap.Error(err))
		return nil, err
	}

	return &tokenString, nil
}
