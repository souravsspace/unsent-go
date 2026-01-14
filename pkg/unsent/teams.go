package unsent

// TeamsClient handles team API endpoints
type TeamsClient struct {
	client *Client
}

// Get retrieves the current team information
func (t *TeamsClient) Get() (*Team, *APIError) {
	return Get[Team](t.client, "/team")
}

// List retrieves all teams
func (t *TeamsClient) List() (*[]Team, *APIError) {
	return Get[[]Team](t.client, "/teams")
}

