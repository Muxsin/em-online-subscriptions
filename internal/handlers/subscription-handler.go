package handlers

import (
	"effective-mobile/online-subscriptions/internal/dto/requests"
	"effective-mobile/online-subscriptions/internal/dto/responses"
	"effective-mobile/online-subscriptions/internal/models"
	"effective-mobile/online-subscriptions/internal/repositories"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"time"
)

type SubscriptionHandlerInterface interface {
	Create(ctx *gin.Context)
	List(ctx *gin.Context)
	GetByID(ctx *gin.Context)
	Update(ctx *gin.Context)
	Delete(ctx *gin.Context)
}

type SubscriptionHandler struct {
	Repository repositories.SubscriptionRepositoryInterface
}

func NewSubscriptionHandler(repository repositories.SubscriptionRepositoryInterface) SubscriptionHandlerInterface {
	return &SubscriptionHandler{
		Repository: repository,
	}
}

func (h *SubscriptionHandler) Create(ctx *gin.Context) {
	var request requests.CreateSubscriptionRequest

	if err := ctx.ShouldBindBodyWithJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	subscription := &models.Subscription{
		ServiceName: request.ServerName,
		Price:       request.Price,
		UserID:      request.UserID,
		StartDate:   request.StartDate,
		EndDate:     request.EndDate,
	}

	if err := h.Repository.Insert(subscription); err != nil {
		log.Printf("Error inserting subscription: %v", err)

		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "Error inserting subscription",
		})
		return
	}

	response := responses.SubscriptionResponse{
		ID:          subscription.ID,
		ServiceName: subscription.ServiceName,
		Price:       subscription.Price,
		UserID:      subscription.UserID,
		StartDate:   subscription.StartDate,
		EndDate:     subscription.EndDate,
		CreatedAt:   subscription.CreatedAt.Format(time.RFC3339),
	}

	ctx.JSON(http.StatusCreated, response)
}

func (h *SubscriptionHandler) List(ctx *gin.Context) {}

func (h *SubscriptionHandler) GetByID(ctx *gin.Context) {}

func (h *SubscriptionHandler) Update(ctx *gin.Context) {}

func (h *SubscriptionHandler) Delete(ctx *gin.Context) {}
