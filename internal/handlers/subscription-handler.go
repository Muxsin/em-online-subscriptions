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
	CalculateTotalCost(ctx *gin.Context)
}

type SubscriptionHandler struct {
	Repository repositories.SubscriptionRepositoryInterface
}

func NewSubscriptionHandler(repository repositories.SubscriptionRepositoryInterface) SubscriptionHandlerInterface {
	return &SubscriptionHandler{
		Repository: repository,
	}
}

// @Summary Create a new subscription
// @Description Creates a new online subscription.
// @Tags Subscriptions
// @Accept json
// @Produce json
// @Param request body requests.CreateSubscriptionRequest true "Subscription creation details"
// @Success 201 {object} responses.SubscriptionResponse "Subscription created successfully"
// @Failure 400 {object} responses.ErrorResponse "Invalid request payload"
// @Failure 500 {object} responses.ErrorResponse "Internal server error"
// @Router /subscriptions [post]
func (h *SubscriptionHandler) Create(ctx *gin.Context) {
	var request requests.CreateSubscriptionRequest

	if err := ctx.ShouldBindBodyWithJSON(&request); err != nil {
		log.Print(err.Error())

		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	_, err := time.Parse(requests.MM_YYYY_FORMAT, request.StartDate)
	if err != nil {
		log.Printf("Invalid StartDate format: %v", err)
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid StartDate format. Expected MM-YYYY (e.g., 07-2025).",
		})
		return
	}

	subscription := &models.Subscription{
		ServiceName: request.ServiceName,
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

// @Summary Get all subscriptions
// @Description Retrieves a list of all online subscriptions.
// @Tags Subscriptions
// @Produce json
// @Success 200 {array} responses.SubscriptionResponse "List of subscriptions retrieved successfully"
// @Failure 500 {object} responses.ErrorResponse "Internal server error"
// @Router /subscriptions [get]
func (h *SubscriptionHandler) List(ctx *gin.Context) {
	products, err := h.Repository.List()

	if err != nil {
		log.Printf("Error listing subscriptions: %v", err)

		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "Error listing subscriptions",
		})
		return
	}

	var response []responses.SubscriptionResponse

	for _, product := range products {
		response = append(response, responses.SubscriptionResponse{
			ID:          product.ID,
			ServiceName: product.ServiceName,
			Price:       product.Price,
			UserID:      product.UserID,
			StartDate:   product.StartDate,
			EndDate:     product.EndDate,
			CreatedAt:   product.CreatedAt.Format(time.RFC3339),
		})
	}

	ctx.JSON(http.StatusOK, response)
}

// @Summary Get subscription by ID
// @Description Retrieves a single online subscription by its ID.
// @Tags Subscriptions
// @Produce json
// @Param id path string true "Subscription ID"
// @Success 200 {object} responses.SubscriptionResponse "Subscription retrieved successfully"
// @Failure 404 {object} responses.ErrorResponse "Subscription not found"
// @Failure 500 {object} responses.ErrorResponse "Internal server error"
// @Router /subscriptions/{id} [get]
func (h *SubscriptionHandler) GetByID(ctx *gin.Context) {
	product, err := h.Repository.GetByID(ctx.Param("id"))

	if err != nil {
		log.Printf("Error getting subscription: %v", err)

		ctx.JSON(http.StatusNotFound, gin.H{
			"error": "Subscription not found",
		})
		return
	}

	response := responses.SubscriptionResponse{
		ID:          product.ID,
		ServiceName: product.ServiceName,
		Price:       product.Price,
		UserID:      product.UserID,
		StartDate:   product.StartDate,
		EndDate:     product.EndDate,
		CreatedAt:   product.CreatedAt.Format(time.RFC3339),
	}

	ctx.JSON(http.StatusOK, response)
}

// @Summary Update an existing subscription
// @Description Updates details of an existing subscription by ID.
// @Tags Subscriptions
// @Accept json
// @Produce json
// @Param id path string true "Subscription ID"
// @Param request body requests.UpdateSubscriptionRequest true "Subscription update details"
// @Success 200 {object} responses.SubscriptionResponse "Subscription updated successfully"
// @Failure 400 {object} responses.ErrorResponse "Invalid request payload or parameters"
// @Failure 404 {object} responses.ErrorResponse "Subscription not found"
// @Failure 500 {object} responses.ErrorResponse "Internal server error"
// @Router /subscriptions/{id} [put]
func (h *SubscriptionHandler) Update(ctx *gin.Context) {
	subscription, err := h.Repository.GetByID(ctx.Param("id"))

	if err != nil {
		log.Printf("Error getting subscription: %v", err)

		ctx.JSON(http.StatusNotFound, gin.H{
			"error": "Subscription not found",
		})
		return
	}

	var request requests.UpdateSubscriptionRequest

	if err := ctx.ShouldBindJSON(&request); err != nil {
		log.Print(err.Error())

		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	subscription.UserID = request.UserID
	subscription.ServiceName = request.ServiceName
	subscription.Price = request.Price
	subscription.StartDate = request.StartDate
	subscription.EndDate = request.EndDate

	if err := h.Repository.Update(subscription); err != nil {
		log.Printf("Error updating subscription: %v", err)

		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "Error updating subscription",
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

	ctx.JSON(http.StatusOK, response)
}

// @Summary Delete a subscription
// @Description Deletes an online subscription by its ID.
// @Tags Subscriptions
// @Param id path string true "Subscription ID"
// @Success 200 "Subscription deleted successfully"
// @Failure 404 {object} responses.ErrorResponse "Subscription not found"
// @Failure 500 {object} responses.ErrorResponse "Internal server error"
// @Router /subscriptions/{id} [delete]
func (h *SubscriptionHandler) Delete(ctx *gin.Context) {
	subscription, err := h.Repository.GetByID(ctx.Param("id"))

	if err != nil {
		log.Printf("Error getting subscription: %v", err)

		ctx.JSON(http.StatusNotFound, gin.H{
			"error": "Subscription not found",
		})
		return
	}

	if err := h.Repository.Delete(subscription); err != nil {
		log.Printf("Error deleting subscription: %v", err)

		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "Error deleting subscription",
		})
		return
	}

	ctx.Status(http.StatusOK)
}

// @Summary Calculate total cost of subscriptions
// @Description Calculates the total cost of subscriptions based on user ID, service name, and a period.
// @Tags Subscriptions
// @Accept json
// @Produce json
// @Param user_id query string false "User ID to filter by" format(uuid)
// @Param service_name query string false "Service name to filter by (case-insensitive) (Yandex Plus)"
// @Param period_start query string false "Start date of the period (RFC3339 format, e.g., 2025-01-01T00:00:00Z)" format(date-time)
// @Param period_end query string false "End date of the period (RFC3339 format, e.g., 2025-12-31T23:59:59Z)" format(date-time)
// @Success 200 {object} responses.TotalCostResponse "Total cost calculated successfully"
// @Failure 400 {object} responses.ErrorResponse "Bad request (e.g., invalid query parameters)"
// @Failure 500 {object} responses.ErrorResponse "Internal server error"
// @Router /subscriptions/total [get]
func (h *SubscriptionHandler) CalculateTotalCost(ctx *gin.Context) {
	var request requests.CalculateTotalCostRequest

	if err := ctx.ShouldBindQuery(&request); err != nil {
		log.Printf("Error binding query parameters: %v", err)
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	total_cost, err := h.Repository.CalculateTotalCost(request.UserID, request.ServiceName, request.PeriodStart, request.PeriodEnd)
	if err != nil {
		log.Printf("Error calculating total cost: %v", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "Error calculating total cost",
		})
		return
	}

	ctx.JSON(http.StatusOK, responses.TotalCostResponse{
		TotalCost: total_cost,
	})
}
