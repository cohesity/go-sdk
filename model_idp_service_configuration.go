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

// IdpServiceConfiguration Specifies the configuration of an IdP service.
type IdpServiceConfiguration struct {
	// Specifies whether to allow local authentication. When IdP is configured, only IdP users are allowed to login to the Cluster. Local login is disabled except for users with admin role. If this flag is set to true, local (non-IdP) logins are allowed for all local and AD users. Local or AD users with admin role can login always independent of this flag's setting.
	AllowLocalAuthentication NullableBool `json:"allowLocalAuthentication,omitempty"`
	// Specifies the certificate generated for the app by the IdP service when the Cluster is registered as an app. This is required to verify the SAML response.
	Certificate NullableString `json:"certificate,omitempty"`
	// Specifies the filename used to upload the certificate.
	CertificateFilename NullableString `json:"certificateFilename,omitempty"`
	// Specifies a unique name for this IdP configuration.
	Domain NullableString `json:"domain,omitempty"`
	// Specifies a flag to enable or disable this IdP service. When it is set to true, IdP service is enabled. When it is set to false, IdP service is disabled. When an IdP service is created, it is set to true.
	Enable NullableBool `json:"enable,omitempty"`
	// Specifies the Id assigned by the Cluster for the IdP service.
	Id NullableInt64 `json:"id,omitempty"`
	// Specifies the IdP provided Issuer ID for the app. For example, exkh1aov1nhHrgFhN0h7.
	IssuerId NullableString `json:"issuerId,omitempty"`
	// Specifies the name of the vendor providing IdP service.
	Name NullableString `json:"name,omitempty"`
	// Specifies a list of roles assigned to an IdP user if samlAttributeName is not given.
	Roles []string `json:"roles,omitempty"`
	// Specifies the SAML attribute name that contains a comma separated list of Cluster roles. Either this field or roles must be set. This field takes higher precedence than the roles field.
	SamlAttributeName NullableString `json:"samlAttributeName,omitempty"`
	// Specifies whether to sign the SAML request or not. When it is set to true, SAML request will be signed. When it is set to false, SAML request is not signed. Default is false. Set this flag to true if the IdP site is configured to expect the SAML request from the Cluster signed. If this is set to true, users must get the Cluster's certificate and upload it on the IdP site.
	SignRequest NullableBool `json:"signRequest,omitempty"`
	// Specifies the SSO URL of the IdP service for the customer. This is the URL given by IdP when the customer created an account. Customers may use this for several clusters that are registered with on IdP site. For example, dev-332534.oktapreview.com
	SsoUrl NullableString `json:"ssoUrl,omitempty"`
	// Specifies the Tenant Id if the IdP is configured for a Tenant. If this is not set, this IdP configuration is used for the Cluster level users and for all users of Tenants not having an IdP configuration.
	TenantId NullableString `json:"tenantId,omitempty"`
}

// NewIdpServiceConfiguration instantiates a new IdpServiceConfiguration object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIdpServiceConfiguration() *IdpServiceConfiguration {
	this := IdpServiceConfiguration{}
	return &this
}

// NewIdpServiceConfigurationWithDefaults instantiates a new IdpServiceConfiguration object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIdpServiceConfigurationWithDefaults() *IdpServiceConfiguration {
	this := IdpServiceConfiguration{}
	return &this
}

// GetAllowLocalAuthentication returns the AllowLocalAuthentication field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IdpServiceConfiguration) GetAllowLocalAuthentication() bool {
	if o == nil || o.AllowLocalAuthentication.Get() == nil {
		var ret bool
		return ret
	}
	return *o.AllowLocalAuthentication.Get()
}

// GetAllowLocalAuthenticationOk returns a tuple with the AllowLocalAuthentication field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IdpServiceConfiguration) GetAllowLocalAuthenticationOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return o.AllowLocalAuthentication.Get(), o.AllowLocalAuthentication.IsSet()
}

// HasAllowLocalAuthentication returns a boolean if a field has been set.
func (o *IdpServiceConfiguration) HasAllowLocalAuthentication() bool {
	if o != nil && o.AllowLocalAuthentication.IsSet() {
		return true
	}

	return false
}

// SetAllowLocalAuthentication gets a reference to the given NullableBool and assigns it to the AllowLocalAuthentication field.
func (o *IdpServiceConfiguration) SetAllowLocalAuthentication(v bool) {
	o.AllowLocalAuthentication.Set(&v)
}
// SetAllowLocalAuthenticationNil sets the value for AllowLocalAuthentication to be an explicit nil
func (o *IdpServiceConfiguration) SetAllowLocalAuthenticationNil() {
	o.AllowLocalAuthentication.Set(nil)
}

// UnsetAllowLocalAuthentication ensures that no value is present for AllowLocalAuthentication, not even an explicit nil
func (o *IdpServiceConfiguration) UnsetAllowLocalAuthentication() {
	o.AllowLocalAuthentication.Unset()
}

// GetCertificate returns the Certificate field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IdpServiceConfiguration) GetCertificate() string {
	if o == nil || o.Certificate.Get() == nil {
		var ret string
		return ret
	}
	return *o.Certificate.Get()
}

// GetCertificateOk returns a tuple with the Certificate field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IdpServiceConfiguration) GetCertificateOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Certificate.Get(), o.Certificate.IsSet()
}

// HasCertificate returns a boolean if a field has been set.
func (o *IdpServiceConfiguration) HasCertificate() bool {
	if o != nil && o.Certificate.IsSet() {
		return true
	}

	return false
}

// SetCertificate gets a reference to the given NullableString and assigns it to the Certificate field.
func (o *IdpServiceConfiguration) SetCertificate(v string) {
	o.Certificate.Set(&v)
}
// SetCertificateNil sets the value for Certificate to be an explicit nil
func (o *IdpServiceConfiguration) SetCertificateNil() {
	o.Certificate.Set(nil)
}

// UnsetCertificate ensures that no value is present for Certificate, not even an explicit nil
func (o *IdpServiceConfiguration) UnsetCertificate() {
	o.Certificate.Unset()
}

// GetCertificateFilename returns the CertificateFilename field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IdpServiceConfiguration) GetCertificateFilename() string {
	if o == nil || o.CertificateFilename.Get() == nil {
		var ret string
		return ret
	}
	return *o.CertificateFilename.Get()
}

// GetCertificateFilenameOk returns a tuple with the CertificateFilename field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IdpServiceConfiguration) GetCertificateFilenameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.CertificateFilename.Get(), o.CertificateFilename.IsSet()
}

// HasCertificateFilename returns a boolean if a field has been set.
func (o *IdpServiceConfiguration) HasCertificateFilename() bool {
	if o != nil && o.CertificateFilename.IsSet() {
		return true
	}

	return false
}

// SetCertificateFilename gets a reference to the given NullableString and assigns it to the CertificateFilename field.
func (o *IdpServiceConfiguration) SetCertificateFilename(v string) {
	o.CertificateFilename.Set(&v)
}
// SetCertificateFilenameNil sets the value for CertificateFilename to be an explicit nil
func (o *IdpServiceConfiguration) SetCertificateFilenameNil() {
	o.CertificateFilename.Set(nil)
}

// UnsetCertificateFilename ensures that no value is present for CertificateFilename, not even an explicit nil
func (o *IdpServiceConfiguration) UnsetCertificateFilename() {
	o.CertificateFilename.Unset()
}

// GetDomain returns the Domain field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IdpServiceConfiguration) GetDomain() string {
	if o == nil || o.Domain.Get() == nil {
		var ret string
		return ret
	}
	return *o.Domain.Get()
}

// GetDomainOk returns a tuple with the Domain field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IdpServiceConfiguration) GetDomainOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Domain.Get(), o.Domain.IsSet()
}

// HasDomain returns a boolean if a field has been set.
func (o *IdpServiceConfiguration) HasDomain() bool {
	if o != nil && o.Domain.IsSet() {
		return true
	}

	return false
}

// SetDomain gets a reference to the given NullableString and assigns it to the Domain field.
func (o *IdpServiceConfiguration) SetDomain(v string) {
	o.Domain.Set(&v)
}
// SetDomainNil sets the value for Domain to be an explicit nil
func (o *IdpServiceConfiguration) SetDomainNil() {
	o.Domain.Set(nil)
}

// UnsetDomain ensures that no value is present for Domain, not even an explicit nil
func (o *IdpServiceConfiguration) UnsetDomain() {
	o.Domain.Unset()
}

// GetEnable returns the Enable field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IdpServiceConfiguration) GetEnable() bool {
	if o == nil || o.Enable.Get() == nil {
		var ret bool
		return ret
	}
	return *o.Enable.Get()
}

// GetEnableOk returns a tuple with the Enable field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IdpServiceConfiguration) GetEnableOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Enable.Get(), o.Enable.IsSet()
}

// HasEnable returns a boolean if a field has been set.
func (o *IdpServiceConfiguration) HasEnable() bool {
	if o != nil && o.Enable.IsSet() {
		return true
	}

	return false
}

// SetEnable gets a reference to the given NullableBool and assigns it to the Enable field.
func (o *IdpServiceConfiguration) SetEnable(v bool) {
	o.Enable.Set(&v)
}
// SetEnableNil sets the value for Enable to be an explicit nil
func (o *IdpServiceConfiguration) SetEnableNil() {
	o.Enable.Set(nil)
}

// UnsetEnable ensures that no value is present for Enable, not even an explicit nil
func (o *IdpServiceConfiguration) UnsetEnable() {
	o.Enable.Unset()
}

// GetId returns the Id field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IdpServiceConfiguration) GetId() int64 {
	if o == nil || o.Id.Get() == nil {
		var ret int64
		return ret
	}
	return *o.Id.Get()
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IdpServiceConfiguration) GetIdOk() (*int64, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Id.Get(), o.Id.IsSet()
}

// HasId returns a boolean if a field has been set.
func (o *IdpServiceConfiguration) HasId() bool {
	if o != nil && o.Id.IsSet() {
		return true
	}

	return false
}

// SetId gets a reference to the given NullableInt64 and assigns it to the Id field.
func (o *IdpServiceConfiguration) SetId(v int64) {
	o.Id.Set(&v)
}
// SetIdNil sets the value for Id to be an explicit nil
func (o *IdpServiceConfiguration) SetIdNil() {
	o.Id.Set(nil)
}

// UnsetId ensures that no value is present for Id, not even an explicit nil
func (o *IdpServiceConfiguration) UnsetId() {
	o.Id.Unset()
}

// GetIssuerId returns the IssuerId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IdpServiceConfiguration) GetIssuerId() string {
	if o == nil || o.IssuerId.Get() == nil {
		var ret string
		return ret
	}
	return *o.IssuerId.Get()
}

// GetIssuerIdOk returns a tuple with the IssuerId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IdpServiceConfiguration) GetIssuerIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.IssuerId.Get(), o.IssuerId.IsSet()
}

// HasIssuerId returns a boolean if a field has been set.
func (o *IdpServiceConfiguration) HasIssuerId() bool {
	if o != nil && o.IssuerId.IsSet() {
		return true
	}

	return false
}

// SetIssuerId gets a reference to the given NullableString and assigns it to the IssuerId field.
func (o *IdpServiceConfiguration) SetIssuerId(v string) {
	o.IssuerId.Set(&v)
}
// SetIssuerIdNil sets the value for IssuerId to be an explicit nil
func (o *IdpServiceConfiguration) SetIssuerIdNil() {
	o.IssuerId.Set(nil)
}

// UnsetIssuerId ensures that no value is present for IssuerId, not even an explicit nil
func (o *IdpServiceConfiguration) UnsetIssuerId() {
	o.IssuerId.Unset()
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IdpServiceConfiguration) GetName() string {
	if o == nil || o.Name.Get() == nil {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IdpServiceConfiguration) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *IdpServiceConfiguration) HasName() bool {
	if o != nil && o.Name.IsSet() {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableString and assigns it to the Name field.
func (o *IdpServiceConfiguration) SetName(v string) {
	o.Name.Set(&v)
}
// SetNameNil sets the value for Name to be an explicit nil
func (o *IdpServiceConfiguration) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *IdpServiceConfiguration) UnsetName() {
	o.Name.Unset()
}

// GetRoles returns the Roles field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IdpServiceConfiguration) GetRoles() []string {
	if o == nil  {
		var ret []string
		return ret
	}
	return o.Roles
}

// GetRolesOk returns a tuple with the Roles field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IdpServiceConfiguration) GetRolesOk() (*[]string, bool) {
	if o == nil || o.Roles == nil {
		return nil, false
	}
	return &o.Roles, true
}

// HasRoles returns a boolean if a field has been set.
func (o *IdpServiceConfiguration) HasRoles() bool {
	if o != nil && o.Roles != nil {
		return true
	}

	return false
}

// SetRoles gets a reference to the given []string and assigns it to the Roles field.
func (o *IdpServiceConfiguration) SetRoles(v []string) {
	o.Roles = v
}

// GetSamlAttributeName returns the SamlAttributeName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IdpServiceConfiguration) GetSamlAttributeName() string {
	if o == nil || o.SamlAttributeName.Get() == nil {
		var ret string
		return ret
	}
	return *o.SamlAttributeName.Get()
}

// GetSamlAttributeNameOk returns a tuple with the SamlAttributeName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IdpServiceConfiguration) GetSamlAttributeNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.SamlAttributeName.Get(), o.SamlAttributeName.IsSet()
}

// HasSamlAttributeName returns a boolean if a field has been set.
func (o *IdpServiceConfiguration) HasSamlAttributeName() bool {
	if o != nil && o.SamlAttributeName.IsSet() {
		return true
	}

	return false
}

// SetSamlAttributeName gets a reference to the given NullableString and assigns it to the SamlAttributeName field.
func (o *IdpServiceConfiguration) SetSamlAttributeName(v string) {
	o.SamlAttributeName.Set(&v)
}
// SetSamlAttributeNameNil sets the value for SamlAttributeName to be an explicit nil
func (o *IdpServiceConfiguration) SetSamlAttributeNameNil() {
	o.SamlAttributeName.Set(nil)
}

// UnsetSamlAttributeName ensures that no value is present for SamlAttributeName, not even an explicit nil
func (o *IdpServiceConfiguration) UnsetSamlAttributeName() {
	o.SamlAttributeName.Unset()
}

// GetSignRequest returns the SignRequest field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IdpServiceConfiguration) GetSignRequest() bool {
	if o == nil || o.SignRequest.Get() == nil {
		var ret bool
		return ret
	}
	return *o.SignRequest.Get()
}

// GetSignRequestOk returns a tuple with the SignRequest field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IdpServiceConfiguration) GetSignRequestOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return o.SignRequest.Get(), o.SignRequest.IsSet()
}

// HasSignRequest returns a boolean if a field has been set.
func (o *IdpServiceConfiguration) HasSignRequest() bool {
	if o != nil && o.SignRequest.IsSet() {
		return true
	}

	return false
}

// SetSignRequest gets a reference to the given NullableBool and assigns it to the SignRequest field.
func (o *IdpServiceConfiguration) SetSignRequest(v bool) {
	o.SignRequest.Set(&v)
}
// SetSignRequestNil sets the value for SignRequest to be an explicit nil
func (o *IdpServiceConfiguration) SetSignRequestNil() {
	o.SignRequest.Set(nil)
}

// UnsetSignRequest ensures that no value is present for SignRequest, not even an explicit nil
func (o *IdpServiceConfiguration) UnsetSignRequest() {
	o.SignRequest.Unset()
}

// GetSsoUrl returns the SsoUrl field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IdpServiceConfiguration) GetSsoUrl() string {
	if o == nil || o.SsoUrl.Get() == nil {
		var ret string
		return ret
	}
	return *o.SsoUrl.Get()
}

// GetSsoUrlOk returns a tuple with the SsoUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IdpServiceConfiguration) GetSsoUrlOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.SsoUrl.Get(), o.SsoUrl.IsSet()
}

// HasSsoUrl returns a boolean if a field has been set.
func (o *IdpServiceConfiguration) HasSsoUrl() bool {
	if o != nil && o.SsoUrl.IsSet() {
		return true
	}

	return false
}

// SetSsoUrl gets a reference to the given NullableString and assigns it to the SsoUrl field.
func (o *IdpServiceConfiguration) SetSsoUrl(v string) {
	o.SsoUrl.Set(&v)
}
// SetSsoUrlNil sets the value for SsoUrl to be an explicit nil
func (o *IdpServiceConfiguration) SetSsoUrlNil() {
	o.SsoUrl.Set(nil)
}

// UnsetSsoUrl ensures that no value is present for SsoUrl, not even an explicit nil
func (o *IdpServiceConfiguration) UnsetSsoUrl() {
	o.SsoUrl.Unset()
}

// GetTenantId returns the TenantId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IdpServiceConfiguration) GetTenantId() string {
	if o == nil || o.TenantId.Get() == nil {
		var ret string
		return ret
	}
	return *o.TenantId.Get()
}

// GetTenantIdOk returns a tuple with the TenantId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IdpServiceConfiguration) GetTenantIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.TenantId.Get(), o.TenantId.IsSet()
}

// HasTenantId returns a boolean if a field has been set.
func (o *IdpServiceConfiguration) HasTenantId() bool {
	if o != nil && o.TenantId.IsSet() {
		return true
	}

	return false
}

// SetTenantId gets a reference to the given NullableString and assigns it to the TenantId field.
func (o *IdpServiceConfiguration) SetTenantId(v string) {
	o.TenantId.Set(&v)
}
// SetTenantIdNil sets the value for TenantId to be an explicit nil
func (o *IdpServiceConfiguration) SetTenantIdNil() {
	o.TenantId.Set(nil)
}

// UnsetTenantId ensures that no value is present for TenantId, not even an explicit nil
func (o *IdpServiceConfiguration) UnsetTenantId() {
	o.TenantId.Unset()
}

func (o IdpServiceConfiguration) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AllowLocalAuthentication.IsSet() {
		toSerialize["allowLocalAuthentication"] = o.AllowLocalAuthentication.Get()
	}
	if o.Certificate.IsSet() {
		toSerialize["certificate"] = o.Certificate.Get()
	}
	if o.CertificateFilename.IsSet() {
		toSerialize["certificateFilename"] = o.CertificateFilename.Get()
	}
	if o.Domain.IsSet() {
		toSerialize["domain"] = o.Domain.Get()
	}
	if o.Enable.IsSet() {
		toSerialize["enable"] = o.Enable.Get()
	}
	if o.Id.IsSet() {
		toSerialize["id"] = o.Id.Get()
	}
	if o.IssuerId.IsSet() {
		toSerialize["issuerId"] = o.IssuerId.Get()
	}
	if o.Name.IsSet() {
		toSerialize["name"] = o.Name.Get()
	}
	if o.Roles != nil {
		toSerialize["roles"] = o.Roles
	}
	if o.SamlAttributeName.IsSet() {
		toSerialize["samlAttributeName"] = o.SamlAttributeName.Get()
	}
	if o.SignRequest.IsSet() {
		toSerialize["signRequest"] = o.SignRequest.Get()
	}
	if o.SsoUrl.IsSet() {
		toSerialize["ssoUrl"] = o.SsoUrl.Get()
	}
	if o.TenantId.IsSet() {
		toSerialize["tenantId"] = o.TenantId.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableIdpServiceConfiguration struct {
	value *IdpServiceConfiguration
	isSet bool
}

func (v NullableIdpServiceConfiguration) Get() *IdpServiceConfiguration {
	return v.value
}

func (v *NullableIdpServiceConfiguration) Set(val *IdpServiceConfiguration) {
	v.value = val
	v.isSet = true
}

func (v NullableIdpServiceConfiguration) IsSet() bool {
	return v.isSet
}

func (v *NullableIdpServiceConfiguration) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIdpServiceConfiguration(val *IdpServiceConfiguration) *NullableIdpServiceConfiguration {
	return &NullableIdpServiceConfiguration{value: val, isSet: true}
}

func (v NullableIdpServiceConfiguration) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIdpServiceConfiguration) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


