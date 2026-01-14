package unsent

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestSettings_Get(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "GET" {
			t.Errorf("expected GET, got %s", r.Method)
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"name": "My Team", "plan": "pro"}`))
	}))
	defer server.Close()

	client, _ := NewClient("key", WithBaseURL(server.URL))
	
	resp, err := client.Settings.Get()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if resp.TeamName != "My Team" {
		t.Errorf("expected My Team, got %s", resp.TeamName)
	}
}
