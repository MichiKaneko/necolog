package controller

import (
	"necolog/model"

	"github.com/gin-gonic/gin"
)

func Index(c *gin.Context) {
	articles, err := model.GetArticles()
	if err != nil {
		c.HTML(500, "500.tmpl", gin.H{})
		return
	}
	c.HTML(200, "index.tmpl", gin.H{"articles": articles})
}
