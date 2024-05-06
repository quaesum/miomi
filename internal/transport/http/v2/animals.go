package v2

import (
	"github.com/gin-gonic/gin"
	"madmax/internal/application"
	"madmax/internal/entity"
	"net/http"
)

func getAnimalsHandler(c *gin.Context) {
	var req entity.SearchRequest
	var err error
	c.ShouldBindJSON(&req)

	if req.Page <= 0 {
		req.Page = 1
	}
	var animals []entity.AnimalsBleve
	if req.Request == "" {
		animals, err = application.GetAllAnimalsFromBleve()
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}
	} else {
		animals, err = application.GetAnimalsFromBleve(req.Request)
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

	resp := entity.SearchAnimalsResponseV2{
		Animals: animals[left:right],
		MaxPage: maxPages,
	}
	c.JSON(200, resp)
	return
}
