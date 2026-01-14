package unsent

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestStatsGet(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/v1/stats" {
			t.Errorf("Expected path /v1/stats, got %s", r.URL.Path)
		}
		if r.Method != "GET" {
			t.Errorf("Expected GET method, got %s", r.Method)
		}

		response := map[string]interface{}{
			"stats": map[string]int{
				"sent":      100,
				"delivered": 95,
				"opened":    25,
			},
		}
		json.NewEncoder(w).Encode(response)
	}))
	defer server.Close()

	client, _ := NewClient("test_key", WithBaseURL(server.URL))
	
	params := GetStatsParams{}
	
	stats, err := client.Stats.Get(params)

	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	if stats == nil {
		t.Error("Expected stats response, got nil")
	}
}
