# Idp

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **NullableString** | Specifies the name of the vendor providing IDP service. | 
**Domain** | **NullableString** | Specifies a unique name for this IdP configuration. | 
**SsoUrl** | **NullableString** | Specifies the SSO URL of the IdP service for the customer. This is the URL given by IdP when the customer created an account. For example, dev-332534.oktapreview.com. | 
**IssuerId** | **NullableString** | Specifies the IdP provided Issuer ID for the app. For example, exkh1aov1nhHrgFhN0h7. | 
**SfAccountId** | Pointer to **NullableString** | Specifies the salesforce account ID linked to this IDP. Either TenantId or SfAccountId should be set for IdP. | [optional] [readonly] 
**Certificate** | **NullableString** | Specifies the certificate generated for the app by the IdP service when the Helios is registered as an app. This is required to verify the SAML response. | 
**DefaultRoles** | Pointer to **[]string** | Specifies a list of default roles assigned to an IdP user if rolesSamlAttributeName is not given. | [optional] 
**DefaultClusters** | Pointer to **[]string** | Specifies a list of default clusterIds assigned to an IdP user if clustersSamlAttributeName is not given. &#39;All&#39; must be specified to give access to all clusters. | [optional] 
**SignRequest** | Pointer to **NullableBool** | Specifies whether to sign the SAML request or not. When it is set to true, SAML request will be signed. When it is set to false, SAML request is not signed. Default is false. Set this flag to true if the IdP site is configured to expect the SAML request from Helios signed. If this is set to true, users must get the Helios certificate and upload it on the IdP site. | [optional] 
**IsEnabled** | Pointer to **NullableBool** | Specifies a flag to enable or disable this IdP service. When it is set to true, IdP service is enabled. When it is set to false, IdP service is disabled. Default value is true. | [optional] 
**TenantId** | Pointer to **NullableString** | Specifies the Tenant Id if the IdP is configured for a Tenant. Either TenantId or SfAccountId should be set for IdP. | [optional] 
**Id** | Pointer to **NullableInt64** | Specifies the Id of the IDP Configuration. | [optional] 

## Methods

### NewIdp

`func NewIdp(name NullableString, domain NullableString, ssoUrl NullableString, issuerId NullableString, certificate NullableString, ) *Idp`

NewIdp instantiates a new Idp object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIdpWithDefaults

`func NewIdpWithDefaults() *Idp`

NewIdpWithDefaults instantiates a new Idp object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *Idp) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Idp) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Idp) SetName(v string)`

SetName sets Name field to given value.


### SetNameNil

`func (o *Idp) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *Idp) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetDomain

`func (o *Idp) GetDomain() string`

GetDomain returns the Domain field if non-nil, zero value otherwise.

### GetDomainOk

`func (o *Idp) GetDomainOk() (*string, bool)`

GetDomainOk returns a tuple with the Domain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomain

`func (o *Idp) SetDomain(v string)`

SetDomain sets Domain field to given value.


### SetDomainNil

`func (o *Idp) SetDomainNil(b bool)`

 SetDomainNil sets the value for Domain to be an explicit nil

### UnsetDomain
`func (o *Idp) UnsetDomain()`

UnsetDomain ensures that no value is present for Domain, not even an explicit nil
### GetSsoUrl

`func (o *Idp) GetSsoUrl() string`

GetSsoUrl returns the SsoUrl field if non-nil, zero value otherwise.

### GetSsoUrlOk

`func (o *Idp) GetSsoUrlOk() (*string, bool)`

GetSsoUrlOk returns a tuple with the SsoUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSsoUrl

`func (o *Idp) SetSsoUrl(v string)`

SetSsoUrl sets SsoUrl field to given value.


### SetSsoUrlNil

`func (o *Idp) SetSsoUrlNil(b bool)`

 SetSsoUrlNil sets the value for SsoUrl to be an explicit nil

### UnsetSsoUrl
`func (o *Idp) UnsetSsoUrl()`

UnsetSsoUrl ensures that no value is present for SsoUrl, not even an explicit nil
### GetIssuerId

`func (o *Idp) GetIssuerId() string`

GetIssuerId returns the IssuerId field if non-nil, zero value otherwise.

### GetIssuerIdOk

`func (o *Idp) GetIssuerIdOk() (*string, bool)`

GetIssuerIdOk returns a tuple with the IssuerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuerId

`func (o *Idp) SetIssuerId(v string)`

SetIssuerId sets IssuerId field to given value.


### SetIssuerIdNil

`func (o *Idp) SetIssuerIdNil(b bool)`

 SetIssuerIdNil sets the value for IssuerId to be an explicit nil

### UnsetIssuerId
`func (o *Idp) UnsetIssuerId()`

UnsetIssuerId ensures that no value is present for IssuerId, not even an explicit nil
### GetSfAccountId

`func (o *Idp) GetSfAccountId() string`

GetSfAccountId returns the SfAccountId field if non-nil, zero value otherwise.

### GetSfAccountIdOk

`func (o *Idp) GetSfAccountIdOk() (*string, bool)`

GetSfAccountIdOk returns a tuple with the SfAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSfAccountId

`func (o *Idp) SetSfAccountId(v string)`

SetSfAccountId sets SfAccountId field to given value.

### HasSfAccountId

`func (o *Idp) HasSfAccountId() bool`

HasSfAccountId returns a boolean if a field has been set.

### SetSfAccountIdNil

`func (o *Idp) SetSfAccountIdNil(b bool)`

 SetSfAccountIdNil sets the value for SfAccountId to be an explicit nil

### UnsetSfAccountId
`func (o *Idp) UnsetSfAccountId()`

UnsetSfAccountId ensures that no value is present for SfAccountId, not even an explicit nil
### GetCertificate

`func (o *Idp) GetCertificate() string`

GetCertificate returns the Certificate field if non-nil, zero value otherwise.

### GetCertificateOk

`func (o *Idp) GetCertificateOk() (*string, bool)`

GetCertificateOk returns a tuple with the Certificate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificate

`func (o *Idp) SetCertificate(v string)`

SetCertificate sets Certificate field to given value.


### SetCertificateNil

`func (o *Idp) SetCertificateNil(b bool)`

 SetCertificateNil sets the value for Certificate to be an explicit nil

### UnsetCertificate
`func (o *Idp) UnsetCertificate()`

UnsetCertificate ensures that no value is present for Certificate, not even an explicit nil
### GetDefaultRoles

`func (o *Idp) GetDefaultRoles() []string`

GetDefaultRoles returns the DefaultRoles field if non-nil, zero value otherwise.

### GetDefaultRolesOk

`func (o *Idp) GetDefaultRolesOk() (*[]string, bool)`

GetDefaultRolesOk returns a tuple with the DefaultRoles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultRoles

`func (o *Idp) SetDefaultRoles(v []string)`

SetDefaultRoles sets DefaultRoles field to given value.

### HasDefaultRoles

`func (o *Idp) HasDefaultRoles() bool`

HasDefaultRoles returns a boolean if a field has been set.

### SetDefaultRolesNil

`func (o *Idp) SetDefaultRolesNil(b bool)`

 SetDefaultRolesNil sets the value for DefaultRoles to be an explicit nil

### UnsetDefaultRoles
`func (o *Idp) UnsetDefaultRoles()`

UnsetDefaultRoles ensures that no value is present for DefaultRoles, not even an explicit nil
### GetDefaultClusters

`func (o *Idp) GetDefaultClusters() []string`

GetDefaultClusters returns the DefaultClusters field if non-nil, zero value otherwise.

### GetDefaultClustersOk

`func (o *Idp) GetDefaultClustersOk() (*[]string, bool)`

GetDefaultClustersOk returns a tuple with the DefaultClusters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultClusters

`func (o *Idp) SetDefaultClusters(v []string)`

SetDefaultClusters sets DefaultClusters field to given value.

### HasDefaultClusters

`func (o *Idp) HasDefaultClusters() bool`

HasDefaultClusters returns a boolean if a field has been set.

### SetDefaultClustersNil

`func (o *Idp) SetDefaultClustersNil(b bool)`

 SetDefaultClustersNil sets the value for DefaultClusters to be an explicit nil

### UnsetDefaultClusters
`func (o *Idp) UnsetDefaultClusters()`

UnsetDefaultClusters ensures that no value is present for DefaultClusters, not even an explicit nil
### GetSignRequest

`func (o *Idp) GetSignRequest() bool`

GetSignRequest returns the SignRequest field if non-nil, zero value otherwise.

### GetSignRequestOk

`func (o *Idp) GetSignRequestOk() (*bool, bool)`

GetSignRequestOk returns a tuple with the SignRequest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignRequest

`func (o *Idp) SetSignRequest(v bool)`

SetSignRequest sets SignRequest field to given value.

### HasSignRequest

`func (o *Idp) HasSignRequest() bool`

HasSignRequest returns a boolean if a field has been set.

### SetSignRequestNil

`func (o *Idp) SetSignRequestNil(b bool)`

 SetSignRequestNil sets the value for SignRequest to be an explicit nil

### UnsetSignRequest
`func (o *Idp) UnsetSignRequest()`

UnsetSignRequest ensures that no value is present for SignRequest, not even an explicit nil
### GetIsEnabled

`func (o *Idp) GetIsEnabled() bool`

GetIsEnabled returns the IsEnabled field if non-nil, zero value otherwise.

### GetIsEnabledOk

`func (o *Idp) GetIsEnabledOk() (*bool, bool)`

GetIsEnabledOk returns a tuple with the IsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsEnabled

`func (o *Idp) SetIsEnabled(v bool)`

SetIsEnabled sets IsEnabled field to given value.

### HasIsEnabled

`func (o *Idp) HasIsEnabled() bool`

HasIsEnabled returns a boolean if a field has been set.

### SetIsEnabledNil

`func (o *Idp) SetIsEnabledNil(b bool)`

 SetIsEnabledNil sets the value for IsEnabled to be an explicit nil

### UnsetIsEnabled
`func (o *Idp) UnsetIsEnabled()`

UnsetIsEnabled ensures that no value is present for IsEnabled, not even an explicit nil
### GetTenantId

`func (o *Idp) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *Idp) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *Idp) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.

### HasTenantId

`func (o *Idp) HasTenantId() bool`

HasTenantId returns a boolean if a field has been set.

### SetTenantIdNil

`func (o *Idp) SetTenantIdNil(b bool)`

 SetTenantIdNil sets the value for TenantId to be an explicit nil

### UnsetTenantId
`func (o *Idp) UnsetTenantId()`

UnsetTenantId ensures that no value is present for TenantId, not even an explicit nil
### GetId

`func (o *Idp) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Idp) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Idp) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *Idp) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *Idp) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *Idp) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


