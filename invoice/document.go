// This file was auto-generated by Fern from our API Definition.

package invoice

type UploadDocumentRequest struct {
	// Base64 encoded image or PDF of invoice document. PNG, JPG, and PDF are supported. 10MB max.
	Document *string `json:"document,omitempty" url:"document,omitempty"`
}
