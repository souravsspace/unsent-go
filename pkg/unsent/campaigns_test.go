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

func TestCampaigns_Get(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "GET" {
			t.Errorf("expected GET, got %s", r.Method)
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"id": "camp1", "name": "Camp 1"}`))
	}))
	defer server.Close()
	client, _ := NewClient("key", WithBaseURL(server.URL))
	resp, err := client.Campaigns.Get("camp1")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if resp.ID != "camp1" {
		t.Errorf("expected camp1, got %s", resp.ID)
	}
}

func TestCampaigns_Schedule(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "POST" {
			t.Errorf("expected POST, got %s", r.Method)
		}
		if r.URL.Path != "/v1/campaigns/camp1/schedule" {
			t.Errorf("expected /v1/campaigns/camp1/schedule, got %s", r.URL.Path)
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"id": "camp1", "status": "SCHEDULED"}`))
	}))
	defer server.Close()
	client, _ := NewClient("key", WithBaseURL(server.URL))
	resp, err := client.Campaigns.Schedule("camp1", ScheduleCampaignJSONBody{})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if resp.ID != "camp1" {
		t.Errorf("expected camp1, got %s", resp.ID)
	}
}

func TestCampaigns_Pause(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "POST" {
			t.Errorf("expected POST, got %s", r.Method)
		}
		if r.URL.Path != "/v1/campaigns/camp1/pause" {
			t.Errorf("expected /v1/campaigns/camp1/pause, got %s", r.URL.Path)
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"id": "camp1", "status": "PAUSED"}`))
	}))
	defer server.Close()
	client, _ := NewClient("key", WithBaseURL(server.URL))
	resp, err := client.Campaigns.Pause("camp1")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if resp.ID != "camp1" {
		t.Errorf("expected camp1, got %s", resp.ID)
	}
}

func TestCampaigns_Resume(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "POST" {
			t.Errorf("expected POST, got %s", r.Method)
		}
		if r.URL.Path != "/v1/campaigns/camp1/resume" {
			t.Errorf("expected /v1/campaigns/camp1/resume, got %s", r.URL.Path)
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"id": "camp1", "status": "ACTIVE"}`))
	}))
	defer server.Close()
	client, _ := NewClient("key", WithBaseURL(server.URL))
	resp, err := client.Campaigns.Resume("camp1")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if resp.ID != "camp1" {
		t.Errorf("expected camp1, got %s", resp.ID)
	}
}
