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

// MsSQLCommonConnectionParams Specifies the common parameters to connect to a SQL node/cluster
type MsSQLCommonConnectionParams struct {
	// Specifies the unique identifier to locate the SQL node or cluster. The host identifier can be IP address or FQDN.
	HostIdentifier string `json:"hostIdentifier"`
}

// NewMsSQLCommonConnectionParams instantiates a new MsSQLCommonConnectionParams object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMsSQLCommonConnectionParams(hostIdentifier string) *MsSQLCommonConnectionParams {
	this := MsSQLCommonConnectionParams{}
	this.HostIdentifier = hostIdentifier
	return &this
}

// NewMsSQLCommonConnectionParamsWithDefaults instantiates a new MsSQLCommonConnectionParams object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMsSQLCommonConnectionParamsWithDefaults() *MsSQLCommonConnectionParams {
	this := MsSQLCommonConnectionParams{}
	return &this
}

// GetHostIdentifier returns the HostIdentifier field value
func (o *MsSQLCommonConnectionParams) GetHostIdentifier() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.HostIdentifier
}

// GetHostIdentifierOk returns a tuple with the HostIdentifier field value
// and a boolean to check if the value has been set.
func (o *MsSQLCommonConnectionParams) GetHostIdentifierOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.HostIdentifier, true
}

// SetHostIdentifier sets field value
func (o *MsSQLCommonConnectionParams) SetHostIdentifier(v string) {
	o.HostIdentifier = v
}

func (o MsSQLCommonConnectionParams) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["hostIdentifier"] = o.HostIdentifier
	}
	return json.Marshal(toSerialize)
}

type NullableMsSQLCommonConnectionParams struct {
	value *MsSQLCommonConnectionParams
	isSet bool
}

func (v NullableMsSQLCommonConnectionParams) Get() *MsSQLCommonConnectionParams {
	return v.value
}

func (v *NullableMsSQLCommonConnectionParams) Set(val *MsSQLCommonConnectionParams) {
	v.value = val
	v.isSet = true
}

func (v NullableMsSQLCommonConnectionParams) IsSet() bool {
	return v.isSet
}

func (v *NullableMsSQLCommonConnectionParams) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMsSQLCommonConnectionParams(val *MsSQLCommonConnectionParams) *NullableMsSQLCommonConnectionParams {
	return &NullableMsSQLCommonConnectionParams{value: val, isSet: true}
}

func (v NullableMsSQLCommonConnectionParams) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMsSQLCommonConnectionParams) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


