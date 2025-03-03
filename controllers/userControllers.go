package controllers

import (
	"minimalist-calories-api/errorHandling"
	"minimalist-calories-api/initializers"
	"minimalist-calories-api/models"
	upsertmodels "minimalist-calories-api/upsertModels"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

func createToken(user models.User, c *gin.Context) string {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":  user.Id,
		"exp": time.Now().Add(time.Hour * 14 * 30).Unix(),
	})

	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET")))
	if err != nil {
		errorHandling.HandleBadRequest(c, err, "Failed to create token")
		return ""
	}

	return tokenString
}

func SignUp(c *gin.Context) {
	var userUpsert upsertmodels.UserUpsert

	err := c.Bind(&userUpsert)
	if err != nil {
		errorHandling.HandleBadRequest(c, err, "Failed to read request")
		return
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(userUpsert.Password), 10)
	if err != nil {
		errorHandling.HandleBadRequest(c, err, "Failed to hash password")
		return
	}

	user := models.User{Name: userUpsert.Name, Email: userUpsert.Email, Password: string(hash)}
	result := initializers.DB.Create(&user)
	if result.Error != nil {
		errorHandling.HandleBadRequest(c, result.Error, "Failed to create user.")
		return
	}

	user.Token = createToken(user, c)
	updateResult := initializers.DB.Save(&user)
	if result.Error != nil {
		errorHandling.HandleBadRequest(c, updateResult.Error, "Failed to create token.")
		return
	}

	c.Set("user", user)

	c.JSON(http.StatusOK, user)
}

func Login(c *gin.Context) {
	var userLogin upsertmodels.UserLogin
	err := c.Bind(&userLogin)
	errorHandling.HandleBadRequest(c, err, "Failed to read request")

	var user models.User
	result := initializers.DB.First(&user, "email = ?", userLogin.Email)
	if result.Error != nil {
		errorHandling.HandleBadRequest(c, result.Error, "No user found")
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(userLogin.Password))
	if err != nil {
		errorHandling.HandleBadRequest(c, err, "Incorrect password")
		return
	}

	user.Token = createToken(user, c)
	updateResult := initializers.DB.Save(&user)
	if updateResult.Error != nil {
		errorHandling.HandleBadRequest(c, updateResult.Error, "Failed to create token")
		return
	}

	c.Set("user", user)

	c.JSON(http.StatusOK, user)
}

func Validate(c *gin.Context) {
	user, _ := c.Get("user")

	c.JSON(http.StatusOK, user)
}

func Logout(c *gin.Context) {
	savedUser, _ := c.Get("user")

	var user models.User
	result := initializers.DB.First(&user, "email = ?", savedUser.(models.User).Email)
	if result.Error != nil {
		errorHandling.HandleBadRequest(c, result.Error, "No user found")
		return
	}

	user.Token = ""
	updateResult := initializers.DB.Save(&user)
	if updateResult.Error != nil {
		errorHandling.HandleBadRequest(c, updateResult.Error, "Failed to log out")
		return
	}

	c.JSON(http.StatusOK, gin.H{})
}
