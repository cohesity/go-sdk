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

// checks if the RecoverGpfsNasVolumeParams type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RecoverGpfsNasVolumeParams{}

// RecoverGpfsNasVolumeParams Specifies the parameters to recover GPFS NAS volumes.
type RecoverGpfsNasVolumeParams struct {
	ElastifileTargetParams NullableRecoverFlashbladeNasVolumeParamsElastifileTargetParams `json:"elastifileTargetParams,omitempty"`
	FlashbladeTargetParams NullableRecoverGenericNasNasVolumeParamsFlashbladeTargetParams `json:"flashbladeTargetParams,omitempty"`
	GenericNasTargetParams NullableRecoverGpfsNasVolumeParamsGenericNasTargetParams `json:"genericNasTargetParams,omitempty"`
	GpfsTargetParams NullableRecoverGpfsNasVolumeParamsGpfsTargetParams `json:"gpfsTargetParams,omitempty"`
	IsilonTargetParams NullableRecoverElastifileNasVolumeParamsIsilonTargetParams `json:"isilonTargetParams,omitempty"`
	NetappTargetParams NullableRecoverElastifileNasVolumeParamsNetappTargetParams `json:"netappTargetParams,omitempty"`
	// Specifies the environment of the recovery target. The corresponding params below must be filled out.
	TargetEnvironment string `json:"targetEnvironment"`
	ViewTargetParams NullableRecoverElastifileNasVolumeParamsViewTargetParams `json:"viewTargetParams,omitempty"`
}

type _RecoverGpfsNasVolumeParams RecoverGpfsNasVolumeParams

// NewRecoverGpfsNasVolumeParams instantiates a new RecoverGpfsNasVolumeParams object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRecoverGpfsNasVolumeParams(targetEnvironment string) *RecoverGpfsNasVolumeParams {
	this := RecoverGpfsNasVolumeParams{}
	this.TargetEnvironment = targetEnvironment
	return &this
}

// NewRecoverGpfsNasVolumeParamsWithDefaults instantiates a new RecoverGpfsNasVolumeParams object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRecoverGpfsNasVolumeParamsWithDefaults() *RecoverGpfsNasVolumeParams {
	this := RecoverGpfsNasVolumeParams{}
	return &this
}

// GetElastifileTargetParams returns the ElastifileTargetParams field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RecoverGpfsNasVolumeParams) GetElastifileTargetParams() RecoverFlashbladeNasVolumeParamsElastifileTargetParams {
	if o == nil || IsNil(o.ElastifileTargetParams.Get()) {
		var ret RecoverFlashbladeNasVolumeParamsElastifileTargetParams
		return ret
	}
	return *o.ElastifileTargetParams.Get()
}

// GetElastifileTargetParamsOk returns a tuple with the ElastifileTargetParams field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RecoverGpfsNasVolumeParams) GetElastifileTargetParamsOk() (*RecoverFlashbladeNasVolumeParamsElastifileTargetParams, bool) {
	if o == nil {
		return nil, false
	}
	return o.ElastifileTargetParams.Get(), o.ElastifileTargetParams.IsSet()
}

// HasElastifileTargetParams returns a boolean if a field has been set.
func (o *RecoverGpfsNasVolumeParams) HasElastifileTargetParams() bool {
	if o != nil && o.ElastifileTargetParams.IsSet() {
		return true
	}

	return false
}

// SetElastifileTargetParams gets a reference to the given NullableRecoverFlashbladeNasVolumeParamsElastifileTargetParams and assigns it to the ElastifileTargetParams field.
func (o *RecoverGpfsNasVolumeParams) SetElastifileTargetParams(v RecoverFlashbladeNasVolumeParamsElastifileTargetParams) {
	o.ElastifileTargetParams.Set(&v)
}
// SetElastifileTargetParamsNil sets the value for ElastifileTargetParams to be an explicit nil
func (o *RecoverGpfsNasVolumeParams) SetElastifileTargetParamsNil() {
	o.ElastifileTargetParams.Set(nil)
}

// UnsetElastifileTargetParams ensures that no value is present for ElastifileTargetParams, not even an explicit nil
func (o *RecoverGpfsNasVolumeParams) UnsetElastifileTargetParams() {
	o.ElastifileTargetParams.Unset()
}

// GetFlashbladeTargetParams returns the FlashbladeTargetParams field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RecoverGpfsNasVolumeParams) GetFlashbladeTargetParams() RecoverGenericNasNasVolumeParamsFlashbladeTargetParams {
	if o == nil || IsNil(o.FlashbladeTargetParams.Get()) {
		var ret RecoverGenericNasNasVolumeParamsFlashbladeTargetParams
		return ret
	}
	return *o.FlashbladeTargetParams.Get()
}

// GetFlashbladeTargetParamsOk returns a tuple with the FlashbladeTargetParams field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RecoverGpfsNasVolumeParams) GetFlashbladeTargetParamsOk() (*RecoverGenericNasNasVolumeParamsFlashbladeTargetParams, bool) {
	if o == nil {
		return nil, false
	}
	return o.FlashbladeTargetParams.Get(), o.FlashbladeTargetParams.IsSet()
}

// HasFlashbladeTargetParams returns a boolean if a field has been set.
func (o *RecoverGpfsNasVolumeParams) HasFlashbladeTargetParams() bool {
	if o != nil && o.FlashbladeTargetParams.IsSet() {
		return true
	}

	return false
}

// SetFlashbladeTargetParams gets a reference to the given NullableRecoverGenericNasNasVolumeParamsFlashbladeTargetParams and assigns it to the FlashbladeTargetParams field.
func (o *RecoverGpfsNasVolumeParams) SetFlashbladeTargetParams(v RecoverGenericNasNasVolumeParamsFlashbladeTargetParams) {
	o.FlashbladeTargetParams.Set(&v)
}
// SetFlashbladeTargetParamsNil sets the value for FlashbladeTargetParams to be an explicit nil
func (o *RecoverGpfsNasVolumeParams) SetFlashbladeTargetParamsNil() {
	o.FlashbladeTargetParams.Set(nil)
}

// UnsetFlashbladeTargetParams ensures that no value is present for FlashbladeTargetParams, not even an explicit nil
func (o *RecoverGpfsNasVolumeParams) UnsetFlashbladeTargetParams() {
	o.FlashbladeTargetParams.Unset()
}

// GetGenericNasTargetParams returns the GenericNasTargetParams field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RecoverGpfsNasVolumeParams) GetGenericNasTargetParams() RecoverGpfsNasVolumeParamsGenericNasTargetParams {
	if o == nil || IsNil(o.GenericNasTargetParams.Get()) {
		var ret RecoverGpfsNasVolumeParamsGenericNasTargetParams
		return ret
	}
	return *o.GenericNasTargetParams.Get()
}

// GetGenericNasTargetParamsOk returns a tuple with the GenericNasTargetParams field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RecoverGpfsNasVolumeParams) GetGenericNasTargetParamsOk() (*RecoverGpfsNasVolumeParamsGenericNasTargetParams, bool) {
	if o == nil {
		return nil, false
	}
	return o.GenericNasTargetParams.Get(), o.GenericNasTargetParams.IsSet()
}

// HasGenericNasTargetParams returns a boolean if a field has been set.
func (o *RecoverGpfsNasVolumeParams) HasGenericNasTargetParams() bool {
	if o != nil && o.GenericNasTargetParams.IsSet() {
		return true
	}

	return false
}

// SetGenericNasTargetParams gets a reference to the given NullableRecoverGpfsNasVolumeParamsGenericNasTargetParams and assigns it to the GenericNasTargetParams field.
func (o *RecoverGpfsNasVolumeParams) SetGenericNasTargetParams(v RecoverGpfsNasVolumeParamsGenericNasTargetParams) {
	o.GenericNasTargetParams.Set(&v)
}
// SetGenericNasTargetParamsNil sets the value for GenericNasTargetParams to be an explicit nil
func (o *RecoverGpfsNasVolumeParams) SetGenericNasTargetParamsNil() {
	o.GenericNasTargetParams.Set(nil)
}

// UnsetGenericNasTargetParams ensures that no value is present for GenericNasTargetParams, not even an explicit nil
func (o *RecoverGpfsNasVolumeParams) UnsetGenericNasTargetParams() {
	o.GenericNasTargetParams.Unset()
}

// GetGpfsTargetParams returns the GpfsTargetParams field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RecoverGpfsNasVolumeParams) GetGpfsTargetParams() RecoverGpfsNasVolumeParamsGpfsTargetParams {
	if o == nil || IsNil(o.GpfsTargetParams.Get()) {
		var ret RecoverGpfsNasVolumeParamsGpfsTargetParams
		return ret
	}
	return *o.GpfsTargetParams.Get()
}

// GetGpfsTargetParamsOk returns a tuple with the GpfsTargetParams field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RecoverGpfsNasVolumeParams) GetGpfsTargetParamsOk() (*RecoverGpfsNasVolumeParamsGpfsTargetParams, bool) {
	if o == nil {
		return nil, false
	}
	return o.GpfsTargetParams.Get(), o.GpfsTargetParams.IsSet()
}

// HasGpfsTargetParams returns a boolean if a field has been set.
func (o *RecoverGpfsNasVolumeParams) HasGpfsTargetParams() bool {
	if o != nil && o.GpfsTargetParams.IsSet() {
		return true
	}

	return false
}

// SetGpfsTargetParams gets a reference to the given NullableRecoverGpfsNasVolumeParamsGpfsTargetParams and assigns it to the GpfsTargetParams field.
func (o *RecoverGpfsNasVolumeParams) SetGpfsTargetParams(v RecoverGpfsNasVolumeParamsGpfsTargetParams) {
	o.GpfsTargetParams.Set(&v)
}
// SetGpfsTargetParamsNil sets the value for GpfsTargetParams to be an explicit nil
func (o *RecoverGpfsNasVolumeParams) SetGpfsTargetParamsNil() {
	o.GpfsTargetParams.Set(nil)
}

// UnsetGpfsTargetParams ensures that no value is present for GpfsTargetParams, not even an explicit nil
func (o *RecoverGpfsNasVolumeParams) UnsetGpfsTargetParams() {
	o.GpfsTargetParams.Unset()
}

// GetIsilonTargetParams returns the IsilonTargetParams field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RecoverGpfsNasVolumeParams) GetIsilonTargetParams() RecoverElastifileNasVolumeParamsIsilonTargetParams {
	if o == nil || IsNil(o.IsilonTargetParams.Get()) {
		var ret RecoverElastifileNasVolumeParamsIsilonTargetParams
		return ret
	}
	return *o.IsilonTargetParams.Get()
}

// GetIsilonTargetParamsOk returns a tuple with the IsilonTargetParams field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RecoverGpfsNasVolumeParams) GetIsilonTargetParamsOk() (*RecoverElastifileNasVolumeParamsIsilonTargetParams, bool) {
	if o == nil {
		return nil, false
	}
	return o.IsilonTargetParams.Get(), o.IsilonTargetParams.IsSet()
}

// HasIsilonTargetParams returns a boolean if a field has been set.
func (o *RecoverGpfsNasVolumeParams) HasIsilonTargetParams() bool {
	if o != nil && o.IsilonTargetParams.IsSet() {
		return true
	}

	return false
}

// SetIsilonTargetParams gets a reference to the given NullableRecoverElastifileNasVolumeParamsIsilonTargetParams and assigns it to the IsilonTargetParams field.
func (o *RecoverGpfsNasVolumeParams) SetIsilonTargetParams(v RecoverElastifileNasVolumeParamsIsilonTargetParams) {
	o.IsilonTargetParams.Set(&v)
}
// SetIsilonTargetParamsNil sets the value for IsilonTargetParams to be an explicit nil
func (o *RecoverGpfsNasVolumeParams) SetIsilonTargetParamsNil() {
	o.IsilonTargetParams.Set(nil)
}

// UnsetIsilonTargetParams ensures that no value is present for IsilonTargetParams, not even an explicit nil
func (o *RecoverGpfsNasVolumeParams) UnsetIsilonTargetParams() {
	o.IsilonTargetParams.Unset()
}

// GetNetappTargetParams returns the NetappTargetParams field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RecoverGpfsNasVolumeParams) GetNetappTargetParams() RecoverElastifileNasVolumeParamsNetappTargetParams {
	if o == nil || IsNil(o.NetappTargetParams.Get()) {
		var ret RecoverElastifileNasVolumeParamsNetappTargetParams
		return ret
	}
	return *o.NetappTargetParams.Get()
}

// GetNetappTargetParamsOk returns a tuple with the NetappTargetParams field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RecoverGpfsNasVolumeParams) GetNetappTargetParamsOk() (*RecoverElastifileNasVolumeParamsNetappTargetParams, bool) {
	if o == nil {
		return nil, false
	}
	return o.NetappTargetParams.Get(), o.NetappTargetParams.IsSet()
}

// HasNetappTargetParams returns a boolean if a field has been set.
func (o *RecoverGpfsNasVolumeParams) HasNetappTargetParams() bool {
	if o != nil && o.NetappTargetParams.IsSet() {
		return true
	}

	return false
}

// SetNetappTargetParams gets a reference to the given NullableRecoverElastifileNasVolumeParamsNetappTargetParams and assigns it to the NetappTargetParams field.
func (o *RecoverGpfsNasVolumeParams) SetNetappTargetParams(v RecoverElastifileNasVolumeParamsNetappTargetParams) {
	o.NetappTargetParams.Set(&v)
}
// SetNetappTargetParamsNil sets the value for NetappTargetParams to be an explicit nil
func (o *RecoverGpfsNasVolumeParams) SetNetappTargetParamsNil() {
	o.NetappTargetParams.Set(nil)
}

// UnsetNetappTargetParams ensures that no value is present for NetappTargetParams, not even an explicit nil
func (o *RecoverGpfsNasVolumeParams) UnsetNetappTargetParams() {
	o.NetappTargetParams.Unset()
}

// GetTargetEnvironment returns the TargetEnvironment field value
func (o *RecoverGpfsNasVolumeParams) GetTargetEnvironment() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TargetEnvironment
}

// GetTargetEnvironmentOk returns a tuple with the TargetEnvironment field value
// and a boolean to check if the value has been set.
func (o *RecoverGpfsNasVolumeParams) GetTargetEnvironmentOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TargetEnvironment, true
}

// SetTargetEnvironment sets field value
func (o *RecoverGpfsNasVolumeParams) SetTargetEnvironment(v string) {
	o.TargetEnvironment = v
}

// GetViewTargetParams returns the ViewTargetParams field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RecoverGpfsNasVolumeParams) GetViewTargetParams() RecoverElastifileNasVolumeParamsViewTargetParams {
	if o == nil || IsNil(o.ViewTargetParams.Get()) {
		var ret RecoverElastifileNasVolumeParamsViewTargetParams
		return ret
	}
	return *o.ViewTargetParams.Get()
}

// GetViewTargetParamsOk returns a tuple with the ViewTargetParams field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RecoverGpfsNasVolumeParams) GetViewTargetParamsOk() (*RecoverElastifileNasVolumeParamsViewTargetParams, bool) {
	if o == nil {
		return nil, false
	}
	return o.ViewTargetParams.Get(), o.ViewTargetParams.IsSet()
}

// HasViewTargetParams returns a boolean if a field has been set.
func (o *RecoverGpfsNasVolumeParams) HasViewTargetParams() bool {
	if o != nil && o.ViewTargetParams.IsSet() {
		return true
	}

	return false
}

// SetViewTargetParams gets a reference to the given NullableRecoverElastifileNasVolumeParamsViewTargetParams and assigns it to the ViewTargetParams field.
func (o *RecoverGpfsNasVolumeParams) SetViewTargetParams(v RecoverElastifileNasVolumeParamsViewTargetParams) {
	o.ViewTargetParams.Set(&v)
}
// SetViewTargetParamsNil sets the value for ViewTargetParams to be an explicit nil
func (o *RecoverGpfsNasVolumeParams) SetViewTargetParamsNil() {
	o.ViewTargetParams.Set(nil)
}

// UnsetViewTargetParams ensures that no value is present for ViewTargetParams, not even an explicit nil
func (o *RecoverGpfsNasVolumeParams) UnsetViewTargetParams() {
	o.ViewTargetParams.Unset()
}

func (o RecoverGpfsNasVolumeParams) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RecoverGpfsNasVolumeParams) ToMap() (map[string]interface{}, error) {
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

func (o *RecoverGpfsNasVolumeParams) UnmarshalJSON(data []byte) (err error) {
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

	varRecoverGpfsNasVolumeParams := _RecoverGpfsNasVolumeParams{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varRecoverGpfsNasVolumeParams)

	if err != nil {
		return err
	}

	*o = RecoverGpfsNasVolumeParams(varRecoverGpfsNasVolumeParams)

	return err
}

type NullableRecoverGpfsNasVolumeParams struct {
	value *RecoverGpfsNasVolumeParams
	isSet bool
}

func (v NullableRecoverGpfsNasVolumeParams) Get() *RecoverGpfsNasVolumeParams {
	return v.value
}

func (v *NullableRecoverGpfsNasVolumeParams) Set(val *RecoverGpfsNasVolumeParams) {
	v.value = val
	v.isSet = true
}

func (v NullableRecoverGpfsNasVolumeParams) IsSet() bool {
	return v.isSet
}

func (v *NullableRecoverGpfsNasVolumeParams) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRecoverGpfsNasVolumeParams(val *RecoverGpfsNasVolumeParams) *NullableRecoverGpfsNasVolumeParams {
	return &NullableRecoverGpfsNasVolumeParams{value: val, isSet: true}
}

func (v NullableRecoverGpfsNasVolumeParams) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRecoverGpfsNasVolumeParams) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


