package util

type ErrorResponse struct {
	Message    string
	StatusCode int
}

func (e *ErrorResponse) Error() string {
	return e.Message
}

func NewErrorResponse(message string, statusCode int) *ErrorResponse {
	return &ErrorResponse{
		Message:    message,
		StatusCode: statusCode,
	}
}
