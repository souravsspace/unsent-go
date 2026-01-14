package unsent

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestContacts_List(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "GET" {
			t.Errorf("expected GET, got %s", r.Method)
		}
		if r.URL.Path != "/v1/contactBooks/book1/contacts" {
			t.Errorf("expected path /v1/contactBooks/book1/contacts, got %s", r.URL.Path)
		}
		// Check query params if needed, but for now just basic list
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`[{"id": "c1", "email": "test@example.com"}]`))
	}))
	defer server.Close()

	client, _ := NewClient("key", WithBaseURL(server.URL))

	resp, err := client.Contacts.List("book1", GetContactsParams{})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(*resp) != 1 {
		t.Errorf("expected 1 contact, got %d", len(*resp))
	}
	if (*resp)[0].ID != "c1" {
		t.Errorf("expected c1, got %s", (*resp)[0].ID)
	}
}

func TestContacts_Create(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "POST" {
			t.Errorf("expected POST, got %s", r.Method)
		}
		if r.URL.Path != "/v1/contactBooks/book1/contacts" {
			t.Errorf("expected path /v1/contactBooks/book1/contacts, got %s", r.URL.Path)
		}
		var body CreateContactJSONBody
		json.NewDecoder(r.Body).Decode(&body)
		if body.Email != "test@example.com" {
			t.Errorf("expected test@example.com, got %s", body.Email)
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"contactId": "c1", "email": "test@example.com"}`))
	}))
	defer server.Close()

	client, _ := NewClient("key", WithBaseURL(server.URL))
	
	resp, err := client.Contacts.Create("book1", CreateContactJSONBody{
		Email: "test@example.com",
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if resp.ID != "c1" {
		t.Errorf("expected c1, got %s", resp.ID)
	}
}

func TestContacts_Get(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "GET" {
			t.Errorf("expected GET, got %s", r.Method)
		}
		if r.URL.Path != "/v1/contactBooks/book1/contacts/c1" {
			t.Errorf("expected path /v1/contactBooks/book1/contacts/c1, got %s", r.URL.Path)
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"id": "c1", "email": "test@example.com"}`))
	}))
	defer server.Close()

	client, _ := NewClient("key", WithBaseURL(server.URL))
	
	resp, err := client.Contacts.Get("book1", "c1")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if resp.Email != "test@example.com" {
		t.Errorf("expected test@example.com, got %s", resp.Email)
	}
}
