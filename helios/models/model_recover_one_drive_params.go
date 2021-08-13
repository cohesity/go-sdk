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

// RecoverOneDriveParams Specifies the parameters to recover an Office 365 OneDrive.
type RecoverOneDriveParams struct {
	// Specifies a list of OneDrive params associated with the objects to recover.
	Objects []ObjectOneDriveParam `json:"objects"`
	// Specifies the target OneDrive to recover to. If not specified, the objects will be recovered to original location.
	TargetDrive *TargetOneDriveParam `json:"targetDrive,omitempty"`
	// Specifies whether to continue recovering other OneDrive items if one of items failed to recover. Default value is false.
	ContinueOnError NullableBool `json:"continueOnError,omitempty"`
}

// NewRecoverOneDriveParams instantiates a new RecoverOneDriveParams object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRecoverOneDriveParams(objects []ObjectOneDriveParam) *RecoverOneDriveParams {
	this := RecoverOneDriveParams{}
	this.Objects = objects
	return &this
}

// NewRecoverOneDriveParamsWithDefaults instantiates a new RecoverOneDriveParams object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRecoverOneDriveParamsWithDefaults() *RecoverOneDriveParams {
	this := RecoverOneDriveParams{}
	return &this
}

// GetObjects returns the Objects field value
// If the value is explicit nil, the zero value for []ObjectOneDriveParam will be returned
func (o *RecoverOneDriveParams) GetObjects() []ObjectOneDriveParam {
	if o == nil {
		var ret []ObjectOneDriveParam
		return ret
	}

	return o.Objects
}

// GetObjectsOk returns a tuple with the Objects field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RecoverOneDriveParams) GetObjectsOk() (*[]ObjectOneDriveParam, bool) {
	if o == nil || o.Objects == nil {
		return nil, false
	}
	return &o.Objects, true
}

// SetObjects sets field value
func (o *RecoverOneDriveParams) SetObjects(v []ObjectOneDriveParam) {
	o.Objects = v
}

// GetTargetDrive returns the TargetDrive field value if set, zero value otherwise.
func (o *RecoverOneDriveParams) GetTargetDrive() TargetOneDriveParam {
	if o == nil || o.TargetDrive == nil {
		var ret TargetOneDriveParam
		return ret
	}
	return *o.TargetDrive
}

// GetTargetDriveOk returns a tuple with the TargetDrive field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RecoverOneDriveParams) GetTargetDriveOk() (*TargetOneDriveParam, bool) {
	if o == nil || o.TargetDrive == nil {
		return nil, false
	}
	return o.TargetDrive, true
}

// HasTargetDrive returns a boolean if a field has been set.
func (o *RecoverOneDriveParams) HasTargetDrive() bool {
	if o != nil && o.TargetDrive != nil {
		return true
	}

	return false
}

// SetTargetDrive gets a reference to the given TargetOneDriveParam and assigns it to the TargetDrive field.
func (o *RecoverOneDriveParams) SetTargetDrive(v TargetOneDriveParam) {
	o.TargetDrive = &v
}

// GetContinueOnError returns the ContinueOnError field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RecoverOneDriveParams) GetContinueOnError() bool {
	if o == nil || o.ContinueOnError.Get() == nil {
		var ret bool
		return ret
	}
	return *o.ContinueOnError.Get()
}

// GetContinueOnErrorOk returns a tuple with the ContinueOnError field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RecoverOneDriveParams) GetContinueOnErrorOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return o.ContinueOnError.Get(), o.ContinueOnError.IsSet()
}

// HasContinueOnError returns a boolean if a field has been set.
func (o *RecoverOneDriveParams) HasContinueOnError() bool {
	if o != nil && o.ContinueOnError.IsSet() {
		return true
	}

	return false
}

// SetContinueOnError gets a reference to the given NullableBool and assigns it to the ContinueOnError field.
func (o *RecoverOneDriveParams) SetContinueOnError(v bool) {
	o.ContinueOnError.Set(&v)
}
// SetContinueOnErrorNil sets the value for ContinueOnError to be an explicit nil
func (o *RecoverOneDriveParams) SetContinueOnErrorNil() {
	o.ContinueOnError.Set(nil)
}

// UnsetContinueOnError ensures that no value is present for ContinueOnError, not even an explicit nil
func (o *RecoverOneDriveParams) UnsetContinueOnError() {
	o.ContinueOnError.Unset()
}

func (o RecoverOneDriveParams) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Objects != nil {
		toSerialize["objects"] = o.Objects
	}
	if o.TargetDrive != nil {
		toSerialize["targetDrive"] = o.TargetDrive
	}
	if o.ContinueOnError.IsSet() {
		toSerialize["continueOnError"] = o.ContinueOnError.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableRecoverOneDriveParams struct {
	value *RecoverOneDriveParams
	isSet bool
}

func (v NullableRecoverOneDriveParams) Get() *RecoverOneDriveParams {
	return v.value
}

func (v *NullableRecoverOneDriveParams) Set(val *RecoverOneDriveParams) {
	v.value = val
	v.isSet = true
}

func (v NullableRecoverOneDriveParams) IsSet() bool {
	return v.isSet
}

func (v *NullableRecoverOneDriveParams) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRecoverOneDriveParams(val *RecoverOneDriveParams) *NullableRecoverOneDriveParams {
	return &NullableRecoverOneDriveParams{value: val, isSet: true}
}

func (v NullableRecoverOneDriveParams) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRecoverOneDriveParams) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


