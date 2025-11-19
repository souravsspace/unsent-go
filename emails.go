package unsent

import "fmt"

// EmailsClient handles email-related API operations
type EmailsClient struct {
	client *Client
}

// Send is an alias for Create
func (e *EmailsClient) Send(payload EmailCreate) (*EmailCreateResponse, *APIError) {
	return e.Create(payload)
}

// Create sends a new email
func (e *EmailsClient) Create(payload EmailCreate) (*EmailCreateResponse, *APIError) {
	data, err := e.client.Post("/emails", payload)
	if err != nil {
		return nil, err
	}

	// Convert interface{} to EmailCreateResponse
	result := &EmailCreateResponse{}
	if dataMap, ok := data.(map[string]interface{}); ok {
		result.ID = dataMap["id"].(string)
		result.To = dataMap["to"].(string)
		result.From = dataMap["from"].(string)
		result.Subject = dataMap["subject"].(string)
		result.Status = dataMap["status"].(string)
	}

	return result, nil
}

// Batch sends multiple emails in a batch
func (e *EmailsClient) Batch(emails []EmailBatchItem) (*EmailBatchResponse, *APIError) {
	data, err := e.client.Post("/emails/batch", emails)
	if err != nil {
		return nil, err
	}

	result := &EmailBatchResponse{}
	if dataMap, ok := data.(map[string]interface{}); ok {
		if emailsData, ok := dataMap["emails"].([]interface{}); ok {
			for _, emailData := range emailsData {
				if emailMap, ok := emailData.(map[string]interface{}); ok {
					email := EmailCreateResponse{
						ID:      emailMap["id"].(string),
						To:      emailMap["to"].(string),
						From:    emailMap["from"].(string),
						Subject: emailMap["subject"].(string),
						Status:  emailMap["status"].(string),
					}
					result.Emails = append(result.Emails, email)
				}
			}
		}
	}

	return result, nil
}

// Get retrieves an email by ID
func (e *EmailsClient) Get(emailID string) (*Email, *APIError) {
	data, err := e.client.Get(fmt.Sprintf("/emails/%s", emailID))
	if err != nil {
		return nil, err
	}

	result := &Email{}
	if dataMap, ok := data.(map[string]interface{}); ok {
		result.ID = dataMap["id"].(string)
		result.To = dataMap["to"].(string)
		result.From = dataMap["from"].(string)
		result.Subject = dataMap["subject"].(string)
		result.Status = dataMap["status"].(string)
	}

	return result, nil
}

// Update updates a scheduled email
func (e *EmailsClient) Update(emailID string, payload EmailUpdate) (*EmailUpdateResponse, *APIError) {
	data, err := e.client.Patch(fmt.Sprintf("/emails/%s", emailID), payload)
	if err != nil {
		return nil, err
	}

	result := &EmailUpdateResponse{}
	if dataMap, ok := data.(map[string]interface{}); ok {
		result.ID = dataMap["id"].(string)
		result.Status = dataMap["status"].(string)
	}

	return result, nil
}

// Cancel cancels a scheduled email
func (e *EmailsClient) Cancel(emailID string) (*EmailCancelResponse, *APIError) {
	data, err := e.client.Post(fmt.Sprintf("/emails/%s/cancel", emailID), map[string]interface{}{})
	if err != nil {
		return nil, err
	}

	result := &EmailCancelResponse{}
	if dataMap, ok := data.(map[string]interface{}); ok {
		result.ID = dataMap["id"].(string)
		result.Status = dataMap["status"].(string)
	}

	return result, nil
}
