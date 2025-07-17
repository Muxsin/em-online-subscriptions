package requests

import (
	"github.com/google/uuid"
	"time"
)

type CalculateTotalCostRequest struct {
	UserID      *uuid.UUID `form:"user_id"`
	ServiceName string     `form:"service_name"`
	PeriodStart *time.Time `form:"period_start"`
	PeriodEnd   *time.Time `form:"period_end"`
}
