package unsent

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestWebhooks_List(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "GET" {
			t.Errorf("expected GET, got %s", r.Method)
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`[{"id": "wh1", "url": "https://hook.site"}]`))
	}))
	defer server.Close()
	client, _ := NewClient("key", WithBaseURL(server.URL))
	resp, err := client.Webhooks.List()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(*resp) != 1 {
		t.Errorf("expected 1 webhook, got %d", len(*resp))
	}
}

func TestWebhooks_Create(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "POST" {
			t.Errorf("expected POST, got %s", r.Method)
		}
		var body CreateWebhookJSONBody
		json.NewDecoder(r.Body).Decode(&body)
		if body.Url != "https://hook.site" {
			t.Errorf("expected https://hook.site, got %s", body.Url)
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"id": "wh1"}`))
	}))
	defer server.Close()

	client, _ := NewClient("key", WithBaseURL(server.URL))
	
	resp, err := client.Webhooks.Create(CreateWebhookJSONBody{
		Url: "https://hook.site",
		EventTypes: []CreateWebhookJSONBodyEventTypes{"email.sent"},
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if resp.ID != "wh1" {
		t.Errorf("expected wh1, got %s", resp.ID)
	}
}

func TestWebhooks_Update(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "PATCH" {
			t.Errorf("expected PATCH, got %s", r.Method)
		}
		if r.URL.Path != "/v1/webhooks/wh1" {
			t.Errorf("expected /v1/webhooks/wh1, got %s", r.URL.Path)
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"success": true}`))
	}))
	defer server.Close()
	client, _ := NewClient("key", WithBaseURL(server.URL))
	url := "https://new.site"
	_, err := client.Webhooks.Update("wh1", UpdateWebhookJSONBody{
		Url: &url,
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
}

func TestWebhooks_Delete(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "DELETE" {
			t.Errorf("expected DELETE, got %s", r.Method)
		}
		if r.URL.Path != "/v1/webhooks/wh1" {
			t.Errorf("expected /v1/webhooks/wh1, got %s", r.URL.Path)
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"success": true}`))
	}))
	defer server.Close()
	client, _ := NewClient("key", WithBaseURL(server.URL))
	_, err := client.Webhooks.Delete("wh1")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
}

func TestWebhooks_Test(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "POST" {
			t.Errorf("expected POST, got %s", r.Method)
		}
		if r.URL.Path != "/v1/webhooks/wh1/test" {
			t.Errorf("expected /v1/webhooks/wh1/test, got %s", r.URL.Path)
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{
			"id": "call_123",
			"type": "contact.created",
			"createdAt": "2024-01-01T00:00:00Z",
			"updatedAt": "2024-01-01T00:00:00Z",
			"teamId": "team_123",
			"status": "success",
			"webhookId": "wh1",
			"payload": "{}",
			"attempt": 1,
			"responseStatus": 200,
			"responseTimeMs": 100
		}`))
	}))
	defer server.Close()

	client, _ := NewClient("key", WithBaseURL(server.URL))
	resp, err := client.Webhooks.Test("wh1")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if resp.ID != "call_123" {
		t.Errorf("expected call_123, got %s", resp.ID)
	}
	if resp.Attempt != 1 {
		t.Errorf("expected 1, got %f", resp.Attempt)
	}
}
