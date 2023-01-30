package model

import (
	"os"
)

func Seed() error {
	adminUser, err := findUserByEmail(os.Getenv("ADMIN_EMAIL"))
	if err != nil {
		return err
	}
	if adminUser.Email == os.Getenv("ADMIN_EMAIL") {
		return nil
	}

	var user = User{
		Email:    os.Getenv("ADMIN_EMAIL"),
		Password: os.Getenv("ADMIN_PASSWORD"),
	}

	if err := user.Create(); err != nil {
		return err
	}
	return nil
}
