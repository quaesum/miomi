package bleve

import (
	"github.com/blevesearch/bleve"
	"madmax/internal/entity"
)

type AnimalBleve struct {
	Index             bleve.Index
	AnimalCreateBleve *entity.AnimalCreateBleve
	AnimalsBleve      *entity.AnimalBleve
}

func NewAnimal() *AnimalBleve {
	return &AnimalBleve{
		Index: bleveDBAnimals,
	}
}

func (a AnimalBleve) Add(animalID string, animal *entity.AnimalCreateBleve) error {
	err := a.Index.Index(animalID, &animal)
	if err != nil {
		return err
	}
	return nil
}

func (a AnimalBleve) Remove(id string) error {
	err := a.Index.Delete(id)
	if err != nil {
		return err
	}
	return nil
}

func (a AnimalBleve) Search(req *entity.SearchRequest, limit int) (*bleve.SearchResult, error) {
	var search *bleve.SearchRequest

	sr := req.Request
	fl := req.Filters
	if sr != "" || !fl.IsEmpty() {
		q := NewQuery()
		if sr != "" {
			q.applySearchRequest(req.Request)
		}
		if !fl.IsEmpty() {
			q.applyAnimalsFilters(&req.Filters)
		}
		search = bleve.NewSearchRequest(q)
	} else {
		search = bleve.NewSearchRequest(bleve.NewMatchAllQuery())
	}
	search.Size = limit
	search.Fields = []string{"age", "name", "sex", "type", "description", "sterilized", "vaccinated", "shelterId", "shelter", "address", "photos"}
	return a.Index.Search(search)
}
