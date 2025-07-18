package responses

type TotalCostResponse struct {
	// TotalCost is the aggregated sum of subscription prices.
	// @example 1200
	TotalCost int `json:"total_cost"`
}
