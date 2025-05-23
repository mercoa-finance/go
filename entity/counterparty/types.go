// This file was auto-generated by Fern from our API Definition.

package counterparty

import (
	mercoafinancego "github.com/mercoa-finance/go"
)

type FindPayeeCounterpartiesRequest struct {
	// Use search instead. Deprecated. Filter counterparties by name. Partial matches are supported.
	Name *string `json:"-" url:"name,omitempty"`
	// Filter counterparties by name or email. Partial matches are supported.
	Search *string `json:"-" url:"search,omitempty"`
	// Filter by network type. By default, only ENTITY counterparties are returned.
	NetworkType []*mercoafinancego.CounterpartyNetworkType `json:"-" url:"networkType,omitempty"`
	// If true, will include counterparty payment methods as part of the response
	PaymentMethods *bool `json:"-" url:"paymentMethods,omitempty"`
	// If true, will include counterparty invoice metrics as part of the response
	InvoiceMetrics *bool `json:"-" url:"invoiceMetrics,omitempty"`
	// Filter by counterparty ids (Foreign ID is supported)
	CounterpartyID []*mercoafinancego.EntityID `json:"-" url:"counterpartyId,omitempty"`
	// Filter counterparties by simple key/value metadata. Each filter will be applied as an AND condition. Duplicate keys will be ignored.
	Metadata []*mercoafinancego.MetadataFilter `json:"-" url:"metadata,omitempty"`
	// If true, will return simple key/value metadata for the counterparties. For more complex metadata, use the Metadata API.
	ReturnMetadata []*string `json:"-" url:"returnMetadata,omitempty"`
	// Number of counterparties to return. Limit can range between 1 and 100, and the default is 10.
	Limit *int `json:"-" url:"limit,omitempty"`
	// The ID of the counterparties to start after. If not provided, the first page of counterparties will be returned.
	StartingAfter *mercoafinancego.EntityID `json:"-" url:"startingAfter,omitempty"`
}

type FindPayorCounterpartiesRequest struct {
	// Use search instead. Deprecated. Filter counterparties by name. Partial matches are supported.
	Name *string `json:"-" url:"name,omitempty"`
	// Filter counterparties by name or email. Partial matches are supported.
	Search *string `json:"-" url:"search,omitempty"`
	// Filter by network type. By default, only ENTITY counterparties are returned.
	NetworkType []*mercoafinancego.CounterpartyNetworkType `json:"-" url:"networkType,omitempty"`
	// If true, will include counterparty payment methods as part of the response
	PaymentMethods *bool `json:"-" url:"paymentMethods,omitempty"`
	// If true, will include counterparty invoice metrics as part of the response
	InvoiceMetrics *bool `json:"-" url:"invoiceMetrics,omitempty"`
	// Filter by counterparty ids (Foreign ID is supported)
	CounterpartyID []*mercoafinancego.EntityID `json:"-" url:"counterpartyId,omitempty"`
	// Filter counterparties by simple key/value metadata. Each filter will be applied as an AND condition. Duplicate keys will be ignored.
	Metadata []*mercoafinancego.MetadataFilter `json:"-" url:"metadata,omitempty"`
	// If true, will return simple key/value metadata for the counterparties. For more complex metadata, use the Metadata API.
	ReturnMetadata []*string `json:"-" url:"returnMetadata,omitempty"`
	// Number of counterparties to return. Limit can range between 1 and 100, and the default is 10.
	Limit *int `json:"-" url:"limit,omitempty"`
	// The ID of the counterparties to start after. If not provided, the first page of counterparties will be returned.
	StartingAfter *mercoafinancego.EntityID `json:"-" url:"startingAfter,omitempty"`
}
