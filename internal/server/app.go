package server

import (
	"effective-mobile/online-subscriptions/internal/database"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/gorm"
)

type AppInterface interface {
	Run()
	LoadRoutes()
}

type App struct {
	Router      *gin.Engine
	Postgres_db *gorm.DB
}

func New() *App {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}

	db := database.Connect()

	app := &App{
		Postgres_db: db,
	}

	app.LoadRoutes()

	return app
}

func (app *App) Run() {
	err := app.Router.Run()
	if err != nil {
		panic(err)
	}
}
