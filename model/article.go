package model

import (
	"necolog/db"

	"gorm.io/gorm"
)

type Article struct {
	gorm.Model
	Id          int    `gorm:"primaryKey"`
	Title       string `gorm:"not null" json:"title"`
	Description string `json:"description"`
	Body        string `json:"body"`
	Image       string `json:"image"`
}

func (a *Article) TableName() string {
	return "articles"
}

func (a *Article) Create() error {
	err := db.Debug().Model(&Article{}).Create(&a).Error
	if err != nil {
		return err
	}
	return nil
}

func (a *Article) Update() error {
	err := db.Debug().Model(&Article{}).Updates(&a).Error
	if err != nil {
		return err
	}
	return nil
}

func (a *Article) Delete() error {
	err := db.Debug().Model(&Article{}).Delete(&a).Error
	if err != nil {
		return err
	}
	return nil
}

func GetArticles() ([]*Article, error) {
	var articles []*Article
	err := db.Debug().Model(&Article{}).Limit(100).Find(&articles).Error
	if err != nil {
		return articles, err
	}
	return articles, nil
}

func GetArticleById(id int) (Article, error) {
	var article Article
	err := db.Debug().Model(&Article{}).Where("id = ?", id).Take(&article).Error
	if err != nil {
		return article, err
	}
	return article, nil
}
