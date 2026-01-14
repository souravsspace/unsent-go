package unsent

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestMetricsGet(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/v1/metrics" {
			t.Errorf("Expected path /v1/metrics, got %s", r.URL.Path)
		}
		if r.Method != "GET" {
			t.Errorf("Expected GET method, got %s", r.Method)
		}

		response := map[string]interface{}{
			"metrics": map[string]interface{}{
				"deliveryRate": 0.95,
				"openRate":     0.25,
			},
		}
		json.NewEncoder(w).Encode(response)
	}))
	defer server.Close()

	client, _ := NewClient("test_key", WithBaseURL(server.URL))
	
	period := GetMetricsParamsPeriodMonth
	params := GetMetricsParams{
		Period: &period,
	}
	
	metrics, err := client.Metrics.Get(params)

	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	if metrics == nil {
		t.Error("Expected metrics response, got nil")
	}
}
