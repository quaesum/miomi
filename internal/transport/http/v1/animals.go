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

func getAnimalByIDHandler(c *gin.Context) {
	id := c.Param("id")
	animalID, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx := context.Background()
	tctx, _ := context.WithTimeout(ctx, time.Second*5)
	animals, err := application.AnimalByID(tctx, animalID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, animals)
	return
}

func getAnimalsHandler(c *gin.Context) {
	var req entity.SearchRequest
	var err error
	c.ShouldBindJSON(&req)

	if req.Page <= 0 {
		req.Page = 1
	}
	ctx := context.Background()
	animals, err := application.GetAllAnimalsFromMysql(ctx)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	if req.Request != "" {
		animals, err = application.GetAnimalsSearchResult(req.Request, animals)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}
	}

	maxPages, err := application.GetMaxPages(len(animals), req.PerPage)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	left, right, err := application.GetRecordsOnCurrentPage(req, len(animals))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	resp := entity.SearchAnimalsResponse{
		Animals: animals[left:right],
		MaxPage: maxPages,
	}
	c.JSON(200, resp)
	return
}

func createAnimalHandler(c *gin.Context) {
	var ucr entity.AnimalCreateRequest
	if err := c.ShouldBindJSON(&ucr); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx := context.Background()
	tctx, _ := context.WithTimeout(ctx, time.Second*5)
	animalID, err := application.AnimalCreate(tctx, &ucr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"id": animalID})
	return
}
func updateAnimalHandler(c *gin.Context) {
	id := c.Param("id")
	animalID, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	var ucr entity.AnimalCreateRequest
	if err := c.ShouldBindJSON(&ucr); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx := context.Background()
	tctx, _ := context.WithTimeout(ctx, time.Second*5)
	err = application.AnimalUpdate(tctx, animalID, &ucr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{})
	return
}

func removeAnimalHandler(c *gin.Context) {
	id := c.Param("id")
	animalID, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx := context.Background()
	tctx, _ := context.WithTimeout(ctx, time.Second*5)
	err = application.RemoveAnimalByID(tctx, animalID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{})
	return
}
