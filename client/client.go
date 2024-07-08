// This file was auto-generated by Fern from our API Definition.

package client

import (
	banklookup "github.com/mercoa-finance/go/banklookup"
	core "github.com/mercoa-finance/go/core"
	custompaymentmethodschema "github.com/mercoa-finance/go/custompaymentmethodschema"
	entityclient "github.com/mercoa-finance/go/entity/client"
	entitygroupclient "github.com/mercoa-finance/go/entitygroup/client"
	fees "github.com/mercoa-finance/go/fees"
	invoiceclient "github.com/mercoa-finance/go/invoice/client"
	ocr "github.com/mercoa-finance/go/ocr"
	option "github.com/mercoa-finance/go/option"
	organizationclient "github.com/mercoa-finance/go/organization/client"
	http "net/http"
)

type Client struct {
	baseURL string
	caller  *core.Caller
	header  http.Header

	EntityGroup               *entitygroupclient.Client
	Entity                    *entityclient.Client
	Invoice                   *invoiceclient.Client
	Organization              *organizationclient.Client
	BankLookup                *banklookup.Client
	CustomPaymentMethodSchema *custompaymentmethodschema.Client
	Fees                      *fees.Client
	Ocr                       *ocr.Client
}

func NewClient(opts ...option.RequestOption) *Client {
	options := core.NewRequestOptions(opts...)
	return &Client{
		baseURL: options.BaseURL,
		caller: core.NewCaller(
			&core.CallerParams{
				Client:      options.HTTPClient,
				MaxAttempts: options.MaxAttempts,
			},
		),
		header:                    options.ToHeader(),
		EntityGroup:               entitygroupclient.NewClient(opts...),
		Entity:                    entityclient.NewClient(opts...),
		Invoice:                   invoiceclient.NewClient(opts...),
		Organization:              organizationclient.NewClient(opts...),
		BankLookup:                banklookup.NewClient(opts...),
		CustomPaymentMethodSchema: custompaymentmethodschema.NewClient(opts...),
		Fees:                      fees.NewClient(opts...),
		Ocr:                       ocr.NewClient(opts...),
	}
}
