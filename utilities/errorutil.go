package utilities

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func HandleServerError(ctx *gin.Context, err error) bool {
	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return true
	}

	return false
}

func HandleBadRequest(ctx *gin.Context, err error) bool {
	if err != nil {
		ctx.AbortWithStatus(http.StatusBadRequest)
		return true
	}

	return false
}
