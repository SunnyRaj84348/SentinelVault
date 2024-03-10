package controllers

import (
	"SentinelVault/utilities"
	"encoding/hex"
	"io"
	"net/http"
	"os"
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
	if utilities.HandleServerError(ctx, err) {
		return
	}

	fileBytes, err := io.ReadAll(file)
	if utilities.HandleServerError(ctx, err) {
		return
	}

	keyBytes, err := utilities.GenAESKey()
	if utilities.HandleServerError(ctx, err) {
		return
	}

	cipherText, err := utilities.EncryptFile(fileBytes, keyBytes)
	if utilities.HandleServerError(ctx, err) {
		return
	}

	err = os.WriteFile("/mnt/e/enc_files/"+fileName, cipherText, 0644)
	if utilities.HandleServerError(ctx, err) {
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"key": hex.EncodeToString(keyBytes),
	})
}
