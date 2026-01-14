package unsent

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestEventsList(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/v1/events" {
			t.Errorf("Expected path /v1/events, got %s", r.URL.Path)
		}
		if r.Method != "GET" {
			t.Errorf("Expected GET method, got %s", r.Method)
		}

		response := map[string]interface{}{
			"events": []interface{}{},
			"total":  0,
		}
		json.NewEncoder(w).Encode(response)
	}))
	defer server.Close()

	client, _ := NewClient("test_key", WithBaseURL(server.URL))
	
	page := 1
	limit := 10
	params := GetEventsParams{
		Page:  &page,
		Limit: &limit,
	}
	
	events, err := client.Events.List(params)

	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	if events == nil {
		t.Error("Expected events response, got nil")
	}
}
