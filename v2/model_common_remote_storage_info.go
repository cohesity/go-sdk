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

// checks if the CommonRemoteStorageInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CommonRemoteStorageInfo{}

// CommonRemoteStorageInfo Specifies the details of common remote storage info.
type CommonRemoteStorageInfo struct {
	// Specifies unique id of the registered remote storage.
	Id NullableInt64 `json:"id,omitempty"`
	// Specifies the product type of the remote storage.
	Product NullableString `json:"product"`
}

type _CommonRemoteStorageInfo CommonRemoteStorageInfo

// NewCommonRemoteStorageInfo instantiates a new CommonRemoteStorageInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCommonRemoteStorageInfo(product NullableString) *CommonRemoteStorageInfo {
	this := CommonRemoteStorageInfo{}
	this.Product = product
	return &this
}

// NewCommonRemoteStorageInfoWithDefaults instantiates a new CommonRemoteStorageInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCommonRemoteStorageInfoWithDefaults() *CommonRemoteStorageInfo {
	this := CommonRemoteStorageInfo{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CommonRemoteStorageInfo) GetId() int64 {
	if o == nil || IsNil(o.Id.Get()) {
		var ret int64
		return ret
	}
	return *o.Id.Get()
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CommonRemoteStorageInfo) GetIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.Id.Get(), o.Id.IsSet()
}

// HasId returns a boolean if a field has been set.
func (o *CommonRemoteStorageInfo) HasId() bool {
	if o != nil && o.Id.IsSet() {
		return true
	}

	return false
}

// SetId gets a reference to the given NullableInt64 and assigns it to the Id field.
func (o *CommonRemoteStorageInfo) SetId(v int64) {
	o.Id.Set(&v)
}
// SetIdNil sets the value for Id to be an explicit nil
func (o *CommonRemoteStorageInfo) SetIdNil() {
	o.Id.Set(nil)
}

// UnsetId ensures that no value is present for Id, not even an explicit nil
func (o *CommonRemoteStorageInfo) UnsetId() {
	o.Id.Unset()
}

// GetProduct returns the Product field value
// If the value is explicit nil, the zero value for string will be returned
func (o *CommonRemoteStorageInfo) GetProduct() string {
	if o == nil || o.Product.Get() == nil {
		var ret string
		return ret
	}

	return *o.Product.Get()
}

// GetProductOk returns a tuple with the Product field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CommonRemoteStorageInfo) GetProductOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Product.Get(), o.Product.IsSet()
}

// SetProduct sets field value
func (o *CommonRemoteStorageInfo) SetProduct(v string) {
	o.Product.Set(&v)
}

func (o CommonRemoteStorageInfo) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CommonRemoteStorageInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Id.IsSet() {
		toSerialize["id"] = o.Id.Get()
	}
	toSerialize["product"] = o.Product.Get()
	return toSerialize, nil
}

func (o *CommonRemoteStorageInfo) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"product",
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

	varCommonRemoteStorageInfo := _CommonRemoteStorageInfo{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCommonRemoteStorageInfo)

	if err != nil {
		return err
	}

	*o = CommonRemoteStorageInfo(varCommonRemoteStorageInfo)

	return err
}

type NullableCommonRemoteStorageInfo struct {
	value *CommonRemoteStorageInfo
	isSet bool
}

func (v NullableCommonRemoteStorageInfo) Get() *CommonRemoteStorageInfo {
	return v.value
}

func (v *NullableCommonRemoteStorageInfo) Set(val *CommonRemoteStorageInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableCommonRemoteStorageInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableCommonRemoteStorageInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCommonRemoteStorageInfo(val *CommonRemoteStorageInfo) *NullableCommonRemoteStorageInfo {
	return &NullableCommonRemoteStorageInfo{value: val, isSet: true}
}

func (v NullableCommonRemoteStorageInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCommonRemoteStorageInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


