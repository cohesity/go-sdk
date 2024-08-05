# UpdateIdpRequestParams

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

### NewUpdateIdpRequestParams

`func NewUpdateIdpRequestParams(certificate NullableString, issuerId NullableString, ssoUrl NullableString, ) *UpdateIdpRequestParams`

NewUpdateIdpRequestParams instantiates a new UpdateIdpRequestParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateIdpRequestParamsWithDefaults

`func NewUpdateIdpRequestParamsWithDefaults() *UpdateIdpRequestParams`

NewUpdateIdpRequestParamsWithDefaults instantiates a new UpdateIdpRequestParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAllowLocalUserLogin

`func (o *UpdateIdpRequestParams) GetAllowLocalUserLogin() bool`

GetAllowLocalUserLogin returns the AllowLocalUserLogin field if non-nil, zero value otherwise.

### GetAllowLocalUserLoginOk

`func (o *UpdateIdpRequestParams) GetAllowLocalUserLoginOk() (*bool, bool)`

GetAllowLocalUserLoginOk returns a tuple with the AllowLocalUserLogin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowLocalUserLogin

`func (o *UpdateIdpRequestParams) SetAllowLocalUserLogin(v bool)`

SetAllowLocalUserLogin sets AllowLocalUserLogin field to given value.

### HasAllowLocalUserLogin

`func (o *UpdateIdpRequestParams) HasAllowLocalUserLogin() bool`

HasAllowLocalUserLogin returns a boolean if a field has been set.

### SetAllowLocalUserLoginNil

`func (o *UpdateIdpRequestParams) SetAllowLocalUserLoginNil(b bool)`

 SetAllowLocalUserLoginNil sets the value for AllowLocalUserLogin to be an explicit nil

### UnsetAllowLocalUserLogin
`func (o *UpdateIdpRequestParams) UnsetAllowLocalUserLogin()`

UnsetAllowLocalUserLogin ensures that no value is present for AllowLocalUserLogin, not even an explicit nil
### GetCertificate

`func (o *UpdateIdpRequestParams) GetCertificate() string`

GetCertificate returns the Certificate field if non-nil, zero value otherwise.

### GetCertificateOk

`func (o *UpdateIdpRequestParams) GetCertificateOk() (*string, bool)`

GetCertificateOk returns a tuple with the Certificate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificate

`func (o *UpdateIdpRequestParams) SetCertificate(v string)`

SetCertificate sets Certificate field to given value.


### SetCertificateNil

`func (o *UpdateIdpRequestParams) SetCertificateNil(b bool)`

 SetCertificateNil sets the value for Certificate to be an explicit nil

### UnsetCertificate
`func (o *UpdateIdpRequestParams) UnsetCertificate()`

UnsetCertificate ensures that no value is present for Certificate, not even an explicit nil
### GetCertificateFilename

`func (o *UpdateIdpRequestParams) GetCertificateFilename() string`

GetCertificateFilename returns the CertificateFilename field if non-nil, zero value otherwise.

### GetCertificateFilenameOk

`func (o *UpdateIdpRequestParams) GetCertificateFilenameOk() (*string, bool)`

GetCertificateFilenameOk returns a tuple with the CertificateFilename field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificateFilename

`func (o *UpdateIdpRequestParams) SetCertificateFilename(v string)`

SetCertificateFilename sets CertificateFilename field to given value.

### HasCertificateFilename

`func (o *UpdateIdpRequestParams) HasCertificateFilename() bool`

HasCertificateFilename returns a boolean if a field has been set.

### SetCertificateFilenameNil

`func (o *UpdateIdpRequestParams) SetCertificateFilenameNil(b bool)`

 SetCertificateFilenameNil sets the value for CertificateFilename to be an explicit nil

### UnsetCertificateFilename
`func (o *UpdateIdpRequestParams) UnsetCertificateFilename()`

UnsetCertificateFilename ensures that no value is present for CertificateFilename, not even an explicit nil
### GetIsEnabled

`func (o *UpdateIdpRequestParams) GetIsEnabled() bool`

GetIsEnabled returns the IsEnabled field if non-nil, zero value otherwise.

### GetIsEnabledOk

`func (o *UpdateIdpRequestParams) GetIsEnabledOk() (*bool, bool)`

GetIsEnabledOk returns a tuple with the IsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsEnabled

`func (o *UpdateIdpRequestParams) SetIsEnabled(v bool)`

SetIsEnabled sets IsEnabled field to given value.

### HasIsEnabled

`func (o *UpdateIdpRequestParams) HasIsEnabled() bool`

HasIsEnabled returns a boolean if a field has been set.

### SetIsEnabledNil

`func (o *UpdateIdpRequestParams) SetIsEnabledNil(b bool)`

 SetIsEnabledNil sets the value for IsEnabled to be an explicit nil

### UnsetIsEnabled
`func (o *UpdateIdpRequestParams) UnsetIsEnabled()`

UnsetIsEnabled ensures that no value is present for IsEnabled, not even an explicit nil
### GetIssuerId

`func (o *UpdateIdpRequestParams) GetIssuerId() string`

GetIssuerId returns the IssuerId field if non-nil, zero value otherwise.

### GetIssuerIdOk

`func (o *UpdateIdpRequestParams) GetIssuerIdOk() (*string, bool)`

GetIssuerIdOk returns a tuple with the IssuerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuerId

`func (o *UpdateIdpRequestParams) SetIssuerId(v string)`

SetIssuerId sets IssuerId field to given value.


### SetIssuerIdNil

`func (o *UpdateIdpRequestParams) SetIssuerIdNil(b bool)`

 SetIssuerIdNil sets the value for IssuerId to be an explicit nil

### UnsetIssuerId
`func (o *UpdateIdpRequestParams) UnsetIssuerId()`

UnsetIssuerId ensures that no value is present for IssuerId, not even an explicit nil
### GetRoles

`func (o *UpdateIdpRequestParams) GetRoles() []string`

GetRoles returns the Roles field if non-nil, zero value otherwise.

### GetRolesOk

`func (o *UpdateIdpRequestParams) GetRolesOk() (*[]string, bool)`

GetRolesOk returns a tuple with the Roles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoles

`func (o *UpdateIdpRequestParams) SetRoles(v []string)`

SetRoles sets Roles field to given value.

### HasRoles

`func (o *UpdateIdpRequestParams) HasRoles() bool`

HasRoles returns a boolean if a field has been set.

### SetRolesNil

`func (o *UpdateIdpRequestParams) SetRolesNil(b bool)`

 SetRolesNil sets the value for Roles to be an explicit nil

### UnsetRoles
`func (o *UpdateIdpRequestParams) UnsetRoles()`

UnsetRoles ensures that no value is present for Roles, not even an explicit nil
### GetSamlAttributeName

`func (o *UpdateIdpRequestParams) GetSamlAttributeName() string`

GetSamlAttributeName returns the SamlAttributeName field if non-nil, zero value otherwise.

### GetSamlAttributeNameOk

`func (o *UpdateIdpRequestParams) GetSamlAttributeNameOk() (*string, bool)`

GetSamlAttributeNameOk returns a tuple with the SamlAttributeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSamlAttributeName

`func (o *UpdateIdpRequestParams) SetSamlAttributeName(v string)`

SetSamlAttributeName sets SamlAttributeName field to given value.

### HasSamlAttributeName

`func (o *UpdateIdpRequestParams) HasSamlAttributeName() bool`

HasSamlAttributeName returns a boolean if a field has been set.

### SetSamlAttributeNameNil

`func (o *UpdateIdpRequestParams) SetSamlAttributeNameNil(b bool)`

 SetSamlAttributeNameNil sets the value for SamlAttributeName to be an explicit nil

### UnsetSamlAttributeName
`func (o *UpdateIdpRequestParams) UnsetSamlAttributeName()`

UnsetSamlAttributeName ensures that no value is present for SamlAttributeName, not even an explicit nil
### GetSignRequest

`func (o *UpdateIdpRequestParams) GetSignRequest() bool`

GetSignRequest returns the SignRequest field if non-nil, zero value otherwise.

### GetSignRequestOk

`func (o *UpdateIdpRequestParams) GetSignRequestOk() (*bool, bool)`

GetSignRequestOk returns a tuple with the SignRequest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignRequest

`func (o *UpdateIdpRequestParams) SetSignRequest(v bool)`

SetSignRequest sets SignRequest field to given value.

### HasSignRequest

`func (o *UpdateIdpRequestParams) HasSignRequest() bool`

HasSignRequest returns a boolean if a field has been set.

### SetSignRequestNil

`func (o *UpdateIdpRequestParams) SetSignRequestNil(b bool)`

 SetSignRequestNil sets the value for SignRequest to be an explicit nil

### UnsetSignRequest
`func (o *UpdateIdpRequestParams) UnsetSignRequest()`

UnsetSignRequest ensures that no value is present for SignRequest, not even an explicit nil
### GetSsoUrl

`func (o *UpdateIdpRequestParams) GetSsoUrl() string`

GetSsoUrl returns the SsoUrl field if non-nil, zero value otherwise.

### GetSsoUrlOk

`func (o *UpdateIdpRequestParams) GetSsoUrlOk() (*string, bool)`

GetSsoUrlOk returns a tuple with the SsoUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSsoUrl

`func (o *UpdateIdpRequestParams) SetSsoUrl(v string)`

SetSsoUrl sets SsoUrl field to given value.


### SetSsoUrlNil

`func (o *UpdateIdpRequestParams) SetSsoUrlNil(b bool)`

 SetSsoUrlNil sets the value for SsoUrl to be an explicit nil

### UnsetSsoUrl
`func (o *UpdateIdpRequestParams) UnsetSsoUrl()`

UnsetSsoUrl ensures that no value is present for SsoUrl, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


