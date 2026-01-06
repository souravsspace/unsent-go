package unsent

import "fmt"

// WebhooksClient handles webhook-related API operations
//
// NOTE: Webhook functionality is currently a placeholder for future implementation.
// These methods are included in anticipation of webhook support being added to the
// Unsent API. When webhook endpoints become available in the API, these methods
// will be functional without requiring SDK updates.
//
// Webhook events will allow you to receive real-time notifications about email
// activities (sent, delivered, opened, clicked, bounced, etc.) by registering
// HTTP endpoints that the Unsent API will call when events occur.
type WebhooksClient struct {
	client *Client
}

// List retrieves all webhooks
//
// NOTE: This is a placeholder method for future webhook functionality.
// The webhook API endpoints are not yet available.
func (c *WebhooksClient) List() (*[]Webhook, *APIError) {
	return Get[[]Webhook](c.client, "/webhooks")
}

// Create creates a new webhook
//
// NOTE: This is a placeholder method for future webhook functionality.
// The webhook API endpoints are not yet available.
func (c *WebhooksClient) Create(payload WebhookCreateRequest) (*WebhookCreateResponse, *APIError) {
	return Post[WebhookCreateResponse](c.client, "/webhooks", payload)
}

// Update updates a webhook
//
// NOTE: This is a placeholder method for future webhook functionality.
// The webhook API endpoints are not yet available.
func (c *WebhooksClient) Update(id string, payload WebhookUpdateRequest) (*WebhookUpdateResponse, *APIError) {
	return Patch[WebhookUpdateResponse](c.client, fmt.Sprintf("/webhooks/%s", id), payload)
}

// Delete deletes a webhook
//
// NOTE: This is a placeholder method for future webhook functionality.
// The webhook API endpoints are not yet available.
func (c *WebhooksClient) Delete(id string) (*WebhookDeleteResponse, *APIError) {
	return Delete[WebhookDeleteResponse](c.client, fmt.Sprintf("/webhooks/%s", id), nil)
}
