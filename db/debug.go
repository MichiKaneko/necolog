package db

import (
	"gorm.io/gorm"
)

func Debug() *gorm.DB {
	return Db.Debug()
}
