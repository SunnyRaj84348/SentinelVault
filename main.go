package main

import (
	"SentinelVault/controllers"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.POST("/upload", controllers.UploadFile)
	router.GET("/download", controllers.DownloadFile)

	err := router.Run(":8080")
	if err != nil {
		log.Fatal(err)
	}
}
