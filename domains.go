package unsent

import "fmt"

// DomainsClient handles domain-related API operations
type DomainsClient struct {
	client *Client
}

// List retrieves all domains
func (d *DomainsClient) List() ([]Domain, *APIError) {
	data, err := d.client.Get("/domains")
	if err != nil {
		return nil, err
	}

	var result []Domain
	if dataArray, ok := data.([]interface{}); ok {
		for _, item := range dataArray {
			if domainMap, ok := item.(map[string]interface{}); ok {
				domain := Domain{
					Domain: domainMap["domain"].(string),
					Status: domainMap["status"].(string),
				}
				if id, ok := domainMap["id"].(float64); ok {
					domain.ID = int(id)
				}
				result = append(result, domain)
			}
		}
	}

	return result, nil
}

// Create creates a new domain
func (d *DomainsClient) Create(payload DomainCreate) (*DomainCreateResponse, *APIError) {
	data, err := d.client.Post("/domains", payload)
	if err != nil {
		return nil, err
	}

	result := &DomainCreateResponse{}
	if dataMap, ok := data.(map[string]interface{}); ok {
		if id, ok := dataMap["id"].(float64); ok {
			result.ID = int(id)
		}
		result.Domain = dataMap["domain"].(string)
		result.Status = dataMap["status"].(string)
	}

	return result, nil
}

// Verify verifies a domain
func (d *DomainsClient) Verify(domainID int) (*DomainVerifyResponse, *APIError) {
	data, err := d.client.Put(fmt.Sprintf("/domains/%d/verify", domainID), map[string]interface{}{})
	if err != nil {
		return nil, err
	}

	result := &DomainVerifyResponse{}
	if dataMap, ok := data.(map[string]interface{}); ok {
		if id, ok := dataMap["id"].(float64); ok {
			result.ID = int(id)
		}
		result.Status = dataMap["status"].(string)
	}

	return result, nil
}

// Get retrieves a domain by ID
func (d *DomainsClient) Get(domainID int) (*Domain, *APIError) {
	data, err := d.client.Get(fmt.Sprintf("/domains/%d", domainID))
	if err != nil {
		return nil, err
	}

	result := &Domain{}
	if dataMap, ok := data.(map[string]interface{}); ok {
		if id, ok := dataMap["id"].(float64); ok {
			result.ID = int(id)
		}
		result.Domain = dataMap["domain"].(string)
		result.Status = dataMap["status"].(string)
	}

	return result, nil
}

// Delete deletes a domain
func (d *DomainsClient) Delete(domainID int) (*DomainDeleteResponse, *APIError) {
	data, err := d.client.Delete(fmt.Sprintf("/domains/%d", domainID), nil)
	if err != nil {
		return nil, err
	}

	result := &DomainDeleteResponse{}
	if dataMap, ok := data.(map[string]interface{}); ok {
		if id, ok := dataMap["id"].(float64); ok {
			result.ID = int(id)
		}
		if deleted, ok := dataMap["deleted"].(bool); ok {
			result.Deleted = deleted
		}
	}

	return result, nil
}
