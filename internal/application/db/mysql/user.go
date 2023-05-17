package mysql

import (
	"context"
	"madmax/internal/entity"
)

func CreateUser(ctx context.Context, info *entity.User) (int64, error) {
	res, err := mioDB.ExecContext(ctx, `
INSERT INTO volunteers
		SET firstName = ?,
			lastName = ?,
			password  = ?,
			email  = ?,
			createdAt = UNIX_TIMESTAMP()
`, info.FirstName, info.LastName, info.Password, info.Email)
	if err != nil {
		return 0, err
	}
	return res.LastInsertId()
}

func UpdateUser(ctx context.Context, info *entity.User) (int64, error) {
	res, err := mioDB.ExecContext(ctx, `
INSERT INTO volunteers
		SET firstName = ?,
			lastName = ?,
			password  = ?,
			email  = ?,
			createdAt = UNIX_TIMESTAMP()
`, info.FirstName, info.LastName, info.Password, info.Email)
	if err != nil {
		return 0, err
	}
	return res.LastInsertId()
}

func UpdateUserByID(ctx context.Context, userID int64) (*entity.User, error) {
	row := mioDB.QueryRowContext(ctx, `
SELECT U.id, U.firstName, U.lastName, U.password, U.email,  U.createdAt
  FROM users AS U
 WHERE U.id = ?`, userID)
	info := new(entity.User)
	err := row.Scan(
		&info.ID, &info.LastName, &info.FirstName, &info.Password,
		&info.Email, &info.CreatedAt,
	)
	return info, err
}

func UpdateUserByEmail(ctx context.Context, email string) (*entity.User, error) {
	row := mioDB.QueryRowContext(ctx, `
SELECT U.id, U.firstName, U.lastName, U.password, U.email,  U.createdAt
  FROM users AS U
 WHERE U.email = ?`, email)
	info := new(entity.User)
	err := row.Scan(
		&info.ID, &info.LastName, &info.FirstName, &info.Password,
		&info.Email, &info.CreatedAt,
	)
	return info, err
}

func GetUserBasicInfo(ctx context.Context, userID int64) (*entity.User, error) {
	row := mioDB.QueryRowContext(ctx, `
SELECT U.id, U.name, U.email
  FROM users AS U
 WHERE U.id = ? `, userID)
	user := new(entity.User)
	err := row.Scan(
		&user.ID, &user.LastName, &user.Email,
	)
	return user, err
}

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
	animals := []entity.Animal{
		{
			ID:          1,
			Age:         3,
			Name:        "Шарик",
			Sex:         1,
			Type:        "CAT",
			Description: "Sed ut perspiciatis, unde omnis iste natus error sit voluptatem accusantium doloremque laudantium, totam rem aperiam eaque ipsa, quae ab illo inventore veritatis et quasi architecto beatae vitae dicta sunt, explicabo. Nemo enim ipsam voluptatem, quia voluptas sit, aspernatur aut odit aut fugit, sed quia consequuntur magni dolores eos, qui ratione voluptatem sequi nesciunt, neque porro quisquam est, qui dolorem ipsum, quia dolor sit, amet, consectetur, adipisci velit, sed quia non numquam eius modi tempora incidunt, ut labore et dolore magnam aliquam quaerat voluptatem. Ut enim ad minima veniam, quis nostrum exercitationem ullam corporis suscipit laboriosam, nisi ut aliquid ex ea commodi consequatur?",
			Sterilized:  true,
			Vaccinated:  false,
			Shelter:     "Super Cat",
		},
		{
			ID:          2,
			Age:         3,
			Name:        "Шарик",
			Sex:         1,
			Type:        "CAT",
			Description: "Sed ut perspiciatis, unde omnis iste natus error sit voluptatem accusantium doloremque laudantium, totam rem aperiam eaque ipsa, quae ab illo inventore veritatis et quasi architecto beatae vitae dicta sunt, explicabo. Nemo enim ipsam voluptatem, quia voluptas sit, aspernatur aut odit aut fugit, sed quia consequuntur magni dolores eos, qui ratione voluptatem sequi nesciunt, neque porro quisquam est, qui dolorem ipsum, quia dolor sit, amet, consectetur, adipisci velit, sed quia non numquam eius modi tempora incidunt, ut labore et dolore magnam aliquam quaerat voluptatem. Ut enim ad minima veniam, quis nostrum exercitationem ullam corporis suscipit laboriosam, nisi ut aliquid ex ea commodi consequatur?",
			Sterilized:  true,
			Vaccinated:  false,
			Shelter:     "Super Cat",
		},
		{
			ID:          3,
			Age:         3,
			Name:        "Шарик",
			Sex:         1,
			Type:        "CAT",
			Description: "Sed ut perspiciatis, unde omnis iste natus error sit voluptatem accusantium doloremque laudantium, totam rem aperiam eaque ipsa, quae ab illo inventore veritatis et quasi architecto beatae vitae dicta sunt, explicabo. Nemo enim ipsam voluptatem, quia voluptas sit, aspernatur aut odit aut fugit, sed quia consequuntur magni dolores eos, qui ratione voluptatem sequi nesciunt, neque porro quisquam est, qui dolorem ipsum, quia dolor sit, amet, consectetur, adipisci velit, sed quia non numquam eius modi tempora incidunt, ut labore et dolore magnam aliquam quaerat voluptatem. Ut enim ad minima veniam, quis nostrum exercitationem ullam corporis suscipit laboriosam, nisi ut aliquid ex ea commodi consequatur?",
			Sterilized:  true,
			Vaccinated:  false,
			Shelter:     "Super Cat",
		},
		{
			ID:          4,
			Age:         3,
			Name:        "Шарик",
			Sex:         1,
			Type:        "CAT",
			Description: "Sed ut perspiciatis, unde omnis iste natus error sit voluptatem accusantium doloremque laudantium, totam rem aperiam eaque ipsa, quae ab illo inventore veritatis et quasi architecto beatae vitae dicta sunt, explicabo. Nemo enim ipsam voluptatem, quia voluptas sit, aspernatur aut odit aut fugit, sed quia consequuntur magni dolores eos, qui ratione voluptatem sequi nesciunt, neque porro quisquam est, qui dolorem ipsum, quia dolor sit, amet, consectetur, adipisci velit, sed quia non numquam eius modi tempora incidunt, ut labore et dolore magnam aliquam quaerat voluptatem. Ut enim ad minima veniam, quis nostrum exercitationem ullam corporis suscipit laboriosam, nisi ut aliquid ex ea commodi consequatur?",
			Sterilized:  true,
			Vaccinated:  false,
			Shelter:     "Super Cat",
		},
	}
	return animals, nil
}
