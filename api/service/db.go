package service

import "gorm.io/gorm"

var DB *gorm.DB

// InitDB はこの処理に必要な内部処理を実行します。
func InitDB(database *gorm.DB) {
	DB = database
}
