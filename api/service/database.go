package service

import "gorm.io/gorm"

var database *gorm.DB

func SetDatabase(db *gorm.DB) {
	database = db
}
