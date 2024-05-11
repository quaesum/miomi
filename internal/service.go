package internal

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/blevesearch/bleve"
	"github.com/gin-gonic/gin"
	rabbitmq2 "github.com/wagslane/go-rabbitmq"
	"log"
	bleve2 "madmax/internal/application/db/bleve"
	"madmax/internal/application/db/mysql"
	"madmax/internal/application/db/rabbitmq"
	"madmax/internal/transport/http/v1"
	v2 "madmax/internal/transport/http/v2"
	"madmax/internal/utils"
	"net/http"
	"os"
	"os/signal"
	"time"
)

// App contains all what needs to run the server
type App struct {
	HTTPServer    http.Server
	SQLDB         *sql.DB
	Config        *Config
	BleveProducts bleve.Index
	BleveAnimals  bleve.Index
	BleveServices bleve.Index
	publisher     *rabbitmq2.Publisher
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
		return nil, fmt.Errorf("failed to initialize mysql client: %w", err)
	}

	err = bleve2.NewBleve()
	if err != nil {
		return nil, fmt.Errorf("failed to initialize bleve: %w", err)
	}

	go utils.RabbitConnect()

	publisher, err := rabbitmq.SetupRabbitMQ()
	if err != nil {
		log.Fatalf("Error setting up RabbitMQ: %v", err)
	}
	defer func() {
		publisher.Close()
	}()

	app.publisher = publisher

	if err != nil {
		return nil, fmt.Errorf("failed to initialize rabbitmq client: %w", err)
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
	router.Use(v1.GinMiddleware("http://localhost:3000"))
	v1.HandlerHTTP(router)
	v2.HandlerHTTP(router)
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
