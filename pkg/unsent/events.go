package unsent

import "fmt"

// EventsClient handles email events API endpoints
type EventsClient struct {
	client *Client
}

// List retrieves all email events with optional filtering
func (e *EventsClient) List(params GetEventsParams) (*EventsListResponse, *APIError) {
	path := "/events"
	
	// Build query parameters
	query := buildQueryParams(map[string]interface{}{
		"page":      params.Page,
		"limit":     params.Limit,
		"status":    params.Status,
		"startDate": params.StartDate,
	})
	
	if query != "" {
		path = fmt.Sprintf("%s?%s", path, query)
	}
	
	return Get[EventsListResponse](e.client, path)
}
