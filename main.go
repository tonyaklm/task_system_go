package main

import (
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"task_system_go/config"
	"task_system_go/controllers"
	"task_system_go/database"
	"task_system_go/middleware"
	"task_system_go/models"
)

func loadEnv() {
	loadedConfig, err := godotenv.Read(".env")
	if err != nil {
		panic(err)
	}
	config.Cfg.Init(loadedConfig)
}

func initDB() {
	if err := database.InitDatabase(); err != nil {
		panic(err)
	}
	err := database.Database.AutoMigrate(&models.User{}, &models.Post{})
	if err != nil {
		panic(err)
	}
}

func initRouter(router *gin.Engine) {
	usersGroup := router.Group("/users")
	{
		usersGroup.POST("/signup", controllers.Signup)
		usersGroup.POST("/login", controllers.Login)
	}

	postsChangeGroup := router.Group("/posts")
	postsChangeGroup.Use(middleware.Authenticate)
	{
		postsChangeGroup.POST("/create", controllers.CreatePost)
		postsChangeGroup.PUT("/update", controllers.UpdatePost)
		postsChangeGroup.DELETE("/:post_id", controllers.DeletePost)
	}
	postsGetGroup := router.Group("/posts")
	{
		postsGetGroup.GET("/:post_id", controllers.GetPostById)
		postsGetGroup.GET("/all", controllers.GetPosts)
	}

}

func main() {
	loadEnv()
	initDB()
	r := gin.Default()
	initRouter(r)

	err := r.Run(":8080")
	if err != nil {
		panic("[Error] failed to start Gin server due to: " + err.Error())
		return
	}
}
