package dto
// ErrorResponse - common error response format for application
type ErrorResponse struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
}
