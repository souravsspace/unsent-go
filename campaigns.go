package unsent

import "fmt"

// CampaignsClient handles campaign-related API operations
type CampaignsClient struct {
	client *Client
}

// Create creates a new campaign
func (c *CampaignsClient) Create(payload CampaignCreate) (*CampaignCreateResponse, *APIError) {
	return Post[CampaignCreateResponse](c.client, "/campaigns", payload)
}

// Get retrieves a campaign by ID
func (c *CampaignsClient) Get(campaignID string) (*Campaign, *APIError) {
	return Get[Campaign](c.client, fmt.Sprintf("/campaigns/%s", campaignID))
}

// Schedule schedules a campaign
func (c *CampaignsClient) Schedule(campaignID string, payload CampaignSchedule) (*CampaignScheduleResponse, *APIError) {
	return Post[CampaignScheduleResponse](c.client, fmt.Sprintf("/campaigns/%s/schedule", campaignID), payload)
}

// Pause pauses a campaign
func (c *CampaignsClient) Pause(campaignID string) (*CampaignActionResponse, *APIError) {
	return Post[CampaignActionResponse](c.client, fmt.Sprintf("/campaigns/%s/pause", campaignID), map[string]interface{}{})
}

// Resume resumes a campaign
func (c *CampaignsClient) Resume(campaignID string) (*CampaignActionResponse, *APIError) {
	return Post[CampaignActionResponse](c.client, fmt.Sprintf("/campaigns/%s/resume", campaignID), map[string]interface{}{})
}
