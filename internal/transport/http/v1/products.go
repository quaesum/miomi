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

type ProductsHttp struct {
	app *application.ProductApplication
}

func NewProductsHttp() *ProductsHttp {
	return &ProductsHttp{
		app: application.NewProductApplication(),
	}
}

func (p *ProductsHttp) Create(c *gin.Context) {
	var pcr entity.ProductCreateRequest
	if err := c.ShouldBindJSON(&pcr); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx := context.Background()
	tctx, _ := context.WithTimeout(ctx, time.Minute)
	productID, err := p.app.Create(tctx, &pcr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"id": productID})
	return
}

func (p *ProductsHttp) GetByID(c *gin.Context) {
	id := c.Param("id")
	productID, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx := context.Background()
	tctx, _ := context.WithTimeout(ctx, time.Second*5)
	product, err := p.app.GetByID(tctx, productID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, product)
	return
}

func (p *ProductsHttp) GetAll(c *gin.Context) {
	var req entity.SearchRequest
	var err error
	c.ShouldBindJSON(&req)

	if req.Page <= 0 {
		req.Page = 1
	}
	if req.PerPage <= 0 {
		req.PerPage = 21
	}
	var products []entity.ProductSearch
	if req.Request == "" {
		products, err = p.app.GetAllFromBleve()
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"search": err.Error()})
			return
		}
	} else {
		products, err = p.app.GetFromBleve(req.Request)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"search": err.Error()})
			return
		}
	}

	maxPages, err := utils.GetMaxPages(len(products), req.PerPage)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"pages": err.Error()})
		return
	}

	left, right, err := utils.GetRecordsOnCurrentPage(req, len(products))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"pages": err.Error()})
		return
	}
	resp := entity.SearchProductsResponse{
		Products: products[left:right],
		MaxPage:  maxPages,
	}
	c.JSON(200, resp)
	return
}

func (p *ProductsHttp) Update(c *gin.Context) {
	id := c.Param("id")
	serviceID, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	var pcr entity.ProductCreateRequest
	if err := c.ShouldBindJSON(&pcr); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx := context.Background()
	tctx, _ := context.WithTimeout(ctx, time.Second*5)
	err = p.app.Update(tctx, serviceID, &pcr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{})
	return
}

func (p *ProductsHttp) Remove(c *gin.Context) {
	id := c.Param("id")
	serviceID, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx := context.Background()
	tctx, _ := context.WithTimeout(ctx, time.Second*5)
	err = p.app.Remove(tctx, serviceID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{})
	return
}
