package unsent

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestClient_Get(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "GET" {
			t.Errorf("expected GET request, got %s", r.Method)
		}
		if r.Header.Get("Authorization") != "Bearer test_key" {
			t.Errorf("expected Authorization header with test_key, got %s", r.Header.Get("Authorization"))
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"success": true}`))
	}))
	defer server.Close()

	client, err := NewClient("test_key", WithBaseURL(server.URL))
	if err != nil {
		t.Fatalf("failed to create client: %v", err)
	}

	// We need to access the private request method or use a public wrapper.
	// But Get is public.
	// We need a response type.
	type Response struct {
		Success bool `json:"success"`
	}

	resp, apiErr := Get[Response](client, "/test")
	if apiErr != nil {
		t.Errorf("expected no error, got %v", apiErr)
	}
	if resp == nil {
		t.Fatal("expected response, got nil")
	}
	if !resp.Success {
		t.Error("expected success true, got false")
	}
}

func TestClient_Post(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "POST" {
			t.Errorf("expected POST request, got %s", r.Method)
		}
		var body map[string]string
		if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
			t.Errorf("failed to decode body: %v", err)
		}
		if body["foo"] != "bar" {
			t.Errorf("expected body foo=bar, got %v", body)
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"id": "123"}`))
	}))
	defer server.Close()

	client, _ := NewClient("test_key", WithBaseURL(server.URL))

	type Response struct {
		ID string `json:"id"`
	}

	resp, apiErr := Post[Response](client, "/test", map[string]string{"foo": "bar"})
	if apiErr != nil {
		t.Errorf("expected no error, got %v", apiErr)
	}
	if resp.ID != "123" {
		t.Errorf("expected ID 123, got %s", resp.ID)
	}
}

func TestClient_Error(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"code": "BAD_REQUEST", "message": "fail"}`))
	}))
	defer server.Close()

	client, _ := NewClient("test_key", WithBaseURL(server.URL))

	_, apiErr := Get[struct{}](client, "/test")
	if apiErr == nil {
		t.Fatal("expected error, got nil")
	}
	if apiErr.Code != "BAD_REQUEST" {
		t.Errorf("expected code BAD_REQUEST, got %s", apiErr.Code)
	}
}
