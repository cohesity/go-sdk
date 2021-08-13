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

// RecoverGenericNasParams Specifies the recovery options specific to Generic NAS environment.
type RecoverGenericNasParams struct {
	// Specifies the list of recover Object parameters.
	Objects []CommonRecoverObjectSnapshotParams `json:"objects"`
	// Specifies the type of recover action to be performed.
	RecoveryAction string `json:"recoveryAction"`
	// Specifies the parameters to recover Nas Volumes.
	RecoverNasVolumeParams NullableRecoverGenericNasNasVolumeParams `json:"recoverNasVolumeParams,omitempty"`
	// Specifies the parameters to recover files.
	RecoverFileAndFolderParams NullableRecoverGenericNasFilesParams `json:"recoverFileAndFolderParams,omitempty"`
	// Specifies the parameters to download files and folders.
	DownloadFileAndFolderParams NullableCommonDownloadFileAndFolderParams `json:"downloadFileAndFolderParams,omitempty"`
}

// NewRecoverGenericNasParams instantiates a new RecoverGenericNasParams object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRecoverGenericNasParams(objects []CommonRecoverObjectSnapshotParams, recoveryAction string) *RecoverGenericNasParams {
	this := RecoverGenericNasParams{}
	this.Objects = objects
	this.RecoveryAction = recoveryAction
	return &this
}

// NewRecoverGenericNasParamsWithDefaults instantiates a new RecoverGenericNasParams object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRecoverGenericNasParamsWithDefaults() *RecoverGenericNasParams {
	this := RecoverGenericNasParams{}
	return &this
}

// GetObjects returns the Objects field value
// If the value is explicit nil, the zero value for []CommonRecoverObjectSnapshotParams will be returned
func (o *RecoverGenericNasParams) GetObjects() []CommonRecoverObjectSnapshotParams {
	if o == nil {
		var ret []CommonRecoverObjectSnapshotParams
		return ret
	}

	return o.Objects
}

// GetObjectsOk returns a tuple with the Objects field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RecoverGenericNasParams) GetObjectsOk() (*[]CommonRecoverObjectSnapshotParams, bool) {
	if o == nil || o.Objects == nil {
		return nil, false
	}
	return &o.Objects, true
}

// SetObjects sets field value
func (o *RecoverGenericNasParams) SetObjects(v []CommonRecoverObjectSnapshotParams) {
	o.Objects = v
}

// GetRecoveryAction returns the RecoveryAction field value
func (o *RecoverGenericNasParams) GetRecoveryAction() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RecoveryAction
}

// GetRecoveryActionOk returns a tuple with the RecoveryAction field value
// and a boolean to check if the value has been set.
func (o *RecoverGenericNasParams) GetRecoveryActionOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.RecoveryAction, true
}

// SetRecoveryAction sets field value
func (o *RecoverGenericNasParams) SetRecoveryAction(v string) {
	o.RecoveryAction = v
}

// GetRecoverNasVolumeParams returns the RecoverNasVolumeParams field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RecoverGenericNasParams) GetRecoverNasVolumeParams() RecoverGenericNasNasVolumeParams {
	if o == nil || o.RecoverNasVolumeParams.Get() == nil {
		var ret RecoverGenericNasNasVolumeParams
		return ret
	}
	return *o.RecoverNasVolumeParams.Get()
}

// GetRecoverNasVolumeParamsOk returns a tuple with the RecoverNasVolumeParams field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RecoverGenericNasParams) GetRecoverNasVolumeParamsOk() (*RecoverGenericNasNasVolumeParams, bool) {
	if o == nil  {
		return nil, false
	}
	return o.RecoverNasVolumeParams.Get(), o.RecoverNasVolumeParams.IsSet()
}

// HasRecoverNasVolumeParams returns a boolean if a field has been set.
func (o *RecoverGenericNasParams) HasRecoverNasVolumeParams() bool {
	if o != nil && o.RecoverNasVolumeParams.IsSet() {
		return true
	}

	return false
}

// SetRecoverNasVolumeParams gets a reference to the given NullableRecoverGenericNasNasVolumeParams and assigns it to the RecoverNasVolumeParams field.
func (o *RecoverGenericNasParams) SetRecoverNasVolumeParams(v RecoverGenericNasNasVolumeParams) {
	o.RecoverNasVolumeParams.Set(&v)
}
// SetRecoverNasVolumeParamsNil sets the value for RecoverNasVolumeParams to be an explicit nil
func (o *RecoverGenericNasParams) SetRecoverNasVolumeParamsNil() {
	o.RecoverNasVolumeParams.Set(nil)
}

// UnsetRecoverNasVolumeParams ensures that no value is present for RecoverNasVolumeParams, not even an explicit nil
func (o *RecoverGenericNasParams) UnsetRecoverNasVolumeParams() {
	o.RecoverNasVolumeParams.Unset()
}

// GetRecoverFileAndFolderParams returns the RecoverFileAndFolderParams field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RecoverGenericNasParams) GetRecoverFileAndFolderParams() RecoverGenericNasFilesParams {
	if o == nil || o.RecoverFileAndFolderParams.Get() == nil {
		var ret RecoverGenericNasFilesParams
		return ret
	}
	return *o.RecoverFileAndFolderParams.Get()
}

// GetRecoverFileAndFolderParamsOk returns a tuple with the RecoverFileAndFolderParams field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RecoverGenericNasParams) GetRecoverFileAndFolderParamsOk() (*RecoverGenericNasFilesParams, bool) {
	if o == nil  {
		return nil, false
	}
	return o.RecoverFileAndFolderParams.Get(), o.RecoverFileAndFolderParams.IsSet()
}

// HasRecoverFileAndFolderParams returns a boolean if a field has been set.
func (o *RecoverGenericNasParams) HasRecoverFileAndFolderParams() bool {
	if o != nil && o.RecoverFileAndFolderParams.IsSet() {
		return true
	}

	return false
}

// SetRecoverFileAndFolderParams gets a reference to the given NullableRecoverGenericNasFilesParams and assigns it to the RecoverFileAndFolderParams field.
func (o *RecoverGenericNasParams) SetRecoverFileAndFolderParams(v RecoverGenericNasFilesParams) {
	o.RecoverFileAndFolderParams.Set(&v)
}
// SetRecoverFileAndFolderParamsNil sets the value for RecoverFileAndFolderParams to be an explicit nil
func (o *RecoverGenericNasParams) SetRecoverFileAndFolderParamsNil() {
	o.RecoverFileAndFolderParams.Set(nil)
}

// UnsetRecoverFileAndFolderParams ensures that no value is present for RecoverFileAndFolderParams, not even an explicit nil
func (o *RecoverGenericNasParams) UnsetRecoverFileAndFolderParams() {
	o.RecoverFileAndFolderParams.Unset()
}

// GetDownloadFileAndFolderParams returns the DownloadFileAndFolderParams field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RecoverGenericNasParams) GetDownloadFileAndFolderParams() CommonDownloadFileAndFolderParams {
	if o == nil || o.DownloadFileAndFolderParams.Get() == nil {
		var ret CommonDownloadFileAndFolderParams
		return ret
	}
	return *o.DownloadFileAndFolderParams.Get()
}

// GetDownloadFileAndFolderParamsOk returns a tuple with the DownloadFileAndFolderParams field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RecoverGenericNasParams) GetDownloadFileAndFolderParamsOk() (*CommonDownloadFileAndFolderParams, bool) {
	if o == nil  {
		return nil, false
	}
	return o.DownloadFileAndFolderParams.Get(), o.DownloadFileAndFolderParams.IsSet()
}

// HasDownloadFileAndFolderParams returns a boolean if a field has been set.
func (o *RecoverGenericNasParams) HasDownloadFileAndFolderParams() bool {
	if o != nil && o.DownloadFileAndFolderParams.IsSet() {
		return true
	}

	return false
}

// SetDownloadFileAndFolderParams gets a reference to the given NullableCommonDownloadFileAndFolderParams and assigns it to the DownloadFileAndFolderParams field.
func (o *RecoverGenericNasParams) SetDownloadFileAndFolderParams(v CommonDownloadFileAndFolderParams) {
	o.DownloadFileAndFolderParams.Set(&v)
}
// SetDownloadFileAndFolderParamsNil sets the value for DownloadFileAndFolderParams to be an explicit nil
func (o *RecoverGenericNasParams) SetDownloadFileAndFolderParamsNil() {
	o.DownloadFileAndFolderParams.Set(nil)
}

// UnsetDownloadFileAndFolderParams ensures that no value is present for DownloadFileAndFolderParams, not even an explicit nil
func (o *RecoverGenericNasParams) UnsetDownloadFileAndFolderParams() {
	o.DownloadFileAndFolderParams.Unset()
}

func (o RecoverGenericNasParams) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Objects != nil {
		toSerialize["objects"] = o.Objects
	}
	if true {
		toSerialize["recoveryAction"] = o.RecoveryAction
	}
	if o.RecoverNasVolumeParams.IsSet() {
		toSerialize["recoverNasVolumeParams"] = o.RecoverNasVolumeParams.Get()
	}
	if o.RecoverFileAndFolderParams.IsSet() {
		toSerialize["recoverFileAndFolderParams"] = o.RecoverFileAndFolderParams.Get()
	}
	if o.DownloadFileAndFolderParams.IsSet() {
		toSerialize["downloadFileAndFolderParams"] = o.DownloadFileAndFolderParams.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableRecoverGenericNasParams struct {
	value *RecoverGenericNasParams
	isSet bool
}

func (v NullableRecoverGenericNasParams) Get() *RecoverGenericNasParams {
	return v.value
}

func (v *NullableRecoverGenericNasParams) Set(val *RecoverGenericNasParams) {
	v.value = val
	v.isSet = true
}

func (v NullableRecoverGenericNasParams) IsSet() bool {
	return v.isSet
}

func (v *NullableRecoverGenericNasParams) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRecoverGenericNasParams(val *RecoverGenericNasParams) *NullableRecoverGenericNasParams {
	return &NullableRecoverGenericNasParams{value: val, isSet: true}
}

func (v NullableRecoverGenericNasParams) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRecoverGenericNasParams) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}




func (o RecoverGenericNasParams) Print() {
		if byteArray, err := o.MarshalJSON(); err != nil {
				fmt.Println(err)
		} else {
				fmt.Println(string(byteArray))
		}
}