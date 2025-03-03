package controllers

import (
	"errors"
	"minimalist-calories-api/errorHandling"
	"minimalist-calories-api/initializers"
	"minimalist-calories-api/models"
	upsertmodels "minimalist-calories-api/upsertModels"
	"net/http"

	"github.com/gin-gonic/gin"
)

func SaveFood(c *gin.Context) {
	var foodItemUpsert upsertmodels.FoodItemUpsert
	err := c.Bind(&foodItemUpsert)
	if err != nil {
		errorHandling.HandleBadRequest(c, err, "Failed to read request")
		return
	}

	user, _ := c.Get("user")

	var foodItem models.FoodItem

	if foodItemUpsert.ID != 0 {
		initializers.DB.Find(&foodItem, foodItemUpsert.ID)

	} else {
		foodItem = models.FoodItem{
			UserID: user.(models.User).Id,
		}
	}

	foodItem.Name = foodItemUpsert.Name
	foodItem.Uom = foodItemUpsert.Uom
	foodItem.Quantity = foodItemUpsert.Quantity
	foodItem.Calories = foodItemUpsert.Calories
	foodItem.Protein = foodItemUpsert.Protein
	foodItem.Carbs = foodItemUpsert.Carbs
	foodItem.Fat = foodItemUpsert.Fat

	updateResult := initializers.DB.Save(&foodItem)
	if updateResult.Error != nil {
		errorHandling.HandleBadRequest(c, updateResult.Error, "Failed to update food item")
		return
	}

	c.JSON(http.StatusOK, foodItem)
}

func DeleteFood(c *gin.Context) {
	var foodItemDelete upsertmodels.FoodItemDelete
	err := c.Bind(&foodItemDelete)
	if err != nil {
		errorHandling.HandleBadRequest(c, err, "Failed to read request")
		return
	}

	updateResult := initializers.DB.Delete(&models.FoodItem{}, foodItemDelete.ID)
	if updateResult.Error != nil {
		errorHandling.HandleBadRequest(c, updateResult.Error, "Failed to delete food item")
		return
	}

	if updateResult.RowsAffected == 0 {
		errorHandling.HandleBadRequest(c, errors.New("no record found"), "Failed to delete food item")
		return
	}

	c.JSON(http.StatusOK, gin.H{"success": updateResult.RowsAffected == 1})
}

func GetFoodList(c *gin.Context) {
	user, _ := c.Get("user")

	var foodItems []models.FoodItem

	fetchResult := initializers.DB.Where("user_id = ?", user.(models.User).Id).Find(&foodItems)
	if fetchResult.Error != nil {
		errorHandling.HandleBadRequest(c, fetchResult.Error, "No records found")
		return
	}

	c.JSON(http.StatusOK, foodItems)
}
