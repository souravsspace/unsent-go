package unsent

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestContactBooks_Create(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "POST" {
			t.Errorf("expected POST, got %s", r.Method)
		}
		var body CreateContactBookJSONBody
		json.NewDecoder(r.Body).Decode(&body)
		if body.Name != "Book 1" {
			t.Errorf("expected Book 1, got %s", body.Name)
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"id": "cb1", "name": "Book 1"}`))
	}))
	defer server.Close()

	client, _ := NewClient("key", WithBaseURL(server.URL))
	
	resp, err := client.ContactBooks.Create(CreateContactBookJSONBody{Name: "Book 1"})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if resp.ID != "cb1" {
		t.Errorf("expected cb1, got %s", resp.ID)
	}
}
