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

func TestWeightNewSave(t *testing.T) {
	router := SetupRouter()
	dbName := "fake_saveNewWeight_test.db"
	defer seedData.CloseConnection(dbName)

	seedData.SetupLiteDb(t, dbName)
	seedData.SeedUser(t)
	user := seedData.GetUserWithToken(t, router)
	seedData.SeedWeight(t, user)

	weight := seedData.NewWeight(t, user)

	weightJson, err := json.Marshal(weight)
	if err != nil {
		t.Fatalf("Failed to marshal user data: %v", err)
	}

	w := httptest.NewRecorder()
	req, err := http.NewRequest("POST", "/weight/save", strings.NewReader(string(weightJson)))
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+user.Token)
	router.ServeHTTP(w, req)

	if !assert.Equal(t, 200, w.Code) {
		t.Fatalf("Wrong code, expected %v, got %v", http.StatusOK, w.Code)
	}

	var data models.Weight
	err = json.Unmarshal(w.Body.Bytes(), &data)

	errorHandling.FatalCheck("Failed to unmarshal", err)

	var weightItem models.Weight
	initializers.DB.First(&weightItem, "id = ?", data.ID)

	assert.ObjectsAreEqualValues(weightItem, data)

	assert.Equal(t, weight.Date.Day(), data.Date.Day())
	assert.Equal(t, weight.Weight, data.Weight)
	assert.Equal(t, weight.UserID, data.UserID)
}

func TestWeightExistingSave(t *testing.T) {
	router := SetupRouter()
	dbName := "fake_saveExistingWeight_test.db"
	defer seedData.CloseConnection(dbName)

	seedData.SetupLiteDb(t, dbName)
	seedData.SeedUser(t)
	user := seedData.GetUserWithToken(t, router)
	existingWeight := seedData.SeedWeight(t, user)

	existingWeight.Weight = 90.0

	weightJson, err := json.Marshal(existingWeight)
	if err != nil {
		t.Fatalf("Failed to marshal user data: %v", err)
	}

	w := httptest.NewRecorder()
	req, err := http.NewRequest("POST", "/weight/save", strings.NewReader(string(weightJson)))
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+user.Token)
	router.ServeHTTP(w, req)

	if !assert.Equal(t, http.StatusOK, w.Code) {
		t.Fatalf("Wrong code, expected %v, got %v", http.StatusOK, w.Code)
	}

	var data models.Weight
	err = json.Unmarshal(w.Body.Bytes(), &data)

	errorHandling.FatalCheck("Failed to unmarshal", err)

	var weightItem models.Weight
	initializers.DB.First(&weightItem, "id = ?", data.ID)

	assert.ObjectsAreEqualValues(weightItem, data)

	assert.Equal(t, existingWeight.Date.Day(), data.Date.Day())
	assert.Equal(t, existingWeight.Weight, data.Weight)
	assert.Equal(t, existingWeight.UserID, data.UserID)
}

func TestWeightSaveFail(t *testing.T) {
	router := SetupRouter()
	dbName := "fake_saveWeightFail_test.db"
	defer seedData.CloseConnection(dbName)

	seedData.SetupLiteDb(t, dbName)
	seedData.SeedUser(t)
	user := seedData.GetUserWithToken(t, router)

	weightJson, err := json.Marshal("")
	if err != nil {
		t.Fatalf("Failed to marshal user data: %v", err)
	}

	w := httptest.NewRecorder()
	req, err := http.NewRequest("POST", "/weight/save", strings.NewReader(string(weightJson)))
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

func TestWeightGetList(t *testing.T) {
	router := SetupRouter()
	dbName := "fake_getListWeight_test.db"
	defer seedData.CloseConnection(dbName)

	seedData.SetupLiteDb(t, dbName)
	seedData.SeedUser(t)
	user := seedData.GetUserWithToken(t, router)
	seedData.SeedWeight(t, user)

	w := httptest.NewRecorder()
	req, err := http.NewRequest("GET", "/weight/getList", nil)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+user.Token)
	router.ServeHTTP(w, req)

	if !assert.Equal(t, 200, w.Code) {
		t.Fatalf("Wrong code, expected %v, got %v", http.StatusOK, w.Code)
	}

	var data []models.Weight
	err = json.Unmarshal(w.Body.Bytes(), &data)

	errorHandling.FatalCheck("Failed to unmarshal", err)

	assert.NotEmpty(t, data)
}

func TestWeightDelete(t *testing.T) {
	router := SetupRouter()
	dbName := "fake_deleteWeight_test.db"
	defer seedData.CloseConnection(dbName)

	seedData.SetupLiteDb(t, dbName)
	seedData.SeedUser(t)
	user := seedData.GetUserWithToken(t, router)
	seedData.SeedWeight(t, user)

	weightItemDelete := upsertmodels.WeightDelete{ID: seedData.ValidWeightId}
	weightJson, err := json.Marshal(weightItemDelete)
	if err != nil {
		t.Fatalf("Failed to marshal user data: %v", err)
	}

	w := httptest.NewRecorder()
	req, err := http.NewRequest("POST", "/weight/save", strings.NewReader(string(weightJson)))
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+user.Token)
	router.ServeHTTP(w, req)

	if !assert.Equal(t, 200, w.Code) {
		t.Fatalf("Wrong code, expected %v, got %v", http.StatusOK, w.Code)
	}

	var data models.Weight
	err = json.Unmarshal(w.Body.Bytes(), &data)

	errorHandling.FatalCheck("Failed to unmarshal", err)

	var weightItem models.Weight
	initializers.DB.First(&weightItem, "id = ?", data.ID)

	assert.ObjectsAreEqualValues(weightItem, data)

	assert.NotNil(t, data.DeletedAt)
}

func TestWeightDeleteFail(t *testing.T) {
	router := SetupRouter()
	dbName := "fake_deleteWeightFail_test.db"
	defer seedData.CloseConnection(dbName)

	seedData.SetupLiteDb(t, dbName)
	seedData.SeedUser(t)
	user := seedData.GetUserWithToken(t, router)
	seedData.SeedWeight(t, user)

	weightItemDelete := upsertmodels.WeightDelete{ID: seedData.InvalidWeightId}
	weightJson, err := json.Marshal(weightItemDelete)
	if err != nil {
		t.Fatalf("Failed to marshal user data: %v", err)
	}

	w := httptest.NewRecorder()
	req, err := http.NewRequest("DELETE", "/weight/delete", strings.NewReader(string(weightJson)))
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
