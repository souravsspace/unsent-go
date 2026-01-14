package unsent

import "fmt"

type ContactBooksClient struct {
	client *Client
}

// List retrieves all contact books
func (c *ContactBooksClient) List() (*[]ContactBook, *APIError) {
	return Get[[]ContactBook](c.client, "/contactBooks")
}

// Get retrieves a contact book by ID
func (c *ContactBooksClient) Get(id string) (*ContactBook, *APIError) {
	return Get[ContactBook](c.client, fmt.Sprintf("/contactBooks/%s", id))
}

// Create creates a new contact book
func (c *ContactBooksClient) Create(payload CreateContactBookJSONBody) (*ContactBookCreateResponse, *APIError) {
	return Post[ContactBookCreateResponse](c.client, "/contactBooks", payload)
}

// Update updates a contact book
func (c *ContactBooksClient) Update(id string, payload UpdateContactBookJSONBody) (*ContactBookUpdateResponse, *APIError) {
	return Patch[ContactBookUpdateResponse](c.client, fmt.Sprintf("/contactBooks/%s", id), payload)
}

// Delete deletes a contact book
func (c *ContactBooksClient) Delete(id string) (*ContactBookDeleteResponse, *APIError) {
	return Delete[ContactBookDeleteResponse](c.client, fmt.Sprintf("/contactBooks/%s", id), nil)
}
