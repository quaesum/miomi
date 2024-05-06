package application

import (
	"context"
	"errors"
	"fmt"
	"log"
	"madmax/internal/application/db/bleve"
	"madmax/internal/application/db/mysql"
	"madmax/internal/entity"
	"strconv"
)

func AnimalCreate(ctx context.Context, animalData *entity.AnimalCreateRequest) (int64, error) {
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
		err = mysql.AddAnimalsPhotos(ctx, animalID, photoID)
		if err != nil {
			return 0, err
		}
	}

	animal, err := AnimalByID(ctx, animalID)
	if err != nil {
		return 0, err
	}
	animalBleve := entity.InserAnimalReqToCreate(*animal)
	err = bleve.AddAnimal(strconv.Itoa(int(animalID)), animalBleve)
	if err != nil {
		return 0, err
	}
	return animalID, nil

}

func AnimalByID(ctx context.Context, id int64) (*entity.Animal, error) {
	animal, err := mysql.GetAnimalBasicInfo(ctx, id)
	if err != nil {
		fmt.Println(err)
	}
	photos, err := mysql.GetPhotosByAnimalID(ctx, id)
	if err != nil {
		return nil, err
	}
	animal.Photos = photos
	return animal, nil
}

func RemoveAnimalByID(ctx context.Context, id int64) error {
	err := mysql.RemoveAnimalByID(ctx, id)
	if err != nil {
		return err
	}

	err = bleve.DeleteAnimal(strconv.Itoa(int(id)))
	if err != nil {
		return err
	}

	return nil
}

func AnimalUpdate(ctx context.Context, animalID int64, animalData *entity.AnimalCreateRequest) error {
	err := mysql.UpdateAnimal(ctx, animalID, animalData)
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
		err = mysql.AddAnimalsPhotos(ctx, animalID, photoID)
		if err != nil {
			return err
		}
	}

	animal, err := AnimalByID(ctx, animalID)
	if err != nil {
		return err
	}
	animalBleve := entity.InserAnimalReqToCreate(*animal)
	err = bleve.AddAnimal(strconv.Itoa(int(animalID)), animalBleve)
	if err != nil {
		return err
	}

	return nil
}

func GetAllAnimalsFromMysql(ctx context.Context) ([]entity.Animal, error) {
	animals, err := mysql.GetAllAnimals(ctx)
	if err != nil {
		return nil, err
	}

	return animals, err
}

func GetAnimalsFromBleve(searchQuery string) ([]entity.AnimalsBleve, error) {
	res, err := bleve.SearchAnimal(searchQuery)
	if err != nil {
		return nil, err
	}
	var animals []entity.AnimalsBleve
	for _, item := range res.Hits {
		result := item.Fields
		id, err := strconv.ParseInt(item.ID, 10, 64)
		if err != nil {
			return nil, err
		}
		animal := entity.AnimalsBleve{
			ID:          id,
			Age:         result["age"].(float64),
			Name:        result["name"].(string),
			Sex:         result["sex"].(float64),
			Type:        result["type"].(string),
			Description: result["description"].(string),
			Sterilized:  result["sterilized"].(bool),
			Vaccinated:  result["vaccinated"].(bool),
			ShelterId:   result["shelterId"].(float64),
		}
		err = processPhotos(result["photos"], &animal)
		if err != nil {
			return nil, err
		}
		animals = append(animals, animal)
	}
	return animals, err
}

func GetAllAnimalsFromBleve() ([]entity.AnimalsBleve, error) {
	res, err := bleve.SearchWOQuery()
	log.Println(res.Hits)
	if err != nil {
		return nil, err
	}
	var animals []entity.AnimalsBleve
	for _, item := range res.Hits {
		result := item.Fields
		log.Println(result)
		id, err := strconv.ParseInt(item.ID, 10, 64)
		if err != nil {
			return nil, err
		}
		animal := entity.AnimalsBleve{
			ID:          id,
			Age:         result["age"].(float64),
			Name:        result["name"].(string),
			Sex:         result["sex"].(float64),
			Type:        result["type"].(string),
			Description: result["description"].(string),
			Sterilized:  result["sterilized"].(bool),
			Vaccinated:  result["vaccinated"].(bool),
			ShelterId:   result["shelterId"].(float64),
			Shelter:     result["shelter"].(string),
			Address:     result["address"].(string),
		}
		err = processPhotos(result["photos"], &animal)
		if err != nil {
			return nil, err
		}
		animals = append(animals, animal)
	}
	return animals, err
}

func processPhotos(photos interface{}, animal *entity.AnimalsBleve) error {
	switch v := photos.(type) {
	case string:
		animal.Photos = append(animal.Photos, v)
	case []interface{}:
		for _, photo := range v {
			if str, ok := photo.(string); ok {
				animal.Photos = append(animal.Photos, str)
			}
		}
	default:
		return errors.New("invalid photo type")
	}
	return nil
}
