package main

import (
	"context"
	"database/sql"
	"encoding/json"
	"github.com/blevesearch/bleve"
	"io"
	"log"
	"madmax/internal"
	bleve2 "madmax/internal/application/db/bleve"
	"madmax/internal/application/db/mysql"
	"madmax/internal/entity"
	"madmax/services/telegram-scrapper/scrapper"
	"os"
	"strconv"
)

func main() {
	config, err := internal.NewConfig()
	if err != nil {
		log.Fatal(err)
	}

	mioDB, err := mysql.NewDB(config.MysqlDSN)
	ind, err := bleve2.NewBleveAnimals()
	if err != nil {
		log.Fatal(err)
	}

	err = RemoveOldAnimals(ind, mioDB)
	if err != nil {
		log.Fatal(err)
	}
	animals, err := getFilteredAnimals()
	if err != nil {
		log.Fatal(err)
	}
	ctx := context.Background()

	var records []AnimalRecord

	for _, animal := range animals {
		animal.ShelterId = 1
		anID, err := AddAnimal(ind, mioDB, ctx, &animal)
		if err != nil {
			log.Fatal(err)
		}
		records = append(records, AnimalRecord{
			AnimalID: anID,
			RecordID: animal.ID,
		})
	}
	recs, err := json.Marshal(records)
	if err != nil {
		log.Fatal(err)
	}

	err = os.WriteFile("animals_upload.json", recs, 0666)
	if err != nil {
		log.Fatal(err)
	}
}

func getFilteredAnimals() ([]entity.Animal, error) {
	file, err := os.Open("animals_gemini.json")
	if err != nil {
		return nil, err
	}
	defer file.Close()

	data, err := io.ReadAll(file)
	if err != nil {
		return nil, err
	}

	var animals []entity.Animal
	err = json.Unmarshal(data, &animals)
	if err != nil {
		return nil, err
	}

	// Чтение animals.json
	fileShort, err := os.Open("animals.json")
	if err != nil {
		return nil, err
	}
	defer fileShort.Close()

	dataShort, err := io.ReadAll(fileShort)
	if err != nil {
		return nil, err
	}

	var animalsShort []scrapper.Animal
	err = json.Unmarshal(dataShort, &animalsShort)
	if err != nil {
		return nil, err
	}

	photosMap := make(map[int][]string)
	for _, animal := range animalsShort {
		photosMap[animal.ID] = animal.Photos
	}

	for i, animal := range animals {
		if photos, exists := photosMap[int(animal.ID)]; exists {
			animals[i].Photos = photos
		} else {
			remove(animals, i)
		}
	}

	// Фильтрация объектов
	var filteredAnimals []entity.Animal
	for _, animal := range animals {
		if (animal.Type == "other" || animal.Type == "") || animal.AgeType == "" || animal.Age == 0 || animal.Description == "" || animal.Photos == nil {
			continue
		}
		filteredAnimals = append(filteredAnimals, animal)
	}

	// Запись обновленных данных обратно в animals_gemini.json
	animalsNew, err := json.Marshal(filteredAnimals)
	if err != nil {
		return nil, err
	}

	err = os.WriteFile("animals_gemini.json", animalsNew, 0666)
	if err != nil {
		return nil, err
	}
	return filteredAnimals, nil
}

func remove(slice []entity.Animal, s int) []entity.Animal {
	return append(slice[:s], slice[s+1:]...)
}

func AddAnimal(ind bleve.Index, db *sql.DB, ctx context.Context, animal *entity.Animal) (int64, error) {
	anID, err := CreateDirectAnimal(db, ctx, animal)
	if err != nil {
		return 0, err
	}
	var anType int64
	var bleveAnType string
	switch animal.Type {
	case "cat":
		anType = 1
		bleveAnType = "Кот"
	case "dog":
		anType = 2
		bleveAnType = "Собака"
	case "other":
		anType = 4
		bleveAnType = "Другое"
	default:
		anType = 4
		bleveAnType = "Другое"
	}

	err = AddAnimalOnType(db, ctx, anType, anID)
	if err != nil {
		return 0, err
	}

	err = AddAnimalOnShelter(db, ctx, animal.ShelterId, anID)
	if err != nil {
		return 0, nil
	}

	for _, photo := range animal.Photos {
		id, err := mysql.CreateFile(ctx, photo)
		if err != nil {
			return 0, err
		}
		err = mysql.AddAnimalsPhotos(ctx, anID, id)
		if err != nil {
			return 0, err
		}
	}

	animalBleve := entity.InserAnimalReqToCreate(*animal)
	animalBleve.Type = bleveAnType
	err = Add(ind, strconv.Itoa(int(anID)), animalBleve)
	if err != nil {
		return 0, err
	}

	return anID, nil
}

func Add(ind bleve.Index, animalID string, animal *entity.AnimalCreateBleve) error {
	err := ind.Index(animalID, &animal)
	if err != nil {
		return err
	}
	return nil
}

func Remove(ind bleve.Index, id string) error {
	err := ind.Delete(id)
	if err != nil {
		return err
	}
	return nil
}

func RemoveOldAnimals(ind bleve.Index, db *sql.DB) error {
	file, err := os.Open("animals_upload.json")
	if err != nil {
		return err
	}
	defer file.Close()

	data, err := io.ReadAll(file)
	if err != nil {
		return err
	}

	var animals []AnimalRecord
	err = json.Unmarshal(data, &animals)
	if err != nil {
		return err
	}

	ctx := context.Background()

	for _, animal := range animals {
		err = RemoveAnimalByID(db, ctx, animal.AnimalID)
		if err != nil {
			log.Println("animal wasn't deleted from mysql, id:", animal.AnimalID)
		}
		err = Remove(ind, strconv.Itoa(int(animal.AnimalID)))
		if err != nil {
			log.Println("animal wasn't deleted from bleve, id:", animal.AnimalID)
		}
	}

	return nil
}
