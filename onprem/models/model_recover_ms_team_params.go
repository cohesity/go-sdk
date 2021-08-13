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

// RecoverMsTeamParams Specifies the parameters to recover Microsoft 365 Teams.
type RecoverMsTeamParams struct {
	// Specifies a list of Microsoft 365 Teams params associated with objects to recover.
	Objects []ObjectMsTeamParam `json:"objects"`
	// This field is deprecated. Use targetTeamNickName and targetTeamFullName instead.
	TargetMsTeam *TargetMsTeamParam `json:"targetMsTeam,omitempty"`
	// Specifies target team nickname in case restoreToOriginal is false.
	TargetTeamNickName NullableString `json:"targetTeamNickName,omitempty"`
	// Specifies target team name in case restoreToOriginal is false. This will be ignored if restoring to alternate existing team (i.e. to a team the nickname of which is same as the one supplied by the end user).
	TargetTeamFullName NullableString `json:"targetTeamFullName,omitempty"`
	// Specifies whether to continue recovering other teams, if some of the teams fail to recover. Default value is false.
	ContinueOnError NullableBool `json:"continueOnError,omitempty"`
}

// NewRecoverMsTeamParams instantiates a new RecoverMsTeamParams object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRecoverMsTeamParams(objects []ObjectMsTeamParam) *RecoverMsTeamParams {
	this := RecoverMsTeamParams{}
	this.Objects = objects
	return &this
}

// NewRecoverMsTeamParamsWithDefaults instantiates a new RecoverMsTeamParams object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRecoverMsTeamParamsWithDefaults() *RecoverMsTeamParams {
	this := RecoverMsTeamParams{}
	return &this
}

// GetObjects returns the Objects field value
// If the value is explicit nil, the zero value for []ObjectMsTeamParam will be returned
func (o *RecoverMsTeamParams) GetObjects() []ObjectMsTeamParam {
	if o == nil {
		var ret []ObjectMsTeamParam
		return ret
	}

	return o.Objects
}

// GetObjectsOk returns a tuple with the Objects field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RecoverMsTeamParams) GetObjectsOk() (*[]ObjectMsTeamParam, bool) {
	if o == nil || o.Objects == nil {
		return nil, false
	}
	return &o.Objects, true
}

// SetObjects sets field value
func (o *RecoverMsTeamParams) SetObjects(v []ObjectMsTeamParam) {
	o.Objects = v
}

// GetTargetMsTeam returns the TargetMsTeam field value if set, zero value otherwise.
func (o *RecoverMsTeamParams) GetTargetMsTeam() TargetMsTeamParam {
	if o == nil || o.TargetMsTeam == nil {
		var ret TargetMsTeamParam
		return ret
	}
	return *o.TargetMsTeam
}

// GetTargetMsTeamOk returns a tuple with the TargetMsTeam field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RecoverMsTeamParams) GetTargetMsTeamOk() (*TargetMsTeamParam, bool) {
	if o == nil || o.TargetMsTeam == nil {
		return nil, false
	}
	return o.TargetMsTeam, true
}

// HasTargetMsTeam returns a boolean if a field has been set.
func (o *RecoverMsTeamParams) HasTargetMsTeam() bool {
	if o != nil && o.TargetMsTeam != nil {
		return true
	}

	return false
}

// SetTargetMsTeam gets a reference to the given TargetMsTeamParam and assigns it to the TargetMsTeam field.
func (o *RecoverMsTeamParams) SetTargetMsTeam(v TargetMsTeamParam) {
	o.TargetMsTeam = &v
}

// GetTargetTeamNickName returns the TargetTeamNickName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RecoverMsTeamParams) GetTargetTeamNickName() string {
	if o == nil || o.TargetTeamNickName.Get() == nil {
		var ret string
		return ret
	}
	return *o.TargetTeamNickName.Get()
}

// GetTargetTeamNickNameOk returns a tuple with the TargetTeamNickName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RecoverMsTeamParams) GetTargetTeamNickNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.TargetTeamNickName.Get(), o.TargetTeamNickName.IsSet()
}

// HasTargetTeamNickName returns a boolean if a field has been set.
func (o *RecoverMsTeamParams) HasTargetTeamNickName() bool {
	if o != nil && o.TargetTeamNickName.IsSet() {
		return true
	}

	return false
}

// SetTargetTeamNickName gets a reference to the given NullableString and assigns it to the TargetTeamNickName field.
func (o *RecoverMsTeamParams) SetTargetTeamNickName(v string) {
	o.TargetTeamNickName.Set(&v)
}
// SetTargetTeamNickNameNil sets the value for TargetTeamNickName to be an explicit nil
func (o *RecoverMsTeamParams) SetTargetTeamNickNameNil() {
	o.TargetTeamNickName.Set(nil)
}

// UnsetTargetTeamNickName ensures that no value is present for TargetTeamNickName, not even an explicit nil
func (o *RecoverMsTeamParams) UnsetTargetTeamNickName() {
	o.TargetTeamNickName.Unset()
}

// GetTargetTeamFullName returns the TargetTeamFullName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RecoverMsTeamParams) GetTargetTeamFullName() string {
	if o == nil || o.TargetTeamFullName.Get() == nil {
		var ret string
		return ret
	}
	return *o.TargetTeamFullName.Get()
}

// GetTargetTeamFullNameOk returns a tuple with the TargetTeamFullName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RecoverMsTeamParams) GetTargetTeamFullNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.TargetTeamFullName.Get(), o.TargetTeamFullName.IsSet()
}

// HasTargetTeamFullName returns a boolean if a field has been set.
func (o *RecoverMsTeamParams) HasTargetTeamFullName() bool {
	if o != nil && o.TargetTeamFullName.IsSet() {
		return true
	}

	return false
}

// SetTargetTeamFullName gets a reference to the given NullableString and assigns it to the TargetTeamFullName field.
func (o *RecoverMsTeamParams) SetTargetTeamFullName(v string) {
	o.TargetTeamFullName.Set(&v)
}
// SetTargetTeamFullNameNil sets the value for TargetTeamFullName to be an explicit nil
func (o *RecoverMsTeamParams) SetTargetTeamFullNameNil() {
	o.TargetTeamFullName.Set(nil)
}

// UnsetTargetTeamFullName ensures that no value is present for TargetTeamFullName, not even an explicit nil
func (o *RecoverMsTeamParams) UnsetTargetTeamFullName() {
	o.TargetTeamFullName.Unset()
}

// GetContinueOnError returns the ContinueOnError field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RecoverMsTeamParams) GetContinueOnError() bool {
	if o == nil || o.ContinueOnError.Get() == nil {
		var ret bool
		return ret
	}
	return *o.ContinueOnError.Get()
}

// GetContinueOnErrorOk returns a tuple with the ContinueOnError field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RecoverMsTeamParams) GetContinueOnErrorOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return o.ContinueOnError.Get(), o.ContinueOnError.IsSet()
}

// HasContinueOnError returns a boolean if a field has been set.
func (o *RecoverMsTeamParams) HasContinueOnError() bool {
	if o != nil && o.ContinueOnError.IsSet() {
		return true
	}

	return false
}

// SetContinueOnError gets a reference to the given NullableBool and assigns it to the ContinueOnError field.
func (o *RecoverMsTeamParams) SetContinueOnError(v bool) {
	o.ContinueOnError.Set(&v)
}
// SetContinueOnErrorNil sets the value for ContinueOnError to be an explicit nil
func (o *RecoverMsTeamParams) SetContinueOnErrorNil() {
	o.ContinueOnError.Set(nil)
}

// UnsetContinueOnError ensures that no value is present for ContinueOnError, not even an explicit nil
func (o *RecoverMsTeamParams) UnsetContinueOnError() {
	o.ContinueOnError.Unset()
}

func (o RecoverMsTeamParams) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Objects != nil {
		toSerialize["objects"] = o.Objects
	}
	if o.TargetMsTeam != nil {
		toSerialize["targetMsTeam"] = o.TargetMsTeam
	}
	if o.TargetTeamNickName.IsSet() {
		toSerialize["targetTeamNickName"] = o.TargetTeamNickName.Get()
	}
	if o.TargetTeamFullName.IsSet() {
		toSerialize["targetTeamFullName"] = o.TargetTeamFullName.Get()
	}
	if o.ContinueOnError.IsSet() {
		toSerialize["continueOnError"] = o.ContinueOnError.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableRecoverMsTeamParams struct {
	value *RecoverMsTeamParams
	isSet bool
}

func (v NullableRecoverMsTeamParams) Get() *RecoverMsTeamParams {
	return v.value
}

func (v *NullableRecoverMsTeamParams) Set(val *RecoverMsTeamParams) {
	v.value = val
	v.isSet = true
}

func (v NullableRecoverMsTeamParams) IsSet() bool {
	return v.isSet
}

func (v *NullableRecoverMsTeamParams) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRecoverMsTeamParams(val *RecoverMsTeamParams) *NullableRecoverMsTeamParams {
	return &NullableRecoverMsTeamParams{value: val, isSet: true}
}

func (v NullableRecoverMsTeamParams) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRecoverMsTeamParams) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}




func (o RecoverMsTeamParams) Print() {
		if byteArray, err := o.MarshalJSON(); err != nil {
				fmt.Println(err)
		} else {
				fmt.Println(string(byteArray))
		}
}