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

// RemoteAdapterHost Specifies params of the remote host.
type RemoteAdapterHost struct {
	// Specifies the Hostname or IP address of the host where the pre and post script will be run.
	Hostname NullableString `json:"hostname,omitempty"`
	// Specifies the username for the host.
	Username NullableString `json:"username,omitempty"`
	// Specifies the Operating system type of the host.
	HostType NullableString `json:"hostType,omitempty"`
	IncrementalBackupScript *CommonPreBackupScriptParams `json:"incrementalBackupScript,omitempty"`
	FullBackupScript *CommonPreBackupScriptParams `json:"fullBackupScript,omitempty"`
	LogBackupScript *CommonPreBackupScriptParams `json:"logBackupScript,omitempty"`
}

// NewRemoteAdapterHost instantiates a new RemoteAdapterHost object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRemoteAdapterHost() *RemoteAdapterHost {
	this := RemoteAdapterHost{}
	return &this
}

// NewRemoteAdapterHostWithDefaults instantiates a new RemoteAdapterHost object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRemoteAdapterHostWithDefaults() *RemoteAdapterHost {
	this := RemoteAdapterHost{}
	return &this
}

// GetHostname returns the Hostname field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RemoteAdapterHost) GetHostname() string {
	if o == nil || o.Hostname.Get() == nil {
		var ret string
		return ret
	}
	return *o.Hostname.Get()
}

// GetHostnameOk returns a tuple with the Hostname field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RemoteAdapterHost) GetHostnameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Hostname.Get(), o.Hostname.IsSet()
}

// HasHostname returns a boolean if a field has been set.
func (o *RemoteAdapterHost) HasHostname() bool {
	if o != nil && o.Hostname.IsSet() {
		return true
	}

	return false
}

// SetHostname gets a reference to the given NullableString and assigns it to the Hostname field.
func (o *RemoteAdapterHost) SetHostname(v string) {
	o.Hostname.Set(&v)
}
// SetHostnameNil sets the value for Hostname to be an explicit nil
func (o *RemoteAdapterHost) SetHostnameNil() {
	o.Hostname.Set(nil)
}

// UnsetHostname ensures that no value is present for Hostname, not even an explicit nil
func (o *RemoteAdapterHost) UnsetHostname() {
	o.Hostname.Unset()
}

// GetUsername returns the Username field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RemoteAdapterHost) GetUsername() string {
	if o == nil || o.Username.Get() == nil {
		var ret string
		return ret
	}
	return *o.Username.Get()
}

// GetUsernameOk returns a tuple with the Username field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RemoteAdapterHost) GetUsernameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Username.Get(), o.Username.IsSet()
}

// HasUsername returns a boolean if a field has been set.
func (o *RemoteAdapterHost) HasUsername() bool {
	if o != nil && o.Username.IsSet() {
		return true
	}

	return false
}

// SetUsername gets a reference to the given NullableString and assigns it to the Username field.
func (o *RemoteAdapterHost) SetUsername(v string) {
	o.Username.Set(&v)
}
// SetUsernameNil sets the value for Username to be an explicit nil
func (o *RemoteAdapterHost) SetUsernameNil() {
	o.Username.Set(nil)
}

// UnsetUsername ensures that no value is present for Username, not even an explicit nil
func (o *RemoteAdapterHost) UnsetUsername() {
	o.Username.Unset()
}

// GetHostType returns the HostType field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RemoteAdapterHost) GetHostType() string {
	if o == nil || o.HostType.Get() == nil {
		var ret string
		return ret
	}
	return *o.HostType.Get()
}

// GetHostTypeOk returns a tuple with the HostType field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RemoteAdapterHost) GetHostTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.HostType.Get(), o.HostType.IsSet()
}

// HasHostType returns a boolean if a field has been set.
func (o *RemoteAdapterHost) HasHostType() bool {
	if o != nil && o.HostType.IsSet() {
		return true
	}

	return false
}

// SetHostType gets a reference to the given NullableString and assigns it to the HostType field.
func (o *RemoteAdapterHost) SetHostType(v string) {
	o.HostType.Set(&v)
}
// SetHostTypeNil sets the value for HostType to be an explicit nil
func (o *RemoteAdapterHost) SetHostTypeNil() {
	o.HostType.Set(nil)
}

// UnsetHostType ensures that no value is present for HostType, not even an explicit nil
func (o *RemoteAdapterHost) UnsetHostType() {
	o.HostType.Unset()
}

// GetIncrementalBackupScript returns the IncrementalBackupScript field value if set, zero value otherwise.
func (o *RemoteAdapterHost) GetIncrementalBackupScript() CommonPreBackupScriptParams {
	if o == nil || o.IncrementalBackupScript == nil {
		var ret CommonPreBackupScriptParams
		return ret
	}
	return *o.IncrementalBackupScript
}

// GetIncrementalBackupScriptOk returns a tuple with the IncrementalBackupScript field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RemoteAdapterHost) GetIncrementalBackupScriptOk() (*CommonPreBackupScriptParams, bool) {
	if o == nil || o.IncrementalBackupScript == nil {
		return nil, false
	}
	return o.IncrementalBackupScript, true
}

// HasIncrementalBackupScript returns a boolean if a field has been set.
func (o *RemoteAdapterHost) HasIncrementalBackupScript() bool {
	if o != nil && o.IncrementalBackupScript != nil {
		return true
	}

	return false
}

// SetIncrementalBackupScript gets a reference to the given CommonPreBackupScriptParams and assigns it to the IncrementalBackupScript field.
func (o *RemoteAdapterHost) SetIncrementalBackupScript(v CommonPreBackupScriptParams) {
	o.IncrementalBackupScript = &v
}

// GetFullBackupScript returns the FullBackupScript field value if set, zero value otherwise.
func (o *RemoteAdapterHost) GetFullBackupScript() CommonPreBackupScriptParams {
	if o == nil || o.FullBackupScript == nil {
		var ret CommonPreBackupScriptParams
		return ret
	}
	return *o.FullBackupScript
}

// GetFullBackupScriptOk returns a tuple with the FullBackupScript field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RemoteAdapterHost) GetFullBackupScriptOk() (*CommonPreBackupScriptParams, bool) {
	if o == nil || o.FullBackupScript == nil {
		return nil, false
	}
	return o.FullBackupScript, true
}

// HasFullBackupScript returns a boolean if a field has been set.
func (o *RemoteAdapterHost) HasFullBackupScript() bool {
	if o != nil && o.FullBackupScript != nil {
		return true
	}

	return false
}

// SetFullBackupScript gets a reference to the given CommonPreBackupScriptParams and assigns it to the FullBackupScript field.
func (o *RemoteAdapterHost) SetFullBackupScript(v CommonPreBackupScriptParams) {
	o.FullBackupScript = &v
}

// GetLogBackupScript returns the LogBackupScript field value if set, zero value otherwise.
func (o *RemoteAdapterHost) GetLogBackupScript() CommonPreBackupScriptParams {
	if o == nil || o.LogBackupScript == nil {
		var ret CommonPreBackupScriptParams
		return ret
	}
	return *o.LogBackupScript
}

// GetLogBackupScriptOk returns a tuple with the LogBackupScript field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RemoteAdapterHost) GetLogBackupScriptOk() (*CommonPreBackupScriptParams, bool) {
	if o == nil || o.LogBackupScript == nil {
		return nil, false
	}
	return o.LogBackupScript, true
}

// HasLogBackupScript returns a boolean if a field has been set.
func (o *RemoteAdapterHost) HasLogBackupScript() bool {
	if o != nil && o.LogBackupScript != nil {
		return true
	}

	return false
}

// SetLogBackupScript gets a reference to the given CommonPreBackupScriptParams and assigns it to the LogBackupScript field.
func (o *RemoteAdapterHost) SetLogBackupScript(v CommonPreBackupScriptParams) {
	o.LogBackupScript = &v
}

func (o RemoteAdapterHost) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Hostname.IsSet() {
		toSerialize["hostname"] = o.Hostname.Get()
	}
	if o.Username.IsSet() {
		toSerialize["username"] = o.Username.Get()
	}
	if o.HostType.IsSet() {
		toSerialize["hostType"] = o.HostType.Get()
	}
	if o.IncrementalBackupScript != nil {
		toSerialize["incrementalBackupScript"] = o.IncrementalBackupScript
	}
	if o.FullBackupScript != nil {
		toSerialize["fullBackupScript"] = o.FullBackupScript
	}
	if o.LogBackupScript != nil {
		toSerialize["logBackupScript"] = o.LogBackupScript
	}
	return json.Marshal(toSerialize)
}

type NullableRemoteAdapterHost struct {
	value *RemoteAdapterHost
	isSet bool
}

func (v NullableRemoteAdapterHost) Get() *RemoteAdapterHost {
	return v.value
}

func (v *NullableRemoteAdapterHost) Set(val *RemoteAdapterHost) {
	v.value = val
	v.isSet = true
}

func (v NullableRemoteAdapterHost) IsSet() bool {
	return v.isSet
}

func (v *NullableRemoteAdapterHost) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRemoteAdapterHost(val *RemoteAdapterHost) *NullableRemoteAdapterHost {
	return &NullableRemoteAdapterHost{value: val, isSet: true}
}

func (v NullableRemoteAdapterHost) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRemoteAdapterHost) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


