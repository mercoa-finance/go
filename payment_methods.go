// This file was auto-generated by Fern from our API Definition.

package mercoa

type FindPaymentMethodsRequest struct {
	// Number of payment methods to return. Limit can range between 1 and 100, and the default is 10.
	Limit *int `json:"-" url:"limit,omitempty"`
	// The ID of the payment method to start after. If not provided, the first page of payment methods will be returned.
	StartingAfter *PaymentMethodID `json:"-" url:"startingAfter,omitempty"`
	// Type of payment method to filter
	Type []*PaymentMethodType `json:"-" url:"type,omitempty"`
	// Entity ID to filter
	EntityID []*EntityID `json:"-" url:"entityId,omitempty"`
}
