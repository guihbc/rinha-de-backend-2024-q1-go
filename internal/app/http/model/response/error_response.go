package response

type ErrorResponse struct {
	Message string `json:"mensagem"`
}

func NewErrorResponse(message string) *ErrorResponse {
	return &ErrorResponse{
		Message: message,
	}
}
