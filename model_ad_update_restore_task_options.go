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

// ADUpdateRestoreTaskOptions struct for ADUpdateRestoreTaskOptions
type ADUpdateRestoreTaskOptions struct {
	ObjectAttributesParam *ADAttributeRestoreParam `json:"objectAttributesParam,omitempty"`
	ObjectParam *ADObjectRestoreParam `json:"objectParam,omitempty"`
	// Specifies the AD restore request type.
	Type NullableInt32 `json:"type,omitempty"`
}

// NewADUpdateRestoreTaskOptions instantiates a new ADUpdateRestoreTaskOptions object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewADUpdateRestoreTaskOptions() *ADUpdateRestoreTaskOptions {
	this := ADUpdateRestoreTaskOptions{}
	return &this
}

// NewADUpdateRestoreTaskOptionsWithDefaults instantiates a new ADUpdateRestoreTaskOptions object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewADUpdateRestoreTaskOptionsWithDefaults() *ADUpdateRestoreTaskOptions {
	this := ADUpdateRestoreTaskOptions{}
	return &this
}

// GetObjectAttributesParam returns the ObjectAttributesParam field value if set, zero value otherwise.
func (o *ADUpdateRestoreTaskOptions) GetObjectAttributesParam() ADAttributeRestoreParam {
	if o == nil || o.ObjectAttributesParam == nil {
		var ret ADAttributeRestoreParam
		return ret
	}
	return *o.ObjectAttributesParam
}

// GetObjectAttributesParamOk returns a tuple with the ObjectAttributesParam field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ADUpdateRestoreTaskOptions) GetObjectAttributesParamOk() (*ADAttributeRestoreParam, bool) {
	if o == nil || o.ObjectAttributesParam == nil {
		return nil, false
	}
	return o.ObjectAttributesParam, true
}

// HasObjectAttributesParam returns a boolean if a field has been set.
func (o *ADUpdateRestoreTaskOptions) HasObjectAttributesParam() bool {
	if o != nil && o.ObjectAttributesParam != nil {
		return true
	}

	return false
}

// SetObjectAttributesParam gets a reference to the given ADAttributeRestoreParam and assigns it to the ObjectAttributesParam field.
func (o *ADUpdateRestoreTaskOptions) SetObjectAttributesParam(v ADAttributeRestoreParam) {
	o.ObjectAttributesParam = &v
}

// GetObjectParam returns the ObjectParam field value if set, zero value otherwise.
func (o *ADUpdateRestoreTaskOptions) GetObjectParam() ADObjectRestoreParam {
	if o == nil || o.ObjectParam == nil {
		var ret ADObjectRestoreParam
		return ret
	}
	return *o.ObjectParam
}

// GetObjectParamOk returns a tuple with the ObjectParam field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ADUpdateRestoreTaskOptions) GetObjectParamOk() (*ADObjectRestoreParam, bool) {
	if o == nil || o.ObjectParam == nil {
		return nil, false
	}
	return o.ObjectParam, true
}

// HasObjectParam returns a boolean if a field has been set.
func (o *ADUpdateRestoreTaskOptions) HasObjectParam() bool {
	if o != nil && o.ObjectParam != nil {
		return true
	}

	return false
}

// SetObjectParam gets a reference to the given ADObjectRestoreParam and assigns it to the ObjectParam field.
func (o *ADUpdateRestoreTaskOptions) SetObjectParam(v ADObjectRestoreParam) {
	o.ObjectParam = &v
}

// GetType returns the Type field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ADUpdateRestoreTaskOptions) GetType() int32 {
	if o == nil || o.Type.Get() == nil {
		var ret int32
		return ret
	}
	return *o.Type.Get()
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ADUpdateRestoreTaskOptions) GetTypeOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Type.Get(), o.Type.IsSet()
}

// HasType returns a boolean if a field has been set.
func (o *ADUpdateRestoreTaskOptions) HasType() bool {
	if o != nil && o.Type.IsSet() {
		return true
	}

	return false
}

// SetType gets a reference to the given NullableInt32 and assigns it to the Type field.
func (o *ADUpdateRestoreTaskOptions) SetType(v int32) {
	o.Type.Set(&v)
}
// SetTypeNil sets the value for Type to be an explicit nil
func (o *ADUpdateRestoreTaskOptions) SetTypeNil() {
	o.Type.Set(nil)
}

// UnsetType ensures that no value is present for Type, not even an explicit nil
func (o *ADUpdateRestoreTaskOptions) UnsetType() {
	o.Type.Unset()
}

func (o ADUpdateRestoreTaskOptions) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ObjectAttributesParam != nil {
		toSerialize["objectAttributesParam"] = o.ObjectAttributesParam
	}
	if o.ObjectParam != nil {
		toSerialize["objectParam"] = o.ObjectParam
	}
	if o.Type.IsSet() {
		toSerialize["type"] = o.Type.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableADUpdateRestoreTaskOptions struct {
	value *ADUpdateRestoreTaskOptions
	isSet bool
}

func (v NullableADUpdateRestoreTaskOptions) Get() *ADUpdateRestoreTaskOptions {
	return v.value
}

func (v *NullableADUpdateRestoreTaskOptions) Set(val *ADUpdateRestoreTaskOptions) {
	v.value = val
	v.isSet = true
}

func (v NullableADUpdateRestoreTaskOptions) IsSet() bool {
	return v.isSet
}

func (v *NullableADUpdateRestoreTaskOptions) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableADUpdateRestoreTaskOptions(val *ADUpdateRestoreTaskOptions) *NullableADUpdateRestoreTaskOptions {
	return &NullableADUpdateRestoreTaskOptions{value: val, isSet: true}
}

func (v NullableADUpdateRestoreTaskOptions) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableADUpdateRestoreTaskOptions) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


