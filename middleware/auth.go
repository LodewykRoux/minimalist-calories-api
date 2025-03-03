package middleware

import (
	"fmt"
	"minimalist-calories-api/initializers"
	"minimalist-calories-api/models"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func Authorize(c *gin.Context) {
	authorization := c.GetHeader("Authorization")

	if authorization == "" {
		c.AbortWithStatus(http.StatusUnauthorized)
	}

	tokenStrings := strings.Split(authorization, " ")
	tokenString := tokenStrings[1]

	if tokenString == "" {
		c.AbortWithStatus(http.StatusUnauthorized)
	}

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method")
		}

		return []byte(os.Getenv("SECRET")), nil
	})

	if err != nil {
		c.AbortWithStatus(http.StatusUnauthorized)
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {

		if float64(time.Now().Unix()) > claims["exp"].(float64) {
			c.AbortWithStatus(http.StatusUnauthorized)
		}

		var user models.User
		initializers.DB.First(&user, claims["id"])

		if user.Id == 0 {
			c.AbortWithStatus(http.StatusUnauthorized)
		}

		c.Set("user", user)

		c.Next()
	} else {
		c.AbortWithStatus(http.StatusUnauthorized)
	}
}
