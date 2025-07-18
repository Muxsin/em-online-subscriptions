package responses

import (
	"github.com/google/uuid"
	"time"
)

// SubscriptionResponse represents the response structure for an online subscription.
type SubscriptionResponse struct {
	// ID of the subscription.
	// @example 1
	ID uint `json:"id"`
	// Name of the subscription service.
	// @example Yandex Plus
	ServiceName string `json:"service_name"`
	// Price of the subscription.
	// @example 400
	Price int `json:"price"`
	// User ID associated with the subscription.
	// @example 60601fee-2bf1-4721-ae6f-7636e79a0cba
	UserID uuid.UUID `json:"user_id"`
	// Start date of the subscription in "MM-YYYY" format.
	// @example 07-2025
	StartDate string `json:"start_date"`
	// Optional end date of the subscription in RFC3339 format.
	// @example 2026-07-31T00:00:00Z
	EndDate *time.Time `json:"end_date"`
	// Creation timestamp of the subscription in RFC3339 format.
	// @example 2025-07-17T10:30:00Z
	CreatedAt string `json:"created_at"`
}
