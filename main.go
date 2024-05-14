package main

import (
	"SentinelVault/controllers"
	"SentinelVault/middlewares"
	"SentinelVault/models"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}

	router := gin.Default()

	router.Use(middlewares.Cors())
	router.Use(middlewares.Sessions())

	models.Connect()
	defer models.Close()

	auth := router.Group("/", middlewares.Auth)
	{
		auth.POST("/upload", controllers.UploadFile)
		auth.GET("/download", controllers.DownloadFile)
		auth.GET("/get-files", controllers.GetFilesData)
	}

	router.POST("/signup", controllers.Signup)
	router.POST("/login", controllers.Login)

	err = router.Run(":8080")
	if err != nil {
		log.Fatal(err)
	}
}
