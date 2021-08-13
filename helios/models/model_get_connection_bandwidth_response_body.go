/*
 * Cohesity REST API
 *
 * Cohesity API provides a RESTful interface to access the various data management operations on Cohesity cluster and Helios.
 *
 * API version: 2.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

// package helios
package models

import (
	"encoding/json"
	. "github.com/cohesity/go-sdk/helios/utils"
)

var _ = NullableBool{}

// GetConnectionBandwidthResponseBody The response body for put and get connection bandwidth.
type GetConnectionBandwidthResponseBody struct {
	// Connection Id for which bandwidth settings are to be returned
	ConnectionId NullableInt64 `json:"connectionId,omitempty"`
	// Specifies the max rate limit at which we download the data.
	Download []BandwidthLimit `json:"download,omitempty"`
	// Specifies the max rate limit at which we upload the data.
	Upload []BandwidthLimit `json:"upload,omitempty"`
	// The tenant Id corresponding to this request.
	TenantId NullableString `json:"tenantId"`
	// Specifies a time zone for the specified time period. The time zone is defined in the following format: 'Area/Location', for example: 'America/New_York'.
	Timezone NullableString `json:"timezone,omitempty"`
}

// NewGetConnectionBandwidthResponseBody instantiates a new GetConnectionBandwidthResponseBody object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetConnectionBandwidthResponseBody(tenantId NullableString) *GetConnectionBandwidthResponseBody {
	this := GetConnectionBandwidthResponseBody{}
	this.TenantId = tenantId
	return &this
}

// NewGetConnectionBandwidthResponseBodyWithDefaults instantiates a new GetConnectionBandwidthResponseBody object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetConnectionBandwidthResponseBodyWithDefaults() *GetConnectionBandwidthResponseBody {
	this := GetConnectionBandwidthResponseBody{}
	return &this
}

// GetConnectionId returns the ConnectionId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GetConnectionBandwidthResponseBody) GetConnectionId() int64 {
	if o == nil || o.ConnectionId.Get() == nil {
		var ret int64
		return ret
	}
	return *o.ConnectionId.Get()
}

// GetConnectionIdOk returns a tuple with the ConnectionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GetConnectionBandwidthResponseBody) GetConnectionIdOk() (*int64, bool) {
	if o == nil  {
		return nil, false
	}
	return o.ConnectionId.Get(), o.ConnectionId.IsSet()
}

// HasConnectionId returns a boolean if a field has been set.
func (o *GetConnectionBandwidthResponseBody) HasConnectionId() bool {
	if o != nil && o.ConnectionId.IsSet() {
		return true
	}

	return false
}

// SetConnectionId gets a reference to the given NullableInt64 and assigns it to the ConnectionId field.
func (o *GetConnectionBandwidthResponseBody) SetConnectionId(v int64) {
	o.ConnectionId.Set(&v)
}
// SetConnectionIdNil sets the value for ConnectionId to be an explicit nil
func (o *GetConnectionBandwidthResponseBody) SetConnectionIdNil() {
	o.ConnectionId.Set(nil)
}

// UnsetConnectionId ensures that no value is present for ConnectionId, not even an explicit nil
func (o *GetConnectionBandwidthResponseBody) UnsetConnectionId() {
	o.ConnectionId.Unset()
}

// GetDownload returns the Download field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GetConnectionBandwidthResponseBody) GetDownload() []BandwidthLimit {
	if o == nil  {
		var ret []BandwidthLimit
		return ret
	}
	return o.Download
}

// GetDownloadOk returns a tuple with the Download field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GetConnectionBandwidthResponseBody) GetDownloadOk() (*[]BandwidthLimit, bool) {
	if o == nil || o.Download == nil {
		return nil, false
	}
	return &o.Download, true
}

// HasDownload returns a boolean if a field has been set.
func (o *GetConnectionBandwidthResponseBody) HasDownload() bool {
	if o != nil && o.Download != nil {
		return true
	}

	return false
}

// SetDownload gets a reference to the given []BandwidthLimit and assigns it to the Download field.
func (o *GetConnectionBandwidthResponseBody) SetDownload(v []BandwidthLimit) {
	o.Download = v
}

// GetUpload returns the Upload field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GetConnectionBandwidthResponseBody) GetUpload() []BandwidthLimit {
	if o == nil  {
		var ret []BandwidthLimit
		return ret
	}
	return o.Upload
}

// GetUploadOk returns a tuple with the Upload field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GetConnectionBandwidthResponseBody) GetUploadOk() (*[]BandwidthLimit, bool) {
	if o == nil || o.Upload == nil {
		return nil, false
	}
	return &o.Upload, true
}

// HasUpload returns a boolean if a field has been set.
func (o *GetConnectionBandwidthResponseBody) HasUpload() bool {
	if o != nil && o.Upload != nil {
		return true
	}

	return false
}

// SetUpload gets a reference to the given []BandwidthLimit and assigns it to the Upload field.
func (o *GetConnectionBandwidthResponseBody) SetUpload(v []BandwidthLimit) {
	o.Upload = v
}

// GetTenantId returns the TenantId field value
// If the value is explicit nil, the zero value for string will be returned
func (o *GetConnectionBandwidthResponseBody) GetTenantId() string {
	if o == nil || o.TenantId.Get() == nil {
		var ret string
		return ret
	}

	return *o.TenantId.Get()
}

// GetTenantIdOk returns a tuple with the TenantId field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GetConnectionBandwidthResponseBody) GetTenantIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.TenantId.Get(), o.TenantId.IsSet()
}

// SetTenantId sets field value
func (o *GetConnectionBandwidthResponseBody) SetTenantId(v string) {
	o.TenantId.Set(&v)
}

// GetTimezone returns the Timezone field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GetConnectionBandwidthResponseBody) GetTimezone() string {
	if o == nil || o.Timezone.Get() == nil {
		var ret string
		return ret
	}
	return *o.Timezone.Get()
}

// GetTimezoneOk returns a tuple with the Timezone field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GetConnectionBandwidthResponseBody) GetTimezoneOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Timezone.Get(), o.Timezone.IsSet()
}

// HasTimezone returns a boolean if a field has been set.
func (o *GetConnectionBandwidthResponseBody) HasTimezone() bool {
	if o != nil && o.Timezone.IsSet() {
		return true
	}

	return false
}

// SetTimezone gets a reference to the given NullableString and assigns it to the Timezone field.
func (o *GetConnectionBandwidthResponseBody) SetTimezone(v string) {
	o.Timezone.Set(&v)
}
// SetTimezoneNil sets the value for Timezone to be an explicit nil
func (o *GetConnectionBandwidthResponseBody) SetTimezoneNil() {
	o.Timezone.Set(nil)
}

// UnsetTimezone ensures that no value is present for Timezone, not even an explicit nil
func (o *GetConnectionBandwidthResponseBody) UnsetTimezone() {
	o.Timezone.Unset()
}

func (o GetConnectionBandwidthResponseBody) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ConnectionId.IsSet() {
		toSerialize["connectionId"] = o.ConnectionId.Get()
	}
	if o.Download != nil {
		toSerialize["download"] = o.Download
	}
	if o.Upload != nil {
		toSerialize["upload"] = o.Upload
	}
	if true {
		toSerialize["tenantId"] = o.TenantId.Get()
	}
	if o.Timezone.IsSet() {
		toSerialize["timezone"] = o.Timezone.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableGetConnectionBandwidthResponseBody struct {
	value *GetConnectionBandwidthResponseBody
	isSet bool
}

func (v NullableGetConnectionBandwidthResponseBody) Get() *GetConnectionBandwidthResponseBody {
	return v.value
}

func (v *NullableGetConnectionBandwidthResponseBody) Set(val *GetConnectionBandwidthResponseBody) {
	v.value = val
	v.isSet = true
}

func (v NullableGetConnectionBandwidthResponseBody) IsSet() bool {
	return v.isSet
}

func (v *NullableGetConnectionBandwidthResponseBody) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetConnectionBandwidthResponseBody(val *GetConnectionBandwidthResponseBody) *NullableGetConnectionBandwidthResponseBody {
	return &NullableGetConnectionBandwidthResponseBody{value: val, isSet: true}
}

func (v NullableGetConnectionBandwidthResponseBody) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetConnectionBandwidthResponseBody) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


