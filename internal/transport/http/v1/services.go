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

func createServiceHandler(c *gin.Context) {
	userID, err := utils.GetUserID(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	var csr entity.ServiceCreateRequest
	if err := c.ShouldBindJSON(&csr); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx := context.Background()
	tctx, _ := context.WithTimeout(ctx, time.Second*5)
	serviceID, err := application.ServiceCreate(tctx, userID, &csr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"id": serviceID})
	return
}

func getServiceByIDHandler(c *gin.Context) {
	id := c.Param("id")
	serviceID, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx := context.Background()
	tctx, _ := context.WithTimeout(ctx, time.Second*5)
	service, err := application.ServiceByID(tctx, serviceID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, service)
	return
}

func getServicesHandler(c *gin.Context) {
	var req entity.SearchRequest
	var err error
	c.ShouldBindJSON(&req)

	if req.Page <= 0 {
		req.Page = 1
	}

	ctx := context.Background()
	tctx, _ := context.WithTimeout(ctx, time.Second*15)

	services, err := application.GetAllServices(tctx)
	//if req.Request != "" {
	//	services, err = application.GetServicesSearchResult(req.Request, services)
	//}

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	maxPages, err := application.GetMaxPages(len(services), req.PerPage)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	left, right, err := application.GetRecordsOnCurrentPage(req, len(services))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	resp := entity.SearchServicesResponse{
		Services: services[left:right],
		MaxPage:  maxPages,
	}
	c.JSON(200, resp)
	return
}

func updateServiceHandler(c *gin.Context) {
	userID, err := utils.GetUserID(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	id := c.Param("id")
	serviceID, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	var csr entity.ServiceCreateRequest
	if err := c.ShouldBindJSON(&csr); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx := context.Background()
	tctx, _ := context.WithTimeout(ctx, time.Second*5)
	err = application.ServiceUpdate(tctx, userID, serviceID, &csr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{})
	return
}

func removeServiceHandler(c *gin.Context) {
	id := c.Param("id")
	serviceID, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx := context.Background()
	tctx, _ := context.WithTimeout(ctx, time.Second*5)
	err = application.RemoveServiceByID(tctx, serviceID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{})
	return
}
