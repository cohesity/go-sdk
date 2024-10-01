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

// ObjectInformation object information
//
// swagger:model ObjectInformation
type ObjectInformation struct {

	// Species the list of user who have access to this object.
	AccessibleUsers []string `json:"accessibleUsers"`

	// Specifies the audit log information.
	AuditLogs []*ClusterAuditLog `json:"auditLogs"`

	// Specifies the copy task information.
	CopyTaskInfo []*GdprCopyTask `json:"copyTaskInfo"`

	// Specifies the protection status of the object.
	IsProtected *bool `json:"isProtected,omitempty"`

	// Specifies the location of the parent source.
	Location *string `json:"location,omitempty"`

	// Specifies the data locations for the protected objects.
	ProtectionInfo []*ProtectionInfo `json:"protectionInfo"`

	// Specifies the id of the root node.
	RootNodeID *int64 `json:"rootNodeId,omitempty"`

	// Specifies the id of the Protection Source.
	SourceID *int64 `json:"sourceId,omitempty"`

	// Specifies the name of the object.
	SourceName *string `json:"sourceName,omitempty"`
}

// Validate validates this object information
func (m *ObjectInformation) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAuditLogs(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCopyTaskInfo(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateProtectionInfo(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ObjectInformation) validateAuditLogs(formats strfmt.Registry) error {
	if swag.IsZero(m.AuditLogs) { // not required
		return nil
	}

	for i := 0; i < len(m.AuditLogs); i++ {
		if swag.IsZero(m.AuditLogs[i]) { // not required
			continue
		}

		if m.AuditLogs[i] != nil {
			if err := m.AuditLogs[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("auditLogs" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("auditLogs" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ObjectInformation) validateCopyTaskInfo(formats strfmt.Registry) error {
	if swag.IsZero(m.CopyTaskInfo) { // not required
		return nil
	}

	for i := 0; i < len(m.CopyTaskInfo); i++ {
		if swag.IsZero(m.CopyTaskInfo[i]) { // not required
			continue
		}

		if m.CopyTaskInfo[i] != nil {
			if err := m.CopyTaskInfo[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("copyTaskInfo" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("copyTaskInfo" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ObjectInformation) validateProtectionInfo(formats strfmt.Registry) error {
	if swag.IsZero(m.ProtectionInfo) { // not required
		return nil
	}

	for i := 0; i < len(m.ProtectionInfo); i++ {
		if swag.IsZero(m.ProtectionInfo[i]) { // not required
			continue
		}

		if m.ProtectionInfo[i] != nil {
			if err := m.ProtectionInfo[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("protectionInfo" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("protectionInfo" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this object information based on the context it is used
func (m *ObjectInformation) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAuditLogs(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateCopyTaskInfo(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateProtectionInfo(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ObjectInformation) contextValidateAuditLogs(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.AuditLogs); i++ {

		if m.AuditLogs[i] != nil {

			if swag.IsZero(m.AuditLogs[i]) { // not required
				return nil
			}

			if err := m.AuditLogs[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("auditLogs" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("auditLogs" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ObjectInformation) contextValidateCopyTaskInfo(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.CopyTaskInfo); i++ {

		if m.CopyTaskInfo[i] != nil {

			if swag.IsZero(m.CopyTaskInfo[i]) { // not required
				return nil
			}

			if err := m.CopyTaskInfo[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("copyTaskInfo" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("copyTaskInfo" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ObjectInformation) contextValidateProtectionInfo(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.ProtectionInfo); i++ {

		if m.ProtectionInfo[i] != nil {

			if swag.IsZero(m.ProtectionInfo[i]) { // not required
				return nil
			}

			if err := m.ProtectionInfo[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("protectionInfo" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("protectionInfo" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *ObjectInformation) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ObjectInformation) UnmarshalBinary(b []byte) error {
	var res ObjectInformation
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
