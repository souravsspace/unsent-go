package unsent

import "fmt"

// StatsClient handles statistics API endpoints
type StatsClient struct {
	client *Client
}

// Get retrieves email statistics
func (s *StatsClient) Get(params GetStatsParams) (*StatsResponse, *APIError) {
	path := "/stats"
	
	// Build query parameters
	query := buildQueryParams(map[string]interface{}{
		"startDate": params.StartDate,
		"endDate":   params.EndDate,
	})
	
	if query != "" {
		path = fmt.Sprintf("%s?%s", path, query)
	}
	
	return Get[StatsResponse](s.client, path)
}
