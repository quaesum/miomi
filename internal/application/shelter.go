package application

import (
	"context"
	"database/sql"
	"madmax/internal/application/db/mysql"
	"madmax/internal/entity"
)

func ShelterByID(ctx context.Context, id int64) (*entity.Shelter, error) {
	shelter, err := mysql.GetShelterByID(ctx, id)
	if err != nil {
		return nil, err
	}
	isVerified, err := mysql.GetShelterStatus(ctx, shelter.ID)
	if err != nil && err != sql.ErrNoRows {
		return nil, err
	}
	if err == sql.ErrNoRows {
		isVerified = false
	}
	shelter.IsVerified = isVerified
	return shelter, nil
}

func SheltersAll(ctx context.Context) ([]entity.Shelter, error) {
	return mysql.GetAllShelters(ctx)
}

func SheltersAllInfo(ctx context.Context) ([]entity.Shelter, error) {
	return mysql.GetAllSheltersInfo(ctx)
}

func AddUserOnShelter(ctx context.Context, uID, shID int64) error {
	return mysql.AddUserOnShelter(ctx, uID, shID)
}

func MoveUserOnNewShelter(ctx context.Context, uID, shID int64) error {
	err := mysql.RemoveUserFromShelter(ctx, uID)
	if err != nil {
		return err
	}
	err = mysql.AddUserOnShelter(ctx, uID, shID)
	if err != nil {
		return err
	}
	participatorCount, err := mysql.GetShelterParticipatorsCount(ctx, shID)
	if err != nil {
		return err
	}
	if participatorCount == 0 {
		err = mysql.RemoveShelterByID(ctx, shID)
		if err != nil {
			return err
		}
	}
	return nil
}

func ShelterCreateRequest(ctx context.Context, shelter *entity.ShelterCreateRequest, userID int64) (int64, error) {
	shID, err := mysql.CreateAnimalShelter(ctx, shelter)
	if err != nil {
		return 0, err
	}
	err = mysql.CreateShelterRequest(ctx, shID, userID)
	if err != nil {
		return 0, err
	}
	return shID, nil
}

func ShelterCreate(ctx context.Context, shelter *entity.ShelterCreateRequest) (int64, error) {
	return mysql.CreateAnimalShelter(ctx, shelter)
}

func ShelterConfirm(ctx context.Context, id int64) (int64, error) {
	return mysql.ConfirmAnimalShelter(ctx, id)
}

func RemoveVolunteerOnShelter(ctx context.Context, uID int64) error {
	return mysql.RemoveUserFromShelter(ctx, uID)
}

func ShelterReject(ctx context.Context, id int64) error {
	err := mysql.RejectAnimalShelter(ctx, id)
	if err != nil {
		return err
	}
	err = mysql.RemoveShelterByID(ctx, id)
	if err != nil {
		return err
	}
	return nil
}

func ShelterUpdate(ctx context.Context, id int64, shelter *entity.ShelterCreateRequest) error {
	return mysql.UpdateShelter(ctx, id, shelter)
}

func ShelterDelete(ctx context.Context, id int64) error {
	return mysql.RemoveShelterByID(ctx, id)
}

func ExitFromShelter(ctx context.Context, user *entity.User) (int64, error) {
	shelterUser, err := mysql.GetShelterByVolunteerID(ctx, user.ID)
	if err != nil {
		return 0, err
	}
	err = mysql.RemoveUserFromShelter(ctx, user.ID)
	if err != nil {
		return 0, err
	}
	participatorCount, err := mysql.GetShelterParticipatorsCount(ctx, shelterUser)
	if err != nil {
		return 0, err
	}
	if participatorCount == 0 {
		err = mysql.RemoveShelterByID(ctx, shelterUser)
		if err != nil {
			return 0, err
		}
	}

	shReq := entity.CompareUserToShelter(user)
	shID, err := mysql.CreateAnimalShelter(ctx, shReq)
	if err != nil {
		return 0, err
	}
	err = mysql.AddUserOnShelter(ctx, user.ID, shID)
	if err != nil {
		return 0, err
	}
	return shID, nil
}

func GetShelterParticipators(ctx context.Context, shelterID int64) ([]entity.User, error) {
	return mysql.GetSheltersParticipators(ctx, shelterID)
}
