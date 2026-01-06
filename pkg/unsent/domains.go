package unsent

import "fmt"

// DomainsClient handles domain-related API operations
type DomainsClient struct {
	client *Client
}

// List retrieves all domains
func (c *DomainsClient) List() (*[]Domain, *APIError) {
	return Get[[]Domain](c.client, "/domains")
}

// Get retrieves a domain by ID
func (c *DomainsClient) Get(domainID string) (*Domain, *APIError) {
	return Get[Domain](c.client, fmt.Sprintf("/domains/%s", domainID))
}

// Create creates a new domain
func (c *DomainsClient) Create(payload CreateDomainJSONBody) (*DomainCreateResponse, *APIError) {
	return Post[DomainCreateResponse](c.client, "/domains", payload)
}

// Verify triggers domain verification
func (c *DomainsClient) Verify(domainID string) (*DomainVerifyResponse, *APIError) {
	return Put[DomainVerifyResponse](c.client, fmt.Sprintf("/domains/%s/verify", domainID), nil)
}

// Delete deletes a domain
func (c *DomainsClient) Delete(domainID string) (*DomainDeleteResponse, *APIError) {
	return Delete[DomainDeleteResponse](c.client, fmt.Sprintf("/domains/%s", domainID), nil)
}
