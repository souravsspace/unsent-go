package unsent

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestActivityGet(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/v1/activity" {
			t.Errorf("Expected path /v1/activity, got %s", r.URL.Path)
		}
		if r.Method != "GET" {
			t.Errorf("Expected GET method, got %s", r.Method)
		}

		response := map[string]interface{}{
			"activity": []interface{}{},
			"total":    0,
		}
		json.NewEncoder(w).Encode(response)
	}))
	defer server.Close()

	client, _ := NewClient("test_key", WithBaseURL(server.URL))
	
	page := 1
	limit := 50
	params := GetActivityParams{
		Page:  &page,
		Limit: &limit,
	}
	
	activity, err := client.Activity.Get(params)

	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	if activity == nil {
		t.Error("Expected activity response, got nil")
	}
}
