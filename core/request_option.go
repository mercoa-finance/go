// This file was auto-generated by Fern from our API Definition.

package core

import (
	http "net/http"
)

// RequestOption adapts the behavior of the client or an individual request.
type RequestOption interface {
	applyRequestOptions(*RequestOptions)
}

// RequestOptions defines all of the possible request options.
//
// This type is primarily used by the generated code and is not meant
// to be used directly; use the option package instead.
type RequestOptions struct {
	BaseURL     string
	HTTPClient  HTTPClient
	HTTPHeader  http.Header
	MaxAttempts uint
	Token       string
}

// NewRequestOptions returns a new *RequestOptions value.
//
// This function is primarily used by the generated code and is not meant
// to be used directly; use RequestOption instead.
func NewRequestOptions(opts ...RequestOption) *RequestOptions {
	options := &RequestOptions{
		HTTPHeader: make(http.Header),
	}
	for _, opt := range opts {
		opt.applyRequestOptions(options)
	}
	return options
}

// ToHeader maps the configured request options into a http.Header used
// for the request(s).
func (r *RequestOptions) ToHeader() http.Header {
	header := r.cloneHeader()
	if r.Token != "" {
		header.Set("Authorization", "Bearer "+r.Token)
	}
	return header
}

func (r *RequestOptions) cloneHeader() http.Header {
	headers := r.HTTPHeader.Clone()
	headers.Set("X-Fern-Language", "Go")
	headers.Set("X-Fern-SDK-Name", "github.com/mercoa-finance/go")
	headers.Set("X-Fern-SDK-Version", "v0.3.38")
	return headers
}

// BaseURLOption implements the RequestOption interface.
type BaseURLOption struct {
	BaseURL string
}

func (b *BaseURLOption) applyRequestOptions(opts *RequestOptions) {
	opts.BaseURL = b.BaseURL
}

// HTTPClientOption implements the RequestOption interface.
type HTTPClientOption struct {
	HTTPClient HTTPClient
}

func (h *HTTPClientOption) applyRequestOptions(opts *RequestOptions) {
	opts.HTTPClient = h.HTTPClient
}

// HTTPHeaderOption implements the RequestOption interface.
type HTTPHeaderOption struct {
	HTTPHeader http.Header
}

func (h *HTTPHeaderOption) applyRequestOptions(opts *RequestOptions) {
	opts.HTTPHeader = h.HTTPHeader
}

// MaxAttemptsOption implements the RequestOption interface.
type MaxAttemptsOption struct {
	MaxAttempts uint
}

func (m *MaxAttemptsOption) applyRequestOptions(opts *RequestOptions) {
	opts.MaxAttempts = m.MaxAttempts
}

// TokenOption implements the RequestOption interface.
type TokenOption struct {
	Token string
}

func (t *TokenOption) applyRequestOptions(opts *RequestOptions) {
	opts.Token = t.Token
}
