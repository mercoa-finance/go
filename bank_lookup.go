// This file was auto-generated by Fern from our API Definition.

package mercoa

import (
	json "encoding/json"
	fmt "fmt"
	core "github.com/mercoa-finance/go/core"
)

type BankLookupRequest struct {
	// Routing number to validate
	RoutingNumber string `json:"-" url:"routingNumber"`
}

type BankLookupResponse struct {
	BankName    string       `json:"bankName" url:"bankName"`
	BankAddress *BankAddress `json:"bankAddress,omitempty" url:"bankAddress,omitempty"`

	extraProperties map[string]interface{}
	_rawJSON        json.RawMessage
}

func (b *BankLookupResponse) GetExtraProperties() map[string]interface{} {
	return b.extraProperties
}

func (b *BankLookupResponse) UnmarshalJSON(data []byte) error {
	type unmarshaler BankLookupResponse
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*b = BankLookupResponse(value)

	extraProperties, err := core.ExtractExtraProperties(data, *b)
	if err != nil {
		return err
	}
	b.extraProperties = extraProperties

	b._rawJSON = json.RawMessage(data)
	return nil
}

func (b *BankLookupResponse) String() string {
	if len(b._rawJSON) > 0 {
		if value, err := core.StringifyJSON(b._rawJSON); err == nil {
			return value
		}
	}
	if value, err := core.StringifyJSON(b); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", b)
}
