package unsent

import "fmt"

type SuppressionsClient struct {
	client *Client
}

// List retrieves all suppressions
func (c *SuppressionsClient) List(params GetSuppressionsParams) (*[]Suppression, *APIError) {
	path := "/suppressions?"
	if params.Page != nil {
		path += fmt.Sprintf("page=%f&", *params.Page)
	}
	if params.Limit != nil {
		path += fmt.Sprintf("limit=%f&", *params.Limit)
	}
	if params.Search != nil {
		path += fmt.Sprintf("search=%s&", *params.Search)
	}
	if params.Reason != nil {
		path += fmt.Sprintf("reason=%s&", *params.Reason)
	}
	return Get[[]Suppression](c.client, path)
}

// Add adds a suppression
func (c *SuppressionsClient) Add(payload AddSuppressionJSONBody) (*SuppressionAddResponse, *APIError) {
	return Post[SuppressionAddResponse](c.client, "/suppressions", payload)
}

// Delete deletes a suppression
// Note: API Ref might use email in path or body. TS SDK uses DELETE /suppressions with body { email } or path param?
// TS SDK: `this.unsent.delete<{ deleted: boolean }>("/suppressions", { email })`
// So it uses a body for DELETE.
func (c *SuppressionsClient) Delete(email string) (*SuppressionDeleteResponse, *APIError) {
	// Using map as body
	return Delete[SuppressionDeleteResponse](c.client, "/suppressions", map[string]string{"email": email})
}
