package seedData

import (
	"math"
	"minimalist-calories-api/initializers"
	"minimalist-calories-api/models"
	"testing"
	"time"
)

var ValidWeightId uint = 1
var InvalidWeightId uint = math.MaxUint32

func SeedWeight(t *testing.T, user models.User) models.Weight {
	weight := models.Weight{
		ID:     ValidWeightId,
		Date:   time.Now(),
		Weight: 80.0,
		UserID: user.Id,
	}

	result := initializers.DB.Create(&weight)

	if result.Error != nil {
		t.Fatalf("Failed to setup user: %v", result.Error)
		return models.Weight{}
	}

	return weight
}

func NewWeight(t *testing.T, user models.User) models.Weight {
	weight := models.Weight{
		ID:     ValidWeightId,
		Date:   time.Now(),
		Weight: 80.0,
		UserID: user.Id,
	}

	return weight
}
