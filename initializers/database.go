package initializers

import (
	"minimalist-calories-api/errorHandling"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectToDb() {
	var err error
	dsn := os.Getenv("DB_URL")
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	errorHandling.FatalCheck("error connecting to db", err)
}
