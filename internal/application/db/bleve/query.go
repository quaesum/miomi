package bleve

import (
	"github.com/blevesearch/bleve"
	"github.com/blevesearch/bleve/search/query"
	"madmax/internal/entity"
	"strconv"
	"strings"
)

type Query struct {
	*query.BooleanQuery
}

func NewQuery() *Query {
	return &Query{
		bleve.NewBooleanQuery(),
	}
}

func (q *Query) applySearchRequest(queryTerm string) {
	boolQuery := bleve.NewBooleanQuery()
	searchName := bleve.NewMatchQuery(queryTerm)
	searchName.SetField("name")
	boolQuery.AddShould(searchName)
	searchDesc := bleve.NewMatchQuery(queryTerm)
	searchDesc.SetField("description")
	boolQuery.AddShould(searchDesc)

	q.AddMust(boolQuery)
}

func (q *Query) applyAnimalsFilters(filters *entity.AnimalFilters) {
	if filters.MaxAge > 0 {
		filters.MaxAge = filters.MaxAge + 1
		field := bleve.NewNumericRangeQuery(&filters.MinAge, &filters.MaxAge)
		field.SetField("age")
		q.AddMust(field)
	}

	if len(filters.Sex) > 0 {
		boolQuery := bleve.NewBooleanQuery()
		for _, sex := range filters.Sex {
			field := bleve.NewTermQuery(sex)
			field.SetField("sex")
			boolQuery.AddShould(field)

		}
		q.AddMust(boolQuery)
	}

	if len(filters.Type) > 0 {
		boolQuery := bleve.NewBooleanQuery()
		for _, animalType := range filters.Type {
			field := bleve.NewTermQuery(strings.ToLower(animalType))
			field.SetField("type")
			boolQuery.AddShould(field)
		}
		q.AddMust(boolQuery)
	}

	if filters.Sterilized {
		field := bleve.NewBoolFieldQuery(filters.Sterilized)
		field.SetField("sterilized")
		q.AddMust(field)
	}

	if filters.Vaccinated {
		field := bleve.NewBoolFieldQuery(filters.Vaccinated)
		field.SetField("vaccinated")
		q.AddMust(field)
	}

	if len(filters.ShelterId) > 0 {
		boolQuery := bleve.NewBooleanQuery()
		for _, shelterId := range filters.ShelterId {
			shId := strconv.Itoa(shelterId)
			field := bleve.NewTermQuery(shId)
			field.SetField("shelterId")
			boolQuery.AddShould(field)
		}
		q.AddMust(boolQuery)
	}
	return
}
