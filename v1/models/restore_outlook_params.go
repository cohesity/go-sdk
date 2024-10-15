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

// RestoreOutlookParams This message defines the Outlook specific restore params.
//
// swagger:model RestoreOutlookParams
type RestoreOutlookParams struct {

	// Human readable prefix that is prepended to the archive recoverable items
	// folder name.
	ArchiveRecoverableItemsPrefix *string `json:"archiveRecoverableItemsPrefix,omitempty"`

	// Governs how items are restored to microsoft. See enum definitions for
	// details.
	ItemRecoveryMethod *int32 `json:"itemRecoveryMethod,omitempty"`

	// In a RestoreJob , user will provide the list of mailboxes
	// to be restored. Provision is there for restoring full AND
	// partial mailbox recovery.
	MailboxVec []*RestoreOutlookParamsMailbox `json:"mailboxVec"`

	// These are the parameters that user can provide for converting/recovering
	// selected EWS items to PST format.
	PstParams *EwsToPstConversionParams `json:"pstParams,omitempty"`

	// Human readable prefix that is prepended to the recoverable items
	// folder name.
	RecoverableItemsPrefix *string `json:"recoverableItemsPrefix,omitempty"`

	// Indicates whether PST conversion should skip mailbox entity permit.
	SkipMbxPermitForPst *bool `json:"skipMbxPermitForPst,omitempty"`

	// Whether to skip recovery of the archive mailbox (or its items).
	SkipRecoverArchiveMailbox *bool `json:"skipRecoverArchiveMailbox,omitempty"`

	// Whether to skip recovery of archive recoverable items folders.
	SkipRecoverArchiveRecoverableItems *bool `json:"skipRecoverArchiveRecoverableItems,omitempty"`

	// Whether to skip recovery of recoverable items folders.
	SkipRecoverRecoverableItems *bool `json:"skipRecoverRecoverableItems,omitempty"`

	// User will type the target folder path.
	// This will always be specified (whether target_mailbox is original mailbox
	// or alternate). If multiple folders are selected, they will all be
	// restored to this folder. The appropriate hierarchy along with the folder
	// names will be preserved.
	TargetFolderPath *string `json:"targetFolderPath,omitempty"`

	// This is the destination mailbox.
	// All mailboxes listed in the mailbox_vec will be restored to this mailbox
	// with appropriate names.
	// Let's say mailbox_vec is A and B; target_mailbox is C.
	// The final folder-hierarchy after restore job is finished will look like
	// this :
	// C : {target_folder_path}/|
	// |A/{whatever is there in mailbox A}
	// |B/{whatever is inside B mailbox B}
	TargetMailbox *EntityProto `json:"targetMailbox,omitempty"`
}

// Validate validates this restore outlook params
func (m *RestoreOutlookParams) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateMailboxVec(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePstParams(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTargetMailbox(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RestoreOutlookParams) validateMailboxVec(formats strfmt.Registry) error {
	if swag.IsZero(m.MailboxVec) { // not required
		return nil
	}

	for i := 0; i < len(m.MailboxVec); i++ {
		if swag.IsZero(m.MailboxVec[i]) { // not required
			continue
		}

		if m.MailboxVec[i] != nil {
			if err := m.MailboxVec[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("mailboxVec" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("mailboxVec" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *RestoreOutlookParams) validatePstParams(formats strfmt.Registry) error {
	if swag.IsZero(m.PstParams) { // not required
		return nil
	}

	if m.PstParams != nil {
		if err := m.PstParams.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("pstParams")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("pstParams")
			}
			return err
		}
	}

	return nil
}

func (m *RestoreOutlookParams) validateTargetMailbox(formats strfmt.Registry) error {
	if swag.IsZero(m.TargetMailbox) { // not required
		return nil
	}

	if m.TargetMailbox != nil {
		if err := m.TargetMailbox.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("targetMailbox")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("targetMailbox")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this restore outlook params based on the context it is used
func (m *RestoreOutlookParams) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateMailboxVec(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidatePstParams(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateTargetMailbox(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RestoreOutlookParams) contextValidateMailboxVec(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.MailboxVec); i++ {

		if m.MailboxVec[i] != nil {

			if swag.IsZero(m.MailboxVec[i]) { // not required
				return nil
			}

			if err := m.MailboxVec[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("mailboxVec" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("mailboxVec" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *RestoreOutlookParams) contextValidatePstParams(ctx context.Context, formats strfmt.Registry) error {

	if m.PstParams != nil {

		if swag.IsZero(m.PstParams) { // not required
			return nil
		}

		if err := m.PstParams.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("pstParams")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("pstParams")
			}
			return err
		}
	}

	return nil
}

func (m *RestoreOutlookParams) contextValidateTargetMailbox(ctx context.Context, formats strfmt.Registry) error {

	if m.TargetMailbox != nil {

		if swag.IsZero(m.TargetMailbox) { // not required
			return nil
		}

		if err := m.TargetMailbox.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("targetMailbox")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("targetMailbox")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *RestoreOutlookParams) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RestoreOutlookParams) UnmarshalBinary(b []byte) error {
	var res RestoreOutlookParams
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
