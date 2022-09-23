package helpers

import (
	"fmt"
	"github.com/golang-jwt/jwt/v4"
	"time"
)

var (
	secretkey string = "secretkeyjwt"
)

func GenerateJwt(userName, role string) (string, error) {
	var mySigningKey = []byte(secretkey)
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["authorized"] = true
	claims["userName"] = userName
	claims["role"] = role
	claims["exp"] = time.Now().Add(time.Minute * 30).Unix()

	tokenString, err := token.SignedString(mySigningKey)
	if err != nil {
		fmt.Errorf("something went wrong:%s", err.Error())
		return "", err
	}
	return tokenString, nil
}
