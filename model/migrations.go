package model

import (
	"necolog/db"
)

func Migrate() error {
	var models = []interface{}{
		&User{},
		&Article{},
	}
	return db.Debug().AutoMigrate(models...)
}
