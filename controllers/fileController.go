package controllers

import (
	"SentinelVault/models"
	"SentinelVault/utilities"
	"crypto/sha256"
	"encoding/hex"
	"io"
	"net/http"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

func UploadFile(ctx *gin.Context) {
	userid, _ := ctx.Get("userid")

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

	hashByte := sha256.Sum256(fileBytes)
	hash := hex.EncodeToString(hashByte[:])

	fileMeta := models.File{Filename: fileName, FileHash: hash}

	fileID, err := models.InsertFile(fileMeta, userid.(string))
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

	err = os.WriteFile("/mnt/e/enc_files/"+fileID, cipherText, 0644)
	if utilities.HandleServerError(ctx, err) {
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"file_id": fileID,
		"enc_key": hex.EncodeToString(keyBytes),
	})
}

func DownloadFile(ctx *gin.Context) {
	userid, _ := ctx.Get("userid")

	fileQuery := struct {
		FileID string `json:"file_id" binding:"required"`
		EncKey string `json:"enc_key" binding:"required"`
	}{}

	err := ctx.BindJSON(&fileQuery)
	if utilities.HandleBadRequest(ctx, err) {
		return
	}

	fileData, err := models.GetFile(fileQuery.FileID, userid.(string))
	if err != nil {
		ctx.AbortWithStatus(http.StatusNotFound)
		return
	}

	cipherText, err := os.ReadFile("/mnt/e/enc_files/" + fileData.FileID)
	if utilities.HandleServerError(ctx, err) {
		return
	}

	key, err := hex.DecodeString(fileQuery.EncKey)
	if utilities.HandleBadRequest(ctx, err) {
		return
	}

	plainText, err := utilities.DecryptFile(cipherText, key)
	if utilities.HandleBadRequest(ctx, err) {
		return
	}

	hashByte := sha256.Sum256(plainText)
	hash := hex.EncodeToString(hashByte[:])

	if hash != fileData.FileHash {
		ctx.AbortWithStatus(http.StatusConflict)
		return
	}

	ctx.Header("Content-Disposition", "attachment; filename="+fileData.Filename)
	ctx.Data(http.StatusOK, "application/octet-stream", plainText)
}
