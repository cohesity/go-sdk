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

// HiveSourceRegistrationUpdateParams Specifies parameters to update registration of Hive source.
type HiveSourceRegistrationUpdateParams struct {
	// IP or hostname of any host from which the Hive configuration file hive-site.xml can be read.
	Host NullableString `json:"host,omitempty"`
	// The directory containing the hive-site.xml.
	ConfigurationDirectory NullableString `json:"configurationDirectory,omitempty"`
	SshPasswordCredentials NullableHbaseSourceRegistrationParamsAllOfSshPasswordCredentials `json:"sshPasswordCredentials,omitempty"`
	SshPrivateKeyCredentials NullableHbaseSourceRegistrationParamsAllOfSshPrivateKeyCredentials `json:"sshPrivateKeyCredentials,omitempty"`
	// The kerberos principal to be used to connect to this Hive source.
	KerberosPrincipal NullableString `json:"kerberosPrincipal,omitempty"`
}

// NewHiveSourceRegistrationUpdateParams instantiates a new HiveSourceRegistrationUpdateParams object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHiveSourceRegistrationUpdateParams() *HiveSourceRegistrationUpdateParams {
	this := HiveSourceRegistrationUpdateParams{}
	return &this
}

// NewHiveSourceRegistrationUpdateParamsWithDefaults instantiates a new HiveSourceRegistrationUpdateParams object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHiveSourceRegistrationUpdateParamsWithDefaults() *HiveSourceRegistrationUpdateParams {
	this := HiveSourceRegistrationUpdateParams{}
	return &this
}

// GetHost returns the Host field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *HiveSourceRegistrationUpdateParams) GetHost() string {
	if o == nil || o.Host.Get() == nil {
		var ret string
		return ret
	}
	return *o.Host.Get()
}

// GetHostOk returns a tuple with the Host field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HiveSourceRegistrationUpdateParams) GetHostOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Host.Get(), o.Host.IsSet()
}

// HasHost returns a boolean if a field has been set.
func (o *HiveSourceRegistrationUpdateParams) HasHost() bool {
	if o != nil && o.Host.IsSet() {
		return true
	}

	return false
}

// SetHost gets a reference to the given NullableString and assigns it to the Host field.
func (o *HiveSourceRegistrationUpdateParams) SetHost(v string) {
	o.Host.Set(&v)
}
// SetHostNil sets the value for Host to be an explicit nil
func (o *HiveSourceRegistrationUpdateParams) SetHostNil() {
	o.Host.Set(nil)
}

// UnsetHost ensures that no value is present for Host, not even an explicit nil
func (o *HiveSourceRegistrationUpdateParams) UnsetHost() {
	o.Host.Unset()
}

// GetConfigurationDirectory returns the ConfigurationDirectory field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *HiveSourceRegistrationUpdateParams) GetConfigurationDirectory() string {
	if o == nil || o.ConfigurationDirectory.Get() == nil {
		var ret string
		return ret
	}
	return *o.ConfigurationDirectory.Get()
}

// GetConfigurationDirectoryOk returns a tuple with the ConfigurationDirectory field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HiveSourceRegistrationUpdateParams) GetConfigurationDirectoryOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.ConfigurationDirectory.Get(), o.ConfigurationDirectory.IsSet()
}

// HasConfigurationDirectory returns a boolean if a field has been set.
func (o *HiveSourceRegistrationUpdateParams) HasConfigurationDirectory() bool {
	if o != nil && o.ConfigurationDirectory.IsSet() {
		return true
	}

	return false
}

// SetConfigurationDirectory gets a reference to the given NullableString and assigns it to the ConfigurationDirectory field.
func (o *HiveSourceRegistrationUpdateParams) SetConfigurationDirectory(v string) {
	o.ConfigurationDirectory.Set(&v)
}
// SetConfigurationDirectoryNil sets the value for ConfigurationDirectory to be an explicit nil
func (o *HiveSourceRegistrationUpdateParams) SetConfigurationDirectoryNil() {
	o.ConfigurationDirectory.Set(nil)
}

// UnsetConfigurationDirectory ensures that no value is present for ConfigurationDirectory, not even an explicit nil
func (o *HiveSourceRegistrationUpdateParams) UnsetConfigurationDirectory() {
	o.ConfigurationDirectory.Unset()
}

// GetSshPasswordCredentials returns the SshPasswordCredentials field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *HiveSourceRegistrationUpdateParams) GetSshPasswordCredentials() HbaseSourceRegistrationParamsAllOfSshPasswordCredentials {
	if o == nil || o.SshPasswordCredentials.Get() == nil {
		var ret HbaseSourceRegistrationParamsAllOfSshPasswordCredentials
		return ret
	}
	return *o.SshPasswordCredentials.Get()
}

// GetSshPasswordCredentialsOk returns a tuple with the SshPasswordCredentials field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HiveSourceRegistrationUpdateParams) GetSshPasswordCredentialsOk() (*HbaseSourceRegistrationParamsAllOfSshPasswordCredentials, bool) {
	if o == nil  {
		return nil, false
	}
	return o.SshPasswordCredentials.Get(), o.SshPasswordCredentials.IsSet()
}

// HasSshPasswordCredentials returns a boolean if a field has been set.
func (o *HiveSourceRegistrationUpdateParams) HasSshPasswordCredentials() bool {
	if o != nil && o.SshPasswordCredentials.IsSet() {
		return true
	}

	return false
}

// SetSshPasswordCredentials gets a reference to the given NullableHbaseSourceRegistrationParamsAllOfSshPasswordCredentials and assigns it to the SshPasswordCredentials field.
func (o *HiveSourceRegistrationUpdateParams) SetSshPasswordCredentials(v HbaseSourceRegistrationParamsAllOfSshPasswordCredentials) {
	o.SshPasswordCredentials.Set(&v)
}
// SetSshPasswordCredentialsNil sets the value for SshPasswordCredentials to be an explicit nil
func (o *HiveSourceRegistrationUpdateParams) SetSshPasswordCredentialsNil() {
	o.SshPasswordCredentials.Set(nil)
}

// UnsetSshPasswordCredentials ensures that no value is present for SshPasswordCredentials, not even an explicit nil
func (o *HiveSourceRegistrationUpdateParams) UnsetSshPasswordCredentials() {
	o.SshPasswordCredentials.Unset()
}

// GetSshPrivateKeyCredentials returns the SshPrivateKeyCredentials field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *HiveSourceRegistrationUpdateParams) GetSshPrivateKeyCredentials() HbaseSourceRegistrationParamsAllOfSshPrivateKeyCredentials {
	if o == nil || o.SshPrivateKeyCredentials.Get() == nil {
		var ret HbaseSourceRegistrationParamsAllOfSshPrivateKeyCredentials
		return ret
	}
	return *o.SshPrivateKeyCredentials.Get()
}

// GetSshPrivateKeyCredentialsOk returns a tuple with the SshPrivateKeyCredentials field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HiveSourceRegistrationUpdateParams) GetSshPrivateKeyCredentialsOk() (*HbaseSourceRegistrationParamsAllOfSshPrivateKeyCredentials, bool) {
	if o == nil  {
		return nil, false
	}
	return o.SshPrivateKeyCredentials.Get(), o.SshPrivateKeyCredentials.IsSet()
}

// HasSshPrivateKeyCredentials returns a boolean if a field has been set.
func (o *HiveSourceRegistrationUpdateParams) HasSshPrivateKeyCredentials() bool {
	if o != nil && o.SshPrivateKeyCredentials.IsSet() {
		return true
	}

	return false
}

// SetSshPrivateKeyCredentials gets a reference to the given NullableHbaseSourceRegistrationParamsAllOfSshPrivateKeyCredentials and assigns it to the SshPrivateKeyCredentials field.
func (o *HiveSourceRegistrationUpdateParams) SetSshPrivateKeyCredentials(v HbaseSourceRegistrationParamsAllOfSshPrivateKeyCredentials) {
	o.SshPrivateKeyCredentials.Set(&v)
}
// SetSshPrivateKeyCredentialsNil sets the value for SshPrivateKeyCredentials to be an explicit nil
func (o *HiveSourceRegistrationUpdateParams) SetSshPrivateKeyCredentialsNil() {
	o.SshPrivateKeyCredentials.Set(nil)
}

// UnsetSshPrivateKeyCredentials ensures that no value is present for SshPrivateKeyCredentials, not even an explicit nil
func (o *HiveSourceRegistrationUpdateParams) UnsetSshPrivateKeyCredentials() {
	o.SshPrivateKeyCredentials.Unset()
}

// GetKerberosPrincipal returns the KerberosPrincipal field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *HiveSourceRegistrationUpdateParams) GetKerberosPrincipal() string {
	if o == nil || o.KerberosPrincipal.Get() == nil {
		var ret string
		return ret
	}
	return *o.KerberosPrincipal.Get()
}

// GetKerberosPrincipalOk returns a tuple with the KerberosPrincipal field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HiveSourceRegistrationUpdateParams) GetKerberosPrincipalOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.KerberosPrincipal.Get(), o.KerberosPrincipal.IsSet()
}

// HasKerberosPrincipal returns a boolean if a field has been set.
func (o *HiveSourceRegistrationUpdateParams) HasKerberosPrincipal() bool {
	if o != nil && o.KerberosPrincipal.IsSet() {
		return true
	}

	return false
}

// SetKerberosPrincipal gets a reference to the given NullableString and assigns it to the KerberosPrincipal field.
func (o *HiveSourceRegistrationUpdateParams) SetKerberosPrincipal(v string) {
	o.KerberosPrincipal.Set(&v)
}
// SetKerberosPrincipalNil sets the value for KerberosPrincipal to be an explicit nil
func (o *HiveSourceRegistrationUpdateParams) SetKerberosPrincipalNil() {
	o.KerberosPrincipal.Set(nil)
}

// UnsetKerberosPrincipal ensures that no value is present for KerberosPrincipal, not even an explicit nil
func (o *HiveSourceRegistrationUpdateParams) UnsetKerberosPrincipal() {
	o.KerberosPrincipal.Unset()
}

func (o HiveSourceRegistrationUpdateParams) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Host.IsSet() {
		toSerialize["host"] = o.Host.Get()
	}
	if o.ConfigurationDirectory.IsSet() {
		toSerialize["configurationDirectory"] = o.ConfigurationDirectory.Get()
	}
	if o.SshPasswordCredentials.IsSet() {
		toSerialize["sshPasswordCredentials"] = o.SshPasswordCredentials.Get()
	}
	if o.SshPrivateKeyCredentials.IsSet() {
		toSerialize["sshPrivateKeyCredentials"] = o.SshPrivateKeyCredentials.Get()
	}
	if o.KerberosPrincipal.IsSet() {
		toSerialize["kerberosPrincipal"] = o.KerberosPrincipal.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableHiveSourceRegistrationUpdateParams struct {
	value *HiveSourceRegistrationUpdateParams
	isSet bool
}

func (v NullableHiveSourceRegistrationUpdateParams) Get() *HiveSourceRegistrationUpdateParams {
	return v.value
}

func (v *NullableHiveSourceRegistrationUpdateParams) Set(val *HiveSourceRegistrationUpdateParams) {
	v.value = val
	v.isSet = true
}

func (v NullableHiveSourceRegistrationUpdateParams) IsSet() bool {
	return v.isSet
}

func (v *NullableHiveSourceRegistrationUpdateParams) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHiveSourceRegistrationUpdateParams(val *HiveSourceRegistrationUpdateParams) *NullableHiveSourceRegistrationUpdateParams {
	return &NullableHiveSourceRegistrationUpdateParams{value: val, isSet: true}
}

func (v NullableHiveSourceRegistrationUpdateParams) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHiveSourceRegistrationUpdateParams) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


