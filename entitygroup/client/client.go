// This file was auto-generated by Fern from our API Definition.

package client

import (
	bytes "bytes"
	context "context"
	json "encoding/json"
	errors "errors"
	mercoafinancego "github.com/mercoa-finance/go"
	core "github.com/mercoa-finance/go/core"
	entitygroup "github.com/mercoa-finance/go/entitygroup"
	counterpartyclient "github.com/mercoa-finance/go/entitygroup/counterparty/client"
	invoice "github.com/mercoa-finance/go/entitygroup/invoice"
	userclient "github.com/mercoa-finance/go/entitygroup/user/client"
	internal "github.com/mercoa-finance/go/internal"
	option "github.com/mercoa-finance/go/option"
	io "io"
	http "net/http"
)

type Client struct {
	baseURL string
	caller  *internal.Caller
	header  http.Header

	Counterparty *counterpartyclient.Client
	User         *userclient.Client
	Invoice      *invoice.Client
}

func NewClient(opts ...option.RequestOption) *Client {
	options := core.NewRequestOptions(opts...)
	return &Client{
		baseURL: options.BaseURL,
		caller: internal.NewCaller(
			&internal.CallerParams{
				Client:      options.HTTPClient,
				MaxAttempts: options.MaxAttempts,
			},
		),
		header:       options.ToHeader(),
		Counterparty: counterpartyclient.NewClient(opts...),
		User:         userclient.NewClient(opts...),
		Invoice:      invoice.NewClient(opts...),
	}
}

// Get all entity groups. If using a JWT, will return all groups the entity is part of. If using an API key, will return all groups for the organization.
func (c *Client) GetAll(
	ctx context.Context,
	request *entitygroup.EntityGroupFindRequest,
	opts ...option.RequestOption,
) (*mercoafinancego.EntityGroupFindResponse, error) {
	options := core.NewRequestOptions(opts...)
	baseURL := internal.ResolveBaseURL(
		options.BaseURL,
		c.baseURL,
		"https://api.mercoa.com",
	)
	endpointURL := baseURL + "/entityGroups"
	queryParams, err := internal.QueryValues(request)
	if err != nil {
		return nil, err
	}
	if len(queryParams) > 0 {
		endpointURL += "?" + queryParams.Encode()
	}
	headers := internal.MergeHeaders(
		c.header.Clone(),
		options.ToHeader(),
	)
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

	var response *mercoafinancego.EntityGroupFindResponse
	if err := c.caller.Call(
		ctx,
		&internal.CallParams{
			URL:             endpointURL,
			Method:          http.MethodGet,
			Headers:         headers,
			MaxAttempts:     options.MaxAttempts,
			BodyProperties:  options.BodyProperties,
			QueryParameters: options.QueryParameters,
			Client:          options.HTTPClient,
			Response:        &response,
			ErrorDecoder:    errorDecoder,
		},
	); err != nil {
		return nil, err
	}
	return response, nil
}

// Create an entity group
func (c *Client) Create(
	ctx context.Context,
	request *mercoafinancego.EntityGroupCreateRequest,
	opts ...option.RequestOption,
) (*mercoafinancego.EntityGroupResponse, error) {
	options := core.NewRequestOptions(opts...)
	baseURL := internal.ResolveBaseURL(
		options.BaseURL,
		c.baseURL,
		"https://api.mercoa.com",
	)
	endpointURL := baseURL + "/entityGroup"
	headers := internal.MergeHeaders(
		c.header.Clone(),
		options.ToHeader(),
	)
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

	var response *mercoafinancego.EntityGroupResponse
	if err := c.caller.Call(
		ctx,
		&internal.CallParams{
			URL:             endpointURL,
			Method:          http.MethodPost,
			Headers:         headers,
			MaxAttempts:     options.MaxAttempts,
			BodyProperties:  options.BodyProperties,
			QueryParameters: options.QueryParameters,
			Client:          options.HTTPClient,
			Request:         request,
			Response:        &response,
			ErrorDecoder:    errorDecoder,
		},
	); err != nil {
		return nil, err
	}
	return response, nil
}

// Get an entity group
func (c *Client) Get(
	ctx context.Context,
	// Entity Group ID or Entity Group ForeignID
	entityGroupID mercoafinancego.EntityGroupID,
	request *entitygroup.EntityGroupGetRequest,
	opts ...option.RequestOption,
) (*mercoafinancego.EntityGroupResponse, error) {
	options := core.NewRequestOptions(opts...)
	baseURL := internal.ResolveBaseURL(
		options.BaseURL,
		c.baseURL,
		"https://api.mercoa.com",
	)
	endpointURL := internal.EncodeURL(
		baseURL+"/entityGroup/%v",
		entityGroupID,
	)
	queryParams, err := internal.QueryValues(request)
	if err != nil {
		return nil, err
	}
	if len(queryParams) > 0 {
		endpointURL += "?" + queryParams.Encode()
	}
	headers := internal.MergeHeaders(
		c.header.Clone(),
		options.ToHeader(),
	)
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

	var response *mercoafinancego.EntityGroupResponse
	if err := c.caller.Call(
		ctx,
		&internal.CallParams{
			URL:             endpointURL,
			Method:          http.MethodGet,
			Headers:         headers,
			MaxAttempts:     options.MaxAttempts,
			BodyProperties:  options.BodyProperties,
			QueryParameters: options.QueryParameters,
			Client:          options.HTTPClient,
			Response:        &response,
			ErrorDecoder:    errorDecoder,
		},
	); err != nil {
		return nil, err
	}
	return response, nil
}

// Update an entity group
func (c *Client) Update(
	ctx context.Context,
	// Entity Group ID or Entity Group ForeignID
	entityGroupID mercoafinancego.EntityGroupID,
	request *mercoafinancego.EntityGroupUpdateRequest,
	opts ...option.RequestOption,
) (*mercoafinancego.EntityGroupResponse, error) {
	options := core.NewRequestOptions(opts...)
	baseURL := internal.ResolveBaseURL(
		options.BaseURL,
		c.baseURL,
		"https://api.mercoa.com",
	)
	endpointURL := internal.EncodeURL(
		baseURL+"/entityGroup/%v",
		entityGroupID,
	)
	headers := internal.MergeHeaders(
		c.header.Clone(),
		options.ToHeader(),
	)
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

	var response *mercoafinancego.EntityGroupResponse
	if err := c.caller.Call(
		ctx,
		&internal.CallParams{
			URL:             endpointURL,
			Method:          http.MethodPost,
			Headers:         headers,
			MaxAttempts:     options.MaxAttempts,
			BodyProperties:  options.BodyProperties,
			QueryParameters: options.QueryParameters,
			Client:          options.HTTPClient,
			Request:         request,
			Response:        &response,
			ErrorDecoder:    errorDecoder,
		},
	); err != nil {
		return nil, err
	}
	return response, nil
}

// Delete an entity group
func (c *Client) Delete(
	ctx context.Context,
	// Entity Group ID or Entity Group ForeignID
	entityGroupID mercoafinancego.EntityGroupID,
	opts ...option.RequestOption,
) error {
	options := core.NewRequestOptions(opts...)
	baseURL := internal.ResolveBaseURL(
		options.BaseURL,
		c.baseURL,
		"https://api.mercoa.com",
	)
	endpointURL := internal.EncodeURL(
		baseURL+"/entityGroup/%v",
		entityGroupID,
	)
	headers := internal.MergeHeaders(
		c.header.Clone(),
		options.ToHeader(),
	)
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
		&internal.CallParams{
			URL:             endpointURL,
			Method:          http.MethodDelete,
			Headers:         headers,
			MaxAttempts:     options.MaxAttempts,
			BodyProperties:  options.BodyProperties,
			QueryParameters: options.QueryParameters,
			Client:          options.HTTPClient,
			ErrorDecoder:    errorDecoder,
		},
	); err != nil {
		return err
	}
	return nil
}

// Generate a JWT token for an entity group with the given options. This token can be used to authenticate to any entity in the entity group in the Mercoa API and iFrame.
//
// <Warning>We recommend using [this endpoint](/api-reference/entity-group/user/get-token). This will enable features such as approvals and comments.</Warning>
func (c *Client) GetToken(
	ctx context.Context,
	// Entity Group ID or Entity Group ForeignID
	entityGroupID mercoafinancego.EntityGroupID,
	request *mercoafinancego.TokenGenerationOptions,
	opts ...option.RequestOption,
) (string, error) {
	options := core.NewRequestOptions(opts...)
	baseURL := internal.ResolveBaseURL(
		options.BaseURL,
		c.baseURL,
		"https://api.mercoa.com",
	)
	endpointURL := internal.EncodeURL(
		baseURL+"/entityGroup/%v/token",
		entityGroupID,
	)
	headers := internal.MergeHeaders(
		c.header.Clone(),
		options.ToHeader(),
	)
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

	var response string
	if err := c.caller.Call(
		ctx,
		&internal.CallParams{
			URL:             endpointURL,
			Method:          http.MethodPost,
			Headers:         headers,
			MaxAttempts:     options.MaxAttempts,
			BodyProperties:  options.BodyProperties,
			QueryParameters: options.QueryParameters,
			Client:          options.HTTPClient,
			Request:         request,
			Response:        &response,
			ErrorDecoder:    errorDecoder,
		},
	); err != nil {
		return "", err
	}
	return response, nil
}

// Add entities to an entity group
func (c *Client) AddEntities(
	ctx context.Context,
	// Entity Group ID or Entity Group ForeignID
	entityGroupID mercoafinancego.EntityGroupID,
	request *mercoafinancego.EntityGroupAddEntitiesRequest,
	opts ...option.RequestOption,
) (*mercoafinancego.EntityGroupResponse, error) {
	options := core.NewRequestOptions(opts...)
	baseURL := internal.ResolveBaseURL(
		options.BaseURL,
		c.baseURL,
		"https://api.mercoa.com",
	)
	endpointURL := internal.EncodeURL(
		baseURL+"/entityGroup/%v/addEntities",
		entityGroupID,
	)
	headers := internal.MergeHeaders(
		c.header.Clone(),
		options.ToHeader(),
	)
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

	var response *mercoafinancego.EntityGroupResponse
	if err := c.caller.Call(
		ctx,
		&internal.CallParams{
			URL:             endpointURL,
			Method:          http.MethodPost,
			Headers:         headers,
			MaxAttempts:     options.MaxAttempts,
			BodyProperties:  options.BodyProperties,
			QueryParameters: options.QueryParameters,
			Client:          options.HTTPClient,
			Request:         request,
			Response:        &response,
			ErrorDecoder:    errorDecoder,
		},
	); err != nil {
		return nil, err
	}
	return response, nil
}

// Remove entities from an entity group
func (c *Client) RemoveEntities(
	ctx context.Context,
	// Entity Group ID or Entity Group ForeignID
	entityGroupID mercoafinancego.EntityGroupID,
	request *mercoafinancego.EntityGroupRemoveEntitiesRequest,
	opts ...option.RequestOption,
) (*mercoafinancego.EntityGroupResponse, error) {
	options := core.NewRequestOptions(opts...)
	baseURL := internal.ResolveBaseURL(
		options.BaseURL,
		c.baseURL,
		"https://api.mercoa.com",
	)
	endpointURL := internal.EncodeURL(
		baseURL+"/entityGroup/%v/removeEntities",
		entityGroupID,
	)
	headers := internal.MergeHeaders(
		c.header.Clone(),
		options.ToHeader(),
	)
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

	var response *mercoafinancego.EntityGroupResponse
	if err := c.caller.Call(
		ctx,
		&internal.CallParams{
			URL:             endpointURL,
			Method:          http.MethodPost,
			Headers:         headers,
			MaxAttempts:     options.MaxAttempts,
			BodyProperties:  options.BodyProperties,
			QueryParameters: options.QueryParameters,
			Client:          options.HTTPClient,
			Request:         request,
			Response:        &response,
			ErrorDecoder:    errorDecoder,
		},
	); err != nil {
		return nil, err
	}
	return response, nil
}
