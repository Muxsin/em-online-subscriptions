package server

import (
	"github.com/gin-gonic/gin"
)

func (app *App) LoadRoutes() {
	router := gin.Default()

	subscription_group := router.Group("/subscriptions")
	app.LoadSubscriptionRoutes(subscription_group)

	app.Router = router
}

func (app *App) LoadSubscriptionRoutes(router *gin.RouterGroup) {
	router.POST("/", app.SubscriptionHandler.Create)
	router.GET("/", app.SubscriptionHandler.List)
	router.GET("/:id", app.SubscriptionHandler.GetByID)
	router.DELETE("/:id", app.SubscriptionHandler.Delete)
	router.PUT("/:id", app.SubscriptionHandler.Update)
	router.GET("/total", app.SubscriptionHandler.CalculateTotalCost)
}
