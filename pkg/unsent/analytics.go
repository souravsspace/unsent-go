package unsent

import "fmt"

type AnalyticsClient struct {
	client *Client
}

// Get retrieves email analytics
func (c *AnalyticsClient) Get() (*Analytics, *APIError) {
	return Get[Analytics](c.client, "/analytics")
}

// GetTimeSeries retrieves analytics data over time
func (a *AnalyticsClient) GetTimeSeries(params GetTimeSeriesParams) (*GetTimelineResponse, *APIError) {
	path := "/analytics/time-series"
	
	// Build query parameters
	query := buildQueryParams(map[string]interface{}{
		"days":   params.Days,
		"domain": params.Domain,
	})
	
	if query != "" {
		path = fmt.Sprintf("%s?%s", path, query)
	}
	
	return Get[GetTimelineResponse](a.client, path)
}

// GetReputation retrieves sender reputation score
func (c *AnalyticsClient) GetReputation(params GetReputationParams) (*AnalyticsReputation, *APIError) {
	path := "/analytics/reputation?"
	if params.Domain != nil {
		path += fmt.Sprintf("domain=%s&", *params.Domain)
	}
	return Get[AnalyticsReputation](c.client, path)
}
