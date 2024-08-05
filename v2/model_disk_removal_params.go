/*
Cohesity REST API

Cohesity API provides a RESTful interface to access the various data management operations on Cohesity cluster and Helios.

API version: 2.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package v2

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the DiskRemovalParams type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DiskRemovalParams{}

// DiskRemovalParams Specifies parameters to initiate/cancel disk removal.
type DiskRemovalParams struct {
	// If true, cancels disk removal which is already in progress.
	Cancel NullableBool `json:"cancel"`
	// Specifies whether request is for pre-check validations only
	IsValidateOnly NullableBool `json:"isValidateOnly,omitempty"`
}

type _DiskRemovalParams DiskRemovalParams

// NewDiskRemovalParams instantiates a new DiskRemovalParams object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDiskRemovalParams(cancel NullableBool) *DiskRemovalParams {
	this := DiskRemovalParams{}
	this.Cancel = cancel
	var isValidateOnly bool = false
	this.IsValidateOnly = *NewNullableBool(&isValidateOnly)
	return &this
}

// NewDiskRemovalParamsWithDefaults instantiates a new DiskRemovalParams object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDiskRemovalParamsWithDefaults() *DiskRemovalParams {
	this := DiskRemovalParams{}
	var isValidateOnly bool = false
	this.IsValidateOnly = *NewNullableBool(&isValidateOnly)
	return &this
}

// GetCancel returns the Cancel field value
// If the value is explicit nil, the zero value for bool will be returned
func (o *DiskRemovalParams) GetCancel() bool {
	if o == nil || o.Cancel.Get() == nil {
		var ret bool
		return ret
	}

	return *o.Cancel.Get()
}

// GetCancelOk returns a tuple with the Cancel field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DiskRemovalParams) GetCancelOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.Cancel.Get(), o.Cancel.IsSet()
}

// SetCancel sets field value
func (o *DiskRemovalParams) SetCancel(v bool) {
	o.Cancel.Set(&v)
}

// GetIsValidateOnly returns the IsValidateOnly field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DiskRemovalParams) GetIsValidateOnly() bool {
	if o == nil || IsNil(o.IsValidateOnly.Get()) {
		var ret bool
		return ret
	}
	return *o.IsValidateOnly.Get()
}

// GetIsValidateOnlyOk returns a tuple with the IsValidateOnly field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DiskRemovalParams) GetIsValidateOnlyOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.IsValidateOnly.Get(), o.IsValidateOnly.IsSet()
}

// HasIsValidateOnly returns a boolean if a field has been set.
func (o *DiskRemovalParams) HasIsValidateOnly() bool {
	if o != nil && o.IsValidateOnly.IsSet() {
		return true
	}

	return false
}

// SetIsValidateOnly gets a reference to the given NullableBool and assigns it to the IsValidateOnly field.
func (o *DiskRemovalParams) SetIsValidateOnly(v bool) {
	o.IsValidateOnly.Set(&v)
}
// SetIsValidateOnlyNil sets the value for IsValidateOnly to be an explicit nil
func (o *DiskRemovalParams) SetIsValidateOnlyNil() {
	o.IsValidateOnly.Set(nil)
}

// UnsetIsValidateOnly ensures that no value is present for IsValidateOnly, not even an explicit nil
func (o *DiskRemovalParams) UnsetIsValidateOnly() {
	o.IsValidateOnly.Unset()
}

func (o DiskRemovalParams) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DiskRemovalParams) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["cancel"] = o.Cancel.Get()
	if o.IsValidateOnly.IsSet() {
		toSerialize["isValidateOnly"] = o.IsValidateOnly.Get()
	}
	return toSerialize, nil
}

func (o *DiskRemovalParams) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"cancel",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varDiskRemovalParams := _DiskRemovalParams{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varDiskRemovalParams)

	if err != nil {
		return err
	}

	*o = DiskRemovalParams(varDiskRemovalParams)

	return err
}

type NullableDiskRemovalParams struct {
	value *DiskRemovalParams
	isSet bool
}

func (v NullableDiskRemovalParams) Get() *DiskRemovalParams {
	return v.value
}

func (v *NullableDiskRemovalParams) Set(val *DiskRemovalParams) {
	v.value = val
	v.isSet = true
}

func (v NullableDiskRemovalParams) IsSet() bool {
	return v.isSet
}

func (v *NullableDiskRemovalParams) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDiskRemovalParams(val *DiskRemovalParams) *NullableDiskRemovalParams {
	return &NullableDiskRemovalParams{value: val, isSet: true}
}

func (v NullableDiskRemovalParams) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDiskRemovalParams) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


