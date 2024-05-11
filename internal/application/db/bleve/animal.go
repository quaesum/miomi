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

func (a AnimalBleve) Search(queryTerm string) (*bleve.SearchResult, error) {
	query := bleve.NewTermQuery(queryTerm)
	search := bleve.NewSearchRequest(query)
	search.Fields = []string{"age", "name", "sex", "type", "description", "sterilized", "vaccinated", "shelterId", "shelter", "address", "photos"}
	return a.Index.Search(search)
}

func (a AnimalBleve) SearchWOQuery() (*bleve.SearchResult, error) {
	//scrollRequest := bleve.NewScrollRequest("scroll_id", 100)
	searchRequest := bleve.NewSearchRequest(bleve.NewMatchAllQuery())
	searchRequest.Fields = []string{"age", "name", "sex", "type", "description", "sterilized", "vaccinated", "shelterId", "shelter", "address", "photos"}
	searchResult, _ := a.Index.Search(searchRequest)
	return searchResult, nil
}
