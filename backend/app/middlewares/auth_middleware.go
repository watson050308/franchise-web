package middleware

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	authHelper "franchise-web/app/helper/admin"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Get Authorization in cookie
		var tokenString string
		tokenStrings := c.GetHeader("Authorization")
		tokenString = tokenStrings
		if tokenString == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header is missing"})
			c.Abort()
			return
		}

		tokenString = strings.TrimPrefix(tokenString, "Bearer ")
		if tokenString == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header is missing or empty"})
			c.Abort()
			return
		}

		// Parse JWT
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			return []byte(os.Getenv("SECRET_KEY")), nil
		})

		// Verify that the JWT is valid
		if err != nil || !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "Invalid token",
			})
			c.Abort()
			return
		}

		tokenID, err := parseTokenID(tokenString)
		log.Println(tokenID)
		if err != nil || authHelper.IsTokenBlacklisted(tokenID) {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			c.Abort()
			return
		}

		// Get claims from token
		if claims, ok := token.Claims.(jwt.MapClaims); ok {
			c.Set("claims", claims)
			log.Println("Parsed claims:", claims)
		} else {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid claims"})
			c.Abort()
			return
		}

		c.Next()
	}
}

func parseTokenID(tokenString string) (string, error) {
	jwtSecret := os.Getenv("SECRET_KEY")
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(jwtSecret), nil
	})
	if err != nil {
		return "", err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return "", fmt.Errorf("invalid token")
	}

	if tokenID, ok := claims["jti"].(string); ok {
		return tokenID, nil
	}

	return "", fmt.Errorf("token ID not found")
}
