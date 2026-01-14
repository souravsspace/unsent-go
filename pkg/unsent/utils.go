package unsent

import (
	"fmt"
	"strings"
	"time"
)

// buildQueryParams constructs a query string from a map of parameters
func buildQueryParams(params map[string]interface{}) string {
	values := make([]string, 0)
	for key, val := range params {
		if val != nil {
			// Handle different types
			switch v := val.(type) {
			case *int:
				if v != nil {
					values = append(values, fmt.Sprintf("%s=%d", key, *v))
				}
			case *string:
				if v != nil {
					values = append(values, fmt.Sprintf("%s=%s", key, *v))
				}
			case *float32:
				if v != nil {
					values = append(values, fmt.Sprintf("%s=%f", key, *v))
				}
			case *GetEventsParamsStatus:
				if v != nil {
					values = append(values, fmt.Sprintf("%s=%s", key, string(*v)))
				}
			case *GetMetricsParamsPeriod:
				if v != nil {
					values = append(values, fmt.Sprintf("%s=%s", key, string(*v)))
				}
			case *GetDomainAnalyticsParamsPeriod:
				if v != nil {
					values = append(values, fmt.Sprintf("%s=%s", key, string(*v)))
				}
			case *GetEmailEventsParamsStatus:
				if v != nil {
					values = append(values, fmt.Sprintf("%s=%s", key, string(*v)))
				}
			case *time.Time:
				if v != nil {
					values = append(values, fmt.Sprintf("%s=%s", key, v.Format(time.RFC3339)))
				}
			}
		}
	}
	if len(values) == 0 {
		return ""
	}
	return strings.Join(values, "&")
}
