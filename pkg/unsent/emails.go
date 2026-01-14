package unsent

import "fmt"

// EmailsClient handles email-related API operations
type EmailsClient struct {
	client *Client
}

// Send is an alias for Create
func (e *EmailsClient) Send(payload SendEmailJSONBody, opts ...RequestOption) (*EmailCreateResponse, *APIError) {
	return e.Create(payload, opts...)
}

// Create sends a new email
func (e *EmailsClient) Create(payload SendEmailJSONBody, opts ...RequestOption) (*EmailCreateResponse, *APIError) {
	return Post[EmailCreateResponse](e.client, "/emails", payload, opts...)
}

// Batch sends multiple emails in a batch
func (e *EmailsClient) Batch(emails SendBatchEmailsJSONBody, opts ...RequestOption) (*EmailBatchResponse, *APIError) {
	return Post[EmailBatchResponse](e.client, "/emails/batch", emails, opts...)
}

// Get retrieves an email by ID
func (e *EmailsClient) Get(emailID string) (*Email, *APIError) {
	return Get[Email](e.client, fmt.Sprintf("/emails/%s", emailID))
}

// Update updates a scheduled email
func (e *EmailsClient) Update(emailID string, payload UpdateEmailJSONBody) (*EmailUpdateResponse, *APIError) {
	return Patch[EmailUpdateResponse](e.client, fmt.Sprintf("/emails/%s", emailID), payload)
}

// Cancel cancels a scheduled email
func (e *EmailsClient) Cancel(emailID string) (*EmailCancelResponse, *APIError) {
	return Post[EmailCancelResponse](e.client, fmt.Sprintf("/emails/%s/cancel", emailID), map[string]interface{}{})
}

// List retrieves a list of sent emails with optional filters
func (e *EmailsClient) List(params ListEmailsParams) (*ListEmailsResponse, *APIError) {
	path := "/emails?"
	if params.Page != nil {
		path += fmt.Sprintf("page=%s&", *params.Page)
	}
	if params.Limit != nil {
		path += fmt.Sprintf("limit=%s&", *params.Limit)
	}
	if params.StartDate != nil {
		path += fmt.Sprintf("startDate=%s&", params.StartDate.Format("2006-01-02T15:04:05Z"))
	}
	if params.EndDate != nil {
		path += fmt.Sprintf("endDate=%s&", params.EndDate.Format("2006-01-02T15:04:05Z"))
	}
	return Get[ListEmailsResponse](e.client, path)
}

// GetBounces retrieves a list of bounced emails
func (e *EmailsClient) GetBounces(params GetBouncesParams) (*GetBouncesResponse, *APIError) {
	path := "/emails/bounces?"
	if params.Page != nil {
		path += fmt.Sprintf("page=%f&", *params.Page)
	}
	if params.Limit != nil {
		path += fmt.Sprintf("limit=%f&", *params.Limit)
	}
	return Get[GetBouncesResponse](e.client, path)
}

// GetComplaints retrieves a list of spam complaints
func (e *EmailsClient) GetComplaints(params GetComplaintsParams) (*GetComplaintsResponse, *APIError) {
	path := "/emails/complaints?"
	if params.Page != nil {
		path += fmt.Sprintf("page=%f&", *params.Page)
	}
	if params.Limit != nil {
		path += fmt.Sprintf("limit=%f&", *params.Limit)
	}
	return Get[GetComplaintsResponse](e.client, path)
}

// GetUnsubscribes retrieves a list of un subscribed emails
func (e *EmailsClient) GetUnsubscribes(params GetUnsubscribesParams) (*GetUnsubscribesResponse, *APIError) {
	path := "/emails/unsubscribes?"
	if params.Page != nil {
		path += fmt.Sprintf("page=%f&", *params.Page)
	}
	if params.Limit != nil {
		path += fmt.Sprintf("limit=%f&", *params.Limit)
	}
	return Get[GetUnsubscribesResponse](e.client, path)
}

// GetEvents retrieves events for a specific email
func (e *EmailsClient) GetEvents(emailID string, params GetEmailEventsParams) (*GetEmailEventsResponse, *APIError) {
	path := fmt.Sprintf("/emails/%s/events", emailID)
	
	// Build query parameters
	query := buildQueryParams(map[string]interface{}{
		"page":      params.Page,
		"limit":     params.Limit,
		"status":    params.Status,
		"startDate": params.StartDate,
	})
	
	if query != "" {
		path = fmt.Sprintf("%s?%s", path, query)
	}
	
	return Get[GetEmailEventsResponse](e.client, path)
}


