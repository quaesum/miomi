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

func createProductHandler(c *gin.Context) {
	var pcr entity.ProductCreateRequest
	if err := c.ShouldBindJSON(&pcr); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx := context.Background()
	tctx, _ := context.WithTimeout(ctx, time.Minute)
	productID, err := application.ProductCreate(tctx, &pcr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"id": productID})
	return
}

func getProductByIDHandler(c *gin.Context) {
	id := c.Param("id")
	productID, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx := context.Background()
	tctx, _ := context.WithTimeout(ctx, time.Second*5)
	product, err := application.ProductByID(tctx, productID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, product)
	return
}

func getProductsHandler(c *gin.Context) {
	var req entity.SearchRequest
	var err error
	c.ShouldBindJSON(&req)

	if req.Page <= 0 {
		req.Page = 1
	}

	products, err := application.GetProductsFromBleve(req.Request)
	//if req.Request != "" {
	//	products, err = application.GetProductsSearchResult(req.Request, products)
	//}

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	//maxPages, err := application.GetMaxPages(len(products), req.PerPage)
	//if err != nil {
	//	c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	//	return
	//}

	//left, right, err := application.GetRecordsOnCurrentPage(req, len(products))
	//if err != nil {
	//	c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	//	return
	//}

	//resp := entity.SearchProductsResponse{
	//	Products: products[left:right],
	//	MaxPage:  maxPages,
	//}
	c.JSON(200, products)
	return
}

func updateProductHandler(c *gin.Context) {
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
	err = application.ProductUpdate(tctx, serviceID, &pcr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{})
	return
}

func removeProductHandler(c *gin.Context) {
	id := c.Param("id")
	serviceID, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx := context.Background()
	tctx, _ := context.WithTimeout(ctx, time.Second*5)
	err = application.RemoveProductByID(tctx, serviceID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{})
	return
}
