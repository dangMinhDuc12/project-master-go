package middleware

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"

	jwt "github.com/golang-jwt/jwt/v4"
)

type JWTClaims struct {
	PhoneNumber string `json:"phoneNumber"`
	jwt.RegisteredClaims
}

var jwtKey = []byte("MY SECRET KEY")


func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		skipAuth := map[string]bool{
				"/assessors/login": true,
		}

		 // Check if the request path is in the skipAuth list
		if skipAuth[c.Request.URL.Path] {
				c.Next()
				return
		}

		authHeader := c.GetHeader("Authorization")

		//validate token exist
		if (authHeader == "") {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header is required"})
			c.Abort() //stop every proccess (controller, other middleware in current route)
			return
		}

		//validate format token
		tokenParts := strings.Split(authHeader, " ")

		if len(tokenParts) != 2 || tokenParts[0] != "Bearer" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid Authorization header format"})
			c.Abort()
			return
		}

		//validate token with secretKey
		tokenString := tokenParts[1]
		jwtClaims := &JWTClaims{}

		jwtToken, errValidateJwtToken := jwt.ParseWithClaims(tokenString, jwtClaims, func(token *jwt.Token) (interface{}, error) {
			return jwtKey, nil
		})

		if errValidateJwtToken != nil || !jwtToken.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			c.Abort()
			return
		}


		c.Next()
	}
}