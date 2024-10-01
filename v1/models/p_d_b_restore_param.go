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

// PDBRestoreParam p d b restore param
//
// swagger:model PDBRestoreParam
type PDBRestoreParam struct {

	// During the restore workflow, drop the pdb if the same name pdb exists.
	DropPdbsIfExists *bool `json:"dropPdbsIfExists,omitempty"`

	// Restore given list of pdbs to an existing CDB.
	ExistingCdb *bool `json:"existingCdb,omitempty"`

	// Whether or not to restore the PDB when restoring the CDB.
	// Before this field was added, agent's behavior was to include the PDBs
	// provided, there was no exclusion. By keeping the value as default true
	// we will ensure the agent behaves the same way if someone upgrades agent
	// service and the magneto service is still old.
	IncludeInRestore *bool `json:"includeInRestore,omitempty"`

	// List of PDB's to restore or to skip based on "include_in_restore".
	PdbEntityInfoVec *CDBEntityInfo `json:"pdbEntityInfoVec,omitempty"`

	// If rename pdb is provided, list of new names for the pdb.
	RenamePdbMap []*PDBRestoreParamRenamePdbMapEntry `json:"renamePdbMap"`
}

// Validate validates this p d b restore param
func (m *PDBRestoreParam) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validatePdbEntityInfoVec(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRenamePdbMap(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PDBRestoreParam) validatePdbEntityInfoVec(formats strfmt.Registry) error {
	if swag.IsZero(m.PdbEntityInfoVec) { // not required
		return nil
	}

	if m.PdbEntityInfoVec != nil {
		if err := m.PdbEntityInfoVec.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("pdbEntityInfoVec")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("pdbEntityInfoVec")
			}
			return err
		}
	}

	return nil
}

func (m *PDBRestoreParam) validateRenamePdbMap(formats strfmt.Registry) error {
	if swag.IsZero(m.RenamePdbMap) { // not required
		return nil
	}

	for i := 0; i < len(m.RenamePdbMap); i++ {
		if swag.IsZero(m.RenamePdbMap[i]) { // not required
			continue
		}

		if m.RenamePdbMap[i] != nil {
			if err := m.RenamePdbMap[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("renamePdbMap" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("renamePdbMap" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this p d b restore param based on the context it is used
func (m *PDBRestoreParam) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidatePdbEntityInfoVec(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateRenamePdbMap(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PDBRestoreParam) contextValidatePdbEntityInfoVec(ctx context.Context, formats strfmt.Registry) error {

	if m.PdbEntityInfoVec != nil {

		if swag.IsZero(m.PdbEntityInfoVec) { // not required
			return nil
		}

		if err := m.PdbEntityInfoVec.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("pdbEntityInfoVec")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("pdbEntityInfoVec")
			}
			return err
		}
	}

	return nil
}

func (m *PDBRestoreParam) contextValidateRenamePdbMap(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.RenamePdbMap); i++ {

		if m.RenamePdbMap[i] != nil {

			if swag.IsZero(m.RenamePdbMap[i]) { // not required
				return nil
			}

			if err := m.RenamePdbMap[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("renamePdbMap" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("renamePdbMap" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *PDBRestoreParam) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PDBRestoreParam) UnmarshalBinary(b []byte) error {
	var res PDBRestoreParam
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
