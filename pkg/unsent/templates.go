package unsent

import "fmt"

type TemplatesClient struct {
	client *Client
}

// List retrieves all templates
func (c *TemplatesClient) List() (*[]Template, *APIError) {
	return Get[[]Template](c.client, "/templates")
}

// Get retrieves a template by ID
func (c *TemplatesClient) Get(id string) (*Template, *APIError) {
	return Get[Template](c.client, fmt.Sprintf("/templates/%s", id))
}

// Create creates a new template
func (c *TemplatesClient) Create(payload CreateTemplateJSONBody) (*TemplateCreateResponse, *APIError) {
	return Post[TemplateCreateResponse](c.client, "/templates", payload)
}

// Update updates a template
func (c *TemplatesClient) Update(id string, payload UpdateTemplateJSONBody) (*TemplateUpdateResponse, *APIError) {
	return Patch[TemplateUpdateResponse](c.client, fmt.Sprintf("/templates/%s", id), payload)
}

// Delete deletes a template
func (c *TemplatesClient) Delete(id string) (*TemplateDeleteResponse, *APIError) {
	return Delete[TemplateDeleteResponse](c.client, fmt.Sprintf("/templates/%s", id), nil)
}
