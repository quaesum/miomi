package internal

import (
	"context"
	"database/sql"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Run() {
	r := gin.Default()
	r.GET("/cmd", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	r.Run()
}

var (
	ctx context.Context
	db  *sql.DB
)

type DBEntity struct {
	id       int
	animal   string
	name     string
	gender   string
	location string
	age      int
	photo    string
}

/* GetBasicInfo fetches and returns basic information about user.
func GetBasicInfo(ctx context.Context, userID int64) (entity.DBEntity, error) {
	row := db.QueryRowContext(ctx, `
SELECT U.id, U.lastName, U.firstName, U.roleName,  U.timeZone,
    U.role, U.companyID, C.Name, IFNULL(C.photo, '') AS logo,
    UE.email, UE.isConfirmed,
    U.photo, U.isSuperAdmin, U.isRegistered, IFNULL(U.advice, '') AS advice,
    C.bank, C.bic, C.iban, C.nds
  FROM users AS U
    INNER JOIN usersEmail AS UE ON U.id = UE.userID
    INNER JOIN companies AS C ON U.companyID = C.id
 WHERE U.id = ? AND NOT U.isDeleted`, userID)
	info := new(entity.UserBasicInfo)
	err := row.Scan(
		&info.ID, &info.LastName, &info.FirstName, &info.RoleName, &info.TimeZone,
		&info.Role, &info.Company.ID, &info.Company.Name, &info.Company.Photo,
		&info.Email, &info.IsVerified,
		&info.Photo, &info.IsSuperAdmin, &info.IsRegistered, &info.Advice,
		&info.Company.Bank, &info.Company.Bic, &info.Company.IBAN, &info.Company.NDS,
	)
	return info, err
}*/
