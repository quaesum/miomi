package internal

import (
	"context"
	"madmax/internal/entity"
	"madmax/internal/mysql"
)

func userCreation() error {
	//utils.GetMD5Hash()
	//mysql.CreateUser()
	return nil
}

func userByID(ctx context.Context, id int64) (*entity.User, error) {
	return mysql.UpdateUserByID(ctx, id)
}
