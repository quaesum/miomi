package bleve

import (
	"github.com/blevesearch/bleve"
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
	indexServices, err := NewBleveServices()
	if err != nil {
		return err
	}
	bleveDBServices = indexServices

	return nil
}

func NewBleveProducts() (bleve.Index, error) {
	indexName := "bleve.products"
	impl := buildProductsMapping()
	index, err := buildIndex(indexName, impl)
	if err != nil {
		return nil, err
	}
	return index, nil
}

func NewBleveAnimals() (bleve.Index, error) {
	indexName := "bleve.animals"
	impl := buildAnimalsMapping()
	index, err := buildIndex(indexName, impl)
	if err != nil {
		return nil, err
	}
	return index, nil
}

func NewBleveServices() (bleve.Index, error) {
	indexName := "bleve.services"
	impl := buildServicesMapping()
	index, err := buildIndex(indexName, impl)
	if err != nil {
		return nil, err
	}
	return index, nil
}

func buildAnimalsMapping() *mapping.IndexMappingImpl {
	textFieldMapping := newTextFieldMapping()
	numericFieldMapping := newNumericFieldMapping()
	booleanFieldMapping := newBooleanFieldMapping()

	animalMapping := bleve.NewDocumentMapping()

	animalMapping.AddFieldMappingsAt("age", numericFieldMapping)
	animalMapping.AddFieldMappingsAt("name", textFieldMapping)
	animalMapping.AddFieldMappingsAt("sex", textFieldMapping)
	animalMapping.AddFieldMappingsAt("type", textFieldMapping)
	animalMapping.AddFieldMappingsAt("description", textFieldMapping)
	animalMapping.AddFieldMappingsAt("sterilized", booleanFieldMapping)
	animalMapping.AddFieldMappingsAt("vaccinated", booleanFieldMapping)
	animalMapping.AddFieldMappingsAt("shelterId", textFieldMapping)
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
	textFieldMapping := newTextFieldMapping()
	numericFieldMapping := newNumericFieldMapping()

	serviceMapping := bleve.NewDocumentMapping()

	serviceMapping.AddFieldMappingsAt("volunteer_id", numericFieldMapping)
	serviceMapping.AddFieldMappingsAt("name", textFieldMapping)
	serviceMapping.AddFieldMappingsAt("description", textFieldMapping)
	serviceMapping.AddFieldMappingsAt("photos", textFieldMapping)

	indexMapping := bleve.NewIndexMapping()
	return indexMapping
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
