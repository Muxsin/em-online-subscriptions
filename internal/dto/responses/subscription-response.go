package responses

import (
	"github.com/google/uuid"
	"time"
)

type SubscriptionResponse struct {
	ID          uint       `json:"id"`
	ServiceName string     `json:"service_name"`
	Price       int        `json:"price"`
	UserID      uuid.UUID  `json:"user_id"`
	StartDate   string     `json:"start_date"`
	EndDate     *time.Time `json:"end_date"`
	CreatedAt   string     `json:"created_at"`
}
