package unsent

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestTemplates_Create(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "POST" {
			t.Errorf("expected POST, got %s", r.Method)
		}
		var body CreateTemplateJSONBody
		json.NewDecoder(r.Body).Decode(&body)
		if body.Name != "Tpl 1" {
			t.Errorf("expected Tpl 1, got %s", body.Name)
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"id": "t1", "name": "Tpl 1"}`))
	}))
	defer server.Close()

	client, _ := NewClient("key", WithBaseURL(server.URL))
	
	resp, err := client.Templates.Create(CreateTemplateJSONBody{
		Name: "Tpl 1",
		Subject: "Hello",
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if resp.ID != "t1" {
		t.Errorf("expected t1, got %s", resp.ID)
	}
}

func TestTemplates_List(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "GET" {
			t.Errorf("expected GET, got %s", r.Method)
		}
		if r.URL.Path != "/v1/templates" {
			t.Errorf("expected /v1/templates, got %s", r.URL.Path)
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`[{"id": "t1", "name": "Tpl 1"}]`))
	}))
	defer server.Close()
	client, _ := NewClient("key", WithBaseURL(server.URL))
	resp, err := client.Templates.List()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(*resp) != 1 {
		t.Errorf("expected 1 template, got %d", len(*resp))
	}
}

func TestTemplates_Get(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "GET" {
			t.Errorf("expected GET, got %s", r.Method)
		}
		if r.URL.Path != "/v1/templates/t1" {
			t.Errorf("expected /v1/templates/t1, got %s", r.URL.Path)
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"id": "t1", "name": "Tpl 1"}`))
	}))
	defer server.Close()
	client, _ := NewClient("key", WithBaseURL(server.URL))
	resp, err := client.Templates.Get("t1")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if resp.ID != "t1" {
		t.Errorf("expected t1, got %s", resp.ID)
	}
}

func TestTemplates_Update(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "PATCH" {
			t.Errorf("expected PATCH, got %s", r.Method)
		}
		if r.URL.Path != "/v1/templates/t1" {
			t.Errorf("expected /v1/templates/t1, got %s", r.URL.Path)
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"id": "t1", "updatedAt": "2024-01-01T00:00:00Z"}`))
	}))
	defer server.Close()
	client, _ := NewClient("key", WithBaseURL(server.URL))
	name := "Updated Tpl"
	resp, err := client.Templates.Update("t1", UpdateTemplateJSONBody{Name: &name})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if resp.ID != "t1" {
		t.Errorf("expected t1, got %s", resp.ID)
	}
}

func TestTemplates_Delete(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "DELETE" {
			t.Errorf("expected DELETE, got %s", r.Method)
		}
		if r.URL.Path != "/v1/templates/t1" {
			t.Errorf("expected /v1/templates/t1, got %s", r.URL.Path)
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"success": true}`))
	}))
	defer server.Close()
	client, _ := NewClient("key", WithBaseURL(server.URL))
	_, err := client.Templates.Delete("t1")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
}
