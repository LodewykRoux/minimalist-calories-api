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

func SaveWeight(c *gin.Context) {
	var weightUpsert upsertmodels.WeightUpsert
	err := c.Bind(&weightUpsert)
	if err != nil {
		errorHandling.HandleBadRequest(c, err, "Failed to read request")
		return
	}

	user, _ := c.Get("user")

	var weightEntry models.Weight

	if weightUpsert.ID != 0 {
		initializers.DB.Find(&weightEntry, weightUpsert.ID)
	} else {
		weightEntry =
			models.Weight{
				UserID: user.(models.User).Id,
			}
	}

	weightEntry.Date = weightUpsert.Date
	weightEntry.Weight = float32(weightUpsert.Weight)

	updateResult := initializers.DB.Save(&weightEntry)
	if updateResult.Error != nil {
		errorHandling.HandleBadRequest(c, updateResult.Error, "Failed to update weight")
		return
	}

	c.JSON(http.StatusOK, weightEntry)
}

func DeleteWeight(c *gin.Context) {
	var weightDelete upsertmodels.WeightDelete
	err := c.Bind(&weightDelete)
	if err != nil {
		errorHandling.HandleBadRequest(c, err, "Failed to read request")
		return
	}

	updateResult := initializers.DB.Delete(&models.Weight{}, weightDelete.ID)
	if updateResult.Error != nil {
		errorHandling.HandleBadRequest(c, updateResult.Error, "Failed to delete weight")
		return
	}

	if updateResult.RowsAffected == 0 {
		errorHandling.HandleBadRequest(c, errors.New("failed to delete weight"), "Failed to delete weight")
		return
	}

	c.JSON(http.StatusOK, gin.H{"success": updateResult.RowsAffected == 1})
}

func GetWeightList(c *gin.Context) {
	user, _ := c.Get("user")

	var weights []models.Weight

	fetchResult := initializers.DB.Where("user_id = ?", user.(models.User).Id).Find(&weights)
	if fetchResult.Error != nil {
		errorHandling.HandleBadRequest(c, fetchResult.Error, "No records found")
		return
	}

	c.JSON(http.StatusOK, weights)
}
