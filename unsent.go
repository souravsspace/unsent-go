package unsent

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

const DefaultBaseURL = "https://api.unsent.dev"

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

// Client is the main client for the Unsent API
type Client struct {
	Key          string
	URL          string
	RaiseOnError bool
	HTTPClient   *http.Client

	// Resource clients
	Emails    *EmailsClient
	Contacts  *ContactsClient
	Campaigns *CampaignsClient
	Domains   *DomainsClient
}

// NewClient creates a new Unsent API client
func NewClient(key string, options ...ClientOption) (*Client, error) {
	if key == "" {
		key = os.Getenv("UNSENT_API_KEY")
	}
	if key == "" {
		return nil, fmt.Errorf("missing API key. Pass it to NewClient or set UNSENT_API_KEY environment variable")
	}

	baseURL := os.Getenv("UNSENT_BASE_URL")
	if baseURL == "" {
		baseURL = DefaultBaseURL
	}

	client := &Client{
		Key:          key,
		URL:          baseURL + "/v1",
		RaiseOnError: true,
		HTTPClient:   &http.Client{},
	}

	// Apply options
	for _, opt := range options {
		opt(client)
	}

	// Initialize resource clients
	client.Emails = &EmailsClient{client: client}
	client.Contacts = &ContactsClient{client: client}
	client.Campaigns = &CampaignsClient{client: client}
	client.Domains = &DomainsClient{client: client}

	return client, nil
}

// ClientOption is a function that configures a Client
type ClientOption func(*Client)

// WithBaseURL sets a custom base URL
func WithBaseURL(url string) ClientOption {
	return func(c *Client) {
		c.URL = url + "/v1"
	}
}

// WithHTTPClient sets a custom HTTP client
func WithHTTPClient(httpClient *http.Client) ClientOption {
	return func(c *Client) {
		c.HTTPClient = httpClient
	}
}

// WithRaiseOnError sets whether to raise errors on non-2xx responses
func WithRaiseOnError(raise bool) ClientOption {
	return func(c *Client) {
		c.RaiseOnError = raise
	}
}

// request performs an HTTP request and returns the response data and error
func (c *Client) request(method, path string, body interface{}) (interface{}, *APIError) {
	var reqBody io.Reader
	if body != nil {
		jsonData, err := json.Marshal(body)
		if err != nil {
			return nil, &APIError{Code: "INTERNAL_ERROR", Message: err.Error()}
		}
		reqBody = bytes.NewBuffer(jsonData)
	}

	req, err := http.NewRequest(method, c.URL+path, reqBody)
	if err != nil {
		return nil, &APIError{Code: "INTERNAL_ERROR", Message: err.Error()}
	}

	req.Header.Set("Authorization", "Bearer "+c.Key)
	req.Header.Set("Content-Type", "application/json")

	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, &APIError{Code: "INTERNAL_ERROR", Message: err.Error()}
	}
	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, &APIError{Code: "INTERNAL_ERROR", Message: err.Error()}
	}

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		var apiErr APIError
		if err := json.Unmarshal(respBody, &apiErr); err != nil {
			apiErr = APIError{Code: "INTERNAL_SERVER_ERROR", Message: resp.Status}
		}
		if c.RaiseOnError {
			return nil, &apiErr
		}
		return nil, &apiErr
	}

	var result interface{}
	if len(respBody) > 0 {
		if err := json.Unmarshal(respBody, &result); err != nil {
			return nil, &APIError{Code: "INTERNAL_ERROR", Message: err.Error()}
		}
	}

	return result, nil
}

// Post performs a POST request
func (c *Client) Post(path string, body interface{}) (interface{}, *APIError) {
	return c.request("POST", path, body)
}

// Get performs a GET request
func (c *Client) Get(path string) (interface{}, *APIError) {
	return c.request("GET", path, nil)
}

// Put performs a PUT request
func (c *Client) Put(path string, body interface{}) (interface{}, *APIError) {
	return c.request("PUT", path, body)
}

// Patch performs a PATCH request
func (c *Client) Patch(path string, body interface{}) (interface{}, *APIError) {
	return c.request("PATCH", path, body)
}

// Delete performs a DELETE request
func (c *Client) Delete(path string, body interface{}) (interface{}, *APIError) {
	return c.request("DELETE", path, body)
}
