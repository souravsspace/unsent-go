package unsent

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestSystemHealth(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/v1/health" {
			t.Errorf("Expected path /v1/health, got %s", r.URL.Path)
		}
		if r.Method != "GET" {
			t.Errorf("Expected GET method, got %s", r.Method)
		}

		response := HealthResponse{
			Status: "ok",
			Uptime: 12345.67,
		}
		json.NewEncoder(w).Encode(response)
	}))
	defer server.Close()

	client, _ := NewClient("test_key", WithBaseURL(server.URL))
	health, err := client.System.Health()

	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	if health.Status != "ok" {
		t.Errorf("Expected status 'ok', got %s", health.Status)
	}
	if health.Uptime != 12345.67 {
		t.Errorf("Expected uptime 12345.67, got %f", health.Uptime)
	}
}

func TestSystemVersion(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/v1/version" {
			t.Errorf("Expected path /v1/version, got %s", r.URL.Path)
		}
		if r.Method != "GET" {
			t.Errorf("Expected GET method, got %s", r.Method)
		}

		response := VersionResponse{
			Version:  "1.0.0",
			Platform: "linux",
		}
		json.NewEncoder(w).Encode(response)
	}))
	defer server.Close()

	client, _ := NewClient("test_key", WithBaseURL(server.URL))
	version, err := client.System.Version()

	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	if version.Version != "1.0.0" {
		t.Errorf("Expected version '1.0.0', got %s", version.Version)
	}
	if version.Platform != "linux" {
		t.Errorf("Expected platform 'linux', got %s", version.Platform)
	}
}
