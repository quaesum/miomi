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
