package mysql

import (
	"context"
	"madmax/internal/entity"
)

func GetUserBasicInfo(ctx context.Context, userID int64) (*entity.User, error) {
	row := mioDB.QueryRowContext(ctx, `
SELECT U.id, U.name, U.email
  FROM users AS U
 WHERE U.id = ? `, userID)
	user := new(entity.User)
	err := row.Scan(
		&user.ID, &user.Name, &user.Email,
	)
	return user, err
}

func GetAnimalBasicInfo(ctx context.Context, animalID int64) (*entity.Animal, error) {
	row := mioDB.QueryRowContext(ctx, `
SELECT A.id, A.age, A.name, A.sex, A.type, A.description, A.castrated, A.sterilized, A.vaccinated, A.shelter
	FROM animals as A
WHERE A.id = ?`, animalID)
	animal := new(entity.Animal)
	err := row.Scan(
		&animal.ID, &animal.Age, &animal.Name, &animal.Sex, &animal.Type, &animal.Description, &animal.Castrated, &animal.Sterilized, &animal.Vaccinated, &animal.Shelter,
	)
	return animal, err
}
