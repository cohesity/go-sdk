# UpdateIdpConfigurationRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AllowLocalAuthentication** | Pointer to **NullableBool** | Specifies whether to allow local authentication. When IdP is configured, only IdP users are allowed to login to the Cluster. Local login is disabled except for users with admin role. If this flag is set to true, local (non-IdP) logins are allowed for all local and AD users. Local or AD users with admin role can login always independent of this flag&#39;s setting. | [optional] 
**Certificate** | Pointer to **NullableString** | Specifies the certificate generated for the app by the IdP service when the Cluster is registered as an app. This is required to verify the SAML response. | [optional] 
**CertificateFilename** | Pointer to **NullableString** | Specifies the filename used to upload the certificate. | [optional] 
**Enable** | Pointer to **NullableBool** | Specifies a flag to enable or disable this IdP service. When it is set to true, IdP service is enabled. When it is set to false, IdP service is disabled. When an IdP service is created, it is set to true. | [optional] 
**IssuerId** | Pointer to **NullableString** | Specifies the IdP provided Issuer ID for the app. For example, exkh1aov1nhHrgFhN0h7. | [optional] 
**Roles** | Pointer to **[]string** | Specifies a list of roles assigned to an IdP user if samlAttributeName is not given. | [optional] 
**SamlAttributeName** | Pointer to **NullableString** | Specifies the SAML attribute name that contains a comma separated list of Cluster roles. Either this field or roles must be set. This field takes higher precedence than the roles field. | [optional] 
**SignRequest** | Pointer to **NullableBool** | Specifies whether to sign the SAML request or not. When it is set to true, SAML request will be signed. When it is set to false, SAML request is not signed. Default is false. Set this flag to true if the IdP site is configured to expect the SAML request from the Cluster signed. If this is set to true, users must get the Cluster&#39;s certificate and upload it on the IdP site. | [optional] 
**SsoUrl** | Pointer to **NullableString** | Specifies the SSO URL of the IdP service for the customer. This is the URL given by IdP when the customer created an account. Customers may use this for several clusters that are registered with on IdP site. For example, dev-332534.oktapreview.com | [optional] 

## Methods

### NewUpdateIdpConfigurationRequest

`func NewUpdateIdpConfigurationRequest() *UpdateIdpConfigurationRequest`

NewUpdateIdpConfigurationRequest instantiates a new UpdateIdpConfigurationRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateIdpConfigurationRequestWithDefaults

`func NewUpdateIdpConfigurationRequestWithDefaults() *UpdateIdpConfigurationRequest`

NewUpdateIdpConfigurationRequestWithDefaults instantiates a new UpdateIdpConfigurationRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAllowLocalAuthentication

`func (o *UpdateIdpConfigurationRequest) GetAllowLocalAuthentication() bool`

GetAllowLocalAuthentication returns the AllowLocalAuthentication field if non-nil, zero value otherwise.

### GetAllowLocalAuthenticationOk

`func (o *UpdateIdpConfigurationRequest) GetAllowLocalAuthenticationOk() (*bool, bool)`

GetAllowLocalAuthenticationOk returns a tuple with the AllowLocalAuthentication field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowLocalAuthentication

`func (o *UpdateIdpConfigurationRequest) SetAllowLocalAuthentication(v bool)`

SetAllowLocalAuthentication sets AllowLocalAuthentication field to given value.

### HasAllowLocalAuthentication

`func (o *UpdateIdpConfigurationRequest) HasAllowLocalAuthentication() bool`

HasAllowLocalAuthentication returns a boolean if a field has been set.

### SetAllowLocalAuthenticationNil

`func (o *UpdateIdpConfigurationRequest) SetAllowLocalAuthenticationNil(b bool)`

 SetAllowLocalAuthenticationNil sets the value for AllowLocalAuthentication to be an explicit nil

### UnsetAllowLocalAuthentication
`func (o *UpdateIdpConfigurationRequest) UnsetAllowLocalAuthentication()`

UnsetAllowLocalAuthentication ensures that no value is present for AllowLocalAuthentication, not even an explicit nil
### GetCertificate

`func (o *UpdateIdpConfigurationRequest) GetCertificate() string`

GetCertificate returns the Certificate field if non-nil, zero value otherwise.

### GetCertificateOk

`func (o *UpdateIdpConfigurationRequest) GetCertificateOk() (*string, bool)`

GetCertificateOk returns a tuple with the Certificate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificate

`func (o *UpdateIdpConfigurationRequest) SetCertificate(v string)`

SetCertificate sets Certificate field to given value.

### HasCertificate

`func (o *UpdateIdpConfigurationRequest) HasCertificate() bool`

HasCertificate returns a boolean if a field has been set.

### SetCertificateNil

`func (o *UpdateIdpConfigurationRequest) SetCertificateNil(b bool)`

 SetCertificateNil sets the value for Certificate to be an explicit nil

### UnsetCertificate
`func (o *UpdateIdpConfigurationRequest) UnsetCertificate()`

UnsetCertificate ensures that no value is present for Certificate, not even an explicit nil
### GetCertificateFilename

`func (o *UpdateIdpConfigurationRequest) GetCertificateFilename() string`

GetCertificateFilename returns the CertificateFilename field if non-nil, zero value otherwise.

### GetCertificateFilenameOk

`func (o *UpdateIdpConfigurationRequest) GetCertificateFilenameOk() (*string, bool)`

GetCertificateFilenameOk returns a tuple with the CertificateFilename field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificateFilename

`func (o *UpdateIdpConfigurationRequest) SetCertificateFilename(v string)`

SetCertificateFilename sets CertificateFilename field to given value.

### HasCertificateFilename

`func (o *UpdateIdpConfigurationRequest) HasCertificateFilename() bool`

HasCertificateFilename returns a boolean if a field has been set.

### SetCertificateFilenameNil

`func (o *UpdateIdpConfigurationRequest) SetCertificateFilenameNil(b bool)`

 SetCertificateFilenameNil sets the value for CertificateFilename to be an explicit nil

### UnsetCertificateFilename
`func (o *UpdateIdpConfigurationRequest) UnsetCertificateFilename()`

UnsetCertificateFilename ensures that no value is present for CertificateFilename, not even an explicit nil
### GetEnable

`func (o *UpdateIdpConfigurationRequest) GetEnable() bool`

GetEnable returns the Enable field if non-nil, zero value otherwise.

### GetEnableOk

`func (o *UpdateIdpConfigurationRequest) GetEnableOk() (*bool, bool)`

GetEnableOk returns a tuple with the Enable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnable

`func (o *UpdateIdpConfigurationRequest) SetEnable(v bool)`

SetEnable sets Enable field to given value.

### HasEnable

`func (o *UpdateIdpConfigurationRequest) HasEnable() bool`

HasEnable returns a boolean if a field has been set.

### SetEnableNil

`func (o *UpdateIdpConfigurationRequest) SetEnableNil(b bool)`

 SetEnableNil sets the value for Enable to be an explicit nil

### UnsetEnable
`func (o *UpdateIdpConfigurationRequest) UnsetEnable()`

UnsetEnable ensures that no value is present for Enable, not even an explicit nil
### GetIssuerId

`func (o *UpdateIdpConfigurationRequest) GetIssuerId() string`

GetIssuerId returns the IssuerId field if non-nil, zero value otherwise.

### GetIssuerIdOk

`func (o *UpdateIdpConfigurationRequest) GetIssuerIdOk() (*string, bool)`

GetIssuerIdOk returns a tuple with the IssuerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuerId

`func (o *UpdateIdpConfigurationRequest) SetIssuerId(v string)`

SetIssuerId sets IssuerId field to given value.

### HasIssuerId

`func (o *UpdateIdpConfigurationRequest) HasIssuerId() bool`

HasIssuerId returns a boolean if a field has been set.

### SetIssuerIdNil

`func (o *UpdateIdpConfigurationRequest) SetIssuerIdNil(b bool)`

 SetIssuerIdNil sets the value for IssuerId to be an explicit nil

### UnsetIssuerId
`func (o *UpdateIdpConfigurationRequest) UnsetIssuerId()`

UnsetIssuerId ensures that no value is present for IssuerId, not even an explicit nil
### GetRoles

`func (o *UpdateIdpConfigurationRequest) GetRoles() []string`

GetRoles returns the Roles field if non-nil, zero value otherwise.

### GetRolesOk

`func (o *UpdateIdpConfigurationRequest) GetRolesOk() (*[]string, bool)`

GetRolesOk returns a tuple with the Roles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoles

`func (o *UpdateIdpConfigurationRequest) SetRoles(v []string)`

SetRoles sets Roles field to given value.

### HasRoles

`func (o *UpdateIdpConfigurationRequest) HasRoles() bool`

HasRoles returns a boolean if a field has been set.

### SetRolesNil

`func (o *UpdateIdpConfigurationRequest) SetRolesNil(b bool)`

 SetRolesNil sets the value for Roles to be an explicit nil

### UnsetRoles
`func (o *UpdateIdpConfigurationRequest) UnsetRoles()`

UnsetRoles ensures that no value is present for Roles, not even an explicit nil
### GetSamlAttributeName

`func (o *UpdateIdpConfigurationRequest) GetSamlAttributeName() string`

GetSamlAttributeName returns the SamlAttributeName field if non-nil, zero value otherwise.

### GetSamlAttributeNameOk

`func (o *UpdateIdpConfigurationRequest) GetSamlAttributeNameOk() (*string, bool)`

GetSamlAttributeNameOk returns a tuple with the SamlAttributeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSamlAttributeName

`func (o *UpdateIdpConfigurationRequest) SetSamlAttributeName(v string)`

SetSamlAttributeName sets SamlAttributeName field to given value.

### HasSamlAttributeName

`func (o *UpdateIdpConfigurationRequest) HasSamlAttributeName() bool`

HasSamlAttributeName returns a boolean if a field has been set.

### SetSamlAttributeNameNil

`func (o *UpdateIdpConfigurationRequest) SetSamlAttributeNameNil(b bool)`

 SetSamlAttributeNameNil sets the value for SamlAttributeName to be an explicit nil

### UnsetSamlAttributeName
`func (o *UpdateIdpConfigurationRequest) UnsetSamlAttributeName()`

UnsetSamlAttributeName ensures that no value is present for SamlAttributeName, not even an explicit nil
### GetSignRequest

`func (o *UpdateIdpConfigurationRequest) GetSignRequest() bool`

GetSignRequest returns the SignRequest field if non-nil, zero value otherwise.

### GetSignRequestOk

`func (o *UpdateIdpConfigurationRequest) GetSignRequestOk() (*bool, bool)`

GetSignRequestOk returns a tuple with the SignRequest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignRequest

`func (o *UpdateIdpConfigurationRequest) SetSignRequest(v bool)`

SetSignRequest sets SignRequest field to given value.

### HasSignRequest

`func (o *UpdateIdpConfigurationRequest) HasSignRequest() bool`

HasSignRequest returns a boolean if a field has been set.

### SetSignRequestNil

`func (o *UpdateIdpConfigurationRequest) SetSignRequestNil(b bool)`

 SetSignRequestNil sets the value for SignRequest to be an explicit nil

### UnsetSignRequest
`func (o *UpdateIdpConfigurationRequest) UnsetSignRequest()`

UnsetSignRequest ensures that no value is present for SignRequest, not even an explicit nil
### GetSsoUrl

`func (o *UpdateIdpConfigurationRequest) GetSsoUrl() string`

GetSsoUrl returns the SsoUrl field if non-nil, zero value otherwise.

### GetSsoUrlOk

`func (o *UpdateIdpConfigurationRequest) GetSsoUrlOk() (*string, bool)`

GetSsoUrlOk returns a tuple with the SsoUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSsoUrl

`func (o *UpdateIdpConfigurationRequest) SetSsoUrl(v string)`

SetSsoUrl sets SsoUrl field to given value.

### HasSsoUrl

`func (o *UpdateIdpConfigurationRequest) HasSsoUrl() bool`

HasSsoUrl returns a boolean if a field has been set.

### SetSsoUrlNil

`func (o *UpdateIdpConfigurationRequest) SetSsoUrlNil(b bool)`

 SetSsoUrlNil sets the value for SsoUrl to be an explicit nil

### UnsetSsoUrl
`func (o *UpdateIdpConfigurationRequest) UnsetSsoUrl()`

UnsetSsoUrl ensures that no value is present for SsoUrl, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


