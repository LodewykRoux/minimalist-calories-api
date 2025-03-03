package seedData

import (
	"encoding/json"
	"minimalist-calories-api/initializers"
	"minimalist-calories-api/models"
	upsertmodels "minimalist-calories-api/upsertModels"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func SeedUser(t *testing.T) {
	hash, err := bcrypt.GenerateFromPassword([]byte("1234"), 10)
	if err != nil {
		t.Fatalf("Failed to setup user password: %v", err)
		return
	}
	user := models.User{Name: "Test User 1", Email: "test1@gmail.com", Password: string(hash)}

	result := initializers.DB.Create(&user)

	if result.Error != nil {
		t.Fatalf("Failed to setup user: %v", result.Error)
		return
	}
}

func GetUserWithToken(t *testing.T, router *gin.Engine) models.User {
	upsertUser := upsertmodels.UserLogin{Email: "test1@gmail.com", Password: "1234"}

	userJson, err := json.Marshal(upsertUser)
	if err != nil {
		t.Fatalf("Failed to marshal user data: %v", err)
	}

	w := httptest.NewRecorder()
	req, err := http.NewRequest("POST", "/users/login", strings.NewReader(string(userJson)))
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	req.Header.Set("Content-Type", "application/json")
	router.ServeHTTP(w, req)

	var data models.User
	json.Unmarshal(w.Body.Bytes(), &data)
	return data
}
