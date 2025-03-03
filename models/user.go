package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	Id           uint           `gorm:"primaryKey" json:"id"`
	Name         string         `json:"name"`
	Email        string         `gorm:"unique" json:"email"`
	Password     string         `json:"-"`
	FoodItems    []FoodItem     `json:"foodItems"`
	DailyEntries []DailyEntry   `json:"dailyEntries"`
	Token        string         `json:"token"`
	CreatedAt    time.Time      `json:"createdAt"`
	UpdatedAt    time.Time      `json:"updatedAt"`
	DeletedAt    gorm.DeletedAt `gorm:"index" json:"deletedAt"`
}
