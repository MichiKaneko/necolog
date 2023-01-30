package controller

import (
	"necolog/model"

	"github.com/gin-gonic/gin"
)

func Index(c *gin.Context) {
	articles, err := model.GetArticles()
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, articles)
}