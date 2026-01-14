package unsent

import "fmt"

// WebhooksClient handles webhook-related API operations
type WebhooksClient struct {
	client *Client
}

// List retrieves all webhooks
func (w *WebhooksClient) List() (*[]Webhook, *APIError) {
	return Get[[]Webhook](w.client, "/webhooks")
}

// Get retrieves a webhook by ID
func (w *WebhooksClient) Get(webhookID string) (*Webhook, *APIError) {
	return Get[Webhook](w.client, fmt.Sprintf("/webhooks/%s", webhookID))
}

// Create creates a new webhook
func (w *WebhooksClient) Create(payload CreateWebhookJSONBody) (*WebhookCreateResponse, *APIError) {
	return Post[WebhookCreateResponse](w.client, "/webhooks", payload)
}

// Update updates a webhook
func (w *WebhooksClient) Update(webhookID string, payload UpdateWebhookJSONBody) (*WebhookUpdateResponse, *APIError) {
	return Patch[WebhookUpdateResponse](w.client, fmt.Sprintf("/webhooks/%s", webhookID), payload)
}

// Delete deletes a webhook
func (w *WebhooksClient) Delete(webhookID string) (*WebhookDeleteResponse, *APIError) {
	return Delete[WebhookDeleteResponse](w.client, fmt.Sprintf("/webhooks/%s", webhookID), nil)
}

// WebhookTestResponse represents the response from testing a webhook
type WebhookTestResponse struct {
	ID             string   `json:"id"`
	Type           string   `json:"type"`
	CreatedAt      string   `json:"createdAt"`
	UpdatedAt      string   `json:"updatedAt"`
	TeamID         string   `json:"teamId"`
	Status         string   `json:"status"`
	WebhookID      string   `json:"webhookId"`
	Payload        string   `json:"payload"`
	Attempt        float32  `json:"attempt"`
	NextAttemptAt  *string  `json:"nextAttemptAt,omitempty"`
	LastError      *string  `json:"lastError,omitempty"`
	ResponseStatus *float32 `json:"responseStatus,omitempty"`
	ResponseTimeMs *float32 `json:"responseTimeMs,omitempty"`
	ResponseText   *string  `json:"responseText,omitempty"`
}

// Test triggers a test event for a webhook
func (w *WebhooksClient) Test(webhookID string) (*WebhookTestResponse, *APIError) {
	return Post[WebhookTestResponse](w.client, fmt.Sprintf("/webhooks/%s/test", webhookID), nil)
}
