package controller

import (
	"necolog/model"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func AdminIndex(c *gin.Context) {
	articles, err := model.GetArticles()
	if err != nil {
		c.HTML(500, "500.tmpl", gin.H{})
		return
	}
	c.HTML(200, "home.tmpl", gin.H{"articles": articles})
}

func Login(c *gin.Context) {
	c.HTML(200, "login.tmpl", gin.H{})
}

func LoginPost(c *gin.Context) {
	user := model.User{}
	user.Email = c.PostForm("email")
	user.Password = c.PostForm("password")

	user, err := user.Login()
	if err != nil {
		c.HTML(401, "login.tmpl", gin.H{"error": err.Error(), "user": user})
		return
	}

	session := sessions.Default(c)
	session.Set("user_id", user.Id)
	session.Save()

	c.Redirect(302, "/admin")
}

func Logout(c *gin.Context) {
	session := sessions.Default(c)
	session.Clear()
	session.Save()
	c.Redirect(302, "/admin/login")
}
