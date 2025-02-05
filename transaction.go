// This file was auto-generated by Fern from our API Definition.

package mercoa

import (
	json "encoding/json"
	fmt "fmt"
	core "github.com/mercoa-finance/go/core"
	time "time"
)

type FindTransactionsRequest struct {
	// Filter transactions by the ID or foreign ID of the entity that is the payer or the vendor of the invoice that created the transaction.
	EntityID []*EntityID `json:"-" url:"entityId,omitempty"`
	// CREATED_AT Start date filter.
	StartDate *time.Time `json:"-" url:"startDate,omitempty"`
	// CREATED_AT End date filter.
	EndDate *time.Time `json:"-" url:"endDate,omitempty"`
	// Number of transactions to return. Limit can range between 1 and 100, and the default is 10.
	Limit *int `json:"-" url:"limit,omitempty"`
	// The ID of the transactions to start after. If not provided, the first page of transactions will be returned.
	StartingAfter *TransactionID `json:"-" url:"startingAfter,omitempty"`
	// Find transactions by vendor name, invoice number, check number, or amount. Partial matches are supported.
	Search *string `json:"-" url:"search,omitempty"`
	// Filter transactions by invoice metadata. Each filter will be applied as an AND condition. Duplicate keys will be ignored.
	Metadata []*MetadataFilter `json:"-" url:"metadata,omitempty"`
	// Filter transactions by invoice line item metadata. Each filter will be applied as an AND condition. Duplicate keys will be ignored.
	LineItemMetadata []*MetadataFilter `json:"-" url:"lineItemMetadata,omitempty"`
	// Filter transactions by invoice line item GL account ID. Each filter will be applied as an OR condition. Duplicate keys will be ignored.
	LineItemGlAccountID []*string `json:"-" url:"lineItemGlAccountId,omitempty"`
	// Filter transactions by payer ID or payer foreign ID.
	PayerID []*EntityID `json:"-" url:"payerId,omitempty"`
	// Filter transactions by vendor ID or vendor foreign ID.
	VendorID []*EntityID `json:"-" url:"vendorId,omitempty"`
	// Filter transactions by the ID or foreign ID of the user that created the invoice that created the transaction.
	CreatorUserID []*EntityUserID `json:"-" url:"creatorUserId,omitempty"`
	// Filter transactions by invoice ID. Does not support foreign ID.
	InvoiceID []*InvoiceID `json:"-" url:"invoiceId,omitempty"`
	// Filter transactions by transaction ID.
	TransactionID []*TransactionID `json:"-" url:"transactionId,omitempty"`
	// Transaction status to filter on
	Status []*TransactionStatus `json:"-" url:"status,omitempty"`
	// Filter transactions by transaction type
	TransactionType []*TransactionType `json:"-" url:"transactionType,omitempty"`
}

type FindTransactionsResponse struct {
	Count   int                    `json:"count" url:"count"`
	HasMore bool                   `json:"hasMore" url:"hasMore"`
	Data    []*TransactionResponse `json:"data,omitempty" url:"data,omitempty"`

	extraProperties map[string]interface{}
	_rawJSON        json.RawMessage
}

func (f *FindTransactionsResponse) GetExtraProperties() map[string]interface{} {
	return f.extraProperties
}

func (f *FindTransactionsResponse) UnmarshalJSON(data []byte) error {
	type unmarshaler FindTransactionsResponse
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*f = FindTransactionsResponse(value)

	extraProperties, err := core.ExtractExtraProperties(data, *f)
	if err != nil {
		return err
	}
	f.extraProperties = extraProperties

	f._rawJSON = json.RawMessage(data)
	return nil
}

func (f *FindTransactionsResponse) String() string {
	if len(f._rawJSON) > 0 {
		if value, err := core.StringifyJSON(f._rawJSON); err == nil {
			return value
		}
	}
	if value, err := core.StringifyJSON(f); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", f)
}

type TransactionResponse struct {
	Type                     string
	BankAccountToBankAccount *TransactionResponseBankToBankWithInvoices
	BankAccountToMailedCheck *TransactionResponseBankToMailedCheckWithInvoices
	Custom                   *TransactionResponseCustomWithInvoices
	OffPlatform              *TransactionResponseCustomWithInvoices
}

func (t *TransactionResponse) UnmarshalJSON(data []byte) error {
	var unmarshaler struct {
		Type string `json:"type"`
	}
	if err := json.Unmarshal(data, &unmarshaler); err != nil {
		return err
	}
	t.Type = unmarshaler.Type
	switch unmarshaler.Type {
	case "bankAccountToBankAccount":
		value := new(TransactionResponseBankToBankWithInvoices)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		t.BankAccountToBankAccount = value
	case "bankAccountToMailedCheck":
		value := new(TransactionResponseBankToMailedCheckWithInvoices)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		t.BankAccountToMailedCheck = value
	case "custom":
		value := new(TransactionResponseCustomWithInvoices)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		t.Custom = value
	case "offPlatform":
		value := new(TransactionResponseCustomWithInvoices)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		t.OffPlatform = value
	}
	return nil
}

func (t TransactionResponse) MarshalJSON() ([]byte, error) {
	if t.BankAccountToBankAccount != nil {
		return core.MarshalJSONWithExtraProperty(t.BankAccountToBankAccount, "type", "bankAccountToBankAccount")
	}
	if t.BankAccountToMailedCheck != nil {
		return core.MarshalJSONWithExtraProperty(t.BankAccountToMailedCheck, "type", "bankAccountToMailedCheck")
	}
	if t.Custom != nil {
		return core.MarshalJSONWithExtraProperty(t.Custom, "type", "custom")
	}
	if t.OffPlatform != nil {
		return core.MarshalJSONWithExtraProperty(t.OffPlatform, "type", "offPlatform")
	}
	return nil, fmt.Errorf("type %T does not define a non-empty union type", t)
}

type TransactionResponseVisitor interface {
	VisitBankAccountToBankAccount(*TransactionResponseBankToBankWithInvoices) error
	VisitBankAccountToMailedCheck(*TransactionResponseBankToMailedCheckWithInvoices) error
	VisitCustom(*TransactionResponseCustomWithInvoices) error
	VisitOffPlatform(*TransactionResponseCustomWithInvoices) error
}

func (t *TransactionResponse) Accept(visitor TransactionResponseVisitor) error {
	if t.BankAccountToBankAccount != nil {
		return visitor.VisitBankAccountToBankAccount(t.BankAccountToBankAccount)
	}
	if t.BankAccountToMailedCheck != nil {
		return visitor.VisitBankAccountToMailedCheck(t.BankAccountToMailedCheck)
	}
	if t.Custom != nil {
		return visitor.VisitCustom(t.Custom)
	}
	if t.OffPlatform != nil {
		return visitor.VisitOffPlatform(t.OffPlatform)
	}
	return fmt.Errorf("type %T does not define a non-empty union type", t)
}

type TransactionResponseBankToBankWithInvoices struct {
	ID                        TransactionID              `json:"id" url:"id"`
	Status                    TransactionStatus          `json:"status" url:"status"`
	Amount                    int                        `json:"amount" url:"amount"`
	Currency                  string                     `json:"currency" url:"currency"`
	PayerID                   EntityID                   `json:"payerId" url:"payerId"`
	Payer                     *CounterpartyResponse      `json:"payer,omitempty" url:"payer,omitempty"`
	PaymentSource             *PaymentMethodResponse     `json:"paymentSource,omitempty" url:"paymentSource,omitempty"`
	PaymentSourceID           PaymentMethodID            `json:"paymentSourceId" url:"paymentSourceId"`
	VendorID                  EntityID                   `json:"vendorId" url:"vendorId"`
	Vendor                    *CounterpartyResponse      `json:"vendor,omitempty" url:"vendor,omitempty"`
	PaymentDestination        *PaymentMethodResponse     `json:"paymentDestination,omitempty" url:"paymentDestination,omitempty"`
	PaymentDestinationID      PaymentMethodID            `json:"paymentDestinationId" url:"paymentDestinationId"`
	PaymentDestinationOptions *PaymentDestinationOptions `json:"paymentDestinationOptions,omitempty" url:"paymentDestinationOptions,omitempty"`
	Fees                      *InvoiceFeesResponse       `json:"fees,omitempty" url:"fees,omitempty"`
	CreatedAt                 time.Time                  `json:"createdAt" url:"createdAt"`
	UpdatedAt                 time.Time                  `json:"updatedAt" url:"updatedAt"`
	// If the invoice failed to be paid, this field will be populated with the reason of failure.
	FailureReason *TransactionFailureReason `json:"failureReason,omitempty" url:"failureReason,omitempty"`
	// Invoices associated with this transaction
	Invoices []*InvoiceResponse `json:"invoices,omitempty" url:"invoices,omitempty"`

	extraProperties map[string]interface{}
	_rawJSON        json.RawMessage
}

func (t *TransactionResponseBankToBankWithInvoices) GetExtraProperties() map[string]interface{} {
	return t.extraProperties
}

func (t *TransactionResponseBankToBankWithInvoices) UnmarshalJSON(data []byte) error {
	type embed TransactionResponseBankToBankWithInvoices
	var unmarshaler = struct {
		embed
		CreatedAt *core.DateTime `json:"createdAt"`
		UpdatedAt *core.DateTime `json:"updatedAt"`
	}{
		embed: embed(*t),
	}
	if err := json.Unmarshal(data, &unmarshaler); err != nil {
		return err
	}
	*t = TransactionResponseBankToBankWithInvoices(unmarshaler.embed)
	t.CreatedAt = unmarshaler.CreatedAt.Time()
	t.UpdatedAt = unmarshaler.UpdatedAt.Time()

	extraProperties, err := core.ExtractExtraProperties(data, *t)
	if err != nil {
		return err
	}
	t.extraProperties = extraProperties

	t._rawJSON = json.RawMessage(data)
	return nil
}

func (t *TransactionResponseBankToBankWithInvoices) MarshalJSON() ([]byte, error) {
	type embed TransactionResponseBankToBankWithInvoices
	var marshaler = struct {
		embed
		CreatedAt *core.DateTime `json:"createdAt"`
		UpdatedAt *core.DateTime `json:"updatedAt"`
	}{
		embed:     embed(*t),
		CreatedAt: core.NewDateTime(t.CreatedAt),
		UpdatedAt: core.NewDateTime(t.UpdatedAt),
	}
	return json.Marshal(marshaler)
}

func (t *TransactionResponseBankToBankWithInvoices) String() string {
	if len(t._rawJSON) > 0 {
		if value, err := core.StringifyJSON(t._rawJSON); err == nil {
			return value
		}
	}
	if value, err := core.StringifyJSON(t); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", t)
}

type TransactionResponseBankToMailedCheckWithInvoices struct {
	ID                        TransactionID              `json:"id" url:"id"`
	Status                    TransactionStatus          `json:"status" url:"status"`
	Amount                    int                        `json:"amount" url:"amount"`
	Currency                  string                     `json:"currency" url:"currency"`
	PayerID                   EntityID                   `json:"payerId" url:"payerId"`
	Payer                     *CounterpartyResponse      `json:"payer,omitempty" url:"payer,omitempty"`
	PaymentSource             *PaymentMethodResponse     `json:"paymentSource,omitempty" url:"paymentSource,omitempty"`
	PaymentSourceID           PaymentMethodID            `json:"paymentSourceId" url:"paymentSourceId"`
	VendorID                  EntityID                   `json:"vendorId" url:"vendorId"`
	Vendor                    *CounterpartyResponse      `json:"vendor,omitempty" url:"vendor,omitempty"`
	PaymentDestination        *PaymentMethodResponse     `json:"paymentDestination,omitempty" url:"paymentDestination,omitempty"`
	PaymentDestinationID      PaymentMethodID            `json:"paymentDestinationId" url:"paymentDestinationId"`
	PaymentDestinationOptions *PaymentDestinationOptions `json:"paymentDestinationOptions,omitempty" url:"paymentDestinationOptions,omitempty"`
	Fees                      *InvoiceFeesResponse       `json:"fees,omitempty" url:"fees,omitempty"`
	CreatedAt                 time.Time                  `json:"createdAt" url:"createdAt"`
	UpdatedAt                 time.Time                  `json:"updatedAt" url:"updatedAt"`
	// The number of the check
	CheckNumber int `json:"checkNumber" url:"checkNumber"`
	// Invoices associated with this transaction
	Invoices []*InvoiceResponse `json:"invoices,omitempty" url:"invoices,omitempty"`

	extraProperties map[string]interface{}
	_rawJSON        json.RawMessage
}

func (t *TransactionResponseBankToMailedCheckWithInvoices) GetExtraProperties() map[string]interface{} {
	return t.extraProperties
}

func (t *TransactionResponseBankToMailedCheckWithInvoices) UnmarshalJSON(data []byte) error {
	type embed TransactionResponseBankToMailedCheckWithInvoices
	var unmarshaler = struct {
		embed
		CreatedAt *core.DateTime `json:"createdAt"`
		UpdatedAt *core.DateTime `json:"updatedAt"`
	}{
		embed: embed(*t),
	}
	if err := json.Unmarshal(data, &unmarshaler); err != nil {
		return err
	}
	*t = TransactionResponseBankToMailedCheckWithInvoices(unmarshaler.embed)
	t.CreatedAt = unmarshaler.CreatedAt.Time()
	t.UpdatedAt = unmarshaler.UpdatedAt.Time()

	extraProperties, err := core.ExtractExtraProperties(data, *t)
	if err != nil {
		return err
	}
	t.extraProperties = extraProperties

	t._rawJSON = json.RawMessage(data)
	return nil
}

func (t *TransactionResponseBankToMailedCheckWithInvoices) MarshalJSON() ([]byte, error) {
	type embed TransactionResponseBankToMailedCheckWithInvoices
	var marshaler = struct {
		embed
		CreatedAt *core.DateTime `json:"createdAt"`
		UpdatedAt *core.DateTime `json:"updatedAt"`
	}{
		embed:     embed(*t),
		CreatedAt: core.NewDateTime(t.CreatedAt),
		UpdatedAt: core.NewDateTime(t.UpdatedAt),
	}
	return json.Marshal(marshaler)
}

func (t *TransactionResponseBankToMailedCheckWithInvoices) String() string {
	if len(t._rawJSON) > 0 {
		if value, err := core.StringifyJSON(t._rawJSON); err == nil {
			return value
		}
	}
	if value, err := core.StringifyJSON(t); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", t)
}

type TransactionResponseCustomWithInvoices struct {
	ID                        TransactionID              `json:"id" url:"id"`
	Status                    TransactionStatus          `json:"status" url:"status"`
	Amount                    int                        `json:"amount" url:"amount"`
	Currency                  string                     `json:"currency" url:"currency"`
	PayerID                   EntityID                   `json:"payerId" url:"payerId"`
	Payer                     *CounterpartyResponse      `json:"payer,omitempty" url:"payer,omitempty"`
	PaymentSource             *PaymentMethodResponse     `json:"paymentSource,omitempty" url:"paymentSource,omitempty"`
	PaymentSourceID           PaymentMethodID            `json:"paymentSourceId" url:"paymentSourceId"`
	VendorID                  EntityID                   `json:"vendorId" url:"vendorId"`
	Vendor                    *CounterpartyResponse      `json:"vendor,omitempty" url:"vendor,omitempty"`
	PaymentDestination        *PaymentMethodResponse     `json:"paymentDestination,omitempty" url:"paymentDestination,omitempty"`
	PaymentDestinationID      PaymentMethodID            `json:"paymentDestinationId" url:"paymentDestinationId"`
	PaymentDestinationOptions *PaymentDestinationOptions `json:"paymentDestinationOptions,omitempty" url:"paymentDestinationOptions,omitempty"`
	Fees                      *InvoiceFeesResponse       `json:"fees,omitempty" url:"fees,omitempty"`
	CreatedAt                 time.Time                  `json:"createdAt" url:"createdAt"`
	UpdatedAt                 time.Time                  `json:"updatedAt" url:"updatedAt"`
	// Invoices associated with this transaction
	Invoices []*InvoiceResponse `json:"invoices,omitempty" url:"invoices,omitempty"`

	extraProperties map[string]interface{}
	_rawJSON        json.RawMessage
}

func (t *TransactionResponseCustomWithInvoices) GetExtraProperties() map[string]interface{} {
	return t.extraProperties
}

func (t *TransactionResponseCustomWithInvoices) UnmarshalJSON(data []byte) error {
	type embed TransactionResponseCustomWithInvoices
	var unmarshaler = struct {
		embed
		CreatedAt *core.DateTime `json:"createdAt"`
		UpdatedAt *core.DateTime `json:"updatedAt"`
	}{
		embed: embed(*t),
	}
	if err := json.Unmarshal(data, &unmarshaler); err != nil {
		return err
	}
	*t = TransactionResponseCustomWithInvoices(unmarshaler.embed)
	t.CreatedAt = unmarshaler.CreatedAt.Time()
	t.UpdatedAt = unmarshaler.UpdatedAt.Time()

	extraProperties, err := core.ExtractExtraProperties(data, *t)
	if err != nil {
		return err
	}
	t.extraProperties = extraProperties

	t._rawJSON = json.RawMessage(data)
	return nil
}

func (t *TransactionResponseCustomWithInvoices) MarshalJSON() ([]byte, error) {
	type embed TransactionResponseCustomWithInvoices
	var marshaler = struct {
		embed
		CreatedAt *core.DateTime `json:"createdAt"`
		UpdatedAt *core.DateTime `json:"updatedAt"`
	}{
		embed:     embed(*t),
		CreatedAt: core.NewDateTime(t.CreatedAt),
		UpdatedAt: core.NewDateTime(t.UpdatedAt),
	}
	return json.Marshal(marshaler)
}

func (t *TransactionResponseCustomWithInvoices) String() string {
	if len(t._rawJSON) > 0 {
		if value, err := core.StringifyJSON(t._rawJSON); err == nil {
			return value
		}
	}
	if value, err := core.StringifyJSON(t); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", t)
}
