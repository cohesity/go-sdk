/*
Cohesity REST API

Cohesity API provides a RESTful interface to access the various data management operations on Cohesity cluster and Helios.

API version: 2.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package v2

import (
	"encoding/json"
)

// checks if the AlertResolutionDetails type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AlertResolutionDetails{}

// AlertResolutionDetails Specifies information about the Alert Resolution.
type AlertResolutionDetails struct {
	// Specifies detailed notes about the Resolution.
	ResolutionDetails NullableString `json:"resolutionDetails,omitempty"`
	// Specifies the unique resolution id assigned in helios.
	ResolutionId NullableInt64 `json:"resolutionId,omitempty"`
	// Specifies short description about the Resolution.
	ResolutionSummary NullableString `json:"resolutionSummary,omitempty"`
	// Specifies unix epoch timestamp (in microseconds) when the Alert was resolved.
	TimestampUsecs NullableInt64 `json:"timestampUsecs,omitempty"`
	// Specifies name of the Cohesity Cluster user who resolved the Alerts.
	UserName NullableString `json:"userName,omitempty"`
}

// NewAlertResolutionDetails instantiates a new AlertResolutionDetails object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAlertResolutionDetails() *AlertResolutionDetails {
	this := AlertResolutionDetails{}
	return &this
}

// NewAlertResolutionDetailsWithDefaults instantiates a new AlertResolutionDetails object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAlertResolutionDetailsWithDefaults() *AlertResolutionDetails {
	this := AlertResolutionDetails{}
	return &this
}

// GetResolutionDetails returns the ResolutionDetails field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AlertResolutionDetails) GetResolutionDetails() string {
	if o == nil || IsNil(o.ResolutionDetails.Get()) {
		var ret string
		return ret
	}
	return *o.ResolutionDetails.Get()
}

// GetResolutionDetailsOk returns a tuple with the ResolutionDetails field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AlertResolutionDetails) GetResolutionDetailsOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ResolutionDetails.Get(), o.ResolutionDetails.IsSet()
}

// HasResolutionDetails returns a boolean if a field has been set.
func (o *AlertResolutionDetails) HasResolutionDetails() bool {
	if o != nil && o.ResolutionDetails.IsSet() {
		return true
	}

	return false
}

// SetResolutionDetails gets a reference to the given NullableString and assigns it to the ResolutionDetails field.
func (o *AlertResolutionDetails) SetResolutionDetails(v string) {
	o.ResolutionDetails.Set(&v)
}
// SetResolutionDetailsNil sets the value for ResolutionDetails to be an explicit nil
func (o *AlertResolutionDetails) SetResolutionDetailsNil() {
	o.ResolutionDetails.Set(nil)
}

// UnsetResolutionDetails ensures that no value is present for ResolutionDetails, not even an explicit nil
func (o *AlertResolutionDetails) UnsetResolutionDetails() {
	o.ResolutionDetails.Unset()
}

// GetResolutionId returns the ResolutionId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AlertResolutionDetails) GetResolutionId() int64 {
	if o == nil || IsNil(o.ResolutionId.Get()) {
		var ret int64
		return ret
	}
	return *o.ResolutionId.Get()
}

// GetResolutionIdOk returns a tuple with the ResolutionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AlertResolutionDetails) GetResolutionIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.ResolutionId.Get(), o.ResolutionId.IsSet()
}

// HasResolutionId returns a boolean if a field has been set.
func (o *AlertResolutionDetails) HasResolutionId() bool {
	if o != nil && o.ResolutionId.IsSet() {
		return true
	}

	return false
}

// SetResolutionId gets a reference to the given NullableInt64 and assigns it to the ResolutionId field.
func (o *AlertResolutionDetails) SetResolutionId(v int64) {
	o.ResolutionId.Set(&v)
}
// SetResolutionIdNil sets the value for ResolutionId to be an explicit nil
func (o *AlertResolutionDetails) SetResolutionIdNil() {
	o.ResolutionId.Set(nil)
}

// UnsetResolutionId ensures that no value is present for ResolutionId, not even an explicit nil
func (o *AlertResolutionDetails) UnsetResolutionId() {
	o.ResolutionId.Unset()
}

// GetResolutionSummary returns the ResolutionSummary field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AlertResolutionDetails) GetResolutionSummary() string {
	if o == nil || IsNil(o.ResolutionSummary.Get()) {
		var ret string
		return ret
	}
	return *o.ResolutionSummary.Get()
}

// GetResolutionSummaryOk returns a tuple with the ResolutionSummary field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AlertResolutionDetails) GetResolutionSummaryOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ResolutionSummary.Get(), o.ResolutionSummary.IsSet()
}

// HasResolutionSummary returns a boolean if a field has been set.
func (o *AlertResolutionDetails) HasResolutionSummary() bool {
	if o != nil && o.ResolutionSummary.IsSet() {
		return true
	}

	return false
}

// SetResolutionSummary gets a reference to the given NullableString and assigns it to the ResolutionSummary field.
func (o *AlertResolutionDetails) SetResolutionSummary(v string) {
	o.ResolutionSummary.Set(&v)
}
// SetResolutionSummaryNil sets the value for ResolutionSummary to be an explicit nil
func (o *AlertResolutionDetails) SetResolutionSummaryNil() {
	o.ResolutionSummary.Set(nil)
}

// UnsetResolutionSummary ensures that no value is present for ResolutionSummary, not even an explicit nil
func (o *AlertResolutionDetails) UnsetResolutionSummary() {
	o.ResolutionSummary.Unset()
}

// GetTimestampUsecs returns the TimestampUsecs field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AlertResolutionDetails) GetTimestampUsecs() int64 {
	if o == nil || IsNil(o.TimestampUsecs.Get()) {
		var ret int64
		return ret
	}
	return *o.TimestampUsecs.Get()
}

// GetTimestampUsecsOk returns a tuple with the TimestampUsecs field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AlertResolutionDetails) GetTimestampUsecsOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.TimestampUsecs.Get(), o.TimestampUsecs.IsSet()
}

// HasTimestampUsecs returns a boolean if a field has been set.
func (o *AlertResolutionDetails) HasTimestampUsecs() bool {
	if o != nil && o.TimestampUsecs.IsSet() {
		return true
	}

	return false
}

// SetTimestampUsecs gets a reference to the given NullableInt64 and assigns it to the TimestampUsecs field.
func (o *AlertResolutionDetails) SetTimestampUsecs(v int64) {
	o.TimestampUsecs.Set(&v)
}
// SetTimestampUsecsNil sets the value for TimestampUsecs to be an explicit nil
func (o *AlertResolutionDetails) SetTimestampUsecsNil() {
	o.TimestampUsecs.Set(nil)
}

// UnsetTimestampUsecs ensures that no value is present for TimestampUsecs, not even an explicit nil
func (o *AlertResolutionDetails) UnsetTimestampUsecs() {
	o.TimestampUsecs.Unset()
}

// GetUserName returns the UserName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AlertResolutionDetails) GetUserName() string {
	if o == nil || IsNil(o.UserName.Get()) {
		var ret string
		return ret
	}
	return *o.UserName.Get()
}

// GetUserNameOk returns a tuple with the UserName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AlertResolutionDetails) GetUserNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.UserName.Get(), o.UserName.IsSet()
}

// HasUserName returns a boolean if a field has been set.
func (o *AlertResolutionDetails) HasUserName() bool {
	if o != nil && o.UserName.IsSet() {
		return true
	}

	return false
}

// SetUserName gets a reference to the given NullableString and assigns it to the UserName field.
func (o *AlertResolutionDetails) SetUserName(v string) {
	o.UserName.Set(&v)
}
// SetUserNameNil sets the value for UserName to be an explicit nil
func (o *AlertResolutionDetails) SetUserNameNil() {
	o.UserName.Set(nil)
}

// UnsetUserName ensures that no value is present for UserName, not even an explicit nil
func (o *AlertResolutionDetails) UnsetUserName() {
	o.UserName.Unset()
}

func (o AlertResolutionDetails) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AlertResolutionDetails) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.ResolutionDetails.IsSet() {
		toSerialize["resolutionDetails"] = o.ResolutionDetails.Get()
	}
	if o.ResolutionId.IsSet() {
		toSerialize["resolutionId"] = o.ResolutionId.Get()
	}
	if o.ResolutionSummary.IsSet() {
		toSerialize["resolutionSummary"] = o.ResolutionSummary.Get()
	}
	if o.TimestampUsecs.IsSet() {
		toSerialize["timestampUsecs"] = o.TimestampUsecs.Get()
	}
	if o.UserName.IsSet() {
		toSerialize["userName"] = o.UserName.Get()
	}
	return toSerialize, nil
}

type NullableAlertResolutionDetails struct {
	value *AlertResolutionDetails
	isSet bool
}

func (v NullableAlertResolutionDetails) Get() *AlertResolutionDetails {
	return v.value
}

func (v *NullableAlertResolutionDetails) Set(val *AlertResolutionDetails) {
	v.value = val
	v.isSet = true
}

func (v NullableAlertResolutionDetails) IsSet() bool {
	return v.isSet
}

func (v *NullableAlertResolutionDetails) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlertResolutionDetails(val *AlertResolutionDetails) *NullableAlertResolutionDetails {
	return &NullableAlertResolutionDetails{value: val, isSet: true}
}

func (v NullableAlertResolutionDetails) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlertResolutionDetails) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


