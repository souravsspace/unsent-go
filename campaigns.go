package unsent

import "fmt"

// CampaignsClient handles campaign-related API operations
type CampaignsClient struct {
	client *Client
}

// Create creates a new campaign
func (c *CampaignsClient) Create(payload CampaignCreate) (*CampaignCreateResponse, *APIError) {
	data, err := c.client.Post("/campaigns", payload)
	if err != nil {
		return nil, err
	}

	result := &CampaignCreateResponse{}
	if dataMap, ok := data.(map[string]interface{}); ok {
		result.ID = dataMap["id"].(string)
		result.Name = dataMap["name"].(string)
		result.Status = dataMap["status"].(string)
	}

	return result, nil
}

// Get retrieves a campaign by ID
func (c *CampaignsClient) Get(campaignID string) (*Campaign, *APIError) {
	data, err := c.client.Get(fmt.Sprintf("/campaigns/%s", campaignID))
	if err != nil {
		return nil, err
	}

	result := &Campaign{}
	if dataMap, ok := data.(map[string]interface{}); ok {
		result.ID = dataMap["id"].(string)
		result.Name = dataMap["name"].(string)
		result.Subject = dataMap["subject"].(string)
		result.HTML = dataMap["html"].(string)
		result.From = dataMap["from"].(string)
		result.ContactBookID = dataMap["contactBookId"].(string)
		result.Status = dataMap["status"].(string)
		if total, ok := dataMap["total"].(float64); ok {
			result.Total = int(total)
		}
		if sent, ok := dataMap["sent"].(float64); ok {
			result.Sent = int(sent)
		}
		if failed, ok := dataMap["failed"].(float64); ok {
			result.Failed = int(failed)
		}
	}

	return result, nil
}

// Schedule schedules a campaign
func (c *CampaignsClient) Schedule(campaignID string, payload CampaignSchedule) (*CampaignScheduleResponse, *APIError) {
	data, err := c.client.Post(fmt.Sprintf("/campaigns/%s/schedule", campaignID), payload)
	if err != nil {
		return nil, err
	}

	result := &CampaignScheduleResponse{}
	if dataMap, ok := data.(map[string]interface{}); ok {
		result.ID = dataMap["id"].(string)
		result.Status = dataMap["status"].(string)
	}

	return result, nil
}

// Pause pauses a campaign
func (c *CampaignsClient) Pause(campaignID string) (*CampaignActionResponse, *APIError) {
	data, err := c.client.Post(fmt.Sprintf("/campaigns/%s/pause", campaignID), map[string]interface{}{})
	if err != nil {
		return nil, err
	}

	result := &CampaignActionResponse{}
	if dataMap, ok := data.(map[string]interface{}); ok {
		result.ID = dataMap["id"].(string)
		result.Status = dataMap["status"].(string)
	}

	return result, nil
}

// Resume resumes a campaign
func (c *CampaignsClient) Resume(campaignID string) (*CampaignActionResponse, *APIError) {
	data, err := c.client.Post(fmt.Sprintf("/campaigns/%s/resume", campaignID), map[string]interface{}{})
	if err != nil {
		return nil, err
	}

	result := &CampaignActionResponse{}
	if dataMap, ok := data.(map[string]interface{}); ok {
		result.ID = dataMap["id"].(string)
		result.Status = dataMap["status"].(string)
	}

	return result, nil
}
