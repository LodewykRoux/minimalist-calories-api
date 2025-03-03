package seedData

import (
	"math"
	"minimalist-calories-api/initializers"
	"minimalist-calories-api/models"
	"testing"
	"time"
)

var ValidDailyEntryId uint = 1
var InvalidDailyEntryId uint = math.MaxUint32

func SeedDailyEntry(t *testing.T, user models.User) models.DailyEntry {
	dailyEntry := models.DailyEntry{
		ID:          ValidDailyEntryId,
		Date:        time.Now(),
		Quantity:    1,
		UserID:      user.Id,
		FoodItemsId: ValidFoodId,
	}

	result := initializers.DB.Create(&dailyEntry)

	if result.Error != nil {
		t.Fatalf("Failed to setup daily entry: %v", result.Error)
		return models.DailyEntry{}
	}

	return dailyEntry
}

func NewDailyEntry(t *testing.T, user models.User) models.DailyEntry {
	dailyEntry := models.DailyEntry{
		ID:          ValidDailyEntryId,
		Date:        time.Now(),
		Quantity:    1,
		UserID:      user.Id,
		FoodItemsId: ValidFoodId,
	}

	return dailyEntry
}
