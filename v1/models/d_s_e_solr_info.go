// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// DSESolrInfo Message to hold information about DSE Solr node.
//
// swagger:model DSESolrInfo
type DSESolrInfo struct {

	// Solr node IP Addresses.
	SolrNodeVec []string `json:"solrNodeVec"`

	// Solr node Port.
	SolrPort *uint32 `json:"solrPort,omitempty"`
}

// Validate validates this d s e solr info
func (m *DSESolrInfo) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this d s e solr info based on context it is used
func (m *DSESolrInfo) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *DSESolrInfo) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DSESolrInfo) UnmarshalBinary(b []byte) error {
	var res DSESolrInfo
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
