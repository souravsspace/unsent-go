package unsent

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestDomains_List(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/v1/domains" {
			t.Errorf("expected path /v1/domains, got %s", r.URL.Path)
		}
		if r.Method != "GET" {
			t.Errorf("expected GET, got %s", r.Method)
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`[{"id": "123", "domain": "example.com"}]`))
	}))
	defer server.Close()

	client, _ := NewClient("key", WithBaseURL(server.URL))
	domains, err := client.Domains.List()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(*domains) != 1 {
		t.Errorf("expected 1 domain, got %d", len(*domains))
	}
	if (*domains)[0].Domain != "example.com" {
		t.Errorf("expected example.com, got %s", (*domains)[0].Domain)
	}
}

func TestDomains_Create(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "POST" {
			t.Errorf("expected POST, got %s", r.Method)
		}
		var body CreateDomainJSONBody
		json.NewDecoder(r.Body).Decode(&body)
		if body.Name != "new.com" {
			t.Errorf("expected new.com, got %s", body.Name)
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"id": "123", "domain": "new.com", "status": "pending"}`))
	}))
	defer server.Close()

	client, _ := NewClient("key", WithBaseURL(server.URL))
	resp, err := client.Domains.Create(CreateDomainJSONBody{Name: "new.com", Region: "us-east-1"})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if resp.Domain != "new.com" {
		t.Errorf("expected new.com, got %s", resp.Domain)
	}
}

func TestDomains_Delete(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "DELETE" {
			t.Errorf("expected DELETE, got %s", r.Method)
		}
		if r.URL.Path != "/v1/domains/123" {
			t.Errorf("expected /v1/domains/123, got %s", r.URL.Path)
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"id": "123", "deleted": true}`))
	}))
	defer server.Close()

	client, _ := NewClient("key", WithBaseURL(server.URL))
	resp, err := client.Domains.Delete("123")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !resp.Deleted {
		t.Error("expected deleted true")
	}
}
