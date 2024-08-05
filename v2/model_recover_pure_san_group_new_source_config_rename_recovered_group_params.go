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

// checks if the RecoverPureSanGroupNewSourceConfigRenameRecoveredGroupParams type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RecoverPureSanGroupNewSourceConfigRenameRecoveredGroupParams{}

// RecoverPureSanGroupNewSourceConfigRenameRecoveredGroupParams Specifies params to rename the recovered SAN group. If not specified, the original names of the group are preserved.
type RecoverPureSanGroupNewSourceConfigRenameRecoveredGroupParams struct {
	// Specifies the prefix string to be added to recovered or cloned object names.
	Prefix NullableString `json:"prefix,omitempty"`
	// Specifies the suffix string to be added to recovered or cloned object names.
	Suffix NullableString `json:"suffix,omitempty"`
}

// NewRecoverPureSanGroupNewSourceConfigRenameRecoveredGroupParams instantiates a new RecoverPureSanGroupNewSourceConfigRenameRecoveredGroupParams object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRecoverPureSanGroupNewSourceConfigRenameRecoveredGroupParams() *RecoverPureSanGroupNewSourceConfigRenameRecoveredGroupParams {
	this := RecoverPureSanGroupNewSourceConfigRenameRecoveredGroupParams{}
	return &this
}

// NewRecoverPureSanGroupNewSourceConfigRenameRecoveredGroupParamsWithDefaults instantiates a new RecoverPureSanGroupNewSourceConfigRenameRecoveredGroupParams object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRecoverPureSanGroupNewSourceConfigRenameRecoveredGroupParamsWithDefaults() *RecoverPureSanGroupNewSourceConfigRenameRecoveredGroupParams {
	this := RecoverPureSanGroupNewSourceConfigRenameRecoveredGroupParams{}
	return &this
}

// GetPrefix returns the Prefix field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RecoverPureSanGroupNewSourceConfigRenameRecoveredGroupParams) GetPrefix() string {
	if o == nil || IsNil(o.Prefix.Get()) {
		var ret string
		return ret
	}
	return *o.Prefix.Get()
}

// GetPrefixOk returns a tuple with the Prefix field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RecoverPureSanGroupNewSourceConfigRenameRecoveredGroupParams) GetPrefixOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Prefix.Get(), o.Prefix.IsSet()
}

// HasPrefix returns a boolean if a field has been set.
func (o *RecoverPureSanGroupNewSourceConfigRenameRecoveredGroupParams) HasPrefix() bool {
	if o != nil && o.Prefix.IsSet() {
		return true
	}

	return false
}

// SetPrefix gets a reference to the given NullableString and assigns it to the Prefix field.
func (o *RecoverPureSanGroupNewSourceConfigRenameRecoveredGroupParams) SetPrefix(v string) {
	o.Prefix.Set(&v)
}
// SetPrefixNil sets the value for Prefix to be an explicit nil
func (o *RecoverPureSanGroupNewSourceConfigRenameRecoveredGroupParams) SetPrefixNil() {
	o.Prefix.Set(nil)
}

// UnsetPrefix ensures that no value is present for Prefix, not even an explicit nil
func (o *RecoverPureSanGroupNewSourceConfigRenameRecoveredGroupParams) UnsetPrefix() {
	o.Prefix.Unset()
}

// GetSuffix returns the Suffix field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RecoverPureSanGroupNewSourceConfigRenameRecoveredGroupParams) GetSuffix() string {
	if o == nil || IsNil(o.Suffix.Get()) {
		var ret string
		return ret
	}
	return *o.Suffix.Get()
}

// GetSuffixOk returns a tuple with the Suffix field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RecoverPureSanGroupNewSourceConfigRenameRecoveredGroupParams) GetSuffixOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Suffix.Get(), o.Suffix.IsSet()
}

// HasSuffix returns a boolean if a field has been set.
func (o *RecoverPureSanGroupNewSourceConfigRenameRecoveredGroupParams) HasSuffix() bool {
	if o != nil && o.Suffix.IsSet() {
		return true
	}

	return false
}

// SetSuffix gets a reference to the given NullableString and assigns it to the Suffix field.
func (o *RecoverPureSanGroupNewSourceConfigRenameRecoveredGroupParams) SetSuffix(v string) {
	o.Suffix.Set(&v)
}
// SetSuffixNil sets the value for Suffix to be an explicit nil
func (o *RecoverPureSanGroupNewSourceConfigRenameRecoveredGroupParams) SetSuffixNil() {
	o.Suffix.Set(nil)
}

// UnsetSuffix ensures that no value is present for Suffix, not even an explicit nil
func (o *RecoverPureSanGroupNewSourceConfigRenameRecoveredGroupParams) UnsetSuffix() {
	o.Suffix.Unset()
}

func (o RecoverPureSanGroupNewSourceConfigRenameRecoveredGroupParams) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RecoverPureSanGroupNewSourceConfigRenameRecoveredGroupParams) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Prefix.IsSet() {
		toSerialize["prefix"] = o.Prefix.Get()
	}
	if o.Suffix.IsSet() {
		toSerialize["suffix"] = o.Suffix.Get()
	}
	return toSerialize, nil
}

type NullableRecoverPureSanGroupNewSourceConfigRenameRecoveredGroupParams struct {
	value *RecoverPureSanGroupNewSourceConfigRenameRecoveredGroupParams
	isSet bool
}

func (v NullableRecoverPureSanGroupNewSourceConfigRenameRecoveredGroupParams) Get() *RecoverPureSanGroupNewSourceConfigRenameRecoveredGroupParams {
	return v.value
}

func (v *NullableRecoverPureSanGroupNewSourceConfigRenameRecoveredGroupParams) Set(val *RecoverPureSanGroupNewSourceConfigRenameRecoveredGroupParams) {
	v.value = val
	v.isSet = true
}

func (v NullableRecoverPureSanGroupNewSourceConfigRenameRecoveredGroupParams) IsSet() bool {
	return v.isSet
}

func (v *NullableRecoverPureSanGroupNewSourceConfigRenameRecoveredGroupParams) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRecoverPureSanGroupNewSourceConfigRenameRecoveredGroupParams(val *RecoverPureSanGroupNewSourceConfigRenameRecoveredGroupParams) *NullableRecoverPureSanGroupNewSourceConfigRenameRecoveredGroupParams {
	return &NullableRecoverPureSanGroupNewSourceConfigRenameRecoveredGroupParams{value: val, isSet: true}
}

func (v NullableRecoverPureSanGroupNewSourceConfigRenameRecoveredGroupParams) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRecoverPureSanGroupNewSourceConfigRenameRecoveredGroupParams) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


