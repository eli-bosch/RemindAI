package db

import (
	"fmt"

	"github.com/eli-bosch/remindAI/config"
	"github.com/eli-bosch/remindAI/internal/models"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func init() {
	db := config.Connect()

	db.AutoMigrate(
		models.User{},
		models.Reminder{},
	)

	fmt.Println("Database is connected...")
}
