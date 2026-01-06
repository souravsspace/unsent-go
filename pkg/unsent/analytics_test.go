package unsent

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestAnalytics_Get(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "GET" {
			t.Errorf("expected GET, got %s", r.Method)
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"total": 100, "sent": 90}`))
	}))
	defer server.Close()

	client, _ := NewClient("key", WithBaseURL(server.URL))
	
	resp, err := client.Analytics.Get()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if resp.Total != 100 {
		t.Errorf("expected 100, got %d", resp.Total)
	}
}
