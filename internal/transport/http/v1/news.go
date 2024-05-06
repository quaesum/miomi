package v1

import (
	"context"
	"github.com/gin-gonic/gin"
	"madmax/internal/application"
	"madmax/internal/entity"
	"net/http"
	"strconv"
	"time"
)

func getAllNewsHandler(c *gin.Context) {
	ctx := context.Background()
	tctx, _ := context.WithTimeout(ctx, time.Second*15)
	news, err := application.NewsAll(tctx)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, news)
	return
}

func createNewsHandler(c *gin.Context) {
	var ncr entity.NewsCreateRequest
	if err := c.ShouldBindJSON(&ncr); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx := context.Background()
	tctx, _ := context.WithTimeout(ctx, time.Second*5)
	newsId, err := application.NewsCreate(tctx, &ncr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"id": newsId})
	return
}

func updateNewsHandler(c *gin.Context) {
	id := c.Param("id")
	newsID, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	var ncr entity.NewsCreateRequest
	if err := c.ShouldBindJSON(&ncr); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx := context.Background()
	tctx, _ := context.WithTimeout(ctx, time.Second*5)
	err = application.NewsUpdate(tctx, newsID, &ncr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{})
	return
}

func removeNewsHandler(c *gin.Context) {
	id := c.Param("id")
	newsId, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx := context.Background()
	tctx, _ := context.WithTimeout(ctx, time.Second*5)
	err = application.RemoveNewsById(tctx, newsId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{})
	return
}
