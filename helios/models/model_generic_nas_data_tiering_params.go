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

// GenericNasDataTieringParams Specifies the parameters which are specific to NAS related Protection Groups.
type GenericNasDataTieringParams struct {
	// Specifies the objects to be included in the Protection Group.
	Objects []ProtectionObjectInput `json:"objects"`
	// Specifies the id of the root of data tiering source.
	SourceId NullableInt64 `json:"sourceId,omitempty"`
}

// NewGenericNasDataTieringParams instantiates a new GenericNasDataTieringParams object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGenericNasDataTieringParams(objects []ProtectionObjectInput) *GenericNasDataTieringParams {
	this := GenericNasDataTieringParams{}
	this.Objects = objects
	return &this
}

// NewGenericNasDataTieringParamsWithDefaults instantiates a new GenericNasDataTieringParams object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGenericNasDataTieringParamsWithDefaults() *GenericNasDataTieringParams {
	this := GenericNasDataTieringParams{}
	return &this
}

// GetObjects returns the Objects field value
func (o *GenericNasDataTieringParams) GetObjects() []ProtectionObjectInput {
	if o == nil {
		var ret []ProtectionObjectInput
		return ret
	}

	return o.Objects
}

// GetObjectsOk returns a tuple with the Objects field value
// and a boolean to check if the value has been set.
func (o *GenericNasDataTieringParams) GetObjectsOk() (*[]ProtectionObjectInput, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Objects, true
}

// SetObjects sets field value
func (o *GenericNasDataTieringParams) SetObjects(v []ProtectionObjectInput) {
	o.Objects = v
}

// GetSourceId returns the SourceId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GenericNasDataTieringParams) GetSourceId() int64 {
	if o == nil || o.SourceId.Get() == nil {
		var ret int64
		return ret
	}
	return *o.SourceId.Get()
}

// GetSourceIdOk returns a tuple with the SourceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GenericNasDataTieringParams) GetSourceIdOk() (*int64, bool) {
	if o == nil  {
		return nil, false
	}
	return o.SourceId.Get(), o.SourceId.IsSet()
}

// HasSourceId returns a boolean if a field has been set.
func (o *GenericNasDataTieringParams) HasSourceId() bool {
	if o != nil && o.SourceId.IsSet() {
		return true
	}

	return false
}

// SetSourceId gets a reference to the given NullableInt64 and assigns it to the SourceId field.
func (o *GenericNasDataTieringParams) SetSourceId(v int64) {
	o.SourceId.Set(&v)
}
// SetSourceIdNil sets the value for SourceId to be an explicit nil
func (o *GenericNasDataTieringParams) SetSourceIdNil() {
	o.SourceId.Set(nil)
}

// UnsetSourceId ensures that no value is present for SourceId, not even an explicit nil
func (o *GenericNasDataTieringParams) UnsetSourceId() {
	o.SourceId.Unset()
}

func (o GenericNasDataTieringParams) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["objects"] = o.Objects
	}
	if o.SourceId.IsSet() {
		toSerialize["sourceId"] = o.SourceId.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableGenericNasDataTieringParams struct {
	value *GenericNasDataTieringParams
	isSet bool
}

func (v NullableGenericNasDataTieringParams) Get() *GenericNasDataTieringParams {
	return v.value
}

func (v *NullableGenericNasDataTieringParams) Set(val *GenericNasDataTieringParams) {
	v.value = val
	v.isSet = true
}

func (v NullableGenericNasDataTieringParams) IsSet() bool {
	return v.isSet
}

func (v *NullableGenericNasDataTieringParams) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGenericNasDataTieringParams(val *GenericNasDataTieringParams) *NullableGenericNasDataTieringParams {
	return &NullableGenericNasDataTieringParams{value: val, isSet: true}
}

func (v NullableGenericNasDataTieringParams) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGenericNasDataTieringParams) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


