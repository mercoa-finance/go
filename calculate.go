// This file was auto-generated by Fern from our API Definition.

package mercoa

import (
	json "encoding/json"
	fmt "fmt"
	core "github.com/mercoa-finance/go/core"
	time "time"
)

type CalculateFeesRequest struct {
	// Total amount in major units. If the entered amount has more decimal places than the currency supports, trailing decimals will be truncated.
	Amount float64 `json:"amount" url:"amount"`
	// Currency code for the amount. Defaults to USD.
	Currency *CurrencyCode `json:"currency,omitempty" url:"currency,omitempty"`
	// ID of payment source.
	PaymentSourceID PaymentMethodID `json:"paymentSourceId" url:"paymentSourceId"`
	// ID of payment destination.
	PaymentDestinationID PaymentMethodID `json:"paymentDestinationId" url:"paymentDestinationId"`
	// Options for the payment destination. Depending on the payment destination, this may include things such as check delivery method.
	PaymentDestinationOptions *PaymentDestinationOptions `json:"paymentDestinationOptions,omitempty" url:"paymentDestinationOptions,omitempty"`

	extraProperties map[string]interface{}
	_rawJSON        json.RawMessage
}

func (c *CalculateFeesRequest) GetExtraProperties() map[string]interface{} {
	return c.extraProperties
}

func (c *CalculateFeesRequest) UnmarshalJSON(data []byte) error {
	type unmarshaler CalculateFeesRequest
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*c = CalculateFeesRequest(value)

	extraProperties, err := core.ExtractExtraProperties(data, *c)
	if err != nil {
		return err
	}
	c.extraProperties = extraProperties

	c._rawJSON = json.RawMessage(data)
	return nil
}

func (c *CalculateFeesRequest) String() string {
	if len(c._rawJSON) > 0 {
		if value, err := core.StringifyJSON(c._rawJSON); err == nil {
			return value
		}
	}
	if value, err := core.StringifyJSON(c); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", c)
}

type CalculatePaymentTimingRequest struct {
	EstimatedTiming *EstimatedTiming
	InvoiceTiming   *InvoiceTiming
}

func (c *CalculatePaymentTimingRequest) UnmarshalJSON(data []byte) error {
	valueEstimatedTiming := new(EstimatedTiming)
	if err := json.Unmarshal(data, &valueEstimatedTiming); err == nil {
		c.EstimatedTiming = valueEstimatedTiming
		return nil
	}
	valueInvoiceTiming := new(InvoiceTiming)
	if err := json.Unmarshal(data, &valueInvoiceTiming); err == nil {
		c.InvoiceTiming = valueInvoiceTiming
		return nil
	}
	return fmt.Errorf("%s cannot be deserialized as a %T", data, c)
}

func (c CalculatePaymentTimingRequest) MarshalJSON() ([]byte, error) {
	if c.EstimatedTiming != nil {
		return json.Marshal(c.EstimatedTiming)
	}
	if c.InvoiceTiming != nil {
		return json.Marshal(c.InvoiceTiming)
	}
	return nil, fmt.Errorf("type %T does not include a non-empty union type", c)
}

type CalculatePaymentTimingRequestVisitor interface {
	VisitEstimatedTiming(*EstimatedTiming) error
	VisitInvoiceTiming(*InvoiceTiming) error
}

func (c *CalculatePaymentTimingRequest) Accept(visitor CalculatePaymentTimingRequestVisitor) error {
	if c.EstimatedTiming != nil {
		return visitor.VisitEstimatedTiming(c.EstimatedTiming)
	}
	if c.InvoiceTiming != nil {
		return visitor.VisitInvoiceTiming(c.InvoiceTiming)
	}
	return fmt.Errorf("type %T does not include a non-empty union type", c)
}

type CalculatePaymentTimingResponse struct {
	// Estimated date the payment will be or was processed.
	EstimatedProcessingDate time.Time `json:"estimatedProcessingDate" url:"estimatedProcessingDate"`
	// Number of business days between the estimated processing date and the estimated settlement date. This does not take into account bank holidays or weekends.
	BusinessDays int `json:"businessDays" url:"businessDays"`
	// Estimated payment time in days. This time takes into account bank holidays and weekends.
	EstimatedProcessingTime int `json:"estimatedProcessingTime" url:"estimatedProcessingTime"`
	// Estimated date the payment will be or was settled. This is the same as the request's deductionDate plus the paymentTiming.
	EstimatedSettlementDate time.Time `json:"estimatedSettlementDate" url:"estimatedSettlementDate"`

	extraProperties map[string]interface{}
	_rawJSON        json.RawMessage
}

func (c *CalculatePaymentTimingResponse) GetExtraProperties() map[string]interface{} {
	return c.extraProperties
}

func (c *CalculatePaymentTimingResponse) UnmarshalJSON(data []byte) error {
	type embed CalculatePaymentTimingResponse
	var unmarshaler = struct {
		embed
		EstimatedProcessingDate *core.DateTime `json:"estimatedProcessingDate"`
		EstimatedSettlementDate *core.DateTime `json:"estimatedSettlementDate"`
	}{
		embed: embed(*c),
	}
	if err := json.Unmarshal(data, &unmarshaler); err != nil {
		return err
	}
	*c = CalculatePaymentTimingResponse(unmarshaler.embed)
	c.EstimatedProcessingDate = unmarshaler.EstimatedProcessingDate.Time()
	c.EstimatedSettlementDate = unmarshaler.EstimatedSettlementDate.Time()

	extraProperties, err := core.ExtractExtraProperties(data, *c)
	if err != nil {
		return err
	}
	c.extraProperties = extraProperties

	c._rawJSON = json.RawMessage(data)
	return nil
}

func (c *CalculatePaymentTimingResponse) MarshalJSON() ([]byte, error) {
	type embed CalculatePaymentTimingResponse
	var marshaler = struct {
		embed
		EstimatedProcessingDate *core.DateTime `json:"estimatedProcessingDate"`
		EstimatedSettlementDate *core.DateTime `json:"estimatedSettlementDate"`
	}{
		embed:                   embed(*c),
		EstimatedProcessingDate: core.NewDateTime(c.EstimatedProcessingDate),
		EstimatedSettlementDate: core.NewDateTime(c.EstimatedSettlementDate),
	}
	return json.Marshal(marshaler)
}

func (c *CalculatePaymentTimingResponse) String() string {
	if len(c._rawJSON) > 0 {
		if value, err := core.StringifyJSON(c._rawJSON); err == nil {
			return value
		}
	}
	if value, err := core.StringifyJSON(c); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", c)
}

type EstimatedTiming struct {
	// Date the payment is scheduled to be deducted from the payer's account. Use this field if the payment has not yet been deducted.
	EstimatedDeductionDate *time.Time `json:"estimatedDeductionDate,omitempty" url:"estimatedDeductionDate,omitempty"`
	// Date the payment was processed. Use this field if the payment has already been deducted.
	ProcessedAt *time.Time `json:"processedAt,omitempty" url:"processedAt,omitempty"`
	// ID of payment source.
	PaymentSourceID PaymentMethodID `json:"paymentSourceId" url:"paymentSourceId"`
	// ID of payment destination.
	PaymentDestinationID PaymentMethodID `json:"paymentDestinationId" url:"paymentDestinationId"`
	// Options for the payment destination. Depending on the payment destination, this may include things such as check delivery method.
	PaymentDestinationOptions *PaymentDestinationOptions `json:"paymentDestinationOptions,omitempty" url:"paymentDestinationOptions,omitempty"`

	extraProperties map[string]interface{}
	_rawJSON        json.RawMessage
}

func (e *EstimatedTiming) GetExtraProperties() map[string]interface{} {
	return e.extraProperties
}

func (e *EstimatedTiming) UnmarshalJSON(data []byte) error {
	type embed EstimatedTiming
	var unmarshaler = struct {
		embed
		EstimatedDeductionDate *core.DateTime `json:"estimatedDeductionDate,omitempty"`
		ProcessedAt            *core.DateTime `json:"processedAt,omitempty"`
	}{
		embed: embed(*e),
	}
	if err := json.Unmarshal(data, &unmarshaler); err != nil {
		return err
	}
	*e = EstimatedTiming(unmarshaler.embed)
	e.EstimatedDeductionDate = unmarshaler.EstimatedDeductionDate.TimePtr()
	e.ProcessedAt = unmarshaler.ProcessedAt.TimePtr()

	extraProperties, err := core.ExtractExtraProperties(data, *e)
	if err != nil {
		return err
	}
	e.extraProperties = extraProperties

	e._rawJSON = json.RawMessage(data)
	return nil
}

func (e *EstimatedTiming) MarshalJSON() ([]byte, error) {
	type embed EstimatedTiming
	var marshaler = struct {
		embed
		EstimatedDeductionDate *core.DateTime `json:"estimatedDeductionDate,omitempty"`
		ProcessedAt            *core.DateTime `json:"processedAt,omitempty"`
	}{
		embed:                  embed(*e),
		EstimatedDeductionDate: core.NewOptionalDateTime(e.EstimatedDeductionDate),
		ProcessedAt:            core.NewOptionalDateTime(e.ProcessedAt),
	}
	return json.Marshal(marshaler)
}

func (e *EstimatedTiming) String() string {
	if len(e._rawJSON) > 0 {
		if value, err := core.StringifyJSON(e._rawJSON); err == nil {
			return value
		}
	}
	if value, err := core.StringifyJSON(e); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", e)
}

type InvoiceTiming struct {
	InvoiceID InvoiceID `json:"invoiceId" url:"invoiceId"`

	extraProperties map[string]interface{}
	_rawJSON        json.RawMessage
}

func (i *InvoiceTiming) GetExtraProperties() map[string]interface{} {
	return i.extraProperties
}

func (i *InvoiceTiming) UnmarshalJSON(data []byte) error {
	type unmarshaler InvoiceTiming
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*i = InvoiceTiming(value)

	extraProperties, err := core.ExtractExtraProperties(data, *i)
	if err != nil {
		return err
	}
	i.extraProperties = extraProperties

	i._rawJSON = json.RawMessage(data)
	return nil
}

func (i *InvoiceTiming) String() string {
	if len(i._rawJSON) > 0 {
		if value, err := core.StringifyJSON(i._rawJSON); err == nil {
			return value
		}
	}
	if value, err := core.StringifyJSON(i); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", i)
}
