package application

import (
	"context"
	"madmax/internal/application/db/mysql"
	"madmax/internal/entity"
)

func ShelterByID(ctx context.Context, id int64) (*entity.Shelter, error) {
	return mysql.GetShelterByID(ctx, id)
}

func SheltersAll(ctx context.Context) ([]entity.Shelter, error) {
	return mysql.GetAllShelters(ctx)
}

func ShelterCreate(ctx context.Context, shelter *entity.ShelterCreateRequest) (int64, error) {
	return mysql.CreateAnimalShelter(ctx, shelter)
}

func ShelterUpdate(ctx context.Context, id int64, shelter *entity.ShelterCreateRequest) error {
	return mysql.UpdateShelter(ctx, id, shelter)
}

func ShelterDelete(ctx context.Context, id int64) error {
	return mysql.RemoveShelterByID(ctx, id)
}
