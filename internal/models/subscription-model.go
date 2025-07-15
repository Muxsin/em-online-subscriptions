package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type Subscription struct {
	gorm.Model

	ServiceName string     `json:"service_name"`
	Price       int        `json:"price"`
	UserID      uuid.UUID  `json:"user_id"`
	StartDate   string     `json:"start_date"`
	EndDate     *time.Time `json:"end_date,omitempty"`
}
