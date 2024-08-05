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

// checks if the UpdateSMTPParams type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateSMTPParams{}

// UpdateSMTPParams Specifies the parameters to update cluster SMTP configuration.
type UpdateSMTPParams struct {
	// Specifies the IP address or the FQDN of the SMTP server.
	Hostname string `json:"hostname"`
	// Specifies if the SMTP configuration is active.
	IsActive NullableBool `json:"isActive,omitempty"`
	// Specifies the SMTP port. Usually 465 or 587. For authenticated connection, it is generally 587.
	Port int32 `json:"port"`
	// This is set to true when the SMTP server uses SSL/TLS without supporting STARTTLS. Typically, this is used for port 465.
	UseSSL NullableBool `json:"useSSL,omitempty"`
	// Specifies the username which will be used to connect to the SMTP server. If username is not specified, then it would imply that SMTP server is set up for unauthenticated access.
	Username NullableString `json:"username,omitempty"`
	// Specifies the password of the SMTP user. This is required if username is specified in the request.
	Password NullableString `json:"password,omitempty"`
}

type _UpdateSMTPParams UpdateSMTPParams

// NewUpdateSMTPParams instantiates a new UpdateSMTPParams object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateSMTPParams(hostname string, port int32) *UpdateSMTPParams {
	this := UpdateSMTPParams{}
	this.Hostname = hostname
	var isActive bool = true
	this.IsActive = *NewNullableBool(&isActive)
	this.Port = port
	var useSSL bool = false
	this.UseSSL = *NewNullableBool(&useSSL)
	return &this
}

// NewUpdateSMTPParamsWithDefaults instantiates a new UpdateSMTPParams object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateSMTPParamsWithDefaults() *UpdateSMTPParams {
	this := UpdateSMTPParams{}
	var isActive bool = true
	this.IsActive = *NewNullableBool(&isActive)
	var useSSL bool = false
	this.UseSSL = *NewNullableBool(&useSSL)
	return &this
}

// GetHostname returns the Hostname field value
func (o *UpdateSMTPParams) GetHostname() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Hostname
}

// GetHostnameOk returns a tuple with the Hostname field value
// and a boolean to check if the value has been set.
func (o *UpdateSMTPParams) GetHostnameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Hostname, true
}

// SetHostname sets field value
func (o *UpdateSMTPParams) SetHostname(v string) {
	o.Hostname = v
}

// GetIsActive returns the IsActive field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UpdateSMTPParams) GetIsActive() bool {
	if o == nil || IsNil(o.IsActive.Get()) {
		var ret bool
		return ret
	}
	return *o.IsActive.Get()
}

// GetIsActiveOk returns a tuple with the IsActive field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UpdateSMTPParams) GetIsActiveOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.IsActive.Get(), o.IsActive.IsSet()
}

// HasIsActive returns a boolean if a field has been set.
func (o *UpdateSMTPParams) HasIsActive() bool {
	if o != nil && o.IsActive.IsSet() {
		return true
	}

	return false
}

// SetIsActive gets a reference to the given NullableBool and assigns it to the IsActive field.
func (o *UpdateSMTPParams) SetIsActive(v bool) {
	o.IsActive.Set(&v)
}
// SetIsActiveNil sets the value for IsActive to be an explicit nil
func (o *UpdateSMTPParams) SetIsActiveNil() {
	o.IsActive.Set(nil)
}

// UnsetIsActive ensures that no value is present for IsActive, not even an explicit nil
func (o *UpdateSMTPParams) UnsetIsActive() {
	o.IsActive.Unset()
}

// GetPort returns the Port field value
func (o *UpdateSMTPParams) GetPort() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Port
}

// GetPortOk returns a tuple with the Port field value
// and a boolean to check if the value has been set.
func (o *UpdateSMTPParams) GetPortOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Port, true
}

// SetPort sets field value
func (o *UpdateSMTPParams) SetPort(v int32) {
	o.Port = v
}

// GetUseSSL returns the UseSSL field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UpdateSMTPParams) GetUseSSL() bool {
	if o == nil || IsNil(o.UseSSL.Get()) {
		var ret bool
		return ret
	}
	return *o.UseSSL.Get()
}

// GetUseSSLOk returns a tuple with the UseSSL field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UpdateSMTPParams) GetUseSSLOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.UseSSL.Get(), o.UseSSL.IsSet()
}

// HasUseSSL returns a boolean if a field has been set.
func (o *UpdateSMTPParams) HasUseSSL() bool {
	if o != nil && o.UseSSL.IsSet() {
		return true
	}

	return false
}

// SetUseSSL gets a reference to the given NullableBool and assigns it to the UseSSL field.
func (o *UpdateSMTPParams) SetUseSSL(v bool) {
	o.UseSSL.Set(&v)
}
// SetUseSSLNil sets the value for UseSSL to be an explicit nil
func (o *UpdateSMTPParams) SetUseSSLNil() {
	o.UseSSL.Set(nil)
}

// UnsetUseSSL ensures that no value is present for UseSSL, not even an explicit nil
func (o *UpdateSMTPParams) UnsetUseSSL() {
	o.UseSSL.Unset()
}

// GetUsername returns the Username field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UpdateSMTPParams) GetUsername() string {
	if o == nil || IsNil(o.Username.Get()) {
		var ret string
		return ret
	}
	return *o.Username.Get()
}

// GetUsernameOk returns a tuple with the Username field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UpdateSMTPParams) GetUsernameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Username.Get(), o.Username.IsSet()
}

// HasUsername returns a boolean if a field has been set.
func (o *UpdateSMTPParams) HasUsername() bool {
	if o != nil && o.Username.IsSet() {
		return true
	}

	return false
}

// SetUsername gets a reference to the given NullableString and assigns it to the Username field.
func (o *UpdateSMTPParams) SetUsername(v string) {
	o.Username.Set(&v)
}
// SetUsernameNil sets the value for Username to be an explicit nil
func (o *UpdateSMTPParams) SetUsernameNil() {
	o.Username.Set(nil)
}

// UnsetUsername ensures that no value is present for Username, not even an explicit nil
func (o *UpdateSMTPParams) UnsetUsername() {
	o.Username.Unset()
}

// GetPassword returns the Password field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UpdateSMTPParams) GetPassword() string {
	if o == nil || IsNil(o.Password.Get()) {
		var ret string
		return ret
	}
	return *o.Password.Get()
}

// GetPasswordOk returns a tuple with the Password field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UpdateSMTPParams) GetPasswordOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Password.Get(), o.Password.IsSet()
}

// HasPassword returns a boolean if a field has been set.
func (o *UpdateSMTPParams) HasPassword() bool {
	if o != nil && o.Password.IsSet() {
		return true
	}

	return false
}

// SetPassword gets a reference to the given NullableString and assigns it to the Password field.
func (o *UpdateSMTPParams) SetPassword(v string) {
	o.Password.Set(&v)
}
// SetPasswordNil sets the value for Password to be an explicit nil
func (o *UpdateSMTPParams) SetPasswordNil() {
	o.Password.Set(nil)
}

// UnsetPassword ensures that no value is present for Password, not even an explicit nil
func (o *UpdateSMTPParams) UnsetPassword() {
	o.Password.Unset()
}

func (o UpdateSMTPParams) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateSMTPParams) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["hostname"] = o.Hostname
	if o.IsActive.IsSet() {
		toSerialize["isActive"] = o.IsActive.Get()
	}
	toSerialize["port"] = o.Port
	if o.UseSSL.IsSet() {
		toSerialize["useSSL"] = o.UseSSL.Get()
	}
	if o.Username.IsSet() {
		toSerialize["username"] = o.Username.Get()
	}
	if o.Password.IsSet() {
		toSerialize["password"] = o.Password.Get()
	}
	return toSerialize, nil
}

func (o *UpdateSMTPParams) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"hostname",
		"port",
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

	varUpdateSMTPParams := _UpdateSMTPParams{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varUpdateSMTPParams)

	if err != nil {
		return err
	}

	*o = UpdateSMTPParams(varUpdateSMTPParams)

	return err
}

type NullableUpdateSMTPParams struct {
	value *UpdateSMTPParams
	isSet bool
}

func (v NullableUpdateSMTPParams) Get() *UpdateSMTPParams {
	return v.value
}

func (v *NullableUpdateSMTPParams) Set(val *UpdateSMTPParams) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateSMTPParams) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateSMTPParams) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateSMTPParams(val *UpdateSMTPParams) *NullableUpdateSMTPParams {
	return &NullableUpdateSMTPParams{value: val, isSet: true}
}

func (v NullableUpdateSMTPParams) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateSMTPParams) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


