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
