package requests

import (
	"github.com/google/uuid"
	"time"
)

type CalculateTotalCostRequest struct {
	// UserID is an optional filter parameter for user ID.
	// @example 60601fee-2bf1-4721-ae6f-7636e79a0cba
	UserID *uuid.UUID `form:"user_id"`
	// ServiceName is an optional filter parameter for service name.
	// @example Yandex Plus
	ServiceName string `form:"service_name"`
	// PeriodStart is an optional start date for the period filter.
	// Format expected in RFC3339 (e.g., "2025-01-01T00:00:00Z").
	// @example 2025-01-01T00:00:00Z
	PeriodStart *time.Time `form:"period_start"`
	// PeriodEnd is an optional end date for the period filter.
	// Format expected in RFC3339.
	// @example 2025-12-31T23:59:59Z
	PeriodEnd *time.Time `form:"period_end"`
}
