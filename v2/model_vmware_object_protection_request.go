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

// checks if the VmwareObjectProtectionRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &VmwareObjectProtectionRequest{}

// VmwareObjectProtectionRequest Specifies the VMware object level settings for object protection.
type VmwareObjectProtectionRequest struct {
	// Specifies a list of disks to exclude from being protected. This is only applicable to VM objects.
	ExcludeDisks []DiskInfo `json:"excludeDisks,omitempty"`
	// Specifies whether or not to truncate MS Exchange logs while taking an app consistent snapshot of this object. This is only applicable to objects which have a registered MS Exchange app.
	TruncateExchangeLogs NullableBool `json:"truncateExchangeLogs,omitempty"`
	// Specifies the list of IDs of the objects to not be protected in this backup. This field only applies if provided object id is non leaf entity such as Tag or a folder. This can be used to ignore specific objects under a parent object which has been included for protection.
	ExcludeObjectIds []*int64 `json:"excludeObjectIds,omitempty"`
	// Specifies the id of the object being protected. This can be a leaf level or non leaf level object.
	Id NullableInt64 `json:"id"`
}

type _VmwareObjectProtectionRequest VmwareObjectProtectionRequest

// NewVmwareObjectProtectionRequest instantiates a new VmwareObjectProtectionRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVmwareObjectProtectionRequest(id NullableInt64) *VmwareObjectProtectionRequest {
	this := VmwareObjectProtectionRequest{}
	this.Id = id
	return &this
}

// NewVmwareObjectProtectionRequestWithDefaults instantiates a new VmwareObjectProtectionRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVmwareObjectProtectionRequestWithDefaults() *VmwareObjectProtectionRequest {
	this := VmwareObjectProtectionRequest{}
	return &this
}

// GetExcludeDisks returns the ExcludeDisks field value if set, zero value otherwise.
func (o *VmwareObjectProtectionRequest) GetExcludeDisks() []DiskInfo {
	if o == nil || IsNil(o.ExcludeDisks) {
		var ret []DiskInfo
		return ret
	}
	return o.ExcludeDisks
}

// GetExcludeDisksOk returns a tuple with the ExcludeDisks field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VmwareObjectProtectionRequest) GetExcludeDisksOk() ([]DiskInfo, bool) {
	if o == nil || IsNil(o.ExcludeDisks) {
		return nil, false
	}
	return o.ExcludeDisks, true
}

// HasExcludeDisks returns a boolean if a field has been set.
func (o *VmwareObjectProtectionRequest) HasExcludeDisks() bool {
	if o != nil && !IsNil(o.ExcludeDisks) {
		return true
	}

	return false
}

// SetExcludeDisks gets a reference to the given []DiskInfo and assigns it to the ExcludeDisks field.
func (o *VmwareObjectProtectionRequest) SetExcludeDisks(v []DiskInfo) {
	o.ExcludeDisks = v
}

// GetTruncateExchangeLogs returns the TruncateExchangeLogs field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VmwareObjectProtectionRequest) GetTruncateExchangeLogs() bool {
	if o == nil || IsNil(o.TruncateExchangeLogs.Get()) {
		var ret bool
		return ret
	}
	return *o.TruncateExchangeLogs.Get()
}

// GetTruncateExchangeLogsOk returns a tuple with the TruncateExchangeLogs field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VmwareObjectProtectionRequest) GetTruncateExchangeLogsOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.TruncateExchangeLogs.Get(), o.TruncateExchangeLogs.IsSet()
}

// HasTruncateExchangeLogs returns a boolean if a field has been set.
func (o *VmwareObjectProtectionRequest) HasTruncateExchangeLogs() bool {
	if o != nil && o.TruncateExchangeLogs.IsSet() {
		return true
	}

	return false
}

// SetTruncateExchangeLogs gets a reference to the given NullableBool and assigns it to the TruncateExchangeLogs field.
func (o *VmwareObjectProtectionRequest) SetTruncateExchangeLogs(v bool) {
	o.TruncateExchangeLogs.Set(&v)
}
// SetTruncateExchangeLogsNil sets the value for TruncateExchangeLogs to be an explicit nil
func (o *VmwareObjectProtectionRequest) SetTruncateExchangeLogsNil() {
	o.TruncateExchangeLogs.Set(nil)
}

// UnsetTruncateExchangeLogs ensures that no value is present for TruncateExchangeLogs, not even an explicit nil
func (o *VmwareObjectProtectionRequest) UnsetTruncateExchangeLogs() {
	o.TruncateExchangeLogs.Unset()
}

// GetExcludeObjectIds returns the ExcludeObjectIds field value if set, zero value otherwise.
func (o *VmwareObjectProtectionRequest) GetExcludeObjectIds() []*int64 {
	if o == nil || IsNil(o.ExcludeObjectIds) {
		var ret []*int64
		return ret
	}
	return o.ExcludeObjectIds
}

// GetExcludeObjectIdsOk returns a tuple with the ExcludeObjectIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VmwareObjectProtectionRequest) GetExcludeObjectIdsOk() ([]*int64, bool) {
	if o == nil || IsNil(o.ExcludeObjectIds) {
		return nil, false
	}
	return o.ExcludeObjectIds, true
}

// HasExcludeObjectIds returns a boolean if a field has been set.
func (o *VmwareObjectProtectionRequest) HasExcludeObjectIds() bool {
	if o != nil && !IsNil(o.ExcludeObjectIds) {
		return true
	}

	return false
}

// SetExcludeObjectIds gets a reference to the given []*int64 and assigns it to the ExcludeObjectIds field.
func (o *VmwareObjectProtectionRequest) SetExcludeObjectIds(v []*int64) {
	o.ExcludeObjectIds = v
}

// GetId returns the Id field value
// If the value is explicit nil, the zero value for int64 will be returned
func (o *VmwareObjectProtectionRequest) GetId() int64 {
	if o == nil || o.Id.Get() == nil {
		var ret int64
		return ret
	}

	return *o.Id.Get()
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VmwareObjectProtectionRequest) GetIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.Id.Get(), o.Id.IsSet()
}

// SetId sets field value
func (o *VmwareObjectProtectionRequest) SetId(v int64) {
	o.Id.Set(&v)
}

func (o VmwareObjectProtectionRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o VmwareObjectProtectionRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ExcludeDisks) {
		toSerialize["excludeDisks"] = o.ExcludeDisks
	}
	if o.TruncateExchangeLogs.IsSet() {
		toSerialize["truncateExchangeLogs"] = o.TruncateExchangeLogs.Get()
	}
	if !IsNil(o.ExcludeObjectIds) {
		toSerialize["excludeObjectIds"] = o.ExcludeObjectIds
	}
	toSerialize["id"] = o.Id.Get()
	return toSerialize, nil
}

func (o *VmwareObjectProtectionRequest) UnmarshalJSON(data []byte) (err error) {
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

	varVmwareObjectProtectionRequest := _VmwareObjectProtectionRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varVmwareObjectProtectionRequest)

	if err != nil {
		return err
	}

	*o = VmwareObjectProtectionRequest(varVmwareObjectProtectionRequest)

	return err
}

type NullableVmwareObjectProtectionRequest struct {
	value *VmwareObjectProtectionRequest
	isSet bool
}

func (v NullableVmwareObjectProtectionRequest) Get() *VmwareObjectProtectionRequest {
	return v.value
}

func (v *NullableVmwareObjectProtectionRequest) Set(val *VmwareObjectProtectionRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableVmwareObjectProtectionRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableVmwareObjectProtectionRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVmwareObjectProtectionRequest(val *VmwareObjectProtectionRequest) *NullableVmwareObjectProtectionRequest {
	return &NullableVmwareObjectProtectionRequest{value: val, isSet: true}
}

func (v NullableVmwareObjectProtectionRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVmwareObjectProtectionRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


