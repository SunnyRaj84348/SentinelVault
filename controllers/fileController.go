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

func DownloadFile(ctx *gin.Context) {
	fileData := struct {
		FileName string `json:"filename" binding:"required"`
		Key      string `json:"key" binding:"required"`
	}{}

	err := ctx.BindJSON(&fileData)
	if err != nil {
		return
	}

	fileData.FileName = filepath.Base(fileData.FileName)

	cipherText, err := os.ReadFile("/mnt/e/enc_files/" + fileData.FileName)
	if utilities.HandleBadRequest(ctx, err) {
		return
	}

	key, err := hex.DecodeString(fileData.Key)
	if utilities.HandleBadRequest(ctx, err) {
		return
	}

	plainText, err := utilities.DecryptFile(cipherText, key)
	if utilities.HandleBadRequest(ctx, err) {
		return
	}

	ctx.Header("Content-Disposition", "attachment; filename="+fileData.FileName)
	ctx.Data(http.StatusOK, "application/octet-stream", plainText)
}
