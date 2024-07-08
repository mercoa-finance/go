// This file was auto-generated by Fern from our API Definition.

package entitygroup

import (
	mercoafinancego "github.com/mercoa-finance/go"
)

type EntityGroupFindRequest struct {
	// The maximum number of results to return. Defaults to 1. Max is 10.
	Limit         *int                           `json:"-" url:"limit,omitempty"`
	StartingAfter *mercoafinancego.EntityGroupID `json:"-" url:"startingAfter,omitempty"`
}
