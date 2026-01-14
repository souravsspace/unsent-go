package unsent

// ListEmailsResponse represents the response for listing emails
type ListEmailsResponse struct {
	Data  []Email `json:"data"`
	Count int     `json:"count"`
}

// GetTimelineResponse represents the response for analytics timeline
type GetTimelineResponse struct {
	Data []AnalyticsTimeSeries `json:"data"`
}

// GetComplaintsResponse represents the response for getting complaints
type GetComplaintsResponse struct {
	Data  []Email `json:"data"`
	Count int     `json:"count"`
}

// GetBouncesResponse represents the response for getting bounces
type GetBouncesResponse struct {
	Data  []Email `json:"data"`
	Count int     `json:"count"`
}

// GetUnsubscribesResponse represents the response for getting unsubscribes
type GetUnsubscribesResponse struct {
	Data  []Email `json:"data"`
	Count int     `json:"count"`
}

// GetEmailEventsResponse represents the response for getting email events
type GetEmailEventsResponse struct {
	Data []map[string]interface{} `json:"data"`
}

// GetEventsResponse represents the response for getting system events
type GetEventsResponse struct {
	Data  []map[string]interface{} `json:"data"`
	Count int                      `json:"count"`
}

// GetActivityResponse represents the response for getting activity
type GetActivityResponse struct {
	Data  []map[string]interface{} `json:"data"`
	Count int                      `json:"count"`
}

// GetTeamsResponse represents the response for listing teams
type GetTeamsResponse struct {
	Data []map[string]interface{} `json:"data"`
}

// GetDomainsResponse represents the response for listing domains
type GetDomainsResponse struct {
	Data []Domain `json:"data"`
}

// GetContactsResponse represents the response for listing contacts
type GetContactsResponse struct {
	Data  []Contact `json:"data"`
	Count int       `json:"count"`
}

// GetTemplatesResponse represents the response for listing templates
type GetTemplatesResponse struct {
	Data []Template `json:"data"`
}

// GetContactBooksResponse represents the response for listing contact books
type GetContactBooksResponse struct {
	Data []ContactBook `json:"data"`
}

// GetApiKeysResponse represents the response for listing API keys
type GetApiKeysResponse struct {
	Data []ApiKey `json:"data"`
}

// GetSuppressionsResponse represents the response for listing suppressions
type GetSuppressionsResponse struct {
	Data []Suppression `json:"data"`
}

// GetWebhooksResponse represents the response for listing webhooks
type GetWebhooksResponse struct {
	Data []Webhook `json:"data"`
}
