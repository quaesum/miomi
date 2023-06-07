package mysql

import (
	"context"
	"fmt"
	"madmax/internal/entity"
)

func GetAnimalBasicInfo(ctx context.Context, animalID int64) (*entity.Animal, error) {
	row := mioDB.QueryRowContext(ctx, `
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
  A.onhappines
FROM 
  animals AS A 
  INNER JOIN animal_types AS ANT 
  LEFT JOIN animals_on_types AS AOT ON A.id = AOT.animalID 
  AND ANT.id = AOT.animal_typeID 
  INNER JOIN animal_shelters AS SH 
  LEFT JOIN animals_on_shelters AS ASH ON A.id = ASH.animalID 
  AND SH.id = ASH.shelterID 
WHERE 
  A.id = ?
  AND A.id = AOT.animalID 
  AND A.id = ASH.animalID 
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
  A.onhappines`, animalID)
	animal := new(entity.Animal)
	err := row.Scan(
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
  SH.adress,
  SH.phone,
  SH.id,
  A.onrainbow, 
  A.onhappines
FROM 
  animals AS A 
  INNER JOIN animal_types AS ANT 
  LEFT JOIN animals_on_types AS AOT ON A.id = AOT.animalID 
  AND ANT.id = AOT.animal_typeID 
  INNER JOIN animal_shelters AS SH 
  LEFT JOIN animals_on_shelters AS ASH ON A.id = ASH.animalID 
  AND SH.id = ASH.shelterID 
WHERE 
  A.id = AOT.animalID 
  AND A.id = ASH.animalID 
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
  SH.adress,
  SH.phone,
  SH.id,
  A.onrainbow, 
  A.onhappines
`)

	if err != nil {
		return nil, err
	}
	var animals []entity.Animal
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
			&animal.Address,
			&animal.Phone,
			&animal.ShelterId,
			&animal.OnRainbow,
			&animal.OnHappiness,
		)
		if err != nil {
			fmt.Println(err)
			return nil, err
		}
		animals = append(animals, animal)
	}

	return animals, nil
}

func CreateAnimal(ctx context.Context, animal *entity.AnimalCreateRequest) (int64, error) {
	res, err := mioDB.ExecContext(ctx, `
INSERT INTO animals  
		SET  age = ?,
		  	name = ?,
 			sex = ?,
   			description = ?,
            sterilized = ?,
            vaccinated = ?
`, animal.Age, animal.Name, animal.Sex, animal.Description, animal.Sterilized, animal.Vaccinated)
	if err != nil {
		return 0, err
	}
	return res.LastInsertId()
}
func AddAnimalOnType(ctx context.Context, typeID, animalID int64) error {
	_, err := mioDB.ExecContext(ctx, `
INSERT INTO animals_on_types
(animal_typeID, animalID)
VALUES(?, ?);
`, typeID, animalID)
	if err != nil {
		return err
	}
	return nil
}
