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

// Idp Specifies the Identity Provider Configuration.
type Idp struct {
	// Specifies the name of the vendor providing IDP service.
	Name NullableString `json:"name"`
	// Specifies a unique name for this IdP configuration.
	Domain NullableString `json:"domain"`
	// Specifies the SSO URL of the IdP service for the customer. This is the URL given by IdP when the customer created an account. For example, dev-332534.oktapreview.com.
	SsoUrl NullableString `json:"ssoUrl"`
	// Specifies the IdP provided Issuer ID for the app. For example, exkh1aov1nhHrgFhN0h7.
	IssuerId NullableString `json:"issuerId"`
	// Specifies the salesforce account ID linked to this IDP. Either TenantId or SfAccountId should be set for IdP.
	SfAccountId NullableString `json:"sfAccountId,omitempty"`
	// Specifies the certificate generated for the app by the IdP service when the Helios is registered as an app. This is required to verify the SAML response.
	Certificate NullableString `json:"certificate"`
	// Specifies a list of default roles assigned to an IdP user if rolesSamlAttributeName is not given.
	DefaultRoles []string `json:"defaultRoles,omitempty"`
	// Specifies a list of default clusterIds assigned to an IdP user if clustersSamlAttributeName is not given. 'All' must be specified to give access to all clusters.
	DefaultClusters []string `json:"defaultClusters,omitempty"`
	// Specifies whether to sign the SAML request or not. When it is set to true, SAML request will be signed. When it is set to false, SAML request is not signed. Default is false. Set this flag to true if the IdP site is configured to expect the SAML request from Helios signed. If this is set to true, users must get the Helios certificate and upload it on the IdP site.
	SignRequest NullableBool `json:"signRequest,omitempty"`
	// Specifies a flag to enable or disable this IdP service. When it is set to true, IdP service is enabled. When it is set to false, IdP service is disabled. Default value is true.
	IsEnabled NullableBool `json:"isEnabled,omitempty"`
	// Specifies the Tenant Id if the IdP is configured for a Tenant. Either TenantId or SfAccountId should be set for IdP.
	TenantId NullableString `json:"tenantId,omitempty"`
	// Specifies the Id of the IDP Configuration.
	Id NullableInt64 `json:"id,omitempty"`
}

// NewIdp instantiates a new Idp object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIdp(name NullableString, domain NullableString, ssoUrl NullableString, issuerId NullableString, certificate NullableString) *Idp {
	this := Idp{}
	this.Name = name
	this.Domain = domain
	this.SsoUrl = ssoUrl
	this.IssuerId = issuerId
	this.Certificate = certificate
	return &this
}

// NewIdpWithDefaults instantiates a new Idp object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIdpWithDefaults() *Idp {
	this := Idp{}
	return &this
}

// GetName returns the Name field value
// If the value is explicit nil, the zero value for string will be returned
func (o *Idp) GetName() string {
	if o == nil || o.Name.Get() == nil {
		var ret string
		return ret
	}

	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Idp) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// SetName sets field value
func (o *Idp) SetName(v string) {
	o.Name.Set(&v)
}

// GetDomain returns the Domain field value
// If the value is explicit nil, the zero value for string will be returned
func (o *Idp) GetDomain() string {
	if o == nil || o.Domain.Get() == nil {
		var ret string
		return ret
	}

	return *o.Domain.Get()
}

// GetDomainOk returns a tuple with the Domain field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Idp) GetDomainOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Domain.Get(), o.Domain.IsSet()
}

// SetDomain sets field value
func (o *Idp) SetDomain(v string) {
	o.Domain.Set(&v)
}

// GetSsoUrl returns the SsoUrl field value
// If the value is explicit nil, the zero value for string will be returned
func (o *Idp) GetSsoUrl() string {
	if o == nil || o.SsoUrl.Get() == nil {
		var ret string
		return ret
	}

	return *o.SsoUrl.Get()
}

// GetSsoUrlOk returns a tuple with the SsoUrl field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Idp) GetSsoUrlOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.SsoUrl.Get(), o.SsoUrl.IsSet()
}

// SetSsoUrl sets field value
func (o *Idp) SetSsoUrl(v string) {
	o.SsoUrl.Set(&v)
}

// GetIssuerId returns the IssuerId field value
// If the value is explicit nil, the zero value for string will be returned
func (o *Idp) GetIssuerId() string {
	if o == nil || o.IssuerId.Get() == nil {
		var ret string
		return ret
	}

	return *o.IssuerId.Get()
}

// GetIssuerIdOk returns a tuple with the IssuerId field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Idp) GetIssuerIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.IssuerId.Get(), o.IssuerId.IsSet()
}

// SetIssuerId sets field value
func (o *Idp) SetIssuerId(v string) {
	o.IssuerId.Set(&v)
}

// GetSfAccountId returns the SfAccountId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Idp) GetSfAccountId() string {
	if o == nil || o.SfAccountId.Get() == nil {
		var ret string
		return ret
	}
	return *o.SfAccountId.Get()
}

// GetSfAccountIdOk returns a tuple with the SfAccountId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Idp) GetSfAccountIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.SfAccountId.Get(), o.SfAccountId.IsSet()
}

// HasSfAccountId returns a boolean if a field has been set.
func (o *Idp) HasSfAccountId() bool {
	if o != nil && o.SfAccountId.IsSet() {
		return true
	}

	return false
}

// SetSfAccountId gets a reference to the given NullableString and assigns it to the SfAccountId field.
func (o *Idp) SetSfAccountId(v string) {
	o.SfAccountId.Set(&v)
}
// SetSfAccountIdNil sets the value for SfAccountId to be an explicit nil
func (o *Idp) SetSfAccountIdNil() {
	o.SfAccountId.Set(nil)
}

// UnsetSfAccountId ensures that no value is present for SfAccountId, not even an explicit nil
func (o *Idp) UnsetSfAccountId() {
	o.SfAccountId.Unset()
}

// GetCertificate returns the Certificate field value
// If the value is explicit nil, the zero value for string will be returned
func (o *Idp) GetCertificate() string {
	if o == nil || o.Certificate.Get() == nil {
		var ret string
		return ret
	}

	return *o.Certificate.Get()
}

// GetCertificateOk returns a tuple with the Certificate field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Idp) GetCertificateOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Certificate.Get(), o.Certificate.IsSet()
}

// SetCertificate sets field value
func (o *Idp) SetCertificate(v string) {
	o.Certificate.Set(&v)
}

// GetDefaultRoles returns the DefaultRoles field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Idp) GetDefaultRoles() []string {
	if o == nil  {
		var ret []string
		return ret
	}
	return o.DefaultRoles
}

// GetDefaultRolesOk returns a tuple with the DefaultRoles field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Idp) GetDefaultRolesOk() (*[]string, bool) {
	if o == nil || o.DefaultRoles == nil {
		return nil, false
	}
	return &o.DefaultRoles, true
}

// HasDefaultRoles returns a boolean if a field has been set.
func (o *Idp) HasDefaultRoles() bool {
	if o != nil && o.DefaultRoles != nil {
		return true
	}

	return false
}

// SetDefaultRoles gets a reference to the given []string and assigns it to the DefaultRoles field.
func (o *Idp) SetDefaultRoles(v []string) {
	o.DefaultRoles = v
}

// GetDefaultClusters returns the DefaultClusters field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Idp) GetDefaultClusters() []string {
	if o == nil  {
		var ret []string
		return ret
	}
	return o.DefaultClusters
}

// GetDefaultClustersOk returns a tuple with the DefaultClusters field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Idp) GetDefaultClustersOk() (*[]string, bool) {
	if o == nil || o.DefaultClusters == nil {
		return nil, false
	}
	return &o.DefaultClusters, true
}

// HasDefaultClusters returns a boolean if a field has been set.
func (o *Idp) HasDefaultClusters() bool {
	if o != nil && o.DefaultClusters != nil {
		return true
	}

	return false
}

// SetDefaultClusters gets a reference to the given []string and assigns it to the DefaultClusters field.
func (o *Idp) SetDefaultClusters(v []string) {
	o.DefaultClusters = v
}

// GetSignRequest returns the SignRequest field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Idp) GetSignRequest() bool {
	if o == nil || o.SignRequest.Get() == nil {
		var ret bool
		return ret
	}
	return *o.SignRequest.Get()
}

// GetSignRequestOk returns a tuple with the SignRequest field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Idp) GetSignRequestOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return o.SignRequest.Get(), o.SignRequest.IsSet()
}

// HasSignRequest returns a boolean if a field has been set.
func (o *Idp) HasSignRequest() bool {
	if o != nil && o.SignRequest.IsSet() {
		return true
	}

	return false
}

// SetSignRequest gets a reference to the given NullableBool and assigns it to the SignRequest field.
func (o *Idp) SetSignRequest(v bool) {
	o.SignRequest.Set(&v)
}
// SetSignRequestNil sets the value for SignRequest to be an explicit nil
func (o *Idp) SetSignRequestNil() {
	o.SignRequest.Set(nil)
}

// UnsetSignRequest ensures that no value is present for SignRequest, not even an explicit nil
func (o *Idp) UnsetSignRequest() {
	o.SignRequest.Unset()
}

// GetIsEnabled returns the IsEnabled field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Idp) GetIsEnabled() bool {
	if o == nil || o.IsEnabled.Get() == nil {
		var ret bool
		return ret
	}
	return *o.IsEnabled.Get()
}

// GetIsEnabledOk returns a tuple with the IsEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Idp) GetIsEnabledOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return o.IsEnabled.Get(), o.IsEnabled.IsSet()
}

// HasIsEnabled returns a boolean if a field has been set.
func (o *Idp) HasIsEnabled() bool {
	if o != nil && o.IsEnabled.IsSet() {
		return true
	}

	return false
}

// SetIsEnabled gets a reference to the given NullableBool and assigns it to the IsEnabled field.
func (o *Idp) SetIsEnabled(v bool) {
	o.IsEnabled.Set(&v)
}
// SetIsEnabledNil sets the value for IsEnabled to be an explicit nil
func (o *Idp) SetIsEnabledNil() {
	o.IsEnabled.Set(nil)
}

// UnsetIsEnabled ensures that no value is present for IsEnabled, not even an explicit nil
func (o *Idp) UnsetIsEnabled() {
	o.IsEnabled.Unset()
}

// GetTenantId returns the TenantId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Idp) GetTenantId() string {
	if o == nil || o.TenantId.Get() == nil {
		var ret string
		return ret
	}
	return *o.TenantId.Get()
}

// GetTenantIdOk returns a tuple with the TenantId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Idp) GetTenantIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.TenantId.Get(), o.TenantId.IsSet()
}

// HasTenantId returns a boolean if a field has been set.
func (o *Idp) HasTenantId() bool {
	if o != nil && o.TenantId.IsSet() {
		return true
	}

	return false
}

// SetTenantId gets a reference to the given NullableString and assigns it to the TenantId field.
func (o *Idp) SetTenantId(v string) {
	o.TenantId.Set(&v)
}
// SetTenantIdNil sets the value for TenantId to be an explicit nil
func (o *Idp) SetTenantIdNil() {
	o.TenantId.Set(nil)
}

// UnsetTenantId ensures that no value is present for TenantId, not even an explicit nil
func (o *Idp) UnsetTenantId() {
	o.TenantId.Unset()
}

// GetId returns the Id field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Idp) GetId() int64 {
	if o == nil || o.Id.Get() == nil {
		var ret int64
		return ret
	}
	return *o.Id.Get()
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Idp) GetIdOk() (*int64, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Id.Get(), o.Id.IsSet()
}

// HasId returns a boolean if a field has been set.
func (o *Idp) HasId() bool {
	if o != nil && o.Id.IsSet() {
		return true
	}

	return false
}

// SetId gets a reference to the given NullableInt64 and assigns it to the Id field.
func (o *Idp) SetId(v int64) {
	o.Id.Set(&v)
}
// SetIdNil sets the value for Id to be an explicit nil
func (o *Idp) SetIdNil() {
	o.Id.Set(nil)
}

// UnsetId ensures that no value is present for Id, not even an explicit nil
func (o *Idp) UnsetId() {
	o.Id.Unset()
}

func (o Idp) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["name"] = o.Name.Get()
	}
	if true {
		toSerialize["domain"] = o.Domain.Get()
	}
	if true {
		toSerialize["ssoUrl"] = o.SsoUrl.Get()
	}
	if true {
		toSerialize["issuerId"] = o.IssuerId.Get()
	}
	if o.SfAccountId.IsSet() {
		toSerialize["sfAccountId"] = o.SfAccountId.Get()
	}
	if true {
		toSerialize["certificate"] = o.Certificate.Get()
	}
	if o.DefaultRoles != nil {
		toSerialize["defaultRoles"] = o.DefaultRoles
	}
	if o.DefaultClusters != nil {
		toSerialize["defaultClusters"] = o.DefaultClusters
	}
	if o.SignRequest.IsSet() {
		toSerialize["signRequest"] = o.SignRequest.Get()
	}
	if o.IsEnabled.IsSet() {
		toSerialize["isEnabled"] = o.IsEnabled.Get()
	}
	if o.TenantId.IsSet() {
		toSerialize["tenantId"] = o.TenantId.Get()
	}
	if o.Id.IsSet() {
		toSerialize["id"] = o.Id.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableIdp struct {
	value *Idp
	isSet bool
}

func (v NullableIdp) Get() *Idp {
	return v.value
}

func (v *NullableIdp) Set(val *Idp) {
	v.value = val
	v.isSet = true
}

func (v NullableIdp) IsSet() bool {
	return v.isSet
}

func (v *NullableIdp) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIdp(val *Idp) *NullableIdp {
	return &NullableIdp{value: val, isSet: true}
}

func (v NullableIdp) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIdp) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


