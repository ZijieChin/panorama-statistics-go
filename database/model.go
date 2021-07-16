package database

import (
	"gorm.io/gorm"
	"time"
)

type Statistics struct {
	ID        int `gorm:"primaryKey"`
	User      string
	Page      string
	CreatedAt time.Time
}

var DB *gorm.DB

func SyncDatabase() {
	DB = ConnectDatabase()
	DB.AutoMigrate(&Statistics{})
}
