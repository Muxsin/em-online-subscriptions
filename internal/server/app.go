package server

import (
	"effective-mobile/online-subscriptions/internal/database"
	"effective-mobile/online-subscriptions/internal/handlers"
	"effective-mobile/online-subscriptions/internal/repositories"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type AppInterface interface {
	Run()
	LoadRoutes()
}

type App struct {
	Router              *gin.Engine
	SubscriptionHandler handlers.SubscriptionHandlerInterface
}

func New() *App {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}

	db := database.Connect()

	subscription_repository := repositories.NewSubscriptionRepository(db)
	subscription_handler := handlers.NewSubscriptionHandler(subscription_repository)

	app := &App{
		SubscriptionHandler: subscription_handler,
	}

	app.LoadRoutes()

	return app
}

func (app *App) Run() {
	app.Router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	err := app.Router.Run()
	if err != nil {
		panic(err)
	}
}
