package middleware

import (
	"TestApp/internal/helpers"
	. "TestApp/internal/model"
	"encoding/json"
	"fmt"
	"github.com/golang-jwt/jwt/v4"
	"net/http"
)

var (
	secretkey string = "secretkeyjwt"
)

func IsAuthorized(handlerFunc http.HandlerFunc) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		if request.Header["Token"] == nil {
			var err Error
			err = helpers.SetError(err, "No token found")
			json.NewEncoder(writer).Encode(err)
		}

		var mySigningKey = []byte(secretkey)
		token, err := jwt.Parse(request.Header["Token"][0], func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("There was an error in parsing token.")
			}
			return mySigningKey, nil
		})

		if err != nil {
			var err Error
			err = helpers.SetError(err, "Your token has been expired")
			json.NewEncoder(writer).Encode(err)
			return
		}

		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			if claims["role"] == "admin" {
				request.Header.Set("Role", "admin")
				handlerFunc.ServeHTTP(writer, request)
				return
			} else if claims["role"] == "user" {
				request.Header.Set("Role", "user")
				handlerFunc.ServeHTTP(writer, request)
				return
			}
		}

		var resErr Error
		resErr = helpers.SetError(resErr, "Not Authorized")
		json.NewEncoder(writer).Encode(resErr)
	}
}
