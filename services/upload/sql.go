package main

import (
	"context"
	"database/sql"
	"madmax/internal/entity"
)

func CreateFile(db *sql.DB, ctx context.Context, name string) (int64, error) {
	res, err := db.ExecContext(ctx, `
INSERT INTO photos
(filename, origin, file_type)
VALUES(?, '', '');
`, name)
	if err != nil {
		return 0, err
	}
	return res.LastInsertId()
}

func CreateDirectAnimal(db *sql.DB, ctx context.Context, animal *entity.Animal) (int64, error) {
	res, err := db.ExecContext(ctx, `
INSERT INTO animals  
		SET  age = ?,
		    ageType = ?,
		  	name = ?,
 			sex = ?,
   			description = ?,
            sterilized = ?,
            vaccinated = ?,
		    onrainbow = false,
            onhappines  = false
`, animal.Age, animal.AgeType, animal.Name, animal.Sex, animal.Description, animal.Sterilized, animal.Vaccinated)
	if err != nil {
		return 0, err
	}
	return res.LastInsertId()
}

func AddAnimalOnType(db *sql.DB, ctx context.Context, typeID, animalID int64) error {
	_, err := db.ExecContext(ctx, `
INSERT INTO animals_on_types
(animal_typeID, animalID)
VALUES(?, ?);
`, typeID, animalID)
	if err != nil {
		return err
	}
	return nil
}

func AddAnimalOnShelter(db *sql.DB, ctx context.Context, shelterID, animalID int64) error {
	_, err := db.ExecContext(ctx, `
INSERT INTO animals_on_shelters
(animalID, shelterID)
VALUES(?, ?);
`, animalID, shelterID)
	return err
}

func AddAnimalsPhotos(db *sql.DB, ctx context.Context, animalID, photoID int64) error {
	_, err := db.ExecContext(ctx, `
INSERT INTO animals_photos
(animalID, photoID)
VALUES(?, ?);
`, animalID, photoID)
	return err
}

func RemoveAnimalByID(db *sql.DB, ctx context.Context, animalID int64) error {
	_, err := db.ExecContext(ctx, `
DELETE FROM animals
WHERE id = ?

`, animalID)
	if err != nil {
		return err
	}
	return nil
}
