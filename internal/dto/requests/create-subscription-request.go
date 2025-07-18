package requests

import (
	"github.com/google/uuid"
	"time"
)

// Date format constant for StartDate.
const MM_YYYY_FORMAT = "01-2006" // "MM-YYYY" format (e.g., "07-2025")

// CreateSubscriptionRequest defines the request payload for creating a new online subscription.
type CreateSubscriptionRequest struct {
	// Name of the subscription service.
	// @example Yandex Plus
	ServiceName string `json:"service_name" binding:"required"`
	// Price of the subscription.
	// @example 400
	Price int `json:"price" binding:"required"`
	// User ID associated with the subscription.
	// @example 60601fee-2bf1-4721-ae6f-7636e79a0cba
	UserID uuid.UUID `json:"user_id" binding:"required"`
	// Start date of the subscription in "MM-YYYY" format.
	// @example 07-2025
	StartDate string `json:"start_date" binding:"required"`
	// Optional end date of the subscription in RFC3339 format.
	// @example 2026-07-31T00:00:00Z
	EndDate *time.Time `json:"end_date,omitempty"`
}
