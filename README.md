# Unsent Go SDK

Official Go SDK for the [Unsent API](https://unsent.dev) - Send transactional emails with ease.

## Prerequisites

- [Unsent API key](https://app.unsent.dev/dev-settings/api-keys)
- [Verified domain](https://app.unsent.dev/domains)
- Go 1.19 or higher

## Installation

```bash
go get github.com/souravsspace/unsent-go
```

## Usage

### Basic Setup

```go
package main

import (
    "fmt"
    "log"
    
    "github.com/souravsspace/unsent-go/pkg/unsent"
)

func main() {
    client, err := unsent.NewClient("un_xxxx")
    if err != nil {
        log.Fatal(err)
    }
    
    // Use the client
}
```

### Environment Variables

You can also set your API key using environment variables:

```go
// Set UNSENT_API_KEY in your environment
// Then initialize without passing the key
client, err := unsent.NewClient("")
```

### Sending Emails

#### Simple Email

```go
email, err := client.Emails.Send(unsent.EmailCreate{
    To:      "hello@acme.com",
    From:    "hello@company.com",
    Subject: "Unsent email",
    HTML:    "<p>Unsent is the best email service provider to send emails</p>",
    Text:    "Unsent is the best email service provider to send emails",
})

if err != nil {
    log.Printf("Error: %v", err)
} else {
    fmt.Printf("Email sent! ID: %s\n", email.ID)
}
```

#### Email with Attachments

```go
email, err := client.Emails.Send(unsent.EmailCreate{
    To:      "hello@acme.com",
    From:    "hello@company.com",
    Subject: "Email with attachment",
    HTML:    "<p>Please find the attachment below</p>",
    Attachments: []unsent.Attachment{
        {
            Filename: "document.pdf",
            Content:  "base64-encoded-content-here",
        },
    },
})
```

#### Scheduled Email

```go
import "time"

scheduledTime := time.Now().Add(1 * time.Hour)

email, err := client.Emails.Send(unsent.EmailCreate{
    To:          "hello@acme.com",
    From:        "hello@company.com",
    Subject:     "Scheduled email",
    HTML:        "<p>This email was scheduled</p>",
    ScheduledAt: &scheduledTime,
})
```

#### Batch Emails

```go
emails := []unsent.EmailBatchItem{
    {
        To:      "user1@example.com",
        From:    "hello@company.com",
        Subject: "Hello User 1",
        HTML:    "<p>Welcome User 1</p>",
    },
    {
        To:      "user2@example.com",
        From:    "hello@company.com",
        Subject: "Hello User 2",
        HTML:    "<p>Welcome User 2</p>",
    },
}

response, err := client.Emails.Batch(emails)
if err != nil {
    log.Printf("Error: %v", err)
} else {
    fmt.Printf("Sent %d emails\n", len(response.Emails))
}
```

#### Idempotent Retries

```go
// Idempotent retries: same payload + same key returns the original response
payload := unsent.EmailCreate{
    To:      "hello@acme.com",
    From:    "hello@company.com",
    Subject: "Welcome!",
    HTML:    "<p>Welcome to our service</p>",
}

resp, err := client.Emails.Send(
    payload,
    unsent.WithIdempotencyKey("signup-123"),
)

// Works for batch requests as well
batchPayload := []unsent.EmailBatchItem{
    // ... items
}

resp, err := client.Emails.Batch(
    batchPayload,
    unsent.WithIdempotencyKey("bulk-welcome-1"),
)

// If the same key is reused with a different payload, the API responds with HTTP 409.
```

### Managing Emails

#### Get Email Details

```go
email, err := client.Emails.Get("email_id")
if err != nil {
    log.Printf("Error: %v", err)
} else {
    fmt.Printf("Email status: %s\n", email.Status)
}
```

#### Update Email

```go
response, err := client.Emails.Update("email_id", unsent.EmailUpdate{
    Subject: "Updated subject",
    HTML:    "<p>Updated content</p>",
})
```

#### Cancel Scheduled Email

```go
response, err := client.Emails.Cancel("email_id")
if err != nil {
    log.Printf("Error: %v", err)
} else {
    fmt.Println("Email cancelled successfully")
}
```

#### List Emails

```go
// List all emails with pagination
emails, err := client.Emails.List(unsent.ListEmailsParams{
    Page:  unsent.StringPtr("1"),
    Limit: unsent.StringPtr("50"),
})

if err != nil {
    log.Printf("Error: %v", err)
} else {
    for _, email := range *emails {
        fmt.Printf("Email ID: %s, Status: %s\n", email.ID, email.Status)
    }
}
```

#### Get Bounced Emails

```go
bounces, err := client.Emails.GetBounces(unsent.GetBouncesParams{
    Page:  unsent.Float32Ptr(1.0),
    Limit: unsent.Float32Ptr(20.0),
})

if err != nil {
    log.Printf("Error: %v", err)
} else {
    fmt.Printf("Found %d bounced emails\n", len(*bounces))
}
```

#### Get Spam Complaints

```go
complaints, err := client.Emails.GetComplaints(unsent.GetComplaintsParams{
    Page:  unsent.Float32Ptr(1.0),
    Limit: unsent.Float32Ptr(20.0),
})

if err != nil {
    log.Printf("Error: %v", err)
}
```

#### Get Unsubscribes

```go
unsubscribes, err := client.Emails.GetUnsubscribes(unsent.GetUnsubscribesParams{
    Page:  unsent.Float32Ptr(1.0),
    Limit: unsent.Float32Ptr(20.0),
})

if err != nil {
    log.Printf("Error: %v", err)
}
```

### Managing Contacts

#### Create Contact

```go
contact, err := client.Contacts.Create("contact_book_id", unsent.ContactCreate{
    Email:     "user@example.com",
    FirstName: "John",
    LastName:  "Doe",
    Metadata: map[string]interface{}{
        "company": "Acme Inc",
        "role":    "Developer",
    },
})
```

#### Get Contact

```go
contact, err := client.Contacts.Get("contact_book_id", "contact_id")
```

#### Update Contact

```go
response, err := client.Contacts.Update("contact_book_id", "contact_id", unsent.ContactUpdate{
    FirstName: "Jane",
    Metadata: map[string]interface{}{
        "role": "Senior Developer",
    },
})
```

#### Upsert Contact

```go
// Creates if doesn't exist, updates if exists
contact, err := client.Contacts.Upsert("contact_book_id", "contact_id", unsent.ContactUpsert{
    Email:     "user@example.com",
    FirstName: "John",
    LastName:  "Doe",
})
```

#### Delete Contact

```go
response, err := client.Contacts.Delete("contact_book_id", "contact_id")
```

### Managing Campaigns

#### Create Campaign

```go
campaign, err := client.Campaigns.Create(unsent.CampaignCreate{
    Name:          "Welcome Series",
    Subject:       "Welcome to our service!",
    HTML:          "<p>Thanks for joining us!</p>",
    From:          "welcome@example.com",
    ContactBookID: "cb_1234567890",
})

if err != nil {
    log.Printf("Error: %v", err)
} else {
    fmt.Printf("Campaign created! ID: %s\n", campaign.ID)
}
```

#### Schedule Campaign

```go
response, err := client.Campaigns.Schedule(campaign.ID, unsent.CampaignSchedule{
    ScheduledAt: "2024-12-01T10:00:00Z",
})
```

#### Pause/Resume Campaigns

```go
// Pause a campaign
pauseResp, err := client.Campaigns.Pause("campaign_123")
if err != nil {
    log.Printf("Error: %v", err)
} else {
    fmt.Println("Campaign paused successfully!")
}

// Resume a campaign
resumeResp, err := client.Campaigns.Resume("campaign_123")
if err != nil {
    log.Printf("Error: %v", err)
} else {
    fmt.Println("Campaign resumed successfully!")
}
```

#### Get Campaign Details

```go
campaign, err := client.Campaigns.Get("campaign_id")
if err != nil {
    log.Printf("Error: %v", err)
} else {
    fmt.Printf("Campaign status: %s\n", campaign.Status)
    fmt.Printf("Recipients: %d\n", campaign.Total)
    fmt.Printf("Sent: %d\n", campaign.Sent)
}
```

### Managing Domains

#### List Domains

```go
domains, err := client.Domains.List()
if err != nil {
    log.Printf("Error: %v", err)
} else {
    for _, domain := range domains {
        fmt.Printf("Domain: %s, Status: %s\n", domain.Domain, domain.Status)
    }
}
```

#### Create Domain

```go
domain, err := client.Domains.Create(unsent.DomainCreate{
    Domain: "example.com",
})
```

#### Verify Domain

```go
response, err := client.Domains.Verify(123)
if err != nil {
    log.Printf("Error: %v", err)
} else {
    fmt.Printf("Verification status: %s\n", response.Status)
}
```

#### Get Domain

```go
domain, err := client.Domains.Get(123)
```

### Error Handling

By default, the SDK returns errors for non-2xx responses:

```go
client, err := unsent.NewClient("un_xxxx")
if err != nil {
    log.Fatal(err)
}

email, err := client.Emails.Send(unsent.EmailCreate{
    To:      "invalid-email",
    From:    "hello@company.com",
    Subject: "Test",
    HTML:    "<p>Test</p>",
})

if err != nil {
    if apiErr, ok := err.(*unsent.APIError); ok {
        fmt.Printf("API Error: %s - %s\n", apiErr.Code, apiErr.Message)
    }
}
```

To disable automatic error raising:

```go
client, err := unsent.NewClient("un_xxxx", unsent.WithRaiseOnError(false))
```

### Custom HTTP Client

For advanced use cases, you can provide your own HTTP client:

```go
import "net/http"

httpClient := &http.Client{
    Timeout: 30 * time.Second,
}

client, err := unsent.NewClient("un_xxxx", unsent.WithHTTPClient(httpClient))
```

## API Reference

### Client Methods

- `NewClient(key string, options ...ClientOption)` - Initialize the client
- `WithBaseURL(url string)` - Set custom base URL
- `WithHTTPClient(client *http.Client)` - Set custom HTTP client
- `WithRaiseOnError(raise bool)` - Set error handling behavior

### Email Methods

- `client.Emails.Send(payload, opts...)` - Send an email (alias for Create)
- `client.Emails.Create(payload, opts...)` - Create and send an email
- `client.Emails.Batch(emails, opts...)` - Send multiple emails in batch
- `client.Emails.List(params)` - List emails with optional filters (page, limit, dates)
- `client.Emails.Get(emailID)` - Get email details
- `client.Emails.Update(emailID, payload)` - Update a scheduled email
- `client.Emails.Cancel(emailID)` - Cancel a scheduled email
- `client.Emails.GetBounces(params)` - Get bounced emails with pagination
- `client.Emails.GetComplaints(params)` - Get spam complaints with pagination
- `client.Emails.GetUnsubscribes(params)` - Get unsubscribed emails with pagination

### Contact Methods

- `client.Contacts.List(bookID, params)` - List contacts in a contact book with filters
- `client.Contacts.Create(bookID, payload)` - Create a contact
- `client.Contacts.Get(bookID, contactID)` - Get contact details
- `client.Contacts.Update(bookID, contactID, payload)` - Update a contact
- `client.Contacts.Upsert(bookID, contactID, payload)` - Upsert a contact
- `client.Contacts.Delete(bookID, contactID)` - Delete a contact

### Contact Book Methods

- `client.ContactBooks.List()` - List all contact books
- `client.ContactBooks.Get(id)` - Get contact book details
- `client.ContactBooks.Create(payload)` - Create a contact book
- `client.ContactBooks.Update(id, payload)` - Update a contact book
- `client.ContactBooks.Delete(id)` - Delete a contact book

### Campaign Methods

- `client.Campaigns.List()` - List all campaigns
- `client.Campaigns.Create(payload)` - Create a campaign
- `client.Campaigns.Get(campaignID)` - Get campaign details
- `client.Campaigns.Schedule(campaignID, payload)` - Schedule a campaign
- `client.Campaigns.Pause(campaignID)` - Pause a campaign
- `client.Campaigns.Resume(campaignID)` - Resume a campaign

### Domain Methods

- `client.Domains.List()` - List all domains
- `client.Domains.Create(payload)` - Create a domain
- `client.Domains.Get(domainID)` - Get domain details
- `client.Domains.Verify(domainID)` - Verify a domain
- `client.Domains.Delete(domainID)` - Delete a domain

### Other Resources

- **Analytics**: `Get()`, `GetTimeSeries(params)`, `GetReputation(params)`
- **API Keys**: `List()`, `Create(payload)`, `Delete(id)`
- **Settings**: `Get()`
- **Suppressions**: `List(params)`, `Add(payload)`, `Delete(email)`
- **Templates**: `List()`, `Get(id)`, `Create(payload)`, `Update(id, payload)`, `Delete(id)`
- **Webhooks** *(Future Feature)*: `List()`, `Create(payload)`, `Update(id, payload)`, `Delete(id)` - *Note: Webhook functionality is currently a placeholder for future implementation when webhook support is added to the Unsent API*

## Requirements

- Go 1.19+

## License

MIT

## Support

- [Documentation](https://docs.unsent.dev)
- [GitHub Issues](https://github.com/souravsspace/unsent-go/issues)
- [Discord Community](https://discord.gg/unsent)
