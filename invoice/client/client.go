// This file was auto-generated by Fern from our API Definition.

package client

import (
	bytes "bytes"
	context "context"
	json "encoding/json"
	errors "errors"
	mercoafinancego "github.com/mercoa-finance/go"
	core "github.com/mercoa-finance/go/core"
	invoice "github.com/mercoa-finance/go/invoice"
	approval "github.com/mercoa-finance/go/invoice/approval"
	bulk "github.com/mercoa-finance/go/invoice/bulk"
	comment "github.com/mercoa-finance/go/invoice/comment"
	document "github.com/mercoa-finance/go/invoice/document"
	lineitemclient "github.com/mercoa-finance/go/invoice/lineitem/client"
	paymentlinks "github.com/mercoa-finance/go/invoice/paymentlinks"
	option "github.com/mercoa-finance/go/option"
	io "io"
	http "net/http"
)

type Client struct {
	baseURL string
	caller  *core.Caller
	header  http.Header

	LineItem     *lineitemclient.Client
	Approval     *approval.Client
	Bulk         *bulk.Client
	Comment      *comment.Client
	Document     *document.Client
	PaymentLinks *paymentlinks.Client
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
		header:       options.ToHeader(),
		LineItem:     lineitemclient.NewClient(opts...),
		Approval:     approval.NewClient(opts...),
		Bulk:         bulk.NewClient(opts...),
		Comment:      comment.NewClient(opts...),
		Document:     document.NewClient(opts...),
		PaymentLinks: paymentlinks.NewClient(opts...),
	}
}

// Search invoices for all entities in the organization
func (c *Client) Find(
	ctx context.Context,
	request *invoice.GetAllInvoicesRequest,
	opts ...option.RequestOption,
) (*mercoafinancego.FindInvoiceResponse, error) {
	options := core.NewRequestOptions(opts...)

	baseURL := "https://api.mercoa.com"
	if c.baseURL != "" {
		baseURL = c.baseURL
	}
	if options.BaseURL != "" {
		baseURL = options.BaseURL
	}
	endpointURL := baseURL + "/invoices"

	queryParams, err := core.QueryValues(request)
	if err != nil {
		return nil, err
	}
	if len(queryParams) > 0 {
		endpointURL += "?" + queryParams.Encode()
	}

	headers := core.MergeHeaders(c.header.Clone(), options.ToHeader())

	errorDecoder := func(statusCode int, body io.Reader) error {
		raw, err := io.ReadAll(body)
		if err != nil {
			return err
		}
		apiError := core.NewAPIError(statusCode, errors.New(string(raw)))
		decoder := json.NewDecoder(bytes.NewReader(raw))
		var discriminant struct {
			ErrorName string          `json:"errorName"`
			Content   json.RawMessage `json:"content"`
		}
		if err := decoder.Decode(&discriminant); err != nil {
			return apiError
		}
		switch discriminant.ErrorName {
		case "BadRequest":
			value := new(mercoafinancego.BadRequest)
			value.APIError = apiError
			if err := json.Unmarshal(discriminant.Content, value); err != nil {
				return apiError
			}
			return value
		case "Unauthorized":
			value := new(mercoafinancego.Unauthorized)
			value.APIError = apiError
			if err := json.Unmarshal(discriminant.Content, value); err != nil {
				return apiError
			}
			return value
		case "Forbidden":
			value := new(mercoafinancego.Forbidden)
			value.APIError = apiError
			if err := json.Unmarshal(discriminant.Content, value); err != nil {
				return apiError
			}
			return value
		case "NotFound":
			value := new(mercoafinancego.NotFound)
			value.APIError = apiError
			if err := json.Unmarshal(discriminant.Content, value); err != nil {
				return apiError
			}
			return value
		case "Conflict":
			value := new(mercoafinancego.Conflict)
			value.APIError = apiError
			if err := json.Unmarshal(discriminant.Content, value); err != nil {
				return apiError
			}
			return value
		case "InternalServerError":
			value := new(mercoafinancego.InternalServerError)
			value.APIError = apiError
			if err := json.Unmarshal(discriminant.Content, value); err != nil {
				return apiError
			}
			return value
		case "Unimplemented":
			value := new(mercoafinancego.Unimplemented)
			value.APIError = apiError
			if err := json.Unmarshal(discriminant.Content, value); err != nil {
				return apiError
			}
			return value
		}
		return apiError
	}

	var response *mercoafinancego.FindInvoiceResponse
	if err := c.caller.Call(
		ctx,
		&core.CallParams{
			URL:          endpointURL,
			Method:       http.MethodGet,
			MaxAttempts:  options.MaxAttempts,
			Headers:      headers,
			Client:       options.HTTPClient,
			Response:     &response,
			ErrorDecoder: errorDecoder,
		},
	); err != nil {
		return nil, err
	}
	return response, nil
}

func (c *Client) Create(
	ctx context.Context,
	request *mercoafinancego.InvoiceCreationRequest,
	opts ...option.RequestOption,
) (*mercoafinancego.InvoiceResponse, error) {
	options := core.NewRequestOptions(opts...)

	baseURL := "https://api.mercoa.com"
	if c.baseURL != "" {
		baseURL = c.baseURL
	}
	if options.BaseURL != "" {
		baseURL = options.BaseURL
	}
	endpointURL := baseURL + "/invoice"

	headers := core.MergeHeaders(c.header.Clone(), options.ToHeader())

	errorDecoder := func(statusCode int, body io.Reader) error {
		raw, err := io.ReadAll(body)
		if err != nil {
			return err
		}
		apiError := core.NewAPIError(statusCode, errors.New(string(raw)))
		decoder := json.NewDecoder(bytes.NewReader(raw))
		var discriminant struct {
			ErrorName string          `json:"errorName"`
			Content   json.RawMessage `json:"content"`
		}
		if err := decoder.Decode(&discriminant); err != nil {
			return apiError
		}
		switch discriminant.ErrorName {
		case "BadRequest":
			value := new(mercoafinancego.BadRequest)
			value.APIError = apiError
			if err := json.Unmarshal(discriminant.Content, value); err != nil {
				return apiError
			}
			return value
		case "Unauthorized":
			value := new(mercoafinancego.Unauthorized)
			value.APIError = apiError
			if err := json.Unmarshal(discriminant.Content, value); err != nil {
				return apiError
			}
			return value
		case "Forbidden":
			value := new(mercoafinancego.Forbidden)
			value.APIError = apiError
			if err := json.Unmarshal(discriminant.Content, value); err != nil {
				return apiError
			}
			return value
		case "NotFound":
			value := new(mercoafinancego.NotFound)
			value.APIError = apiError
			if err := json.Unmarshal(discriminant.Content, value); err != nil {
				return apiError
			}
			return value
		case "Conflict":
			value := new(mercoafinancego.Conflict)
			value.APIError = apiError
			if err := json.Unmarshal(discriminant.Content, value); err != nil {
				return apiError
			}
			return value
		case "InternalServerError":
			value := new(mercoafinancego.InternalServerError)
			value.APIError = apiError
			if err := json.Unmarshal(discriminant.Content, value); err != nil {
				return apiError
			}
			return value
		case "Unimplemented":
			value := new(mercoafinancego.Unimplemented)
			value.APIError = apiError
			if err := json.Unmarshal(discriminant.Content, value); err != nil {
				return apiError
			}
			return value
		}
		return apiError
	}

	var response *mercoafinancego.InvoiceResponse
	if err := c.caller.Call(
		ctx,
		&core.CallParams{
			URL:          endpointURL,
			Method:       http.MethodPost,
			MaxAttempts:  options.MaxAttempts,
			Headers:      headers,
			Client:       options.HTTPClient,
			Request:      request,
			Response:     &response,
			ErrorDecoder: errorDecoder,
		},
	); err != nil {
		return nil, err
	}
	return response, nil
}

func (c *Client) Get(
	ctx context.Context,
	// Invoice ID or Invoice ForeignID
	invoiceID mercoafinancego.InvoiceID,
	opts ...option.RequestOption,
) (*mercoafinancego.InvoiceResponse, error) {
	options := core.NewRequestOptions(opts...)

	baseURL := "https://api.mercoa.com"
	if c.baseURL != "" {
		baseURL = c.baseURL
	}
	if options.BaseURL != "" {
		baseURL = options.BaseURL
	}
	endpointURL := core.EncodeURL(baseURL+"/invoice/%v", invoiceID)

	headers := core.MergeHeaders(c.header.Clone(), options.ToHeader())

	errorDecoder := func(statusCode int, body io.Reader) error {
		raw, err := io.ReadAll(body)
		if err != nil {
			return err
		}
		apiError := core.NewAPIError(statusCode, errors.New(string(raw)))
		decoder := json.NewDecoder(bytes.NewReader(raw))
		var discriminant struct {
			ErrorName string          `json:"errorName"`
			Content   json.RawMessage `json:"content"`
		}
		if err := decoder.Decode(&discriminant); err != nil {
			return apiError
		}
		switch discriminant.ErrorName {
		case "BadRequest":
			value := new(mercoafinancego.BadRequest)
			value.APIError = apiError
			if err := json.Unmarshal(discriminant.Content, value); err != nil {
				return apiError
			}
			return value
		case "Unauthorized":
			value := new(mercoafinancego.Unauthorized)
			value.APIError = apiError
			if err := json.Unmarshal(discriminant.Content, value); err != nil {
				return apiError
			}
			return value
		case "Forbidden":
			value := new(mercoafinancego.Forbidden)
			value.APIError = apiError
			if err := json.Unmarshal(discriminant.Content, value); err != nil {
				return apiError
			}
			return value
		case "NotFound":
			value := new(mercoafinancego.NotFound)
			value.APIError = apiError
			if err := json.Unmarshal(discriminant.Content, value); err != nil {
				return apiError
			}
			return value
		case "Conflict":
			value := new(mercoafinancego.Conflict)
			value.APIError = apiError
			if err := json.Unmarshal(discriminant.Content, value); err != nil {
				return apiError
			}
			return value
		case "InternalServerError":
			value := new(mercoafinancego.InternalServerError)
			value.APIError = apiError
			if err := json.Unmarshal(discriminant.Content, value); err != nil {
				return apiError
			}
			return value
		case "Unimplemented":
			value := new(mercoafinancego.Unimplemented)
			value.APIError = apiError
			if err := json.Unmarshal(discriminant.Content, value); err != nil {
				return apiError
			}
			return value
		}
		return apiError
	}

	var response *mercoafinancego.InvoiceResponse
	if err := c.caller.Call(
		ctx,
		&core.CallParams{
			URL:          endpointURL,
			Method:       http.MethodGet,
			MaxAttempts:  options.MaxAttempts,
			Headers:      headers,
			Client:       options.HTTPClient,
			Response:     &response,
			ErrorDecoder: errorDecoder,
		},
	); err != nil {
		return nil, err
	}
	return response, nil
}

func (c *Client) Update(
	ctx context.Context,
	// Invoice ID or Invoice ForeignID
	invoiceID mercoafinancego.InvoiceID,
	request *mercoafinancego.InvoiceUpdateRequest,
	opts ...option.RequestOption,
) (*mercoafinancego.InvoiceResponse, error) {
	options := core.NewRequestOptions(opts...)

	baseURL := "https://api.mercoa.com"
	if c.baseURL != "" {
		baseURL = c.baseURL
	}
	if options.BaseURL != "" {
		baseURL = options.BaseURL
	}
	endpointURL := core.EncodeURL(baseURL+"/invoice/%v", invoiceID)

	headers := core.MergeHeaders(c.header.Clone(), options.ToHeader())

	errorDecoder := func(statusCode int, body io.Reader) error {
		raw, err := io.ReadAll(body)
		if err != nil {
			return err
		}
		apiError := core.NewAPIError(statusCode, errors.New(string(raw)))
		decoder := json.NewDecoder(bytes.NewReader(raw))
		var discriminant struct {
			ErrorName string          `json:"errorName"`
			Content   json.RawMessage `json:"content"`
		}
		if err := decoder.Decode(&discriminant); err != nil {
			return apiError
		}
		switch discriminant.ErrorName {
		case "BadRequest":
			value := new(mercoafinancego.BadRequest)
			value.APIError = apiError
			if err := json.Unmarshal(discriminant.Content, value); err != nil {
				return apiError
			}
			return value
		case "Unauthorized":
			value := new(mercoafinancego.Unauthorized)
			value.APIError = apiError
			if err := json.Unmarshal(discriminant.Content, value); err != nil {
				return apiError
			}
			return value
		case "Forbidden":
			value := new(mercoafinancego.Forbidden)
			value.APIError = apiError
			if err := json.Unmarshal(discriminant.Content, value); err != nil {
				return apiError
			}
			return value
		case "NotFound":
			value := new(mercoafinancego.NotFound)
			value.APIError = apiError
			if err := json.Unmarshal(discriminant.Content, value); err != nil {
				return apiError
			}
			return value
		case "Conflict":
			value := new(mercoafinancego.Conflict)
			value.APIError = apiError
			if err := json.Unmarshal(discriminant.Content, value); err != nil {
				return apiError
			}
			return value
		case "InternalServerError":
			value := new(mercoafinancego.InternalServerError)
			value.APIError = apiError
			if err := json.Unmarshal(discriminant.Content, value); err != nil {
				return apiError
			}
			return value
		case "Unimplemented":
			value := new(mercoafinancego.Unimplemented)
			value.APIError = apiError
			if err := json.Unmarshal(discriminant.Content, value); err != nil {
				return apiError
			}
			return value
		}
		return apiError
	}

	var response *mercoafinancego.InvoiceResponse
	if err := c.caller.Call(
		ctx,
		&core.CallParams{
			URL:          endpointURL,
			Method:       http.MethodPost,
			MaxAttempts:  options.MaxAttempts,
			Headers:      headers,
			Client:       options.HTTPClient,
			Request:      request,
			Response:     &response,
			ErrorDecoder: errorDecoder,
		},
	); err != nil {
		return nil, err
	}
	return response, nil
}

// Only invoices in the UNASSIGNED and DRAFT statuses can be deleted.
func (c *Client) Delete(
	ctx context.Context,
	// Invoice ID or Invoice ForeignID
	invoiceID mercoafinancego.InvoiceID,
	opts ...option.RequestOption,
) error {
	options := core.NewRequestOptions(opts...)

	baseURL := "https://api.mercoa.com"
	if c.baseURL != "" {
		baseURL = c.baseURL
	}
	if options.BaseURL != "" {
		baseURL = options.BaseURL
	}
	endpointURL := core.EncodeURL(baseURL+"/invoice/%v", invoiceID)

	headers := core.MergeHeaders(c.header.Clone(), options.ToHeader())

	errorDecoder := func(statusCode int, body io.Reader) error {
		raw, err := io.ReadAll(body)
		if err != nil {
			return err
		}
		apiError := core.NewAPIError(statusCode, errors.New(string(raw)))
		decoder := json.NewDecoder(bytes.NewReader(raw))
		var discriminant struct {
			ErrorName string          `json:"errorName"`
			Content   json.RawMessage `json:"content"`
		}
		if err := decoder.Decode(&discriminant); err != nil {
			return apiError
		}
		switch discriminant.ErrorName {
		case "BadRequest":
			value := new(mercoafinancego.BadRequest)
			value.APIError = apiError
			if err := json.Unmarshal(discriminant.Content, value); err != nil {
				return apiError
			}
			return value
		case "Unauthorized":
			value := new(mercoafinancego.Unauthorized)
			value.APIError = apiError
			if err := json.Unmarshal(discriminant.Content, value); err != nil {
				return apiError
			}
			return value
		case "Forbidden":
			value := new(mercoafinancego.Forbidden)
			value.APIError = apiError
			if err := json.Unmarshal(discriminant.Content, value); err != nil {
				return apiError
			}
			return value
		case "NotFound":
			value := new(mercoafinancego.NotFound)
			value.APIError = apiError
			if err := json.Unmarshal(discriminant.Content, value); err != nil {
				return apiError
			}
			return value
		case "Conflict":
			value := new(mercoafinancego.Conflict)
			value.APIError = apiError
			if err := json.Unmarshal(discriminant.Content, value); err != nil {
				return apiError
			}
			return value
		case "InternalServerError":
			value := new(mercoafinancego.InternalServerError)
			value.APIError = apiError
			if err := json.Unmarshal(discriminant.Content, value); err != nil {
				return apiError
			}
			return value
		case "Unimplemented":
			value := new(mercoafinancego.Unimplemented)
			value.APIError = apiError
			if err := json.Unmarshal(discriminant.Content, value); err != nil {
				return apiError
			}
			return value
		}
		return apiError
	}

	if err := c.caller.Call(
		ctx,
		&core.CallParams{
			URL:          endpointURL,
			Method:       http.MethodDelete,
			MaxAttempts:  options.MaxAttempts,
			Headers:      headers,
			Client:       options.HTTPClient,
			ErrorDecoder: errorDecoder,
		},
	); err != nil {
		return err
	}
	return nil
}

// Get all events for an invoice
func (c *Client) Events(
	ctx context.Context,
	// Invoice ID or Invoice ForeignID
	invoiceID mercoafinancego.InvoiceID,
	request *invoice.InvoiceInvoiceGetEventsRequest,
	opts ...option.RequestOption,
) (*mercoafinancego.InvoiceEventsResponse, error) {
	options := core.NewRequestOptions(opts...)

	baseURL := "https://api.mercoa.com"
	if c.baseURL != "" {
		baseURL = c.baseURL
	}
	if options.BaseURL != "" {
		baseURL = options.BaseURL
	}
	endpointURL := core.EncodeURL(baseURL+"/invoice/%v/events", invoiceID)

	queryParams, err := core.QueryValues(request)
	if err != nil {
		return nil, err
	}
	if len(queryParams) > 0 {
		endpointURL += "?" + queryParams.Encode()
	}

	headers := core.MergeHeaders(c.header.Clone(), options.ToHeader())

	errorDecoder := func(statusCode int, body io.Reader) error {
		raw, err := io.ReadAll(body)
		if err != nil {
			return err
		}
		apiError := core.NewAPIError(statusCode, errors.New(string(raw)))
		decoder := json.NewDecoder(bytes.NewReader(raw))
		var discriminant struct {
			ErrorName string          `json:"errorName"`
			Content   json.RawMessage `json:"content"`
		}
		if err := decoder.Decode(&discriminant); err != nil {
			return apiError
		}
		switch discriminant.ErrorName {
		case "BadRequest":
			value := new(mercoafinancego.BadRequest)
			value.APIError = apiError
			if err := json.Unmarshal(discriminant.Content, value); err != nil {
				return apiError
			}
			return value
		case "Unauthorized":
			value := new(mercoafinancego.Unauthorized)
			value.APIError = apiError
			if err := json.Unmarshal(discriminant.Content, value); err != nil {
				return apiError
			}
			return value
		case "Forbidden":
			value := new(mercoafinancego.Forbidden)
			value.APIError = apiError
			if err := json.Unmarshal(discriminant.Content, value); err != nil {
				return apiError
			}
			return value
		case "NotFound":
			value := new(mercoafinancego.NotFound)
			value.APIError = apiError
			if err := json.Unmarshal(discriminant.Content, value); err != nil {
				return apiError
			}
			return value
		case "Conflict":
			value := new(mercoafinancego.Conflict)
			value.APIError = apiError
			if err := json.Unmarshal(discriminant.Content, value); err != nil {
				return apiError
			}
			return value
		case "InternalServerError":
			value := new(mercoafinancego.InternalServerError)
			value.APIError = apiError
			if err := json.Unmarshal(discriminant.Content, value); err != nil {
				return apiError
			}
			return value
		case "Unimplemented":
			value := new(mercoafinancego.Unimplemented)
			value.APIError = apiError
			if err := json.Unmarshal(discriminant.Content, value); err != nil {
				return apiError
			}
			return value
		}
		return apiError
	}

	var response *mercoafinancego.InvoiceEventsResponse
	if err := c.caller.Call(
		ctx,
		&core.CallParams{
			URL:          endpointURL,
			Method:       http.MethodGet,
			MaxAttempts:  options.MaxAttempts,
			Headers:      headers,
			Client:       options.HTTPClient,
			Response:     &response,
			ErrorDecoder: errorDecoder,
		},
	); err != nil {
		return nil, err
	}
	return response, nil
}
