// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// DataEvent Data event to be published/shared to magneto for further processing.
//
// DataEvent extensions                     Location                        No.
// =============================================================================
// vmware.VMwareDataEvent                   vmware_event.proto              100
// magneto.MongoDBDataEvent                 mongodb_event.proto             101
// =============================================================================
//
// swagger:model DataEvent
type DataEvent struct {

	// Amount of data written to SnapFS in this event.
	DataSize *int64 `json:"dataSize,omitempty"`

	// ID of the epoch where this DataEvent belongs. This is only set for the
	// entities with sharded log journal.
	EpochID *int64 `json:"epochId,omitempty"`

	// File in which this data is written. This file path is the absolute path
	// starting from the view_name in the AtomEntityProto. This is not set for
	// entities using CDP sharding.
	FilePath *string `json:"filePath,omitempty"`

	// This field denotes the location of the associated shard locator file.
	// This is supposed to be set for the entities using sharding to store the
	// data.
	ShardLocatorFilePath *string `json:"shardLocatorFilePath,omitempty"`
}

// Validate validates this data event
func (m *DataEvent) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this data event based on context it is used
func (m *DataEvent) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *DataEvent) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DataEvent) UnmarshalBinary(b []byte) error {
	var res DataEvent
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
