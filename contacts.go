package unsent

import "fmt"

// ContactsClient handles contact-related API operations
type ContactsClient struct {
	client *Client
}

// Create creates a new contact
func (c *ContactsClient) Create(bookID string, payload ContactCreate) (*ContactCreateResponse, *APIError) {
	return Post[ContactCreateResponse](c.client, fmt.Sprintf("/contactBooks/%s/contacts", bookID), payload)
}

// Get retrieves a contact by ID
func (c *ContactsClient) Get(bookID, contactID string) (*Contact, *APIError) {
	return Get[Contact](c.client, fmt.Sprintf("/contactBooks/%s/contacts/%s", bookID, contactID))
}

// Update updates a contact
func (c *ContactsClient) Update(bookID, contactID string, payload ContactUpdate) (*ContactUpdateResponse, *APIError) {
	return Patch[ContactUpdateResponse](c.client, fmt.Sprintf("/contactBooks/%s/contacts/%s", bookID, contactID), payload)
}

// Upsert creates or updates a contact
func (c *ContactsClient) Upsert(bookID, contactID string, payload ContactUpsert) (*ContactUpsertResponse, *APIError) {
	return Put[ContactUpsertResponse](c.client, fmt.Sprintf("/contactBooks/%s/contacts/%s", bookID, contactID), payload)
}

// Delete deletes a contact
func (c *ContactsClient) Delete(bookID, contactID string) (*ContactDeleteResponse, *APIError) {
	return Delete[ContactDeleteResponse](c.client, fmt.Sprintf("/contactBooks/%s/contacts/%s", bookID, contactID), nil)
}
