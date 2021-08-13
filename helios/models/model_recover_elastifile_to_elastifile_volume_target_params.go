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

// RecoverElastifileToElastifileVolumeTargetParams Specifies the params of the Elastifile recovery target.
type RecoverElastifileToElastifileVolumeTargetParams struct {
	// Specifies the parameter whether the recovery should be performed to a new or the original Elastifile target.
	RecoverToNewSource bool `json:"recoverToNewSource"`
	// Specifies the new destination Source configuration parameters where the volumes will be recovered. This is mandatory if recoverToNewSource is set to true.
	NewSourceConfig NullableRecoverOtherNasToElastifileVolumeTargetParams `json:"newSourceConfig,omitempty"`
	// Specifies the Source configuration if volumes are being recovered to original Source. If not specified, all the configuration parameters will be retained.
	OriginalSourceConfig NullableOriginalElastifileTargetParams `json:"originalSourceConfig,omitempty"`
}

// NewRecoverElastifileToElastifileVolumeTargetParams instantiates a new RecoverElastifileToElastifileVolumeTargetParams object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRecoverElastifileToElastifileVolumeTargetParams(recoverToNewSource bool) *RecoverElastifileToElastifileVolumeTargetParams {
	this := RecoverElastifileToElastifileVolumeTargetParams{}
	this.RecoverToNewSource = recoverToNewSource
	return &this
}

// NewRecoverElastifileToElastifileVolumeTargetParamsWithDefaults instantiates a new RecoverElastifileToElastifileVolumeTargetParams object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRecoverElastifileToElastifileVolumeTargetParamsWithDefaults() *RecoverElastifileToElastifileVolumeTargetParams {
	this := RecoverElastifileToElastifileVolumeTargetParams{}
	return &this
}

// GetRecoverToNewSource returns the RecoverToNewSource field value
func (o *RecoverElastifileToElastifileVolumeTargetParams) GetRecoverToNewSource() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.RecoverToNewSource
}

// GetRecoverToNewSourceOk returns a tuple with the RecoverToNewSource field value
// and a boolean to check if the value has been set.
func (o *RecoverElastifileToElastifileVolumeTargetParams) GetRecoverToNewSourceOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.RecoverToNewSource, true
}

// SetRecoverToNewSource sets field value
func (o *RecoverElastifileToElastifileVolumeTargetParams) SetRecoverToNewSource(v bool) {
	o.RecoverToNewSource = v
}

// GetNewSourceConfig returns the NewSourceConfig field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RecoverElastifileToElastifileVolumeTargetParams) GetNewSourceConfig() RecoverOtherNasToElastifileVolumeTargetParams {
	if o == nil || o.NewSourceConfig.Get() == nil {
		var ret RecoverOtherNasToElastifileVolumeTargetParams
		return ret
	}
	return *o.NewSourceConfig.Get()
}

// GetNewSourceConfigOk returns a tuple with the NewSourceConfig field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RecoverElastifileToElastifileVolumeTargetParams) GetNewSourceConfigOk() (*RecoverOtherNasToElastifileVolumeTargetParams, bool) {
	if o == nil  {
		return nil, false
	}
	return o.NewSourceConfig.Get(), o.NewSourceConfig.IsSet()
}

// HasNewSourceConfig returns a boolean if a field has been set.
func (o *RecoverElastifileToElastifileVolumeTargetParams) HasNewSourceConfig() bool {
	if o != nil && o.NewSourceConfig.IsSet() {
		return true
	}

	return false
}

// SetNewSourceConfig gets a reference to the given NullableRecoverOtherNasToElastifileVolumeTargetParams and assigns it to the NewSourceConfig field.
func (o *RecoverElastifileToElastifileVolumeTargetParams) SetNewSourceConfig(v RecoverOtherNasToElastifileVolumeTargetParams) {
	o.NewSourceConfig.Set(&v)
}
// SetNewSourceConfigNil sets the value for NewSourceConfig to be an explicit nil
func (o *RecoverElastifileToElastifileVolumeTargetParams) SetNewSourceConfigNil() {
	o.NewSourceConfig.Set(nil)
}

// UnsetNewSourceConfig ensures that no value is present for NewSourceConfig, not even an explicit nil
func (o *RecoverElastifileToElastifileVolumeTargetParams) UnsetNewSourceConfig() {
	o.NewSourceConfig.Unset()
}

// GetOriginalSourceConfig returns the OriginalSourceConfig field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RecoverElastifileToElastifileVolumeTargetParams) GetOriginalSourceConfig() OriginalElastifileTargetParams {
	if o == nil || o.OriginalSourceConfig.Get() == nil {
		var ret OriginalElastifileTargetParams
		return ret
	}
	return *o.OriginalSourceConfig.Get()
}

// GetOriginalSourceConfigOk returns a tuple with the OriginalSourceConfig field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RecoverElastifileToElastifileVolumeTargetParams) GetOriginalSourceConfigOk() (*OriginalElastifileTargetParams, bool) {
	if o == nil  {
		return nil, false
	}
	return o.OriginalSourceConfig.Get(), o.OriginalSourceConfig.IsSet()
}

// HasOriginalSourceConfig returns a boolean if a field has been set.
func (o *RecoverElastifileToElastifileVolumeTargetParams) HasOriginalSourceConfig() bool {
	if o != nil && o.OriginalSourceConfig.IsSet() {
		return true
	}

	return false
}

// SetOriginalSourceConfig gets a reference to the given NullableOriginalElastifileTargetParams and assigns it to the OriginalSourceConfig field.
func (o *RecoverElastifileToElastifileVolumeTargetParams) SetOriginalSourceConfig(v OriginalElastifileTargetParams) {
	o.OriginalSourceConfig.Set(&v)
}
// SetOriginalSourceConfigNil sets the value for OriginalSourceConfig to be an explicit nil
func (o *RecoverElastifileToElastifileVolumeTargetParams) SetOriginalSourceConfigNil() {
	o.OriginalSourceConfig.Set(nil)
}

// UnsetOriginalSourceConfig ensures that no value is present for OriginalSourceConfig, not even an explicit nil
func (o *RecoverElastifileToElastifileVolumeTargetParams) UnsetOriginalSourceConfig() {
	o.OriginalSourceConfig.Unset()
}

func (o RecoverElastifileToElastifileVolumeTargetParams) MarshalJSON() ([]byte, error) {
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

type NullableRecoverElastifileToElastifileVolumeTargetParams struct {
	value *RecoverElastifileToElastifileVolumeTargetParams
	isSet bool
}

func (v NullableRecoverElastifileToElastifileVolumeTargetParams) Get() *RecoverElastifileToElastifileVolumeTargetParams {
	return v.value
}

func (v *NullableRecoverElastifileToElastifileVolumeTargetParams) Set(val *RecoverElastifileToElastifileVolumeTargetParams) {
	v.value = val
	v.isSet = true
}

func (v NullableRecoverElastifileToElastifileVolumeTargetParams) IsSet() bool {
	return v.isSet
}

func (v *NullableRecoverElastifileToElastifileVolumeTargetParams) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRecoverElastifileToElastifileVolumeTargetParams(val *RecoverElastifileToElastifileVolumeTargetParams) *NullableRecoverElastifileToElastifileVolumeTargetParams {
	return &NullableRecoverElastifileToElastifileVolumeTargetParams{value: val, isSet: true}
}

func (v NullableRecoverElastifileToElastifileVolumeTargetParams) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRecoverElastifileToElastifileVolumeTargetParams) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


