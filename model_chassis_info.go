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

// ChassisInfo ChassisInfo is the struct for the Chassis.
type ChassisInfo struct {
	// ChassisId is a unique id assigned to the chassis.
	ChassisId NullableInt64 `json:"chassisId,omitempty"`
	// ChassisName is the name of the chassis. This could be the chassis serial number by default.
	ChassisName NullableString `json:"chassisName,omitempty"`
	// Chassis serial.
	ChassisSerial NullableString `json:"chassisSerial,omitempty"`
	// Location is the location of the chassis within the rack.
	Location NullableString `json:"location,omitempty"`
	// Rack is the rack within which this chassis lives.
	RackId NullableInt64 `json:"rackId,omitempty"`
}

// NewChassisInfo instantiates a new ChassisInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewChassisInfo() *ChassisInfo {
	this := ChassisInfo{}
	return &this
}

// NewChassisInfoWithDefaults instantiates a new ChassisInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewChassisInfoWithDefaults() *ChassisInfo {
	this := ChassisInfo{}
	return &this
}

// GetChassisId returns the ChassisId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ChassisInfo) GetChassisId() int64 {
	if o == nil || o.ChassisId.Get() == nil {
		var ret int64
		return ret
	}
	return *o.ChassisId.Get()
}

// GetChassisIdOk returns a tuple with the ChassisId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ChassisInfo) GetChassisIdOk() (*int64, bool) {
	if o == nil  {
		return nil, false
	}
	return o.ChassisId.Get(), o.ChassisId.IsSet()
}

// HasChassisId returns a boolean if a field has been set.
func (o *ChassisInfo) HasChassisId() bool {
	if o != nil && o.ChassisId.IsSet() {
		return true
	}

	return false
}

// SetChassisId gets a reference to the given NullableInt64 and assigns it to the ChassisId field.
func (o *ChassisInfo) SetChassisId(v int64) {
	o.ChassisId.Set(&v)
}
// SetChassisIdNil sets the value for ChassisId to be an explicit nil
func (o *ChassisInfo) SetChassisIdNil() {
	o.ChassisId.Set(nil)
}

// UnsetChassisId ensures that no value is present for ChassisId, not even an explicit nil
func (o *ChassisInfo) UnsetChassisId() {
	o.ChassisId.Unset()
}

// GetChassisName returns the ChassisName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ChassisInfo) GetChassisName() string {
	if o == nil || o.ChassisName.Get() == nil {
		var ret string
		return ret
	}
	return *o.ChassisName.Get()
}

// GetChassisNameOk returns a tuple with the ChassisName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ChassisInfo) GetChassisNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.ChassisName.Get(), o.ChassisName.IsSet()
}

// HasChassisName returns a boolean if a field has been set.
func (o *ChassisInfo) HasChassisName() bool {
	if o != nil && o.ChassisName.IsSet() {
		return true
	}

	return false
}

// SetChassisName gets a reference to the given NullableString and assigns it to the ChassisName field.
func (o *ChassisInfo) SetChassisName(v string) {
	o.ChassisName.Set(&v)
}
// SetChassisNameNil sets the value for ChassisName to be an explicit nil
func (o *ChassisInfo) SetChassisNameNil() {
	o.ChassisName.Set(nil)
}

// UnsetChassisName ensures that no value is present for ChassisName, not even an explicit nil
func (o *ChassisInfo) UnsetChassisName() {
	o.ChassisName.Unset()
}

// GetChassisSerial returns the ChassisSerial field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ChassisInfo) GetChassisSerial() string {
	if o == nil || o.ChassisSerial.Get() == nil {
		var ret string
		return ret
	}
	return *o.ChassisSerial.Get()
}

// GetChassisSerialOk returns a tuple with the ChassisSerial field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ChassisInfo) GetChassisSerialOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.ChassisSerial.Get(), o.ChassisSerial.IsSet()
}

// HasChassisSerial returns a boolean if a field has been set.
func (o *ChassisInfo) HasChassisSerial() bool {
	if o != nil && o.ChassisSerial.IsSet() {
		return true
	}

	return false
}

// SetChassisSerial gets a reference to the given NullableString and assigns it to the ChassisSerial field.
func (o *ChassisInfo) SetChassisSerial(v string) {
	o.ChassisSerial.Set(&v)
}
// SetChassisSerialNil sets the value for ChassisSerial to be an explicit nil
func (o *ChassisInfo) SetChassisSerialNil() {
	o.ChassisSerial.Set(nil)
}

// UnsetChassisSerial ensures that no value is present for ChassisSerial, not even an explicit nil
func (o *ChassisInfo) UnsetChassisSerial() {
	o.ChassisSerial.Unset()
}

// GetLocation returns the Location field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ChassisInfo) GetLocation() string {
	if o == nil || o.Location.Get() == nil {
		var ret string
		return ret
	}
	return *o.Location.Get()
}

// GetLocationOk returns a tuple with the Location field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ChassisInfo) GetLocationOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Location.Get(), o.Location.IsSet()
}

// HasLocation returns a boolean if a field has been set.
func (o *ChassisInfo) HasLocation() bool {
	if o != nil && o.Location.IsSet() {
		return true
	}

	return false
}

// SetLocation gets a reference to the given NullableString and assigns it to the Location field.
func (o *ChassisInfo) SetLocation(v string) {
	o.Location.Set(&v)
}
// SetLocationNil sets the value for Location to be an explicit nil
func (o *ChassisInfo) SetLocationNil() {
	o.Location.Set(nil)
}

// UnsetLocation ensures that no value is present for Location, not even an explicit nil
func (o *ChassisInfo) UnsetLocation() {
	o.Location.Unset()
}

// GetRackId returns the RackId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ChassisInfo) GetRackId() int64 {
	if o == nil || o.RackId.Get() == nil {
		var ret int64
		return ret
	}
	return *o.RackId.Get()
}

// GetRackIdOk returns a tuple with the RackId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ChassisInfo) GetRackIdOk() (*int64, bool) {
	if o == nil  {
		return nil, false
	}
	return o.RackId.Get(), o.RackId.IsSet()
}

// HasRackId returns a boolean if a field has been set.
func (o *ChassisInfo) HasRackId() bool {
	if o != nil && o.RackId.IsSet() {
		return true
	}

	return false
}

// SetRackId gets a reference to the given NullableInt64 and assigns it to the RackId field.
func (o *ChassisInfo) SetRackId(v int64) {
	o.RackId.Set(&v)
}
// SetRackIdNil sets the value for RackId to be an explicit nil
func (o *ChassisInfo) SetRackIdNil() {
	o.RackId.Set(nil)
}

// UnsetRackId ensures that no value is present for RackId, not even an explicit nil
func (o *ChassisInfo) UnsetRackId() {
	o.RackId.Unset()
}

func (o ChassisInfo) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ChassisId.IsSet() {
		toSerialize["chassisId"] = o.ChassisId.Get()
	}
	if o.ChassisName.IsSet() {
		toSerialize["chassisName"] = o.ChassisName.Get()
	}
	if o.ChassisSerial.IsSet() {
		toSerialize["chassisSerial"] = o.ChassisSerial.Get()
	}
	if o.Location.IsSet() {
		toSerialize["location"] = o.Location.Get()
	}
	if o.RackId.IsSet() {
		toSerialize["rackId"] = o.RackId.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableChassisInfo struct {
	value *ChassisInfo
	isSet bool
}

func (v NullableChassisInfo) Get() *ChassisInfo {
	return v.value
}

func (v *NullableChassisInfo) Set(val *ChassisInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableChassisInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableChassisInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableChassisInfo(val *ChassisInfo) *NullableChassisInfo {
	return &NullableChassisInfo{value: val, isSet: true}
}

func (v NullableChassisInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableChassisInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


