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

// PrivateUdaEntity private uda entity
//
// swagger:model PrivateUdaEntity
type PrivateUdaEntity struct {

	// Information about the agents linked to this UDA entity.
	AgentStatusVec []*HostAgentStatus `json:"agentStatusVec"`

	// Information of a UDA cluster, only valid for an entity of type
	// kCluster.
	ClusterInfo *PrivateClusterInfo `json:"clusterInfo,omitempty"`

	// Fully qualified name for the object. E.g: For a table this could be
	// <database_name>.<table_name>.
	Name *string `json:"name,omitempty"`

	// Information of a generic UDA entity. only valid for an entity
	// of type kObject.
	ObjectInfo *ObjectInfo `json:"objectInfo,omitempty"`

	// Universal Data Adapter (UDA) source type.
	// Must be same as RegisteredEntityUdaParams::source_type.
	SourceType *string `json:"sourceType,omitempty"`

	// The type of entity this proto refers to.
	Type *int32 `json:"type,omitempty"`

	// The UUID of the object.
	UUID *string `json:"uuid,omitempty"`
}

// Validate validates this private uda entity
func (m *PrivateUdaEntity) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAgentStatusVec(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateClusterInfo(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateObjectInfo(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PrivateUdaEntity) validateAgentStatusVec(formats strfmt.Registry) error {
	if swag.IsZero(m.AgentStatusVec) { // not required
		return nil
	}

	for i := 0; i < len(m.AgentStatusVec); i++ {
		if swag.IsZero(m.AgentStatusVec[i]) { // not required
			continue
		}

		if m.AgentStatusVec[i] != nil {
			if err := m.AgentStatusVec[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("agentStatusVec" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("agentStatusVec" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *PrivateUdaEntity) validateClusterInfo(formats strfmt.Registry) error {
	if swag.IsZero(m.ClusterInfo) { // not required
		return nil
	}

	if m.ClusterInfo != nil {
		if err := m.ClusterInfo.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("clusterInfo")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("clusterInfo")
			}
			return err
		}
	}

	return nil
}

func (m *PrivateUdaEntity) validateObjectInfo(formats strfmt.Registry) error {
	if swag.IsZero(m.ObjectInfo) { // not required
		return nil
	}

	if m.ObjectInfo != nil {
		if err := m.ObjectInfo.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("objectInfo")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("objectInfo")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this private uda entity based on the context it is used
func (m *PrivateUdaEntity) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAgentStatusVec(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateClusterInfo(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateObjectInfo(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PrivateUdaEntity) contextValidateAgentStatusVec(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.AgentStatusVec); i++ {

		if m.AgentStatusVec[i] != nil {

			if swag.IsZero(m.AgentStatusVec[i]) { // not required
				return nil
			}

			if err := m.AgentStatusVec[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("agentStatusVec" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("agentStatusVec" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *PrivateUdaEntity) contextValidateClusterInfo(ctx context.Context, formats strfmt.Registry) error {

	if m.ClusterInfo != nil {

		if swag.IsZero(m.ClusterInfo) { // not required
			return nil
		}

		if err := m.ClusterInfo.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("clusterInfo")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("clusterInfo")
			}
			return err
		}
	}

	return nil
}

func (m *PrivateUdaEntity) contextValidateObjectInfo(ctx context.Context, formats strfmt.Registry) error {

	if m.ObjectInfo != nil {

		if swag.IsZero(m.ObjectInfo) { // not required
			return nil
		}

		if err := m.ObjectInfo.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("objectInfo")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("objectInfo")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *PrivateUdaEntity) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PrivateUdaEntity) UnmarshalBinary(b []byte) error {
	var res PrivateUdaEntity
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
