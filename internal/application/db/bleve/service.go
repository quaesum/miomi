package bleve

import (
	"github.com/blevesearch/bleve"
	"madmax/internal/entity"
)

func AddService(productID string, product *entity.ProductCreateBleve) error {
	err := bleveDBServices.Index(productID, &product)
	if err != nil {
		return err
	}
	return nil
}

func DeleteService(id string) error {
	err := bleveDBServices.Delete(id)
	if err != nil {
		return err
	}
	return nil
}

func SearchService(queryTerm string) (*bleve.SearchResult, error) {
	query := bleve.NewTermQuery(queryTerm)
	search := bleve.NewSearchRequest(query)
	search.Fields = []string{"id", "name", "description", "photos", "link"}
	return bleveDBServices.Search(search)
}
