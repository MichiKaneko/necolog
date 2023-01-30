package middleware

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func AuthCheckMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		userId := session.Get("user_id")

		if userId == nil {
			c.HTML(401, "404.html", gin.H{})
			c.Abort()
			return
		}
		c.Next()
	}
}