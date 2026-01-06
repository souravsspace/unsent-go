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
