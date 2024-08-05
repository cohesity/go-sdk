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

// checks if the SuccessResp type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SuccessResp{}

// SuccessResp Specifies the firewall profile names to be removed.
type SuccessResp struct {
	// Specifies the response message.
	Message NullableString `json:"message"`
}

type _SuccessResp SuccessResp

// NewSuccessResp instantiates a new SuccessResp object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSuccessResp(message NullableString) *SuccessResp {
	this := SuccessResp{}
	this.Message = message
	return &this
}

// NewSuccessRespWithDefaults instantiates a new SuccessResp object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSuccessRespWithDefaults() *SuccessResp {
	this := SuccessResp{}
	return &this
}

// GetMessage returns the Message field value
// If the value is explicit nil, the zero value for string will be returned
func (o *SuccessResp) GetMessage() string {
	if o == nil || o.Message.Get() == nil {
		var ret string
		return ret
	}

	return *o.Message.Get()
}

// GetMessageOk returns a tuple with the Message field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SuccessResp) GetMessageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Message.Get(), o.Message.IsSet()
}

// SetMessage sets field value
func (o *SuccessResp) SetMessage(v string) {
	o.Message.Set(&v)
}

func (o SuccessResp) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SuccessResp) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["message"] = o.Message.Get()
	return toSerialize, nil
}

func (o *SuccessResp) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"message",
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

	varSuccessResp := _SuccessResp{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varSuccessResp)

	if err != nil {
		return err
	}

	*o = SuccessResp(varSuccessResp)

	return err
}

type NullableSuccessResp struct {
	value *SuccessResp
	isSet bool
}

func (v NullableSuccessResp) Get() *SuccessResp {
	return v.value
}

func (v *NullableSuccessResp) Set(val *SuccessResp) {
	v.value = val
	v.isSet = true
}

func (v NullableSuccessResp) IsSet() bool {
	return v.isSet
}

func (v *NullableSuccessResp) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSuccessResp(val *SuccessResp) *NullableSuccessResp {
	return &NullableSuccessResp{value: val, isSet: true}
}

func (v NullableSuccessResp) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSuccessResp) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


