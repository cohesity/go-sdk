/*
 * Cohesity REST API
 *
 * This API list provides operations for interfacing with the Cohesity Cluster.
 *
 * API version: 1.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package cohesitysdk

import (
	"encoding/json"
)

// CloneAppViewInfoProto This message encapsulates the information of Clone operation of backup view of an App.
type CloneAppViewInfoProto struct {
	OracleAppViewRestoreInfo *CloneAppViewInfoOracle `json:"oracleAppViewRestoreInfo,omitempty"`
}

// NewCloneAppViewInfoProto instantiates a new CloneAppViewInfoProto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCloneAppViewInfoProto() *CloneAppViewInfoProto {
	this := CloneAppViewInfoProto{}
	return &this
}

// NewCloneAppViewInfoProtoWithDefaults instantiates a new CloneAppViewInfoProto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCloneAppViewInfoProtoWithDefaults() *CloneAppViewInfoProto {
	this := CloneAppViewInfoProto{}
	return &this
}

// GetOracleAppViewRestoreInfo returns the OracleAppViewRestoreInfo field value if set, zero value otherwise.
func (o *CloneAppViewInfoProto) GetOracleAppViewRestoreInfo() CloneAppViewInfoOracle {
	if o == nil || o.OracleAppViewRestoreInfo == nil {
		var ret CloneAppViewInfoOracle
		return ret
	}
	return *o.OracleAppViewRestoreInfo
}

// GetOracleAppViewRestoreInfoOk returns a tuple with the OracleAppViewRestoreInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloneAppViewInfoProto) GetOracleAppViewRestoreInfoOk() (*CloneAppViewInfoOracle, bool) {
	if o == nil || o.OracleAppViewRestoreInfo == nil {
		return nil, false
	}
	return o.OracleAppViewRestoreInfo, true
}

// HasOracleAppViewRestoreInfo returns a boolean if a field has been set.
func (o *CloneAppViewInfoProto) HasOracleAppViewRestoreInfo() bool {
	if o != nil && o.OracleAppViewRestoreInfo != nil {
		return true
	}

	return false
}

// SetOracleAppViewRestoreInfo gets a reference to the given CloneAppViewInfoOracle and assigns it to the OracleAppViewRestoreInfo field.
func (o *CloneAppViewInfoProto) SetOracleAppViewRestoreInfo(v CloneAppViewInfoOracle) {
	o.OracleAppViewRestoreInfo = &v
}

func (o CloneAppViewInfoProto) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.OracleAppViewRestoreInfo != nil {
		toSerialize["oracleAppViewRestoreInfo"] = o.OracleAppViewRestoreInfo
	}
	return json.Marshal(toSerialize)
}

type NullableCloneAppViewInfoProto struct {
	value *CloneAppViewInfoProto
	isSet bool
}

func (v NullableCloneAppViewInfoProto) Get() *CloneAppViewInfoProto {
	return v.value
}

func (v *NullableCloneAppViewInfoProto) Set(val *CloneAppViewInfoProto) {
	v.value = val
	v.isSet = true
}

func (v NullableCloneAppViewInfoProto) IsSet() bool {
	return v.isSet
}

func (v *NullableCloneAppViewInfoProto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCloneAppViewInfoProto(val *CloneAppViewInfoProto) *NullableCloneAppViewInfoProto {
	return &NullableCloneAppViewInfoProto{value: val, isSet: true}
}

func (v NullableCloneAppViewInfoProto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCloneAppViewInfoProto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


