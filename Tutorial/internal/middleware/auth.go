package middleware

import (
	"github.com/gin-gonic/gin"
)

var (
	secretkey string = "secretkeyjwt"
)

func IsAuthorized(handlerFunc gin.HandlerFunc) gin.HandlerFunc {

}
