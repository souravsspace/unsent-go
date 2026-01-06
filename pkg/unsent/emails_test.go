package unsent

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestEmails_Send(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "POST" {
			t.Errorf("expected POST, got %s", r.Method)
		}
		var body SendEmailJSONBody
		json.NewDecoder(r.Body).Decode(&body)
		if body.From != "me@test.com" {
			t.Errorf("expected me@test.com, got %s", body.From)
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"emailId": "email_123"}`))
	}))
	defer server.Close()

	client, _ := NewClient("key", WithBaseURL(server.URL))
	

	
	resp, err := client.Emails.Send(SendEmailJSONBody{
		From: "me@test.com",
		To: SendEmailJSONBody_To{
			union: json.RawMessage(`"you@test.com"`), // Using manual json for simplicity
		},
		Subject: stringPtr("Hello"),
		Html: stringPtr("<p>Hi</p>"),
	})
	// Note: Union type handling in tests might be tricky if we don't assume the helper logic.
	// But `SendEmailJSONBody` has `To SendEmailJSONBody_To`.
	// Correct way to set To for test:
	// We need to marshal it correctly or let encoding/json handle it.
	// In the test setup above, I used `json.RawMessage`. 
	// But the client code will marshal `SendEmailJSONBody`.
	// We should construct it properly.
	
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if resp.EmailID != "email_123" {
		t.Errorf("expected email_123, got %s", resp.EmailID)
	}
}

func stringPtr(s string) *string {
	return &s
}

func float32Ptr(f float32) *float32 {
	return &f
}

func TestEmails_List(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "GET" {
			t.Errorf("expected GET, got %s", r.Method)
		}
		if r.URL.Path != "/v1/emails" {
			t.Errorf("expected /v1/emails, got %s", r.URL.Path)
		}
		// Check query params
		query := r.URL.Query()
		if query.Get("page") != "1" {
			t.Errorf("expected page=1, got %s", query.Get("page"))
		}
		if query.Get("limit") != "50" {
			t.Errorf("expected limit=50, got %s", query.Get("limit"))
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`[{"id": "email_1", "to": "test@example.com", "from": "sender@example.com", "subject": "Test", "status": "sent", "createdAt": "2024-01-01T00:00:00Z", "updatedAt": "2024-01-01T00:00:00Z"}]`))
	}))
	defer server.Close()

	client, _ := NewClient("key", WithBaseURL(server.URL))

	resp, err := client.Emails.List(ListEmailsParams{
		Page:  stringPtr("1"),
		Limit: stringPtr("50"),
	})

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(*resp) != 1 {
		t.Errorf("expected 1 email, got %d", len(*resp))
	}
	if (*resp)[0].ID != "email_1" {
		t.Errorf("expected email_1, got %s", (*resp)[0].ID)
	}
}

func TestEmails_GetBounces(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "GET" {
			t.Errorf("expected GET, got %s", r.Method)
		}
		if r.URL.Path != "/v1/emails/bounces" {
			t.Errorf("expected /v1/emails/bounces, got %s", r.URL.Path)
		}
		query := r.URL.Query()
		if query.Get("page") != "1.000000" {
			t.Errorf("expected page=1.000000, got %s", query.Get("page"))
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`[{"id": "email_bounce_1", "to": "bounce@example.com", "from": "sender@example.com", "subject": "Bounced", "status": "bounced", "createdAt": "2024-01-01T00:00:00Z", "updatedAt": "2024-01-01T00:00:00Z"}]`))
	}))
	defer server.Close()

	client, _ := NewClient("key", WithBaseURL(server.URL))

	resp, err := client.Emails.GetBounces(GetBouncesParams{
		Page:  float32Ptr(1.0),
		Limit: float32Ptr(20.0),
	})

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(*resp) != 1 {
		t.Errorf("expected 1 bounced email, got %d", len(*resp))
	}
	if (*resp)[0].Status != "bounced" {
		t.Errorf("expected status bounced, got %s", (*resp)[0].Status)
	}
}

func TestEmails_GetComplaints(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "GET" {
			t.Errorf("expected GET, got %s", r.Method)
		}
		if r.URL.Path != "/v1/emails/complaints" {
			t.Errorf("expected /v1/emails/complaints, got %s", r.URL.Path)
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`[{"id": "email_complaint_1", "to": "complaint@example.com", "from": "sender@example.com", "subject": "Complained", "status": "complained", "createdAt": "2024-01-01T00:00:00Z", "updatedAt": "2024-01-01T00:00:00Z"}]`))
	}))
	defer server.Close()

	client, _ := NewClient("key", WithBaseURL(server.URL))

	resp, err := client.Emails.GetComplaints(GetComplaintsParams{
		Page:  float32Ptr(1.0),
		Limit: float32Ptr(20.0),
	})

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(*resp) != 1 {
		t.Errorf("expected 1 complaint email, got %d", len(*resp))
	}
}

func TestEmails_GetUnsubscribes(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "GET" {
			t.Errorf("expected GET, got %s", r.Method)
		}
		if r.URL.Path != "/v1/emails/unsubscribes" {
			t.Errorf("expected /v1/emails/unsubscribes, got %s", r.URL.Path)
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`[{"id": "email_unsub_1", "to": "unsub@example.com", "from": "sender@example.com", "subject": "Unsubscribed", "status": "unsubscribed", "createdAt": "2024-01-01T00:00:00Z", "updatedAt": "2024-01-01T00:00:00Z"}]`))
	}))
	defer server.Close()

	client, _ := NewClient("key", WithBaseURL(server.URL))

	resp, err := client.Emails.GetUnsubscribes(GetUnsubscribesParams{
		Page:  float32Ptr(1.0),
		Limit: float32Ptr(20.0),
	})

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(*resp) != 1 {
		t.Errorf("expected 1 unsubscribed email, got %d", len(*resp))
	}
}
