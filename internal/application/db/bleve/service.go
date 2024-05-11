package bleve

import (
	"github.com/blevesearch/bleve"
	"madmax/internal/entity"
)

type ServiceBleve struct {
	Index bleve.Index
}

func NewSerivceBleve() *ServiceBleve {
	return &ServiceBleve{
		Index: bleveDBServices,
	}
}

func (s *ServiceBleve) Add(serviceID string, service *entity.ServiceCreateBleve) error {
	err := s.Index.Index(serviceID, &service)
	if err != nil {
		return err
	}
	return nil
}

func (s *ServiceBleve) Remove(id string) error {
	err := s.Index.Delete(id)
	if err != nil {
		return err
	}
	return nil
}

func (s *ServiceBleve) Search(queryTerm string) (*bleve.SearchResult, error) {
	query := bleve.NewTermQuery(queryTerm)
	search := bleve.NewSearchRequest(query)
	search.Fields = []string{"volunteer_id", "name", "description", "photos"}
	return s.Index.Search(search)
}

func (s *ServiceBleve) SearchWOQuery() (*bleve.SearchResult, error) {
	//scrollRequest := bleve.NewScrollRequest("scroll_id", 100)
	searchRequest := bleve.NewSearchRequest(bleve.NewMatchAllQuery())
	searchRequest.Fields = []string{"volunteer_id", "name", "description", "photos"}
	searchResult, _ := s.Index.Search(searchRequest)
	return searchResult, nil
}
