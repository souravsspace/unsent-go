package unsent

import "fmt"

// ContactsClient handles contact-related API operations
type ContactsClient struct {
	client *Client
}

// List retrieves contacts from a contact book
func (c *ContactsClient) List(bookID string, params GetContactsParams) (*[]Contact, *APIError) {
	path := fmt.Sprintf("/contactBooks/%s/contacts?", bookID)
	if params.Emails != nil {
		path += fmt.Sprintf("emails=%s&", *params.Emails)
	}
	if params.Page != nil {
		path += fmt.Sprintf("page=%f&", *params.Page)
	}
	if params.Limit != nil {
		path += fmt.Sprintf("limit=%f&", *params.Limit)
	}
	if params.Ids != nil {
		path += fmt.Sprintf("ids=%s&", *params.Ids)
	}
	return Get[[]Contact](c.client, path)
}

// Create creates a new contact
func (c *ContactsClient) Create(bookID string, payload CreateContactJSONBody) (*ContactCreateResponse, *APIError) {
	return Post[ContactCreateResponse](c.client, fmt.Sprintf("/contactBooks/%s/contacts", bookID), payload)
}

// Get retrieves a contact by ID
func (c *ContactsClient) Get(bookID, contactID string) (*Contact, *APIError) {
	return Get[Contact](c.client, fmt.Sprintf("/contactBooks/%s/contacts/%s", bookID, contactID))
}

// Update updates a contact
func (c *ContactsClient) Update(bookID, contactID string, payload UpdateContactJSONBody) (*ContactUpdateResponse, *APIError) {
	return Patch[ContactUpdateResponse](c.client, fmt.Sprintf("/contactBooks/%s/contacts/%s", bookID, contactID), payload)
}

// Upsert creates or updates a contact
func (c *ContactsClient) Upsert(bookID, contactID string, payload UpsertContactJSONBody) (*ContactUpsertResponse, *APIError) {
	return Put[ContactUpsertResponse](c.client, fmt.Sprintf("/contactBooks/%s/contacts/%s", bookID, contactID), payload)
}

// Delete deletes a contact
func (c *ContactsClient) Delete(bookID, contactID string) (*ContactDeleteResponse, *APIError) {
	return Delete[ContactDeleteResponse](c.client, fmt.Sprintf("/contactBooks/%s/contacts/%s", bookID, contactID), nil)
}
