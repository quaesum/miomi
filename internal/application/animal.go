package application

import (
	"context"
	"fmt"
	"github.com/samber/lo"
	anBleve "madmax/internal/application/db/bleve"
	"madmax/internal/application/db/mysql"
	"madmax/internal/entity"
	"madmax/internal/utils"
	"sort"
	"strconv"
	"strings"
)

type AnimalApplication struct {
	Bl *anBleve.AnimalBleve
}

func NewAnimalApplication() *AnimalApplication {
	return &AnimalApplication{
		Bl: anBleve.NewAnimal(),
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
	err = a.Bl.Add(strconv.Itoa(int(animalID)), animalBleve)
	if err != nil {
		return 0, err
	}
	return animalID, nil

}

func (a *AnimalApplication) CreateDirect(ctx context.Context, animalData *entity.Animal) error {
	var req = &entity.AnimalCreateRequest{
		Age:         animalData.Age,
		AgeType:     animalData.AgeType,
		Name:        animalData.Name,
		Sex:         animalData.Sex,
		Description: animalData.Description,
		Sterilized:  animalData.Sterilized,
		Vaccinated:  animalData.Vaccinated,
		Onhappines:  animalData.OnHappiness,
		Onrainbow:   animalData.OnRainbow,
	}
	animalID, err := mysql.CreateAnimal(ctx, req)
	if err != nil {
		return err
	}
	anType := func() int64 {
		switch animalData.Type {
		case "cat":
			return 1
		case "dog":
			return 2
		case "other":
			return 4
		}
		return 4
	}
	err = mysql.AddAnimalOnType(ctx, anType(), animalID)
	if err != nil {
		return err
	}
	err = mysql.AddAnimalOnShelter(ctx, animalData.ShelterId, animalID)
	if err != nil {
		return err
	}
	for _, photo := range animalData.Photos {
		id, err := mysql.CreateFile(ctx, photo)
		if err != nil {
			return err
		}
		err = mysql.AddAnimalsPhotos(ctx, animalID, id)
		if err != nil {
			return err
		}
	}

	animal, err := a.GetByID(ctx, animalID)
	if err != nil {
		return err
	}
	animalBleve := entity.InserAnimalReqToCreate(*animal)
	err = a.Bl.Add(strconv.Itoa(int(animalID)), animalBleve)
	if err != nil {
		return err
	}

	return nil
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

func (a *AnimalApplication) GetByShelterID(ctx context.Context, shID int64) ([]entity.Animal, error) {
	animals, err := mysql.GetAnimalsByShelterID(ctx, shID)
	if err != nil {
		return nil, err
	}
	return animals, nil
}

func (a *AnimalApplication) Remove(ctx context.Context, id int64) error {
	err := mysql.RemoveAnimalByID(ctx, id)
	if err != nil {
		return err
	}
	err = a.Bl.Remove(strconv.Itoa(int(id)))
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
	err = a.Bl.Add(strconv.Itoa(int(animalID)), animalBleve)
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

func (a *AnimalApplication) GetFromBleve(req *entity.SearchRequest, limit int) ([]entity.AnimalBleve, error) {
	animalBleve := anBleve.NewAnimal()
	res, err := animalBleve.Search(req, limit)
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
			Age:         safeFloat64(result["age"]),
			AgeType:     safeStringPtr(result["ageType"]),
			Name:        safeString(result["name"]),
			Sex:         safeString(result["sex"]),
			Type:        safeString(result["type"]),
			Description: safeString(result["description"]),
			Sterilized:  safeBool(result["sterilized"]),
			Vaccinated:  safeBool(result["vaccinated"]),
			ShelterId:   safeString(result["shelterId"]),
			Shelter:     safeString(result["shelter"]),
			Address:     safeString(result["address"]),
		}
		animal.Photos, err = utils.ProcessPhotos(result["photos"])
		if err != nil {
			return nil, err
		}
		animals = append(animals, animal)
	}
	return animals, err
}

func GetAnimalTypes(ctx context.Context) ([]entity.AnimalTypes, error) {
	return mysql.GetAllAnimalTypes(ctx)
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

func safeString(i interface{}) string {
	if i == nil {
		return ""
	}
	if v, ok := i.(string); ok {
		return v
	}
	return ""
}

// safeStringPtr получает указатель на строку из интерфейса
func safeStringPtr(i interface{}) *string {
	if i == nil {
		return nil
	}
	if v, ok := i.(string); ok {
		return &v
	}
	return nil
}

// safeFloat64 получает float64 значение из интерфейса
func safeFloat64(i interface{}) float64 {
	if i == nil {
		return 0
	}
	if v, ok := i.(float64); ok {
		return v
	}
	return 0
}

// safeBool получает bool значение из интерфейса
func safeBool(i interface{}) bool {
	if i == nil {
		return false
	}
	if v, ok := i.(bool); ok {
		return v
	}
	return false
}

// safeStringSlice получает срез строк из интерфейса
func safeStringSlice(i interface{}) []string {
	if i == nil {
		return nil
	}
	if v, ok := i.([]string); ok {
		return v
	}
	return nil
}
