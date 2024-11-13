// This file was auto-generated by Fern from our API Definition.

package invoicetemplate

import (
	mercoafinancego "github.com/mercoa-finance/go"
)

type GetDocumentsRequest struct {
	// Filter by document type
	Type []*mercoafinancego.DocumentType `json:"-" url:"type,omitempty"`
}

type UploadDocumentRequest struct {
	// Base64 encoded image or PDF of invoice document. PNG, JPG, WEBP, and PDF are supported. 10MB max.
	Document string `json:"document" url:"-"`
	// Specify Document Type, defaults to INVOICE
	Type *mercoafinancego.DocumentType `json:"type,omitempty" url:"-"`
}
