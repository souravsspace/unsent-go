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
func (c *AnalyticsClient) GetTimeSeries(params GetTimeSeriesParams) (*[]AnalyticsTimeSeries, *APIError) {
	// Query params handling needs to be implemented in Request or helper.
	// But `Get` helper currently doesn't support query params easily unless passed in path or opts.
	// `types.go` defines `GetTimeSeriesParams` struct with `form` tags.
	// I need to serialize struct to query string.
	// Since `request` function in `unsent.go` doesn't handle query params automatically from a struct, 
	// I might need to append them to the path string manually or update `unsent.go`.
	// For now, I will construct the path manually or assume `unsent` client is updated to handle query params.
	// Wait, `Get` takes `opts ...RequestOption`. One option could be `WithQueryParams`.
	// But `types.go` params are structs.
	
	// Let's defer strict query param handling or do basic manual construction if simple.
	// "days" and "domain".
	
	path := "/analytics/time-series?"
	if params.Days != nil {
		path += fmt.Sprintf("days=%s&", *params.Days)
	}
	if params.Domain != nil {
		path += fmt.Sprintf("domain=%s&", *params.Domain)
	}
	return Get[[]AnalyticsTimeSeries](c.client, path)
}

// GetReputation retrieves sender reputation score
func (c *AnalyticsClient) GetReputation(params GetReputationParams) (*AnalyticsReputation, *APIError) {
	path := "/analytics/reputation?"
	if params.Domain != nil {
		path += fmt.Sprintf("domain=%s&", *params.Domain)
	}
	return Get[AnalyticsReputation](c.client, path)
}
