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
const Version = "1.0.2"



// Client is the main client for the Unsent API
type Client struct {
	Key          string
	URL          string
	RaiseOnError bool
	HTTPClient   *http.Client

	// Resource clients
	Emails       *EmailsClient
	Contacts     *ContactsClient
	Campaigns    *CampaignsClient
	Domains      *DomainsClient
	Analytics    *AnalyticsClient
	ApiKeys      *ApiKeysClient
	ContactBooks *ContactBooksClient
	Settings     *SettingsClient
	Suppressions *SuppressionsClient
	Templates    *TemplatesClient
	Webhooks     *WebhooksClient
	System       *SystemClient
	Events       *EventsClient
	Metrics      *MetricsClient
	Stats        *StatsClient
	Activity     *ActivityClient
	Teams        *TeamsClient
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
	client.Analytics = &AnalyticsClient{client: client}
	client.ApiKeys = &ApiKeysClient{client: client}
	client.ContactBooks = &ContactBooksClient{client: client}
	client.Settings = &SettingsClient{client: client}
	client.Suppressions = &SuppressionsClient{client: client}
	client.Templates = &TemplatesClient{client: client}
	client.Webhooks = &WebhooksClient{client: client}
	client.System = &SystemClient{client: client}
	client.Events = &EventsClient{client: client}
	client.Metrics = &MetricsClient{client: client}
	client.Stats = &StatsClient{client: client}
	client.Activity = &ActivityClient{client: client}
	client.Teams = &TeamsClient{client: client}

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

// RequestOption is a function that configures a request
type RequestOption func(*http.Request)

// WithHeader sets a header on the request
func WithHeader(key, value string) RequestOption {
	return func(req *http.Request) {
		req.Header.Set(key, value)
	}
}

// WithIdempotencyKey sets the Idempotency-Key header
func WithIdempotencyKey(key string) RequestOption {
	return WithHeader("Idempotency-Key", key)
}

// request performs an HTTP request and returns the response data and error
func request[T any](c *Client, method, path string, body interface{}, opts ...RequestOption) (*T, *APIError) {
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

	// Apply request options
	for _, opt := range opts {
		opt(req)
	}

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
		// Try unmarshaling into flat APIError
		if err := json.Unmarshal(respBody, &apiErr); err != nil || apiErr.Code == "" {
			// If flat failed or resulted in empty struct, try nested error object
			var nestedErr struct {
				Error APIError `json:"error"`
			}
			if err2 := json.Unmarshal(respBody, &nestedErr); err2 == nil && nestedErr.Error.Code != "" {
				apiErr = nestedErr.Error
			} else {
				// Fallback if both fail
				apiErr = APIError{Code: "INTERNAL_SERVER_ERROR", Message: resp.Status}
			}
		}
		
		if c.RaiseOnError {
			return nil, &apiErr
		}
		return nil, &apiErr
	}

	var result T
	if len(respBody) > 0 {
		if err := json.Unmarshal(respBody, &result); err != nil {
			return nil, &APIError{Code: "INTERNAL_ERROR", Message: err.Error()}
		}
	}

	return &result, nil
}

// Post performs a POST request
func Post[T any](c *Client, path string, body interface{}, opts ...RequestOption) (*T, *APIError) {
	return request[T](c, "POST", path, body, opts...)
}

// Get performs a GET request
func Get[T any](c *Client, path string, opts ...RequestOption) (*T, *APIError) {
	return request[T](c, "GET", path, nil, opts...)
}

// Put performs a PUT request
func Put[T any](c *Client, path string, body interface{}, opts ...RequestOption) (*T, *APIError) {
	return request[T](c, "PUT", path, body, opts...)
}

// Patch performs a PATCH request
func Patch[T any](c *Client, path string, body interface{}, opts ...RequestOption) (*T, *APIError) {
	return request[T](c, "PATCH", path, body, opts...)
}

// Delete performs a DELETE request
func Delete[T any](c *Client, path string, body interface{}, opts ...RequestOption) (*T, *APIError) {
	return request[T](c, "DELETE", path, body, opts...)
}
