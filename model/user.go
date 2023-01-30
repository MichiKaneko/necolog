package model

import (
	"necolog/db"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Id       int    `gorm:"primaryKey"`
	Email    string `gorm:"not null;unique" json:"email"`
	Password string `gorm:"not null" json:"password"`
}

func (u *User) TableName() string {
	return "users"
}

func findUserByEmail(email string) (User, error) {
	var user User
	err := db.Debug().Model(&User{}).Where("email = ?", email).Take(&user).Error
	if err != nil {
		return user, err
	}
	return user, nil
}

func (u *User) Create() error {

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	u.Password = string(hashedPassword)

	err = db.Debug().Model(&User{}).Create(&u).Error
	if err != nil {
		return err
	}
	return nil
}

func (u *User) Login() (User, error) {
	user, err := findUserByEmail(u.Email)
	if err != nil {
		return user, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(u.Password))
	if err != nil {
		return user, err
	}

	return user, nil
}
