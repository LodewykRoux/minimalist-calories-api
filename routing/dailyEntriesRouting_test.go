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

func TestDailyEntryNewSave(t *testing.T) {
	router := SetupRouter()
	dbName := "fake_saveNewDailyEntry_test.db"
	defer seedData.CloseConnection(dbName)

	seedData.SetupLiteDb(t, dbName)
	seedData.SeedUser(t)
	user := seedData.GetUserWithToken(t, router)
	seedData.SeedDailyEntry(t, user)

	dailyEntry := seedData.NewDailyEntry(t, user)

	dailyEntriesJson, err := json.Marshal(dailyEntry)
	if err != nil {
		t.Fatalf("Failed to marshal user data: %v", err)
	}

	w := httptest.NewRecorder()
	req, err := http.NewRequest("POST", "/dailyEntries/save", strings.NewReader(string(dailyEntriesJson)))
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+user.Token)
	router.ServeHTTP(w, req)

	if !assert.Equal(t, 200, w.Code) {
		t.Fatalf("Wrong code, expected %v, got %v", http.StatusOK, w.Code)
	}

	var data models.DailyEntry
	err = json.Unmarshal(w.Body.Bytes(), &data)

	errorHandling.FatalCheck("Failed to unmarshal", err)

	var dailyEntriesItem models.DailyEntry
	initializers.DB.First(&dailyEntriesItem, "id = ?", data.ID)

	assert.ObjectsAreEqualValues(dailyEntriesItem, data)

	assert.Equal(t, dailyEntry.Date.Day(), data.Date.Day())
	assert.Equal(t, dailyEntry.FoodItemsId, data.FoodItemsId)
	assert.Equal(t, dailyEntry.Quantity, data.Quantity)
	assert.Equal(t, dailyEntry.UserID, data.UserID)
}

func TestDailyEntryExistingSave(t *testing.T) {
	router := SetupRouter()
	dbName := "fake_saveExistingDailyEntry_test.db"
	defer seedData.CloseConnection(dbName)

	seedData.SetupLiteDb(t, dbName)
	seedData.SeedUser(t)
	user := seedData.GetUserWithToken(t, router)
	existingDailyEntry := seedData.SeedDailyEntry(t, user)

	existingDailyEntry.Quantity = 2.0

	dailyEntryJson, err := json.Marshal(existingDailyEntry)
	if err != nil {
		t.Fatalf("Failed to marshal user data: %v", err)
	}

	w := httptest.NewRecorder()
	req, err := http.NewRequest("POST", "/dailyEntries/save", strings.NewReader(string(dailyEntryJson)))
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+user.Token)
	router.ServeHTTP(w, req)

	if !assert.Equal(t, http.StatusOK, w.Code) {
		t.Fatalf("Wrong code, expected %v, got %v", http.StatusOK, w.Code)
	}

	var data models.DailyEntry
	err = json.Unmarshal(w.Body.Bytes(), &data)

	errorHandling.FatalCheck("Failed to unmarshal", err)

	var dailyEntryItem models.DailyEntry
	initializers.DB.First(&dailyEntryItem, "id = ?", data.ID)

	assert.ObjectsAreEqualValues(dailyEntryItem, data)

	assert.Equal(t, existingDailyEntry.Date.Day(), data.Date.Day())
	assert.Equal(t, existingDailyEntry.Quantity, data.Quantity)
	assert.Equal(t, existingDailyEntry.FoodItemsId, data.FoodItemsId)
	assert.Equal(t, existingDailyEntry.UserID, data.UserID)
}

func TestDailyEntrySaveFail(t *testing.T) {
	router := SetupRouter()
	dbName := "fake_saveDailyEntryFail_test.db"
	defer seedData.CloseConnection(dbName)

	seedData.SetupLiteDb(t, dbName)
	seedData.SeedUser(t)
	user := seedData.GetUserWithToken(t, router)

	dailyEntryJson, err := json.Marshal("")
	if err != nil {
		t.Fatalf("Failed to marshal user data: %v", err)
	}

	w := httptest.NewRecorder()
	req, err := http.NewRequest("POST", "/dailyEntries/save", strings.NewReader(string(dailyEntryJson)))
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

func TestDailyEntryGetList(t *testing.T) {
	router := SetupRouter()
	dbName := "fake_getListDailyEntry_test.db"
	defer seedData.CloseConnection(dbName)

	seedData.SetupLiteDb(t, dbName)
	seedData.SeedUser(t)
	user := seedData.GetUserWithToken(t, router)
	seedData.SeedDailyEntry(t, user)

	w := httptest.NewRecorder()
	req, err := http.NewRequest("GET", "/dailyEntries/getList", nil)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+user.Token)
	router.ServeHTTP(w, req)

	if !assert.Equal(t, 200, w.Code) {
		t.Fatalf("Wrong code, expected %v, got %v", http.StatusOK, w.Code)
	}

	var data []models.DailyEntry
	err = json.Unmarshal(w.Body.Bytes(), &data)

	errorHandling.FatalCheck("Failed to unmarshal", err)

	assert.NotEmpty(t, data)
}

func TestDailyEntryDelete(t *testing.T) {
	router := SetupRouter()
	dbName := "fake_deleteDailyEntry_test.db"
	defer seedData.CloseConnection(dbName)

	seedData.SetupLiteDb(t, dbName)
	seedData.SeedUser(t)
	user := seedData.GetUserWithToken(t, router)
	seedData.SeedDailyEntry(t, user)

	dailyEntriesItemDelete := upsertmodels.DailyEntryDelete{ID: seedData.ValidDailyEntryId}
	dailyEntriesJson, err := json.Marshal(dailyEntriesItemDelete)
	if err != nil {
		t.Fatalf("Failed to marshal user data: %v", err)
	}

	w := httptest.NewRecorder()
	req, err := http.NewRequest("POST", "/dailyEntries/save", strings.NewReader(string(dailyEntriesJson)))
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+user.Token)
	router.ServeHTTP(w, req)

	if !assert.Equal(t, 200, w.Code) {
		t.Fatalf("Wrong code, expected %v, got %v", http.StatusOK, w.Code)
	}

	var data models.DailyEntry
	err = json.Unmarshal(w.Body.Bytes(), &data)

	errorHandling.FatalCheck("Failed to unmarshal", err)

	var dailyEntriesItem models.DailyEntry
	initializers.DB.First(&dailyEntriesItem, "id = ?", data.ID)

	assert.ObjectsAreEqualValues(dailyEntriesItem, data)

	assert.NotNil(t, data.DeletedAt)
}

func TestDailyEntryDeleteFail(t *testing.T) {
	router := SetupRouter()
	dbName := "fake_deleteDailyEntryFail_test.db"
	defer seedData.CloseConnection(dbName)

	seedData.SetupLiteDb(t, dbName)
	seedData.SeedUser(t)
	user := seedData.GetUserWithToken(t, router)
	seedData.SeedDailyEntry(t, user)

	dailyEntryItemDelete := upsertmodels.DailyEntryDelete{ID: seedData.InvalidDailyEntryId}
	dailyEntryJson, err := json.Marshal(dailyEntryItemDelete)
	if err != nil {
		t.Fatalf("Failed to marshal user data: %v", err)
	}

	w := httptest.NewRecorder()
	req, err := http.NewRequest("DELETE", "/dailyEntries/delete", strings.NewReader(string(dailyEntryJson)))
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
