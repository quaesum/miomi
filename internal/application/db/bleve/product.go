package bleve

import (
	"github.com/blevesearch/bleve"
	"madmax/internal/entity"
)

func AddProduct(productID string, product *entity.ProductCreateBleve) error {
	err := bleveDBProducts.Index(productID, &product)
	if err != nil {
		return err
	}
	return nil
}

func DeleteProduct(id string) error {
	err := bleveDBProducts.Delete(id)
	if err != nil {
		return err
	}
	return nil
}

func SearchProducts(queryTerm string) (*bleve.SearchResult, error) {
	query := bleve.NewTermQuery(queryTerm)
	search := bleve.NewSearchRequest(query)
	search.Fields = []string{"name", "description", "photos", "link"}
	return bleveDBProducts.Search(search)
}
