package seedData

import (
	"math"
	"minimalist-calories-api/initializers"
	"minimalist-calories-api/models"
	"testing"
)

var ValidFoodId uint = 1
var InvalidFoodId uint = math.MaxUint32

func SeedFood(t *testing.T, user models.User) models.FoodItem {
	food := models.FoodItem{
		ID:       ValidFoodId,
		Name:     "Food 1",
		Uom:      1,
		Quantity: 100,
		Calories: 100,
		Protein:  25,
		Carbs:    20,
		Fat:      5,
		UserID:   user.Id,
	}

	result := initializers.DB.Create(&food)

	if result.Error != nil {
		t.Fatalf("Failed to setup user: %v", result.Error)
		return models.FoodItem{}
	}

	return food
}

func NewFood(t *testing.T, user models.User) models.FoodItem {
	food := models.FoodItem{
		Name:     "Food New",
		Uom:      1,
		Quantity: 100,
		Calories: 100,
		Protein:  25,
		Carbs:    20,
		Fat:      5,
		UserID:   user.Id,
	}

	return food
}
