package unsent

import "time"

// EmailCreateResponse represents the response from creating an email
type EmailCreateResponse struct {
	EmailID string `json:"emailId"`
}

// Email represents an email
type Email struct {
	ID          string                   `json:"id"`
	To          string                   `json:"to"`
	From        string                   `json:"from"`
	Subject     string                   `json:"subject"`
	HTML        string                   `json:"html,omitempty"`
	Text        string                   `json:"text,omitempty"`
	Status      string                   `json:"status"`
	Attachments []map[string]interface{} `json:"attachments,omitempty"`
	ScheduledAt *time.Time               `json:"scheduledAt,omitempty"`
	SentAt      *time.Time               `json:"sentAt,omitempty"`
	CreatedAt   time.Time                `json:"createdAt"`
	UpdatedAt   time.Time                `json:"updatedAt"`
}

// EmailUpdateResponse represents the response from updating an email
type EmailUpdateResponse struct {
	EmailID string `json:"emailId"`
}

// EmailCancelResponse represents the response from canceling an email
type EmailCancelResponse struct {
	EmailID string `json:"emailId"`
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

// ContactCreateResponse represents the response from creating a contact
type ContactCreateResponse struct {
	ID        string    `json:"id"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"createdAt"`
}

// ContactUpdateResponse represents the response from updating a contact
type ContactUpdateResponse struct {
	ID        string    `json:"id"`
	UpdatedAt time.Time `json:"updatedAt"`
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

// CampaignCreateResponse represents the response from creating a campaign
type CampaignCreateResponse struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	Status    string    `json:"status"`
	CreatedAt time.Time `json:"createdAt"`
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
	ID        string    `json:"id"`
	Domain    string    `json:"domain"`
	Status    string    `json:"status"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

// DomainCreateResponse represents the response from creating a domain
type DomainCreateResponse struct {
	ID        string    `json:"id"`
	Domain    string    `json:"domain"`
	Status    string    `json:"status"`
	CreatedAt time.Time `json:"createdAt"`
}

// DomainVerifyResponse represents the response from verifying a domain
type DomainVerifyResponse struct {
	ID        string    `json:"id"`
	Status    string    `json:"status"`
	UpdatedAt time.Time `json:"updatedAt"`
}

// DomainDeleteResponse represents the response from deleting a domain
type DomainDeleteResponse struct {
	ID      string `json:"id"`
	Deleted bool   `json:"deleted"`
}

// Analytics Metrics
type Analytics struct {
	Total      int `json:"total"`
	Sent       int `json:"sent"`
	Delivered  int `json:"delivered"`
	Opened     int `json:"opened"`
	Clicked    int `json:"clicked"`
	Failed     int `json:"failed"`
	Complained int `json:"complained"`
	Bounced    int `json:"bounced"`
}

type AnalyticsTimeSeries struct {
	Date       string `json:"date"`
	Total      int    `json:"total"`
	Sent       int    `json:"sent"`
	Delivered  int    `json:"delivered"`
	Opened     int    `json:"opened"`
	Clicked    int    `json:"clicked"`
	Failed     int    `json:"failed"`
	Complained int    `json:"complained"`
	Bounced    int    `json:"bounced"`
}

type AnalyticsReputation struct {
	Domain     string `json:"domain"`
	Reputation int    `json:"reputation"`
}

// ApiKey
type ApiKey struct {
	ID           string `json:"id"`
	Name         string `json:"name"`
	PartialToken string `json:"partialToken"`
	Permission   string `json:"permission"`
	CreatedAt    string `json:"createdAt"`
}

type ApiKeyCreateResponse struct {
	ID    string `json:"id"`
	Token string `json:"token"`
}

type ApiKeyDeleteResponse struct {
	ID      string `json:"id"`
	Deleted bool   `json:"deleted"`
}

// ContactBook
type ContactBook struct {
	ID         string            `json:"id"`
	Name       string            `json:"name"`
	Emoji      string            `json:"emoji,omitempty"`
	Properties map[string]string `json:"properties,omitempty"`
	CreatedAt  time.Time         `json:"createdAt"`
	UpdatedAt  time.Time         `json:"updatedAt"`
	Total      int               `json:"total"`
}

type ContactBookCreateResponse struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"createdAt"`
}

type ContactBookUpdateResponse struct {
	ID        string    `json:"id"`
	UpdatedAt time.Time `json:"updatedAt"`
}

type ContactBookDeleteResponse struct {
	ID      string `json:"id"`
	Deleted bool   `json:"deleted"`
}

// Settings
type Settings struct {
	TeamName string `json:"teamName"`
	Plan     string `json:"plan"`
}

// Suppression
type Suppression struct {
	Email     string    `json:"email"`
	Reason    string    `json:"reason"`
	Source    string    `json:"source"`
	CreatedAt time.Time `json:"createdAt"`
}

type SuppressionAddResponse struct {
	Email     string    `json:"email"`
	Reason    string    `json:"reason"`
	CreatedAt time.Time `json:"createdAt"`
}

type SuppressionDeleteResponse struct {
	Email   string `json:"email"`
	Deleted bool   `json:"deleted"`
}

// Template
type Template struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	Subject   string    `json:"subject"`
	Content   string    `json:"content,omitempty"`
	HTML      string    `json:"html,omitempty"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

type TemplateCreateResponse struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"createdAt"`
}

type TemplateUpdateResponse struct {
	ID        string    `json:"id"`
	UpdatedAt time.Time `json:"updatedAt"`
}

type TemplateDeleteResponse struct {
	ID      string `json:"id"`
	Deleted bool   `json:"deleted"`
}

// Webhook
type Webhook struct {
	ID        string    `json:"id"`
	Url       string    `json:"url"`
	Events    []string  `json:"events"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

type WebhookCreateRequest struct {
	Url    string   `json:"url"`
	Events []string `json:"events"`
}

type WebhookUpdateRequest struct {
	Url    string   `json:"url,omitempty"`
	Events []string `json:"events,omitempty"`
}

type WebhookCreateResponse struct {
	ID string `json:"id"`
}

type WebhookUpdateResponse struct {
	Success bool `json:"success"`
}

type WebhookDeleteResponse struct {
	Success bool `json:"success"`
}

// Pagination support
type PaginationMeta struct {
	Total      int `json:"total"`
	Page       int `json:"page"`
	Limit      int `json:"limit"`
	TotalPages int `json:"totalPages"`
}

// Enhanced Campaign response with all API fields
type CampaignDetail struct {
	ID                 string    `json:"id"`
	Name               string    `json:"name"`
	From               string    `json:"from"`
	Subject            string    `json:"subject"`
	PreviewText        *string   `json:"previewText,omitempty"`
	ContactBookID      *string   `json:"contactBookId,omitempty"`
	HTML               *string   `json:"html,omitempty"`
	Content            *string   `json:"content,omitempty"`
	Status             string    `json:"status"`
	ScheduledAt        *string   `json:"scheduledAt,omitempty"`
	BatchSize          int       `json:"batchSize"`
	BatchWindowMinutes int       `json:"batchWindowMinutes"`
	Total              int       `json:"total"`
	Sent               int       `json:"sent"`
	Delivered          int       `json:"delivered"`
	Opened             int       `json:"opened"`
	Clicked            int       `json:"clicked"`
	Unsubscribed       int       `json:"unsubscribed"`
	Bounced            int       `json:"bounced"`
	HardBounced        int       `json:"hardBounced"`
	Complained         int       `json:"complained"`
	ReplyTo            []string  `json:"replyTo"`
	Cc                 []string  `json:"cc"`
	Bcc                []string  `json:"bcc"`
	CreatedAt          string    `json:"createdAt"`
	UpdatedAt          string    `json:"updatedAt"`
}

// ContactBook with details as returned by API
type ContactBookDetail struct {
	ID         string            `json:"id"`
	Name       string            `json:"name"`
	Emoji      string            `json:"emoji"`
	Properties map[string]string `json:"properties,omitempty"`
	TeamID     string            `json:"teamId"`
	CreatedAt  string            `json:"createdAt"`
	UpdatedAt  string            `json:"updatedAt"`
	Details    struct {
		TotalContacts        int           `json:"totalContacts"`
		UnsubscribedContacts int           `json:"unsubscribedContacts"`
		Campaigns            []interface{} `json:"campaigns"`
	} `json:"details,omitempty"`
}
