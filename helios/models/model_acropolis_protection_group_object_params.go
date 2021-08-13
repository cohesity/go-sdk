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

// AcropolisProtectionGroupObjectParams Specifies an object protected by a Acropolis Protection Group.
type AcropolisProtectionGroupObjectParams struct {
	// Specifies the ID of the object.
	Id NullableInt64 `json:"id"`
	// Specifies the name of the virtual machine.
	Name NullableString `json:"name,omitempty"`
}

// NewAcropolisProtectionGroupObjectParams instantiates a new AcropolisProtectionGroupObjectParams object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAcropolisProtectionGroupObjectParams(id NullableInt64) *AcropolisProtectionGroupObjectParams {
	this := AcropolisProtectionGroupObjectParams{}
	this.Id = id
	return &this
}

// NewAcropolisProtectionGroupObjectParamsWithDefaults instantiates a new AcropolisProtectionGroupObjectParams object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAcropolisProtectionGroupObjectParamsWithDefaults() *AcropolisProtectionGroupObjectParams {
	this := AcropolisProtectionGroupObjectParams{}
	return &this
}

// GetId returns the Id field value
// If the value is explicit nil, the zero value for int64 will be returned
func (o *AcropolisProtectionGroupObjectParams) GetId() int64 {
	if o == nil || o.Id.Get() == nil {
		var ret int64
		return ret
	}

	return *o.Id.Get()
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AcropolisProtectionGroupObjectParams) GetIdOk() (*int64, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Id.Get(), o.Id.IsSet()
}

// SetId sets field value
func (o *AcropolisProtectionGroupObjectParams) SetId(v int64) {
	o.Id.Set(&v)
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AcropolisProtectionGroupObjectParams) GetName() string {
	if o == nil || o.Name.Get() == nil {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AcropolisProtectionGroupObjectParams) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *AcropolisProtectionGroupObjectParams) HasName() bool {
	if o != nil && o.Name.IsSet() {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableString and assigns it to the Name field.
func (o *AcropolisProtectionGroupObjectParams) SetName(v string) {
	o.Name.Set(&v)
}
// SetNameNil sets the value for Name to be an explicit nil
func (o *AcropolisProtectionGroupObjectParams) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *AcropolisProtectionGroupObjectParams) UnsetName() {
	o.Name.Unset()
}

func (o AcropolisProtectionGroupObjectParams) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["id"] = o.Id.Get()
	}
	if o.Name.IsSet() {
		toSerialize["name"] = o.Name.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableAcropolisProtectionGroupObjectParams struct {
	value *AcropolisProtectionGroupObjectParams
	isSet bool
}

func (v NullableAcropolisProtectionGroupObjectParams) Get() *AcropolisProtectionGroupObjectParams {
	return v.value
}

func (v *NullableAcropolisProtectionGroupObjectParams) Set(val *AcropolisProtectionGroupObjectParams) {
	v.value = val
	v.isSet = true
}

func (v NullableAcropolisProtectionGroupObjectParams) IsSet() bool {
	return v.isSet
}

func (v *NullableAcropolisProtectionGroupObjectParams) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAcropolisProtectionGroupObjectParams(val *AcropolisProtectionGroupObjectParams) *NullableAcropolisProtectionGroupObjectParams {
	return &NullableAcropolisProtectionGroupObjectParams{value: val, isSet: true}
}

func (v NullableAcropolisProtectionGroupObjectParams) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAcropolisProtectionGroupObjectParams) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


