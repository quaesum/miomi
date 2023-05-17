package application

import (
	"context"
	"madmax/internal/application/db/mysql"
	"madmax/internal/entity"
)

func userCreation() error {
	//utils.GetMD5Hash()
	//mysql.CreateUser()
	return nil
}

func UserByID(ctx context.Context, id int64) (*entity.User, error) {
	return mysql.UpdateUserByID(ctx, id)
}

func AnimalsAll(ctx context.Context) ([]entity.Animal, error) {
	return mysql.GetAllAnimals(ctx)
}
