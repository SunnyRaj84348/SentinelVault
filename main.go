package main

import (
	"SentinelVault/controllers"
	"SentinelVault/models"
	"log"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	router := gin.Default()

	router.Use(cors.Default())

	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}

	models.Connect()
	defer models.Close()

	router.POST("/upload", controllers.UploadFile)
	router.GET("/download", controllers.DownloadFile)

	router.POST("/signup", controllers.Signup)
	router.POST("/login", controllers.Login)

	err = router.Run(":8080")
	if err != nil {
		log.Fatal(err)
	}
}
