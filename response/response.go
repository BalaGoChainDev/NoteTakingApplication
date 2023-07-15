package response

type ErrorResponder interface {
	Error() map[string]any
	GetStatusCode() int
}

type ErrorResponse struct {
	Message    string `json:"message"`
	StatusCode int    `json:"statusCode"`
}

// NewErrorResponse creates a new error response.
func NewErrorResponse(statusCode int, message string) *ErrorResponse {
	return &ErrorResponse{
		Message:    message,
		StatusCode: statusCode,
	}
}

// Error returns the error message.
func (e *ErrorResponse) Error() map[string]any {
	return map[string]any{"message": e.Message}
}

// GetStatusCode returns the HTTP status code.
func (e *ErrorResponse) GetStatusCode() int {
	return e.StatusCode
}
