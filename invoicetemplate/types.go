// This file was auto-generated by Fern from our API Definition.

package invoicetemplate

import (
	mercoafinancego "github.com/mercoa-finance/go"
	time "time"
)

type GetAllInvoiceTemplatesRequest struct {
	// Filter invoice templates by the ID or foreign ID of the entity that is the payer or the vendor of the invoice template.
	EntityID []*mercoafinancego.EntityID `json:"-" url:"entityId,omitempty"`
	// Start date filter. Defaults to CREATED_AT unless specified the dateType is specified
	StartDate *time.Time `json:"-" url:"startDate,omitempty"`
	// End date filter. Defaults to CREATED_AT unless specified the dateType is specified
	EndDate *time.Time `json:"-" url:"endDate,omitempty"`
	// Type of date to filter by if startDate and endDate filters are provided. Defaults to CREATED_AT.
	DateType *mercoafinancego.InvoiceDateFilter `json:"-" url:"dateType,omitempty"`
	// Field to order invoice templates by. Defaults to CREATED_AT.
	OrderBy *mercoafinancego.InvoiceOrderByField `json:"-" url:"orderBy,omitempty"`
	// Direction to order invoice templates by. Defaults to asc.
	OrderDirection *mercoafinancego.OrderDirection `json:"-" url:"orderDirection,omitempty"`
	// Number of invoice templates to return. Limit can range between 1 and 100, and the default is 10.
	Limit *int `json:"-" url:"limit,omitempty"`
	// The ID of the invoice template to start after. If not provided, the first page of invoice templates will be returned.
	StartingAfter *mercoafinancego.InvoiceTemplateID `json:"-" url:"startingAfter,omitempty"`
	// Find invoice templates by vendor name, check number, invoice number, or amount. Partial matches are supported.
	Search *string `json:"-" url:"search,omitempty"`
	// Filter invoice templates by metadata. Each filter will be applied as an AND condition. Duplicate keys will be ignored.
	Metadata []*mercoafinancego.MetadataFilter `json:"-" url:"metadata,omitempty"`
	// Filter invoice templates by line item metadata. Each filter will be applied as an AND condition. Duplicate keys will be ignored.
	LineItemMetadata []*mercoafinancego.MetadataFilter `json:"-" url:"lineItemMetadata,omitempty"`
	// Filter invoice templates by line item GL account ID. Each filter will be applied as an OR condition. Duplicate keys will be ignored.
	LineItemGlAccountID []*string `json:"-" url:"lineItemGlAccountId,omitempty"`
	// Filter invoice templates by payer ID or payer foreign ID.
	PayerID []*mercoafinancego.EntityID `json:"-" url:"payerId,omitempty"`
	// Filter invoice templates by vendor ID or vendor foreign ID.
	VendorID []*mercoafinancego.EntityID `json:"-" url:"vendorId,omitempty"`
	// Filter invoices by the ID or foreign ID of the user that created the invoice.
	CreatorUserID []*mercoafinancego.EntityUserID `json:"-" url:"creatorUserId,omitempty"`
	// Filter invoice templates by assigned approver user ID.
	ApproverID []*mercoafinancego.EntityUserID `json:"-" url:"approverId,omitempty"`
	// Filter invoice templates by approver action. Needs to be used with approverId. For example, if you want to find all invoice templates that have been approved by a specific user, you would use approverId and approverAction=APPROVE.
	ApproverAction []*mercoafinancego.ApproverAction `json:"-" url:"approverAction,omitempty"`
	// Filter invoice templates by invoice ID.
	InvoiceID []*mercoafinancego.InvoiceID `json:"-" url:"invoiceId,omitempty"`
	// Invoice status to filter on
	Status []*mercoafinancego.InvoiceStatus `json:"-" url:"status,omitempty"`
	// Filter invoice templates by recurring status
	PaymentType []mercoafinancego.PaymentType `json:"-" url:"paymentType,omitempty"`
}
