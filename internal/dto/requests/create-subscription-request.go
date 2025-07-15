package requests

import (
	"github.com/google/uuid"
	"time"
)

type CreateSubscriptionRequest struct {
	ServerName string     `json:"server_name"`
	Price      int        `json:"price"`
	UserID     uuid.UUID  `json:"user_id"`
	StartDate  string     `json:"start_date"`
	EndDate    *time.Time `json:"end_date,omitempty"`
}
