package application

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"madmax/internal/application/db/mysql"
	"madmax/internal/entity"
)

func UserCreate(ctx context.Context, userData *entity.UserCreateRequest) (string, error) {
	_, err := mysql.GetUserByEmail(ctx, userData.Email)
	if err != nil && err == sql.ErrNoRows {
		return "", errors.New("user exist")
	}
	fmt.Println(userData)
	if userData.ShelterID == 0 {
		sc := entity.ShelterCreateRequest{
			Name:    userData.FirstName + " " + userData.LastName,
			Address: userData.Address,
			Phone:   userData.Phone,
			Email:   userData.Email,
		}
		shID, err := mysql.CreateAnimalShelter(ctx, &sc)
		if err != nil {
			return "", err
		}
		userData.ShelterID = shID
	}
	u := entity.User{
		FirstName: userData.FirstName,
		LastName:  userData.LastName,
		Password:  userData.Password,
		Email:     userData.Email,
	}
	userID, err := mysql.CreateUser(ctx, &u)
	if err != nil {
		return "", err
	}

	err = mysql.AddUserOnShelter(ctx, userID, userData.ShelterID)
	if err != nil {
		return "", err
	}

	return "", nil
}

func UserByID(ctx context.Context, id int64) (*entity.User, error) {
	return mysql.GetUserByID(ctx, id)
}
