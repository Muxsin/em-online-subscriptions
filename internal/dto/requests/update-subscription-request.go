package requests

import (
	"github.com/google/uuid"
	"time"
)

type UpdateSubscriptionRequest struct {
	ServiceName string     `json:"service_name" binding:"required"`
	Price       int        `json:"price" binding:"required"`
	UserID      uuid.UUID  `json:"user_id" binding:"required"`
	StartDate   string     `json:"start_date" binding:"required"`
	EndDate     *time.Time `json:"end_date,omitempty"`
}
