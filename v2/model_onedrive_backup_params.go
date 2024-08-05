/*
Cohesity REST API

Cohesity API provides a RESTful interface to access the various data management operations on Cohesity cluster and Helios.

API version: 2.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package v2

import (
	"encoding/json"
)

// checks if the OnedriveBackupParams type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OnedriveBackupParams{}

// OnedriveBackupParams Specifies OneDrive job parameters applicable for all Office365 Environment type Protection Sources in a Protection group.
type OnedriveBackupParams struct {
	FilePathFilter *FileFilteringPolicy `json:"filePathFilter,omitempty"`
	ShouldBackupOnedrive NullableBool `json:"shouldBackupOnedrive,omitempty"`
}

// NewOnedriveBackupParams instantiates a new OnedriveBackupParams object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOnedriveBackupParams() *OnedriveBackupParams {
	this := OnedriveBackupParams{}
	return &this
}

// NewOnedriveBackupParamsWithDefaults instantiates a new OnedriveBackupParams object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOnedriveBackupParamsWithDefaults() *OnedriveBackupParams {
	this := OnedriveBackupParams{}
	return &this
}

// GetFilePathFilter returns the FilePathFilter field value if set, zero value otherwise.
func (o *OnedriveBackupParams) GetFilePathFilter() FileFilteringPolicy {
	if o == nil || IsNil(o.FilePathFilter) {
		var ret FileFilteringPolicy
		return ret
	}
	return *o.FilePathFilter
}

// GetFilePathFilterOk returns a tuple with the FilePathFilter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OnedriveBackupParams) GetFilePathFilterOk() (*FileFilteringPolicy, bool) {
	if o == nil || IsNil(o.FilePathFilter) {
		return nil, false
	}
	return o.FilePathFilter, true
}

// HasFilePathFilter returns a boolean if a field has been set.
func (o *OnedriveBackupParams) HasFilePathFilter() bool {
	if o != nil && !IsNil(o.FilePathFilter) {
		return true
	}

	return false
}

// SetFilePathFilter gets a reference to the given FileFilteringPolicy and assigns it to the FilePathFilter field.
func (o *OnedriveBackupParams) SetFilePathFilter(v FileFilteringPolicy) {
	o.FilePathFilter = &v
}

// GetShouldBackupOnedrive returns the ShouldBackupOnedrive field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OnedriveBackupParams) GetShouldBackupOnedrive() bool {
	if o == nil || IsNil(o.ShouldBackupOnedrive.Get()) {
		var ret bool
		return ret
	}
	return *o.ShouldBackupOnedrive.Get()
}

// GetShouldBackupOnedriveOk returns a tuple with the ShouldBackupOnedrive field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OnedriveBackupParams) GetShouldBackupOnedriveOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.ShouldBackupOnedrive.Get(), o.ShouldBackupOnedrive.IsSet()
}

// HasShouldBackupOnedrive returns a boolean if a field has been set.
func (o *OnedriveBackupParams) HasShouldBackupOnedrive() bool {
	if o != nil && o.ShouldBackupOnedrive.IsSet() {
		return true
	}

	return false
}

// SetShouldBackupOnedrive gets a reference to the given NullableBool and assigns it to the ShouldBackupOnedrive field.
func (o *OnedriveBackupParams) SetShouldBackupOnedrive(v bool) {
	o.ShouldBackupOnedrive.Set(&v)
}
// SetShouldBackupOnedriveNil sets the value for ShouldBackupOnedrive to be an explicit nil
func (o *OnedriveBackupParams) SetShouldBackupOnedriveNil() {
	o.ShouldBackupOnedrive.Set(nil)
}

// UnsetShouldBackupOnedrive ensures that no value is present for ShouldBackupOnedrive, not even an explicit nil
func (o *OnedriveBackupParams) UnsetShouldBackupOnedrive() {
	o.ShouldBackupOnedrive.Unset()
}

func (o OnedriveBackupParams) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OnedriveBackupParams) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.FilePathFilter) {
		toSerialize["filePathFilter"] = o.FilePathFilter
	}
	if o.ShouldBackupOnedrive.IsSet() {
		toSerialize["shouldBackupOnedrive"] = o.ShouldBackupOnedrive.Get()
	}
	return toSerialize, nil
}

type NullableOnedriveBackupParams struct {
	value *OnedriveBackupParams
	isSet bool
}

func (v NullableOnedriveBackupParams) Get() *OnedriveBackupParams {
	return v.value
}

func (v *NullableOnedriveBackupParams) Set(val *OnedriveBackupParams) {
	v.value = val
	v.isSet = true
}

func (v NullableOnedriveBackupParams) IsSet() bool {
	return v.isSet
}

func (v *NullableOnedriveBackupParams) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOnedriveBackupParams(val *OnedriveBackupParams) *NullableOnedriveBackupParams {
	return &NullableOnedriveBackupParams{value: val, isSet: true}
}

func (v NullableOnedriveBackupParams) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOnedriveBackupParams) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


