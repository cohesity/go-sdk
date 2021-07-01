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

// ChangeServiceStateResult Specifies the result returned after a successful request to change the state of services running on the Cluster.
type ChangeServiceStateResult struct {
	// Specifies a description of the result of the operation.
	Message NullableString `json:"message,omitempty"`
	// Specifies a URL which can be queried to check the status of the operation.
	StatusUrl NullableString `json:"statusUrl,omitempty"`
}

// NewChangeServiceStateResult instantiates a new ChangeServiceStateResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewChangeServiceStateResult() *ChangeServiceStateResult {
	this := ChangeServiceStateResult{}
	return &this
}

// NewChangeServiceStateResultWithDefaults instantiates a new ChangeServiceStateResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewChangeServiceStateResultWithDefaults() *ChangeServiceStateResult {
	this := ChangeServiceStateResult{}
	return &this
}

// GetMessage returns the Message field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ChangeServiceStateResult) GetMessage() string {
	if o == nil || o.Message.Get() == nil {
		var ret string
		return ret
	}
	return *o.Message.Get()
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ChangeServiceStateResult) GetMessageOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Message.Get(), o.Message.IsSet()
}

// HasMessage returns a boolean if a field has been set.
func (o *ChangeServiceStateResult) HasMessage() bool {
	if o != nil && o.Message.IsSet() {
		return true
	}

	return false
}

// SetMessage gets a reference to the given NullableString and assigns it to the Message field.
func (o *ChangeServiceStateResult) SetMessage(v string) {
	o.Message.Set(&v)
}
// SetMessageNil sets the value for Message to be an explicit nil
func (o *ChangeServiceStateResult) SetMessageNil() {
	o.Message.Set(nil)
}

// UnsetMessage ensures that no value is present for Message, not even an explicit nil
func (o *ChangeServiceStateResult) UnsetMessage() {
	o.Message.Unset()
}

// GetStatusUrl returns the StatusUrl field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ChangeServiceStateResult) GetStatusUrl() string {
	if o == nil || o.StatusUrl.Get() == nil {
		var ret string
		return ret
	}
	return *o.StatusUrl.Get()
}

// GetStatusUrlOk returns a tuple with the StatusUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ChangeServiceStateResult) GetStatusUrlOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.StatusUrl.Get(), o.StatusUrl.IsSet()
}

// HasStatusUrl returns a boolean if a field has been set.
func (o *ChangeServiceStateResult) HasStatusUrl() bool {
	if o != nil && o.StatusUrl.IsSet() {
		return true
	}

	return false
}

// SetStatusUrl gets a reference to the given NullableString and assigns it to the StatusUrl field.
func (o *ChangeServiceStateResult) SetStatusUrl(v string) {
	o.StatusUrl.Set(&v)
}
// SetStatusUrlNil sets the value for StatusUrl to be an explicit nil
func (o *ChangeServiceStateResult) SetStatusUrlNil() {
	o.StatusUrl.Set(nil)
}

// UnsetStatusUrl ensures that no value is present for StatusUrl, not even an explicit nil
func (o *ChangeServiceStateResult) UnsetStatusUrl() {
	o.StatusUrl.Unset()
}

func (o ChangeServiceStateResult) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Message.IsSet() {
		toSerialize["message"] = o.Message.Get()
	}
	if o.StatusUrl.IsSet() {
		toSerialize["statusUrl"] = o.StatusUrl.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableChangeServiceStateResult struct {
	value *ChangeServiceStateResult
	isSet bool
}

func (v NullableChangeServiceStateResult) Get() *ChangeServiceStateResult {
	return v.value
}

func (v *NullableChangeServiceStateResult) Set(val *ChangeServiceStateResult) {
	v.value = val
	v.isSet = true
}

func (v NullableChangeServiceStateResult) IsSet() bool {
	return v.isSet
}

func (v *NullableChangeServiceStateResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableChangeServiceStateResult(val *ChangeServiceStateResult) *NullableChangeServiceStateResult {
	return &NullableChangeServiceStateResult{value: val, isSet: true}
}

func (v NullableChangeServiceStateResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableChangeServiceStateResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


