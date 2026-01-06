package unsent

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

// NOTE: Webhook tests are placeholders for future functionality.
// These tests verify the SDK's webhook client structure and methods,
// but the actual webhook API endpoints are not yet available.
// These tests will be functional once webhook support is added to the API.

func TestWebhooks_Create(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "POST" {
			t.Errorf("expected POST, got %s", r.Method)
		}
		var body WebhookCreateRequest
		json.NewDecoder(r.Body).Decode(&body)
		if body.Url != "https://hook.site" {
			t.Errorf("expected https://hook.site, got %s", body.Url)
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"id": "wh1"}`))
	}))
	defer server.Close()

	client, _ := NewClient("key", WithBaseURL(server.URL))
	
	resp, err := client.Webhooks.Create(WebhookCreateRequest{
		Url: "https://hook.site",
		Events: []string{"email.sent"},
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if resp.ID != "wh1" {
		t.Errorf("expected wh1, got %s", resp.ID)
	}
}
