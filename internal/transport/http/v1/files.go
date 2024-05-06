package v1

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"madmax/internal/application"
	"net/http"
	"path/filepath"
	"strings"
	"time"
)

func attachAnimalFileHandler(c *gin.Context) {
	attachFile(c, "animals")
}

func attachNewsFileHandler(c *gin.Context) {
	attachFile(c, "news")
}

func attachProductFileHandler(c *gin.Context) {
	attachFile(c, "products")
}

func attachServiceFileHandler(c *gin.Context) {
	attachFile(c, "service")
}

func attachFile(c *gin.Context, target string) {
	file, err := c.FormFile("file")

	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "No file is received",
		})
		return
	}

	// Retrieve file information
	newFileName := uuid.New().String()
	ctx := context.Background()
	f, err := file.Open()

	resp, err := application.AddFile(ctx, file.Size, fmt.Sprintf("%s%s", newFileName, strings.ToLower(filepath.Ext(file.Filename))), f, target)
	if err != nil {
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error add file": err.Error()})
			return
		}
	}
	c.JSON(200, gin.H{
		"data": resp,
	})
	return
}

func getAllFileNamesAndIdsHandler(c *gin.Context) {
	ctx := context.Background()
	tctx, _ := context.WithTimeout(ctx, time.Second*10)
	photos, err := application.GetFileNameAndId(tctx)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, photos)
	return
}
