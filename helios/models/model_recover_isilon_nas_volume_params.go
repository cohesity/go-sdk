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

// RecoverIsilonNasVolumeParams Specifies the parameters to recover Isilon NAS volumes.
type RecoverIsilonNasVolumeParams struct {
	// Specifies the environment of the recovery target. The corresponding params below must be filled out.
	TargetEnvironment string `json:"targetEnvironment"`
	// Specifies the params for a Isilon recovery target.
	IsilonTargetParams NullableRecoverIsilonToIsilonVolumeTargetParams `json:"isilonTargetParams,omitempty"`
	// Specifies the params for an Elastifile recovery target.
	ElastifileTargetParams NullableRecoverOtherNasToElastifileVolumeTargetParams `json:"elastifileTargetParams,omitempty"`
	// Specifies the params for a Flashblade recovery target.
	FlashbladeTargetParams NullableRecoverOtherNasToFlashbladeVolumeTargetParams `json:"flashbladeTargetParams,omitempty"`
	// Specifies the params for a generic NAS recovery target.
	GenericNasTargetParams NullableRecoverOtherNasToGenericNasVolumeTargetParams `json:"genericNasTargetParams,omitempty"`
	// Specifies the params for a GPFS recovery target.
	GpfsTargetParams NullableRecoverOtherNasToGpfsVolumeTargetParams `json:"gpfsTargetParams,omitempty"`
	// Specifies the params for an Netapp recovery target.
	NetappTargetParams NullableRecoverOtherNasToNetappVolumeTargetParams `json:"netappTargetParams,omitempty"`
	// Specifies the params for a Cohesity view recovery target.
	ViewTargetParams NullableRecoverNasVolumeToViewParams `json:"viewTargetParams,omitempty"`
}

// NewRecoverIsilonNasVolumeParams instantiates a new RecoverIsilonNasVolumeParams object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRecoverIsilonNasVolumeParams(targetEnvironment string) *RecoverIsilonNasVolumeParams {
	this := RecoverIsilonNasVolumeParams{}
	this.TargetEnvironment = targetEnvironment
	return &this
}

// NewRecoverIsilonNasVolumeParamsWithDefaults instantiates a new RecoverIsilonNasVolumeParams object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRecoverIsilonNasVolumeParamsWithDefaults() *RecoverIsilonNasVolumeParams {
	this := RecoverIsilonNasVolumeParams{}
	return &this
}

// GetTargetEnvironment returns the TargetEnvironment field value
func (o *RecoverIsilonNasVolumeParams) GetTargetEnvironment() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TargetEnvironment
}

// GetTargetEnvironmentOk returns a tuple with the TargetEnvironment field value
// and a boolean to check if the value has been set.
func (o *RecoverIsilonNasVolumeParams) GetTargetEnvironmentOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.TargetEnvironment, true
}

// SetTargetEnvironment sets field value
func (o *RecoverIsilonNasVolumeParams) SetTargetEnvironment(v string) {
	o.TargetEnvironment = v
}

// GetIsilonTargetParams returns the IsilonTargetParams field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RecoverIsilonNasVolumeParams) GetIsilonTargetParams() RecoverIsilonToIsilonVolumeTargetParams {
	if o == nil || o.IsilonTargetParams.Get() == nil {
		var ret RecoverIsilonToIsilonVolumeTargetParams
		return ret
	}
	return *o.IsilonTargetParams.Get()
}

// GetIsilonTargetParamsOk returns a tuple with the IsilonTargetParams field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RecoverIsilonNasVolumeParams) GetIsilonTargetParamsOk() (*RecoverIsilonToIsilonVolumeTargetParams, bool) {
	if o == nil  {
		return nil, false
	}
	return o.IsilonTargetParams.Get(), o.IsilonTargetParams.IsSet()
}

// HasIsilonTargetParams returns a boolean if a field has been set.
func (o *RecoverIsilonNasVolumeParams) HasIsilonTargetParams() bool {
	if o != nil && o.IsilonTargetParams.IsSet() {
		return true
	}

	return false
}

// SetIsilonTargetParams gets a reference to the given NullableRecoverIsilonToIsilonVolumeTargetParams and assigns it to the IsilonTargetParams field.
func (o *RecoverIsilonNasVolumeParams) SetIsilonTargetParams(v RecoverIsilonToIsilonVolumeTargetParams) {
	o.IsilonTargetParams.Set(&v)
}
// SetIsilonTargetParamsNil sets the value for IsilonTargetParams to be an explicit nil
func (o *RecoverIsilonNasVolumeParams) SetIsilonTargetParamsNil() {
	o.IsilonTargetParams.Set(nil)
}

// UnsetIsilonTargetParams ensures that no value is present for IsilonTargetParams, not even an explicit nil
func (o *RecoverIsilonNasVolumeParams) UnsetIsilonTargetParams() {
	o.IsilonTargetParams.Unset()
}

// GetElastifileTargetParams returns the ElastifileTargetParams field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RecoverIsilonNasVolumeParams) GetElastifileTargetParams() RecoverOtherNasToElastifileVolumeTargetParams {
	if o == nil || o.ElastifileTargetParams.Get() == nil {
		var ret RecoverOtherNasToElastifileVolumeTargetParams
		return ret
	}
	return *o.ElastifileTargetParams.Get()
}

// GetElastifileTargetParamsOk returns a tuple with the ElastifileTargetParams field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RecoverIsilonNasVolumeParams) GetElastifileTargetParamsOk() (*RecoverOtherNasToElastifileVolumeTargetParams, bool) {
	if o == nil  {
		return nil, false
	}
	return o.ElastifileTargetParams.Get(), o.ElastifileTargetParams.IsSet()
}

// HasElastifileTargetParams returns a boolean if a field has been set.
func (o *RecoverIsilonNasVolumeParams) HasElastifileTargetParams() bool {
	if o != nil && o.ElastifileTargetParams.IsSet() {
		return true
	}

	return false
}

// SetElastifileTargetParams gets a reference to the given NullableRecoverOtherNasToElastifileVolumeTargetParams and assigns it to the ElastifileTargetParams field.
func (o *RecoverIsilonNasVolumeParams) SetElastifileTargetParams(v RecoverOtherNasToElastifileVolumeTargetParams) {
	o.ElastifileTargetParams.Set(&v)
}
// SetElastifileTargetParamsNil sets the value for ElastifileTargetParams to be an explicit nil
func (o *RecoverIsilonNasVolumeParams) SetElastifileTargetParamsNil() {
	o.ElastifileTargetParams.Set(nil)
}

// UnsetElastifileTargetParams ensures that no value is present for ElastifileTargetParams, not even an explicit nil
func (o *RecoverIsilonNasVolumeParams) UnsetElastifileTargetParams() {
	o.ElastifileTargetParams.Unset()
}

// GetFlashbladeTargetParams returns the FlashbladeTargetParams field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RecoverIsilonNasVolumeParams) GetFlashbladeTargetParams() RecoverOtherNasToFlashbladeVolumeTargetParams {
	if o == nil || o.FlashbladeTargetParams.Get() == nil {
		var ret RecoverOtherNasToFlashbladeVolumeTargetParams
		return ret
	}
	return *o.FlashbladeTargetParams.Get()
}

// GetFlashbladeTargetParamsOk returns a tuple with the FlashbladeTargetParams field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RecoverIsilonNasVolumeParams) GetFlashbladeTargetParamsOk() (*RecoverOtherNasToFlashbladeVolumeTargetParams, bool) {
	if o == nil  {
		return nil, false
	}
	return o.FlashbladeTargetParams.Get(), o.FlashbladeTargetParams.IsSet()
}

// HasFlashbladeTargetParams returns a boolean if a field has been set.
func (o *RecoverIsilonNasVolumeParams) HasFlashbladeTargetParams() bool {
	if o != nil && o.FlashbladeTargetParams.IsSet() {
		return true
	}

	return false
}

// SetFlashbladeTargetParams gets a reference to the given NullableRecoverOtherNasToFlashbladeVolumeTargetParams and assigns it to the FlashbladeTargetParams field.
func (o *RecoverIsilonNasVolumeParams) SetFlashbladeTargetParams(v RecoverOtherNasToFlashbladeVolumeTargetParams) {
	o.FlashbladeTargetParams.Set(&v)
}
// SetFlashbladeTargetParamsNil sets the value for FlashbladeTargetParams to be an explicit nil
func (o *RecoverIsilonNasVolumeParams) SetFlashbladeTargetParamsNil() {
	o.FlashbladeTargetParams.Set(nil)
}

// UnsetFlashbladeTargetParams ensures that no value is present for FlashbladeTargetParams, not even an explicit nil
func (o *RecoverIsilonNasVolumeParams) UnsetFlashbladeTargetParams() {
	o.FlashbladeTargetParams.Unset()
}

// GetGenericNasTargetParams returns the GenericNasTargetParams field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RecoverIsilonNasVolumeParams) GetGenericNasTargetParams() RecoverOtherNasToGenericNasVolumeTargetParams {
	if o == nil || o.GenericNasTargetParams.Get() == nil {
		var ret RecoverOtherNasToGenericNasVolumeTargetParams
		return ret
	}
	return *o.GenericNasTargetParams.Get()
}

// GetGenericNasTargetParamsOk returns a tuple with the GenericNasTargetParams field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RecoverIsilonNasVolumeParams) GetGenericNasTargetParamsOk() (*RecoverOtherNasToGenericNasVolumeTargetParams, bool) {
	if o == nil  {
		return nil, false
	}
	return o.GenericNasTargetParams.Get(), o.GenericNasTargetParams.IsSet()
}

// HasGenericNasTargetParams returns a boolean if a field has been set.
func (o *RecoverIsilonNasVolumeParams) HasGenericNasTargetParams() bool {
	if o != nil && o.GenericNasTargetParams.IsSet() {
		return true
	}

	return false
}

// SetGenericNasTargetParams gets a reference to the given NullableRecoverOtherNasToGenericNasVolumeTargetParams and assigns it to the GenericNasTargetParams field.
func (o *RecoverIsilonNasVolumeParams) SetGenericNasTargetParams(v RecoverOtherNasToGenericNasVolumeTargetParams) {
	o.GenericNasTargetParams.Set(&v)
}
// SetGenericNasTargetParamsNil sets the value for GenericNasTargetParams to be an explicit nil
func (o *RecoverIsilonNasVolumeParams) SetGenericNasTargetParamsNil() {
	o.GenericNasTargetParams.Set(nil)
}

// UnsetGenericNasTargetParams ensures that no value is present for GenericNasTargetParams, not even an explicit nil
func (o *RecoverIsilonNasVolumeParams) UnsetGenericNasTargetParams() {
	o.GenericNasTargetParams.Unset()
}

// GetGpfsTargetParams returns the GpfsTargetParams field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RecoverIsilonNasVolumeParams) GetGpfsTargetParams() RecoverOtherNasToGpfsVolumeTargetParams {
	if o == nil || o.GpfsTargetParams.Get() == nil {
		var ret RecoverOtherNasToGpfsVolumeTargetParams
		return ret
	}
	return *o.GpfsTargetParams.Get()
}

// GetGpfsTargetParamsOk returns a tuple with the GpfsTargetParams field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RecoverIsilonNasVolumeParams) GetGpfsTargetParamsOk() (*RecoverOtherNasToGpfsVolumeTargetParams, bool) {
	if o == nil  {
		return nil, false
	}
	return o.GpfsTargetParams.Get(), o.GpfsTargetParams.IsSet()
}

// HasGpfsTargetParams returns a boolean if a field has been set.
func (o *RecoverIsilonNasVolumeParams) HasGpfsTargetParams() bool {
	if o != nil && o.GpfsTargetParams.IsSet() {
		return true
	}

	return false
}

// SetGpfsTargetParams gets a reference to the given NullableRecoverOtherNasToGpfsVolumeTargetParams and assigns it to the GpfsTargetParams field.
func (o *RecoverIsilonNasVolumeParams) SetGpfsTargetParams(v RecoverOtherNasToGpfsVolumeTargetParams) {
	o.GpfsTargetParams.Set(&v)
}
// SetGpfsTargetParamsNil sets the value for GpfsTargetParams to be an explicit nil
func (o *RecoverIsilonNasVolumeParams) SetGpfsTargetParamsNil() {
	o.GpfsTargetParams.Set(nil)
}

// UnsetGpfsTargetParams ensures that no value is present for GpfsTargetParams, not even an explicit nil
func (o *RecoverIsilonNasVolumeParams) UnsetGpfsTargetParams() {
	o.GpfsTargetParams.Unset()
}

// GetNetappTargetParams returns the NetappTargetParams field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RecoverIsilonNasVolumeParams) GetNetappTargetParams() RecoverOtherNasToNetappVolumeTargetParams {
	if o == nil || o.NetappTargetParams.Get() == nil {
		var ret RecoverOtherNasToNetappVolumeTargetParams
		return ret
	}
	return *o.NetappTargetParams.Get()
}

// GetNetappTargetParamsOk returns a tuple with the NetappTargetParams field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RecoverIsilonNasVolumeParams) GetNetappTargetParamsOk() (*RecoverOtherNasToNetappVolumeTargetParams, bool) {
	if o == nil  {
		return nil, false
	}
	return o.NetappTargetParams.Get(), o.NetappTargetParams.IsSet()
}

// HasNetappTargetParams returns a boolean if a field has been set.
func (o *RecoverIsilonNasVolumeParams) HasNetappTargetParams() bool {
	if o != nil && o.NetappTargetParams.IsSet() {
		return true
	}

	return false
}

// SetNetappTargetParams gets a reference to the given NullableRecoverOtherNasToNetappVolumeTargetParams and assigns it to the NetappTargetParams field.
func (o *RecoverIsilonNasVolumeParams) SetNetappTargetParams(v RecoverOtherNasToNetappVolumeTargetParams) {
	o.NetappTargetParams.Set(&v)
}
// SetNetappTargetParamsNil sets the value for NetappTargetParams to be an explicit nil
func (o *RecoverIsilonNasVolumeParams) SetNetappTargetParamsNil() {
	o.NetappTargetParams.Set(nil)
}

// UnsetNetappTargetParams ensures that no value is present for NetappTargetParams, not even an explicit nil
func (o *RecoverIsilonNasVolumeParams) UnsetNetappTargetParams() {
	o.NetappTargetParams.Unset()
}

// GetViewTargetParams returns the ViewTargetParams field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RecoverIsilonNasVolumeParams) GetViewTargetParams() RecoverNasVolumeToViewParams {
	if o == nil || o.ViewTargetParams.Get() == nil {
		var ret RecoverNasVolumeToViewParams
		return ret
	}
	return *o.ViewTargetParams.Get()
}

// GetViewTargetParamsOk returns a tuple with the ViewTargetParams field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RecoverIsilonNasVolumeParams) GetViewTargetParamsOk() (*RecoverNasVolumeToViewParams, bool) {
	if o == nil  {
		return nil, false
	}
	return o.ViewTargetParams.Get(), o.ViewTargetParams.IsSet()
}

// HasViewTargetParams returns a boolean if a field has been set.
func (o *RecoverIsilonNasVolumeParams) HasViewTargetParams() bool {
	if o != nil && o.ViewTargetParams.IsSet() {
		return true
	}

	return false
}

// SetViewTargetParams gets a reference to the given NullableRecoverNasVolumeToViewParams and assigns it to the ViewTargetParams field.
func (o *RecoverIsilonNasVolumeParams) SetViewTargetParams(v RecoverNasVolumeToViewParams) {
	o.ViewTargetParams.Set(&v)
}
// SetViewTargetParamsNil sets the value for ViewTargetParams to be an explicit nil
func (o *RecoverIsilonNasVolumeParams) SetViewTargetParamsNil() {
	o.ViewTargetParams.Set(nil)
}

// UnsetViewTargetParams ensures that no value is present for ViewTargetParams, not even an explicit nil
func (o *RecoverIsilonNasVolumeParams) UnsetViewTargetParams() {
	o.ViewTargetParams.Unset()
}

func (o RecoverIsilonNasVolumeParams) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["targetEnvironment"] = o.TargetEnvironment
	}
	if o.IsilonTargetParams.IsSet() {
		toSerialize["isilonTargetParams"] = o.IsilonTargetParams.Get()
	}
	if o.ElastifileTargetParams.IsSet() {
		toSerialize["elastifileTargetParams"] = o.ElastifileTargetParams.Get()
	}
	if o.FlashbladeTargetParams.IsSet() {
		toSerialize["flashbladeTargetParams"] = o.FlashbladeTargetParams.Get()
	}
	if o.GenericNasTargetParams.IsSet() {
		toSerialize["genericNasTargetParams"] = o.GenericNasTargetParams.Get()
	}
	if o.GpfsTargetParams.IsSet() {
		toSerialize["gpfsTargetParams"] = o.GpfsTargetParams.Get()
	}
	if o.NetappTargetParams.IsSet() {
		toSerialize["netappTargetParams"] = o.NetappTargetParams.Get()
	}
	if o.ViewTargetParams.IsSet() {
		toSerialize["viewTargetParams"] = o.ViewTargetParams.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableRecoverIsilonNasVolumeParams struct {
	value *RecoverIsilonNasVolumeParams
	isSet bool
}

func (v NullableRecoverIsilonNasVolumeParams) Get() *RecoverIsilonNasVolumeParams {
	return v.value
}

func (v *NullableRecoverIsilonNasVolumeParams) Set(val *RecoverIsilonNasVolumeParams) {
	v.value = val
	v.isSet = true
}

func (v NullableRecoverIsilonNasVolumeParams) IsSet() bool {
	return v.isSet
}

func (v *NullableRecoverIsilonNasVolumeParams) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRecoverIsilonNasVolumeParams(val *RecoverIsilonNasVolumeParams) *NullableRecoverIsilonNasVolumeParams {
	return &NullableRecoverIsilonNasVolumeParams{value: val, isSet: true}
}

func (v NullableRecoverIsilonNasVolumeParams) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRecoverIsilonNasVolumeParams) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


