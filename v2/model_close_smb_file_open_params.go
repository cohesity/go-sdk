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

// checks if the CloseSmbFileOpenParams type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CloseSmbFileOpenParams{}

// CloseSmbFileOpenParams Specifies params to close active SMB file open.
type CloseSmbFileOpenParams struct {
	// Specifies the filepath in the Cohesity View relative to the root filesystem. If this field is specified, viewName field must also be specified.
	FilePath NullableString `json:"filePath"`
	// Specifies the id of the active open.
	OpenId NullableInt64 `json:"openId"`
	// Specifies the name of the Cohesity View in which to search. If a view name is not specified, all the views in the Cluster are searched. This field is mandatory if filePath field is specified.
	ViewName NullableString `json:"viewName"`
}

type _CloseSmbFileOpenParams CloseSmbFileOpenParams

// NewCloseSmbFileOpenParams instantiates a new CloseSmbFileOpenParams object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCloseSmbFileOpenParams(filePath NullableString, openId NullableInt64, viewName NullableString) *CloseSmbFileOpenParams {
	this := CloseSmbFileOpenParams{}
	this.FilePath = filePath
	this.OpenId = openId
	this.ViewName = viewName
	return &this
}

// NewCloseSmbFileOpenParamsWithDefaults instantiates a new CloseSmbFileOpenParams object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCloseSmbFileOpenParamsWithDefaults() *CloseSmbFileOpenParams {
	this := CloseSmbFileOpenParams{}
	return &this
}

// GetFilePath returns the FilePath field value
// If the value is explicit nil, the zero value for string will be returned
func (o *CloseSmbFileOpenParams) GetFilePath() string {
	if o == nil || o.FilePath.Get() == nil {
		var ret string
		return ret
	}

	return *o.FilePath.Get()
}

// GetFilePathOk returns a tuple with the FilePath field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CloseSmbFileOpenParams) GetFilePathOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.FilePath.Get(), o.FilePath.IsSet()
}

// SetFilePath sets field value
func (o *CloseSmbFileOpenParams) SetFilePath(v string) {
	o.FilePath.Set(&v)
}

// GetOpenId returns the OpenId field value
// If the value is explicit nil, the zero value for int64 will be returned
func (o *CloseSmbFileOpenParams) GetOpenId() int64 {
	if o == nil || o.OpenId.Get() == nil {
		var ret int64
		return ret
	}

	return *o.OpenId.Get()
}

// GetOpenIdOk returns a tuple with the OpenId field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CloseSmbFileOpenParams) GetOpenIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.OpenId.Get(), o.OpenId.IsSet()
}

// SetOpenId sets field value
func (o *CloseSmbFileOpenParams) SetOpenId(v int64) {
	o.OpenId.Set(&v)
}

// GetViewName returns the ViewName field value
// If the value is explicit nil, the zero value for string will be returned
func (o *CloseSmbFileOpenParams) GetViewName() string {
	if o == nil || o.ViewName.Get() == nil {
		var ret string
		return ret
	}

	return *o.ViewName.Get()
}

// GetViewNameOk returns a tuple with the ViewName field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CloseSmbFileOpenParams) GetViewNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ViewName.Get(), o.ViewName.IsSet()
}

// SetViewName sets field value
func (o *CloseSmbFileOpenParams) SetViewName(v string) {
	o.ViewName.Set(&v)
}

func (o CloseSmbFileOpenParams) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CloseSmbFileOpenParams) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["filePath"] = o.FilePath.Get()
	toSerialize["openId"] = o.OpenId.Get()
	toSerialize["viewName"] = o.ViewName.Get()
	return toSerialize, nil
}

func (o *CloseSmbFileOpenParams) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"filePath",
		"openId",
		"viewName",
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

	varCloseSmbFileOpenParams := _CloseSmbFileOpenParams{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCloseSmbFileOpenParams)

	if err != nil {
		return err
	}

	*o = CloseSmbFileOpenParams(varCloseSmbFileOpenParams)

	return err
}

type NullableCloseSmbFileOpenParams struct {
	value *CloseSmbFileOpenParams
	isSet bool
}

func (v NullableCloseSmbFileOpenParams) Get() *CloseSmbFileOpenParams {
	return v.value
}

func (v *NullableCloseSmbFileOpenParams) Set(val *CloseSmbFileOpenParams) {
	v.value = val
	v.isSet = true
}

func (v NullableCloseSmbFileOpenParams) IsSet() bool {
	return v.isSet
}

func (v *NullableCloseSmbFileOpenParams) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCloseSmbFileOpenParams(val *CloseSmbFileOpenParams) *NullableCloseSmbFileOpenParams {
	return &NullableCloseSmbFileOpenParams{value: val, isSet: true}
}

func (v NullableCloseSmbFileOpenParams) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCloseSmbFileOpenParams) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


