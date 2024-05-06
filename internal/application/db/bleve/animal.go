package bleve

import (
	"github.com/blevesearch/bleve"
	"madmax/internal/entity"
)

func AddAnimal(animalID string, animal *entity.AnimalCreateBleve) error {
	err := bleveDBAnimals.Index(animalID, &animal)
	if err != nil {
		return err
	}
	return nil
}

func DeleteAnimal(id string) error {
	err := bleveDBAnimals.Delete(id)
	if err != nil {
		return err
	}
	return nil
}

func SearchAnimal(queryTerm string) (*bleve.SearchResult, error) {
	query := bleve.NewTermQuery(queryTerm)
	search := bleve.NewSearchRequest(query)
	search.Fields = []string{"age", "name", "sex", "type", "description", "sterilized", "vaccinated", "shelterId", "shelter", "address", "photos"}
	return bleveDBProducts.Search(search)
}

func SearchWOQuery() (*bleve.SearchResult, error) {
	//scrollRequest := bleve.NewScrollRequest("scroll_id", 100)
	searchRequest := bleve.NewSearchRequest(bleve.NewMatchAllQuery())
	searchRequest.Fields = []string{"age", "name", "sex", "type", "description", "sterilized", "vaccinated", "shelterId", "shelter", "address", "photos"}
	searchResult, _ := bleveDBAnimals.Search(searchRequest)
	return searchResult, nil
}
