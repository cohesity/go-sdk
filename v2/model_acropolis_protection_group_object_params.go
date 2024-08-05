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

// checks if the AcropolisProtectionGroupObjectParams type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AcropolisProtectionGroupObjectParams{}

// AcropolisProtectionGroupObjectParams Specifies an object protected by a Acropolis Protection Group.
type AcropolisProtectionGroupObjectParams struct {
	// Specifies a list of disks to exclude from being protected. This is only applicable to VM objects.
	ExcludeDisks []AcropolisDiskInfo `json:"excludeDisks,omitempty"`
	// Specifies the ID of the object.
	Id NullableInt64 `json:"id"`
	// Specifies a list of disks to include in the protection. This is only applicable to VM objects.
	IncludeDisks []AcropolisDiskInfo `json:"includeDisks,omitempty"`
	// Specifies the name of the virtual machine.
	Name NullableString `json:"name,omitempty"`
}

type _AcropolisProtectionGroupObjectParams AcropolisProtectionGroupObjectParams

// NewAcropolisProtectionGroupObjectParams instantiates a new AcropolisProtectionGroupObjectParams object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAcropolisProtectionGroupObjectParams(id NullableInt64) *AcropolisProtectionGroupObjectParams {
	this := AcropolisProtectionGroupObjectParams{}
	this.Id = id
	return &this
}

// NewAcropolisProtectionGroupObjectParamsWithDefaults instantiates a new AcropolisProtectionGroupObjectParams object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAcropolisProtectionGroupObjectParamsWithDefaults() *AcropolisProtectionGroupObjectParams {
	this := AcropolisProtectionGroupObjectParams{}
	return &this
}

// GetExcludeDisks returns the ExcludeDisks field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AcropolisProtectionGroupObjectParams) GetExcludeDisks() []AcropolisDiskInfo {
	if o == nil {
		var ret []AcropolisDiskInfo
		return ret
	}
	return o.ExcludeDisks
}

// GetExcludeDisksOk returns a tuple with the ExcludeDisks field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AcropolisProtectionGroupObjectParams) GetExcludeDisksOk() ([]AcropolisDiskInfo, bool) {
	if o == nil || IsNil(o.ExcludeDisks) {
		return nil, false
	}
	return o.ExcludeDisks, true
}

// HasExcludeDisks returns a boolean if a field has been set.
func (o *AcropolisProtectionGroupObjectParams) HasExcludeDisks() bool {
	if o != nil && !IsNil(o.ExcludeDisks) {
		return true
	}

	return false
}

// SetExcludeDisks gets a reference to the given []AcropolisDiskInfo and assigns it to the ExcludeDisks field.
func (o *AcropolisProtectionGroupObjectParams) SetExcludeDisks(v []AcropolisDiskInfo) {
	o.ExcludeDisks = v
}

// GetId returns the Id field value
// If the value is explicit nil, the zero value for int64 will be returned
func (o *AcropolisProtectionGroupObjectParams) GetId() int64 {
	if o == nil || o.Id.Get() == nil {
		var ret int64
		return ret
	}

	return *o.Id.Get()
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AcropolisProtectionGroupObjectParams) GetIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.Id.Get(), o.Id.IsSet()
}

// SetId sets field value
func (o *AcropolisProtectionGroupObjectParams) SetId(v int64) {
	o.Id.Set(&v)
}

// GetIncludeDisks returns the IncludeDisks field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AcropolisProtectionGroupObjectParams) GetIncludeDisks() []AcropolisDiskInfo {
	if o == nil {
		var ret []AcropolisDiskInfo
		return ret
	}
	return o.IncludeDisks
}

// GetIncludeDisksOk returns a tuple with the IncludeDisks field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AcropolisProtectionGroupObjectParams) GetIncludeDisksOk() ([]AcropolisDiskInfo, bool) {
	if o == nil || IsNil(o.IncludeDisks) {
		return nil, false
	}
	return o.IncludeDisks, true
}

// HasIncludeDisks returns a boolean if a field has been set.
func (o *AcropolisProtectionGroupObjectParams) HasIncludeDisks() bool {
	if o != nil && !IsNil(o.IncludeDisks) {
		return true
	}

	return false
}

// SetIncludeDisks gets a reference to the given []AcropolisDiskInfo and assigns it to the IncludeDisks field.
func (o *AcropolisProtectionGroupObjectParams) SetIncludeDisks(v []AcropolisDiskInfo) {
	o.IncludeDisks = v
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AcropolisProtectionGroupObjectParams) GetName() string {
	if o == nil || IsNil(o.Name.Get()) {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AcropolisProtectionGroupObjectParams) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *AcropolisProtectionGroupObjectParams) HasName() bool {
	if o != nil && o.Name.IsSet() {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableString and assigns it to the Name field.
func (o *AcropolisProtectionGroupObjectParams) SetName(v string) {
	o.Name.Set(&v)
}
// SetNameNil sets the value for Name to be an explicit nil
func (o *AcropolisProtectionGroupObjectParams) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *AcropolisProtectionGroupObjectParams) UnsetName() {
	o.Name.Unset()
}

func (o AcropolisProtectionGroupObjectParams) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AcropolisProtectionGroupObjectParams) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.ExcludeDisks != nil {
		toSerialize["excludeDisks"] = o.ExcludeDisks
	}
	toSerialize["id"] = o.Id.Get()
	if o.IncludeDisks != nil {
		toSerialize["includeDisks"] = o.IncludeDisks
	}
	if o.Name.IsSet() {
		toSerialize["name"] = o.Name.Get()
	}
	return toSerialize, nil
}

func (o *AcropolisProtectionGroupObjectParams) UnmarshalJSON(data []byte) (err error) {
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

	varAcropolisProtectionGroupObjectParams := _AcropolisProtectionGroupObjectParams{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varAcropolisProtectionGroupObjectParams)

	if err != nil {
		return err
	}

	*o = AcropolisProtectionGroupObjectParams(varAcropolisProtectionGroupObjectParams)

	return err
}

type NullableAcropolisProtectionGroupObjectParams struct {
	value *AcropolisProtectionGroupObjectParams
	isSet bool
}

func (v NullableAcropolisProtectionGroupObjectParams) Get() *AcropolisProtectionGroupObjectParams {
	return v.value
}

func (v *NullableAcropolisProtectionGroupObjectParams) Set(val *AcropolisProtectionGroupObjectParams) {
	v.value = val
	v.isSet = true
}

func (v NullableAcropolisProtectionGroupObjectParams) IsSet() bool {
	return v.isSet
}

func (v *NullableAcropolisProtectionGroupObjectParams) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAcropolisProtectionGroupObjectParams(val *AcropolisProtectionGroupObjectParams) *NullableAcropolisProtectionGroupObjectParams {
	return &NullableAcropolisProtectionGroupObjectParams{value: val, isSet: true}
}

func (v NullableAcropolisProtectionGroupObjectParams) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAcropolisProtectionGroupObjectParams) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


