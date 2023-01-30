package controller

import (
	"necolog/model"
	"strconv"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func GetArticles(c *gin.Context) {
	articles, err := model.GetArticles()
	if err != nil {
		c.HTML(404, "404.tmpl", gin.H{})
		return
	}

	session := sessions.Default(c)
	userId := session.Get("user_id")

	if userId != nil {
		c.HTML(200, "articles/index.tmpl", gin.H{"articles": articles, "user_id": userId})
		return
	}
	c.HTML(200, "articles/index.tmpl", gin.H{"articles": articles})
}

func GetArticle(c *gin.Context) {
	id := c.Param("id")
	id_int, err := strconv.Atoi(id)
	if err != nil {
		c.HTML(404, "404.tmpl", gin.H{})
		return
	}

	article, err := model.GetArticleById(id_int)
	if err != nil {
		c.HTML(404, "404.tmpl", gin.H{})
		return
	}

	session := sessions.Default(c)
	userId := session.Get("user_id")

	if userId != nil {
		c.HTML(200, "articles/show.tmpl", gin.H{"article": article, "user_id": userId})
		return
	}
	c.HTML(200, "articles/show.tmpl", gin.H{"article": article})
}

func CreateArticlePage(c *gin.Context) {
	c.HTML(200, "new.tmpl", gin.H{})
}

func CreateArticle(c *gin.Context) {
	title := c.PostForm("title")
	body := c.PostForm("body")

	article := model.Article{Title: title, Body: body}
	err := article.Create()

	if err != nil {
		c.HTML(302, "articles/new.tmpl", gin.H{
			"error":   err.Error(),
			"article": article,
		})
		return
	}
	c.Redirect(302, "/article/"+strconv.Itoa(article.Id))
}

func UpdateArticlePage(c *gin.Context) {
	id := c.Param("id")
	id_int, err := strconv.Atoi(id)
	if err != nil {
		c.HTML(404, "404.tmpl", gin.H{})
		return
	}

	article, err := model.GetArticleById(id_int)
	if err != nil {
		c.HTML(404, "404.tmpl", gin.H{})
		return
	}
	c.HTML(200, "articles/edit.tmpl", gin.H{"article": article})
}

func UpdateArticle(c *gin.Context) {
	id := c.Param("id")
	id_int, err := strconv.Atoi(id)
	if err != nil {
		c.HTML(404, "404.tmpl", gin.H{})
		return
	}

	article, err := model.GetArticleById(id_int)
	if err != nil {
		c.HTML(404, "404.tmpl", gin.H{})
		return
	}

	title := c.PostForm("title")
	body := c.PostForm("body")

	article.Title = title
	article.Body = body

	err = article.Update()
	if err != nil {
		c.HTML(302, "articles/edit.tmpl", gin.H{
			"error":   err.Error(),
			"article": article,
		})
		return
	}
	c.Redirect(302, "/admin/article/"+strconv.Itoa(article.Id))
}

func DeleteArticle(c *gin.Context) {
	id := c.Param("id")
	id_int, err := strconv.Atoi(id)
	if err != nil {
		c.HTML(404, "404.tmpl", gin.H{})
		return
	}

	article, err := model.GetArticleById(id_int)
	if err != nil {
		c.HTML(404, "404.tmpl", gin.H{})
		return
	}

	err = article.Delete()
	if err != nil {
		c.Redirect(302, "/admin/article")
		return
	}
	c.Redirect(302, "/admin/article")
}
