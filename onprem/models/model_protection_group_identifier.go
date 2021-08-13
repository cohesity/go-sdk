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

// ProtectionGroupIdentifier Specifies Protection Group Identifier.
type ProtectionGroupIdentifier struct {
	// Specifies Protection Group id.
	ProtectionGroupId NullableString `json:"protectionGroupId,omitempty"`
	// Specifies Protection Group name.
	ProtectionGroupName NullableString `json:"protectionGroupName,omitempty"`
}

// NewProtectionGroupIdentifier instantiates a new ProtectionGroupIdentifier object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProtectionGroupIdentifier() *ProtectionGroupIdentifier {
	this := ProtectionGroupIdentifier{}
	return &this
}

// NewProtectionGroupIdentifierWithDefaults instantiates a new ProtectionGroupIdentifier object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProtectionGroupIdentifierWithDefaults() *ProtectionGroupIdentifier {
	this := ProtectionGroupIdentifier{}
	return &this
}

// GetProtectionGroupId returns the ProtectionGroupId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ProtectionGroupIdentifier) GetProtectionGroupId() string {
	if o == nil || o.ProtectionGroupId.Get() == nil {
		var ret string
		return ret
	}
	return *o.ProtectionGroupId.Get()
}

// GetProtectionGroupIdOk returns a tuple with the ProtectionGroupId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ProtectionGroupIdentifier) GetProtectionGroupIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.ProtectionGroupId.Get(), o.ProtectionGroupId.IsSet()
}

// HasProtectionGroupId returns a boolean if a field has been set.
func (o *ProtectionGroupIdentifier) HasProtectionGroupId() bool {
	if o != nil && o.ProtectionGroupId.IsSet() {
		return true
	}

	return false
}

// SetProtectionGroupId gets a reference to the given NullableString and assigns it to the ProtectionGroupId field.
func (o *ProtectionGroupIdentifier) SetProtectionGroupId(v string) {
	o.ProtectionGroupId.Set(&v)
}
// SetProtectionGroupIdNil sets the value for ProtectionGroupId to be an explicit nil
func (o *ProtectionGroupIdentifier) SetProtectionGroupIdNil() {
	o.ProtectionGroupId.Set(nil)
}

// UnsetProtectionGroupId ensures that no value is present for ProtectionGroupId, not even an explicit nil
func (o *ProtectionGroupIdentifier) UnsetProtectionGroupId() {
	o.ProtectionGroupId.Unset()
}

// GetProtectionGroupName returns the ProtectionGroupName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ProtectionGroupIdentifier) GetProtectionGroupName() string {
	if o == nil || o.ProtectionGroupName.Get() == nil {
		var ret string
		return ret
	}
	return *o.ProtectionGroupName.Get()
}

// GetProtectionGroupNameOk returns a tuple with the ProtectionGroupName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ProtectionGroupIdentifier) GetProtectionGroupNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.ProtectionGroupName.Get(), o.ProtectionGroupName.IsSet()
}

// HasProtectionGroupName returns a boolean if a field has been set.
func (o *ProtectionGroupIdentifier) HasProtectionGroupName() bool {
	if o != nil && o.ProtectionGroupName.IsSet() {
		return true
	}

	return false
}

// SetProtectionGroupName gets a reference to the given NullableString and assigns it to the ProtectionGroupName field.
func (o *ProtectionGroupIdentifier) SetProtectionGroupName(v string) {
	o.ProtectionGroupName.Set(&v)
}
// SetProtectionGroupNameNil sets the value for ProtectionGroupName to be an explicit nil
func (o *ProtectionGroupIdentifier) SetProtectionGroupNameNil() {
	o.ProtectionGroupName.Set(nil)
}

// UnsetProtectionGroupName ensures that no value is present for ProtectionGroupName, not even an explicit nil
func (o *ProtectionGroupIdentifier) UnsetProtectionGroupName() {
	o.ProtectionGroupName.Unset()
}

func (o ProtectionGroupIdentifier) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ProtectionGroupId.IsSet() {
		toSerialize["protectionGroupId"] = o.ProtectionGroupId.Get()
	}
	if o.ProtectionGroupName.IsSet() {
		toSerialize["protectionGroupName"] = o.ProtectionGroupName.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableProtectionGroupIdentifier struct {
	value *ProtectionGroupIdentifier
	isSet bool
}

func (v NullableProtectionGroupIdentifier) Get() *ProtectionGroupIdentifier {
	return v.value
}

func (v *NullableProtectionGroupIdentifier) Set(val *ProtectionGroupIdentifier) {
	v.value = val
	v.isSet = true
}

func (v NullableProtectionGroupIdentifier) IsSet() bool {
	return v.isSet
}

func (v *NullableProtectionGroupIdentifier) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProtectionGroupIdentifier(val *ProtectionGroupIdentifier) *NullableProtectionGroupIdentifier {
	return &NullableProtectionGroupIdentifier{value: val, isSet: true}
}

func (v NullableProtectionGroupIdentifier) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProtectionGroupIdentifier) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}




func (o ProtectionGroupIdentifier) Print() {
		if byteArray, err := o.MarshalJSON(); err != nil {
				fmt.Println(err)
		} else {
				fmt.Println(string(byteArray))
		}
}