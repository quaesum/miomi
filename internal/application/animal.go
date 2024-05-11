package application

import (
	"context"
	"fmt"
	"github.com/samber/lo"
	"madmax/internal/application/db/bleve"
	"madmax/internal/application/db/mysql"
	"madmax/internal/entity"
	"madmax/internal/utils"
	"sort"
	"strconv"
	"strings"
)

type AnimalApplication struct {
	bleve *bleve.AnimalBleve
}

func NewAnimalApplication() *AnimalApplication {
	return &AnimalApplication{
		bleve: bleve.NewAnimal(),
	}
}

func (a *AnimalApplication) Create(ctx context.Context, animalData *entity.AnimalCreateRequest) (int64, error) {
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

	animal, err := a.GetByID(ctx, animalID)
	if err != nil {
		return 0, err
	}
	animalBleve := entity.InserAnimalReqToCreate(*animal)
	err = a.bleve.Add(strconv.Itoa(int(animalID)), animalBleve)
	if err != nil {
		return 0, err
	}
	return animalID, nil

}

func (a *AnimalApplication) GetByID(ctx context.Context, id int64) (*entity.Animal, error) {
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

func (a *AnimalApplication) Remove(ctx context.Context, id int64) error {
	err := mysql.RemoveAnimalByID(ctx, id)
	if err != nil {
		return err
	}
	err = a.bleve.Remove(strconv.Itoa(int(id)))
	if err != nil {
		return err
	}

	return nil
}

func (a *AnimalApplication) Update(ctx context.Context, animalID int64, animalData *entity.AnimalCreateRequest) error {
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

	animal, err := a.GetByID(ctx, animalID)
	if err != nil {
		return err
	}
	animalBleve := entity.InserAnimalReqToCreate(*animal)
	err = a.bleve.Add(strconv.Itoa(int(animalID)), animalBleve)
	if err != nil {
		return err
	}

	return nil
}

func (a *AnimalApplication) GetAllFromMYSQL(ctx context.Context) ([]entity.Animal, error) {
	animals, err := mysql.GetAllAnimals(ctx)
	if err != nil {
		return nil, err
	}

	return animals, err
}

func (a *AnimalApplication) GetFromBleve(searchQuery string) ([]entity.AnimalBleve, error) {
	animalBleve := bleve.NewAnimal()
	res, err := animalBleve.Search(searchQuery)
	if err != nil {
		return nil, err
	}
	var animals []entity.AnimalBleve
	for _, item := range res.Hits {
		result := item.Fields
		id, err := strconv.ParseInt(item.ID, 10, 64)
		if err != nil {
			return nil, err
		}
		animal := entity.AnimalBleve{
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
		animal.Photos, err = utils.ProcessPhotos(result["photos"])
		if err != nil {
			return nil, err
		}
		animals = append(animals, animal)
	}
	return animals, err
}

func (a *AnimalApplication) GetAllFromBleve() ([]entity.AnimalBleve, error) {
	animalBleve := bleve.NewAnimal()
	res, err := animalBleve.SearchWOQuery()
	if err != nil {
		return nil, err
	}
	var animals []entity.AnimalBleve
	for _, item := range res.Hits {
		result := item.Fields
		id, err := strconv.ParseInt(item.ID, 10, 64)
		if err != nil {
			return nil, err
		}
		animal := entity.AnimalBleve{
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
		animal.Photos, err = utils.ProcessPhotos(result["photos"])
		if err != nil {
			return nil, err
		}
		animals = append(animals, animal)
	}
	return animals, err
}

func GetAnimalsSearchResult(searchTerm string, animals []entity.Animal) ([]entity.Animal, error) {
	searchTerm = utils.CleanQuery(searchTerm)

	for i := range animals {
		animals[i].Score = calculateAnimalsScore(animals[i], searchTerm)
	}

	sort.Slice(animals, func(i, j int) bool {
		return animals[i].Score > animals[j].Score
	})

	animals = lo.Filter(animals, func(animal entity.Animal, _ int) bool {
		return animal.Score > 0
	})

	return animals, nil
}

func calculateAnimalsScore(animal entity.Animal, searchTerm string) int64 {
	searchTerm = strings.ToLower(searchTerm)
	nameCount := strings.Count(strings.ToLower(animal.Name), searchTerm)
	descriptionCount := strings.Count(strings.ToLower(animal.Description), searchTerm)
	return int64(nameCount + descriptionCount)
}
