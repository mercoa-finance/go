// This file was auto-generated by Fern from our API Definition.

package mercoa

import (
	json "encoding/json"
	fmt "fmt"
	internal "github.com/mercoa-finance/go/internal"
	time "time"
)

type EmailLog struct {
	ID        EmailLogID `json:"id" url:"id"`
	Subject   string     `json:"subject" url:"subject"`
	From      string     `json:"from" url:"from"`
	To        string     `json:"to" url:"to"`
	HTMLBody  string     `json:"htmlBody" url:"htmlBody"`
	TextBody  string     `json:"textBody" url:"textBody"`
	CreatedAt time.Time  `json:"createdAt" url:"createdAt"`
	InvoiceID *InvoiceID `json:"invoiceId,omitempty" url:"invoiceId,omitempty"`

	extraProperties map[string]interface{}
	rawJSON         json.RawMessage
}

func (e *EmailLog) GetID() EmailLogID {
	if e == nil {
		return ""
	}
	return e.ID
}

func (e *EmailLog) GetSubject() string {
	if e == nil {
		return ""
	}
	return e.Subject
}

func (e *EmailLog) GetFrom() string {
	if e == nil {
		return ""
	}
	return e.From
}

func (e *EmailLog) GetTo() string {
	if e == nil {
		return ""
	}
	return e.To
}

func (e *EmailLog) GetHTMLBody() string {
	if e == nil {
		return ""
	}
	return e.HTMLBody
}

func (e *EmailLog) GetTextBody() string {
	if e == nil {
		return ""
	}
	return e.TextBody
}

func (e *EmailLog) GetCreatedAt() time.Time {
	if e == nil {
		return time.Time{}
	}
	return e.CreatedAt
}

func (e *EmailLog) GetInvoiceID() *InvoiceID {
	if e == nil {
		return nil
	}
	return e.InvoiceID
}

func (e *EmailLog) GetExtraProperties() map[string]interface{} {
	return e.extraProperties
}

func (e *EmailLog) UnmarshalJSON(data []byte) error {
	type embed EmailLog
	var unmarshaler = struct {
		embed
		CreatedAt *internal.DateTime `json:"createdAt"`
	}{
		embed: embed(*e),
	}
	if err := json.Unmarshal(data, &unmarshaler); err != nil {
		return err
	}
	*e = EmailLog(unmarshaler.embed)
	e.CreatedAt = unmarshaler.CreatedAt.Time()
	extraProperties, err := internal.ExtractExtraProperties(data, *e)
	if err != nil {
		return err
	}
	e.extraProperties = extraProperties
	e.rawJSON = json.RawMessage(data)
	return nil
}

func (e *EmailLog) MarshalJSON() ([]byte, error) {
	type embed EmailLog
	var marshaler = struct {
		embed
		CreatedAt *internal.DateTime `json:"createdAt"`
	}{
		embed:     embed(*e),
		CreatedAt: internal.NewDateTime(e.CreatedAt),
	}
	return json.Marshal(marshaler)
}

func (e *EmailLog) String() string {
	if len(e.rawJSON) > 0 {
		if value, err := internal.StringifyJSON(e.rawJSON); err == nil {
			return value
		}
	}
	if value, err := internal.StringifyJSON(e); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", e)
}

type EmailLogID = string

type EmailLogResponse struct {
	// Total number of logs for the given filters. This value is not limited by the limit parameter. It is provided so that you can determine how many pages of results are available.
	Count int `json:"count" url:"count"`
	// True if there are more logs available for the given filters.
	HasMore bool        `json:"hasMore" url:"hasMore"`
	Data    []*EmailLog `json:"data,omitempty" url:"data,omitempty"`

	extraProperties map[string]interface{}
	rawJSON         json.RawMessage
}

func (e *EmailLogResponse) GetCount() int {
	if e == nil {
		return 0
	}
	return e.Count
}

func (e *EmailLogResponse) GetHasMore() bool {
	if e == nil {
		return false
	}
	return e.HasMore
}

func (e *EmailLogResponse) GetData() []*EmailLog {
	if e == nil {
		return nil
	}
	return e.Data
}

func (e *EmailLogResponse) GetExtraProperties() map[string]interface{} {
	return e.extraProperties
}

func (e *EmailLogResponse) UnmarshalJSON(data []byte) error {
	type unmarshaler EmailLogResponse
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*e = EmailLogResponse(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *e)
	if err != nil {
		return err
	}
	e.extraProperties = extraProperties
	e.rawJSON = json.RawMessage(data)
	return nil
}

func (e *EmailLogResponse) String() string {
	if len(e.rawJSON) > 0 {
		if value, err := internal.StringifyJSON(e.rawJSON); err == nil {
			return value
		}
	}
	if value, err := internal.StringifyJSON(e); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", e)
}
