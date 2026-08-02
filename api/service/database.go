package service

import "gorm.io/gorm"

var database *gorm.DB

// SetDatabase はこの処理に必要な内部処理を実行します。
func SetDatabase(db *gorm.DB) {
	database = db
}
