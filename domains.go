package unsent

import "fmt"

// DomainsClient handles domain-related API operations
type DomainsClient struct {
	client *Client
}

// List retrieves all domains
func (d *DomainsClient) List() ([]Domain, *APIError) {
	// Note: The API returns a list of domains directly, not wrapped in an object
	// We need to handle this case. Our generic Get returns *T.
	// So we need *[]Domain
	result, err := Get[[]Domain](d.client, "/domains")
	if err != nil {
		return nil, err
	}
	return *result, nil
}

// Create creates a new domain
func (d *DomainsClient) Create(payload DomainCreate) (*DomainCreateResponse, *APIError) {
	return Post[DomainCreateResponse](d.client, "/domains", payload)
}

// Verify verifies a domain
func (d *DomainsClient) Verify(domainID int) (*DomainVerifyResponse, *APIError) {
	return Put[DomainVerifyResponse](d.client, fmt.Sprintf("/domains/%d/verify", domainID), map[string]interface{}{})
}

// Get retrieves a domain by ID
func (d *DomainsClient) Get(domainID int) (*Domain, *APIError) {
	return Get[Domain](d.client, fmt.Sprintf("/domains/%d", domainID))
}

// Delete deletes a domain
func (d *DomainsClient) Delete(domainID int) (*DomainDeleteResponse, *APIError) {
	return Delete[DomainDeleteResponse](d.client, fmt.Sprintf("/domains/%d", domainID), nil)
}
