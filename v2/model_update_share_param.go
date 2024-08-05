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

// checks if the UpdateShareParam type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateShareParam{}

// UpdateShareParam Specifies the parameter to update a Share.
type UpdateShareParam struct {
	// List of external client subnet IPs that are allowed to access the share.
	ClientSubnetWhitelist []Subnet `json:"clientSubnetWhitelist,omitempty"`
	// This field is currently deprecated. Specifies if Filer Audit Logging is enabled for this Share.
	EnableFilerAuditLogging NullableBool `json:"enableFilerAuditLogging,omitempty"`
	// Specifies the state of File Audit logging for this Share. Inherited: Audit log setting is inherited from the  View. Enabled: Audit log is enabled for this Share. Disabled: Audit log is disabled for this Share.
	FileAuditLoggingState NullableString `json:"fileAuditLoggingState,omitempty"`
	SmbConfig *UpdateShareParamSmbConfig `json:"smbConfig,omitempty"`
}

// NewUpdateShareParam instantiates a new UpdateShareParam object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateShareParam() *UpdateShareParam {
	this := UpdateShareParam{}
	return &this
}

// NewUpdateShareParamWithDefaults instantiates a new UpdateShareParam object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateShareParamWithDefaults() *UpdateShareParam {
	this := UpdateShareParam{}
	return &this
}

// GetClientSubnetWhitelist returns the ClientSubnetWhitelist field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UpdateShareParam) GetClientSubnetWhitelist() []Subnet {
	if o == nil {
		var ret []Subnet
		return ret
	}
	return o.ClientSubnetWhitelist
}

// GetClientSubnetWhitelistOk returns a tuple with the ClientSubnetWhitelist field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UpdateShareParam) GetClientSubnetWhitelistOk() ([]Subnet, bool) {
	if o == nil || IsNil(o.ClientSubnetWhitelist) {
		return nil, false
	}
	return o.ClientSubnetWhitelist, true
}

// HasClientSubnetWhitelist returns a boolean if a field has been set.
func (o *UpdateShareParam) HasClientSubnetWhitelist() bool {
	if o != nil && !IsNil(o.ClientSubnetWhitelist) {
		return true
	}

	return false
}

// SetClientSubnetWhitelist gets a reference to the given []Subnet and assigns it to the ClientSubnetWhitelist field.
func (o *UpdateShareParam) SetClientSubnetWhitelist(v []Subnet) {
	o.ClientSubnetWhitelist = v
}

// GetEnableFilerAuditLogging returns the EnableFilerAuditLogging field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UpdateShareParam) GetEnableFilerAuditLogging() bool {
	if o == nil || IsNil(o.EnableFilerAuditLogging.Get()) {
		var ret bool
		return ret
	}
	return *o.EnableFilerAuditLogging.Get()
}

// GetEnableFilerAuditLoggingOk returns a tuple with the EnableFilerAuditLogging field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UpdateShareParam) GetEnableFilerAuditLoggingOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.EnableFilerAuditLogging.Get(), o.EnableFilerAuditLogging.IsSet()
}

// HasEnableFilerAuditLogging returns a boolean if a field has been set.
func (o *UpdateShareParam) HasEnableFilerAuditLogging() bool {
	if o != nil && o.EnableFilerAuditLogging.IsSet() {
		return true
	}

	return false
}

// SetEnableFilerAuditLogging gets a reference to the given NullableBool and assigns it to the EnableFilerAuditLogging field.
func (o *UpdateShareParam) SetEnableFilerAuditLogging(v bool) {
	o.EnableFilerAuditLogging.Set(&v)
}
// SetEnableFilerAuditLoggingNil sets the value for EnableFilerAuditLogging to be an explicit nil
func (o *UpdateShareParam) SetEnableFilerAuditLoggingNil() {
	o.EnableFilerAuditLogging.Set(nil)
}

// UnsetEnableFilerAuditLogging ensures that no value is present for EnableFilerAuditLogging, not even an explicit nil
func (o *UpdateShareParam) UnsetEnableFilerAuditLogging() {
	o.EnableFilerAuditLogging.Unset()
}

// GetFileAuditLoggingState returns the FileAuditLoggingState field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UpdateShareParam) GetFileAuditLoggingState() string {
	if o == nil || IsNil(o.FileAuditLoggingState.Get()) {
		var ret string
		return ret
	}
	return *o.FileAuditLoggingState.Get()
}

// GetFileAuditLoggingStateOk returns a tuple with the FileAuditLoggingState field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UpdateShareParam) GetFileAuditLoggingStateOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.FileAuditLoggingState.Get(), o.FileAuditLoggingState.IsSet()
}

// HasFileAuditLoggingState returns a boolean if a field has been set.
func (o *UpdateShareParam) HasFileAuditLoggingState() bool {
	if o != nil && o.FileAuditLoggingState.IsSet() {
		return true
	}

	return false
}

// SetFileAuditLoggingState gets a reference to the given NullableString and assigns it to the FileAuditLoggingState field.
func (o *UpdateShareParam) SetFileAuditLoggingState(v string) {
	o.FileAuditLoggingState.Set(&v)
}
// SetFileAuditLoggingStateNil sets the value for FileAuditLoggingState to be an explicit nil
func (o *UpdateShareParam) SetFileAuditLoggingStateNil() {
	o.FileAuditLoggingState.Set(nil)
}

// UnsetFileAuditLoggingState ensures that no value is present for FileAuditLoggingState, not even an explicit nil
func (o *UpdateShareParam) UnsetFileAuditLoggingState() {
	o.FileAuditLoggingState.Unset()
}

// GetSmbConfig returns the SmbConfig field value if set, zero value otherwise.
func (o *UpdateShareParam) GetSmbConfig() UpdateShareParamSmbConfig {
	if o == nil || IsNil(o.SmbConfig) {
		var ret UpdateShareParamSmbConfig
		return ret
	}
	return *o.SmbConfig
}

// GetSmbConfigOk returns a tuple with the SmbConfig field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateShareParam) GetSmbConfigOk() (*UpdateShareParamSmbConfig, bool) {
	if o == nil || IsNil(o.SmbConfig) {
		return nil, false
	}
	return o.SmbConfig, true
}

// HasSmbConfig returns a boolean if a field has been set.
func (o *UpdateShareParam) HasSmbConfig() bool {
	if o != nil && !IsNil(o.SmbConfig) {
		return true
	}

	return false
}

// SetSmbConfig gets a reference to the given UpdateShareParamSmbConfig and assigns it to the SmbConfig field.
func (o *UpdateShareParam) SetSmbConfig(v UpdateShareParamSmbConfig) {
	o.SmbConfig = &v
}

func (o UpdateShareParam) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateShareParam) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.ClientSubnetWhitelist != nil {
		toSerialize["clientSubnetWhitelist"] = o.ClientSubnetWhitelist
	}
	if o.EnableFilerAuditLogging.IsSet() {
		toSerialize["enableFilerAuditLogging"] = o.EnableFilerAuditLogging.Get()
	}
	if o.FileAuditLoggingState.IsSet() {
		toSerialize["fileAuditLoggingState"] = o.FileAuditLoggingState.Get()
	}
	if !IsNil(o.SmbConfig) {
		toSerialize["smbConfig"] = o.SmbConfig
	}
	return toSerialize, nil
}

type NullableUpdateShareParam struct {
	value *UpdateShareParam
	isSet bool
}

func (v NullableUpdateShareParam) Get() *UpdateShareParam {
	return v.value
}

func (v *NullableUpdateShareParam) Set(val *UpdateShareParam) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateShareParam) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateShareParam) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateShareParam(val *UpdateShareParam) *NullableUpdateShareParam {
	return &NullableUpdateShareParam{value: val, isSet: true}
}

func (v NullableUpdateShareParam) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateShareParam) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


