package unsent

type SettingsClient struct {
	client *Client
}

// Get retrieves team settings
func (c *SettingsClient) Get() (*Settings, *APIError) {
	return Get[Settings](c.client, "/settings")
}
