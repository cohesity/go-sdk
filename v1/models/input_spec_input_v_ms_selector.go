// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// InputSpecInputVMsSelector input spec input v ms selector
//
// swagger:model InputSpec_InputVMsSelector
type InputSpecInputVMsSelector struct {

	// Time filter for file's last change time.
	FileTimeFilter *InputSpecFileTimeFilter `json:"fileTimeFilter,omitempty"`

	// After VMDKs are selected as above, the files within them can be selected
	// by using these predicates.
	FilenameGlob []string `json:"filenameGlob"`

	// job ids
	JobIds []int64 `json:"jobIds"`

	// Exclusive end of snapshot_timestamp range. If missing, inf will be used
	// as the timestamp range.
	MaxSnapshotTimestamp *int64 `json:"maxSnapshotTimestamp,omitempty"`

	// Inclusive. If missing, 0 will the default lower end of timestamp range
	MinSnapshotTimestamp *int64 `json:"minSnapshotTimestamp,omitempty"`

	// Filters are AND of ORs.
	PartitionIds []int64 `json:"partitionIds"`

	// Boolean flag to indicate if only latest snapshot of each object should
	// be processed.
	ProcessLatestOnly *bool `json:"processLatestOnly,omitempty"`

	// Within each volume, traversal will be rooted at this directory. A
	// typical value here might be /home
	RootDir *string `json:"rootDir,omitempty"`

	// source entity ids
	SourceEntityIds []int64 `json:"sourceEntityIds"`

	// view box ids
	ViewBoxIds []int64 `json:"viewBoxIds"`

	// view names
	ViewNames []string `json:"viewNames"`
}

// Validate validates this input spec input v ms selector
func (m *InputSpecInputVMsSelector) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateFileTimeFilter(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *InputSpecInputVMsSelector) validateFileTimeFilter(formats strfmt.Registry) error {
	if swag.IsZero(m.FileTimeFilter) { // not required
		return nil
	}

	if m.FileTimeFilter != nil {
		if err := m.FileTimeFilter.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("fileTimeFilter")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("fileTimeFilter")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this input spec input v ms selector based on the context it is used
func (m *InputSpecInputVMsSelector) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateFileTimeFilter(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *InputSpecInputVMsSelector) contextValidateFileTimeFilter(ctx context.Context, formats strfmt.Registry) error {

	if m.FileTimeFilter != nil {

		if swag.IsZero(m.FileTimeFilter) { // not required
			return nil
		}

		if err := m.FileTimeFilter.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("fileTimeFilter")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("fileTimeFilter")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *InputSpecInputVMsSelector) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *InputSpecInputVMsSelector) UnmarshalBinary(b []byte) error {
	var res InputSpecInputVMsSelector
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
