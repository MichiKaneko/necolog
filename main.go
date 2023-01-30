package main

import (
	"fmt"
	"log"
	"necolog/controller"
	"necolog/db"
	"necolog/middleware"
	"necolog/model"
	"os"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

func main() {
	r := Router()
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	r.Run(fmt.Sprintf(":%s", port))
}

func init() {
	err := LoadEnv()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	err = db.Connect(DatabaseConfig())
	if err != nil {
		log.Fatal("Error connecting to database")
	}

	err = model.Migrate()
	if err != nil {
		log.Fatal("Error migrating database")
	}

	err = model.Seed()
	if err != nil {
		log.Fatal("Error seeding database")
	}

}

func Router() *gin.Engine {
	r := gin.New()

	r.Use(gin.Logger(), gin.Recovery())

	secret := os.Getenv("SECRET")
	store := cookie.NewStore([]byte(secret))
	r.Use(sessions.Sessions("necolog_admin", store))

	r.LoadHTMLGlob("views/**/*")

	r.Static("/static", "./static")

	r.GET("/", controller.Index)

	article := r.Group("/article")
	{
		article.GET("/", controller.GetArticles)
		article.GET("/:id", controller.GetArticle)
	}

	r.GET("/admin/login", controller.Login)
	r.POST("/admin/login", controller.LoginPost)

	admin := r.Group("/admin")
	admin.Use(middleware.AuthCheckMiddleware())
	{
		admin.GET("/", controller.AdminIndex)
		admin.GET("/logout", controller.Logout)

		admin_article := admin.Group("/article")
		{
			admin_article.GET("/", controller.GetArticles)
			admin_article.GET("/:id", controller.GetArticle)
			admin_article.GET("/create", controller.CreateArticlePage)
			admin_article.POST("/create", controller.CreateArticle)
			admin_article.GET("/:id/update", controller.UpdateArticlePage)
			admin_article.POST("/:id/update", controller.UpdateArticle)
			admin_article.POST("/:id/delete", controller.DeleteArticle)
		}
	}

	return r
}
