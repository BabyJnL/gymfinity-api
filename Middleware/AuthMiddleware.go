package Middleware  

import (
	"net/http"

	"github.com/gin-gonic/gin"
    "github.com/golang-jwt/jwt"
)

var JwtyKey = []byte("ASUS_TUF_GAMING")

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
        tokenString := c.GetHeader("Authorization")

        if tokenString == "" {
            c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header missing"})
            c.Abort()
            return
        }

        claims := &jwt.StandardClaims{}

        token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
            return JwtyKey, nil
        })

        if err != nil || !token.Valid {
            c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
            c.Abort()
            return
        }

        c.Set("user", claims.Subject)
        c.Next()
    }
}