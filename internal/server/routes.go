package server

import (
	"effective-mobile/online-subscriptions/internal/handlers"
	"effective-mobile/online-subscriptions/internal/repositories"
	"github.com/gin-gonic/gin"
)

func (app *App) LoadRoutes() {
	router := gin.Default()

	subscription_group := router.Group("/subscriptions")
	app.LoadSubscriptionRoutes(subscription_group)

	app.Router = router
}

func (app *App) LoadSubscriptionRoutes(router *gin.RouterGroup) {
	subscription_repository := repositories.NewSubscriptionRepository(app.Postgres_db)
	subscription_handler := handlers.NewSubscriptionHandler(subscription_repository)

	router.POST("/", subscription_handler.Create)
	router.GET("/", subscription_handler.List)
	router.GET("/:id", subscription_handler.GetByID)
	router.DELETE("/:id", subscription_handler.Delete)
	router.PUT("/:id", subscription_handler.Update)
	router.GET("/total", subscription_handler.CalculateTotalCost)
}
