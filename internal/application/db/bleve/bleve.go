package bleve

import (
	"github.com/blevesearch/bleve"
	"github.com/blevesearch/bleve/analysis/lang/ru"
	"github.com/blevesearch/bleve/mapping"
)

var (
	bleveDBProducts bleve.Index
	bleveDBAnimals  bleve.Index
	bleveDBServices bleve.Index
)

func NewBleve() error {
	indexProducts, err := NewBleveProducts()
	if err != nil {
		return err
	}
	bleveDBProducts = indexProducts

	indexAnimals, err := NewBleveAnimals()
	if err != nil {
		return err
	}
	bleveDBAnimals = indexAnimals
	bleveDBServices = indexProducts

	return nil
}

func NewBleveProducts() (bleve.Index, error) {
	indexName := "bleve.products"
	impl := buildProductsMapping()
	index, err := buildIndex(indexName, impl)
	if err != nil {
		return nil, err
	}
	bleveDBProducts = index
	return index, nil
}

func NewBleveAnimals() (bleve.Index, error) {
	indexName := "bleve.animals"
	impl := buildAnimalsMapping()
	index, err := buildIndex(indexName, impl)
	if err != nil {
		return nil, err
	}
	bleveDBAnimals = index
	return index, nil
}

//func NewBleveServices() (bleve.Index, error) {
//	indexName := "belve.services"
//	index, err := buildIndex(indexName)
//	if err != nil {
//		return nil, err
//	}
//	bleveDBServices = index
//	return index, nil
//}

//func buildIndex(indexName string, impl *mapping.IndexMappingImpl) (bleve.Index, error) {
//	index, err := bleve.Open(indexName)
//	if err == bleve.ErrorIndexPathDoesNotExist {
//		kvStore := goleveldb.Name
//		kvConfig := map[string]interface{}{
//			"create_if_missing": true,
//			//		"write_buffer_size":         536870912,
//			//		"lru_cache_capacity":        536870912,
//			//		"bloom_filter_bits_per_key": 10,
//		}
//
//		index, err = bleve.NewUsing(indexName, impl, "upside_down", kvStore, kvConfig)
//	}
//	if err != nil {
//		return nil, err
//	}
//	return index, nil
//}

func buildAnimalsMapping() *mapping.IndexMappingImpl {
	textFieldMapping := newTextFieldMapping()
	numericFieldMapping := newNumericFieldMapping()
	booleanFieldMapping := newBooleanFieldMapping()

	animalMapping := bleve.NewDocumentMapping()

	animalMapping.AddFieldMappingsAt("age", numericFieldMapping)
	animalMapping.AddFieldMappingsAt("name", textFieldMapping)
	animalMapping.AddFieldMappingsAt("sex", numericFieldMapping)
	animalMapping.AddFieldMappingsAt("type", textFieldMapping)
	animalMapping.AddFieldMappingsAt("description", textFieldMapping)
	animalMapping.AddFieldMappingsAt("sterilized", booleanFieldMapping)
	animalMapping.AddFieldMappingsAt("vaccinated", booleanFieldMapping)
	animalMapping.AddFieldMappingsAt("shelterId", numericFieldMapping)
	animalMapping.AddFieldMappingsAt("shelter", textFieldMapping)
	animalMapping.AddFieldMappingsAt("address", textFieldMapping)
	animalMapping.AddFieldMappingsAt("photos", textFieldMapping)

	indexMapping := bleve.NewIndexMapping()
	indexMapping.DefaultMapping = animalMapping
	return indexMapping
}

func buildProductsMapping() *mapping.IndexMappingImpl {
	textFieldMapping := newTextFieldMapping()

	productMapping := bleve.NewDocumentMapping()

	productMapping.AddFieldMappingsAt("name", textFieldMapping)
	productMapping.AddFieldMappingsAt("description", textFieldMapping)
	productMapping.AddFieldMappingsAt("photos", textFieldMapping)
	productMapping.AddFieldMappingsAt("link", textFieldMapping)

	indexMapping := bleve.NewIndexMapping()
	indexMapping.DefaultMapping = productMapping
	return indexMapping
}

func buildServicesMapping() *mapping.IndexMappingImpl {
	return nil
}

func buildMapping() *mapping.IndexMappingImpl {
	ruFieldMapping := bleve.NewTextFieldMapping()
	ruFieldMapping.Analyzer = ru.AnalyzerName

	eventMapping := bleve.NewDocumentMapping()
	eventMapping.AddFieldMappingsAt("name", ruFieldMapping)

	mapping := bleve.NewIndexMapping()
	mapping.DefaultMapping = eventMapping
	mapping.DefaultAnalyzer = ru.AnalyzerName
	return mapping
}

func buildIndex(indexName string, impl *mapping.IndexMappingImpl) (bleve.Index, error) {

	index, err := bleve.Open(indexName)
	if err == bleve.ErrorIndexPathDoesNotExist {
		index, err = bleve.New(indexName, impl)
		if err != nil {
			return nil, err
		}
	}
	if err != nil {
		return nil, err
	}

	return index, nil
}

func newTextFieldMapping() *mapping.FieldMapping {
	textFieldMapping := bleve.NewTextFieldMapping()
	textFieldMapping.Store = true
	textFieldMapping.IncludeTermVectors = true
	textFieldMapping.IncludeInAll = true
	return textFieldMapping
}

func newNumericFieldMapping() *mapping.FieldMapping {
	numericFieldMapping := bleve.NewNumericFieldMapping()
	numericFieldMapping.Store = true
	numericFieldMapping.IncludeInAll = true
	return numericFieldMapping
}

func newBooleanFieldMapping() *mapping.FieldMapping {
	booleanFieldMapping := bleve.NewBooleanFieldMapping()
	booleanFieldMapping.Store = true
	booleanFieldMapping.IncludeInAll = true
	return booleanFieldMapping
}
