# CommonIdentityProviderConfiguration

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

## Methods

### NewCommonIdentityProviderConfiguration

`func NewCommonIdentityProviderConfiguration(certificate NullableString, issuerId NullableString, ssoUrl NullableString, ) *CommonIdentityProviderConfiguration`

NewCommonIdentityProviderConfiguration instantiates a new CommonIdentityProviderConfiguration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCommonIdentityProviderConfigurationWithDefaults

`func NewCommonIdentityProviderConfigurationWithDefaults() *CommonIdentityProviderConfiguration`

NewCommonIdentityProviderConfigurationWithDefaults instantiates a new CommonIdentityProviderConfiguration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAllowLocalUserLogin

`func (o *CommonIdentityProviderConfiguration) GetAllowLocalUserLogin() bool`

GetAllowLocalUserLogin returns the AllowLocalUserLogin field if non-nil, zero value otherwise.

### GetAllowLocalUserLoginOk

`func (o *CommonIdentityProviderConfiguration) GetAllowLocalUserLoginOk() (*bool, bool)`

GetAllowLocalUserLoginOk returns a tuple with the AllowLocalUserLogin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowLocalUserLogin

`func (o *CommonIdentityProviderConfiguration) SetAllowLocalUserLogin(v bool)`

SetAllowLocalUserLogin sets AllowLocalUserLogin field to given value.

### HasAllowLocalUserLogin

`func (o *CommonIdentityProviderConfiguration) HasAllowLocalUserLogin() bool`

HasAllowLocalUserLogin returns a boolean if a field has been set.

### SetAllowLocalUserLoginNil

`func (o *CommonIdentityProviderConfiguration) SetAllowLocalUserLoginNil(b bool)`

 SetAllowLocalUserLoginNil sets the value for AllowLocalUserLogin to be an explicit nil

### UnsetAllowLocalUserLogin
`func (o *CommonIdentityProviderConfiguration) UnsetAllowLocalUserLogin()`

UnsetAllowLocalUserLogin ensures that no value is present for AllowLocalUserLogin, not even an explicit nil
### GetCertificate

`func (o *CommonIdentityProviderConfiguration) GetCertificate() string`

GetCertificate returns the Certificate field if non-nil, zero value otherwise.

### GetCertificateOk

`func (o *CommonIdentityProviderConfiguration) GetCertificateOk() (*string, bool)`

GetCertificateOk returns a tuple with the Certificate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificate

`func (o *CommonIdentityProviderConfiguration) SetCertificate(v string)`

SetCertificate sets Certificate field to given value.


### SetCertificateNil

`func (o *CommonIdentityProviderConfiguration) SetCertificateNil(b bool)`

 SetCertificateNil sets the value for Certificate to be an explicit nil

### UnsetCertificate
`func (o *CommonIdentityProviderConfiguration) UnsetCertificate()`

UnsetCertificate ensures that no value is present for Certificate, not even an explicit nil
### GetCertificateFilename

`func (o *CommonIdentityProviderConfiguration) GetCertificateFilename() string`

GetCertificateFilename returns the CertificateFilename field if non-nil, zero value otherwise.

### GetCertificateFilenameOk

`func (o *CommonIdentityProviderConfiguration) GetCertificateFilenameOk() (*string, bool)`

GetCertificateFilenameOk returns a tuple with the CertificateFilename field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificateFilename

`func (o *CommonIdentityProviderConfiguration) SetCertificateFilename(v string)`

SetCertificateFilename sets CertificateFilename field to given value.

### HasCertificateFilename

`func (o *CommonIdentityProviderConfiguration) HasCertificateFilename() bool`

HasCertificateFilename returns a boolean if a field has been set.

### SetCertificateFilenameNil

`func (o *CommonIdentityProviderConfiguration) SetCertificateFilenameNil(b bool)`

 SetCertificateFilenameNil sets the value for CertificateFilename to be an explicit nil

### UnsetCertificateFilename
`func (o *CommonIdentityProviderConfiguration) UnsetCertificateFilename()`

UnsetCertificateFilename ensures that no value is present for CertificateFilename, not even an explicit nil
### GetIsEnabled

`func (o *CommonIdentityProviderConfiguration) GetIsEnabled() bool`

GetIsEnabled returns the IsEnabled field if non-nil, zero value otherwise.

### GetIsEnabledOk

`func (o *CommonIdentityProviderConfiguration) GetIsEnabledOk() (*bool, bool)`

GetIsEnabledOk returns a tuple with the IsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsEnabled

`func (o *CommonIdentityProviderConfiguration) SetIsEnabled(v bool)`

SetIsEnabled sets IsEnabled field to given value.

### HasIsEnabled

`func (o *CommonIdentityProviderConfiguration) HasIsEnabled() bool`

HasIsEnabled returns a boolean if a field has been set.

### SetIsEnabledNil

`func (o *CommonIdentityProviderConfiguration) SetIsEnabledNil(b bool)`

 SetIsEnabledNil sets the value for IsEnabled to be an explicit nil

### UnsetIsEnabled
`func (o *CommonIdentityProviderConfiguration) UnsetIsEnabled()`

UnsetIsEnabled ensures that no value is present for IsEnabled, not even an explicit nil
### GetIssuerId

`func (o *CommonIdentityProviderConfiguration) GetIssuerId() string`

GetIssuerId returns the IssuerId field if non-nil, zero value otherwise.

### GetIssuerIdOk

`func (o *CommonIdentityProviderConfiguration) GetIssuerIdOk() (*string, bool)`

GetIssuerIdOk returns a tuple with the IssuerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuerId

`func (o *CommonIdentityProviderConfiguration) SetIssuerId(v string)`

SetIssuerId sets IssuerId field to given value.


### SetIssuerIdNil

`func (o *CommonIdentityProviderConfiguration) SetIssuerIdNil(b bool)`

 SetIssuerIdNil sets the value for IssuerId to be an explicit nil

### UnsetIssuerId
`func (o *CommonIdentityProviderConfiguration) UnsetIssuerId()`

UnsetIssuerId ensures that no value is present for IssuerId, not even an explicit nil
### GetRoles

`func (o *CommonIdentityProviderConfiguration) GetRoles() []string`

GetRoles returns the Roles field if non-nil, zero value otherwise.

### GetRolesOk

`func (o *CommonIdentityProviderConfiguration) GetRolesOk() (*[]string, bool)`

GetRolesOk returns a tuple with the Roles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoles

`func (o *CommonIdentityProviderConfiguration) SetRoles(v []string)`

SetRoles sets Roles field to given value.

### HasRoles

`func (o *CommonIdentityProviderConfiguration) HasRoles() bool`

HasRoles returns a boolean if a field has been set.

### SetRolesNil

`func (o *CommonIdentityProviderConfiguration) SetRolesNil(b bool)`

 SetRolesNil sets the value for Roles to be an explicit nil

### UnsetRoles
`func (o *CommonIdentityProviderConfiguration) UnsetRoles()`

UnsetRoles ensures that no value is present for Roles, not even an explicit nil
### GetSamlAttributeName

`func (o *CommonIdentityProviderConfiguration) GetSamlAttributeName() string`

GetSamlAttributeName returns the SamlAttributeName field if non-nil, zero value otherwise.

### GetSamlAttributeNameOk

`func (o *CommonIdentityProviderConfiguration) GetSamlAttributeNameOk() (*string, bool)`

GetSamlAttributeNameOk returns a tuple with the SamlAttributeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSamlAttributeName

`func (o *CommonIdentityProviderConfiguration) SetSamlAttributeName(v string)`

SetSamlAttributeName sets SamlAttributeName field to given value.

### HasSamlAttributeName

`func (o *CommonIdentityProviderConfiguration) HasSamlAttributeName() bool`

HasSamlAttributeName returns a boolean if a field has been set.

### SetSamlAttributeNameNil

`func (o *CommonIdentityProviderConfiguration) SetSamlAttributeNameNil(b bool)`

 SetSamlAttributeNameNil sets the value for SamlAttributeName to be an explicit nil

### UnsetSamlAttributeName
`func (o *CommonIdentityProviderConfiguration) UnsetSamlAttributeName()`

UnsetSamlAttributeName ensures that no value is present for SamlAttributeName, not even an explicit nil
### GetSignRequest

`func (o *CommonIdentityProviderConfiguration) GetSignRequest() bool`

GetSignRequest returns the SignRequest field if non-nil, zero value otherwise.

### GetSignRequestOk

`func (o *CommonIdentityProviderConfiguration) GetSignRequestOk() (*bool, bool)`

GetSignRequestOk returns a tuple with the SignRequest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignRequest

`func (o *CommonIdentityProviderConfiguration) SetSignRequest(v bool)`

SetSignRequest sets SignRequest field to given value.

### HasSignRequest

`func (o *CommonIdentityProviderConfiguration) HasSignRequest() bool`

HasSignRequest returns a boolean if a field has been set.

### SetSignRequestNil

`func (o *CommonIdentityProviderConfiguration) SetSignRequestNil(b bool)`

 SetSignRequestNil sets the value for SignRequest to be an explicit nil

### UnsetSignRequest
`func (o *CommonIdentityProviderConfiguration) UnsetSignRequest()`

UnsetSignRequest ensures that no value is present for SignRequest, not even an explicit nil
### GetSsoUrl

`func (o *CommonIdentityProviderConfiguration) GetSsoUrl() string`

GetSsoUrl returns the SsoUrl field if non-nil, zero value otherwise.

### GetSsoUrlOk

`func (o *CommonIdentityProviderConfiguration) GetSsoUrlOk() (*string, bool)`

GetSsoUrlOk returns a tuple with the SsoUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSsoUrl

`func (o *CommonIdentityProviderConfiguration) SetSsoUrl(v string)`

SetSsoUrl sets SsoUrl field to given value.


### SetSsoUrlNil

`func (o *CommonIdentityProviderConfiguration) SetSsoUrlNil(b bool)`

 SetSsoUrlNil sets the value for SsoUrl to be an explicit nil

### UnsetSsoUrl
`func (o *CommonIdentityProviderConfiguration) UnsetSsoUrl()`

UnsetSsoUrl ensures that no value is present for SsoUrl, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


