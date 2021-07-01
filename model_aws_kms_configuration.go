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

// AwsKmsConfiguration struct for AwsKmsConfiguration
type AwsKmsConfiguration struct {
	// Access key id needed to access the cloud account. When update cluster config, should encrypte accessKeyId with cluster ID.
	AccessKeyId NullableString `json:"accessKeyId,omitempty"`
	// Specify the ca certificate path.
	CaCertificate NullableString `json:"caCertificate,omitempty"`
	// The string alias of the CMK.
	CmkAlias NullableString `json:"cmkAlias,omitempty"`
	// The Amazon Resource Number of AWS Customer Managed Key.
	CmkArn NullableString `json:"cmkArn,omitempty"`
	// AWS keyId, and alias. Only need one of them to connect AWS. Alias is better, because keyId maybe rotated by AWS. The unique key id of the CMK.
	CmkKeyId NullableString `json:"cmkKeyId,omitempty"`
	// AWS region, e.g. us-east-1, us-west-2, for the AWS Glacier service to be used to authenticate resources within this region by the configured AWS account.
	Region NullableString `json:"region,omitempty"`
	// Secret access key needed to access the cloud account. This is encrypted with the cluster id.
	SecretAccessKey NullableString `json:"secretAccessKey,omitempty"`
	// Specify whether to verify SSL when connect with AWS KMS. Default is true.
	VerifySSL NullableBool `json:"verifySSL,omitempty"`
}

// NewAwsKmsConfiguration instantiates a new AwsKmsConfiguration object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAwsKmsConfiguration() *AwsKmsConfiguration {
	this := AwsKmsConfiguration{}
	return &this
}

// NewAwsKmsConfigurationWithDefaults instantiates a new AwsKmsConfiguration object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAwsKmsConfigurationWithDefaults() *AwsKmsConfiguration {
	this := AwsKmsConfiguration{}
	return &this
}

// GetAccessKeyId returns the AccessKeyId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AwsKmsConfiguration) GetAccessKeyId() string {
	if o == nil || o.AccessKeyId.Get() == nil {
		var ret string
		return ret
	}
	return *o.AccessKeyId.Get()
}

// GetAccessKeyIdOk returns a tuple with the AccessKeyId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AwsKmsConfiguration) GetAccessKeyIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.AccessKeyId.Get(), o.AccessKeyId.IsSet()
}

// HasAccessKeyId returns a boolean if a field has been set.
func (o *AwsKmsConfiguration) HasAccessKeyId() bool {
	if o != nil && o.AccessKeyId.IsSet() {
		return true
	}

	return false
}

// SetAccessKeyId gets a reference to the given NullableString and assigns it to the AccessKeyId field.
func (o *AwsKmsConfiguration) SetAccessKeyId(v string) {
	o.AccessKeyId.Set(&v)
}
// SetAccessKeyIdNil sets the value for AccessKeyId to be an explicit nil
func (o *AwsKmsConfiguration) SetAccessKeyIdNil() {
	o.AccessKeyId.Set(nil)
}

// UnsetAccessKeyId ensures that no value is present for AccessKeyId, not even an explicit nil
func (o *AwsKmsConfiguration) UnsetAccessKeyId() {
	o.AccessKeyId.Unset()
}

// GetCaCertificate returns the CaCertificate field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AwsKmsConfiguration) GetCaCertificate() string {
	if o == nil || o.CaCertificate.Get() == nil {
		var ret string
		return ret
	}
	return *o.CaCertificate.Get()
}

// GetCaCertificateOk returns a tuple with the CaCertificate field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AwsKmsConfiguration) GetCaCertificateOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.CaCertificate.Get(), o.CaCertificate.IsSet()
}

// HasCaCertificate returns a boolean if a field has been set.
func (o *AwsKmsConfiguration) HasCaCertificate() bool {
	if o != nil && o.CaCertificate.IsSet() {
		return true
	}

	return false
}

// SetCaCertificate gets a reference to the given NullableString and assigns it to the CaCertificate field.
func (o *AwsKmsConfiguration) SetCaCertificate(v string) {
	o.CaCertificate.Set(&v)
}
// SetCaCertificateNil sets the value for CaCertificate to be an explicit nil
func (o *AwsKmsConfiguration) SetCaCertificateNil() {
	o.CaCertificate.Set(nil)
}

// UnsetCaCertificate ensures that no value is present for CaCertificate, not even an explicit nil
func (o *AwsKmsConfiguration) UnsetCaCertificate() {
	o.CaCertificate.Unset()
}

// GetCmkAlias returns the CmkAlias field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AwsKmsConfiguration) GetCmkAlias() string {
	if o == nil || o.CmkAlias.Get() == nil {
		var ret string
		return ret
	}
	return *o.CmkAlias.Get()
}

// GetCmkAliasOk returns a tuple with the CmkAlias field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AwsKmsConfiguration) GetCmkAliasOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.CmkAlias.Get(), o.CmkAlias.IsSet()
}

// HasCmkAlias returns a boolean if a field has been set.
func (o *AwsKmsConfiguration) HasCmkAlias() bool {
	if o != nil && o.CmkAlias.IsSet() {
		return true
	}

	return false
}

// SetCmkAlias gets a reference to the given NullableString and assigns it to the CmkAlias field.
func (o *AwsKmsConfiguration) SetCmkAlias(v string) {
	o.CmkAlias.Set(&v)
}
// SetCmkAliasNil sets the value for CmkAlias to be an explicit nil
func (o *AwsKmsConfiguration) SetCmkAliasNil() {
	o.CmkAlias.Set(nil)
}

// UnsetCmkAlias ensures that no value is present for CmkAlias, not even an explicit nil
func (o *AwsKmsConfiguration) UnsetCmkAlias() {
	o.CmkAlias.Unset()
}

// GetCmkArn returns the CmkArn field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AwsKmsConfiguration) GetCmkArn() string {
	if o == nil || o.CmkArn.Get() == nil {
		var ret string
		return ret
	}
	return *o.CmkArn.Get()
}

// GetCmkArnOk returns a tuple with the CmkArn field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AwsKmsConfiguration) GetCmkArnOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.CmkArn.Get(), o.CmkArn.IsSet()
}

// HasCmkArn returns a boolean if a field has been set.
func (o *AwsKmsConfiguration) HasCmkArn() bool {
	if o != nil && o.CmkArn.IsSet() {
		return true
	}

	return false
}

// SetCmkArn gets a reference to the given NullableString and assigns it to the CmkArn field.
func (o *AwsKmsConfiguration) SetCmkArn(v string) {
	o.CmkArn.Set(&v)
}
// SetCmkArnNil sets the value for CmkArn to be an explicit nil
func (o *AwsKmsConfiguration) SetCmkArnNil() {
	o.CmkArn.Set(nil)
}

// UnsetCmkArn ensures that no value is present for CmkArn, not even an explicit nil
func (o *AwsKmsConfiguration) UnsetCmkArn() {
	o.CmkArn.Unset()
}

// GetCmkKeyId returns the CmkKeyId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AwsKmsConfiguration) GetCmkKeyId() string {
	if o == nil || o.CmkKeyId.Get() == nil {
		var ret string
		return ret
	}
	return *o.CmkKeyId.Get()
}

// GetCmkKeyIdOk returns a tuple with the CmkKeyId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AwsKmsConfiguration) GetCmkKeyIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.CmkKeyId.Get(), o.CmkKeyId.IsSet()
}

// HasCmkKeyId returns a boolean if a field has been set.
func (o *AwsKmsConfiguration) HasCmkKeyId() bool {
	if o != nil && o.CmkKeyId.IsSet() {
		return true
	}

	return false
}

// SetCmkKeyId gets a reference to the given NullableString and assigns it to the CmkKeyId field.
func (o *AwsKmsConfiguration) SetCmkKeyId(v string) {
	o.CmkKeyId.Set(&v)
}
// SetCmkKeyIdNil sets the value for CmkKeyId to be an explicit nil
func (o *AwsKmsConfiguration) SetCmkKeyIdNil() {
	o.CmkKeyId.Set(nil)
}

// UnsetCmkKeyId ensures that no value is present for CmkKeyId, not even an explicit nil
func (o *AwsKmsConfiguration) UnsetCmkKeyId() {
	o.CmkKeyId.Unset()
}

// GetRegion returns the Region field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AwsKmsConfiguration) GetRegion() string {
	if o == nil || o.Region.Get() == nil {
		var ret string
		return ret
	}
	return *o.Region.Get()
}

// GetRegionOk returns a tuple with the Region field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AwsKmsConfiguration) GetRegionOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Region.Get(), o.Region.IsSet()
}

// HasRegion returns a boolean if a field has been set.
func (o *AwsKmsConfiguration) HasRegion() bool {
	if o != nil && o.Region.IsSet() {
		return true
	}

	return false
}

// SetRegion gets a reference to the given NullableString and assigns it to the Region field.
func (o *AwsKmsConfiguration) SetRegion(v string) {
	o.Region.Set(&v)
}
// SetRegionNil sets the value for Region to be an explicit nil
func (o *AwsKmsConfiguration) SetRegionNil() {
	o.Region.Set(nil)
}

// UnsetRegion ensures that no value is present for Region, not even an explicit nil
func (o *AwsKmsConfiguration) UnsetRegion() {
	o.Region.Unset()
}

// GetSecretAccessKey returns the SecretAccessKey field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AwsKmsConfiguration) GetSecretAccessKey() string {
	if o == nil || o.SecretAccessKey.Get() == nil {
		var ret string
		return ret
	}
	return *o.SecretAccessKey.Get()
}

// GetSecretAccessKeyOk returns a tuple with the SecretAccessKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AwsKmsConfiguration) GetSecretAccessKeyOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.SecretAccessKey.Get(), o.SecretAccessKey.IsSet()
}

// HasSecretAccessKey returns a boolean if a field has been set.
func (o *AwsKmsConfiguration) HasSecretAccessKey() bool {
	if o != nil && o.SecretAccessKey.IsSet() {
		return true
	}

	return false
}

// SetSecretAccessKey gets a reference to the given NullableString and assigns it to the SecretAccessKey field.
func (o *AwsKmsConfiguration) SetSecretAccessKey(v string) {
	o.SecretAccessKey.Set(&v)
}
// SetSecretAccessKeyNil sets the value for SecretAccessKey to be an explicit nil
func (o *AwsKmsConfiguration) SetSecretAccessKeyNil() {
	o.SecretAccessKey.Set(nil)
}

// UnsetSecretAccessKey ensures that no value is present for SecretAccessKey, not even an explicit nil
func (o *AwsKmsConfiguration) UnsetSecretAccessKey() {
	o.SecretAccessKey.Unset()
}

// GetVerifySSL returns the VerifySSL field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AwsKmsConfiguration) GetVerifySSL() bool {
	if o == nil || o.VerifySSL.Get() == nil {
		var ret bool
		return ret
	}
	return *o.VerifySSL.Get()
}

// GetVerifySSLOk returns a tuple with the VerifySSL field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AwsKmsConfiguration) GetVerifySSLOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return o.VerifySSL.Get(), o.VerifySSL.IsSet()
}

// HasVerifySSL returns a boolean if a field has been set.
func (o *AwsKmsConfiguration) HasVerifySSL() bool {
	if o != nil && o.VerifySSL.IsSet() {
		return true
	}

	return false
}

// SetVerifySSL gets a reference to the given NullableBool and assigns it to the VerifySSL field.
func (o *AwsKmsConfiguration) SetVerifySSL(v bool) {
	o.VerifySSL.Set(&v)
}
// SetVerifySSLNil sets the value for VerifySSL to be an explicit nil
func (o *AwsKmsConfiguration) SetVerifySSLNil() {
	o.VerifySSL.Set(nil)
}

// UnsetVerifySSL ensures that no value is present for VerifySSL, not even an explicit nil
func (o *AwsKmsConfiguration) UnsetVerifySSL() {
	o.VerifySSL.Unset()
}

func (o AwsKmsConfiguration) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AccessKeyId.IsSet() {
		toSerialize["accessKeyId"] = o.AccessKeyId.Get()
	}
	if o.CaCertificate.IsSet() {
		toSerialize["caCertificate"] = o.CaCertificate.Get()
	}
	if o.CmkAlias.IsSet() {
		toSerialize["cmkAlias"] = o.CmkAlias.Get()
	}
	if o.CmkArn.IsSet() {
		toSerialize["cmkArn"] = o.CmkArn.Get()
	}
	if o.CmkKeyId.IsSet() {
		toSerialize["cmkKeyId"] = o.CmkKeyId.Get()
	}
	if o.Region.IsSet() {
		toSerialize["region"] = o.Region.Get()
	}
	if o.SecretAccessKey.IsSet() {
		toSerialize["secretAccessKey"] = o.SecretAccessKey.Get()
	}
	if o.VerifySSL.IsSet() {
		toSerialize["verifySSL"] = o.VerifySSL.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableAwsKmsConfiguration struct {
	value *AwsKmsConfiguration
	isSet bool
}

func (v NullableAwsKmsConfiguration) Get() *AwsKmsConfiguration {
	return v.value
}

func (v *NullableAwsKmsConfiguration) Set(val *AwsKmsConfiguration) {
	v.value = val
	v.isSet = true
}

func (v NullableAwsKmsConfiguration) IsSet() bool {
	return v.isSet
}

func (v *NullableAwsKmsConfiguration) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAwsKmsConfiguration(val *AwsKmsConfiguration) *NullableAwsKmsConfiguration {
	return &NullableAwsKmsConfiguration{value: val, isSet: true}
}

func (v NullableAwsKmsConfiguration) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAwsKmsConfiguration) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


