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

type ReportsHttp struct {
	Controller
	app application.ReportApp
}

func NewReportsHttp() *ReportsHttp {
	return &ReportsHttp{}
}

func (r *ReportsHttp) GetAll(c *gin.Context) {
	ctx := context.Background()
	tctx, _ := context.WithTimeout(ctx, 5*time.Second)

	reports, err := r.app.GetAll(tctx)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"reports": reports})
	return
}

func (r *ReportsHttp) GetByID(c *gin.Context) {
	id := c.Param("id")
	reportID, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx := context.Background()
	tctx, _ := context.WithTimeout(ctx, 1*time.Second)

	report, err := r.app.GetByID(tctx, reportID)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"report": report})
	return
}

func (r *ReportsHttp) Create(c *gin.Context) {
	userID, err := utils.GetUserID(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var ucr entity.ReportCreateRequest
	if err = c.ShouldBindJSON(&ucr); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx := context.Background()
	tctx, _ := context.WithTimeout(ctx, 1*time.Second)

	id, err := r.app.Create(tctx, &ucr, userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, id)
	return
}

func (r *ReportsHttp) Remove(c *gin.Context) {
	id := c.Param("id")
	reportID, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx := context.Background()
	tctx, _ := context.WithTimeout(ctx, 1*time.Second)

	err = r.app.Remove(tctx, reportID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
}
