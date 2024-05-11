package v1

import (
	"context"
	"github.com/gin-gonic/gin"
	"madmax/internal/application"
	"madmax/internal/entity"
	"madmax/internal/utils"
	"net/http"
	"strconv"
	"time"
)

type AnimalsHttp struct {
	app *application.AnimalApplication
}

func NewAnimalsHttp() *AnimalsHttp {
	return &AnimalsHttp{
		app: application.NewAnimalApplication(),
	}
}

func (a *AnimalsHttp) GetByID(c *gin.Context) {
	id := c.Param("id")
	animalID, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx := context.Background()
	tctx, _ := context.WithTimeout(ctx, time.Second*5)
	animals, err := a.app.GetByID(tctx, animalID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, animals)
	c.Done()
	return
}

func (a *AnimalsHttp) GetAll(c *gin.Context) {
	var req entity.SearchRequest
	var err error
	c.ShouldBindJSON(&req)

	if req.Page <= 0 {
		req.Page = 1
	}
	ctx := context.Background()
	animals, err := a.app.GetAllFromMYSQL(ctx)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	if req.Request != "" {
		animals, err = application.GetAnimalsSearchResult(req.Request, animals)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}
	}

	maxPages, err := utils.GetMaxPages(len(animals), req.PerPage)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	left, right, err := utils.GetRecordsOnCurrentPage(req, len(animals))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	resp := entity.SearchAnimalsResponse{
		Animals: animals[left:right],
		MaxPage: maxPages,
	}
	c.JSON(200, resp)
	c.Done()
	return
}

func (a *AnimalsHttp) Create(c *gin.Context) {
	var ucr entity.AnimalCreateRequest
	if err := c.ShouldBindJSON(&ucr); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx := context.Background()
	tctx, _ := context.WithTimeout(ctx, time.Second*5)
	animalID, err := a.app.Create(tctx, &ucr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"id": animalID})
	c.Done()
	return
}
func (a *AnimalsHttp) Update(c *gin.Context) {
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
	err = a.app.Update(tctx, animalID, &ucr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{})
	c.Done()
	return
}

func (a *AnimalsHttp) Remove(c *gin.Context) {
	id := c.Param("id")
	animalID, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx := context.Background()
	tctx, _ := context.WithTimeout(ctx, time.Second*5)
	err = a.app.Remove(tctx, animalID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{})
	c.Done()
	return
}
