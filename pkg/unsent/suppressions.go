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
	resp, err := Get[GetSuppressionsResponse](c.client, path)
	if err != nil {
		return nil, err
	}
	return &resp.Data, nil
}

// Add adds a suppression
func (c *SuppressionsClient) Add(payload AddSuppressionJSONBody) (*SuppressionAddResponse, *APIError) {
	return Post[SuppressionAddResponse](c.client, "/suppressions", payload)
}

// Delete deletes a suppression
// Note: API Ref might use email in path or body. TS SDK uses DELETE /suppressions with body { email } or path param?
// TS SDK: `this.unsent.delete<{ deleted: boolean }>("/suppressions", { email })`
// So it uses a body for DELETE.
// Delete removes an email from the suppression list
func (c *SuppressionsClient) Delete(email string) (*SuppressionDeleteResponse, *APIError) {
	return Delete[SuppressionDeleteResponse](c.client, fmt.Sprintf("/suppressions/email/%s", email), nil)
}
