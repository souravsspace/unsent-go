package unsent

// SystemClient handles system-level API endpoints
type SystemClient struct {
	client *Client
}

// Health checks if the API is running correctly
func (s *SystemClient) Health() (*HealthResponse, *APIError) {
	return Get[HealthResponse](s.client, "/health")
}

// Version retrieves API version information
func (s *SystemClient) Version() (*VersionResponse, *APIError) {
	return Get[VersionResponse](s.client, "/version")
}
