package controllers

import (
	"SentinelVault/models"
	"SentinelVault/utilities"
	"encoding/hex"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strconv"

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

	fileID, err := models.InsertFile(fileName, userid.(int64))
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

	err = os.WriteFile("/mnt/e/enc_files/"+strconv.Itoa(int(fileID)), cipherText, 0644)
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
		FileID int64  `json:"file_id" binding:"required"`
		EncKey string `json:"enc_key" binding:"required"`
	}{}

	err := ctx.BindJSON(&fileQuery)
	if utilities.HandleBadRequest(ctx, err) {
		return
	}

	fileData, err := models.GetFile(fileQuery.FileID, userid.(int64))
	if err != nil {
		ctx.AbortWithStatus(http.StatusNotFound)
		return
	}

	cipherText, err := os.ReadFile("/mnt/e/enc_files/" + strconv.Itoa(int(fileData.FileID)))
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

	ctx.Header("Content-Disposition", "attachment; filename="+fileData.Filename)
	ctx.Data(http.StatusOK, "application/octet-stream", plainText)
}

func GetFilesData(ctx *gin.Context) {
	userid, _ := ctx.Get("userid")

	filesData, err := models.GetAllFiles(userid.(int64))
	if utilities.HandleServerError(ctx, err) {
		return
	}

	ctx.JSON(http.StatusOK, filesData)
}

func ShareFile(ctx *gin.Context) {
	userid, _ := ctx.Get("userid")

	fileInfo := struct {
		FileID   int64  `json:"file_id" binding:"required"`
		UserName string `json:"username" binding:"required"`
	}{}

	err := ctx.BindJSON(&fileInfo)
	if utilities.HandleBadRequest(ctx, err) {
		return
	}

	targetUser, err := models.GetUser(fileInfo.UserName)
	if err != nil {
		ctx.AbortWithStatus(http.StatusNotFound)
		return
	}

	if userid.(int64) == targetUser.UserID {
		ctx.AbortWithStatus(http.StatusConflict)
		return
	}

	_, err = models.GetFile(fileInfo.FileID, userid.(int64))
	if err != nil {
		ctx.AbortWithStatus(http.StatusForbidden)
		return
	}

	err = models.InsertSharedFile(fileInfo.FileID, targetUser.UserID)
	if err != nil {
		ctx.AbortWithStatus(http.StatusConflict)
	}
}
