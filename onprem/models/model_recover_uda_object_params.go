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

// RecoverUdaObjectParams Specifies details of objects to be recovered.
type RecoverUdaObjectParams struct {
	// Specifies the fully qualified name of the object to be restored.
	ObjectName NullableString `json:"objectName,omitempty"`
	// Set to true to overwrite an existing object at the destination. If set to false, and the same object exists at the destination, then recovery will fail for that object.
	Overwrite NullableBool `json:"overwrite,omitempty"`
	// Specifies the new name to which the object should be renamed to after the recovery.
	RenameTo NullableString `json:"renameTo,omitempty"`
}

// NewRecoverUdaObjectParams instantiates a new RecoverUdaObjectParams object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRecoverUdaObjectParams() *RecoverUdaObjectParams {
	this := RecoverUdaObjectParams{}
	return &this
}

// NewRecoverUdaObjectParamsWithDefaults instantiates a new RecoverUdaObjectParams object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRecoverUdaObjectParamsWithDefaults() *RecoverUdaObjectParams {
	this := RecoverUdaObjectParams{}
	return &this
}

// GetObjectName returns the ObjectName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RecoverUdaObjectParams) GetObjectName() string {
	if o == nil || o.ObjectName.Get() == nil {
		var ret string
		return ret
	}
	return *o.ObjectName.Get()
}

// GetObjectNameOk returns a tuple with the ObjectName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RecoverUdaObjectParams) GetObjectNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.ObjectName.Get(), o.ObjectName.IsSet()
}

// HasObjectName returns a boolean if a field has been set.
func (o *RecoverUdaObjectParams) HasObjectName() bool {
	if o != nil && o.ObjectName.IsSet() {
		return true
	}

	return false
}

// SetObjectName gets a reference to the given NullableString and assigns it to the ObjectName field.
func (o *RecoverUdaObjectParams) SetObjectName(v string) {
	o.ObjectName.Set(&v)
}
// SetObjectNameNil sets the value for ObjectName to be an explicit nil
func (o *RecoverUdaObjectParams) SetObjectNameNil() {
	o.ObjectName.Set(nil)
}

// UnsetObjectName ensures that no value is present for ObjectName, not even an explicit nil
func (o *RecoverUdaObjectParams) UnsetObjectName() {
	o.ObjectName.Unset()
}

// GetOverwrite returns the Overwrite field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RecoverUdaObjectParams) GetOverwrite() bool {
	if o == nil || o.Overwrite.Get() == nil {
		var ret bool
		return ret
	}
	return *o.Overwrite.Get()
}

// GetOverwriteOk returns a tuple with the Overwrite field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RecoverUdaObjectParams) GetOverwriteOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Overwrite.Get(), o.Overwrite.IsSet()
}

// HasOverwrite returns a boolean if a field has been set.
func (o *RecoverUdaObjectParams) HasOverwrite() bool {
	if o != nil && o.Overwrite.IsSet() {
		return true
	}

	return false
}

// SetOverwrite gets a reference to the given NullableBool and assigns it to the Overwrite field.
func (o *RecoverUdaObjectParams) SetOverwrite(v bool) {
	o.Overwrite.Set(&v)
}
// SetOverwriteNil sets the value for Overwrite to be an explicit nil
func (o *RecoverUdaObjectParams) SetOverwriteNil() {
	o.Overwrite.Set(nil)
}

// UnsetOverwrite ensures that no value is present for Overwrite, not even an explicit nil
func (o *RecoverUdaObjectParams) UnsetOverwrite() {
	o.Overwrite.Unset()
}

// GetRenameTo returns the RenameTo field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RecoverUdaObjectParams) GetRenameTo() string {
	if o == nil || o.RenameTo.Get() == nil {
		var ret string
		return ret
	}
	return *o.RenameTo.Get()
}

// GetRenameToOk returns a tuple with the RenameTo field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RecoverUdaObjectParams) GetRenameToOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.RenameTo.Get(), o.RenameTo.IsSet()
}

// HasRenameTo returns a boolean if a field has been set.
func (o *RecoverUdaObjectParams) HasRenameTo() bool {
	if o != nil && o.RenameTo.IsSet() {
		return true
	}

	return false
}

// SetRenameTo gets a reference to the given NullableString and assigns it to the RenameTo field.
func (o *RecoverUdaObjectParams) SetRenameTo(v string) {
	o.RenameTo.Set(&v)
}
// SetRenameToNil sets the value for RenameTo to be an explicit nil
func (o *RecoverUdaObjectParams) SetRenameToNil() {
	o.RenameTo.Set(nil)
}

// UnsetRenameTo ensures that no value is present for RenameTo, not even an explicit nil
func (o *RecoverUdaObjectParams) UnsetRenameTo() {
	o.RenameTo.Unset()
}

func (o RecoverUdaObjectParams) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ObjectName.IsSet() {
		toSerialize["objectName"] = o.ObjectName.Get()
	}
	if o.Overwrite.IsSet() {
		toSerialize["overwrite"] = o.Overwrite.Get()
	}
	if o.RenameTo.IsSet() {
		toSerialize["renameTo"] = o.RenameTo.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableRecoverUdaObjectParams struct {
	value *RecoverUdaObjectParams
	isSet bool
}

func (v NullableRecoverUdaObjectParams) Get() *RecoverUdaObjectParams {
	return v.value
}

func (v *NullableRecoverUdaObjectParams) Set(val *RecoverUdaObjectParams) {
	v.value = val
	v.isSet = true
}

func (v NullableRecoverUdaObjectParams) IsSet() bool {
	return v.isSet
}

func (v *NullableRecoverUdaObjectParams) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRecoverUdaObjectParams(val *RecoverUdaObjectParams) *NullableRecoverUdaObjectParams {
	return &NullableRecoverUdaObjectParams{value: val, isSet: true}
}

func (v NullableRecoverUdaObjectParams) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRecoverUdaObjectParams) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}




func (o RecoverUdaObjectParams) Print() {
		if byteArray, err := o.MarshalJSON(); err != nil {
				fmt.Println(err)
		} else {
				fmt.Println(string(byteArray))
		}
}