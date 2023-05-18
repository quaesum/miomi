package mysql

import (
	"context"
	"fmt"
	"madmax/internal/entity"
	"strings"
)

func GetAnimalBasicInfo(ctx context.Context, animalID int64) (*entity.Animal, error) {
	row := mioDB.QueryRowContext(ctx, `
SELECT A.id, A.age, A.name, A.sex, A.type, A.description,  A.sterilized, A.vaccinated, A.shelter
	FROM animals as A
WHERE A.id = ?`, animalID)
	animal := new(entity.Animal)
	err := row.Scan(
		&animal.ID, &animal.Age, &animal.Name, &animal.Sex, &animal.Type, &animal.Description, &animal.Sterilized, &animal.Vaccinated, &animal.Shelter,
	)
	return animal, err
}

func GetAllAnimals(ctx context.Context) ([]entity.Animal, error) {
	rows, err := mioDB.QueryContext(ctx, `
SELECT 
  A.id, 
  A.age, 
  A.name, 
  A.sex, 
  ANT.name, 
  A.description, 
  A.sterilized, 
  A.vaccinated, 
  SH.shelter_name, 
  A.onrainbow, 
  A.onhappines, 
  GROUP_CONCAT(
    DISTINCT P.filename 
    ORDER BY 
      P.id ASC SEPARATOR ';'
  ) AS a_photos 
FROM 
  animals AS A 
  INNER JOIN animal_types AS ANT 
  LEFT JOIN animals_on_types AS AOT ON A.id = AOT.animalID 
  AND ANT.id = AOT.animal_typeID 
  INNER JOIN animal_shelters AS SH 
  LEFT JOIN animals_on_shelters AS ASH ON A.id = ASH.animalID 
  AND SH.id = ASH.shelterID 
  INNER JOIN photos AS P 
  LEFT JOIN animals_photos AS PH ON A.id = PH.animalID 
  AND P.id = PH.photoID 
WHERE 
  A.id = AOT.animalID 
  AND A.id = ASH.animalID 
  AND A.id = PH.animalID 
GROUP BY 
  A.id, 
  A.age, 
  A.name, 
  A.sex, 
  ANT.name, 
  A.description, 
  A.sterilized, 
  A.vaccinated, 
  SH.shelter_name, 
  A.onrainbow, 
  A.onhappines
`)

	if err != nil {
		return nil, err
	}
	var animals []entity.Animal
	var photos string
	for rows.Next() {
		var animal entity.Animal
		err := rows.Scan(
			&animal.ID,
			&animal.Age,
			&animal.Name,
			&animal.Sex,
			&animal.Type,
			&animal.Description,
			&animal.Sterilized,
			&animal.Vaccinated,
			&animal.Shelter,
			&animal.OnRainbow,
			&animal.OnHappiness,
			&photos,
		)
		if err != nil {
			fmt.Println(err)
			return nil, err
		}
		animal.Photos = strings.Split(photos, ";")
		animals = append(animals, animal)
	}

	return animals, nil
}
