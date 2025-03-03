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

func TestFoodNewSave(t *testing.T) {
	router := SetupRouter()
	dbName := "fake_saveNewFood_test.db"
	defer seedData.CloseConnection(dbName)

	seedData.SetupLiteDb(t, dbName)
	seedData.SeedUser(t)
	user := seedData.GetUserWithToken(t, router)
	seedData.SeedFood(t, user)

	food := seedData.NewFood(t, user)

	foodJson, err := json.Marshal(food)
	if err != nil {
		t.Fatalf("Failed to marshal user data: %v", err)
	}

	w := httptest.NewRecorder()
	req, err := http.NewRequest("POST", "/food/save", strings.NewReader(string(foodJson)))
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+user.Token)
	router.ServeHTTP(w, req)

	if !assert.Equal(t, http.StatusOK, w.Code) {
		t.Fatalf("Wrong code, expected %v, got %v", http.StatusOK, w.Code)
	}

	var data models.FoodItem
	err = json.Unmarshal(w.Body.Bytes(), &data)

	errorHandling.FatalCheck("Failed to unmarshal", err)

	var foodItem models.FoodItem
	initializers.DB.First(&foodItem, "id = ?", data.ID)

	assert.ObjectsAreEqualValues(foodItem, data)

	assert.Equal(t, food.Name, data.Name)
	assert.Equal(t, food.Uom, data.Uom)
	assert.Equal(t, food.Quantity, data.Quantity)
	assert.Equal(t, food.Calories, data.Calories)
	assert.Equal(t, food.Protein, data.Protein)
	assert.Equal(t, food.Carbs, data.Carbs)
	assert.Equal(t, food.Fat, data.Fat)
	assert.Equal(t, food.UserID, data.UserID)
}

func TestFoodExistingSave(t *testing.T) {
	router := SetupRouter()
	dbName := "fake_saveExistingFood_test.db"
	defer seedData.CloseConnection(dbName)

	seedData.SetupLiteDb(t, dbName)
	seedData.SeedUser(t)
	user := seedData.GetUserWithToken(t, router)
	existingFood := seedData.SeedFood(t, user)

	existingFood.Name = "New Name"

	foodJson, err := json.Marshal(existingFood)
	if err != nil {
		t.Fatalf("Failed to marshal user data: %v", err)
	}

	w := httptest.NewRecorder()
	req, err := http.NewRequest("POST", "/food/save", strings.NewReader(string(foodJson)))
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+user.Token)
	router.ServeHTTP(w, req)

	if !assert.Equal(t, http.StatusOK, w.Code) {
		t.Fatalf("Wrong code, expected %v, got %v", http.StatusOK, w.Code)
	}

	var data models.FoodItem
	err = json.Unmarshal(w.Body.Bytes(), &data)

	errorHandling.FatalCheck("Failed to unmarshal", err)

	var foodItem models.FoodItem
	initializers.DB.First(&foodItem, "id = ?", data.ID)

	assert.ObjectsAreEqualValues(foodItem, data)

	assert.Equal(t, existingFood.Name, data.Name)
	assert.Equal(t, existingFood.Uom, data.Uom)
	assert.Equal(t, existingFood.Quantity, data.Quantity)
	assert.Equal(t, existingFood.Calories, data.Calories)
	assert.Equal(t, existingFood.Protein, data.Protein)
	assert.Equal(t, existingFood.Carbs, data.Carbs)
	assert.Equal(t, existingFood.Fat, data.Fat)
	assert.Equal(t, existingFood.UserID, data.UserID)
}

func TestFoodSaveFail(t *testing.T) {
	router := SetupRouter()
	dbName := "fake_saveFoodFail_test.db"
	defer seedData.CloseConnection(dbName)

	seedData.SetupLiteDb(t, dbName)
	seedData.SeedUser(t)
	user := seedData.GetUserWithToken(t, router)

	foodJson, err := json.Marshal("")
	if err != nil {
		t.Fatalf("Failed to marshal user data: %v", err)
	}

	w := httptest.NewRecorder()
	req, err := http.NewRequest("POST", "/food/save", strings.NewReader(string(foodJson)))
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+user.Token)
	router.ServeHTTP(w, req)

	if !assert.Equal(t, http.StatusBadRequest, w.Code) {
		t.Fatalf("Wrong code, expected %v, got %v", http.StatusBadRequest, w.Code)
	}
}

func TestFoodGetList(t *testing.T) {
	router := SetupRouter()
	dbName := "fake_getListFood_test.db"
	defer seedData.CloseConnection(dbName)

	seedData.SetupLiteDb(t, dbName)
	seedData.SeedUser(t)
	user := seedData.GetUserWithToken(t, router)
	seedData.SeedFood(t, user)

	w := httptest.NewRecorder()
	req, err := http.NewRequest("GET", "/food/getList", nil)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+user.Token)
	router.ServeHTTP(w, req)

	if !assert.Equal(t, http.StatusOK, w.Code) {
		t.Fatalf("Wrong code, expected %v, got %v", http.StatusOK, w.Code)
	}

	var data []models.FoodItem
	err = json.Unmarshal(w.Body.Bytes(), &data)

	errorHandling.FatalCheck("Failed to unmarshal", err)

	assert.NotEmpty(t, data)
}

func TestFoodDelete(t *testing.T) {
	router := SetupRouter()
	dbName := "fake_deleteFood_test.db"
	defer seedData.CloseConnection(dbName)

	seedData.SetupLiteDb(t, dbName)
	seedData.SeedUser(t)
	user := seedData.GetUserWithToken(t, router)
	seedData.SeedFood(t, user)

	foodItemDelete := upsertmodels.FoodItemDelete{ID: seedData.ValidFoodId}
	foodJson, err := json.Marshal(foodItemDelete)
	if err != nil {
		t.Fatalf("Failed to marshal user data: %v", err)
	}

	w := httptest.NewRecorder()
	req, err := http.NewRequest("DELETE", "/food/delete", strings.NewReader(string(foodJson)))
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+user.Token)
	router.ServeHTTP(w, req)

	if !assert.Equal(t, http.StatusOK, w.Code) {
		t.Fatalf("Wrong code, expected %v, got %v", http.StatusOK, w.Code)
	}

	var data models.FoodItem
	err = json.Unmarshal(w.Body.Bytes(), &data)

	errorHandling.FatalCheck("Failed to unmarshal", err)

	var foodItem models.FoodItem
	initializers.DB.First(&foodItem, "id = ?", data.ID)

	assert.ObjectsAreEqualValues(foodItem, data)

	assert.NotNil(t, data.DeletedAt)
}

func TestFoodDeleteFail(t *testing.T) {
	router := SetupRouter()
	dbName := "fake_deleteFoodFail_test.db"
	defer seedData.CloseConnection(dbName)

	seedData.SetupLiteDb(t, dbName)
	seedData.SeedUser(t)
	user := seedData.GetUserWithToken(t, router)
	seedData.SeedFood(t, user)

	foodItemDelete := upsertmodels.FoodItemDelete{ID: seedData.InvalidFoodId}
	foodJson, err := json.Marshal(foodItemDelete)
	if err != nil {
		t.Fatalf("Failed to marshal user data: %v", err)
	}

	w := httptest.NewRecorder()
	req, err := http.NewRequest("DELETE", "/food/delete", strings.NewReader(string(foodJson)))
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+user.Token)
	router.ServeHTTP(w, req)

	if !assert.Equal(t, http.StatusBadRequest, w.Code) {
		t.Fatalf("Wrong code, expected %v, got %v", http.StatusBadRequest, w.Code)
	}
}
