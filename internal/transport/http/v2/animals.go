package v2

import (
	"github.com/gin-gonic/gin"
	"madmax/internal/application"
	"madmax/internal/entity"
	"madmax/internal/utils"
	"net/http"
)

type AnimalsHttp struct {
	*application.AnimalApplication
}

func NewAnimalsHttp() *AnimalsHttp {
	return &AnimalsHttp{
		application.NewAnimalApplication(),
	}
}

func (a *AnimalsHttp) GetAll(c *gin.Context) {
	var req entity.SearchRequest
	var err error
	c.ShouldBindJSON(&req)

	if req.Page <= 0 {
		req.Page = 1
	}
	maxPossibleResults := (req.Page * req.PerPage) - 1
	var animals []entity.AnimalBleve
	animals, err = a.GetFromBleve(&req, maxPossibleResults)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
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

	if left > len(animals) || right > len(animals) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
		return
	}

	resp := entity.SearchAnimalsResponseV2{
		Animals: animals[left:right],
		MaxPage: maxPages,
	}
	c.JSON(200, resp)
	return
}
