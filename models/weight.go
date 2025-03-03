package models

import (
	"time"

	"gorm.io/gorm"
)

type Weight struct {
	ID        uint           `gorm:"primaryKey" json:"id"`
	Date      time.Time      `json:"date"`
	Weight    float32        `json:"weight"`
	UserID    uint           `json:"userId"`
	CreatedAt time.Time      `json:"createdAt"`
	UpdatedAt time.Time      `json:"updatedAt"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deletedAt"`
}
