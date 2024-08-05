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

// checks if the FilerLifecycleAgingPolicy type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FilerLifecycleAgingPolicy{}

// FilerLifecycleAgingPolicy Specifies the aging policy. Note: Both the fields days and dateInUsecs are mutually exclusive to each other.
type FilerLifecycleAgingPolicy struct {
	// Specifies the criteria for aging
	AgingCriteria NullableString `json:"agingCriteria"`
	// Files that possess timestamps exceeding the specified value will be eligible for selection.
	DateInUsecs NullableInt64 `json:"dateInUsecs,omitempty"`
	// Files that possess timestamps older than the specified value in days will be eligible for selection.
	Days NullableInt32 `json:"days,omitempty"`
}

type _FilerLifecycleAgingPolicy FilerLifecycleAgingPolicy

// NewFilerLifecycleAgingPolicy instantiates a new FilerLifecycleAgingPolicy object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFilerLifecycleAgingPolicy(agingCriteria NullableString) *FilerLifecycleAgingPolicy {
	this := FilerLifecycleAgingPolicy{}
	this.AgingCriteria = agingCriteria
	return &this
}

// NewFilerLifecycleAgingPolicyWithDefaults instantiates a new FilerLifecycleAgingPolicy object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFilerLifecycleAgingPolicyWithDefaults() *FilerLifecycleAgingPolicy {
	this := FilerLifecycleAgingPolicy{}
	return &this
}

// GetAgingCriteria returns the AgingCriteria field value
// If the value is explicit nil, the zero value for string will be returned
func (o *FilerLifecycleAgingPolicy) GetAgingCriteria() string {
	if o == nil || o.AgingCriteria.Get() == nil {
		var ret string
		return ret
	}

	return *o.AgingCriteria.Get()
}

// GetAgingCriteriaOk returns a tuple with the AgingCriteria field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FilerLifecycleAgingPolicy) GetAgingCriteriaOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.AgingCriteria.Get(), o.AgingCriteria.IsSet()
}

// SetAgingCriteria sets field value
func (o *FilerLifecycleAgingPolicy) SetAgingCriteria(v string) {
	o.AgingCriteria.Set(&v)
}

// GetDateInUsecs returns the DateInUsecs field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FilerLifecycleAgingPolicy) GetDateInUsecs() int64 {
	if o == nil || IsNil(o.DateInUsecs.Get()) {
		var ret int64
		return ret
	}
	return *o.DateInUsecs.Get()
}

// GetDateInUsecsOk returns a tuple with the DateInUsecs field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FilerLifecycleAgingPolicy) GetDateInUsecsOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.DateInUsecs.Get(), o.DateInUsecs.IsSet()
}

// HasDateInUsecs returns a boolean if a field has been set.
func (o *FilerLifecycleAgingPolicy) HasDateInUsecs() bool {
	if o != nil && o.DateInUsecs.IsSet() {
		return true
	}

	return false
}

// SetDateInUsecs gets a reference to the given NullableInt64 and assigns it to the DateInUsecs field.
func (o *FilerLifecycleAgingPolicy) SetDateInUsecs(v int64) {
	o.DateInUsecs.Set(&v)
}
// SetDateInUsecsNil sets the value for DateInUsecs to be an explicit nil
func (o *FilerLifecycleAgingPolicy) SetDateInUsecsNil() {
	o.DateInUsecs.Set(nil)
}

// UnsetDateInUsecs ensures that no value is present for DateInUsecs, not even an explicit nil
func (o *FilerLifecycleAgingPolicy) UnsetDateInUsecs() {
	o.DateInUsecs.Unset()
}

// GetDays returns the Days field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FilerLifecycleAgingPolicy) GetDays() int32 {
	if o == nil || IsNil(o.Days.Get()) {
		var ret int32
		return ret
	}
	return *o.Days.Get()
}

// GetDaysOk returns a tuple with the Days field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FilerLifecycleAgingPolicy) GetDaysOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Days.Get(), o.Days.IsSet()
}

// HasDays returns a boolean if a field has been set.
func (o *FilerLifecycleAgingPolicy) HasDays() bool {
	if o != nil && o.Days.IsSet() {
		return true
	}

	return false
}

// SetDays gets a reference to the given NullableInt32 and assigns it to the Days field.
func (o *FilerLifecycleAgingPolicy) SetDays(v int32) {
	o.Days.Set(&v)
}
// SetDaysNil sets the value for Days to be an explicit nil
func (o *FilerLifecycleAgingPolicy) SetDaysNil() {
	o.Days.Set(nil)
}

// UnsetDays ensures that no value is present for Days, not even an explicit nil
func (o *FilerLifecycleAgingPolicy) UnsetDays() {
	o.Days.Unset()
}

func (o FilerLifecycleAgingPolicy) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FilerLifecycleAgingPolicy) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["agingCriteria"] = o.AgingCriteria.Get()
	if o.DateInUsecs.IsSet() {
		toSerialize["dateInUsecs"] = o.DateInUsecs.Get()
	}
	if o.Days.IsSet() {
		toSerialize["days"] = o.Days.Get()
	}
	return toSerialize, nil
}

func (o *FilerLifecycleAgingPolicy) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"agingCriteria",
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

	varFilerLifecycleAgingPolicy := _FilerLifecycleAgingPolicy{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varFilerLifecycleAgingPolicy)

	if err != nil {
		return err
	}

	*o = FilerLifecycleAgingPolicy(varFilerLifecycleAgingPolicy)

	return err
}

type NullableFilerLifecycleAgingPolicy struct {
	value *FilerLifecycleAgingPolicy
	isSet bool
}

func (v NullableFilerLifecycleAgingPolicy) Get() *FilerLifecycleAgingPolicy {
	return v.value
}

func (v *NullableFilerLifecycleAgingPolicy) Set(val *FilerLifecycleAgingPolicy) {
	v.value = val
	v.isSet = true
}

func (v NullableFilerLifecycleAgingPolicy) IsSet() bool {
	return v.isSet
}

func (v *NullableFilerLifecycleAgingPolicy) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFilerLifecycleAgingPolicy(val *FilerLifecycleAgingPolicy) *NullableFilerLifecycleAgingPolicy {
	return &NullableFilerLifecycleAgingPolicy{value: val, isSet: true}
}

func (v NullableFilerLifecycleAgingPolicy) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFilerLifecycleAgingPolicy) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


