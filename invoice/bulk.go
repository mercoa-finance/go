// This file was auto-generated by Fern from our API Definition.

package invoice

import (
	json "encoding/json"
	mercoafinancego "github.com/mercoa-finance/go"
)

type BulkInvoiceApprovalRequest struct {
	// If true, webhooks will be emitted for each invoice that is approved. By default, webhooks are not emitted.
	EmitWebhooks *bool                                       `json:"-" url:"emitWebhooks,omitempty"`
	Body         *mercoafinancego.BulkInvoiceApprovalRequest `json:"-" url:"-"`
}

func (b *BulkInvoiceApprovalRequest) UnmarshalJSON(data []byte) error {
	body := new(mercoafinancego.BulkInvoiceApprovalRequest)
	if err := json.Unmarshal(data, &body); err != nil {
		return err
	}
	b.Body = body
	return nil
}

func (b *BulkInvoiceApprovalRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(b.Body)
}

type BulkInvoiceCreationRequest struct {
	// If true, webhooks will be emitted for each invoice that is created. By default, webhooks are not emitted.
	EmitWebhooks *bool                                       `json:"-" url:"emitWebhooks,omitempty"`
	Body         *mercoafinancego.BulkInvoiceCreationRequest `json:"-" url:"-"`
}

func (b *BulkInvoiceCreationRequest) UnmarshalJSON(data []byte) error {
	body := new(mercoafinancego.BulkInvoiceCreationRequest)
	if err := json.Unmarshal(data, &body); err != nil {
		return err
	}
	b.Body = body
	return nil
}

func (b *BulkInvoiceCreationRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(b.Body)
}

type BulkInvoiceUpdateRequest struct {
	// If true, webhooks will be emitted for each invoice that is updated. By default, webhooks are not emitted.
	EmitWebhooks *bool                                     `json:"-" url:"emitWebhooks,omitempty"`
	Body         *mercoafinancego.BulkInvoiceUpdateRequest `json:"-" url:"-"`
}

func (b *BulkInvoiceUpdateRequest) UnmarshalJSON(data []byte) error {
	body := new(mercoafinancego.BulkInvoiceUpdateRequest)
	if err := json.Unmarshal(data, &body); err != nil {
		return err
	}
	b.Body = body
	return nil
}

func (b *BulkInvoiceUpdateRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(b.Body)
}
