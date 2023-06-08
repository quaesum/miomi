package application

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"madmax/internal/application/db/mysql"
	"madmax/internal/entity"
)

func AnimalCreate(ctx context.Context, userID int64, animalData *entity.AnimalCreateRequest) (int64, error) {
	_, err := mysql.GetUserByID(ctx, userID)
	if err != nil && err != sql.ErrNoRows {
		return 0, errors.New("user exist")
	}
	animalID, err := mysql.CreateAnimal(ctx, animalData)
	if err != nil {
		return 0, err
	}
	err = mysql.AddAnimalOnType(ctx, animalData.Type, animalID)
	if err != nil {
		return 0, err
	}
	fmt.Println(animalData.ShelterId, animalID)
	err = mysql.AddAnimalOnShelter(ctx, animalData.ShelterId, animalID)
	if err != nil {
		fmt.Println(err)
		return 0, err
	}
	for _, photoID := range animalData.Photos {
		err = mysql.AddanimalsPhotos(ctx, animalID, photoID)
		if err != nil {
			return 0, err
		}
	}

	return animalID, nil

}

func AnimalByID(ctx context.Context, id int64) (*entity.Animal, error) {
	animal, err := mysql.GetAnimalBasicInfo(ctx, id)
	if err != nil && err == sql.ErrNoRows {
		fmt.Println(err)
		return nil, errors.New("animal not exist")
	}
	photos, err := mysql.GetPhotosByAnimalID(ctx, id)
	if err != nil && err != sql.ErrNoRows {
		return nil, err
	}
	animal.Photos = photos
	return animal, nil
}

func RemoveAnimalByID(ctx context.Context, id int64) error {
	err := mysql.RemoveAnimalByID(ctx, id)
	if err != nil && err != sql.ErrNoRows {
		return err
	}
	return nil
}

func AnimalUpdate(ctx context.Context, userID, animalID int64, animalData *entity.AnimalCreateRequest) error {
	_, err := mysql.GetUserByID(ctx, userID)
	if err != nil && err != sql.ErrNoRows {
		return errors.New("user exist")
	}
	err = mysql.UpdateAnimal(ctx, animalID, animalData)
	if err != nil {
		return err
	}

	err = mysql.RemoveAnimalOnShelter(ctx, animalID)
	if err != nil {
		return err
	}
	fmt.Println(animalData.ShelterId, animalID)
	err = mysql.AddAnimalOnShelter(ctx, animalData.ShelterId, animalID)
	if err != nil {
		fmt.Println(err)
		return err
	}
	err = mysql.RemoveAnimalsPhotos(ctx, animalID)
	if err != nil {
		return err
	}
	for _, photoID := range animalData.Photos {
		err = mysql.AddanimalsPhotos(ctx, animalID, photoID)
		if err != nil {
			return err
		}
	}
	return nil
}
