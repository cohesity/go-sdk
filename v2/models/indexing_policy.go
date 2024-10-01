// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// IndexingPolicy Indexing Policy.
//
// Specifies settings for indexing files found in an Object (such as a VM) so these files can be searched and recovered. This also specifies inclusion and exclusion rules that determine the directories to index.
//
// swagger:model IndexingPolicy
type IndexingPolicy struct {

	// Specifies if the files found in an Object (such as a VM) should be indexed. If true (the default), files are indexed.
	// Required: true
	EnableIndexing *bool `json:"enableIndexing"`

	// Array of Indexed Directories. Specifies a list of directories to index. Regular expression can also be specified to provide the directory paths. Example: /Users/<wildcard>/AppData
	IncludePaths []string `json:"includePaths"`

	// Array of Excluded Directories. Specifies a list of directories to exclude from indexing.Regular expression can also be specified to provide the directory paths. Example: /Users/<wildcard>/AppData
	ExcludePaths []string `json:"excludePaths"`
}

// Validate validates this indexing policy
func (m *IndexingPolicy) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateEnableIndexing(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *IndexingPolicy) validateEnableIndexing(formats strfmt.Registry) error {

	if err := validate.Required("enableIndexing", "body", m.EnableIndexing); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this indexing policy based on context it is used
func (m *IndexingPolicy) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *IndexingPolicy) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *IndexingPolicy) UnmarshalBinary(b []byte) error {
	var res IndexingPolicy
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
