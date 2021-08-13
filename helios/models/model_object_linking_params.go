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

// ObjectLinkingParams Specifies the parameters required for linking objects. This is currently used as a part of vm migration workflow.
type ObjectLinkingParams struct {
	// Specifies the objectMap that will be used to perform bulk actions such as linking and unlinking.
	ObjectMap []ActionObjectMapping `json:"objectMap,omitempty"`
}

// NewObjectLinkingParams instantiates a new ObjectLinkingParams object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewObjectLinkingParams() *ObjectLinkingParams {
	this := ObjectLinkingParams{}
	return &this
}

// NewObjectLinkingParamsWithDefaults instantiates a new ObjectLinkingParams object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewObjectLinkingParamsWithDefaults() *ObjectLinkingParams {
	this := ObjectLinkingParams{}
	return &this
}

// GetObjectMap returns the ObjectMap field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ObjectLinkingParams) GetObjectMap() []ActionObjectMapping {
	if o == nil  {
		var ret []ActionObjectMapping
		return ret
	}
	return o.ObjectMap
}

// GetObjectMapOk returns a tuple with the ObjectMap field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ObjectLinkingParams) GetObjectMapOk() (*[]ActionObjectMapping, bool) {
	if o == nil || o.ObjectMap == nil {
		return nil, false
	}
	return &o.ObjectMap, true
}

// HasObjectMap returns a boolean if a field has been set.
func (o *ObjectLinkingParams) HasObjectMap() bool {
	if o != nil && o.ObjectMap != nil {
		return true
	}

	return false
}

// SetObjectMap gets a reference to the given []ActionObjectMapping and assigns it to the ObjectMap field.
func (o *ObjectLinkingParams) SetObjectMap(v []ActionObjectMapping) {
	o.ObjectMap = v
}

func (o ObjectLinkingParams) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ObjectMap != nil {
		toSerialize["objectMap"] = o.ObjectMap
	}
	return json.Marshal(toSerialize)
}

type NullableObjectLinkingParams struct {
	value *ObjectLinkingParams
	isSet bool
}

func (v NullableObjectLinkingParams) Get() *ObjectLinkingParams {
	return v.value
}

func (v *NullableObjectLinkingParams) Set(val *ObjectLinkingParams) {
	v.value = val
	v.isSet = true
}

func (v NullableObjectLinkingParams) IsSet() bool {
	return v.isSet
}

func (v *NullableObjectLinkingParams) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableObjectLinkingParams(val *ObjectLinkingParams) *NullableObjectLinkingParams {
	return &NullableObjectLinkingParams{value: val, isSet: true}
}

func (v NullableObjectLinkingParams) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableObjectLinkingParams) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


