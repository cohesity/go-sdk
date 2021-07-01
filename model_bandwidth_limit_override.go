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

// BandwidthLimitOverride Specifies bandwidth limit override value to be enforced during the specified daily time period for the specified days of the week.
type BandwidthLimitOverride struct {
	// Specifies the value to override the regular maximum bandwidth rate (rateLimitBytesPerSec) for the specified time period. The value is specified in bytes per second.
	BytesPerSecond NullableInt64 `json:"bytesPerSecond,omitempty"`
	// Specifies the value to override the default IO rate for the specified time period.
	IoRate NullableInt32 `json:"ioRate,omitempty"`
	TimePeriods *TimeOfAWeek `json:"timePeriods,omitempty"`
}

// NewBandwidthLimitOverride instantiates a new BandwidthLimitOverride object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBandwidthLimitOverride() *BandwidthLimitOverride {
	this := BandwidthLimitOverride{}
	return &this
}

// NewBandwidthLimitOverrideWithDefaults instantiates a new BandwidthLimitOverride object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBandwidthLimitOverrideWithDefaults() *BandwidthLimitOverride {
	this := BandwidthLimitOverride{}
	return &this
}

// GetBytesPerSecond returns the BytesPerSecond field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BandwidthLimitOverride) GetBytesPerSecond() int64 {
	if o == nil || o.BytesPerSecond.Get() == nil {
		var ret int64
		return ret
	}
	return *o.BytesPerSecond.Get()
}

// GetBytesPerSecondOk returns a tuple with the BytesPerSecond field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BandwidthLimitOverride) GetBytesPerSecondOk() (*int64, bool) {
	if o == nil  {
		return nil, false
	}
	return o.BytesPerSecond.Get(), o.BytesPerSecond.IsSet()
}

// HasBytesPerSecond returns a boolean if a field has been set.
func (o *BandwidthLimitOverride) HasBytesPerSecond() bool {
	if o != nil && o.BytesPerSecond.IsSet() {
		return true
	}

	return false
}

// SetBytesPerSecond gets a reference to the given NullableInt64 and assigns it to the BytesPerSecond field.
func (o *BandwidthLimitOverride) SetBytesPerSecond(v int64) {
	o.BytesPerSecond.Set(&v)
}
// SetBytesPerSecondNil sets the value for BytesPerSecond to be an explicit nil
func (o *BandwidthLimitOverride) SetBytesPerSecondNil() {
	o.BytesPerSecond.Set(nil)
}

// UnsetBytesPerSecond ensures that no value is present for BytesPerSecond, not even an explicit nil
func (o *BandwidthLimitOverride) UnsetBytesPerSecond() {
	o.BytesPerSecond.Unset()
}

// GetIoRate returns the IoRate field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BandwidthLimitOverride) GetIoRate() int32 {
	if o == nil || o.IoRate.Get() == nil {
		var ret int32
		return ret
	}
	return *o.IoRate.Get()
}

// GetIoRateOk returns a tuple with the IoRate field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BandwidthLimitOverride) GetIoRateOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return o.IoRate.Get(), o.IoRate.IsSet()
}

// HasIoRate returns a boolean if a field has been set.
func (o *BandwidthLimitOverride) HasIoRate() bool {
	if o != nil && o.IoRate.IsSet() {
		return true
	}

	return false
}

// SetIoRate gets a reference to the given NullableInt32 and assigns it to the IoRate field.
func (o *BandwidthLimitOverride) SetIoRate(v int32) {
	o.IoRate.Set(&v)
}
// SetIoRateNil sets the value for IoRate to be an explicit nil
func (o *BandwidthLimitOverride) SetIoRateNil() {
	o.IoRate.Set(nil)
}

// UnsetIoRate ensures that no value is present for IoRate, not even an explicit nil
func (o *BandwidthLimitOverride) UnsetIoRate() {
	o.IoRate.Unset()
}

// GetTimePeriods returns the TimePeriods field value if set, zero value otherwise.
func (o *BandwidthLimitOverride) GetTimePeriods() TimeOfAWeek {
	if o == nil || o.TimePeriods == nil {
		var ret TimeOfAWeek
		return ret
	}
	return *o.TimePeriods
}

// GetTimePeriodsOk returns a tuple with the TimePeriods field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BandwidthLimitOverride) GetTimePeriodsOk() (*TimeOfAWeek, bool) {
	if o == nil || o.TimePeriods == nil {
		return nil, false
	}
	return o.TimePeriods, true
}

// HasTimePeriods returns a boolean if a field has been set.
func (o *BandwidthLimitOverride) HasTimePeriods() bool {
	if o != nil && o.TimePeriods != nil {
		return true
	}

	return false
}

// SetTimePeriods gets a reference to the given TimeOfAWeek and assigns it to the TimePeriods field.
func (o *BandwidthLimitOverride) SetTimePeriods(v TimeOfAWeek) {
	o.TimePeriods = &v
}

func (o BandwidthLimitOverride) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.BytesPerSecond.IsSet() {
		toSerialize["bytesPerSecond"] = o.BytesPerSecond.Get()
	}
	if o.IoRate.IsSet() {
		toSerialize["ioRate"] = o.IoRate.Get()
	}
	if o.TimePeriods != nil {
		toSerialize["timePeriods"] = o.TimePeriods
	}
	return json.Marshal(toSerialize)
}

type NullableBandwidthLimitOverride struct {
	value *BandwidthLimitOverride
	isSet bool
}

func (v NullableBandwidthLimitOverride) Get() *BandwidthLimitOverride {
	return v.value
}

func (v *NullableBandwidthLimitOverride) Set(val *BandwidthLimitOverride) {
	v.value = val
	v.isSet = true
}

func (v NullableBandwidthLimitOverride) IsSet() bool {
	return v.isSet
}

func (v *NullableBandwidthLimitOverride) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBandwidthLimitOverride(val *BandwidthLimitOverride) *NullableBandwidthLimitOverride {
	return &NullableBandwidthLimitOverride{value: val, isSet: true}
}

func (v NullableBandwidthLimitOverride) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBandwidthLimitOverride) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


