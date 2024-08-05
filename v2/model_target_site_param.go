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

// checks if the TargetSiteParam type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TargetSiteParam{}

// TargetSiteParam Specifies the target Site to recover to.
type TargetSiteParam struct {
	// Specifies the id of the object.
	Id NullableInt64 `json:"id"`
	// Specifies the name of the object.
	Name NullableString `json:"name,omitempty"`
	// Specifies the name for the target doc lib.
	TargetDocLibName NullableString `json:"targetDocLibName,omitempty"`
	// Specifies the prefix for the target doc lib.
	TargetDocLibPrefix NullableString `json:"targetDocLibPrefix,omitempty"`
}

type _TargetSiteParam TargetSiteParam

// NewTargetSiteParam instantiates a new TargetSiteParam object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTargetSiteParam(id NullableInt64) *TargetSiteParam {
	this := TargetSiteParam{}
	this.Id = id
	return &this
}

// NewTargetSiteParamWithDefaults instantiates a new TargetSiteParam object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTargetSiteParamWithDefaults() *TargetSiteParam {
	this := TargetSiteParam{}
	return &this
}

// GetId returns the Id field value
// If the value is explicit nil, the zero value for int64 will be returned
func (o *TargetSiteParam) GetId() int64 {
	if o == nil || o.Id.Get() == nil {
		var ret int64
		return ret
	}

	return *o.Id.Get()
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TargetSiteParam) GetIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.Id.Get(), o.Id.IsSet()
}

// SetId sets field value
func (o *TargetSiteParam) SetId(v int64) {
	o.Id.Set(&v)
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TargetSiteParam) GetName() string {
	if o == nil || IsNil(o.Name.Get()) {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TargetSiteParam) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *TargetSiteParam) HasName() bool {
	if o != nil && o.Name.IsSet() {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableString and assigns it to the Name field.
func (o *TargetSiteParam) SetName(v string) {
	o.Name.Set(&v)
}
// SetNameNil sets the value for Name to be an explicit nil
func (o *TargetSiteParam) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *TargetSiteParam) UnsetName() {
	o.Name.Unset()
}

// GetTargetDocLibName returns the TargetDocLibName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TargetSiteParam) GetTargetDocLibName() string {
	if o == nil || IsNil(o.TargetDocLibName.Get()) {
		var ret string
		return ret
	}
	return *o.TargetDocLibName.Get()
}

// GetTargetDocLibNameOk returns a tuple with the TargetDocLibName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TargetSiteParam) GetTargetDocLibNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.TargetDocLibName.Get(), o.TargetDocLibName.IsSet()
}

// HasTargetDocLibName returns a boolean if a field has been set.
func (o *TargetSiteParam) HasTargetDocLibName() bool {
	if o != nil && o.TargetDocLibName.IsSet() {
		return true
	}

	return false
}

// SetTargetDocLibName gets a reference to the given NullableString and assigns it to the TargetDocLibName field.
func (o *TargetSiteParam) SetTargetDocLibName(v string) {
	o.TargetDocLibName.Set(&v)
}
// SetTargetDocLibNameNil sets the value for TargetDocLibName to be an explicit nil
func (o *TargetSiteParam) SetTargetDocLibNameNil() {
	o.TargetDocLibName.Set(nil)
}

// UnsetTargetDocLibName ensures that no value is present for TargetDocLibName, not even an explicit nil
func (o *TargetSiteParam) UnsetTargetDocLibName() {
	o.TargetDocLibName.Unset()
}

// GetTargetDocLibPrefix returns the TargetDocLibPrefix field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TargetSiteParam) GetTargetDocLibPrefix() string {
	if o == nil || IsNil(o.TargetDocLibPrefix.Get()) {
		var ret string
		return ret
	}
	return *o.TargetDocLibPrefix.Get()
}

// GetTargetDocLibPrefixOk returns a tuple with the TargetDocLibPrefix field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TargetSiteParam) GetTargetDocLibPrefixOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.TargetDocLibPrefix.Get(), o.TargetDocLibPrefix.IsSet()
}

// HasTargetDocLibPrefix returns a boolean if a field has been set.
func (o *TargetSiteParam) HasTargetDocLibPrefix() bool {
	if o != nil && o.TargetDocLibPrefix.IsSet() {
		return true
	}

	return false
}

// SetTargetDocLibPrefix gets a reference to the given NullableString and assigns it to the TargetDocLibPrefix field.
func (o *TargetSiteParam) SetTargetDocLibPrefix(v string) {
	o.TargetDocLibPrefix.Set(&v)
}
// SetTargetDocLibPrefixNil sets the value for TargetDocLibPrefix to be an explicit nil
func (o *TargetSiteParam) SetTargetDocLibPrefixNil() {
	o.TargetDocLibPrefix.Set(nil)
}

// UnsetTargetDocLibPrefix ensures that no value is present for TargetDocLibPrefix, not even an explicit nil
func (o *TargetSiteParam) UnsetTargetDocLibPrefix() {
	o.TargetDocLibPrefix.Unset()
}

func (o TargetSiteParam) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TargetSiteParam) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id.Get()
	if o.Name.IsSet() {
		toSerialize["name"] = o.Name.Get()
	}
	if o.TargetDocLibName.IsSet() {
		toSerialize["targetDocLibName"] = o.TargetDocLibName.Get()
	}
	if o.TargetDocLibPrefix.IsSet() {
		toSerialize["targetDocLibPrefix"] = o.TargetDocLibPrefix.Get()
	}
	return toSerialize, nil
}

func (o *TargetSiteParam) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
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

	varTargetSiteParam := _TargetSiteParam{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varTargetSiteParam)

	if err != nil {
		return err
	}

	*o = TargetSiteParam(varTargetSiteParam)

	return err
}

type NullableTargetSiteParam struct {
	value *TargetSiteParam
	isSet bool
}

func (v NullableTargetSiteParam) Get() *TargetSiteParam {
	return v.value
}

func (v *NullableTargetSiteParam) Set(val *TargetSiteParam) {
	v.value = val
	v.isSet = true
}

func (v NullableTargetSiteParam) IsSet() bool {
	return v.isSet
}

func (v *NullableTargetSiteParam) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTargetSiteParam(val *TargetSiteParam) *NullableTargetSiteParam {
	return &NullableTargetSiteParam{value: val, isSet: true}
}

func (v NullableTargetSiteParam) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTargetSiteParam) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


