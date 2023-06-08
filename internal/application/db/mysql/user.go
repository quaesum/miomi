package mysql

import (
	"context"
	"madmax/internal/entity"
	"time"
)

func CreateUser(ctx context.Context, info *entity.User) (int64, error) {
	res, err := mioDB.ExecContext(ctx, `
INSERT INTO volunteers
		SET firstName = ?,
			lastName = ?,
			password  = ?,
			email  = ?,
			user_role  = ?,
			createdAt = UNIX_TIMESTAMP()
`, info.FirstName, info.LastName, info.Password, info.Email, info.Role)
	if err != nil {
		return 0, err
	}
	return res.LastInsertId()
}

func UpdateUser(ctx context.Context, info *entity.User) error {
	_, err := mioDB.ExecContext(ctx, `
UPDATE volunteers
   SET firstName = ?,
	   lastName = ?,
       email  = ?                
 WHERE id = ? 
`, info.FirstName, info.LastName, info.Email, info.ID)
	if err != nil {
		return err
	}
	return nil
}

func GetUserByID(ctx context.Context, userID int64) (*entity.User, error) {
	row := mioDB.QueryRowContext(ctx, `
SELECT U.id, U.firstName, U.lastName, U.password, U.email,  U.createdAt, U.user_role, ASH.id
  FROM volunteers AS U
  INNER JOIN animal_shelters as ASH
  LEFT JOIN volunteers_on_shelters AS VOSH ON U.id = VOSH.volunteerID
  AND ASH.id = VOSH.shelterID
 WHERE U.id = ?`, userID)
	info := new(entity.User)
	var createdAt int64
	err := row.Scan(
		&info.ID, &info.FirstName, &info.LastName, &info.Password,
		&info.Email, &createdAt, &info.Role, &info.ShelterID,
	)
	info.CreatedAt = time.Unix(createdAt, 0).Format(time.RFC3339)
	return info, err
}

func GetUserByEmail(ctx context.Context, email string) (*entity.User, error) {
	row := mioDB.QueryRowContext(ctx, `
SELECT U.id, U.firstName, U.lastName, U.password, U.email,  U.createdAt, U.user_role
  FROM volunteers AS U
 WHERE U.email = ?`, email)
	info := new(entity.User)
	var createdAt int64
	err := row.Scan(
		&info.ID, &info.FirstName, &info.LastName, &info.Password,
		&info.Email, &createdAt, &info.Role,
	)
	info.CreatedAt = time.Unix(createdAt, 0).Format(time.RFC3339)
	if err != nil {
		return nil, err
	}
	return info, err
}

func GetUserBasicInfo(ctx context.Context, userID int64) (*entity.User, error) {
	row := mioDB.QueryRowContext(ctx, `
SELECT U.id, U.name, U.email
  FROM volunteers AS U
 WHERE U.id = ? `, userID)
	user := new(entity.User)
	err := row.Scan(
		&user.ID, &user.LastName, &user.Email,
	)
	return user, err
}
