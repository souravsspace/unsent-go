package unsent

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestSuppressions_Add(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "POST" {
			t.Errorf("expected POST, got %s", r.Method)
		}
		var body AddSuppressionJSONBody
		json.NewDecoder(r.Body).Decode(&body)
		if body.Reason != AddSuppressionJSONBodyReasonMANUAL {
			t.Errorf("expected MANUAL, got %s", body.Reason)
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"email": "bad@guy.com", "reason": "MANUAL"}`))
	}))
	defer server.Close()

	client, _ := NewClient("key", WithBaseURL(server.URL))
	
	resp, err := client.Suppressions.Add(AddSuppressionJSONBody{
		Email: "bad@guy.com",
		Reason: AddSuppressionJSONBodyReasonMANUAL,
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if resp.Email != "bad@guy.com" {
		t.Errorf("expected bad@guy.com, got %s", resp.Email)
	}
}

func TestSuppressions_Delete(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "DELETE" {
			t.Errorf("expected DELETE, got %s", r.Method)
		}
		if r.URL.Path != "/v1/suppressions/email/bad@guy.com" {
			t.Errorf("expected /v1/suppressions/email/bad@guy.com, got %s", r.URL.Path)
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"success": true}`))
	}))
	defer server.Close()

	client, _ := NewClient("key", WithBaseURL(server.URL))
	
	_, err := client.Suppressions.Delete("bad@guy.com")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
}
