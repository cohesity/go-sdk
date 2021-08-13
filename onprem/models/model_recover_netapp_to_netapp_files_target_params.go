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

// RecoverNetappToNetappFilesTargetParams Specifies the params of the Netapp recovery target.
type RecoverNetappToNetappFilesTargetParams struct {
	// Specifies the parameter whether the recovery should be performed to a new or the original Netapp target.
	RecoverToNewSource bool `json:"recoverToNewSource"`
	// Specifies the new destination Source configuration parameters where the files will be recovered. This is mandatory if recoverToNewSource is set to true.
	NewSourceConfig NullableRecoverOtherNasToNetappFilesTargetParams `json:"newSourceConfig,omitempty"`
	// Specifies the Source configuration if files are being recovered to original Source. If not specified, all the configuration parameters will be retained.
	OriginalSourceConfig NullableOriginalNetappFilesTargetParams `json:"originalSourceConfig,omitempty"`
}

// NewRecoverNetappToNetappFilesTargetParams instantiates a new RecoverNetappToNetappFilesTargetParams object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRecoverNetappToNetappFilesTargetParams(recoverToNewSource bool) *RecoverNetappToNetappFilesTargetParams {
	this := RecoverNetappToNetappFilesTargetParams{}
	this.RecoverToNewSource = recoverToNewSource
	return &this
}

// NewRecoverNetappToNetappFilesTargetParamsWithDefaults instantiates a new RecoverNetappToNetappFilesTargetParams object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRecoverNetappToNetappFilesTargetParamsWithDefaults() *RecoverNetappToNetappFilesTargetParams {
	this := RecoverNetappToNetappFilesTargetParams{}
	return &this
}

// GetRecoverToNewSource returns the RecoverToNewSource field value
func (o *RecoverNetappToNetappFilesTargetParams) GetRecoverToNewSource() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.RecoverToNewSource
}

// GetRecoverToNewSourceOk returns a tuple with the RecoverToNewSource field value
// and a boolean to check if the value has been set.
func (o *RecoverNetappToNetappFilesTargetParams) GetRecoverToNewSourceOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.RecoverToNewSource, true
}

// SetRecoverToNewSource sets field value
func (o *RecoverNetappToNetappFilesTargetParams) SetRecoverToNewSource(v bool) {
	o.RecoverToNewSource = v
}

// GetNewSourceConfig returns the NewSourceConfig field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RecoverNetappToNetappFilesTargetParams) GetNewSourceConfig() RecoverOtherNasToNetappFilesTargetParams {
	if o == nil || o.NewSourceConfig.Get() == nil {
		var ret RecoverOtherNasToNetappFilesTargetParams
		return ret
	}
	return *o.NewSourceConfig.Get()
}

// GetNewSourceConfigOk returns a tuple with the NewSourceConfig field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RecoverNetappToNetappFilesTargetParams) GetNewSourceConfigOk() (*RecoverOtherNasToNetappFilesTargetParams, bool) {
	if o == nil  {
		return nil, false
	}
	return o.NewSourceConfig.Get(), o.NewSourceConfig.IsSet()
}

// HasNewSourceConfig returns a boolean if a field has been set.
func (o *RecoverNetappToNetappFilesTargetParams) HasNewSourceConfig() bool {
	if o != nil && o.NewSourceConfig.IsSet() {
		return true
	}

	return false
}

// SetNewSourceConfig gets a reference to the given NullableRecoverOtherNasToNetappFilesTargetParams and assigns it to the NewSourceConfig field.
func (o *RecoverNetappToNetappFilesTargetParams) SetNewSourceConfig(v RecoverOtherNasToNetappFilesTargetParams) {
	o.NewSourceConfig.Set(&v)
}
// SetNewSourceConfigNil sets the value for NewSourceConfig to be an explicit nil
func (o *RecoverNetappToNetappFilesTargetParams) SetNewSourceConfigNil() {
	o.NewSourceConfig.Set(nil)
}

// UnsetNewSourceConfig ensures that no value is present for NewSourceConfig, not even an explicit nil
func (o *RecoverNetappToNetappFilesTargetParams) UnsetNewSourceConfig() {
	o.NewSourceConfig.Unset()
}

// GetOriginalSourceConfig returns the OriginalSourceConfig field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RecoverNetappToNetappFilesTargetParams) GetOriginalSourceConfig() OriginalNetappFilesTargetParams {
	if o == nil || o.OriginalSourceConfig.Get() == nil {
		var ret OriginalNetappFilesTargetParams
		return ret
	}
	return *o.OriginalSourceConfig.Get()
}

// GetOriginalSourceConfigOk returns a tuple with the OriginalSourceConfig field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RecoverNetappToNetappFilesTargetParams) GetOriginalSourceConfigOk() (*OriginalNetappFilesTargetParams, bool) {
	if o == nil  {
		return nil, false
	}
	return o.OriginalSourceConfig.Get(), o.OriginalSourceConfig.IsSet()
}

// HasOriginalSourceConfig returns a boolean if a field has been set.
func (o *RecoverNetappToNetappFilesTargetParams) HasOriginalSourceConfig() bool {
	if o != nil && o.OriginalSourceConfig.IsSet() {
		return true
	}

	return false
}

// SetOriginalSourceConfig gets a reference to the given NullableOriginalNetappFilesTargetParams and assigns it to the OriginalSourceConfig field.
func (o *RecoverNetappToNetappFilesTargetParams) SetOriginalSourceConfig(v OriginalNetappFilesTargetParams) {
	o.OriginalSourceConfig.Set(&v)
}
// SetOriginalSourceConfigNil sets the value for OriginalSourceConfig to be an explicit nil
func (o *RecoverNetappToNetappFilesTargetParams) SetOriginalSourceConfigNil() {
	o.OriginalSourceConfig.Set(nil)
}

// UnsetOriginalSourceConfig ensures that no value is present for OriginalSourceConfig, not even an explicit nil
func (o *RecoverNetappToNetappFilesTargetParams) UnsetOriginalSourceConfig() {
	o.OriginalSourceConfig.Unset()
}

func (o RecoverNetappToNetappFilesTargetParams) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["recoverToNewSource"] = o.RecoverToNewSource
	}
	if o.NewSourceConfig.IsSet() {
		toSerialize["newSourceConfig"] = o.NewSourceConfig.Get()
	}
	if o.OriginalSourceConfig.IsSet() {
		toSerialize["originalSourceConfig"] = o.OriginalSourceConfig.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableRecoverNetappToNetappFilesTargetParams struct {
	value *RecoverNetappToNetappFilesTargetParams
	isSet bool
}

func (v NullableRecoverNetappToNetappFilesTargetParams) Get() *RecoverNetappToNetappFilesTargetParams {
	return v.value
}

func (v *NullableRecoverNetappToNetappFilesTargetParams) Set(val *RecoverNetappToNetappFilesTargetParams) {
	v.value = val
	v.isSet = true
}

func (v NullableRecoverNetappToNetappFilesTargetParams) IsSet() bool {
	return v.isSet
}

func (v *NullableRecoverNetappToNetappFilesTargetParams) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRecoverNetappToNetappFilesTargetParams(val *RecoverNetappToNetappFilesTargetParams) *NullableRecoverNetappToNetappFilesTargetParams {
	return &NullableRecoverNetappToNetappFilesTargetParams{value: val, isSet: true}
}

func (v NullableRecoverNetappToNetappFilesTargetParams) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRecoverNetappToNetappFilesTargetParams) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}




func (o RecoverNetappToNetappFilesTargetParams) Print() {
		if byteArray, err := o.MarshalJSON(); err != nil {
				fmt.Println(err)
		} else {
				fmt.Println(string(byteArray))
		}
}