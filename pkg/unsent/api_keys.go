package unsent

import "fmt"

type ApiKeysClient struct {
	client *Client
}

// List retrieves all API keys
func (c *ApiKeysClient) List() (*[]ApiKey, *APIError) {
	return Get[[]ApiKey](c.client, "/api-keys")
}

// Create creates a new API key
func (c *ApiKeysClient) Create(payload CreateApiKeyJSONBody) (*ApiKeyCreateResponse, *APIError) {
	return Post[ApiKeyCreateResponse](c.client, "/api-keys", payload)
}

// Delete deletes an API key
func (c *ApiKeysClient) Delete(id string) (*ApiKeyDeleteResponse, *APIError) {
	return Delete[ApiKeyDeleteResponse](c.client, fmt.Sprintf("/api-keys/%s", id), nil)
}
