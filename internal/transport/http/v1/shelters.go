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

func getShelterByIDHandler(c *gin.Context) {
	id := c.Param("id")
	userID, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx := context.Background()
	tctx, _ := context.WithTimeout(ctx, time.Minute*2)
	user, err := application.ShelterByID(tctx, userID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"data": user})
	return
}
func getAllSheltersHandler(c *gin.Context) {
	ctx := context.Background()
	tctx, _ := context.WithTimeout(ctx, time.Second*15)
	animals, err := application.SheltersAll(tctx)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, animals)
	return
}
func createShelterHandler(c *gin.Context) {
	var scr entity.ShelterCreateRequest
	if err := c.ShouldBindJSON(&scr); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	ctx := context.Background()
	tctx, _ := context.WithTimeout(ctx, time.Second*15)
	shID, err := application.ShelterCreate(tctx, &scr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	c.JSON(200, gin.H{"id": shID})
}
func updateShelterHandler(c *gin.Context) {
	id := c.Param("id")
	shID, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	var scr entity.ShelterCreateRequest
	if err := c.ShouldBindJSON(&scr); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	ctx := context.Background()
	tctx, _ := context.WithTimeout(ctx, time.Second*15)

	err = application.ShelterUpdate(tctx, shID, &scr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	c.JSON(200, gin.H{})
}

func removeShelterHandler(c *gin.Context) {
	id := c.Param("id")
	shID, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	ctx := context.Background()
	tctx, _ := context.WithTimeout(ctx, time.Second*15)

	err = application.RemoveNewsById(tctx, shID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	c.JSON(200, gin.H{})
}
