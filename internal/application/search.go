package application

import (
	"context"
	"github.com/samber/lo"
	"madmax/internal/application/db/mysql"
	"madmax/internal/entity"
	"sort"
	"strings"
)

func GetAllAnimals(ctx context.Context) ([]entity.Animal, error) {
	animals, err := mysql.GetAllAnimals(ctx)
	if err != nil {
		return nil, err
	}

	return animals, err
}

func GetSearchResult(searchTerm string, animals []entity.Animal) ([]entity.Animal, error) {
	searchTerm = cleanQuery(searchTerm)

	for i := range animals {
		animals[i].Score = calculateScore(animals[i], searchTerm)
	}

	sort.Slice(animals, func(i, j int) bool {
		return animals[i].Score > animals[j].Score
	})

	animals = lo.Filter(animals, func(animal entity.Animal, _ int) bool {
		return animal.Score > 0
	})

	return animals, nil
}

func calculateScore(animal entity.Animal, searchTerm string) int64 {
	searchTerm = strings.ToLower(searchTerm)
	nameCount := strings.Count(strings.ToLower(animal.Name), searchTerm)
	descriptionCount := strings.Count(strings.ToLower(animal.Description), searchTerm)
	return int64(nameCount + descriptionCount)
}

func cleanQuery(query string) string {
	words := strings.Fields(query)

	cleanedWords := lo.Filter(words, func(item string, _ int) bool {
		return len(item) > 2
	})

	// Собираем очищенный запрос из слов
	cleanedQuery := strings.Join(cleanedWords, " ")
	query = strings.TrimSpace(cleanedQuery)

	return cleanedQuery
}
