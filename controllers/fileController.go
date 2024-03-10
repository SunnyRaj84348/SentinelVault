package controllers

import (
	"io"
	"net/http"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

func UploadFile(ctx *gin.Context) {
	fileHeader, err := ctx.FormFile("file")
	if err != nil {
		ctx.AbortWithStatus(http.StatusBadRequest)
		return
	}

	fileName := filepath.Base(fileHeader.Filename)

	file, err := fileHeader.Open()
	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	fileBytes, err := io.ReadAll(file)
	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}
}
