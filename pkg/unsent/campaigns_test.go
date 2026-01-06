package unsent

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCampaigns_List(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "GET" {
			t.Errorf("expected GET, got %s", r.Method)
		}
		if r.URL.Path != "/v1/campaigns" {
			t.Errorf("expected path /v1/campaigns, got %s", r.URL.Path)
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`[{"id": "camp1", "name": "Camp 1", "status": "DRAFT"}]`))
	}))
	defer server.Close()

	client, _ := NewClient("key", WithBaseURL(server.URL))

	resp, err := client.Campaigns.List()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(*resp) != 1 {
		t.Errorf("expected 1 campaign, got %d", len(*resp))
	}
	if (*resp)[0].ID != "camp1" {
		t.Errorf("expected camp1, got %s", (*resp)[0].ID)
	}
}

func TestCampaigns_Create(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "POST" {
			t.Errorf("expected POST, got %s", r.Method)
		}
		var body CreateCampaignJSONBody
		json.NewDecoder(r.Body).Decode(&body)
		if body.Name != "Camp 1" {
			t.Errorf("expected Camp 1, got %s", body.Name)
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"id": "camp1", "name": "Camp 1", "status": "DRAFT"}`))
	}))
	defer server.Close()

	client, _ := NewClient("key", WithBaseURL(server.URL))
	
	resp, err := client.Campaigns.Create(CreateCampaignJSONBody{
		Name: "Camp 1",
		Subject: "Hello",
		From: "me@test.com",
		ContactBookId: "book1",
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if resp.ID != "camp1" {
		t.Errorf("expected camp1, got %s", resp.ID)
	}
}
