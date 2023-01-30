package model

import (
	"necolog/db"
)

func Migrate() error {
	var models = []interface{}{
		&Article{},
	}
	return db.Debug().AutoMigrate(models...)
}
