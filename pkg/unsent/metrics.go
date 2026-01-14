package unsent

import "fmt"

// MetricsClient handles metrics API endpoints
type MetricsClient struct {
	client *Client
}

// Get retrieves performance metrics
func (m *MetricsClient) Get(params GetMetricsParams) (*MetricsResponse, *APIError) {
	path := "/metrics"
	
	// Build query parameters
	query := buildQueryParams(map[string]interface{}{
		"period": params.Period,
	})
	
	if query != "" {
		path = fmt.Sprintf("%s?%s", path, query)
	}
	
	return Get[MetricsResponse](m.client, path)
}
