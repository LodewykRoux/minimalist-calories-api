package models

import (
	"time"

	"gorm.io/gorm"
)

type DailyEntry struct {
	ID          uint           `gorm:"primaryKey" json:"id"`
	Date        time.Time      `json:"date"`
	Quantity    float32        `json:"quantity"`
	UserID      uint           `json:"userId"`
	FoodItemsId uint           `json:"foodItemsId"`
	CreatedAt   time.Time      `json:"createdAt"`
	UpdatedAt   time.Time      `json:"updatedAt"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"deletedAt"`
}
