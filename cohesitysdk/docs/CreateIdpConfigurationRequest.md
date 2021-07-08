# CreateIdpConfigurationRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AllowLocalAuthentication** | Pointer to **NullableBool** | Specifies whether to allow local authentication. When IdP is configured, only IdP users are allowed to login to the Cluster. Local login is disabled except for users with admin role. If this flag is set to true, local (non-IdP) logins are allowed for all local and AD users. Local or AD users with admin role can login always independent of this flag&#39;s setting. | [optional] 
**Certificate** | Pointer to **NullableString** | Specifies the certificate generated for the app by the IdP service when the Cluster is registered as an app. This is required to verify the SAML response. | [optional] 
**CertificateFilename** | Pointer to **NullableString** | Specifies the filename used to upload the certificate. | [optional] 
**Domain** | Pointer to **NullableString** | Specifies a unique name for this IdP configuration. | [optional] 
**Enable** | Pointer to **NullableBool** | Specifies a flag to enable or disable this IdP service. When it is set to true, IdP service is enabled. When it is set to false, IdP service is disabled. When an IdP service is created, it is set to true. | [optional] 
**IssuerId** | Pointer to **NullableString** | Specifies the IdP provided Issuer ID for the app. For example, exkh1aov1nhHrgFhN0h7. | [optional] 
**Name** | Pointer to **NullableString** | Specifies the name of the vendor providing IdP service. | [optional] 
**Roles** | Pointer to **[]string** | Specifies a list of roles assigned to an IdP user if samlAttributeName is not given. | [optional] 
**SamlAttributeName** | Pointer to **NullableString** | Specifies the SAML attribute name that contains a comma separated list of Cluster roles. Either this field or roles must be set. This field takes higher precedence than the roles field. | [optional] 
**SignRequest** | Pointer to **NullableBool** | Specifies whether to sign the SAML request or not. When it is set to true, SAML request will be signed. When it is set to false, SAML request is not signed. Default is false. Set this flag to true if the IdP site is configured to expect the SAML request from the Cluster signed. If this is set to true, users must get the Cluster&#39;s certificate and upload it on the IdP site. | [optional] 
**SsoUrl** | Pointer to **NullableString** | Specifies the SSO URL of the IdP service for the customer. This is the URL given by IdP when the customer created an account. Customers may use this for several clusters that are registered with on IdP site. For example, dev-332534.oktapreview.com | [optional] 
**TenantId** | Pointer to **NullableString** | Specifies the Tenant Id if the IdP is configured for a Tenant. If this is not set, this IdP configuration is used for the Cluster level users and for all users of Tenants not having an IdP configuration. | [optional] 

## Methods

### NewCreateIdpConfigurationRequest

`func NewCreateIdpConfigurationRequest() *CreateIdpConfigurationRequest`

NewCreateIdpConfigurationRequest instantiates a new CreateIdpConfigurationRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateIdpConfigurationRequestWithDefaults

`func NewCreateIdpConfigurationRequestWithDefaults() *CreateIdpConfigurationRequest`

NewCreateIdpConfigurationRequestWithDefaults instantiates a new CreateIdpConfigurationRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAllowLocalAuthentication

`func (o *CreateIdpConfigurationRequest) GetAllowLocalAuthentication() bool`

GetAllowLocalAuthentication returns the AllowLocalAuthentication field if non-nil, zero value otherwise.

### GetAllowLocalAuthenticationOk

`func (o *CreateIdpConfigurationRequest) GetAllowLocalAuthenticationOk() (*bool, bool)`

GetAllowLocalAuthenticationOk returns a tuple with the AllowLocalAuthentication field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowLocalAuthentication

`func (o *CreateIdpConfigurationRequest) SetAllowLocalAuthentication(v bool)`

SetAllowLocalAuthentication sets AllowLocalAuthentication field to given value.

### HasAllowLocalAuthentication

`func (o *CreateIdpConfigurationRequest) HasAllowLocalAuthentication() bool`

HasAllowLocalAuthentication returns a boolean if a field has been set.

### SetAllowLocalAuthenticationNil

`func (o *CreateIdpConfigurationRequest) SetAllowLocalAuthenticationNil(b bool)`

 SetAllowLocalAuthenticationNil sets the value for AllowLocalAuthentication to be an explicit nil

### UnsetAllowLocalAuthentication
`func (o *CreateIdpConfigurationRequest) UnsetAllowLocalAuthentication()`

UnsetAllowLocalAuthentication ensures that no value is present for AllowLocalAuthentication, not even an explicit nil
### GetCertificate

`func (o *CreateIdpConfigurationRequest) GetCertificate() string`

GetCertificate returns the Certificate field if non-nil, zero value otherwise.

### GetCertificateOk

`func (o *CreateIdpConfigurationRequest) GetCertificateOk() (*string, bool)`

GetCertificateOk returns a tuple with the Certificate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificate

`func (o *CreateIdpConfigurationRequest) SetCertificate(v string)`

SetCertificate sets Certificate field to given value.

### HasCertificate

`func (o *CreateIdpConfigurationRequest) HasCertificate() bool`

HasCertificate returns a boolean if a field has been set.

### SetCertificateNil

`func (o *CreateIdpConfigurationRequest) SetCertificateNil(b bool)`

 SetCertificateNil sets the value for Certificate to be an explicit nil

### UnsetCertificate
`func (o *CreateIdpConfigurationRequest) UnsetCertificate()`

UnsetCertificate ensures that no value is present for Certificate, not even an explicit nil
### GetCertificateFilename

`func (o *CreateIdpConfigurationRequest) GetCertificateFilename() string`

GetCertificateFilename returns the CertificateFilename field if non-nil, zero value otherwise.

### GetCertificateFilenameOk

`func (o *CreateIdpConfigurationRequest) GetCertificateFilenameOk() (*string, bool)`

GetCertificateFilenameOk returns a tuple with the CertificateFilename field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificateFilename

`func (o *CreateIdpConfigurationRequest) SetCertificateFilename(v string)`

SetCertificateFilename sets CertificateFilename field to given value.

### HasCertificateFilename

`func (o *CreateIdpConfigurationRequest) HasCertificateFilename() bool`

HasCertificateFilename returns a boolean if a field has been set.

### SetCertificateFilenameNil

`func (o *CreateIdpConfigurationRequest) SetCertificateFilenameNil(b bool)`

 SetCertificateFilenameNil sets the value for CertificateFilename to be an explicit nil

### UnsetCertificateFilename
`func (o *CreateIdpConfigurationRequest) UnsetCertificateFilename()`

UnsetCertificateFilename ensures that no value is present for CertificateFilename, not even an explicit nil
### GetDomain

`func (o *CreateIdpConfigurationRequest) GetDomain() string`

GetDomain returns the Domain field if non-nil, zero value otherwise.

### GetDomainOk

`func (o *CreateIdpConfigurationRequest) GetDomainOk() (*string, bool)`

GetDomainOk returns a tuple with the Domain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomain

`func (o *CreateIdpConfigurationRequest) SetDomain(v string)`

SetDomain sets Domain field to given value.

### HasDomain

`func (o *CreateIdpConfigurationRequest) HasDomain() bool`

HasDomain returns a boolean if a field has been set.

### SetDomainNil

`func (o *CreateIdpConfigurationRequest) SetDomainNil(b bool)`

 SetDomainNil sets the value for Domain to be an explicit nil

### UnsetDomain
`func (o *CreateIdpConfigurationRequest) UnsetDomain()`

UnsetDomain ensures that no value is present for Domain, not even an explicit nil
### GetEnable

`func (o *CreateIdpConfigurationRequest) GetEnable() bool`

GetEnable returns the Enable field if non-nil, zero value otherwise.

### GetEnableOk

`func (o *CreateIdpConfigurationRequest) GetEnableOk() (*bool, bool)`

GetEnableOk returns a tuple with the Enable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnable

`func (o *CreateIdpConfigurationRequest) SetEnable(v bool)`

SetEnable sets Enable field to given value.

### HasEnable

`func (o *CreateIdpConfigurationRequest) HasEnable() bool`

HasEnable returns a boolean if a field has been set.

### SetEnableNil

`func (o *CreateIdpConfigurationRequest) SetEnableNil(b bool)`

 SetEnableNil sets the value for Enable to be an explicit nil

### UnsetEnable
`func (o *CreateIdpConfigurationRequest) UnsetEnable()`

UnsetEnable ensures that no value is present for Enable, not even an explicit nil
### GetIssuerId

`func (o *CreateIdpConfigurationRequest) GetIssuerId() string`

GetIssuerId returns the IssuerId field if non-nil, zero value otherwise.

### GetIssuerIdOk

`func (o *CreateIdpConfigurationRequest) GetIssuerIdOk() (*string, bool)`

GetIssuerIdOk returns a tuple with the IssuerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuerId

`func (o *CreateIdpConfigurationRequest) SetIssuerId(v string)`

SetIssuerId sets IssuerId field to given value.

### HasIssuerId

`func (o *CreateIdpConfigurationRequest) HasIssuerId() bool`

HasIssuerId returns a boolean if a field has been set.

### SetIssuerIdNil

`func (o *CreateIdpConfigurationRequest) SetIssuerIdNil(b bool)`

 SetIssuerIdNil sets the value for IssuerId to be an explicit nil

### UnsetIssuerId
`func (o *CreateIdpConfigurationRequest) UnsetIssuerId()`

UnsetIssuerId ensures that no value is present for IssuerId, not even an explicit nil
### GetName

`func (o *CreateIdpConfigurationRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateIdpConfigurationRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateIdpConfigurationRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CreateIdpConfigurationRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *CreateIdpConfigurationRequest) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *CreateIdpConfigurationRequest) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetRoles

`func (o *CreateIdpConfigurationRequest) GetRoles() []string`

GetRoles returns the Roles field if non-nil, zero value otherwise.

### GetRolesOk

`func (o *CreateIdpConfigurationRequest) GetRolesOk() (*[]string, bool)`

GetRolesOk returns a tuple with the Roles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoles

`func (o *CreateIdpConfigurationRequest) SetRoles(v []string)`

SetRoles sets Roles field to given value.

### HasRoles

`func (o *CreateIdpConfigurationRequest) HasRoles() bool`

HasRoles returns a boolean if a field has been set.

### SetRolesNil

`func (o *CreateIdpConfigurationRequest) SetRolesNil(b bool)`

 SetRolesNil sets the value for Roles to be an explicit nil

### UnsetRoles
`func (o *CreateIdpConfigurationRequest) UnsetRoles()`

UnsetRoles ensures that no value is present for Roles, not even an explicit nil
### GetSamlAttributeName

`func (o *CreateIdpConfigurationRequest) GetSamlAttributeName() string`

GetSamlAttributeName returns the SamlAttributeName field if non-nil, zero value otherwise.

### GetSamlAttributeNameOk

`func (o *CreateIdpConfigurationRequest) GetSamlAttributeNameOk() (*string, bool)`

GetSamlAttributeNameOk returns a tuple with the SamlAttributeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSamlAttributeName

`func (o *CreateIdpConfigurationRequest) SetSamlAttributeName(v string)`

SetSamlAttributeName sets SamlAttributeName field to given value.

### HasSamlAttributeName

`func (o *CreateIdpConfigurationRequest) HasSamlAttributeName() bool`

HasSamlAttributeName returns a boolean if a field has been set.

### SetSamlAttributeNameNil

`func (o *CreateIdpConfigurationRequest) SetSamlAttributeNameNil(b bool)`

 SetSamlAttributeNameNil sets the value for SamlAttributeName to be an explicit nil

### UnsetSamlAttributeName
`func (o *CreateIdpConfigurationRequest) UnsetSamlAttributeName()`

UnsetSamlAttributeName ensures that no value is present for SamlAttributeName, not even an explicit nil
### GetSignRequest

`func (o *CreateIdpConfigurationRequest) GetSignRequest() bool`

GetSignRequest returns the SignRequest field if non-nil, zero value otherwise.

### GetSignRequestOk

`func (o *CreateIdpConfigurationRequest) GetSignRequestOk() (*bool, bool)`

GetSignRequestOk returns a tuple with the SignRequest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignRequest

`func (o *CreateIdpConfigurationRequest) SetSignRequest(v bool)`

SetSignRequest sets SignRequest field to given value.

### HasSignRequest

`func (o *CreateIdpConfigurationRequest) HasSignRequest() bool`

HasSignRequest returns a boolean if a field has been set.

### SetSignRequestNil

`func (o *CreateIdpConfigurationRequest) SetSignRequestNil(b bool)`

 SetSignRequestNil sets the value for SignRequest to be an explicit nil

### UnsetSignRequest
`func (o *CreateIdpConfigurationRequest) UnsetSignRequest()`

UnsetSignRequest ensures that no value is present for SignRequest, not even an explicit nil
### GetSsoUrl

`func (o *CreateIdpConfigurationRequest) GetSsoUrl() string`

GetSsoUrl returns the SsoUrl field if non-nil, zero value otherwise.

### GetSsoUrlOk

`func (o *CreateIdpConfigurationRequest) GetSsoUrlOk() (*string, bool)`

GetSsoUrlOk returns a tuple with the SsoUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSsoUrl

`func (o *CreateIdpConfigurationRequest) SetSsoUrl(v string)`

SetSsoUrl sets SsoUrl field to given value.

### HasSsoUrl

`func (o *CreateIdpConfigurationRequest) HasSsoUrl() bool`

HasSsoUrl returns a boolean if a field has been set.

### SetSsoUrlNil

`func (o *CreateIdpConfigurationRequest) SetSsoUrlNil(b bool)`

 SetSsoUrlNil sets the value for SsoUrl to be an explicit nil

### UnsetSsoUrl
`func (o *CreateIdpConfigurationRequest) UnsetSsoUrl()`

UnsetSsoUrl ensures that no value is present for SsoUrl, not even an explicit nil
### GetTenantId

`func (o *CreateIdpConfigurationRequest) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *CreateIdpConfigurationRequest) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *CreateIdpConfigurationRequest) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.

### HasTenantId

`func (o *CreateIdpConfigurationRequest) HasTenantId() bool`

HasTenantId returns a boolean if a field has been set.

### SetTenantIdNil

`func (o *CreateIdpConfigurationRequest) SetTenantIdNil(b bool)`

 SetTenantIdNil sets the value for TenantId to be an explicit nil

### UnsetTenantId
`func (o *CreateIdpConfigurationRequest) UnsetTenantId()`

UnsetTenantId ensures that no value is present for TenantId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


