package controllers

import (
	"SentinelVault/models"
	"SentinelVault/utilities"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Signup(ctx *gin.Context) {
	cred := models.User{}

	err := ctx.BindJSON(&cred)
	if utilities.HandleBadRequest(ctx, err) {
		return
	}

	err = models.InsertUser(cred)
	if err != nil {
		ctx.AbortWithStatus(http.StatusConflict)
	}
}

func Login(ctx *gin.Context) {
	cred := models.User{}

	err := ctx.BindJSON(&cred)
	if utilities.HandleBadRequest(ctx, err) {
		return
	}

	user, err := models.GetUser(cred.Username)
	if (err != nil) || (user.Password != cred.Password) {
		ctx.AbortWithStatus(http.StatusUnauthorized)
		return
	}
}
