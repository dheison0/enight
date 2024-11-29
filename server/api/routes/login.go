package routes

import (
	"net/http"
	"server/database"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

const AUTH_LIFETIME = 6 * time.Hour

var jwtToken string

func SetJwtToken(token string) {
	jwtToken = token
}

func Login(c *gin.Context) {
	password := c.PostForm("password")

	if !database.CheckPassword(password) {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "incorrect password"})
		return
	}

	// Create a new token object, specifying signing method and the claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"password": password,
		"exp":      time.Now().Add(AUTH_LIFETIME).Unix(), // Token expiration time
	})

	// Sign and get the complete encoded token as a string
	tokenString, err := token.SignedString([]byte(jwtToken))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to generate token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": tokenString})
}

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.GetHeader("Authorization")
		if strings.Contains(tokenString, "Bearer") {
			tokenString = strings.Split(tokenString, "Bearer ")[1]
		}
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, http.ErrAbortHandler
			}
			return []byte(jwtToken), nil
		})
		if err != nil || !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
			c.Abort()
			return
		}
		c.Next()
	}
}
