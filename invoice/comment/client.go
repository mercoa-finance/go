// This file was auto-generated by Fern from our API Definition.

package comment

import (
	bytes "bytes"
	context "context"
	json "encoding/json"
	errors "errors"
	mercoafinancego "github.com/mercoa-finance/go"
	core "github.com/mercoa-finance/go/core"
	option "github.com/mercoa-finance/go/option"
	io "io"
	http "net/http"
)

type Client struct {
	baseURL string
	caller  *core.Caller
	header  http.Header
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
		header: options.ToHeader(),
	}
}

// Get all comments associated with this invoice
func (c *Client) GetAll(
	ctx context.Context,
	// Invoice ID or Invoice ForeignID
	invoiceID mercoafinancego.InvoiceID,
	opts ...option.RequestOption,
) ([]*mercoafinancego.CommentResponse, error) {
	options := core.NewRequestOptions(opts...)

	baseURL := "https://api.mercoa.com"
	if c.baseURL != "" {
		baseURL = c.baseURL
	}
	if options.BaseURL != "" {
		baseURL = options.BaseURL
	}
	endpointURL := core.EncodeURL(baseURL+"/invoice/%v/comments", invoiceID)

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

	var response []*mercoafinancego.CommentResponse
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

// Add a comment to this invoice
func (c *Client) Create(
	ctx context.Context,
	// Invoice ID or Invoice ForeignID
	invoiceID mercoafinancego.InvoiceID,
	request *mercoafinancego.CommentRequest,
	opts ...option.RequestOption,
) (*mercoafinancego.CommentResponse, error) {
	options := core.NewRequestOptions(opts...)

	baseURL := "https://api.mercoa.com"
	if c.baseURL != "" {
		baseURL = c.baseURL
	}
	if options.BaseURL != "" {
		baseURL = options.BaseURL
	}
	endpointURL := core.EncodeURL(baseURL+"/invoice/%v/comment", invoiceID)

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

	var response *mercoafinancego.CommentResponse
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
	commentID mercoafinancego.CommentID,
	opts ...option.RequestOption,
) (*mercoafinancego.CommentResponse, error) {
	options := core.NewRequestOptions(opts...)

	baseURL := "https://api.mercoa.com"
	if c.baseURL != "" {
		baseURL = c.baseURL
	}
	if options.BaseURL != "" {
		baseURL = options.BaseURL
	}
	endpointURL := core.EncodeURL(
		baseURL+"/invoice/%v/comment/%v",
		invoiceID,
		commentID,
	)

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

	var response *mercoafinancego.CommentResponse
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

// Edit a comment on this invoice
func (c *Client) Update(
	ctx context.Context,
	// Invoice ID or Invoice ForeignID
	invoiceID mercoafinancego.InvoiceID,
	commentID mercoafinancego.CommentID,
	request *mercoafinancego.CommentRequest,
	opts ...option.RequestOption,
) (*mercoafinancego.CommentResponse, error) {
	options := core.NewRequestOptions(opts...)

	baseURL := "https://api.mercoa.com"
	if c.baseURL != "" {
		baseURL = c.baseURL
	}
	if options.BaseURL != "" {
		baseURL = options.BaseURL
	}
	endpointURL := core.EncodeURL(
		baseURL+"/invoice/%v/comment/%v",
		invoiceID,
		commentID,
	)

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

	var response *mercoafinancego.CommentResponse
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

// Delete a comment on this invoice
func (c *Client) Delete(
	ctx context.Context,
	// Invoice ID or Invoice ForeignID
	invoiceID mercoafinancego.InvoiceID,
	commentID mercoafinancego.CommentID,
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
	endpointURL := core.EncodeURL(
		baseURL+"/invoice/%v/comment/%v",
		invoiceID,
		commentID,
	)

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
