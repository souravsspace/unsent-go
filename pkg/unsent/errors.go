package unsent

import "fmt"

// APIError represents an error response from the API
type APIError struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

func (e *APIError) Error() string {
	if e == nil {
		return "nil APIError"
	}
	return fmt.Sprintf("API Error: %s (code: %s)", e.Message, e.Code)
}

// HTTPError represents an HTTP error from the API
type HTTPError struct {
	StatusCode int
	APIErr     APIError
	Method     string
	Path       string
}

func (e *HTTPError) Error() string {
	return fmt.Sprintf("%s %s -> %d %s: %s", e.Method, e.Path, e.StatusCode, e.APIErr.Code, e.APIErr.Message)
}
