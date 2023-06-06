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
