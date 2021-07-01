/*
 * Cohesity REST API
 *
 * This API list provides operations for interfacing with the Cohesity Cluster.
 *
 * API version: 1.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package cohesitysdk

import (
	"encoding/json"
)

// AppInstanceIdParameter struct for AppInstanceIdParameter
type AppInstanceIdParameter struct {
	// Specifies the app instance Id. In: path
	AppInstanceId NullableInt64 `json:"appInstanceId"`
}

// NewAppInstanceIdParameter instantiates a new AppInstanceIdParameter object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAppInstanceIdParameter(appInstanceId NullableInt64) *AppInstanceIdParameter {
	this := AppInstanceIdParameter{}
	this.AppInstanceId = appInstanceId
	return &this
}

// NewAppInstanceIdParameterWithDefaults instantiates a new AppInstanceIdParameter object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAppInstanceIdParameterWithDefaults() *AppInstanceIdParameter {
	this := AppInstanceIdParameter{}
	return &this
}

// GetAppInstanceId returns the AppInstanceId field value
// If the value is explicit nil, the zero value for int64 will be returned
func (o *AppInstanceIdParameter) GetAppInstanceId() int64 {
	if o == nil || o.AppInstanceId.Get() == nil {
		var ret int64
		return ret
	}

	return *o.AppInstanceId.Get()
}

// GetAppInstanceIdOk returns a tuple with the AppInstanceId field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AppInstanceIdParameter) GetAppInstanceIdOk() (*int64, bool) {
	if o == nil  {
		return nil, false
	}
	return o.AppInstanceId.Get(), o.AppInstanceId.IsSet()
}

// SetAppInstanceId sets field value
func (o *AppInstanceIdParameter) SetAppInstanceId(v int64) {
	o.AppInstanceId.Set(&v)
}

func (o AppInstanceIdParameter) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["appInstanceId"] = o.AppInstanceId.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableAppInstanceIdParameter struct {
	value *AppInstanceIdParameter
	isSet bool
}

func (v NullableAppInstanceIdParameter) Get() *AppInstanceIdParameter {
	return v.value
}

func (v *NullableAppInstanceIdParameter) Set(val *AppInstanceIdParameter) {
	v.value = val
	v.isSet = true
}

func (v NullableAppInstanceIdParameter) IsSet() bool {
	return v.isSet
}

func (v *NullableAppInstanceIdParameter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAppInstanceIdParameter(val *AppInstanceIdParameter) *NullableAppInstanceIdParameter {
	return &NullableAppInstanceIdParameter{value: val, isSet: true}
}

func (v NullableAppInstanceIdParameter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAppInstanceIdParameter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


