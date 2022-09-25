package middleware

import (
	"AwesomeProject/internal/helpers"
	"github.com/golang-jwt/jwt/v4"
	"net/http"
	"os"
)

var (
	secretKey          = os.Getenv("SecretKey")
	ErrorTokenNotFound = helpers.NewError("token_not_found", http.StatusBadRequest)
	ErrorTokenParsing  = helpers.NewError("parsing_error", http.StatusBadRequest)
	ErrorTokenExpired  = helpers.NewError("active_token_expired", http.StatusBadRequest)
	ErrorNotAuthorized = helpers.NewError("not_authorized", http.StatusUnauthorized)
)

func IsAuthorized(handler http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		if r.Header["Token"] == nil {
			helpers.JSON(w, http.StatusBadRequest, ErrorTokenNotFound)
			return
		}

		var mySigningKey = []byte(secretKey)

		token, err := jwt.Parse(r.Header["Token"][0], func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, ErrorTokenParsing
			}
			return mySigningKey, nil
		})

		if err != nil {
			helpers.JSON(w, http.StatusUnauthorized, ErrorTokenExpired)
			return
		}

		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			if claims["role"] == "admin" {
				r.Header.Set("Role", "admin")
				handler.ServeHTTP(w, r)
				return

			} else if claims["role"] == "user" {
				r.Header.Set("Role", "user")
				handler.ServeHTTP(w, r)
				return

			}
		}

		helpers.JSON(w, http.StatusUnauthorized, ErrorNotAuthorized)
	}
}
