package unsent

import "fmt"

// ActivityClient handles activity feed API endpoints
type ActivityClient struct {
	client *Client
}

// Get retrieves the activity feed with email events and email details
func (a *ActivityClient) Get(params GetActivityParams) (*ActivityResponse, *APIError) {
	path := "/activity"
	
	// Build query parameters
	query := buildQueryParams(map[string]interface{}{
		"page":  params.Page,
		"limit": params.Limit,
	})
	
	if query != "" {
		path = fmt.Sprintf("%s?%s", path, query)
	}
	
	return Get[ActivityResponse](a.client, path)
}
