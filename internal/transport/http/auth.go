package http

import (
	"errors"
	"fmt"
	"net/http"
	"strings"
	"time"

	appUser "github.com/amajakai14/admin-renewal/internal/user"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
			return
		}

		authParts := strings.Split(authHeader, " ")
		if len(authParts) != 2 || strings.ToLower(authParts[0]) != "bearer" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
			return
		}

		token, err := extractToken(authParts[1])
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
			return
		}

		if !token.Valid {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
			return
		}
		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "claim not found"})
			return
		}
		userId, ok := claims["userId"]
		if !ok {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "user_id not found"})
			return
		}

		corporationId, ok := claims["corporationId"]
		if !ok {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "corporation_id not found"})
			return
		}
		role, ok := claims["role"]
		if !ok {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "role not found"})
			return
		}
		fmt.Println(userId, corporationId, role)
		c.Set("userId", userId)
		c.Set("corporationId", corporationId)
		c.Set("role", role)
		c.Next()
	}
}

func extractToken(accessToken string) (*jwt.Token, error) {
	var signingKey = []byte("secret")
	token, err := jwt.Parse(
		accessToken,
		func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, errors.New("unexpected signing method")
			}
			return signingKey, nil
		},
	)

	if err != nil {
		return nil, err
	}
	return token, nil
}

func Generate(user *appUser.User) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userId":        user.ID,
		"corporationId": user.CorporationId,
		"role":          user.Role,
		"iat":           time.Now().Unix(),
		"exp":           time.Now().Add(time.Hour * 24).Unix(),
	})

	return token.SignedString([]byte("secret"))
}

