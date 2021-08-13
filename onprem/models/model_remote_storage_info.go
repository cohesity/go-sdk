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

// RemoteStorageInfo Specifies the remote storage Registration parameters.
type RemoteStorageInfo struct {
	// Specifies unique id of the registered remote storage.
	Id NullableInt64 `json:"id,omitempty"`
	// Specifies the product type of the remote storage.
	Product NullableString `json:"product"`
	FlashbladeParams *FlashbladeParams `json:"flashbladeParams,omitempty"`
}

// NewRemoteStorageInfo instantiates a new RemoteStorageInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRemoteStorageInfo(product NullableString) *RemoteStorageInfo {
	this := RemoteStorageInfo{}
	this.Product = product
	return &this
}

// NewRemoteStorageInfoWithDefaults instantiates a new RemoteStorageInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRemoteStorageInfoWithDefaults() *RemoteStorageInfo {
	this := RemoteStorageInfo{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RemoteStorageInfo) GetId() int64 {
	if o == nil || o.Id.Get() == nil {
		var ret int64
		return ret
	}
	return *o.Id.Get()
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RemoteStorageInfo) GetIdOk() (*int64, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Id.Get(), o.Id.IsSet()
}

// HasId returns a boolean if a field has been set.
func (o *RemoteStorageInfo) HasId() bool {
	if o != nil && o.Id.IsSet() {
		return true
	}

	return false
}

// SetId gets a reference to the given NullableInt64 and assigns it to the Id field.
func (o *RemoteStorageInfo) SetId(v int64) {
	o.Id.Set(&v)
}
// SetIdNil sets the value for Id to be an explicit nil
func (o *RemoteStorageInfo) SetIdNil() {
	o.Id.Set(nil)
}

// UnsetId ensures that no value is present for Id, not even an explicit nil
func (o *RemoteStorageInfo) UnsetId() {
	o.Id.Unset()
}

// GetProduct returns the Product field value
// If the value is explicit nil, the zero value for string will be returned
func (o *RemoteStorageInfo) GetProduct() string {
	if o == nil || o.Product.Get() == nil {
		var ret string
		return ret
	}

	return *o.Product.Get()
}

// GetProductOk returns a tuple with the Product field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RemoteStorageInfo) GetProductOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Product.Get(), o.Product.IsSet()
}

// SetProduct sets field value
func (o *RemoteStorageInfo) SetProduct(v string) {
	o.Product.Set(&v)
}

// GetFlashbladeParams returns the FlashbladeParams field value if set, zero value otherwise.
func (o *RemoteStorageInfo) GetFlashbladeParams() FlashbladeParams {
	if o == nil || o.FlashbladeParams == nil {
		var ret FlashbladeParams
		return ret
	}
	return *o.FlashbladeParams
}

// GetFlashbladeParamsOk returns a tuple with the FlashbladeParams field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RemoteStorageInfo) GetFlashbladeParamsOk() (*FlashbladeParams, bool) {
	if o == nil || o.FlashbladeParams == nil {
		return nil, false
	}
	return o.FlashbladeParams, true
}

// HasFlashbladeParams returns a boolean if a field has been set.
func (o *RemoteStorageInfo) HasFlashbladeParams() bool {
	if o != nil && o.FlashbladeParams != nil {
		return true
	}

	return false
}

// SetFlashbladeParams gets a reference to the given FlashbladeParams and assigns it to the FlashbladeParams field.
func (o *RemoteStorageInfo) SetFlashbladeParams(v FlashbladeParams) {
	o.FlashbladeParams = &v
}

func (o RemoteStorageInfo) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id.IsSet() {
		toSerialize["id"] = o.Id.Get()
	}
	if true {
		toSerialize["product"] = o.Product.Get()
	}
	if o.FlashbladeParams != nil {
		toSerialize["flashbladeParams"] = o.FlashbladeParams
	}
	return json.Marshal(toSerialize)
}

type NullableRemoteStorageInfo struct {
	value *RemoteStorageInfo
	isSet bool
}

func (v NullableRemoteStorageInfo) Get() *RemoteStorageInfo {
	return v.value
}

func (v *NullableRemoteStorageInfo) Set(val *RemoteStorageInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableRemoteStorageInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableRemoteStorageInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRemoteStorageInfo(val *RemoteStorageInfo) *NullableRemoteStorageInfo {
	return &NullableRemoteStorageInfo{value: val, isSet: true}
}

func (v NullableRemoteStorageInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRemoteStorageInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}




func (o RemoteStorageInfo) Print() {
		if byteArray, err := o.MarshalJSON(); err != nil {
				fmt.Println(err)
		} else {
				fmt.Println(string(byteArray))
		}
}