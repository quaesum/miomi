package application

import (
	"context"
	"database/sql"
	"errors"
	"madmax/internal/application/db/mysql"
	"madmax/internal/entity"
)

func AnimalsAll(ctx context.Context) ([]entity.Animal, error) {
	return mysql.GetAllAnimals(ctx)
}
func AnimalByID(ctx context.Context, id int64) (*entity.Animal, error) {
	animal, err := mysql.GetAnimalBasicInfo(ctx, id)
	if err != nil && err == sql.ErrNoRows {
		return nil, errors.New("animal not exist")
	}
	mysql.GetPhotosByAnimalID(ctx, id)
	photos, err := mysql.GetPhotosByAnimalID(ctx, id)
	if err != nil && err != sql.ErrNoRows {
		return nil, err
	}
	animal.Photos = photos
	return animal, nil
}

func NewsAll(ctx context.Context) ([]entity.News, error) {
	return mysql.GetNewsInfo(ctx)
}
