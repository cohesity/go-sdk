/*
 * Cohesity REST API
 *
 * Cohesity API provides a RESTful interface to access the various data management operations on Cohesity cluster and Helios.
 *
 * API version: 2.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

// package onprem
package models

import (
	"encoding/json"
	. "github.com/cohesity/go-sdk/onprem/utils"
	"fmt"
)

var _ = NullableBool{}

// MssqlConnectionParams Specifies the parameters to connect to a SQL node/cluster using given IP or hostname FQDN.
type MssqlConnectionParams struct {
	// Specifies the unique identifier to locate the SQL node or cluster. The host identifier can be IP address or FQDN.
	HostIdentifier string `json:"hostIdentifier"`
}

// NewMssqlConnectionParams instantiates a new MssqlConnectionParams object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMssqlConnectionParams(hostIdentifier string) *MssqlConnectionParams {
	this := MssqlConnectionParams{}
	this.HostIdentifier = hostIdentifier
	return &this
}

// NewMssqlConnectionParamsWithDefaults instantiates a new MssqlConnectionParams object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMssqlConnectionParamsWithDefaults() *MssqlConnectionParams {
	this := MssqlConnectionParams{}
	return &this
}

// GetHostIdentifier returns the HostIdentifier field value
func (o *MssqlConnectionParams) GetHostIdentifier() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.HostIdentifier
}

// GetHostIdentifierOk returns a tuple with the HostIdentifier field value
// and a boolean to check if the value has been set.
func (o *MssqlConnectionParams) GetHostIdentifierOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.HostIdentifier, true
}

// SetHostIdentifier sets field value
func (o *MssqlConnectionParams) SetHostIdentifier(v string) {
	o.HostIdentifier = v
}

func (o MssqlConnectionParams) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["hostIdentifier"] = o.HostIdentifier
	}
	return json.Marshal(toSerialize)
}

type NullableMssqlConnectionParams struct {
	value *MssqlConnectionParams
	isSet bool
}

func (v NullableMssqlConnectionParams) Get() *MssqlConnectionParams {
	return v.value
}

func (v *NullableMssqlConnectionParams) Set(val *MssqlConnectionParams) {
	v.value = val
	v.isSet = true
}

func (v NullableMssqlConnectionParams) IsSet() bool {
	return v.isSet
}

func (v *NullableMssqlConnectionParams) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMssqlConnectionParams(val *MssqlConnectionParams) *NullableMssqlConnectionParams {
	return &NullableMssqlConnectionParams{value: val, isSet: true}
}

func (v NullableMssqlConnectionParams) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMssqlConnectionParams) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}




func (o MssqlConnectionParams) Print() {
		if byteArray, err := o.MarshalJSON(); err != nil {
				fmt.Println(err)
		} else {
				fmt.Println(string(byteArray))
		}
}