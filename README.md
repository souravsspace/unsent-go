# Unsent Go SDK

Official Go SDK for the [Unsent API](https://unsent.dev) - Send transactional emails with ease.

## Prerequisites

- [Unsent API key](https://app.unsent.dev/dev-settings/api-keys)
- [Verified domain](https://app.unsent.dev/domains)
- Go 1.19 or higher

## Installation

```bash
go get github.com/souravsspace/unsent-go/pkg/unsent
```

## Usage

### Basic Setup

```go
package main

import (
    "fmt"
    "log"
    "encoding/json"
    
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

### Helper Functions

The SDK uses pointer types for optional fields and union types for complex fields like email recipients. Here are recommended helper functions:

```go
// String pointer helper
func stringPtr(s string) *string {
    return &s
}

// Int pointer helper
func intPtr(i int) *int {
    return &i
}

// Float32 pointer helper
func float32Ptr(f float32) *float32 {
    return &f
}

// Bool pointer helper
func boolPtr(b bool) *bool {
    return &b
}

// Email recipient constructor for single emails
func toEmail(email string) unsent.SendEmailJSONBody_To {
    var to unsent.SendEmailJSONBody_To
    json.Unmarshal([]byte(`"`+email+`"`), &to)
    return to
}
```

### Type Helpers

Since the SDK uses generated union types for fields like `To`, `Cc`, and `Bcc`, it is recommended to use a helper function to construct them:

```go
func toEmail(email string) unsent.SendEmailJSONBody_To {
    var to unsent.SendEmailJSONBody_To
    // We marshal the string to JSON, then let the struct handle the unmarshaling if needed
    // or direct internal assignment if possible. 
    // Since the internal field is unexported, we may rely on usage like:
    // This is a placeholder since currently the generated types are strict.
    // For now, please check the 'types.go' for manual construction if available.
    // Assuming standard JSON marshaling works:
    json.Unmarshal([]byte(`"`+email+`"`), &to)
    return to
}

func stringPtr(s string) *string {
    return &s
}
```

### Sending Emails

#### Simple Email

```go
email, err := client.Emails.Send(unsent.SendEmailJSONBody{
    To:      toEmail("hello@acme.com"),
    From:    "hello@company.com",
    Subject: stringPtr("Unsent email"),
    Html:    stringPtr("<p>Unsent is the best email service provider to send emails</p>"),
    Text:    stringPtr("Unsent is the best email service provider to send emails"),
})

if err != nil {
    log.Printf("Error: %v", err)
} else {
    fmt.Printf("Email sent! ID: %s\n", email.EmailID)
}
```

#### Email with Attachments

```go
attachments := []map[string]interface{}{
    {
        "filename": "document.pdf",
        "content":  "base64-encoded-content-here",
    },
}

email, err := client.Emails.Send(unsent.SendEmailJSONBody{
    To:          toEmail("hello@acme.com"),
    From:        "hello@company.com",
    Subject:     stringPtr("Email with attachment"),
    Html:        stringPtr("<p>Please find the attachment below</p>"),
    Attachments: &attachments,
})
```

#### Scheduled Email

```go
import "time"

scheduledTime := time.Now().Add(1 * time.Hour)

email, err := client.Emails.Send(unsent.SendEmailJSONBody{
    To:          toEmail("hello@acme.com"),
    From:        "hello@company.com",
    Subject:     stringPtr("Scheduled email"),
    Html:        stringPtr("<p>This email was scheduled</p>"),
    ScheduledAt: &scheduledTime,
})
```

#### Batch Emails

```go
emails := unsent.SendBatchEmailsJSONBody{
    {
        To:      toEmail("user1@example.com"),
        From:    "hello@company.com",
        Subject: stringPtr("Hello User 1"),
        Html:    stringPtr("<p>Welcome User 1</p>"),
    },
    {
        To:      toEmail("user2@example.com"),
        From:    "hello@company.com",
        Subject: stringPtr("Hello User 2"),
        Html:    stringPtr("<p>Welcome User 2</p>"),
    },
}

response, err := client.Emails.Batch(emails)
if err != nil {
    log.Printf("Error: %v", err)
} else {
    fmt.Printf("Sent %d emails\n", len(response.Data))
}
```

#### Idempotent Retries

```go
// Idempotent retries: same payload + same key returns the original response
payload := unsent.SendEmailJSONBody{
    To:      toEmail("hello@acme.com"),
    From:    "hello@company.com",
    Subject: stringPtr("Welcome!"),
    Html:    stringPtr("<p>Welcome to our service</p>"),
}

resp, err := client.Emails.Send(
    payload,
    unsent.WithIdempotencyKey("signup-123"),
)
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

#### Get Email Events

```go
// Get events for a specific email
events, err := client.Emails.GetEvents("email_id", unsent.GetEmailEventsParams{
    Page: intPtr(1),
    Limit: intPtr(20),
})
```

#### List Emails

```go
// List all emails with pagination
emails, err := client.Emails.List(unsent.ListEmailsParams{
    Page:  stringPtr("1"),
    Limit: stringPtr("50"),
})
```

#### Get Bounces

```go
bounces, err := client.Emails.GetBounces(unsent.GetBouncesParams{
    Page:  float32Ptr(1.0),
    Limit: float32Ptr(20.0),
})
```

### Managing Contacts & Contact Books

#### List Contact Books

```go
books, err := client.ContactBooks.List()
for _, book := range *books {
    fmt.Printf("Book: %s (ID: %s)\n", book.Name, book.ID)
}
```

#### Create Contact

```go
contact, err := client.Contacts.Create("contact_book_id", unsent.CreateContactJSONBody{
    Email:     "user@example.com",
    FirstName: stringPtr("John"),
    LastName:  stringPtr("Doe"),
    Properties: &map[string]string{
        "company": "Acme Inc",
        "role":    "Developer",
    },
})
```

### Managing Campaigns

#### Create Campaign

```go
campaign, err := client.Campaigns.Create(unsent.CreateCampaignJSONBody{
    Name:          "Welcome Series",
    Subject:       "Welcome to our service!",
    Html:          stringPtr("<p>Thanks for joining us!</p>"),
    From:          "welcome@example.com",
    ContactBookId: "cb_1234567890",
})
```

#### Schedule Campaign

```go
schedule := "2024-12-01T10:00:00Z"
response, err := client.Campaigns.Schedule(campaign.ID, unsent.ScheduleCampaignJSONBody{
    ScheduledAt: &schedule,
})
```

### Analytics & Stats

#### Get Overview

```go
analytics, err := client.Analytics.Get()
fmt.Printf("Sent: %d, Opened: %d\n", analytics.Sent, analytics.Opened)
```

#### Get Time Series

```go
timeline, err := client.Analytics.GetTimeSeries(unsent.GetTimeSeriesParams{
    Days: stringPtr("30"),
})
```

```go
timeline, err := client.Analytics.GetTimeSeries(unsent.GetTimeSeriesParams{
    Days: stringPtr(\"30\"),
})
```

#### Get Reputation

```go
reputation, err := client.Analytics.GetReputation(unsent.GetReputationParams{})
if err != nil {
    log.Printf(\"Error: %v\", err)  
} else {
    fmt.Printf(\"Reputation Score: %.2f\\n\", reputation.Score)
}
```

### Events

Retrieve all email events across your account.

#### List Events

```go
page := 1
limit := 50
status := unsent.GetEventsParamsStatusDELIVERED

events, err := client.Events.List(unsent.GetEventsParams{
    Page:   &page,
    Limit:  &limit,
    Status: &status,
})

if err != nil {
    log.Printf(\"Error: %v\", err)
} else {
    fmt.Printf(\"Found %d events\\n\", len(events.Data))
    for _, event := range events.Data {
        fmt.Printf(\"Event: %s - %s\\n\", event.EmailId, event.Status)
    }
}
```

### Activity Feed

Get a combined feed of email events with email details.

```go
page := 1
limit := 20

activity, err := client.Activity.Get(unsent.GetActivityParams{
    Page:  &page,
    Limit: &limit,
})

if err != nil {
    log.Printf(\"Error: %v\", err)
} else {
    fmt.Printf(\"Activity items: %d\\n\", len(activity.Activity))
    for _, item := range activity.Activity {
        if item.Email != nil {
            fmt.Printf(\"Email %s: %s\\n\", item.Email.Subject, item.Status)
        }
    }
}
```

### Metrics

Get performance metrics over different time periods.

```go
period := unsent.GetMetricsParamsPeriodMonth

metrics, err := client.Metrics.Get(unsent.GetMetricsParams{
    Period: &period,
})

if err != nil {
    log.Printf(\"Error: %v\", err)
} else {
    fmt.Printf(\"Open Rate: %.2f%%\\n\", metrics.OpenRate*100)
    fmt.Printf(\"Click Rate: %.2f%%\\n\", metrics.ClickRate*100)
}
```

### Statistics

Retrieve email statistics with date range filtering.

```go
startDate := \"2024-01-01T00:00:00Z\"
endDate := \"2024-12-31T23:59:59Z\"

stats, err := client.Stats.Get(unsent.GetStatsParams{
    StartDate: &startDate,
    EndDate:   &endDate,
})

if err != nil {
    log.Printf(\"Error: %v\", err)
} else {
    fmt.Printf(\"Total Sent: %d\\n\", stats.Sent)
    fmt.Printf(\"Total Delivered: %d\\n\", stats.Delivered)
}
```

### System

Check API health and version information.

#### Health Check

```go
health, err := client.System.Health()
if err != nil {
    log.Printf(\"API is down: %v\", err)
} else {
    fmt.Printf(\"Status: %s, Uptime: %.2fs\\n\", health.Status, health.Uptime)
}
```

#### Get Version

```go
version, err := client.System.Version()
if err != nil {
    log.Printf(\"Error: %v\", err)
} else {
    fmt.Printf(\"API Version: %s\\n\", version.Version)
    fmt.Printf(\"Platform: %s\\n\", version.Platform)
}
```

### Teams

Manage team information.

#### Get Current Team

```go
team, err := client.Teams.Get()
if err != nil {
    log.Printf(\"Error: %v\", err)
} else {
    fmt.Printf(\"Team: %s (ID: %s)\\n\", team.Name, team.ID)
}
```

#### List All Teams

```go
teams, err := client.Teams.List()
if err != nil {
    log.Printf(\"Error: %v\", err)
} else {
    for _, team := range *teams {
        fmt.Printf(\"Team: %s\\n\", team.Name)
    }
}
```

### Email Events

Get events for a specific email.

```go
page := 1
limit := 20

events, err := client.Emails.GetEvents(\"email_id\", unsent.GetEmailEventsParams{
    Page:  &page,
    Limit: &limit,
})

if err != nil {
    log.Printf(\"Error: %v\", err)
} else {
    fmt.Printf(\"Found %d events for this email\\n\", len(events.Data))
}
```

### SDK Resources Summary

The SDK provides clients for all Unsent resources:

- **Activity**: `client.Activity.Get(params)` - Get activity feed with email events and details
- **Analytics**: `client.Analytics.Get()`, `GetTimeSeries(params)`, `GetReputation(params)` - Comprehensive analytics
- **ApiKeys**: `client.ApiKeys.List()`, `Create(payload)`, `Delete(id)` - Manage API keys
- **Campaigns**: `client.Campaigns.List()`, `Create(payload)`, `Schedule(id, payload)`, `Pause(id)`, `Resume(id)` - Campaign management
- **ContactBooks**: `client.ContactBooks.List()`, `Create(payload)`, `Get(id)`, `Update(id, payload)`, `Delete(id)` - Contact book operations
- **Contacts**: `client.Contacts.List(bookId, params)`, `Create(bookId, payload)`, `Get(bookId, id)`, `Update(bookId, id, payload)`, `Delete(bookId, id)` - Contact management
- **Domains**: `client.Domains.List()`, `Create(payload)`, `Get(id)`, `Verify(id)`, `Delete(id)`, `GetAnalytics(id, params)`, `GetStats(id, params)` - Domain operations
- **Emails**: `client.Emails.Send(payload)`, `Batch(payload)`, `List(params)`, `Get(id)`, `Update(id, payload)`, `Cancel(id)`, `GetEvents(id, params)`, `GetBounces(params)`, `GetComplaints(params)`, `GetUnsubscribes(params)` - Email operations
- **Events**: `client.Events.List(params)` - Get all email events
- **Metrics**: `client.Metrics.Get(params)` - Performance metrics
- **Settings**: `client.Settings.Get()` - Account settings
- **Stats**: `client.Stats.Get(params)` - Email statistics
- **Suppressions**: `client.Suppressions.List(params)`, `Add(payload)`, `Delete(email)` - Suppression list management
- **System**: `client.System.Health()`, `Version()` - System information
- **Teams**: `client.Teams.Get()`, `List()` - Team information
- **Templates**: `client.Templates.List()`, `Create(payload)`, `Get(id)`, `Update(id, payload)`, `Delete(id)` - Template operations
- **Webhooks**: `client.Webhooks.List()`, `Create(payload)`, `Get(id)`, `Update(id, payload)`, `Delete(id)`, `Test(id)` - Webhook management

## Error Handling

By default, the SDK returns `*unsent.APIError` for non-2xx responses.

```go
if err != nil {
    fmt.Printf("API Error: %s\n", err.Message)
}
```

To disable automatic error raising:

```go
client, err := unsent.NewClient("un_xxxx", unsent.WithRaiseOnError(false))
```

## License

MIT
