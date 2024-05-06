package application

import (
	"github.com/samber/lo"
	"madmax/internal/entity"
	"sort"
	"strings"
)

func GetAnimalsSearchResult(searchTerm string, animals []entity.Animal) ([]entity.Animal, error) {
	searchTerm = cleanQuery(searchTerm)

	for i := range animals {
		animals[i].Score = calculateAnimalsScore(animals[i], searchTerm)
	}

	sort.Slice(animals, func(i, j int) bool {
		return animals[i].Score > animals[j].Score
	})

	animals = lo.Filter(animals, func(animal entity.Animal, _ int) bool {
		return animal.Score > 0
	})

	return animals, nil
}

func calculateAnimalsScore(animal entity.Animal, searchTerm string) int64 {
	searchTerm = strings.ToLower(searchTerm)
	nameCount := strings.Count(strings.ToLower(animal.Name), searchTerm)
	descriptionCount := strings.Count(strings.ToLower(animal.Description), searchTerm)
	return int64(nameCount + descriptionCount)
}

func calculateServicesScore(service entity.Service, searchTerm string) int64 {
	searchTerm = strings.ToLower(searchTerm)
	nameCount := strings.Count(strings.ToLower(service.Label), searchTerm)
	descriptionCount := strings.Count(strings.ToLower(service.Description), searchTerm)
	return int64(nameCount + descriptionCount)
}

func calculateProductsScore(service entity.Product, searchTerm string) int64 {
	searchTerm = strings.ToLower(searchTerm)
	nameCount := strings.Count(strings.ToLower(service.Name), searchTerm)
	descriptionCount := strings.Count(strings.ToLower(service.Description), searchTerm)
	return int64(nameCount + descriptionCount)
}

func cleanQuery(query string) string {
	words := strings.Fields(query)

	cleanedWords := lo.Filter(words, func(item string, _ int) bool {
		return len(item) > 2
	})

	cleanedQuery := strings.Join(cleanedWords, " ")
	query = strings.TrimSpace(cleanedQuery)

	return cleanedQuery
}
