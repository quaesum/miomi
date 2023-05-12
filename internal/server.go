package internal

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"madmax/internal/mysql"
	"net/http"
	"os"
	"os/signal"
	"time"
)

// App contains all what needs to run the server
type App struct {
	HTTPServer http.Server
	SQLDB      *sql.DB
	Config     *Config
}

var application *App

func Run() error {
	app, err := NewApp()
	if err != nil {
		return fmt.Errorf("failed to init app: %w", err)
	}
	application = app
	return app.Run()
}

func NewApp() (*App, error) {
	var app = new(App)
	var err error

	app.Config, err = NewConfig()
	if err != nil {
		return nil, fmt.Errorf("failed to open config: %w", err)
	}

	app.SQLDB, err = mysql.NewDB(app.Config.MysqlDSN)
	if err != nil {
		return nil, fmt.Errorf("failed to initialize  mysql client: %w", err)
	}
	return app, err
}

func (app *App) Run() error {
	var errc = app.Start()

	var quit = make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)

	select {
	case <-quit:
		log.Println("caught os signal")
	case err := <-errc:
		log.Printf("caught error: %s", err)
	}

	log.Println("trying to shutdown server")

	return app.Shutdown(context.TODO())
}

// Start runs App and doesn't wait
func (app *App) Start() <-chan error {
	var errc = make(chan error, 1)

	router := gin.Default()
	HTTPHandler(router)
	srv := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	go func() {
		// service connections
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
			errc <- err
		}
	}()

	return errc
}

// Shutdown can be run to clean up all what was run
func (app *App) Shutdown(ctx context.Context) error {
	ctx, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()

	if err := app.HTTPServer.Shutdown(ctx); err != nil {
		return fmt.Errorf("failed to shutdown http server")
	}

	log.Println("server has been shutdown")
	return nil
}

type DBEntity struct {
	id       int
	animal   string
	name     string
	gender   string
	location string
	age      int
	photo    string
}

/*
	r.GET("/cmd", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

GetBasicInfo fetches and returns basic information about user.
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
