package application

import (
	"context"
	"madmax/internal/application/db/mysql"
	"madmax/internal/entity"
)

func AnimalsAll(ctx context.Context) ([]entity.Animal, error) {
	return mysql.GetAllAnimals(ctx)
}

func NewsAll(ctx context.Context) ([]entity.News, error) {
	return mysql.GetNewsInfo(ctx)
}
