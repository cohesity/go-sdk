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

// RecoverNetappNasVolumeParams Specifies the parameters to recover Netapp NAS volumes.
type RecoverNetappNasVolumeParams struct {
	// Specifies if the snapshot trying to recover is from a source initiated protection.
	IsFromSourceInitiatedProtection NullableBool `json:"isFromSourceInitiatedProtection,omitempty"`
	// Specifies the environment of the recovery target. The corresponding params below must be filled out.
	TargetEnvironment string `json:"targetEnvironment"`
	// Specifies the params for a Netapp recovery target.
	NetappTargetParams NullableRecoverNetappToNetappVolumeTargetParams `json:"netappTargetParams,omitempty"`
	// Specifies the params for an Elastifile recovery target.
	ElastifileTargetParams NullableRecoverOtherNasToElastifileVolumeTargetParams `json:"elastifileTargetParams,omitempty"`
	// Specifies the params for a Flashblade recovery target.
	FlashbladeTargetParams NullableRecoverOtherNasToFlashbladeVolumeTargetParams `json:"flashbladeTargetParams,omitempty"`
	// Specifies the params for a generic NAS recovery target.
	GenericNasTargetParams NullableRecoverOtherNasToGenericNasVolumeTargetParams `json:"genericNasTargetParams,omitempty"`
	// Specifies the params for a GPFS recovery target.
	GpfsTargetParams NullableRecoverOtherNasToGpfsVolumeTargetParams `json:"gpfsTargetParams,omitempty"`
	// Specifies the params for an Isilon recovery target.
	IsilonTargetParams NullableRecoverOtherNasToIsilonVolumeTargetParams `json:"isilonTargetParams,omitempty"`
	// Specifies the params for a Cohesity view recovery target.
	ViewTargetParams NullableRecoverNasVolumeToViewParams `json:"viewTargetParams,omitempty"`
}

// NewRecoverNetappNasVolumeParams instantiates a new RecoverNetappNasVolumeParams object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRecoverNetappNasVolumeParams(targetEnvironment string) *RecoverNetappNasVolumeParams {
	this := RecoverNetappNasVolumeParams{}
	this.TargetEnvironment = targetEnvironment
	return &this
}

// NewRecoverNetappNasVolumeParamsWithDefaults instantiates a new RecoverNetappNasVolumeParams object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRecoverNetappNasVolumeParamsWithDefaults() *RecoverNetappNasVolumeParams {
	this := RecoverNetappNasVolumeParams{}
	return &this
}

// GetIsFromSourceInitiatedProtection returns the IsFromSourceInitiatedProtection field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RecoverNetappNasVolumeParams) GetIsFromSourceInitiatedProtection() bool {
	if o == nil || o.IsFromSourceInitiatedProtection.Get() == nil {
		var ret bool
		return ret
	}
	return *o.IsFromSourceInitiatedProtection.Get()
}

// GetIsFromSourceInitiatedProtectionOk returns a tuple with the IsFromSourceInitiatedProtection field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RecoverNetappNasVolumeParams) GetIsFromSourceInitiatedProtectionOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return o.IsFromSourceInitiatedProtection.Get(), o.IsFromSourceInitiatedProtection.IsSet()
}

// HasIsFromSourceInitiatedProtection returns a boolean if a field has been set.
func (o *RecoverNetappNasVolumeParams) HasIsFromSourceInitiatedProtection() bool {
	if o != nil && o.IsFromSourceInitiatedProtection.IsSet() {
		return true
	}

	return false
}

// SetIsFromSourceInitiatedProtection gets a reference to the given NullableBool and assigns it to the IsFromSourceInitiatedProtection field.
func (o *RecoverNetappNasVolumeParams) SetIsFromSourceInitiatedProtection(v bool) {
	o.IsFromSourceInitiatedProtection.Set(&v)
}
// SetIsFromSourceInitiatedProtectionNil sets the value for IsFromSourceInitiatedProtection to be an explicit nil
func (o *RecoverNetappNasVolumeParams) SetIsFromSourceInitiatedProtectionNil() {
	o.IsFromSourceInitiatedProtection.Set(nil)
}

// UnsetIsFromSourceInitiatedProtection ensures that no value is present for IsFromSourceInitiatedProtection, not even an explicit nil
func (o *RecoverNetappNasVolumeParams) UnsetIsFromSourceInitiatedProtection() {
	o.IsFromSourceInitiatedProtection.Unset()
}

// GetTargetEnvironment returns the TargetEnvironment field value
func (o *RecoverNetappNasVolumeParams) GetTargetEnvironment() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TargetEnvironment
}

// GetTargetEnvironmentOk returns a tuple with the TargetEnvironment field value
// and a boolean to check if the value has been set.
func (o *RecoverNetappNasVolumeParams) GetTargetEnvironmentOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.TargetEnvironment, true
}

// SetTargetEnvironment sets field value
func (o *RecoverNetappNasVolumeParams) SetTargetEnvironment(v string) {
	o.TargetEnvironment = v
}

// GetNetappTargetParams returns the NetappTargetParams field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RecoverNetappNasVolumeParams) GetNetappTargetParams() RecoverNetappToNetappVolumeTargetParams {
	if o == nil || o.NetappTargetParams.Get() == nil {
		var ret RecoverNetappToNetappVolumeTargetParams
		return ret
	}
	return *o.NetappTargetParams.Get()
}

// GetNetappTargetParamsOk returns a tuple with the NetappTargetParams field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RecoverNetappNasVolumeParams) GetNetappTargetParamsOk() (*RecoverNetappToNetappVolumeTargetParams, bool) {
	if o == nil  {
		return nil, false
	}
	return o.NetappTargetParams.Get(), o.NetappTargetParams.IsSet()
}

// HasNetappTargetParams returns a boolean if a field has been set.
func (o *RecoverNetappNasVolumeParams) HasNetappTargetParams() bool {
	if o != nil && o.NetappTargetParams.IsSet() {
		return true
	}

	return false
}

// SetNetappTargetParams gets a reference to the given NullableRecoverNetappToNetappVolumeTargetParams and assigns it to the NetappTargetParams field.
func (o *RecoverNetappNasVolumeParams) SetNetappTargetParams(v RecoverNetappToNetappVolumeTargetParams) {
	o.NetappTargetParams.Set(&v)
}
// SetNetappTargetParamsNil sets the value for NetappTargetParams to be an explicit nil
func (o *RecoverNetappNasVolumeParams) SetNetappTargetParamsNil() {
	o.NetappTargetParams.Set(nil)
}

// UnsetNetappTargetParams ensures that no value is present for NetappTargetParams, not even an explicit nil
func (o *RecoverNetappNasVolumeParams) UnsetNetappTargetParams() {
	o.NetappTargetParams.Unset()
}

// GetElastifileTargetParams returns the ElastifileTargetParams field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RecoverNetappNasVolumeParams) GetElastifileTargetParams() RecoverOtherNasToElastifileVolumeTargetParams {
	if o == nil || o.ElastifileTargetParams.Get() == nil {
		var ret RecoverOtherNasToElastifileVolumeTargetParams
		return ret
	}
	return *o.ElastifileTargetParams.Get()
}

// GetElastifileTargetParamsOk returns a tuple with the ElastifileTargetParams field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RecoverNetappNasVolumeParams) GetElastifileTargetParamsOk() (*RecoverOtherNasToElastifileVolumeTargetParams, bool) {
	if o == nil  {
		return nil, false
	}
	return o.ElastifileTargetParams.Get(), o.ElastifileTargetParams.IsSet()
}

// HasElastifileTargetParams returns a boolean if a field has been set.
func (o *RecoverNetappNasVolumeParams) HasElastifileTargetParams() bool {
	if o != nil && o.ElastifileTargetParams.IsSet() {
		return true
	}

	return false
}

// SetElastifileTargetParams gets a reference to the given NullableRecoverOtherNasToElastifileVolumeTargetParams and assigns it to the ElastifileTargetParams field.
func (o *RecoverNetappNasVolumeParams) SetElastifileTargetParams(v RecoverOtherNasToElastifileVolumeTargetParams) {
	o.ElastifileTargetParams.Set(&v)
}
// SetElastifileTargetParamsNil sets the value for ElastifileTargetParams to be an explicit nil
func (o *RecoverNetappNasVolumeParams) SetElastifileTargetParamsNil() {
	o.ElastifileTargetParams.Set(nil)
}

// UnsetElastifileTargetParams ensures that no value is present for ElastifileTargetParams, not even an explicit nil
func (o *RecoverNetappNasVolumeParams) UnsetElastifileTargetParams() {
	o.ElastifileTargetParams.Unset()
}

// GetFlashbladeTargetParams returns the FlashbladeTargetParams field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RecoverNetappNasVolumeParams) GetFlashbladeTargetParams() RecoverOtherNasToFlashbladeVolumeTargetParams {
	if o == nil || o.FlashbladeTargetParams.Get() == nil {
		var ret RecoverOtherNasToFlashbladeVolumeTargetParams
		return ret
	}
	return *o.FlashbladeTargetParams.Get()
}

// GetFlashbladeTargetParamsOk returns a tuple with the FlashbladeTargetParams field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RecoverNetappNasVolumeParams) GetFlashbladeTargetParamsOk() (*RecoverOtherNasToFlashbladeVolumeTargetParams, bool) {
	if o == nil  {
		return nil, false
	}
	return o.FlashbladeTargetParams.Get(), o.FlashbladeTargetParams.IsSet()
}

// HasFlashbladeTargetParams returns a boolean if a field has been set.
func (o *RecoverNetappNasVolumeParams) HasFlashbladeTargetParams() bool {
	if o != nil && o.FlashbladeTargetParams.IsSet() {
		return true
	}

	return false
}

// SetFlashbladeTargetParams gets a reference to the given NullableRecoverOtherNasToFlashbladeVolumeTargetParams and assigns it to the FlashbladeTargetParams field.
func (o *RecoverNetappNasVolumeParams) SetFlashbladeTargetParams(v RecoverOtherNasToFlashbladeVolumeTargetParams) {
	o.FlashbladeTargetParams.Set(&v)
}
// SetFlashbladeTargetParamsNil sets the value for FlashbladeTargetParams to be an explicit nil
func (o *RecoverNetappNasVolumeParams) SetFlashbladeTargetParamsNil() {
	o.FlashbladeTargetParams.Set(nil)
}

// UnsetFlashbladeTargetParams ensures that no value is present for FlashbladeTargetParams, not even an explicit nil
func (o *RecoverNetappNasVolumeParams) UnsetFlashbladeTargetParams() {
	o.FlashbladeTargetParams.Unset()
}

// GetGenericNasTargetParams returns the GenericNasTargetParams field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RecoverNetappNasVolumeParams) GetGenericNasTargetParams() RecoverOtherNasToGenericNasVolumeTargetParams {
	if o == nil || o.GenericNasTargetParams.Get() == nil {
		var ret RecoverOtherNasToGenericNasVolumeTargetParams
		return ret
	}
	return *o.GenericNasTargetParams.Get()
}

// GetGenericNasTargetParamsOk returns a tuple with the GenericNasTargetParams field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RecoverNetappNasVolumeParams) GetGenericNasTargetParamsOk() (*RecoverOtherNasToGenericNasVolumeTargetParams, bool) {
	if o == nil  {
		return nil, false
	}
	return o.GenericNasTargetParams.Get(), o.GenericNasTargetParams.IsSet()
}

// HasGenericNasTargetParams returns a boolean if a field has been set.
func (o *RecoverNetappNasVolumeParams) HasGenericNasTargetParams() bool {
	if o != nil && o.GenericNasTargetParams.IsSet() {
		return true
	}

	return false
}

// SetGenericNasTargetParams gets a reference to the given NullableRecoverOtherNasToGenericNasVolumeTargetParams and assigns it to the GenericNasTargetParams field.
func (o *RecoverNetappNasVolumeParams) SetGenericNasTargetParams(v RecoverOtherNasToGenericNasVolumeTargetParams) {
	o.GenericNasTargetParams.Set(&v)
}
// SetGenericNasTargetParamsNil sets the value for GenericNasTargetParams to be an explicit nil
func (o *RecoverNetappNasVolumeParams) SetGenericNasTargetParamsNil() {
	o.GenericNasTargetParams.Set(nil)
}

// UnsetGenericNasTargetParams ensures that no value is present for GenericNasTargetParams, not even an explicit nil
func (o *RecoverNetappNasVolumeParams) UnsetGenericNasTargetParams() {
	o.GenericNasTargetParams.Unset()
}

// GetGpfsTargetParams returns the GpfsTargetParams field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RecoverNetappNasVolumeParams) GetGpfsTargetParams() RecoverOtherNasToGpfsVolumeTargetParams {
	if o == nil || o.GpfsTargetParams.Get() == nil {
		var ret RecoverOtherNasToGpfsVolumeTargetParams
		return ret
	}
	return *o.GpfsTargetParams.Get()
}

// GetGpfsTargetParamsOk returns a tuple with the GpfsTargetParams field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RecoverNetappNasVolumeParams) GetGpfsTargetParamsOk() (*RecoverOtherNasToGpfsVolumeTargetParams, bool) {
	if o == nil  {
		return nil, false
	}
	return o.GpfsTargetParams.Get(), o.GpfsTargetParams.IsSet()
}

// HasGpfsTargetParams returns a boolean if a field has been set.
func (o *RecoverNetappNasVolumeParams) HasGpfsTargetParams() bool {
	if o != nil && o.GpfsTargetParams.IsSet() {
		return true
	}

	return false
}

// SetGpfsTargetParams gets a reference to the given NullableRecoverOtherNasToGpfsVolumeTargetParams and assigns it to the GpfsTargetParams field.
func (o *RecoverNetappNasVolumeParams) SetGpfsTargetParams(v RecoverOtherNasToGpfsVolumeTargetParams) {
	o.GpfsTargetParams.Set(&v)
}
// SetGpfsTargetParamsNil sets the value for GpfsTargetParams to be an explicit nil
func (o *RecoverNetappNasVolumeParams) SetGpfsTargetParamsNil() {
	o.GpfsTargetParams.Set(nil)
}

// UnsetGpfsTargetParams ensures that no value is present for GpfsTargetParams, not even an explicit nil
func (o *RecoverNetappNasVolumeParams) UnsetGpfsTargetParams() {
	o.GpfsTargetParams.Unset()
}

// GetIsilonTargetParams returns the IsilonTargetParams field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RecoverNetappNasVolumeParams) GetIsilonTargetParams() RecoverOtherNasToIsilonVolumeTargetParams {
	if o == nil || o.IsilonTargetParams.Get() == nil {
		var ret RecoverOtherNasToIsilonVolumeTargetParams
		return ret
	}
	return *o.IsilonTargetParams.Get()
}

// GetIsilonTargetParamsOk returns a tuple with the IsilonTargetParams field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RecoverNetappNasVolumeParams) GetIsilonTargetParamsOk() (*RecoverOtherNasToIsilonVolumeTargetParams, bool) {
	if o == nil  {
		return nil, false
	}
	return o.IsilonTargetParams.Get(), o.IsilonTargetParams.IsSet()
}

// HasIsilonTargetParams returns a boolean if a field has been set.
func (o *RecoverNetappNasVolumeParams) HasIsilonTargetParams() bool {
	if o != nil && o.IsilonTargetParams.IsSet() {
		return true
	}

	return false
}

// SetIsilonTargetParams gets a reference to the given NullableRecoverOtherNasToIsilonVolumeTargetParams and assigns it to the IsilonTargetParams field.
func (o *RecoverNetappNasVolumeParams) SetIsilonTargetParams(v RecoverOtherNasToIsilonVolumeTargetParams) {
	o.IsilonTargetParams.Set(&v)
}
// SetIsilonTargetParamsNil sets the value for IsilonTargetParams to be an explicit nil
func (o *RecoverNetappNasVolumeParams) SetIsilonTargetParamsNil() {
	o.IsilonTargetParams.Set(nil)
}

// UnsetIsilonTargetParams ensures that no value is present for IsilonTargetParams, not even an explicit nil
func (o *RecoverNetappNasVolumeParams) UnsetIsilonTargetParams() {
	o.IsilonTargetParams.Unset()
}

// GetViewTargetParams returns the ViewTargetParams field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RecoverNetappNasVolumeParams) GetViewTargetParams() RecoverNasVolumeToViewParams {
	if o == nil || o.ViewTargetParams.Get() == nil {
		var ret RecoverNasVolumeToViewParams
		return ret
	}
	return *o.ViewTargetParams.Get()
}

// GetViewTargetParamsOk returns a tuple with the ViewTargetParams field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RecoverNetappNasVolumeParams) GetViewTargetParamsOk() (*RecoverNasVolumeToViewParams, bool) {
	if o == nil  {
		return nil, false
	}
	return o.ViewTargetParams.Get(), o.ViewTargetParams.IsSet()
}

// HasViewTargetParams returns a boolean if a field has been set.
func (o *RecoverNetappNasVolumeParams) HasViewTargetParams() bool {
	if o != nil && o.ViewTargetParams.IsSet() {
		return true
	}

	return false
}

// SetViewTargetParams gets a reference to the given NullableRecoverNasVolumeToViewParams and assigns it to the ViewTargetParams field.
func (o *RecoverNetappNasVolumeParams) SetViewTargetParams(v RecoverNasVolumeToViewParams) {
	o.ViewTargetParams.Set(&v)
}
// SetViewTargetParamsNil sets the value for ViewTargetParams to be an explicit nil
func (o *RecoverNetappNasVolumeParams) SetViewTargetParamsNil() {
	o.ViewTargetParams.Set(nil)
}

// UnsetViewTargetParams ensures that no value is present for ViewTargetParams, not even an explicit nil
func (o *RecoverNetappNasVolumeParams) UnsetViewTargetParams() {
	o.ViewTargetParams.Unset()
}

func (o RecoverNetappNasVolumeParams) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.IsFromSourceInitiatedProtection.IsSet() {
		toSerialize["isFromSourceInitiatedProtection"] = o.IsFromSourceInitiatedProtection.Get()
	}
	if true {
		toSerialize["targetEnvironment"] = o.TargetEnvironment
	}
	if o.NetappTargetParams.IsSet() {
		toSerialize["netappTargetParams"] = o.NetappTargetParams.Get()
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
	if o.IsilonTargetParams.IsSet() {
		toSerialize["isilonTargetParams"] = o.IsilonTargetParams.Get()
	}
	if o.ViewTargetParams.IsSet() {
		toSerialize["viewTargetParams"] = o.ViewTargetParams.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableRecoverNetappNasVolumeParams struct {
	value *RecoverNetappNasVolumeParams
	isSet bool
}

func (v NullableRecoverNetappNasVolumeParams) Get() *RecoverNetappNasVolumeParams {
	return v.value
}

func (v *NullableRecoverNetappNasVolumeParams) Set(val *RecoverNetappNasVolumeParams) {
	v.value = val
	v.isSet = true
}

func (v NullableRecoverNetappNasVolumeParams) IsSet() bool {
	return v.isSet
}

func (v *NullableRecoverNetappNasVolumeParams) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRecoverNetappNasVolumeParams(val *RecoverNetappNasVolumeParams) *NullableRecoverNetappNasVolumeParams {
	return &NullableRecoverNetappNasVolumeParams{value: val, isSet: true}
}

func (v NullableRecoverNetappNasVolumeParams) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRecoverNetappNasVolumeParams) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


