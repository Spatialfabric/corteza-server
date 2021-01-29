package types

import "time"

type (
	// TemplateType defines the content type of the Template
	TemplateType int

	// Template describes some source that can be used to produce a Document
	Template struct {
		ID          uint64       `json:"id,string"`
		Type        TemplateType `json:"type"`
		Name        string       `json:"name"`
		Description string       `json:"description,omitempty"`
		Revision    uint         `json:"revision,omitempty"`
		// Errors define a set of issues that occurred when performing preprocessing
		//
		// @note might remove this one and just keep Document.Errors
		Errors []error `json:"errors,omitempty"`

		OwnerID uint64 `json:"ownerID,string"`

		CreatedAt  time.Time  `json:"createdAt,omitempty"`
		UpdatedAt  *time.Time `json:"updatedAt,omitempty"`
		DeletedAt  *time.Time `json:"deletedAt,omitempty"`
		LastUsedAt *time.Time `json:"lastUsedAt,omitempty"`
	}
	TemplateSet []*Template
)

const (
	TemplateTypeHTML TemplateType = iota
	TemplateTypeJSON
	// ...
)
