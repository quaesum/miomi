package mysql

import (
	"context"
	"madmax/internal/entity"
)

func CreateAnimalShelter(ctx context.Context, shelter *entity.ShelterCreateRequest) (int64, error) {
	res, err := mioDB.ExecContext(ctx, `
INSERT INTO animal_shelters
		SET shelter_name = ?,
			description = ?,
			logo  = ?,
			phone  = ?,
			adress   = ?
`, shelter.Name, shelter.Description, shelter.Logo, shelter.Phone, shelter.Address)
	if err != nil {
		return 0, err
	}
	return res.LastInsertId()
}

func AddUserOnShelter(ctx context.Context, userID, shelterID int64) error {
	_, err := mioDB.ExecContext(ctx, `
INSERT INTO volunteers_on_shelters
		SET volunteerID = ?,
			shelterID = ?
`, userID, shelterID)
	return err
}

func GetShelterByID(ctx context.Context, shelterID int64) (*entity.Shelter, error) {
	row := mioDB.QueryRowContext(ctx, `
SELECT SH.id, SH.shelter_name, SH.description, SH.logo, SH.phone, SH.adress
  FROM animal_shelters AS SH 
 WHERE SH.id = ?`, shelterID)
	info := new(entity.Shelter)
	err := row.Scan(
		&info.ID, &info.Name, &info.Description, &info.Logo, &info.Phone, &info.Address,
	)
	return info, err
}
func GetAllShelters(ctx context.Context) ([]entity.Shelter, error) {
	rows, err := mioDB.QueryContext(ctx, `
SELECT SH.id, SH.shelter_name, SH.description, SH.logo, SH.phone, SH.adress
  FROM animal_shelters AS SH
`)
	if err != nil {
		return nil, err
	}
	var shelters []entity.Shelter
	for rows.Next() {
		var info entity.Shelter
		err := rows.Scan(
			&info.ID, &info.Name, &info.Description, &info.Logo, &info.Phone, &info.Address,
		)
		if err != nil {
			return nil, err
		}
		shelters = append(shelters, info)
	}

	return shelters, nil
}

func AddAnimalOnShelter(ctx context.Context, shelterID, animalID int64) error {
	_, err := mioDB.ExecContext(ctx, `
INSERT INTO animals_on_shelters
(animalID, shelterID)
VALUES(?, ?);
`, animalID, shelterID)
	return err
}

func RemoveAnimalOnShelter(ctx context.Context, animalID int64) error {
	_, err := mioDB.ExecContext(ctx, `
DELETE FROM animals_on_shelters
		WHERE  animalID = ?
`, animalID)
	return err
}

func UpdateShelter(ctx context.Context, shID int64, shelter *entity.ShelterCreateRequest) error {
	_, err := mioDB.ExecContext(ctx, `
UPDATE animal_shelters
   SET shelter_name = ?,
	description = ?,
	logo = ?,
	phone = ?,
	adress = ?
 WHERE id = ? 
`, shelter.Name, shelter.Description, shelter.Logo, shelter.Phone, shelter.Address, shID)
	if err != nil {
		return err
	}
	return nil
}

func RemoveShelterByID(ctx context.Context, shID int64) error {
	_, err := mioDB.ExecContext(ctx, `
DELETE FROM animal_shelters
WHERE id = ?

`, shID)
	if err != nil {
		return err
	}
	return nil
}
