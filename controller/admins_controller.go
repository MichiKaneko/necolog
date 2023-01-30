package controller

import (
	"necolog/model"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	c.HTML(200, "login.html", gin.H{})
}

func LoginPost(c *gin.Context) {
	user := model.User{}
	user.Email = c.PostForm("email")
	user.Password = c.PostForm("password")

	user, err := user.Login()
	if err != nil {
		c.HTML(401, "login.html", gin.H{"error": err.Error(), "user": user})
		return
	}

	session := sessions.Default(c)
	session.Set("user_id", user.Id)
	session.Save()

	c.Redirect(302, "/admin")
}
