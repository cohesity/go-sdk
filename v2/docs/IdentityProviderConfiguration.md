# IdentityProviderConfiguration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AllowLocalUserLogin** | Pointer to **NullableBool** | Specifies if local user login is allowed. When idp is configured, only idp users are allowed to login to the cluster, local login is disabled except for users with admin role. If this flag is set to true, local (non-idp) logins are allowed for all local and AD users. Local or AD users with admin role can login always independent of this flag&#39;s setting. By default there is no local authentication i.e the value is false. | [optional] 
**Certificate** | **NullableString** | Specifies the certificate generated for the app by the idp service when the cluster is registered as an app. This is required to verify the SAML response. | 
**CertificateFilename** | Pointer to **NullableString** | Specifies the filename used for the certificate. The default value is idp_certificate.pem | [optional] 
**IsEnabled** | Pointer to **NullableBool** | Specifies a flag to enable or disable this idp service. When it is set to true, idp service is enabled. When it is set to false, idp service is disabled. By defaut idp is enabled i.e the value is true. | [optional] 
**IssuerId** | **NullableString** | Specifies identity provider issuer id | 
**Roles** | Pointer to **[]string** | Specifies the default roles assined for all SSO users | [optional] 
**SamlAttributeName** | Pointer to **NullableString** | Specifies the SAML attribute name that contains a comma separated list of cluster roles. This sets the default roles for all SSO users. Either this field or roles must be set, this field takes higher precedence than the roles field. | [optional] 
**SignRequest** | Pointer to **NullableBool** | Specifies whether to sign the SAML request or not. When it is set to true, SAML request will be signed. When it is set to false, SAML request is not signed. Default is false, set this flag to true if the idp site is configured to expect the SAML request from the Cluster signed. If this is set to true, users must get the cluster&#39;s certificate and upload it on the idp site. | [optional] 
**SsoUrl** | **NullableString** | Specifies the identity provider SSO url | 
**Domain** | Pointer to **NullableString** | Specifies domain of idp configuration | [optional] 
**Id** | Pointer to **NullableInt64** | Specifies id of idp configuration | [optional] 
**Name** | Pointer to **NullableString** | Specifies name of the vendor providing idp service | [optional] 
**TenantId** | Pointer to **NullableString** | Specifies the tenant id if the idp is configured for a tenant. If this is not set, this idp configuration is used for the cluster level users and for all users of tenants not having an idp configuration. | [optional] 

## Methods

### NewIdentityProviderConfiguration

`func NewIdentityProviderConfiguration(certificate NullableString, issuerId NullableString, ssoUrl NullableString, ) *IdentityProviderConfiguration`

NewIdentityProviderConfiguration instantiates a new IdentityProviderConfiguration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIdentityProviderConfigurationWithDefaults

`func NewIdentityProviderConfigurationWithDefaults() *IdentityProviderConfiguration`

NewIdentityProviderConfigurationWithDefaults instantiates a new IdentityProviderConfiguration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAllowLocalUserLogin

`func (o *IdentityProviderConfiguration) GetAllowLocalUserLogin() bool`

GetAllowLocalUserLogin returns the AllowLocalUserLogin field if non-nil, zero value otherwise.

### GetAllowLocalUserLoginOk

`func (o *IdentityProviderConfiguration) GetAllowLocalUserLoginOk() (*bool, bool)`

GetAllowLocalUserLoginOk returns a tuple with the AllowLocalUserLogin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowLocalUserLogin

`func (o *IdentityProviderConfiguration) SetAllowLocalUserLogin(v bool)`

SetAllowLocalUserLogin sets AllowLocalUserLogin field to given value.

### HasAllowLocalUserLogin

`func (o *IdentityProviderConfiguration) HasAllowLocalUserLogin() bool`

HasAllowLocalUserLogin returns a boolean if a field has been set.

### SetAllowLocalUserLoginNil

`func (o *IdentityProviderConfiguration) SetAllowLocalUserLoginNil(b bool)`

 SetAllowLocalUserLoginNil sets the value for AllowLocalUserLogin to be an explicit nil

### UnsetAllowLocalUserLogin
`func (o *IdentityProviderConfiguration) UnsetAllowLocalUserLogin()`

UnsetAllowLocalUserLogin ensures that no value is present for AllowLocalUserLogin, not even an explicit nil
### GetCertificate

`func (o *IdentityProviderConfiguration) GetCertificate() string`

GetCertificate returns the Certificate field if non-nil, zero value otherwise.

### GetCertificateOk

`func (o *IdentityProviderConfiguration) GetCertificateOk() (*string, bool)`

GetCertificateOk returns a tuple with the Certificate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificate

`func (o *IdentityProviderConfiguration) SetCertificate(v string)`

SetCertificate sets Certificate field to given value.


### SetCertificateNil

`func (o *IdentityProviderConfiguration) SetCertificateNil(b bool)`

 SetCertificateNil sets the value for Certificate to be an explicit nil

### UnsetCertificate
`func (o *IdentityProviderConfiguration) UnsetCertificate()`

UnsetCertificate ensures that no value is present for Certificate, not even an explicit nil
### GetCertificateFilename

`func (o *IdentityProviderConfiguration) GetCertificateFilename() string`

GetCertificateFilename returns the CertificateFilename field if non-nil, zero value otherwise.

### GetCertificateFilenameOk

`func (o *IdentityProviderConfiguration) GetCertificateFilenameOk() (*string, bool)`

GetCertificateFilenameOk returns a tuple with the CertificateFilename field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificateFilename

`func (o *IdentityProviderConfiguration) SetCertificateFilename(v string)`

SetCertificateFilename sets CertificateFilename field to given value.

### HasCertificateFilename

`func (o *IdentityProviderConfiguration) HasCertificateFilename() bool`

HasCertificateFilename returns a boolean if a field has been set.

### SetCertificateFilenameNil

`func (o *IdentityProviderConfiguration) SetCertificateFilenameNil(b bool)`

 SetCertificateFilenameNil sets the value for CertificateFilename to be an explicit nil

### UnsetCertificateFilename
`func (o *IdentityProviderConfiguration) UnsetCertificateFilename()`

UnsetCertificateFilename ensures that no value is present for CertificateFilename, not even an explicit nil
### GetIsEnabled

`func (o *IdentityProviderConfiguration) GetIsEnabled() bool`

GetIsEnabled returns the IsEnabled field if non-nil, zero value otherwise.

### GetIsEnabledOk

`func (o *IdentityProviderConfiguration) GetIsEnabledOk() (*bool, bool)`

GetIsEnabledOk returns a tuple with the IsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsEnabled

`func (o *IdentityProviderConfiguration) SetIsEnabled(v bool)`

SetIsEnabled sets IsEnabled field to given value.

### HasIsEnabled

`func (o *IdentityProviderConfiguration) HasIsEnabled() bool`

HasIsEnabled returns a boolean if a field has been set.

### SetIsEnabledNil

`func (o *IdentityProviderConfiguration) SetIsEnabledNil(b bool)`

 SetIsEnabledNil sets the value for IsEnabled to be an explicit nil

### UnsetIsEnabled
`func (o *IdentityProviderConfiguration) UnsetIsEnabled()`

UnsetIsEnabled ensures that no value is present for IsEnabled, not even an explicit nil
### GetIssuerId

`func (o *IdentityProviderConfiguration) GetIssuerId() string`

GetIssuerId returns the IssuerId field if non-nil, zero value otherwise.

### GetIssuerIdOk

`func (o *IdentityProviderConfiguration) GetIssuerIdOk() (*string, bool)`

GetIssuerIdOk returns a tuple with the IssuerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuerId

`func (o *IdentityProviderConfiguration) SetIssuerId(v string)`

SetIssuerId sets IssuerId field to given value.


### SetIssuerIdNil

`func (o *IdentityProviderConfiguration) SetIssuerIdNil(b bool)`

 SetIssuerIdNil sets the value for IssuerId to be an explicit nil

### UnsetIssuerId
`func (o *IdentityProviderConfiguration) UnsetIssuerId()`

UnsetIssuerId ensures that no value is present for IssuerId, not even an explicit nil
### GetRoles

`func (o *IdentityProviderConfiguration) GetRoles() []string`

GetRoles returns the Roles field if non-nil, zero value otherwise.

### GetRolesOk

`func (o *IdentityProviderConfiguration) GetRolesOk() (*[]string, bool)`

GetRolesOk returns a tuple with the Roles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoles

`func (o *IdentityProviderConfiguration) SetRoles(v []string)`

SetRoles sets Roles field to given value.

### HasRoles

`func (o *IdentityProviderConfiguration) HasRoles() bool`

HasRoles returns a boolean if a field has been set.

### SetRolesNil

`func (o *IdentityProviderConfiguration) SetRolesNil(b bool)`

 SetRolesNil sets the value for Roles to be an explicit nil

### UnsetRoles
`func (o *IdentityProviderConfiguration) UnsetRoles()`

UnsetRoles ensures that no value is present for Roles, not even an explicit nil
### GetSamlAttributeName

`func (o *IdentityProviderConfiguration) GetSamlAttributeName() string`

GetSamlAttributeName returns the SamlAttributeName field if non-nil, zero value otherwise.

### GetSamlAttributeNameOk

`func (o *IdentityProviderConfiguration) GetSamlAttributeNameOk() (*string, bool)`

GetSamlAttributeNameOk returns a tuple with the SamlAttributeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSamlAttributeName

`func (o *IdentityProviderConfiguration) SetSamlAttributeName(v string)`

SetSamlAttributeName sets SamlAttributeName field to given value.

### HasSamlAttributeName

`func (o *IdentityProviderConfiguration) HasSamlAttributeName() bool`

HasSamlAttributeName returns a boolean if a field has been set.

### SetSamlAttributeNameNil

`func (o *IdentityProviderConfiguration) SetSamlAttributeNameNil(b bool)`

 SetSamlAttributeNameNil sets the value for SamlAttributeName to be an explicit nil

### UnsetSamlAttributeName
`func (o *IdentityProviderConfiguration) UnsetSamlAttributeName()`

UnsetSamlAttributeName ensures that no value is present for SamlAttributeName, not even an explicit nil
### GetSignRequest

`func (o *IdentityProviderConfiguration) GetSignRequest() bool`

GetSignRequest returns the SignRequest field if non-nil, zero value otherwise.

### GetSignRequestOk

`func (o *IdentityProviderConfiguration) GetSignRequestOk() (*bool, bool)`

GetSignRequestOk returns a tuple with the SignRequest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignRequest

`func (o *IdentityProviderConfiguration) SetSignRequest(v bool)`

SetSignRequest sets SignRequest field to given value.

### HasSignRequest

`func (o *IdentityProviderConfiguration) HasSignRequest() bool`

HasSignRequest returns a boolean if a field has been set.

### SetSignRequestNil

`func (o *IdentityProviderConfiguration) SetSignRequestNil(b bool)`

 SetSignRequestNil sets the value for SignRequest to be an explicit nil

### UnsetSignRequest
`func (o *IdentityProviderConfiguration) UnsetSignRequest()`

UnsetSignRequest ensures that no value is present for SignRequest, not even an explicit nil
### GetSsoUrl

`func (o *IdentityProviderConfiguration) GetSsoUrl() string`

GetSsoUrl returns the SsoUrl field if non-nil, zero value otherwise.

### GetSsoUrlOk

`func (o *IdentityProviderConfiguration) GetSsoUrlOk() (*string, bool)`

GetSsoUrlOk returns a tuple with the SsoUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSsoUrl

`func (o *IdentityProviderConfiguration) SetSsoUrl(v string)`

SetSsoUrl sets SsoUrl field to given value.


### SetSsoUrlNil

`func (o *IdentityProviderConfiguration) SetSsoUrlNil(b bool)`

 SetSsoUrlNil sets the value for SsoUrl to be an explicit nil

### UnsetSsoUrl
`func (o *IdentityProviderConfiguration) UnsetSsoUrl()`

UnsetSsoUrl ensures that no value is present for SsoUrl, not even an explicit nil
### GetDomain

`func (o *IdentityProviderConfiguration) GetDomain() string`

GetDomain returns the Domain field if non-nil, zero value otherwise.

### GetDomainOk

`func (o *IdentityProviderConfiguration) GetDomainOk() (*string, bool)`

GetDomainOk returns a tuple with the Domain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomain

`func (o *IdentityProviderConfiguration) SetDomain(v string)`

SetDomain sets Domain field to given value.

### HasDomain

`func (o *IdentityProviderConfiguration) HasDomain() bool`

HasDomain returns a boolean if a field has been set.

### SetDomainNil

`func (o *IdentityProviderConfiguration) SetDomainNil(b bool)`

 SetDomainNil sets the value for Domain to be an explicit nil

### UnsetDomain
`func (o *IdentityProviderConfiguration) UnsetDomain()`

UnsetDomain ensures that no value is present for Domain, not even an explicit nil
### GetId

`func (o *IdentityProviderConfiguration) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *IdentityProviderConfiguration) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *IdentityProviderConfiguration) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *IdentityProviderConfiguration) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *IdentityProviderConfiguration) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *IdentityProviderConfiguration) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetName

`func (o *IdentityProviderConfiguration) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *IdentityProviderConfiguration) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *IdentityProviderConfiguration) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *IdentityProviderConfiguration) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *IdentityProviderConfiguration) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *IdentityProviderConfiguration) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetTenantId

`func (o *IdentityProviderConfiguration) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *IdentityProviderConfiguration) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *IdentityProviderConfiguration) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.

### HasTenantId

`func (o *IdentityProviderConfiguration) HasTenantId() bool`

HasTenantId returns a boolean if a field has been set.

### SetTenantIdNil

`func (o *IdentityProviderConfiguration) SetTenantIdNil(b bool)`

 SetTenantIdNil sets the value for TenantId to be an explicit nil

### UnsetTenantId
`func (o *IdentityProviderConfiguration) UnsetTenantId()`

UnsetTenantId ensures that no value is present for TenantId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


