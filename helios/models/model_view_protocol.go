/*
 * Cohesity REST API
 *
 * Cohesity API provides a RESTful interface to access the various data management operations on Cohesity cluster and Helios.
 *
 * API version: 2.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

// package helios
package models

import (
	"encoding/json"
	. "github.com/cohesity/go-sdk/helios/utils"
)

var _ = NullableBool{}

// ViewProtocol Specifies the protocol options for view.
type ViewProtocol struct {
	// Type of protocol. Specifies the supported Protocols for the View.   'NFS' enables protocol access to NFS v3.   'NFS4' enables protocol access to NFS v4.1.   'SMB' enables protocol access to SMB.   'S3' enables protocol access to S3.   'Swift' enables protocol access to Swift.
	Type *string `json:"type,omitempty"`
	// Mode of protocol access.   'ReadOnly'   'ReadWrite'
	Mode *string `json:"mode,omitempty"`
}

// NewViewProtocol instantiates a new ViewProtocol object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewViewProtocol() *ViewProtocol {
	this := ViewProtocol{}
	return &this
}

// NewViewProtocolWithDefaults instantiates a new ViewProtocol object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewViewProtocolWithDefaults() *ViewProtocol {
	this := ViewProtocol{}
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *ViewProtocol) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ViewProtocol) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *ViewProtocol) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *ViewProtocol) SetType(v string) {
	o.Type = &v
}

// GetMode returns the Mode field value if set, zero value otherwise.
func (o *ViewProtocol) GetMode() string {
	if o == nil || o.Mode == nil {
		var ret string
		return ret
	}
	return *o.Mode
}

// GetModeOk returns a tuple with the Mode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ViewProtocol) GetModeOk() (*string, bool) {
	if o == nil || o.Mode == nil {
		return nil, false
	}
	return o.Mode, true
}

// HasMode returns a boolean if a field has been set.
func (o *ViewProtocol) HasMode() bool {
	if o != nil && o.Mode != nil {
		return true
	}

	return false
}

// SetMode gets a reference to the given string and assigns it to the Mode field.
func (o *ViewProtocol) SetMode(v string) {
	o.Mode = &v
}

func (o ViewProtocol) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.Mode != nil {
		toSerialize["mode"] = o.Mode
	}
	return json.Marshal(toSerialize)
}

type NullableViewProtocol struct {
	value *ViewProtocol
	isSet bool
}

func (v NullableViewProtocol) Get() *ViewProtocol {
	return v.value
}

func (v *NullableViewProtocol) Set(val *ViewProtocol) {
	v.value = val
	v.isSet = true
}

func (v NullableViewProtocol) IsSet() bool {
	return v.isSet
}

func (v *NullableViewProtocol) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableViewProtocol(val *ViewProtocol) *NullableViewProtocol {
	return &NullableViewProtocol{value: val, isSet: true}
}

func (v NullableViewProtocol) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableViewProtocol) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


