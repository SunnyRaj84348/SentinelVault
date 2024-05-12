package controllers

import (
	"SentinelVault/models"
	"SentinelVault/utilities"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Signup(ctx *gin.Context) {
	user := models.User{}

	err := ctx.BindJSON(&user)
	if utilities.HandleBadRequest(ctx, err) {
		return
	}

	err = models.InsertUser(user)
	if err != nil {
		ctx.AbortWithStatus(http.StatusConflict)
	}
}
