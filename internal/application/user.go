package application

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"madmax/internal/application/db/mysql"
	"madmax/internal/entity"
	"madmax/internal/utils"
	"strconv"
)

func UserCreate(ctx context.Context, userData *entity.UserCreateRequest) (string, error) {
	_, err := mysql.GetUserByEmail(ctx, userData.Email)
	if err != sql.ErrNoRows {
		return "", errors.New("user exist")
	}
	if userData.ShelterID == 0 {
		sc := entity.ShelterCreateRequest{
			Name:    userData.FirstName + " " + userData.LastName,
			Address: userData.Address,
			Phone:   userData.Phone,
		}
		shID, err := mysql.CreateAnimalShelter(ctx, &sc)
		if err != nil {
			return "", err
		}
		userData.ShelterID = shID
	}
	u := entity.User{
		FirstName: userData.FirstName,
		LastName:  userData.LastName,
		Password:  utils.GetMD5Hash(userData.Password),
		Email:     userData.Email,
		Phone:     userData.Phone,
		Role:      utils.UserRoleVolunteer,
	}
	userID, err := mysql.CreateUser(ctx, &u)
	if err != nil {
		return "", err
	}

	err = mysql.AddUserOnShelter(ctx, userID, userData.ShelterID)
	if err != nil {
		return "", err
	}

	token, err := utils.GenerateToken(strconv.FormatInt(userID, 10), utils.UserRoleVolunteer)
	if err != nil {
		return "", err
	}
	in := entity.Email{
		Email:             userData.Email,
		UserID:            userID,
		VerificationToken: token,
		IsVerified:        false,
	}
	err = addEmail(ctx, &in)
	if err != nil {
		return "", err
	}
	return token, nil

}

func LogIn(ctx context.Context, email, pass string) (string, error) {
	user, err := mysql.GetUserByEmail(ctx, email)
	if err != nil {
		fmt.Println(err)
		return "", errors.New("user not exist")
	}
	if utils.GetMD5Hash(pass) != user.Password {
		return "", errors.New("password does not match")
	}
	token, err := utils.GenerateToken(strconv.FormatInt(user.ID, 10), user.Role)
	if err != nil {
		return "", err
	}
	return token, nil
}

func UserByID(ctx context.Context, id int64) (*entity.User, error) {
	return mysql.GetUserByID(ctx, id)
}

func UserUpdate(ctx context.Context, userID int64, userData *entity.UserCreateRequest) error {
	u, err := mysql.GetUserByID(ctx, userID)
	if err != nil {
		return errors.New("user not exist")
	}
	fmt.Println(userData)
	u.Email = userData.Email
	u.FirstName = userData.FirstName
	u.LastName = userData.LastName

	err = mysql.UpdateUser(ctx, u)
	if err != nil {
		return err
	}
	return nil
}

func GetAllUsers(ctx context.Context) ([]entity.User, error) {
	return mysql.GetAllUsers(ctx)
}

func GetAllUsersBasicInfo(ctx context.Context) ([]entity.User, error) {
	return mysql.GetAllUsersBasicInfo(ctx)
}

func RemoveUser(ctx context.Context, id int64) error {
	err := mysql.RemoveServiceByUserID(ctx, id)
	if err != nil {
		return err
	}
	shID, err := mysql.GetShelterByVolunteerID(ctx, id)
	if err != nil {
		return err
	}
	count, err := mysql.GetShelterParticipatorsCount(ctx, shID)

	err = mysql.RemoveUserFromShelter(ctx, id)
	if err != nil {
		return err
	}
	err = mysql.RemoveUserEmail(ctx, id)
	if err != nil {
		return err
	}

	if count > 1 {
		return nil
	}

	animalsIDs, err := mysql.GetAnimalsByShID(ctx, shID)
	if err != nil {
		return err
	}
	for _, anID := range animalsIDs {
		err = mysql.RemoveAnimalOnShelter(ctx, anID)
		if err != nil {
			return err
		}
		err = mysql.RemoveAnimalOnType(ctx, anID)

		photoIds, err := mysql.GetPhotosIdByAnimalID(ctx, anID)
		if err != nil {
			return err
		}
		for _, pID := range photoIds {
			err = mysql.RemovePhoto(ctx, pID)
			if err != nil {
				return err
			}
		}
	}

	err = mysql.RemoveShelterByID(ctx, shID)
	if err != nil {
		return err
	}

	err = mysql.DeleteUser(ctx, id)
	if err != nil {
		return err
	}

	return nil
}

func VerifyUserEmail(ctx context.Context, userID int64) error {
	userBI, err := UserByID(ctx, userID)
	if err != nil {
		return err // Error is formatted within fetchUserBasicInfo
	}

	emailVerificationToken, err := generateEmailVerificationToken(userID, userBI)
	if err != nil {
		return err
	}

	if err = mysql.UpdateEmailVerification(ctx, userID, emailVerificationToken); err != nil {
		return err
	}

	return sendEmailVerificationMessage(userBI.Email, emailVerificationToken)
}

func VerifyEmail(ctx context.Context, token string) error {
	rowsAffected, err := mysql.VerifyEmail(ctx, token)
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return errors.New("invalid token")
	}
	return nil
}

func addEmail(ctx context.Context, in *entity.Email) error {
	err := mysql.AddEmail(ctx, in)
	if err != nil {
		return err
	}
	return nil
}

func GetSheltersConfirmRequests(ctx context.Context) (*[]entity.ShelterConfirmRequestInfo, error) {
	return mysql.GetSheltersRequests(ctx)
}

func GetInvitationsByID(ctx context.Context, uID int64) (*[]entity.ShelterInvitation, error) {
	return mysql.GetInvitationsByUserID(ctx, uID)
}

func InviteUserToShelter(ctx context.Context, shID, recID, sendID int64) error {
	return mysql.CreateShelterInvitation(ctx, shID, sendID, recID)
}

func AcceptInvitation(ctx context.Context, id, uID int64) error {
	shID, recID, _, err := mysql.GetInvitationByID(ctx, id)
	if err != nil {
		return err
	}
	if recID != uID {
		return errors.New("ids do not match")
	}
	err = MoveUserOnNewShelter(ctx, uID, shID)
	if err != nil {
		return err
	}
	err = mysql.DeleteInvitation(ctx, id)
	if err != nil {
		return err
	}
	return nil
}

func RejectInvitation(ctx context.Context, id, uID int64) error {
	_, recID, _, err := mysql.GetInvitationByID(ctx, id)
	if err != nil {
		return err
	}
	if recID != uID {
		return errors.New("ids do not match")
	}
	err = mysql.DeleteInvitation(ctx, id)
	if err != nil {
		return err
	}
	return nil
}
