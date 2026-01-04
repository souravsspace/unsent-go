package unsent

import (
	"fmt"
	"time"
)

// APIError represents an error response from the API
type APIError struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

func (e *APIError) Error() string {
	if e == nil {
		return "nil APIError"
	}
	return fmt.Sprintf("API Error: %s (code: %s)", e.Message, e.Code)
}

// Attachment represents an email attachment
type Attachment struct {
	Filename string `json:"filename"`
	Content  string `json:"content"`
}

// EmailCreate represents the payload for creating an email
type EmailCreate struct {
	To          string       `json:"to"`
	From        string       `json:"from"`
	Subject     string       `json:"subject"`
	HTML        string       `json:"html,omitempty"`
	Text        string       `json:"text,omitempty"`
	ReplyTo     string       `json:"replyTo,omitempty"`
	CC          []string     `json:"cc,omitempty"`
	BCC         []string     `json:"bcc,omitempty"`
	Attachments []Attachment `json:"attachments,omitempty"`
	ScheduledAt *time.Time   `json:"scheduledAt,omitempty"`
}

// EmailCreateResponse represents the response from creating an email
type EmailCreateResponse struct {
	EmailID string `json:"emailId"`
}

// Email represents an email
type Email struct {
	ID          string       `json:"id"`
	To          string       `json:"to"`
	From        string       `json:"from"`
	Subject     string       `json:"subject"`
	HTML        string       `json:"html,omitempty"`
	Text        string       `json:"text,omitempty"`
	Status      string       `json:"status"`
	Attachments []Attachment `json:"attachments,omitempty"`
	ScheduledAt *time.Time   `json:"scheduledAt,omitempty"`
	SentAt      *time.Time   `json:"sentAt,omitempty"`
	CreatedAt   time.Time    `json:"createdAt"`
	UpdatedAt   time.Time    `json:"updatedAt"`
}

// EmailUpdate represents the payload for updating an email
type EmailUpdate struct {
	Subject     string       `json:"subject,omitempty"`
	HTML        string       `json:"html,omitempty"`
	Text        string       `json:"text,omitempty"`
	ScheduledAt *time.Time   `json:"scheduledAt,omitempty"`
	Attachments []Attachment `json:"attachments,omitempty"`
}

// EmailUpdateResponse represents the response from updating an email
type EmailUpdateResponse struct {
	EmailID string `json:"emailId"`
}

// EmailCancelResponse represents the response from canceling an email
type EmailCancelResponse struct {
	EmailID string `json:"emailId"`
}

// EmailBatchItem represents a single email in a batch
type EmailBatchItem struct {
	To          string       `json:"to"`
	From        string       `json:"from"`
	Subject     string       `json:"subject"`
	HTML        string       `json:"html,omitempty"`
	Text        string       `json:"text,omitempty"`
	Attachments []Attachment `json:"attachments,omitempty"`
	ScheduledAt *time.Time   `json:"scheduledAt,omitempty"`
}

// EmailBatchResponse represents the response from sending batch emails
type EmailBatchResponse struct {
	Data []EmailCreateResponse `json:"data"`
}

// Contact represents a contact
type Contact struct {
	ID        string                 `json:"id"`
	Email     string                 `json:"email"`
	FirstName string                 `json:"firstName,omitempty"`
	LastName  string                 `json:"lastName,omitempty"`
	Metadata  map[string]interface{} `json:"metadata,omitempty"`
	CreatedAt time.Time              `json:"createdAt"`
	UpdatedAt time.Time              `json:"updatedAt"`
}

// ContactCreate represents the payload for creating a contact
type ContactCreate struct {
	Email     string                 `json:"email"`
	FirstName string                 `json:"firstName,omitempty"`
	LastName  string                 `json:"lastName,omitempty"`
	Metadata  map[string]interface{} `json:"metadata,omitempty"`
}

// ContactCreateResponse represents the response from creating a contact
type ContactCreateResponse struct {
	ID        string    `json:"id"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"createdAt"`
}

// ContactUpdate represents the payload for updating a contact
type ContactUpdate struct {
	FirstName string                 `json:"firstName,omitempty"`
	LastName  string                 `json:"lastName,omitempty"`
	Metadata  map[string]interface{} `json:"metadata,omitempty"`
}

// ContactUpdateResponse represents the response from updating a contact
type ContactUpdateResponse struct {
	ID        string    `json:"id"`
	UpdatedAt time.Time `json:"updatedAt"`
}

// ContactUpsert represents the payload for upserting a contact
type ContactUpsert struct {
	Email     string                 `json:"email"`
	FirstName string                 `json:"firstName,omitempty"`
	LastName  string                 `json:"lastName,omitempty"`
	Metadata  map[string]interface{} `json:"metadata,omitempty"`
}

// ContactUpsertResponse represents the response from upserting a contact
type ContactUpsertResponse struct {
	ID        string    `json:"id"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

// ContactDeleteResponse represents the response from deleting a contact
type ContactDeleteResponse struct {
	ID      string `json:"id"`
	Deleted bool   `json:"deleted"`
}

// Campaign represents a campaign
type Campaign struct {
	ID            string    `json:"id"`
	Name          string    `json:"name"`
	Subject       string    `json:"subject"`
	HTML          string    `json:"html"`
	From          string    `json:"from"`
	ContactBookID string    `json:"contactBookId"`
	Status        string    `json:"status"`
	Total         int       `json:"total"`
	Sent          int       `json:"sent"`
	Failed        int       `json:"failed"`
	ScheduledAt   time.Time `json:"scheduledAt,omitempty"`
	CreatedAt     time.Time `json:"createdAt"`
	UpdatedAt     time.Time `json:"updatedAt"`
}

// CampaignCreate represents the payload for creating a campaign
type CampaignCreate struct {
	Name          string `json:"name"`
	Subject       string `json:"subject"`
	HTML          string `json:"html"`
	From          string `json:"from"`
	ContactBookID string `json:"contactBookId"`
}

// CampaignCreateResponse represents the response from creating a campaign
type CampaignCreateResponse struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	Status    string    `json:"status"`
	CreatedAt time.Time `json:"createdAt"`
}

// CampaignSchedule represents the payload for scheduling a campaign
type CampaignSchedule struct {
	ScheduledAt string `json:"scheduledAt"`
}

// CampaignScheduleResponse represents the response from scheduling a campaign
type CampaignScheduleResponse struct {
	ID          string    `json:"id"`
	Status      string    `json:"status"`
	ScheduledAt time.Time `json:"scheduledAt"`
}

// CampaignActionResponse represents the response from pausing/resuming a campaign
type CampaignActionResponse struct {
	ID        string    `json:"id"`
	Status    string    `json:"status"`
	UpdatedAt time.Time `json:"updatedAt"`
}

// Domain represents a domain
type Domain struct {
	ID        int       `json:"id"`
	Domain    string    `json:"domain"`
	Status    string    `json:"status"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

// DomainCreate represents the payload for creating a domain
type DomainCreate struct {
	Domain string `json:"domain"`
}

// DomainCreateResponse represents the response from creating a domain
type DomainCreateResponse struct {
	ID        int       `json:"id"`
	Domain    string    `json:"domain"`
	Status    string    `json:"status"`
	CreatedAt time.Time `json:"createdAt"`
}

// DomainVerifyResponse represents the response from verifying a domain
type DomainVerifyResponse struct {
	ID        int       `json:"id"`
	Status    string    `json:"status"`
	UpdatedAt time.Time `json:"updatedAt"`
}

// DomainDeleteResponse represents the response from deleting a domain
type DomainDeleteResponse struct {
	ID      int  `json:"id"`
	Deleted bool `json:"deleted"`
}
