package main

import (
	"minimalist-calories-api/initializers"
	"minimalist-calories-api/models"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDb()
}

func main() {
	//TODO: fix migration stuffs
	initializers.DB.AutoMigrate(&models.User{}, &models.FoodItem{}, &models.DailyEntry{}, &models.Weight{})

	initializers.DB.Migrator().CreateConstraint(&models.User{}, "FoodItems")
	initializers.DB.Migrator().CreateConstraint(&models.User{}, "fk_users_food_items")

	initializers.DB.Migrator().CreateConstraint(&models.User{}, "DailyEntries")
	initializers.DB.Migrator().CreateConstraint(&models.User{}, "fk_users_daily_entries")

	initializers.DB.Migrator().CreateConstraint(&models.User{}, "Weight")
	initializers.DB.Migrator().CreateConstraint(&models.User{}, "fk_users_weight")

	initializers.DB.Migrator().CreateConstraint(&models.DailyEntry{}, "FoodItem")
	initializers.DB.Migrator().CreateConstraint(&models.DailyEntry{}, "fk_daily_entry_food_item")
}
