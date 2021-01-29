package types

import "time"

type (
	DocumentType int

	// Document is a final result of a rendering process based on a Template and data
	Document struct {
		ID          uint64       `json:"id,string"`
		Template    uint64       `json:"templateID,string"`
		Type        DocumentType `json:"type"`
		Name        string       `json:"name"`
		Description string       `json:"description"`
		// Errors define a set of issues that occurred when rendering the document
		Errors []error `json:"errors,omitempty"`

		OwnerID uint64 `json:"ownerID,string"`

		CreatedAt time.Time  `json:"createdAt,omitempty"`
		UpdatedAt *time.Time `json:"updatedAt,omitempty"`
		DeletedAt *time.Time `json:"deletedAt,omitempty"`
	}
	DocumentSet []*Document
)

const (
	DocumentTypeHTML DocumentType = iota
	DocumentTypeJSON
	// ...
)
