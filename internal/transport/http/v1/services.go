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

type ServiceHttp struct {
	app *application.ServiceApplication
}

func NewServicesHttp() *ServiceHttp {
	return &ServiceHttp{
		app: application.NewServiceApplication(),
	}
}

func (s *ServiceHttp) Create(c *gin.Context) {
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
	tctx, _ := context.WithTimeout(ctx, time.Minute*5)
	serviceID, err := s.app.Create(tctx, userID, &csr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"id": serviceID})
	return
}

func (s *ServiceHttp) GetByID(c *gin.Context) {
	id := c.Param("id")
	serviceID, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx := context.Background()
	tctx, _ := context.WithTimeout(ctx, time.Second*5)
	service, err := s.app.GetByID(tctx, serviceID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, service)
	return
}

func (s *ServiceHttp) GetAll(c *gin.Context) {
	var req entity.SearchRequest
	var err error
	c.ShouldBindJSON(&req)

	if req.Page <= 0 {
		req.Page = 1
	}
	var services []entity.ServiceBleve
	if req.Request == "" {
		services, err = s.app.GetAllFromBleve()
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}
	} else {
		services, err = s.app.GetFromBleve(req.Request)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}
	}

	maxPages, err := utils.GetMaxPages(len(services), req.PerPage)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	left, right, err := utils.GetRecordsOnCurrentPage(req, len(services))
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

func (s *ServiceHttp) Update(c *gin.Context) {
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
	err = s.app.Update(tctx, userID, serviceID, &csr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{})
	return
}

func (s *ServiceHttp) Remove(c *gin.Context) {
	id := c.Param("id")
	serviceID, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx := context.Background()
	tctx, _ := context.WithTimeout(ctx, time.Second*5)
	err = s.app.Remove(tctx, serviceID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{})
	return
}
