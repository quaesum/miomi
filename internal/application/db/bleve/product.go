package bleve

import (
	"github.com/blevesearch/bleve"
	"madmax/internal/entity"
)

type ProductBleve struct {
	Index bleve.Index
}

func NewProductBleve() *ProductBleve {
	return &ProductBleve{
		Index: bleveDBProducts,
	}
}

func (p *ProductBleve) Add(productID string, product *entity.ProductCreateBleve) error {
	err := p.Index.Index(productID, &product)
	if err != nil {
		return err
	}
	return nil
}

func (p *ProductBleve) Remove(id string) error {
	err := p.Index.Delete(id)
	if err != nil {
		return err
	}
	return nil
}

func (p *ProductBleve) Search(queryTerm string) (*bleve.SearchResult, error) {
	query := bleve.NewTermQuery(queryTerm)
	search := bleve.NewSearchRequest(query)
	search.Fields = []string{"name", "description", "photos", "link"}
	return p.Index.Search(search)
}

func (p *ProductBleve) SearchWOQuery() (*bleve.SearchResult, error) {
	//scrollRequest := bleve.NewScrollRequest("scroll_id", 100)
	searchRequest := bleve.NewSearchRequest(bleve.NewMatchAllQuery())
	searchRequest.Fields = []string{"name", "description", "photos", "link"}
	searchResult, _ := p.Index.Search(searchRequest)
	return searchResult, nil
}
