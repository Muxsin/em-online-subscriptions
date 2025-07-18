package responses

// ErrorResponse represents a generic error response structure for API failures.
type ErrorResponse struct {
	// Message describes the error that occurred.
	// @example "Subscription not found"
	Message string `json:"error"`
}
