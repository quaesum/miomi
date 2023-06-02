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
