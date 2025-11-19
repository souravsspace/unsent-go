package unsent

import "fmt"

// ContactsClient handles contact-related API operations
type ContactsClient struct {
	client *Client
}

// Create creates a new contact
func (c *ContactsClient) Create(bookID string, payload ContactCreate) (*ContactCreateResponse, *APIError) {
	data, err := c.client.Post(fmt.Sprintf("/contactBooks/%s/contacts", bookID), payload)
	if err != nil {
		return nil, err
	}

	result := &ContactCreateResponse{}
	if dataMap, ok := data.(map[string]interface{}); ok {
		result.ID = dataMap["id"].(string)
		result.Email = dataMap["email"].(string)
	}

	return result, nil
}

// Get retrieves a contact by ID
func (c *ContactsClient) Get(bookID, contactID string) (*Contact, *APIError) {
	data, err := c.client.Get(fmt.Sprintf("/contactBooks/%s/contacts/%s", bookID, contactID))
	if err != nil {
		return nil, err
	}

	result := &Contact{}
	if dataMap, ok := data.(map[string]interface{}); ok {
		result.ID = dataMap["id"].(string)
		result.Email = dataMap["email"].(string)
		if firstName, ok := dataMap["firstName"].(string); ok {
			result.FirstName = firstName
		}
		if lastName, ok := dataMap["lastName"].(string); ok {
			result.LastName = lastName
		}
		if metadata, ok := dataMap["metadata"].(map[string]interface{}); ok {
			result.Metadata = metadata
		}
	}

	return result, nil
}

// Update updates a contact
func (c *ContactsClient) Update(bookID, contactID string, payload ContactUpdate) (*ContactUpdateResponse, *APIError) {
	data, err := c.client.Patch(fmt.Sprintf("/contactBooks/%s/contacts/%s", bookID, contactID), payload)
	if err != nil {
		return nil, err
	}

	result := &ContactUpdateResponse{}
	if dataMap, ok := data.(map[string]interface{}); ok {
		result.ID = dataMap["id"].(string)
	}

	return result, nil
}

// Upsert creates or updates a contact
func (c *ContactsClient) Upsert(bookID, contactID string, payload ContactUpsert) (*ContactUpsertResponse, *APIError) {
	data, err := c.client.Put(fmt.Sprintf("/contactBooks/%s/contacts/%s", bookID, contactID), payload)
	if err != nil {
		return nil, err
	}

	result := &ContactUpsertResponse{}
	if dataMap, ok := data.(map[string]interface{}); ok {
		result.ID = dataMap["id"].(string)
		result.Email = dataMap["email"].(string)
	}

	return result, nil
}

// Delete deletes a contact
func (c *ContactsClient) Delete(bookID, contactID string) (*ContactDeleteResponse, *APIError) {
	data, err := c.client.Delete(fmt.Sprintf("/contactBooks/%s/contacts/%s", bookID, contactID), nil)
	if err != nil {
		return nil, err
	}

	result := &ContactDeleteResponse{}
	if dataMap, ok := data.(map[string]interface{}); ok {
		result.ID = dataMap["id"].(string)
		if deleted, ok := dataMap["deleted"].(bool); ok {
			result.Deleted = deleted
		}
	}

	return result, nil
}
