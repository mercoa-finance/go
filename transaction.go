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
	// Find transactions by vendor name, invoice number, or amount. Partial matches are supported.
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
	// Filter transactions by invoice ID.
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

type TransactionStatus string

const (
	TransactionStatusCreated   TransactionStatus = "CREATED"
	TransactionStatusPending   TransactionStatus = "PENDING"
	TransactionStatusCompleted TransactionStatus = "COMPLETED"
	TransactionStatusFailed    TransactionStatus = "FAILED"
	TransactionStatusReversed  TransactionStatus = "REVERSED"
	TransactionStatusQueued    TransactionStatus = "QUEUED"
	TransactionStatusCanceled  TransactionStatus = "CANCELED"
)

func NewTransactionStatusFromString(s string) (TransactionStatus, error) {
	switch s {
	case "CREATED":
		return TransactionStatusCreated, nil
	case "PENDING":
		return TransactionStatusPending, nil
	case "COMPLETED":
		return TransactionStatusCompleted, nil
	case "FAILED":
		return TransactionStatusFailed, nil
	case "REVERSED":
		return TransactionStatusReversed, nil
	case "QUEUED":
		return TransactionStatusQueued, nil
	case "CANCELED":
		return TransactionStatusCanceled, nil
	}
	var t TransactionStatus
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (t TransactionStatus) Ptr() *TransactionStatus {
	return &t
}

type TransactionType string

const (
	TransactionTypeBankAccountToBankAccount TransactionType = "bankAccountToBankAccount"
	TransactionTypeBankAccountToMailedCheck TransactionType = "bankAccountToMailedCheck"
	TransactionTypeCustom                   TransactionType = "custom"
	TransactionTypeOffPlatform              TransactionType = "offPlatform"
)

func NewTransactionTypeFromString(s string) (TransactionType, error) {
	switch s {
	case "bankAccountToBankAccount":
		return TransactionTypeBankAccountToBankAccount, nil
	case "bankAccountToMailedCheck":
		return TransactionTypeBankAccountToMailedCheck, nil
	case "custom":
		return TransactionTypeCustom, nil
	case "offPlatform":
		return TransactionTypeOffPlatform, nil
	}
	var t TransactionType
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (t TransactionType) Ptr() *TransactionType {
	return &t
}
