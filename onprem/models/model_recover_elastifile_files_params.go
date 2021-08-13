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

// RecoverElastifileFilesParams Specifies the parameters to recover Elastifile files.
type RecoverElastifileFilesParams struct {
	// Specifies the info about the files and folders to be recovered.
	FilesAndFolders []CommonRecoverFileAndFolderInfo `json:"filesAndFolders"`
	// Specifies the environment of the recovery target. The corresponding params below must be filled out.
	TargetEnvironment string `json:"targetEnvironment"`
	// Specifies the params for a Elastifile recovery target.
	ElastifileTargetParams NullableRecoverElastifileToElastifileFilesTargetParams `json:"elastifileTargetParams,omitempty"`
	// Specifies the params for a Flashblade recovery target.
	FlashbladeTargetParams NullableRecoverOtherNasToFlashbladeFilesTargetParams `json:"flashbladeTargetParams,omitempty"`
	// Specifies the params for a generic NAS recovery target.
	GenericNasTargetParams NullableRecoverOtherNasToGenericNasFilesTargetParams `json:"genericNasTargetParams,omitempty"`
	// Specifies the params for a GPFS recovery target.
	GpfsTargetParams NullableRecoverOtherNasToGpfsFilesTargetParams `json:"gpfsTargetParams,omitempty"`
	// Specifies the params for an Isilon recovery target.
	IsilonTargetParams NullableRecoverOtherNasToIsilonFilesTargetParams `json:"isilonTargetParams,omitempty"`
	// Specifies the params for an Netapp recovery target.
	NetappTargetParams NullableRecoverOtherNasToNetappFilesTargetParams `json:"netappTargetParams,omitempty"`
}

// NewRecoverElastifileFilesParams instantiates a new RecoverElastifileFilesParams object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRecoverElastifileFilesParams(filesAndFolders []CommonRecoverFileAndFolderInfo, targetEnvironment string) *RecoverElastifileFilesParams {
	this := RecoverElastifileFilesParams{}
	this.FilesAndFolders = filesAndFolders
	this.TargetEnvironment = targetEnvironment
	return &this
}

// NewRecoverElastifileFilesParamsWithDefaults instantiates a new RecoverElastifileFilesParams object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRecoverElastifileFilesParamsWithDefaults() *RecoverElastifileFilesParams {
	this := RecoverElastifileFilesParams{}
	return &this
}

// GetFilesAndFolders returns the FilesAndFolders field value
// If the value is explicit nil, the zero value for []CommonRecoverFileAndFolderInfo will be returned
func (o *RecoverElastifileFilesParams) GetFilesAndFolders() []CommonRecoverFileAndFolderInfo {
	if o == nil {
		var ret []CommonRecoverFileAndFolderInfo
		return ret
	}

	return o.FilesAndFolders
}

// GetFilesAndFoldersOk returns a tuple with the FilesAndFolders field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RecoverElastifileFilesParams) GetFilesAndFoldersOk() (*[]CommonRecoverFileAndFolderInfo, bool) {
	if o == nil || o.FilesAndFolders == nil {
		return nil, false
	}
	return &o.FilesAndFolders, true
}

// SetFilesAndFolders sets field value
func (o *RecoverElastifileFilesParams) SetFilesAndFolders(v []CommonRecoverFileAndFolderInfo) {
	o.FilesAndFolders = v
}

// GetTargetEnvironment returns the TargetEnvironment field value
func (o *RecoverElastifileFilesParams) GetTargetEnvironment() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TargetEnvironment
}

// GetTargetEnvironmentOk returns a tuple with the TargetEnvironment field value
// and a boolean to check if the value has been set.
func (o *RecoverElastifileFilesParams) GetTargetEnvironmentOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.TargetEnvironment, true
}

// SetTargetEnvironment sets field value
func (o *RecoverElastifileFilesParams) SetTargetEnvironment(v string) {
	o.TargetEnvironment = v
}

// GetElastifileTargetParams returns the ElastifileTargetParams field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RecoverElastifileFilesParams) GetElastifileTargetParams() RecoverElastifileToElastifileFilesTargetParams {
	if o == nil || o.ElastifileTargetParams.Get() == nil {
		var ret RecoverElastifileToElastifileFilesTargetParams
		return ret
	}
	return *o.ElastifileTargetParams.Get()
}

// GetElastifileTargetParamsOk returns a tuple with the ElastifileTargetParams field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RecoverElastifileFilesParams) GetElastifileTargetParamsOk() (*RecoverElastifileToElastifileFilesTargetParams, bool) {
	if o == nil  {
		return nil, false
	}
	return o.ElastifileTargetParams.Get(), o.ElastifileTargetParams.IsSet()
}

// HasElastifileTargetParams returns a boolean if a field has been set.
func (o *RecoverElastifileFilesParams) HasElastifileTargetParams() bool {
	if o != nil && o.ElastifileTargetParams.IsSet() {
		return true
	}

	return false
}

// SetElastifileTargetParams gets a reference to the given NullableRecoverElastifileToElastifileFilesTargetParams and assigns it to the ElastifileTargetParams field.
func (o *RecoverElastifileFilesParams) SetElastifileTargetParams(v RecoverElastifileToElastifileFilesTargetParams) {
	o.ElastifileTargetParams.Set(&v)
}
// SetElastifileTargetParamsNil sets the value for ElastifileTargetParams to be an explicit nil
func (o *RecoverElastifileFilesParams) SetElastifileTargetParamsNil() {
	o.ElastifileTargetParams.Set(nil)
}

// UnsetElastifileTargetParams ensures that no value is present for ElastifileTargetParams, not even an explicit nil
func (o *RecoverElastifileFilesParams) UnsetElastifileTargetParams() {
	o.ElastifileTargetParams.Unset()
}

// GetFlashbladeTargetParams returns the FlashbladeTargetParams field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RecoverElastifileFilesParams) GetFlashbladeTargetParams() RecoverOtherNasToFlashbladeFilesTargetParams {
	if o == nil || o.FlashbladeTargetParams.Get() == nil {
		var ret RecoverOtherNasToFlashbladeFilesTargetParams
		return ret
	}
	return *o.FlashbladeTargetParams.Get()
}

// GetFlashbladeTargetParamsOk returns a tuple with the FlashbladeTargetParams field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RecoverElastifileFilesParams) GetFlashbladeTargetParamsOk() (*RecoverOtherNasToFlashbladeFilesTargetParams, bool) {
	if o == nil  {
		return nil, false
	}
	return o.FlashbladeTargetParams.Get(), o.FlashbladeTargetParams.IsSet()
}

// HasFlashbladeTargetParams returns a boolean if a field has been set.
func (o *RecoverElastifileFilesParams) HasFlashbladeTargetParams() bool {
	if o != nil && o.FlashbladeTargetParams.IsSet() {
		return true
	}

	return false
}

// SetFlashbladeTargetParams gets a reference to the given NullableRecoverOtherNasToFlashbladeFilesTargetParams and assigns it to the FlashbladeTargetParams field.
func (o *RecoverElastifileFilesParams) SetFlashbladeTargetParams(v RecoverOtherNasToFlashbladeFilesTargetParams) {
	o.FlashbladeTargetParams.Set(&v)
}
// SetFlashbladeTargetParamsNil sets the value for FlashbladeTargetParams to be an explicit nil
func (o *RecoverElastifileFilesParams) SetFlashbladeTargetParamsNil() {
	o.FlashbladeTargetParams.Set(nil)
}

// UnsetFlashbladeTargetParams ensures that no value is present for FlashbladeTargetParams, not even an explicit nil
func (o *RecoverElastifileFilesParams) UnsetFlashbladeTargetParams() {
	o.FlashbladeTargetParams.Unset()
}

// GetGenericNasTargetParams returns the GenericNasTargetParams field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RecoverElastifileFilesParams) GetGenericNasTargetParams() RecoverOtherNasToGenericNasFilesTargetParams {
	if o == nil || o.GenericNasTargetParams.Get() == nil {
		var ret RecoverOtherNasToGenericNasFilesTargetParams
		return ret
	}
	return *o.GenericNasTargetParams.Get()
}

// GetGenericNasTargetParamsOk returns a tuple with the GenericNasTargetParams field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RecoverElastifileFilesParams) GetGenericNasTargetParamsOk() (*RecoverOtherNasToGenericNasFilesTargetParams, bool) {
	if o == nil  {
		return nil, false
	}
	return o.GenericNasTargetParams.Get(), o.GenericNasTargetParams.IsSet()
}

// HasGenericNasTargetParams returns a boolean if a field has been set.
func (o *RecoverElastifileFilesParams) HasGenericNasTargetParams() bool {
	if o != nil && o.GenericNasTargetParams.IsSet() {
		return true
	}

	return false
}

// SetGenericNasTargetParams gets a reference to the given NullableRecoverOtherNasToGenericNasFilesTargetParams and assigns it to the GenericNasTargetParams field.
func (o *RecoverElastifileFilesParams) SetGenericNasTargetParams(v RecoverOtherNasToGenericNasFilesTargetParams) {
	o.GenericNasTargetParams.Set(&v)
}
// SetGenericNasTargetParamsNil sets the value for GenericNasTargetParams to be an explicit nil
func (o *RecoverElastifileFilesParams) SetGenericNasTargetParamsNil() {
	o.GenericNasTargetParams.Set(nil)
}

// UnsetGenericNasTargetParams ensures that no value is present for GenericNasTargetParams, not even an explicit nil
func (o *RecoverElastifileFilesParams) UnsetGenericNasTargetParams() {
	o.GenericNasTargetParams.Unset()
}

// GetGpfsTargetParams returns the GpfsTargetParams field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RecoverElastifileFilesParams) GetGpfsTargetParams() RecoverOtherNasToGpfsFilesTargetParams {
	if o == nil || o.GpfsTargetParams.Get() == nil {
		var ret RecoverOtherNasToGpfsFilesTargetParams
		return ret
	}
	return *o.GpfsTargetParams.Get()
}

// GetGpfsTargetParamsOk returns a tuple with the GpfsTargetParams field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RecoverElastifileFilesParams) GetGpfsTargetParamsOk() (*RecoverOtherNasToGpfsFilesTargetParams, bool) {
	if o == nil  {
		return nil, false
	}
	return o.GpfsTargetParams.Get(), o.GpfsTargetParams.IsSet()
}

// HasGpfsTargetParams returns a boolean if a field has been set.
func (o *RecoverElastifileFilesParams) HasGpfsTargetParams() bool {
	if o != nil && o.GpfsTargetParams.IsSet() {
		return true
	}

	return false
}

// SetGpfsTargetParams gets a reference to the given NullableRecoverOtherNasToGpfsFilesTargetParams and assigns it to the GpfsTargetParams field.
func (o *RecoverElastifileFilesParams) SetGpfsTargetParams(v RecoverOtherNasToGpfsFilesTargetParams) {
	o.GpfsTargetParams.Set(&v)
}
// SetGpfsTargetParamsNil sets the value for GpfsTargetParams to be an explicit nil
func (o *RecoverElastifileFilesParams) SetGpfsTargetParamsNil() {
	o.GpfsTargetParams.Set(nil)
}

// UnsetGpfsTargetParams ensures that no value is present for GpfsTargetParams, not even an explicit nil
func (o *RecoverElastifileFilesParams) UnsetGpfsTargetParams() {
	o.GpfsTargetParams.Unset()
}

// GetIsilonTargetParams returns the IsilonTargetParams field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RecoverElastifileFilesParams) GetIsilonTargetParams() RecoverOtherNasToIsilonFilesTargetParams {
	if o == nil || o.IsilonTargetParams.Get() == nil {
		var ret RecoverOtherNasToIsilonFilesTargetParams
		return ret
	}
	return *o.IsilonTargetParams.Get()
}

// GetIsilonTargetParamsOk returns a tuple with the IsilonTargetParams field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RecoverElastifileFilesParams) GetIsilonTargetParamsOk() (*RecoverOtherNasToIsilonFilesTargetParams, bool) {
	if o == nil  {
		return nil, false
	}
	return o.IsilonTargetParams.Get(), o.IsilonTargetParams.IsSet()
}

// HasIsilonTargetParams returns a boolean if a field has been set.
func (o *RecoverElastifileFilesParams) HasIsilonTargetParams() bool {
	if o != nil && o.IsilonTargetParams.IsSet() {
		return true
	}

	return false
}

// SetIsilonTargetParams gets a reference to the given NullableRecoverOtherNasToIsilonFilesTargetParams and assigns it to the IsilonTargetParams field.
func (o *RecoverElastifileFilesParams) SetIsilonTargetParams(v RecoverOtherNasToIsilonFilesTargetParams) {
	o.IsilonTargetParams.Set(&v)
}
// SetIsilonTargetParamsNil sets the value for IsilonTargetParams to be an explicit nil
func (o *RecoverElastifileFilesParams) SetIsilonTargetParamsNil() {
	o.IsilonTargetParams.Set(nil)
}

// UnsetIsilonTargetParams ensures that no value is present for IsilonTargetParams, not even an explicit nil
func (o *RecoverElastifileFilesParams) UnsetIsilonTargetParams() {
	o.IsilonTargetParams.Unset()
}

// GetNetappTargetParams returns the NetappTargetParams field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RecoverElastifileFilesParams) GetNetappTargetParams() RecoverOtherNasToNetappFilesTargetParams {
	if o == nil || o.NetappTargetParams.Get() == nil {
		var ret RecoverOtherNasToNetappFilesTargetParams
		return ret
	}
	return *o.NetappTargetParams.Get()
}

// GetNetappTargetParamsOk returns a tuple with the NetappTargetParams field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RecoverElastifileFilesParams) GetNetappTargetParamsOk() (*RecoverOtherNasToNetappFilesTargetParams, bool) {
	if o == nil  {
		return nil, false
	}
	return o.NetappTargetParams.Get(), o.NetappTargetParams.IsSet()
}

// HasNetappTargetParams returns a boolean if a field has been set.
func (o *RecoverElastifileFilesParams) HasNetappTargetParams() bool {
	if o != nil && o.NetappTargetParams.IsSet() {
		return true
	}

	return false
}

// SetNetappTargetParams gets a reference to the given NullableRecoverOtherNasToNetappFilesTargetParams and assigns it to the NetappTargetParams field.
func (o *RecoverElastifileFilesParams) SetNetappTargetParams(v RecoverOtherNasToNetappFilesTargetParams) {
	o.NetappTargetParams.Set(&v)
}
// SetNetappTargetParamsNil sets the value for NetappTargetParams to be an explicit nil
func (o *RecoverElastifileFilesParams) SetNetappTargetParamsNil() {
	o.NetappTargetParams.Set(nil)
}

// UnsetNetappTargetParams ensures that no value is present for NetappTargetParams, not even an explicit nil
func (o *RecoverElastifileFilesParams) UnsetNetappTargetParams() {
	o.NetappTargetParams.Unset()
}

func (o RecoverElastifileFilesParams) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.FilesAndFolders != nil {
		toSerialize["filesAndFolders"] = o.FilesAndFolders
	}
	if true {
		toSerialize["targetEnvironment"] = o.TargetEnvironment
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
	if o.NetappTargetParams.IsSet() {
		toSerialize["netappTargetParams"] = o.NetappTargetParams.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableRecoverElastifileFilesParams struct {
	value *RecoverElastifileFilesParams
	isSet bool
}

func (v NullableRecoverElastifileFilesParams) Get() *RecoverElastifileFilesParams {
	return v.value
}

func (v *NullableRecoverElastifileFilesParams) Set(val *RecoverElastifileFilesParams) {
	v.value = val
	v.isSet = true
}

func (v NullableRecoverElastifileFilesParams) IsSet() bool {
	return v.isSet
}

func (v *NullableRecoverElastifileFilesParams) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRecoverElastifileFilesParams(val *RecoverElastifileFilesParams) *NullableRecoverElastifileFilesParams {
	return &NullableRecoverElastifileFilesParams{value: val, isSet: true}
}

func (v NullableRecoverElastifileFilesParams) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRecoverElastifileFilesParams) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}




func (o RecoverElastifileFilesParams) Print() {
		if byteArray, err := o.MarshalJSON(); err != nil {
				fmt.Println(err)
		} else {
				fmt.Println(string(byteArray))
		}
}