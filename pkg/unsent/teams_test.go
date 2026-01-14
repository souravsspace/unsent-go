package unsent

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestTeamsGet(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/v1/team" {
			t.Errorf("Expected path /v1/team, got %s", r.URL.Path)
		}
		if r.Method != "GET" {
			t.Errorf("Expected GET method, got %s", r.Method)
		}

		response := map[string]interface{}{
			"id":   "team_123",
			"name": "Test Team",
		}
		json.NewEncoder(w).Encode(response)
	}))
	defer server.Close()

	client, _ := NewClient("test_key", WithBaseURL(server.URL))
	
	team, err := client.Teams.Get()

	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	if team == nil {
		t.Error("Expected team response, got nil")
	}
}

func TestTeamsList(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/v1/teams" {
			t.Errorf("Expected path /v1/teams, got %s", r.URL.Path)
		}
		if r.Method != "GET" {
			t.Errorf("Expected GET method, got %s", r.Method)
		}

		response := []interface{}{
			map[string]interface{}{
				"id":   "team_123",
				"name": "Test Team",
			},
		}
		json.NewEncoder(w).Encode(response)
	}))
	defer server.Close()

	client, _ := NewClient("test_key", WithBaseURL(server.URL))
	
	resp, err := client.Teams.List()

	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	if resp == nil {
		t.Error("Expected teams response, got nil")
	}
	if len(*resp) != 1 {
		t.Errorf("Expected 1 team, got %d", len(*resp))
	}
}
