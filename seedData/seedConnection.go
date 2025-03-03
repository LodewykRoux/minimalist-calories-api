package seedData

import (
	"database/sql/driver"
	"minimalist-calories-api/initializers"
	"minimalist-calories-api/models"
	"os"
	"testing"
	"time"

	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
)

type AnyTime struct{}
type AnyValue struct{}

func (a AnyTime) Match(v driver.Value) bool {
	_, ok := v.(time.Time)
	return ok || v == nil
}

func (a AnyValue) Match(v driver.Value) bool {
	return true
}

func SetupLiteDb(t *testing.T, dbName string) {
	var err error
	initializers.DB, err = gorm.Open(sqlite.Open(dbName), &gorm.Config{
		// Logger: logger.Default.LogMode(logger.Info), // Turn on logging here
	})
	if err != nil {
		t.Fatalf("Failed to setup gorm mock: %v", err)
	}

	initializers.DB.AutoMigrate(&models.User{}, &models.FoodItem{}, &models.DailyEntry{}, &models.Weight{})
}

func CloseConnection(dbName string) {
	db, _ := initializers.DB.DB()
	db.Close()
	os.Remove(dbName)
}
