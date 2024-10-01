// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// FileSearchResults File/Folder Search Result.
//
// Specifies an array of found files and folders. In addition, a count is
// provided to indicate if additional requests must be made to get
// the full result.
//
// swagger:model FileSearchResults
type FileSearchResults struct {

	// Array of Files and Folders.
	//
	// Specifies the list of files and folders returned by this request that
	// match the specified search and filter criteria.
	// The number of files returned is limited by the pageCount field.
	Files []*FileSearchResult `json:"files"`

	// Specifies cookie for resuming search if pagination is being used.
	// For Librarian queries only.
	PaginationCookie *string `json:"paginationCookie,omitempty"`

	// Specifies the total number of files and folders that match the filter and
	// search criteria. Use this value to determine how many additional
	// requests are required to get the full result.
	TotalCount *int64 `json:"totalCount,omitempty"`
}

// Validate validates this file search results
func (m *FileSearchResults) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateFiles(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *FileSearchResults) validateFiles(formats strfmt.Registry) error {
	if swag.IsZero(m.Files) { // not required
		return nil
	}

	for i := 0; i < len(m.Files); i++ {
		if swag.IsZero(m.Files[i]) { // not required
			continue
		}

		if m.Files[i] != nil {
			if err := m.Files[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("files" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("files" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this file search results based on the context it is used
func (m *FileSearchResults) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateFiles(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *FileSearchResults) contextValidateFiles(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Files); i++ {

		if m.Files[i] != nil {

			if swag.IsZero(m.Files[i]) { // not required
				return nil
			}

			if err := m.Files[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("files" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("files" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *FileSearchResults) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *FileSearchResults) UnmarshalBinary(b []byte) error {
	var res FileSearchResults
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
