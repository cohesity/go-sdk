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

// RecoverPureParams Specifies the recovery options specific to Pure environment.
type RecoverPureParams struct {
	// Specifies the list of recover object parameters.
	Objects []CommonRecoverObjectSnapshotParams `json:"objects"`
	// Specifies the type of recovery action to be performed. The corresponding recovery action params must be filled out.
	RecoveryAction string `json:"recoveryAction"`
	// Specifies the parameters to recover SAN Volume.
	RecoverSanVolumeParams NullableRecoverPureSanVolumeParams `json:"recoverSanVolumeParams,omitempty"`
}

// NewRecoverPureParams instantiates a new RecoverPureParams object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRecoverPureParams(objects []CommonRecoverObjectSnapshotParams, recoveryAction string) *RecoverPureParams {
	this := RecoverPureParams{}
	this.Objects = objects
	this.RecoveryAction = recoveryAction
	return &this
}

// NewRecoverPureParamsWithDefaults instantiates a new RecoverPureParams object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRecoverPureParamsWithDefaults() *RecoverPureParams {
	this := RecoverPureParams{}
	return &this
}

// GetObjects returns the Objects field value
// If the value is explicit nil, the zero value for []CommonRecoverObjectSnapshotParams will be returned
func (o *RecoverPureParams) GetObjects() []CommonRecoverObjectSnapshotParams {
	if o == nil {
		var ret []CommonRecoverObjectSnapshotParams
		return ret
	}

	return o.Objects
}

// GetObjectsOk returns a tuple with the Objects field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RecoverPureParams) GetObjectsOk() (*[]CommonRecoverObjectSnapshotParams, bool) {
	if o == nil || o.Objects == nil {
		return nil, false
	}
	return &o.Objects, true
}

// SetObjects sets field value
func (o *RecoverPureParams) SetObjects(v []CommonRecoverObjectSnapshotParams) {
	o.Objects = v
}

// GetRecoveryAction returns the RecoveryAction field value
func (o *RecoverPureParams) GetRecoveryAction() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RecoveryAction
}

// GetRecoveryActionOk returns a tuple with the RecoveryAction field value
// and a boolean to check if the value has been set.
func (o *RecoverPureParams) GetRecoveryActionOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.RecoveryAction, true
}

// SetRecoveryAction sets field value
func (o *RecoverPureParams) SetRecoveryAction(v string) {
	o.RecoveryAction = v
}

// GetRecoverSanVolumeParams returns the RecoverSanVolumeParams field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RecoverPureParams) GetRecoverSanVolumeParams() RecoverPureSanVolumeParams {
	if o == nil || o.RecoverSanVolumeParams.Get() == nil {
		var ret RecoverPureSanVolumeParams
		return ret
	}
	return *o.RecoverSanVolumeParams.Get()
}

// GetRecoverSanVolumeParamsOk returns a tuple with the RecoverSanVolumeParams field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RecoverPureParams) GetRecoverSanVolumeParamsOk() (*RecoverPureSanVolumeParams, bool) {
	if o == nil  {
		return nil, false
	}
	return o.RecoverSanVolumeParams.Get(), o.RecoverSanVolumeParams.IsSet()
}

// HasRecoverSanVolumeParams returns a boolean if a field has been set.
func (o *RecoverPureParams) HasRecoverSanVolumeParams() bool {
	if o != nil && o.RecoverSanVolumeParams.IsSet() {
		return true
	}

	return false
}

// SetRecoverSanVolumeParams gets a reference to the given NullableRecoverPureSanVolumeParams and assigns it to the RecoverSanVolumeParams field.
func (o *RecoverPureParams) SetRecoverSanVolumeParams(v RecoverPureSanVolumeParams) {
	o.RecoverSanVolumeParams.Set(&v)
}
// SetRecoverSanVolumeParamsNil sets the value for RecoverSanVolumeParams to be an explicit nil
func (o *RecoverPureParams) SetRecoverSanVolumeParamsNil() {
	o.RecoverSanVolumeParams.Set(nil)
}

// UnsetRecoverSanVolumeParams ensures that no value is present for RecoverSanVolumeParams, not even an explicit nil
func (o *RecoverPureParams) UnsetRecoverSanVolumeParams() {
	o.RecoverSanVolumeParams.Unset()
}

func (o RecoverPureParams) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Objects != nil {
		toSerialize["objects"] = o.Objects
	}
	if true {
		toSerialize["recoveryAction"] = o.RecoveryAction
	}
	if o.RecoverSanVolumeParams.IsSet() {
		toSerialize["recoverSanVolumeParams"] = o.RecoverSanVolumeParams.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableRecoverPureParams struct {
	value *RecoverPureParams
	isSet bool
}

func (v NullableRecoverPureParams) Get() *RecoverPureParams {
	return v.value
}

func (v *NullableRecoverPureParams) Set(val *RecoverPureParams) {
	v.value = val
	v.isSet = true
}

func (v NullableRecoverPureParams) IsSet() bool {
	return v.isSet
}

func (v *NullableRecoverPureParams) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRecoverPureParams(val *RecoverPureParams) *NullableRecoverPureParams {
	return &NullableRecoverPureParams{value: val, isSet: true}
}

func (v NullableRecoverPureParams) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRecoverPureParams) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


