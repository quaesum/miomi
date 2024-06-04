package mysql

import (
	"context"
	"fmt"
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
			phone = ?,
			role  = ?,
			createdAt = UNIX_TIMESTAMP()
`, info.FirstName, info.LastName, info.Password, info.Email, info.Phone, info.Role)
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
       email  = ?,
       phone = ?
 WHERE id = ? 
`, info.FirstName, info.LastName, info.Email, info.Phone, info.ID)
	if err != nil {
		return err
	}
	return nil
}

func DeleteUser(ctx context.Context, userID int64) error {
	_, err := mioDB.ExecContext(ctx, `
DELETE FROM volunteers
WHERE id = ?

`, userID)
	if err != nil {
		return err
	}
	return nil
}

func GetUserByID(ctx context.Context, userID int64) (*entity.User, error) {
	row := mioDB.QueryRowContext(ctx, `
SELECT U.id, U.firstName, U.lastName, U.password, U.email, UE.isConfirmed, U.phone, U.createdAt, U.role, ASH.id
  FROM volunteers AS U
  INNER JOIN animal_shelters as ASH
  RIGHT JOIN volunteers_on_shelters AS VOSH ON U.id = VOSH.volunteerID
  AND ASH.id = VOSH.shelterID
  LEFT JOIN usersEmail AS UE ON U.email = UE.email
 WHERE U.id = ?`, userID)
	info := new(entity.User)
	var createdAt int64
	err := row.Scan(
		&info.ID, &info.FirstName, &info.LastName, &info.Password,
		&info.Email, &info.IsVerified, &info.Phone, &createdAt, &info.Role, &info.ShelterID,
	)
	info.CreatedAt = time.Unix(createdAt, 0).Format(time.RFC3339)
	return info, err
}

func GetUserByEmail(ctx context.Context, email string) (*entity.User, error) {
	row := mioDB.QueryRowContext(ctx, `
SELECT U.id, U.firstName, U.lastName, U.password, U.email, U.phone,  U.createdAt, U.role
  FROM volunteers AS U
 WHERE U.email = ?`, email)
	info := new(entity.User)
	var createdAt int64
	err := row.Scan(
		&info.ID, &info.FirstName, &info.LastName, &info.Password,
		&info.Email, &info.Phone, &createdAt, &info.Role,
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

func GetAllUsers(ctx context.Context) ([]entity.User, error) {
	rows, err := mioDB.QueryContext(ctx, `
SELECT 
  V.id, 
  V.firstName, 
  V.lastName, 
  V.role, 
  V.email, 
  V.phone
FROM 
  volunteers AS V 
GROUP BY 
  V.id, 
  V.firstName, 
  V.lastName, 
  V.role, 
  V.email, 
`)
	if err != nil {
		return nil, err
	}
	var users []entity.User
	for rows.Next() {
		var user entity.User
		err = rows.Scan(
			&user.ID,
			&user.FirstName,
			&user.LastName,
			&user.Role,
			&user.Email,
			&user.Phone,
		)
		if err != nil {
			fmt.Println(err)
			return nil, err
		}
		users = append(users, user)
	}

	return users, nil
}

func GetAllUsersBasicInfo(ctx context.Context) ([]entity.User, error) {
	rows, err := mioDB.QueryContext(ctx, `
SELECT 
  V.id, 
  V.firstName, 
  V.lastName, 
  V.email
FROM 
  volunteers AS V 
GROUP BY 
  V.id, 
  V.firstName, 
  V.lastName, 
  V.email
`)
	if err != nil {
		return nil, err
	}
	var users []entity.User
	for rows.Next() {
		var user entity.User
		err = rows.Scan(
			&user.ID,
			&user.FirstName,
			&user.LastName,
			&user.Email,
		)
		if err != nil {
			fmt.Println(err)
			return nil, err
		}
		users = append(users, user)
	}

	return users, nil
}

func UpdateEmailVerification(ctx context.Context, userID int64, token string) error {
	_, err := mioDB.ExecContext(ctx, `
       UPDATE usersEmail
		SET createdAt = UNIX_TIMESTAMP(),
			verificationToken = ?
        WHERE volunteerID = ?`,
		token, userID)
	return err
}

func VerifyEmail(ctx context.Context, token string) (int64, error) {
	res, err := mioDB.ExecContext(ctx, `
UPDATE usersEmail 
SET isConfirmed = TRUE, 
    verificationToken = NULL
WHERE verificationToken = ?`, token)
	if err != nil {
		return 0, err
	}
	return res.RowsAffected()
}

func AddEmail(ctx context.Context, email *entity.Email) error {
	_, err := mioDB.ExecContext(ctx, `
INSERT INTO usersEmail
		SET email = ?,
			volunteerID = ?,
			isConfirmed = ?,
			createdAt = UNIX_TIMESTAMP(),
			verificationToken = ?`,
		email.Email, email.UserID, email.IsVerified, email.VerificationToken)
	return err
}

func GetEmailStatus(ctx context.Context, userID int64) (bool, error) {
	res := mioDB.QueryRowContext(ctx, `
SELECT UE.isConfirmed
FROM usersEmail
WHERE UE.volunteerID = ?`, userID)
	var isConfirmed bool
	err := res.Scan(&isConfirmed)
	if err != nil {
		return false, err
	}
	return isConfirmed, nil
}

func RemoveUserEmail(ctx context.Context, userID int64) error {
	_, err := mioDB.ExecContext(ctx, `
	DELETE FROM usersEmail
	where volunteerID = ?`, userID)
	if err != nil {
		return err
	}
	return nil
}

func GetSheltersRequests(ctx context.Context) (*[]entity.ShelterConfirmRequestInfo, error) {
	query := `
		SELECT
    usersShelters.createdAt AS createdAt,
    volunteers.id AS volunteerID,
    volunteers.firstName AS volunteerFirstName,
    volunteers.lastName AS volunteerLastName,
    volunteers.email AS volunteerEmail,
    volunteers.phone AS volunteerPhone,
    animal_shelters.id AS shelterID,
    animal_shelters.shelter_name AS shelterName,
    animal_shelters.description AS shelterDescription,
    animal_shelters.phone AS shelterPhone,
    animal_shelters.adress AS shelterAddress,
	usersEmail.isConfirmed AS isVolunteerConfirmed
FROM
    usersShelters
INNER JOIN
    volunteers ON usersShelters.requestedBy = volunteers.id
INNER JOIN
    animal_shelters ON usersShelters.shelterID = animal_shelters.id
INNER JOIN 
    usersEmail ON usersShelters.requestedBy = usersEmail.volunteerID
WHERE
    usersShelters.isConfirmed = FALSE
	`

	rows, err := mioDB.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var results []entity.ShelterConfirmRequestInfo
	for rows.Next() {
		var (
			createdAt            string
			volunteerID          int64
			volunteerFirstName   string
			volunteerLastName    string
			volunteerEmail       string
			volunteerPhone       string
			shelterID            int64
			shelterName          string
			shelterDescription   string
			shelterPhone         string
			shelterAddress       string
			isVolunteerConfirmed bool
		)

		if err = rows.Scan(
			&createdAt,
			&volunteerID,
			&volunteerFirstName,
			&volunteerLastName,
			&volunteerEmail,
			&volunteerPhone,
			&shelterID,
			&shelterName,
			&shelterDescription,
			&shelterPhone,
			&shelterAddress,
			&isVolunteerConfirmed,
		); err != nil {
			return nil, err
		}

		requestedVolunteer := entity.User{
			ID:         volunteerID,
			FirstName:  volunteerFirstName,
			LastName:   volunteerLastName,
			Email:      volunteerEmail,
			Phone:      volunteerPhone,
			IsVerified: isVolunteerConfirmed,
		}

		shelter := entity.Shelter{
			ID:          shelterID,
			Name:        shelterName,
			Description: shelterDescription,
			Phone:       shelterPhone,
			Address:     shelterAddress,
		}

		results = append(results, entity.ShelterConfirmRequestInfo{
			RequestedVolunteer: requestedVolunteer,
			Shelter:            shelter,
			CreatedAt:          createdAt,
		})
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return &results, nil
}

func GetInvitationByID(ctx context.Context, id int64) (shID, recID, sendID int64, err error) {
	res := mioDB.QueryRowContext(ctx, `
SELECT senderID, recipientID, shelterID
FROM volunteersInvitationsOnShelters
WHERE id=?`, id)

	err = res.Scan(&sendID, &recID, &shID)
	if err != nil {
		return 0, 0, 0, err
	}
	return shID, recID, sendID, nil
}

func GetInvitationsByUserID(ctx context.Context, uID int64) (*[]entity.ShelterInvitation, error) {
	rows, err := mioDB.QueryContext(ctx, `
	SELECT 	
	    	v.id AS id,
            v1.firstName AS senderFirstName,
            v1.lastName AS senderLastName,
            s.shelter_name AS shelterName,
            s.phone AS shelterPhone,
            s.adress AS shelterAddress
        FROM 
            volunteersInvitationsOnShelters v
        JOIN 
            volunteers v1 ON v.senderID = v1.id
        JOIN 
            volunteers v2 ON v.recipientID = v2.id
        JOIN 
            animal_shelters s ON v.shelterID = s.id
        WHERE 
            v.recipientID = ?
`, uID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	invitations := make([]entity.ShelterInvitation, 0)

	for rows.Next() {
		var invitation entity.ShelterInvitation
		err = rows.Scan(
			&invitation.ID,
			&invitation.From.FirstName,
			&invitation.From.LastName,
			&invitation.InvitedTo.Name,
			&invitation.InvitedTo.Phone,
			&invitation.InvitedTo.Address,
		)
		if err != nil {
			return nil, err
		}
		invitations = append(invitations, invitation)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}
	return &invitations, nil
}

func DeleteInvitation(ctx context.Context, id int64) error {
	_, err := mioDB.ExecContext(ctx, `
	DELETE FROM volunteersInvitationsOnShelters WHERE id=?`, id)
	if err != nil {
		return err
	}
	return nil
}
