package unsent

import "fmt"

// EmailsClient handles email-related API operations
type EmailsClient struct {
	client *Client
}

// Send is an alias for Create
func (e *EmailsClient) Send(payload EmailCreate, opts ...RequestOption) (*EmailCreateResponse, *APIError) {
	return e.Create(payload, opts...)
}

// Create sends a new email
func (e *EmailsClient) Create(payload EmailCreate, opts ...RequestOption) (*EmailCreateResponse, *APIError) {
	return Post[EmailCreateResponse](e.client, "/emails", payload, opts...)
}

// Batch sends multiple emails in a batch
func (e *EmailsClient) Batch(emails []EmailBatchItem, opts ...RequestOption) (*EmailBatchResponse, *APIError) {
	return Post[EmailBatchResponse](e.client, "/emails/batch", emails, opts...)
}

// Get retrieves an email by ID
func (e *EmailsClient) Get(emailID string) (*Email, *APIError) {
	return Get[Email](e.client, fmt.Sprintf("/emails/%s", emailID))
}

// Update updates a scheduled email
func (e *EmailsClient) Update(emailID string, payload EmailUpdate) (*EmailUpdateResponse, *APIError) {
	return Patch[EmailUpdateResponse](e.client, fmt.Sprintf("/emails/%s", emailID), payload)
}

// Cancel cancels a scheduled email
func (e *EmailsClient) Cancel(emailID string) (*EmailCancelResponse, *APIError) {
	return Post[EmailCancelResponse](e.client, fmt.Sprintf("/emails/%s/cancel", emailID), map[string]interface{}{})
}
