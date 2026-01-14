package unsent

import "encoding/json"

// MakeSendEmailJSONBodyTo creates a SendEmailJSONBody_To from a string or []string
func MakeSendEmailJSONBodyTo(v interface{}) SendEmailJSONBody_To {
	b, _ := json.Marshal(v)
	return SendEmailJSONBody_To{union: b}
}

// MakeBatchEmailTo creates a SendBatchEmailsJSONBody_To from a string or []string
func MakeBatchEmailTo(v interface{}) SendBatchEmailsJSONBody_To {
	b, _ := json.Marshal(v)
	return SendBatchEmailsJSONBody_To{union: b}
}

// MarshalJSON marshals the SendEmailJSONBody_To union
func (t SendEmailJSONBody_To) MarshalJSON() ([]byte, error) {
	if t.union == nil {
		return []byte("null"), nil
	}
	return t.union, nil
}

// UnmarshalJSON unmarshals the SendEmailJSONBody_To union
func (t *SendEmailJSONBody_To) UnmarshalJSON(data []byte) error {
	t.union = make(json.RawMessage, len(data))
	copy(t.union, data)
	return nil
}

// MarshalJSON marshals the SendBatchEmailsJSONBody_To union
func (t SendBatchEmailsJSONBody_To) MarshalJSON() ([]byte, error) {
	if t.union == nil {
		return []byte("null"), nil
	}
	return t.union, nil
}

// UnmarshalJSON unmarshals the SendBatchEmailsJSONBody_To union
func (t *SendBatchEmailsJSONBody_To) UnmarshalJSON(data []byte) error {
	t.union = make(json.RawMessage, len(data))
	copy(t.union, data)
	return nil
}
