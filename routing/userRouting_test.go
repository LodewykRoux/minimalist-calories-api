package routing

import (
	"encoding/json"
	"minimalist-calories-api/errorHandling"
	"minimalist-calories-api/initializers"
	"minimalist-calories-api/models"
	"minimalist-calories-api/seedData"
	upsertmodels "minimalist-calories-api/upsertModels"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUserSignUp(t *testing.T) {
	router := SetupRouter()

	dbName := "fake_setup_test.db"
	seedData.SetupLiteDb(t, dbName)

	upsertUser := upsertmodels.UserUpsert{Name: "Test", Email: "test@gmail.com", Password: "Password"}

	userJson, err := json.Marshal(upsertUser)
	if err != nil {
		t.Fatalf("Failed to marshal user data: %v", err)
	}

	w := httptest.NewRecorder()
	req, err := http.NewRequest("POST", "/users/signup", strings.NewReader(string(userJson)))
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}
	req.Header.Set("Content-Type", "application/json")
	router.ServeHTTP(w, req)

	if !assert.Equal(t, http.StatusOK, w.Code) {
		t.Fatalf("Wrong code, expected %v, got %v", http.StatusOK, w.Code)
	}
	var data models.User
	err = json.Unmarshal(w.Body.Bytes(), &data)

	errorHandling.FatalCheck("Failed to unmarshal", err)

	var user models.User
	initializers.DB.First(&user, "email = ?", upsertUser.Email)

	assert.Equal(t, upsertUser.Name, data.Name)
	assert.Equal(t, upsertUser.Email, data.Email)

	assert.Equal(t, upsertUser.Name, user.Name)
	assert.Equal(t, upsertUser.Email, user.Email)

	seedData.CloseConnection(dbName)
}

func TestUserLogin(t *testing.T) {
	router := SetupRouter()
	dbName := "fake_login_test.db"
	defer seedData.CloseConnection(dbName)

	seedData.SetupLiteDb(t, dbName)
	seedData.SeedUser(t)

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

	if !assert.Equal(t, 200, w.Code) {
		t.Fatalf("Wrong code, expected %v, got %v", http.StatusOK, w.Code)
	}
	var data models.User
	err = json.Unmarshal(w.Body.Bytes(), &data)

	errorHandling.FatalCheck("Failed to unmarshal", err)
	assert.Equal(t, upsertUser.Email, data.Email)
}

func TestUserValidation(t *testing.T) {
	router := SetupRouter()
	dbName := "fake_validate_test.db"
	defer seedData.CloseConnection(dbName)

	seedData.SetupLiteDb(t, dbName)
	seedData.SeedUser(t)
	user := seedData.GetUserWithToken(t, router)

	w := httptest.NewRecorder()
	req, err := http.NewRequest("GET", "/users/validate", strings.NewReader(string(user.Token)))
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+user.Token)
	router.ServeHTTP(w, req)

	if !assert.Equal(t, 200, w.Code) {
		t.Fatalf("Wrong code, expected %v, got %v", http.StatusOK, w.Code)
	}
	var data models.User
	err = json.Unmarshal(w.Body.Bytes(), &data)

	errorHandling.FatalCheck("Failed to unmarshal", err)
	assert.Equal(t, user.Email, data.Email)
	assert.Equal(t, user.Email, data.Email)
}

func TestUserLogout(t *testing.T) {
	router := SetupRouter()
	dbName := "fake_logout_test.db"
	defer seedData.CloseConnection(dbName)

	seedData.SetupLiteDb(t, dbName)
	seedData.SeedUser(t)
	user := seedData.GetUserWithToken(t, router)

	w := httptest.NewRecorder()
	req, err := http.NewRequest("POST", "/users/logout", nil)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+user.Token)
	router.ServeHTTP(w, req)

	if !assert.Equal(t, 200, w.Code) {
		t.Fatalf("Wrong code, expected %v, got %v", http.StatusOK, w.Code)
	}
}
