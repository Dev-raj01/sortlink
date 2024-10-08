package database

import (
	"fmt"

	"github.com/Dev-raj01/sortlink/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DBConn *gorm.DB

func InitDB() {
	var err error
	DBConn, err = gorm.Open(sqlite.Open("database.db"))
	if err != nil {
		panic("Cannot connect to the database")
	}
	fmt.Println("Connected to the database")
	DBConn.AutoMigrate(&models.URL{})
	fmt.Println("Migration Successful")
}
