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

func GetShelterStatus(ctx context.Context, shelterID int64) (bool, error) {
	res := mioDB.QueryRowContext(ctx, `
SELECT isConfirmed
FROM usersShelters
WHERE shelterID = ?`, shelterID)
	var isConfirmed bool
	err := res.Scan(&isConfirmed)
	if err != nil {
		return false, err
	}
	return isConfirmed, nil
}

func GetAllSheltersInfo(ctx context.Context) ([]entity.Shelter, error) {
	rows, err := mioDB.QueryContext(ctx, `
SELECT SH.id, SH.shelter_name
  FROM animal_shelters AS SH
`)
	if err != nil {
		return nil, err
	}
	var shelters []entity.Shelter
	for rows.Next() {
		var info entity.Shelter
		err := rows.Scan(
			&info.ID, &info.Name,
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

func GetShelterByVolunteerID(ctx context.Context, volunteerID int64) (int64, error) {
	row := mioDB.QueryRowContext(ctx, `
	SELECT VOS.shelterID
	FROM volunteers_on_shelters VOS
	WHERE VOS.volunteerID = ?
`, volunteerID)
	var id int64
	err := row.Scan(&id)
	return id, err
}

func GetShelterParticipatorsCount(ctx context.Context, shelterID int64) (int64, error) {
	rows, err := mioDB.Query(`SELECT count(*) FROM volunteers_on_shelters WHERE shelterID = ?`, shelterID)
	if err != nil {
		return 0, err
	}
	var count int64
	defer rows.Close()
	for rows.Next() {
		err = rows.Scan(&count)
	}
	return count, err
}

func GetSheltersParticipators(ctx context.Context, shID int64) ([]entity.User, error) {
	rows, err := mioDB.QueryContext(ctx, `
SELECT v.id, v.firstName, v.lastName, v.role, v.email, v.phone
        FROM volunteers v
        JOIN volunteers_on_shelters vos ON v.id = vos.volunteerID
        WHERE vos.shelterID = ?
`, shID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var users []entity.User
	for rows.Next() {
		var user entity.User
		err = rows.Scan(&user.ID, &user.FirstName, &user.LastName, &user.Role, &user.Email, &user.Phone)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return users, nil
}

func RemoveUserFromShelter(ctx context.Context, userID int64) error {
	_, err := mioDB.ExecContext(ctx, `
DELETE FROM volunteers_on_shelters
WHERE volunteerID = ?
`, userID)
	if err != nil {
		return err
	}
	return nil
}

func ConfirmAnimalShelter(ctx context.Context, shelterID int64) (int64, error) {
	var requestedBy int64
	err := mioDB.QueryRowContext(ctx, `
		SELECT requestedBy
		FROM usersShelters
		WHERE shelterID = ?
	`, shelterID).Scan(&requestedBy)
	if err != nil {
		return 0, err
	}

	_, err = mioDB.ExecContext(ctx, `
UPDATE usersShelters
   SET isConfirmed = TRUE
 WHERE shelterID = ? 
`, shelterID)
	if err != nil {
		return 0, err
	}
	return requestedBy, err
}

func RejectAnimalShelter(ctx context.Context, shelterID int64) error {
	_, err := mioDB.ExecContext(ctx, `
DELETE FROM usersShelters
	WHERE shelterID = ?
`, shelterID)
	if err != nil {
		return err
	}
	return nil
}

func CreateShelterRequest(ctx context.Context, shID int64, userID int64) error {
	_, err := mioDB.ExecContext(ctx, `
INSERT INTO usersShelters
		SET requestedBy = ?,
			shelterID = ?,
			isConfirmed = FALSE,
			createdAt = UNIX_TIMESTAMP()`, userID, shID)
	return err
}

func CreateShelterInvitation(ctx context.Context, shID, sendID, recID int64) error {
	_, err := mioDB.ExecContext(ctx, `
INSERT INTO volunteersInvitationsOnShelters
	SET senderID = ?,
		recipientID = ?,
		shelterID = ?
`, sendID, recID, shID)
	if err != nil {
		return err
	}
	return nil
}
