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

func SaveDailyEntry(c *gin.Context) {
	var dailyEntryUpsert upsertmodels.DailyEntryUpsert
	err := c.Bind(&dailyEntryUpsert)
	if err != nil {
		errorHandling.HandleBadRequest(c, err, "Failed to update daily entry")
		return
	}

	user, _ := c.Get("user")

	var dailyEntry models.DailyEntry

	if dailyEntryUpsert.ID != 0 {
		findResult := initializers.DB.Find(&dailyEntry, dailyEntryUpsert.ID)

		if findResult.Error != nil {
			errorHandling.HandleBadRequest(c, findResult.Error, "Failed to find entry")
			return
		}
	} else {
		dailyEntry =
			models.DailyEntry{
				UserID: user.(models.User).Id,
			}
	}

	dailyEntry.Date = dailyEntryUpsert.Date
	dailyEntry.FoodItemsId = dailyEntryUpsert.FoodItemsId
	dailyEntry.Quantity = float32(dailyEntryUpsert.Quantity)

	updateResult := initializers.DB.Save(&dailyEntry)
	if updateResult.Error != nil {
		errorHandling.HandleBadRequest(c, updateResult.Error, "Failed to update entry")
		return
	}

	c.JSON(http.StatusOK, dailyEntry)
}

func DeleteDailyEntry(c *gin.Context) {
	var dailyEntryDelete upsertmodels.DailyEntryDelete
	err := c.Bind(&dailyEntryDelete)
	if err != nil {
		errorHandling.HandleBadRequest(c, err, "Failed to read request")
		return
	}

	updateResult := initializers.DB.Delete(&models.DailyEntry{}, dailyEntryDelete.ID)

	if updateResult.Error != nil {
		errorHandling.HandleBadRequest(c, updateResult.Error, "Failed to delete daily entry")
		return
	}

	if updateResult.RowsAffected == 0 {
		errorHandling.HandleBadRequest(c, errors.New("failed to delete item"), "Failed to delete daily entry")
		return
	}

	c.JSON(http.StatusOK, gin.H{"success": updateResult.RowsAffected == 1})
}

func GetDailyEntryList(c *gin.Context) {
	user, _ := c.Get("user")

	var dailyEntries []models.DailyEntry

	fetchResult := initializers.DB.Where("user_id = ?", user.(models.User).Id).Find(&dailyEntries)
	if fetchResult.Error != nil {
		errorHandling.HandleBadRequest(c, fetchResult.Error, "No records found")
		return
	}

	c.JSON(http.StatusOK, dailyEntries)
}
