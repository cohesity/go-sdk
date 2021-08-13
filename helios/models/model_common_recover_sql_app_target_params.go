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

// CommonRecoverSqlAppTargetParams Specifies the snapshot parameters to recover Sql databases.
type CommonRecoverSqlAppTargetParams struct {
	// Specifies the parameter whether the recovery should be performed to a new sources or an original Source Target.
	RecoverToNewSource bool `json:"recoverToNewSource"`
	// Specifies the destination Source configuration parameters where the databases will be recovered. This is mandatory if recoverToNewSource is set to true.
	NewSourceConfig NullableRecoverSqlAppNewSourceConfig `json:"newSourceConfig,omitempty"`
	// Specifies the Source configuration if databases are being recovered to Original Source. If not specified, all the configuration parameters will be retained.
	OriginalSourceConfig NullableRecoverSqlAppOriginalSourceConfig `json:"originalSourceConfig,omitempty"`
}

// NewCommonRecoverSqlAppTargetParams instantiates a new CommonRecoverSqlAppTargetParams object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCommonRecoverSqlAppTargetParams(recoverToNewSource bool) *CommonRecoverSqlAppTargetParams {
	this := CommonRecoverSqlAppTargetParams{}
	this.RecoverToNewSource = recoverToNewSource
	return &this
}

// NewCommonRecoverSqlAppTargetParamsWithDefaults instantiates a new CommonRecoverSqlAppTargetParams object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCommonRecoverSqlAppTargetParamsWithDefaults() *CommonRecoverSqlAppTargetParams {
	this := CommonRecoverSqlAppTargetParams{}
	return &this
}

// GetRecoverToNewSource returns the RecoverToNewSource field value
func (o *CommonRecoverSqlAppTargetParams) GetRecoverToNewSource() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.RecoverToNewSource
}

// GetRecoverToNewSourceOk returns a tuple with the RecoverToNewSource field value
// and a boolean to check if the value has been set.
func (o *CommonRecoverSqlAppTargetParams) GetRecoverToNewSourceOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.RecoverToNewSource, true
}

// SetRecoverToNewSource sets field value
func (o *CommonRecoverSqlAppTargetParams) SetRecoverToNewSource(v bool) {
	o.RecoverToNewSource = v
}

// GetNewSourceConfig returns the NewSourceConfig field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CommonRecoverSqlAppTargetParams) GetNewSourceConfig() RecoverSqlAppNewSourceConfig {
	if o == nil || o.NewSourceConfig.Get() == nil {
		var ret RecoverSqlAppNewSourceConfig
		return ret
	}
	return *o.NewSourceConfig.Get()
}

// GetNewSourceConfigOk returns a tuple with the NewSourceConfig field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CommonRecoverSqlAppTargetParams) GetNewSourceConfigOk() (*RecoverSqlAppNewSourceConfig, bool) {
	if o == nil  {
		return nil, false
	}
	return o.NewSourceConfig.Get(), o.NewSourceConfig.IsSet()
}

// HasNewSourceConfig returns a boolean if a field has been set.
func (o *CommonRecoverSqlAppTargetParams) HasNewSourceConfig() bool {
	if o != nil && o.NewSourceConfig.IsSet() {
		return true
	}

	return false
}

// SetNewSourceConfig gets a reference to the given NullableRecoverSqlAppNewSourceConfig and assigns it to the NewSourceConfig field.
func (o *CommonRecoverSqlAppTargetParams) SetNewSourceConfig(v RecoverSqlAppNewSourceConfig) {
	o.NewSourceConfig.Set(&v)
}
// SetNewSourceConfigNil sets the value for NewSourceConfig to be an explicit nil
func (o *CommonRecoverSqlAppTargetParams) SetNewSourceConfigNil() {
	o.NewSourceConfig.Set(nil)
}

// UnsetNewSourceConfig ensures that no value is present for NewSourceConfig, not even an explicit nil
func (o *CommonRecoverSqlAppTargetParams) UnsetNewSourceConfig() {
	o.NewSourceConfig.Unset()
}

// GetOriginalSourceConfig returns the OriginalSourceConfig field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CommonRecoverSqlAppTargetParams) GetOriginalSourceConfig() RecoverSqlAppOriginalSourceConfig {
	if o == nil || o.OriginalSourceConfig.Get() == nil {
		var ret RecoverSqlAppOriginalSourceConfig
		return ret
	}
	return *o.OriginalSourceConfig.Get()
}

// GetOriginalSourceConfigOk returns a tuple with the OriginalSourceConfig field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CommonRecoverSqlAppTargetParams) GetOriginalSourceConfigOk() (*RecoverSqlAppOriginalSourceConfig, bool) {
	if o == nil  {
		return nil, false
	}
	return o.OriginalSourceConfig.Get(), o.OriginalSourceConfig.IsSet()
}

// HasOriginalSourceConfig returns a boolean if a field has been set.
func (o *CommonRecoverSqlAppTargetParams) HasOriginalSourceConfig() bool {
	if o != nil && o.OriginalSourceConfig.IsSet() {
		return true
	}

	return false
}

// SetOriginalSourceConfig gets a reference to the given NullableRecoverSqlAppOriginalSourceConfig and assigns it to the OriginalSourceConfig field.
func (o *CommonRecoverSqlAppTargetParams) SetOriginalSourceConfig(v RecoverSqlAppOriginalSourceConfig) {
	o.OriginalSourceConfig.Set(&v)
}
// SetOriginalSourceConfigNil sets the value for OriginalSourceConfig to be an explicit nil
func (o *CommonRecoverSqlAppTargetParams) SetOriginalSourceConfigNil() {
	o.OriginalSourceConfig.Set(nil)
}

// UnsetOriginalSourceConfig ensures that no value is present for OriginalSourceConfig, not even an explicit nil
func (o *CommonRecoverSqlAppTargetParams) UnsetOriginalSourceConfig() {
	o.OriginalSourceConfig.Unset()
}

func (o CommonRecoverSqlAppTargetParams) MarshalJSON() ([]byte, error) {
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

type NullableCommonRecoverSqlAppTargetParams struct {
	value *CommonRecoverSqlAppTargetParams
	isSet bool
}

func (v NullableCommonRecoverSqlAppTargetParams) Get() *CommonRecoverSqlAppTargetParams {
	return v.value
}

func (v *NullableCommonRecoverSqlAppTargetParams) Set(val *CommonRecoverSqlAppTargetParams) {
	v.value = val
	v.isSet = true
}

func (v NullableCommonRecoverSqlAppTargetParams) IsSet() bool {
	return v.isSet
}

func (v *NullableCommonRecoverSqlAppTargetParams) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCommonRecoverSqlAppTargetParams(val *CommonRecoverSqlAppTargetParams) *NullableCommonRecoverSqlAppTargetParams {
	return &NullableCommonRecoverSqlAppTargetParams{value: val, isSet: true}
}

func (v NullableCommonRecoverSqlAppTargetParams) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCommonRecoverSqlAppTargetParams) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


