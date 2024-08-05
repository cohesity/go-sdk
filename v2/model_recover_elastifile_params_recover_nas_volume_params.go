/*
Cohesity REST API

Cohesity API provides a RESTful interface to access the various data management operations on Cohesity cluster and Helios.

API version: 2.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package v2

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the RecoverElastifileParamsRecoverNasVolumeParams type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RecoverElastifileParamsRecoverNasVolumeParams{}

// RecoverElastifileParamsRecoverNasVolumeParams Specifies the parameters to recover Nas Volumes.
type RecoverElastifileParamsRecoverNasVolumeParams struct {
	ElastifileTargetParams NullableRecoverElastifileNasVolumeParamsElastifileTargetParams `json:"elastifileTargetParams,omitempty"`
	FlashbladeTargetParams NullableRecoverElastifileNasVolumeParamsFlashbladeTargetParams `json:"flashbladeTargetParams,omitempty"`
	GenericNasTargetParams NullableRecoverElastifileNasVolumeParamsGenericNasTargetParams `json:"genericNasTargetParams,omitempty"`
	GpfsTargetParams NullableRecoverElastifileNasVolumeParamsGpfsTargetParams `json:"gpfsTargetParams,omitempty"`
	IsilonTargetParams NullableRecoverElastifileNasVolumeParamsIsilonTargetParams `json:"isilonTargetParams,omitempty"`
	NetappTargetParams NullableRecoverElastifileNasVolumeParamsNetappTargetParams `json:"netappTargetParams,omitempty"`
	// Specifies the environment of the recovery target. The corresponding params below must be filled out.
	TargetEnvironment string `json:"targetEnvironment"`
	ViewTargetParams NullableRecoverElastifileNasVolumeParamsViewTargetParams `json:"viewTargetParams,omitempty"`
}

type _RecoverElastifileParamsRecoverNasVolumeParams RecoverElastifileParamsRecoverNasVolumeParams

// NewRecoverElastifileParamsRecoverNasVolumeParams instantiates a new RecoverElastifileParamsRecoverNasVolumeParams object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRecoverElastifileParamsRecoverNasVolumeParams(targetEnvironment string) *RecoverElastifileParamsRecoverNasVolumeParams {
	this := RecoverElastifileParamsRecoverNasVolumeParams{}
	this.TargetEnvironment = targetEnvironment
	return &this
}

// NewRecoverElastifileParamsRecoverNasVolumeParamsWithDefaults instantiates a new RecoverElastifileParamsRecoverNasVolumeParams object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRecoverElastifileParamsRecoverNasVolumeParamsWithDefaults() *RecoverElastifileParamsRecoverNasVolumeParams {
	this := RecoverElastifileParamsRecoverNasVolumeParams{}
	return &this
}

// GetElastifileTargetParams returns the ElastifileTargetParams field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RecoverElastifileParamsRecoverNasVolumeParams) GetElastifileTargetParams() RecoverElastifileNasVolumeParamsElastifileTargetParams {
	if o == nil || IsNil(o.ElastifileTargetParams.Get()) {
		var ret RecoverElastifileNasVolumeParamsElastifileTargetParams
		return ret
	}
	return *o.ElastifileTargetParams.Get()
}

// GetElastifileTargetParamsOk returns a tuple with the ElastifileTargetParams field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RecoverElastifileParamsRecoverNasVolumeParams) GetElastifileTargetParamsOk() (*RecoverElastifileNasVolumeParamsElastifileTargetParams, bool) {
	if o == nil {
		return nil, false
	}
	return o.ElastifileTargetParams.Get(), o.ElastifileTargetParams.IsSet()
}

// HasElastifileTargetParams returns a boolean if a field has been set.
func (o *RecoverElastifileParamsRecoverNasVolumeParams) HasElastifileTargetParams() bool {
	if o != nil && o.ElastifileTargetParams.IsSet() {
		return true
	}

	return false
}

// SetElastifileTargetParams gets a reference to the given NullableRecoverElastifileNasVolumeParamsElastifileTargetParams and assigns it to the ElastifileTargetParams field.
func (o *RecoverElastifileParamsRecoverNasVolumeParams) SetElastifileTargetParams(v RecoverElastifileNasVolumeParamsElastifileTargetParams) {
	o.ElastifileTargetParams.Set(&v)
}
// SetElastifileTargetParamsNil sets the value for ElastifileTargetParams to be an explicit nil
func (o *RecoverElastifileParamsRecoverNasVolumeParams) SetElastifileTargetParamsNil() {
	o.ElastifileTargetParams.Set(nil)
}

// UnsetElastifileTargetParams ensures that no value is present for ElastifileTargetParams, not even an explicit nil
func (o *RecoverElastifileParamsRecoverNasVolumeParams) UnsetElastifileTargetParams() {
	o.ElastifileTargetParams.Unset()
}

// GetFlashbladeTargetParams returns the FlashbladeTargetParams field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RecoverElastifileParamsRecoverNasVolumeParams) GetFlashbladeTargetParams() RecoverElastifileNasVolumeParamsFlashbladeTargetParams {
	if o == nil || IsNil(o.FlashbladeTargetParams.Get()) {
		var ret RecoverElastifileNasVolumeParamsFlashbladeTargetParams
		return ret
	}
	return *o.FlashbladeTargetParams.Get()
}

// GetFlashbladeTargetParamsOk returns a tuple with the FlashbladeTargetParams field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RecoverElastifileParamsRecoverNasVolumeParams) GetFlashbladeTargetParamsOk() (*RecoverElastifileNasVolumeParamsFlashbladeTargetParams, bool) {
	if o == nil {
		return nil, false
	}
	return o.FlashbladeTargetParams.Get(), o.FlashbladeTargetParams.IsSet()
}

// HasFlashbladeTargetParams returns a boolean if a field has been set.
func (o *RecoverElastifileParamsRecoverNasVolumeParams) HasFlashbladeTargetParams() bool {
	if o != nil && o.FlashbladeTargetParams.IsSet() {
		return true
	}

	return false
}

// SetFlashbladeTargetParams gets a reference to the given NullableRecoverElastifileNasVolumeParamsFlashbladeTargetParams and assigns it to the FlashbladeTargetParams field.
func (o *RecoverElastifileParamsRecoverNasVolumeParams) SetFlashbladeTargetParams(v RecoverElastifileNasVolumeParamsFlashbladeTargetParams) {
	o.FlashbladeTargetParams.Set(&v)
}
// SetFlashbladeTargetParamsNil sets the value for FlashbladeTargetParams to be an explicit nil
func (o *RecoverElastifileParamsRecoverNasVolumeParams) SetFlashbladeTargetParamsNil() {
	o.FlashbladeTargetParams.Set(nil)
}

// UnsetFlashbladeTargetParams ensures that no value is present for FlashbladeTargetParams, not even an explicit nil
func (o *RecoverElastifileParamsRecoverNasVolumeParams) UnsetFlashbladeTargetParams() {
	o.FlashbladeTargetParams.Unset()
}

// GetGenericNasTargetParams returns the GenericNasTargetParams field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RecoverElastifileParamsRecoverNasVolumeParams) GetGenericNasTargetParams() RecoverElastifileNasVolumeParamsGenericNasTargetParams {
	if o == nil || IsNil(o.GenericNasTargetParams.Get()) {
		var ret RecoverElastifileNasVolumeParamsGenericNasTargetParams
		return ret
	}
	return *o.GenericNasTargetParams.Get()
}

// GetGenericNasTargetParamsOk returns a tuple with the GenericNasTargetParams field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RecoverElastifileParamsRecoverNasVolumeParams) GetGenericNasTargetParamsOk() (*RecoverElastifileNasVolumeParamsGenericNasTargetParams, bool) {
	if o == nil {
		return nil, false
	}
	return o.GenericNasTargetParams.Get(), o.GenericNasTargetParams.IsSet()
}

// HasGenericNasTargetParams returns a boolean if a field has been set.
func (o *RecoverElastifileParamsRecoverNasVolumeParams) HasGenericNasTargetParams() bool {
	if o != nil && o.GenericNasTargetParams.IsSet() {
		return true
	}

	return false
}

// SetGenericNasTargetParams gets a reference to the given NullableRecoverElastifileNasVolumeParamsGenericNasTargetParams and assigns it to the GenericNasTargetParams field.
func (o *RecoverElastifileParamsRecoverNasVolumeParams) SetGenericNasTargetParams(v RecoverElastifileNasVolumeParamsGenericNasTargetParams) {
	o.GenericNasTargetParams.Set(&v)
}
// SetGenericNasTargetParamsNil sets the value for GenericNasTargetParams to be an explicit nil
func (o *RecoverElastifileParamsRecoverNasVolumeParams) SetGenericNasTargetParamsNil() {
	o.GenericNasTargetParams.Set(nil)
}

// UnsetGenericNasTargetParams ensures that no value is present for GenericNasTargetParams, not even an explicit nil
func (o *RecoverElastifileParamsRecoverNasVolumeParams) UnsetGenericNasTargetParams() {
	o.GenericNasTargetParams.Unset()
}

// GetGpfsTargetParams returns the GpfsTargetParams field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RecoverElastifileParamsRecoverNasVolumeParams) GetGpfsTargetParams() RecoverElastifileNasVolumeParamsGpfsTargetParams {
	if o == nil || IsNil(o.GpfsTargetParams.Get()) {
		var ret RecoverElastifileNasVolumeParamsGpfsTargetParams
		return ret
	}
	return *o.GpfsTargetParams.Get()
}

// GetGpfsTargetParamsOk returns a tuple with the GpfsTargetParams field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RecoverElastifileParamsRecoverNasVolumeParams) GetGpfsTargetParamsOk() (*RecoverElastifileNasVolumeParamsGpfsTargetParams, bool) {
	if o == nil {
		return nil, false
	}
	return o.GpfsTargetParams.Get(), o.GpfsTargetParams.IsSet()
}

// HasGpfsTargetParams returns a boolean if a field has been set.
func (o *RecoverElastifileParamsRecoverNasVolumeParams) HasGpfsTargetParams() bool {
	if o != nil && o.GpfsTargetParams.IsSet() {
		return true
	}

	return false
}

// SetGpfsTargetParams gets a reference to the given NullableRecoverElastifileNasVolumeParamsGpfsTargetParams and assigns it to the GpfsTargetParams field.
func (o *RecoverElastifileParamsRecoverNasVolumeParams) SetGpfsTargetParams(v RecoverElastifileNasVolumeParamsGpfsTargetParams) {
	o.GpfsTargetParams.Set(&v)
}
// SetGpfsTargetParamsNil sets the value for GpfsTargetParams to be an explicit nil
func (o *RecoverElastifileParamsRecoverNasVolumeParams) SetGpfsTargetParamsNil() {
	o.GpfsTargetParams.Set(nil)
}

// UnsetGpfsTargetParams ensures that no value is present for GpfsTargetParams, not even an explicit nil
func (o *RecoverElastifileParamsRecoverNasVolumeParams) UnsetGpfsTargetParams() {
	o.GpfsTargetParams.Unset()
}

// GetIsilonTargetParams returns the IsilonTargetParams field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RecoverElastifileParamsRecoverNasVolumeParams) GetIsilonTargetParams() RecoverElastifileNasVolumeParamsIsilonTargetParams {
	if o == nil || IsNil(o.IsilonTargetParams.Get()) {
		var ret RecoverElastifileNasVolumeParamsIsilonTargetParams
		return ret
	}
	return *o.IsilonTargetParams.Get()
}

// GetIsilonTargetParamsOk returns a tuple with the IsilonTargetParams field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RecoverElastifileParamsRecoverNasVolumeParams) GetIsilonTargetParamsOk() (*RecoverElastifileNasVolumeParamsIsilonTargetParams, bool) {
	if o == nil {
		return nil, false
	}
	return o.IsilonTargetParams.Get(), o.IsilonTargetParams.IsSet()
}

// HasIsilonTargetParams returns a boolean if a field has been set.
func (o *RecoverElastifileParamsRecoverNasVolumeParams) HasIsilonTargetParams() bool {
	if o != nil && o.IsilonTargetParams.IsSet() {
		return true
	}

	return false
}

// SetIsilonTargetParams gets a reference to the given NullableRecoverElastifileNasVolumeParamsIsilonTargetParams and assigns it to the IsilonTargetParams field.
func (o *RecoverElastifileParamsRecoverNasVolumeParams) SetIsilonTargetParams(v RecoverElastifileNasVolumeParamsIsilonTargetParams) {
	o.IsilonTargetParams.Set(&v)
}
// SetIsilonTargetParamsNil sets the value for IsilonTargetParams to be an explicit nil
func (o *RecoverElastifileParamsRecoverNasVolumeParams) SetIsilonTargetParamsNil() {
	o.IsilonTargetParams.Set(nil)
}

// UnsetIsilonTargetParams ensures that no value is present for IsilonTargetParams, not even an explicit nil
func (o *RecoverElastifileParamsRecoverNasVolumeParams) UnsetIsilonTargetParams() {
	o.IsilonTargetParams.Unset()
}

// GetNetappTargetParams returns the NetappTargetParams field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RecoverElastifileParamsRecoverNasVolumeParams) GetNetappTargetParams() RecoverElastifileNasVolumeParamsNetappTargetParams {
	if o == nil || IsNil(o.NetappTargetParams.Get()) {
		var ret RecoverElastifileNasVolumeParamsNetappTargetParams
		return ret
	}
	return *o.NetappTargetParams.Get()
}

// GetNetappTargetParamsOk returns a tuple with the NetappTargetParams field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RecoverElastifileParamsRecoverNasVolumeParams) GetNetappTargetParamsOk() (*RecoverElastifileNasVolumeParamsNetappTargetParams, bool) {
	if o == nil {
		return nil, false
	}
	return o.NetappTargetParams.Get(), o.NetappTargetParams.IsSet()
}

// HasNetappTargetParams returns a boolean if a field has been set.
func (o *RecoverElastifileParamsRecoverNasVolumeParams) HasNetappTargetParams() bool {
	if o != nil && o.NetappTargetParams.IsSet() {
		return true
	}

	return false
}

// SetNetappTargetParams gets a reference to the given NullableRecoverElastifileNasVolumeParamsNetappTargetParams and assigns it to the NetappTargetParams field.
func (o *RecoverElastifileParamsRecoverNasVolumeParams) SetNetappTargetParams(v RecoverElastifileNasVolumeParamsNetappTargetParams) {
	o.NetappTargetParams.Set(&v)
}
// SetNetappTargetParamsNil sets the value for NetappTargetParams to be an explicit nil
func (o *RecoverElastifileParamsRecoverNasVolumeParams) SetNetappTargetParamsNil() {
	o.NetappTargetParams.Set(nil)
}

// UnsetNetappTargetParams ensures that no value is present for NetappTargetParams, not even an explicit nil
func (o *RecoverElastifileParamsRecoverNasVolumeParams) UnsetNetappTargetParams() {
	o.NetappTargetParams.Unset()
}

// GetTargetEnvironment returns the TargetEnvironment field value
func (o *RecoverElastifileParamsRecoverNasVolumeParams) GetTargetEnvironment() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TargetEnvironment
}

// GetTargetEnvironmentOk returns a tuple with the TargetEnvironment field value
// and a boolean to check if the value has been set.
func (o *RecoverElastifileParamsRecoverNasVolumeParams) GetTargetEnvironmentOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TargetEnvironment, true
}

// SetTargetEnvironment sets field value
func (o *RecoverElastifileParamsRecoverNasVolumeParams) SetTargetEnvironment(v string) {
	o.TargetEnvironment = v
}

// GetViewTargetParams returns the ViewTargetParams field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RecoverElastifileParamsRecoverNasVolumeParams) GetViewTargetParams() RecoverElastifileNasVolumeParamsViewTargetParams {
	if o == nil || IsNil(o.ViewTargetParams.Get()) {
		var ret RecoverElastifileNasVolumeParamsViewTargetParams
		return ret
	}
	return *o.ViewTargetParams.Get()
}

// GetViewTargetParamsOk returns a tuple with the ViewTargetParams field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RecoverElastifileParamsRecoverNasVolumeParams) GetViewTargetParamsOk() (*RecoverElastifileNasVolumeParamsViewTargetParams, bool) {
	if o == nil {
		return nil, false
	}
	return o.ViewTargetParams.Get(), o.ViewTargetParams.IsSet()
}

// HasViewTargetParams returns a boolean if a field has been set.
func (o *RecoverElastifileParamsRecoverNasVolumeParams) HasViewTargetParams() bool {
	if o != nil && o.ViewTargetParams.IsSet() {
		return true
	}

	return false
}

// SetViewTargetParams gets a reference to the given NullableRecoverElastifileNasVolumeParamsViewTargetParams and assigns it to the ViewTargetParams field.
func (o *RecoverElastifileParamsRecoverNasVolumeParams) SetViewTargetParams(v RecoverElastifileNasVolumeParamsViewTargetParams) {
	o.ViewTargetParams.Set(&v)
}
// SetViewTargetParamsNil sets the value for ViewTargetParams to be an explicit nil
func (o *RecoverElastifileParamsRecoverNasVolumeParams) SetViewTargetParamsNil() {
	o.ViewTargetParams.Set(nil)
}

// UnsetViewTargetParams ensures that no value is present for ViewTargetParams, not even an explicit nil
func (o *RecoverElastifileParamsRecoverNasVolumeParams) UnsetViewTargetParams() {
	o.ViewTargetParams.Unset()
}

func (o RecoverElastifileParamsRecoverNasVolumeParams) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RecoverElastifileParamsRecoverNasVolumeParams) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
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
	if o.NetappTargetParams.IsSet() {
		toSerialize["netappTargetParams"] = o.NetappTargetParams.Get()
	}
	toSerialize["targetEnvironment"] = o.TargetEnvironment
	if o.ViewTargetParams.IsSet() {
		toSerialize["viewTargetParams"] = o.ViewTargetParams.Get()
	}
	return toSerialize, nil
}

func (o *RecoverElastifileParamsRecoverNasVolumeParams) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"targetEnvironment",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varRecoverElastifileParamsRecoverNasVolumeParams := _RecoverElastifileParamsRecoverNasVolumeParams{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varRecoverElastifileParamsRecoverNasVolumeParams)

	if err != nil {
		return err
	}

	*o = RecoverElastifileParamsRecoverNasVolumeParams(varRecoverElastifileParamsRecoverNasVolumeParams)

	return err
}

type NullableRecoverElastifileParamsRecoverNasVolumeParams struct {
	value *RecoverElastifileParamsRecoverNasVolumeParams
	isSet bool
}

func (v NullableRecoverElastifileParamsRecoverNasVolumeParams) Get() *RecoverElastifileParamsRecoverNasVolumeParams {
	return v.value
}

func (v *NullableRecoverElastifileParamsRecoverNasVolumeParams) Set(val *RecoverElastifileParamsRecoverNasVolumeParams) {
	v.value = val
	v.isSet = true
}

func (v NullableRecoverElastifileParamsRecoverNasVolumeParams) IsSet() bool {
	return v.isSet
}

func (v *NullableRecoverElastifileParamsRecoverNasVolumeParams) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRecoverElastifileParamsRecoverNasVolumeParams(val *RecoverElastifileParamsRecoverNasVolumeParams) *NullableRecoverElastifileParamsRecoverNasVolumeParams {
	return &NullableRecoverElastifileParamsRecoverNasVolumeParams{value: val, isSet: true}
}

func (v NullableRecoverElastifileParamsRecoverNasVolumeParams) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRecoverElastifileParamsRecoverNasVolumeParams) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


