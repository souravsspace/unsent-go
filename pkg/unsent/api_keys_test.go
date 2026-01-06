package unsent

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestApiKeys_Create(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "POST" {
			t.Errorf("expected POST, got %s", r.Method)
		}
		var body CreateApiKeyJSONBody
		json.NewDecoder(r.Body).Decode(&body)
		if body.Name != "Key1" {
			t.Errorf("expected Key1, got %s", body.Name)
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"id": "k1", "token": "secret"}`))
	}))
	defer server.Close()

	client, _ := NewClient("key", WithBaseURL(server.URL))
	
	resp, err := client.ApiKeys.Create(CreateApiKeyJSONBody{Name: "Key1"})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if resp.Token != "secret" {
		t.Errorf("expected secret, got %s", resp.Token)
	}
}
