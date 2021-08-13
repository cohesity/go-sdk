/*
 * Cohesity REST API
 *
 * Cohesity API provides a RESTful interface to access the various data management operations on Cohesity cluster and Helios.
 *
 * API version: 2.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

// package onprem
package models

import (
	"encoding/json"
	. "github.com/cohesity/go-sdk/onprem/utils"
	"fmt"
)

var _ = NullableBool{}

// AwsCloudC2SParams Specifies the parameters which are specific to AWS related External Targets with Cloud Type C2S.
type AwsCloudC2SParams struct {
	// Specifies base url of the External Target.
	BaseURL NullableString `json:"baseURL"`
	// Specifies agency of the External Target.
	Agency NullableString `json:"agency"`
	// Specifies mission of the External Target
	Mission NullableString `json:"mission"`
	// Specifies role of the External Target
	Role NullableString `json:"role"`
	// Specifies client certificate password of the External Target
	ClientCertificatePassword NullableString `json:"clientCertificatePassword"`
	// Specifies server CA trusted certificate of the External Target
	ServerCATrustedCertificate NullableString `json:"serverCATrustedCertificate"`
	// Specifies client certificate of the External Target
	ClientCertificate NullableString `json:"clientCertificate"`
	// Specifies client private key of the External Target
	ClientPrivateKey NullableString `json:"clientPrivateKey"`
}

// NewAwsCloudC2SParams instantiates a new AwsCloudC2SParams object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAwsCloudC2SParams(baseURL NullableString, agency NullableString, mission NullableString, role NullableString, clientCertificatePassword NullableString, serverCATrustedCertificate NullableString, clientCertificate NullableString, clientPrivateKey NullableString) *AwsCloudC2SParams {
	this := AwsCloudC2SParams{}
	this.BaseURL = baseURL
	this.Agency = agency
	this.Mission = mission
	this.Role = role
	this.ClientCertificatePassword = clientCertificatePassword
	this.ServerCATrustedCertificate = serverCATrustedCertificate
	this.ClientCertificate = clientCertificate
	this.ClientPrivateKey = clientPrivateKey
	return &this
}

// NewAwsCloudC2SParamsWithDefaults instantiates a new AwsCloudC2SParams object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAwsCloudC2SParamsWithDefaults() *AwsCloudC2SParams {
	this := AwsCloudC2SParams{}
	return &this
}

// GetBaseURL returns the BaseURL field value
// If the value is explicit nil, the zero value for string will be returned
func (o *AwsCloudC2SParams) GetBaseURL() string {
	if o == nil || o.BaseURL.Get() == nil {
		var ret string
		return ret
	}

	return *o.BaseURL.Get()
}

// GetBaseURLOk returns a tuple with the BaseURL field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AwsCloudC2SParams) GetBaseURLOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.BaseURL.Get(), o.BaseURL.IsSet()
}

// SetBaseURL sets field value
func (o *AwsCloudC2SParams) SetBaseURL(v string) {
	o.BaseURL.Set(&v)
}

// GetAgency returns the Agency field value
// If the value is explicit nil, the zero value for string will be returned
func (o *AwsCloudC2SParams) GetAgency() string {
	if o == nil || o.Agency.Get() == nil {
		var ret string
		return ret
	}

	return *o.Agency.Get()
}

// GetAgencyOk returns a tuple with the Agency field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AwsCloudC2SParams) GetAgencyOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Agency.Get(), o.Agency.IsSet()
}

// SetAgency sets field value
func (o *AwsCloudC2SParams) SetAgency(v string) {
	o.Agency.Set(&v)
}

// GetMission returns the Mission field value
// If the value is explicit nil, the zero value for string will be returned
func (o *AwsCloudC2SParams) GetMission() string {
	if o == nil || o.Mission.Get() == nil {
		var ret string
		return ret
	}

	return *o.Mission.Get()
}

// GetMissionOk returns a tuple with the Mission field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AwsCloudC2SParams) GetMissionOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Mission.Get(), o.Mission.IsSet()
}

// SetMission sets field value
func (o *AwsCloudC2SParams) SetMission(v string) {
	o.Mission.Set(&v)
}

// GetRole returns the Role field value
// If the value is explicit nil, the zero value for string will be returned
func (o *AwsCloudC2SParams) GetRole() string {
	if o == nil || o.Role.Get() == nil {
		var ret string
		return ret
	}

	return *o.Role.Get()
}

// GetRoleOk returns a tuple with the Role field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AwsCloudC2SParams) GetRoleOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Role.Get(), o.Role.IsSet()
}

// SetRole sets field value
func (o *AwsCloudC2SParams) SetRole(v string) {
	o.Role.Set(&v)
}

// GetClientCertificatePassword returns the ClientCertificatePassword field value
// If the value is explicit nil, the zero value for string will be returned
func (o *AwsCloudC2SParams) GetClientCertificatePassword() string {
	if o == nil || o.ClientCertificatePassword.Get() == nil {
		var ret string
		return ret
	}

	return *o.ClientCertificatePassword.Get()
}

// GetClientCertificatePasswordOk returns a tuple with the ClientCertificatePassword field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AwsCloudC2SParams) GetClientCertificatePasswordOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.ClientCertificatePassword.Get(), o.ClientCertificatePassword.IsSet()
}

// SetClientCertificatePassword sets field value
func (o *AwsCloudC2SParams) SetClientCertificatePassword(v string) {
	o.ClientCertificatePassword.Set(&v)
}

// GetServerCATrustedCertificate returns the ServerCATrustedCertificate field value
// If the value is explicit nil, the zero value for string will be returned
func (o *AwsCloudC2SParams) GetServerCATrustedCertificate() string {
	if o == nil || o.ServerCATrustedCertificate.Get() == nil {
		var ret string
		return ret
	}

	return *o.ServerCATrustedCertificate.Get()
}

// GetServerCATrustedCertificateOk returns a tuple with the ServerCATrustedCertificate field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AwsCloudC2SParams) GetServerCATrustedCertificateOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.ServerCATrustedCertificate.Get(), o.ServerCATrustedCertificate.IsSet()
}

// SetServerCATrustedCertificate sets field value
func (o *AwsCloudC2SParams) SetServerCATrustedCertificate(v string) {
	o.ServerCATrustedCertificate.Set(&v)
}

// GetClientCertificate returns the ClientCertificate field value
// If the value is explicit nil, the zero value for string will be returned
func (o *AwsCloudC2SParams) GetClientCertificate() string {
	if o == nil || o.ClientCertificate.Get() == nil {
		var ret string
		return ret
	}

	return *o.ClientCertificate.Get()
}

// GetClientCertificateOk returns a tuple with the ClientCertificate field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AwsCloudC2SParams) GetClientCertificateOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.ClientCertificate.Get(), o.ClientCertificate.IsSet()
}

// SetClientCertificate sets field value
func (o *AwsCloudC2SParams) SetClientCertificate(v string) {
	o.ClientCertificate.Set(&v)
}

// GetClientPrivateKey returns the ClientPrivateKey field value
// If the value is explicit nil, the zero value for string will be returned
func (o *AwsCloudC2SParams) GetClientPrivateKey() string {
	if o == nil || o.ClientPrivateKey.Get() == nil {
		var ret string
		return ret
	}

	return *o.ClientPrivateKey.Get()
}

// GetClientPrivateKeyOk returns a tuple with the ClientPrivateKey field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AwsCloudC2SParams) GetClientPrivateKeyOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.ClientPrivateKey.Get(), o.ClientPrivateKey.IsSet()
}

// SetClientPrivateKey sets field value
func (o *AwsCloudC2SParams) SetClientPrivateKey(v string) {
	o.ClientPrivateKey.Set(&v)
}

func (o AwsCloudC2SParams) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["baseURL"] = o.BaseURL.Get()
	}
	if true {
		toSerialize["agency"] = o.Agency.Get()
	}
	if true {
		toSerialize["mission"] = o.Mission.Get()
	}
	if true {
		toSerialize["role"] = o.Role.Get()
	}
	if true {
		toSerialize["clientCertificatePassword"] = o.ClientCertificatePassword.Get()
	}
	if true {
		toSerialize["serverCATrustedCertificate"] = o.ServerCATrustedCertificate.Get()
	}
	if true {
		toSerialize["clientCertificate"] = o.ClientCertificate.Get()
	}
	if true {
		toSerialize["clientPrivateKey"] = o.ClientPrivateKey.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableAwsCloudC2SParams struct {
	value *AwsCloudC2SParams
	isSet bool
}

func (v NullableAwsCloudC2SParams) Get() *AwsCloudC2SParams {
	return v.value
}

func (v *NullableAwsCloudC2SParams) Set(val *AwsCloudC2SParams) {
	v.value = val
	v.isSet = true
}

func (v NullableAwsCloudC2SParams) IsSet() bool {
	return v.isSet
}

func (v *NullableAwsCloudC2SParams) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAwsCloudC2SParams(val *AwsCloudC2SParams) *NullableAwsCloudC2SParams {
	return &NullableAwsCloudC2SParams{value: val, isSet: true}
}

func (v NullableAwsCloudC2SParams) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAwsCloudC2SParams) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}




func (o AwsCloudC2SParams) Print() {
		if byteArray, err := o.MarshalJSON(); err != nil {
				fmt.Println(err)
		} else {
				fmt.Println(string(byteArray))
		}
}