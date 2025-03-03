package models

import (
	"time"

	"gorm.io/gorm"
)

type FoodItem struct {
	ID        uint           `gorm:"primaryKey" json:"id"`
	Name      string         `json:"name"`
	Uom       int            `gorm:"uom" json:"uoM"`
	Quantity  float32        `json:"quantity"`
	Calories  float32        `json:"calories"`
	Protein   float32        `json:"protein"`
	Carbs     float32        `json:"carbs"`
	Fat       float32        `json:"fat"`
	UserID    uint           `json:"userId"`
	CreatedAt time.Time      `json:"createdAt"`
	UpdatedAt time.Time      `json:"updatedAt"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deletedAt"`
}
