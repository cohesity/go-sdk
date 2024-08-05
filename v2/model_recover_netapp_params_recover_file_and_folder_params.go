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

// checks if the RecoverNetappParamsRecoverFileAndFolderParams type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RecoverNetappParamsRecoverFileAndFolderParams{}

// RecoverNetappParamsRecoverFileAndFolderParams Specifies the parameters to recover files.
type RecoverNetappParamsRecoverFileAndFolderParams struct {
	ElastifileTargetParams NullableRecoverNetappFilesParamsElastifileTargetParams `json:"elastifileTargetParams,omitempty"`
	// Specifies the list of info about the netapp files and folders to be recovered.
	FilesAndFolders []NetappRecoverFileAndFolderInfo `json:"filesAndFolders"`
	FlashbladeTargetParams NullableRecoverNetappFilesParamsFlashbladeTargetParams `json:"flashbladeTargetParams,omitempty"`
	GenericNasTargetParams NullableRecoverNetappFilesParamsGenericNasTargetParams `json:"genericNasTargetParams,omitempty"`
	GpfsTargetParams NullableRecoverNetappFilesParamsGpfsTargetParams `json:"gpfsTargetParams,omitempty"`
	// Specifies if the snapshot trying to recover is from a source initiated protection.
	IsFromSourceInitiatedProtection NullableBool `json:"isFromSourceInitiatedProtection,omitempty"`
	IsilonTargetParams NullableRecoverNetappFilesParamsIsilonTargetParams `json:"isilonTargetParams,omitempty"`
	NetappTargetParams NullableRecoverNetappFilesParamsNetappTargetParams `json:"netappTargetParams,omitempty"`
	// Specifies the environment of the recovery target. The corresponding params below must be filled out.
	TargetEnvironment string `json:"targetEnvironment"`
}

type _RecoverNetappParamsRecoverFileAndFolderParams RecoverNetappParamsRecoverFileAndFolderParams

// NewRecoverNetappParamsRecoverFileAndFolderParams instantiates a new RecoverNetappParamsRecoverFileAndFolderParams object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRecoverNetappParamsRecoverFileAndFolderParams(filesAndFolders []NetappRecoverFileAndFolderInfo, targetEnvironment string) *RecoverNetappParamsRecoverFileAndFolderParams {
	this := RecoverNetappParamsRecoverFileAndFolderParams{}
	this.FilesAndFolders = filesAndFolders
	this.TargetEnvironment = targetEnvironment
	return &this
}

// NewRecoverNetappParamsRecoverFileAndFolderParamsWithDefaults instantiates a new RecoverNetappParamsRecoverFileAndFolderParams object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRecoverNetappParamsRecoverFileAndFolderParamsWithDefaults() *RecoverNetappParamsRecoverFileAndFolderParams {
	this := RecoverNetappParamsRecoverFileAndFolderParams{}
	return &this
}

// GetElastifileTargetParams returns the ElastifileTargetParams field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RecoverNetappParamsRecoverFileAndFolderParams) GetElastifileTargetParams() RecoverNetappFilesParamsElastifileTargetParams {
	if o == nil || IsNil(o.ElastifileTargetParams.Get()) {
		var ret RecoverNetappFilesParamsElastifileTargetParams
		return ret
	}
	return *o.ElastifileTargetParams.Get()
}

// GetElastifileTargetParamsOk returns a tuple with the ElastifileTargetParams field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RecoverNetappParamsRecoverFileAndFolderParams) GetElastifileTargetParamsOk() (*RecoverNetappFilesParamsElastifileTargetParams, bool) {
	if o == nil {
		return nil, false
	}
	return o.ElastifileTargetParams.Get(), o.ElastifileTargetParams.IsSet()
}

// HasElastifileTargetParams returns a boolean if a field has been set.
func (o *RecoverNetappParamsRecoverFileAndFolderParams) HasElastifileTargetParams() bool {
	if o != nil && o.ElastifileTargetParams.IsSet() {
		return true
	}

	return false
}

// SetElastifileTargetParams gets a reference to the given NullableRecoverNetappFilesParamsElastifileTargetParams and assigns it to the ElastifileTargetParams field.
func (o *RecoverNetappParamsRecoverFileAndFolderParams) SetElastifileTargetParams(v RecoverNetappFilesParamsElastifileTargetParams) {
	o.ElastifileTargetParams.Set(&v)
}
// SetElastifileTargetParamsNil sets the value for ElastifileTargetParams to be an explicit nil
func (o *RecoverNetappParamsRecoverFileAndFolderParams) SetElastifileTargetParamsNil() {
	o.ElastifileTargetParams.Set(nil)
}

// UnsetElastifileTargetParams ensures that no value is present for ElastifileTargetParams, not even an explicit nil
func (o *RecoverNetappParamsRecoverFileAndFolderParams) UnsetElastifileTargetParams() {
	o.ElastifileTargetParams.Unset()
}

// GetFilesAndFolders returns the FilesAndFolders field value
// If the value is explicit nil, the zero value for []NetappRecoverFileAndFolderInfo will be returned
func (o *RecoverNetappParamsRecoverFileAndFolderParams) GetFilesAndFolders() []NetappRecoverFileAndFolderInfo {
	if o == nil {
		var ret []NetappRecoverFileAndFolderInfo
		return ret
	}

	return o.FilesAndFolders
}

// GetFilesAndFoldersOk returns a tuple with the FilesAndFolders field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RecoverNetappParamsRecoverFileAndFolderParams) GetFilesAndFoldersOk() ([]NetappRecoverFileAndFolderInfo, bool) {
	if o == nil || IsNil(o.FilesAndFolders) {
		return nil, false
	}
	return o.FilesAndFolders, true
}

// SetFilesAndFolders sets field value
func (o *RecoverNetappParamsRecoverFileAndFolderParams) SetFilesAndFolders(v []NetappRecoverFileAndFolderInfo) {
	o.FilesAndFolders = v
}

// GetFlashbladeTargetParams returns the FlashbladeTargetParams field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RecoverNetappParamsRecoverFileAndFolderParams) GetFlashbladeTargetParams() RecoverNetappFilesParamsFlashbladeTargetParams {
	if o == nil || IsNil(o.FlashbladeTargetParams.Get()) {
		var ret RecoverNetappFilesParamsFlashbladeTargetParams
		return ret
	}
	return *o.FlashbladeTargetParams.Get()
}

// GetFlashbladeTargetParamsOk returns a tuple with the FlashbladeTargetParams field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RecoverNetappParamsRecoverFileAndFolderParams) GetFlashbladeTargetParamsOk() (*RecoverNetappFilesParamsFlashbladeTargetParams, bool) {
	if o == nil {
		return nil, false
	}
	return o.FlashbladeTargetParams.Get(), o.FlashbladeTargetParams.IsSet()
}

// HasFlashbladeTargetParams returns a boolean if a field has been set.
func (o *RecoverNetappParamsRecoverFileAndFolderParams) HasFlashbladeTargetParams() bool {
	if o != nil && o.FlashbladeTargetParams.IsSet() {
		return true
	}

	return false
}

// SetFlashbladeTargetParams gets a reference to the given NullableRecoverNetappFilesParamsFlashbladeTargetParams and assigns it to the FlashbladeTargetParams field.
func (o *RecoverNetappParamsRecoverFileAndFolderParams) SetFlashbladeTargetParams(v RecoverNetappFilesParamsFlashbladeTargetParams) {
	o.FlashbladeTargetParams.Set(&v)
}
// SetFlashbladeTargetParamsNil sets the value for FlashbladeTargetParams to be an explicit nil
func (o *RecoverNetappParamsRecoverFileAndFolderParams) SetFlashbladeTargetParamsNil() {
	o.FlashbladeTargetParams.Set(nil)
}

// UnsetFlashbladeTargetParams ensures that no value is present for FlashbladeTargetParams, not even an explicit nil
func (o *RecoverNetappParamsRecoverFileAndFolderParams) UnsetFlashbladeTargetParams() {
	o.FlashbladeTargetParams.Unset()
}

// GetGenericNasTargetParams returns the GenericNasTargetParams field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RecoverNetappParamsRecoverFileAndFolderParams) GetGenericNasTargetParams() RecoverNetappFilesParamsGenericNasTargetParams {
	if o == nil || IsNil(o.GenericNasTargetParams.Get()) {
		var ret RecoverNetappFilesParamsGenericNasTargetParams
		return ret
	}
	return *o.GenericNasTargetParams.Get()
}

// GetGenericNasTargetParamsOk returns a tuple with the GenericNasTargetParams field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RecoverNetappParamsRecoverFileAndFolderParams) GetGenericNasTargetParamsOk() (*RecoverNetappFilesParamsGenericNasTargetParams, bool) {
	if o == nil {
		return nil, false
	}
	return o.GenericNasTargetParams.Get(), o.GenericNasTargetParams.IsSet()
}

// HasGenericNasTargetParams returns a boolean if a field has been set.
func (o *RecoverNetappParamsRecoverFileAndFolderParams) HasGenericNasTargetParams() bool {
	if o != nil && o.GenericNasTargetParams.IsSet() {
		return true
	}

	return false
}

// SetGenericNasTargetParams gets a reference to the given NullableRecoverNetappFilesParamsGenericNasTargetParams and assigns it to the GenericNasTargetParams field.
func (o *RecoverNetappParamsRecoverFileAndFolderParams) SetGenericNasTargetParams(v RecoverNetappFilesParamsGenericNasTargetParams) {
	o.GenericNasTargetParams.Set(&v)
}
// SetGenericNasTargetParamsNil sets the value for GenericNasTargetParams to be an explicit nil
func (o *RecoverNetappParamsRecoverFileAndFolderParams) SetGenericNasTargetParamsNil() {
	o.GenericNasTargetParams.Set(nil)
}

// UnsetGenericNasTargetParams ensures that no value is present for GenericNasTargetParams, not even an explicit nil
func (o *RecoverNetappParamsRecoverFileAndFolderParams) UnsetGenericNasTargetParams() {
	o.GenericNasTargetParams.Unset()
}

// GetGpfsTargetParams returns the GpfsTargetParams field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RecoverNetappParamsRecoverFileAndFolderParams) GetGpfsTargetParams() RecoverNetappFilesParamsGpfsTargetParams {
	if o == nil || IsNil(o.GpfsTargetParams.Get()) {
		var ret RecoverNetappFilesParamsGpfsTargetParams
		return ret
	}
	return *o.GpfsTargetParams.Get()
}

// GetGpfsTargetParamsOk returns a tuple with the GpfsTargetParams field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RecoverNetappParamsRecoverFileAndFolderParams) GetGpfsTargetParamsOk() (*RecoverNetappFilesParamsGpfsTargetParams, bool) {
	if o == nil {
		return nil, false
	}
	return o.GpfsTargetParams.Get(), o.GpfsTargetParams.IsSet()
}

// HasGpfsTargetParams returns a boolean if a field has been set.
func (o *RecoverNetappParamsRecoverFileAndFolderParams) HasGpfsTargetParams() bool {
	if o != nil && o.GpfsTargetParams.IsSet() {
		return true
	}

	return false
}

// SetGpfsTargetParams gets a reference to the given NullableRecoverNetappFilesParamsGpfsTargetParams and assigns it to the GpfsTargetParams field.
func (o *RecoverNetappParamsRecoverFileAndFolderParams) SetGpfsTargetParams(v RecoverNetappFilesParamsGpfsTargetParams) {
	o.GpfsTargetParams.Set(&v)
}
// SetGpfsTargetParamsNil sets the value for GpfsTargetParams to be an explicit nil
func (o *RecoverNetappParamsRecoverFileAndFolderParams) SetGpfsTargetParamsNil() {
	o.GpfsTargetParams.Set(nil)
}

// UnsetGpfsTargetParams ensures that no value is present for GpfsTargetParams, not even an explicit nil
func (o *RecoverNetappParamsRecoverFileAndFolderParams) UnsetGpfsTargetParams() {
	o.GpfsTargetParams.Unset()
}

// GetIsFromSourceInitiatedProtection returns the IsFromSourceInitiatedProtection field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RecoverNetappParamsRecoverFileAndFolderParams) GetIsFromSourceInitiatedProtection() bool {
	if o == nil || IsNil(o.IsFromSourceInitiatedProtection.Get()) {
		var ret bool
		return ret
	}
	return *o.IsFromSourceInitiatedProtection.Get()
}

// GetIsFromSourceInitiatedProtectionOk returns a tuple with the IsFromSourceInitiatedProtection field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RecoverNetappParamsRecoverFileAndFolderParams) GetIsFromSourceInitiatedProtectionOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.IsFromSourceInitiatedProtection.Get(), o.IsFromSourceInitiatedProtection.IsSet()
}

// HasIsFromSourceInitiatedProtection returns a boolean if a field has been set.
func (o *RecoverNetappParamsRecoverFileAndFolderParams) HasIsFromSourceInitiatedProtection() bool {
	if o != nil && o.IsFromSourceInitiatedProtection.IsSet() {
		return true
	}

	return false
}

// SetIsFromSourceInitiatedProtection gets a reference to the given NullableBool and assigns it to the IsFromSourceInitiatedProtection field.
func (o *RecoverNetappParamsRecoverFileAndFolderParams) SetIsFromSourceInitiatedProtection(v bool) {
	o.IsFromSourceInitiatedProtection.Set(&v)
}
// SetIsFromSourceInitiatedProtectionNil sets the value for IsFromSourceInitiatedProtection to be an explicit nil
func (o *RecoverNetappParamsRecoverFileAndFolderParams) SetIsFromSourceInitiatedProtectionNil() {
	o.IsFromSourceInitiatedProtection.Set(nil)
}

// UnsetIsFromSourceInitiatedProtection ensures that no value is present for IsFromSourceInitiatedProtection, not even an explicit nil
func (o *RecoverNetappParamsRecoverFileAndFolderParams) UnsetIsFromSourceInitiatedProtection() {
	o.IsFromSourceInitiatedProtection.Unset()
}

// GetIsilonTargetParams returns the IsilonTargetParams field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RecoverNetappParamsRecoverFileAndFolderParams) GetIsilonTargetParams() RecoverNetappFilesParamsIsilonTargetParams {
	if o == nil || IsNil(o.IsilonTargetParams.Get()) {
		var ret RecoverNetappFilesParamsIsilonTargetParams
		return ret
	}
	return *o.IsilonTargetParams.Get()
}

// GetIsilonTargetParamsOk returns a tuple with the IsilonTargetParams field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RecoverNetappParamsRecoverFileAndFolderParams) GetIsilonTargetParamsOk() (*RecoverNetappFilesParamsIsilonTargetParams, bool) {
	if o == nil {
		return nil, false
	}
	return o.IsilonTargetParams.Get(), o.IsilonTargetParams.IsSet()
}

// HasIsilonTargetParams returns a boolean if a field has been set.
func (o *RecoverNetappParamsRecoverFileAndFolderParams) HasIsilonTargetParams() bool {
	if o != nil && o.IsilonTargetParams.IsSet() {
		return true
	}

	return false
}

// SetIsilonTargetParams gets a reference to the given NullableRecoverNetappFilesParamsIsilonTargetParams and assigns it to the IsilonTargetParams field.
func (o *RecoverNetappParamsRecoverFileAndFolderParams) SetIsilonTargetParams(v RecoverNetappFilesParamsIsilonTargetParams) {
	o.IsilonTargetParams.Set(&v)
}
// SetIsilonTargetParamsNil sets the value for IsilonTargetParams to be an explicit nil
func (o *RecoverNetappParamsRecoverFileAndFolderParams) SetIsilonTargetParamsNil() {
	o.IsilonTargetParams.Set(nil)
}

// UnsetIsilonTargetParams ensures that no value is present for IsilonTargetParams, not even an explicit nil
func (o *RecoverNetappParamsRecoverFileAndFolderParams) UnsetIsilonTargetParams() {
	o.IsilonTargetParams.Unset()
}

// GetNetappTargetParams returns the NetappTargetParams field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RecoverNetappParamsRecoverFileAndFolderParams) GetNetappTargetParams() RecoverNetappFilesParamsNetappTargetParams {
	if o == nil || IsNil(o.NetappTargetParams.Get()) {
		var ret RecoverNetappFilesParamsNetappTargetParams
		return ret
	}
	return *o.NetappTargetParams.Get()
}

// GetNetappTargetParamsOk returns a tuple with the NetappTargetParams field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RecoverNetappParamsRecoverFileAndFolderParams) GetNetappTargetParamsOk() (*RecoverNetappFilesParamsNetappTargetParams, bool) {
	if o == nil {
		return nil, false
	}
	return o.NetappTargetParams.Get(), o.NetappTargetParams.IsSet()
}

// HasNetappTargetParams returns a boolean if a field has been set.
func (o *RecoverNetappParamsRecoverFileAndFolderParams) HasNetappTargetParams() bool {
	if o != nil && o.NetappTargetParams.IsSet() {
		return true
	}

	return false
}

// SetNetappTargetParams gets a reference to the given NullableRecoverNetappFilesParamsNetappTargetParams and assigns it to the NetappTargetParams field.
func (o *RecoverNetappParamsRecoverFileAndFolderParams) SetNetappTargetParams(v RecoverNetappFilesParamsNetappTargetParams) {
	o.NetappTargetParams.Set(&v)
}
// SetNetappTargetParamsNil sets the value for NetappTargetParams to be an explicit nil
func (o *RecoverNetappParamsRecoverFileAndFolderParams) SetNetappTargetParamsNil() {
	o.NetappTargetParams.Set(nil)
}

// UnsetNetappTargetParams ensures that no value is present for NetappTargetParams, not even an explicit nil
func (o *RecoverNetappParamsRecoverFileAndFolderParams) UnsetNetappTargetParams() {
	o.NetappTargetParams.Unset()
}

// GetTargetEnvironment returns the TargetEnvironment field value
func (o *RecoverNetappParamsRecoverFileAndFolderParams) GetTargetEnvironment() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TargetEnvironment
}

// GetTargetEnvironmentOk returns a tuple with the TargetEnvironment field value
// and a boolean to check if the value has been set.
func (o *RecoverNetappParamsRecoverFileAndFolderParams) GetTargetEnvironmentOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TargetEnvironment, true
}

// SetTargetEnvironment sets field value
func (o *RecoverNetappParamsRecoverFileAndFolderParams) SetTargetEnvironment(v string) {
	o.TargetEnvironment = v
}

func (o RecoverNetappParamsRecoverFileAndFolderParams) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RecoverNetappParamsRecoverFileAndFolderParams) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.ElastifileTargetParams.IsSet() {
		toSerialize["elastifileTargetParams"] = o.ElastifileTargetParams.Get()
	}
	if o.FilesAndFolders != nil {
		toSerialize["filesAndFolders"] = o.FilesAndFolders
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
	if o.IsFromSourceInitiatedProtection.IsSet() {
		toSerialize["isFromSourceInitiatedProtection"] = o.IsFromSourceInitiatedProtection.Get()
	}
	if o.IsilonTargetParams.IsSet() {
		toSerialize["isilonTargetParams"] = o.IsilonTargetParams.Get()
	}
	if o.NetappTargetParams.IsSet() {
		toSerialize["netappTargetParams"] = o.NetappTargetParams.Get()
	}
	toSerialize["targetEnvironment"] = o.TargetEnvironment
	return toSerialize, nil
}

func (o *RecoverNetappParamsRecoverFileAndFolderParams) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"filesAndFolders",
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

	varRecoverNetappParamsRecoverFileAndFolderParams := _RecoverNetappParamsRecoverFileAndFolderParams{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varRecoverNetappParamsRecoverFileAndFolderParams)

	if err != nil {
		return err
	}

	*o = RecoverNetappParamsRecoverFileAndFolderParams(varRecoverNetappParamsRecoverFileAndFolderParams)

	return err
}

type NullableRecoverNetappParamsRecoverFileAndFolderParams struct {
	value *RecoverNetappParamsRecoverFileAndFolderParams
	isSet bool
}

func (v NullableRecoverNetappParamsRecoverFileAndFolderParams) Get() *RecoverNetappParamsRecoverFileAndFolderParams {
	return v.value
}

func (v *NullableRecoverNetappParamsRecoverFileAndFolderParams) Set(val *RecoverNetappParamsRecoverFileAndFolderParams) {
	v.value = val
	v.isSet = true
}

func (v NullableRecoverNetappParamsRecoverFileAndFolderParams) IsSet() bool {
	return v.isSet
}

func (v *NullableRecoverNetappParamsRecoverFileAndFolderParams) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRecoverNetappParamsRecoverFileAndFolderParams(val *RecoverNetappParamsRecoverFileAndFolderParams) *NullableRecoverNetappParamsRecoverFileAndFolderParams {
	return &NullableRecoverNetappParamsRecoverFileAndFolderParams{value: val, isSet: true}
}

func (v NullableRecoverNetappParamsRecoverFileAndFolderParams) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRecoverNetappParamsRecoverFileAndFolderParams) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


