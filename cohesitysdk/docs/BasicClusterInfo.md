# BasicClusterInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AuthenticationType** | Pointer to **NullableString** | Specifies the authentication scheme for the cluster. &#39;kPasswordOnly&#39; indicates the normal cohesity authentication type. &#39;kCertificateOnly&#39; indicates that certificate based authentication has been enabled and the password based authentication has been turned off. &#39;kPasswordAndCertificate&#39; indicates that both the authenticatio schemes are required. | [optional] 
**BannerEnabled** | Pointer to **NullableBool** | Specifies if banner is enabled on the cluster. | [optional] 
**ClusterSoftwareVersion** | Pointer to **NullableString** | This field is deprecated. Specifies the current release of the Cohesity software running on this Cohesity Cluster. deprecated: true | [optional] 
**ClusterType** | Pointer to **NullableString** | Specifies the type of Cohesity Cluster. &#39;kPhysical&#39; indicates the Cohesity Cluster is hosted directly on hardware. &#39;kVirtualRobo&#39; indicates the Cohesity Cluster is hosted in a VM on a ESXi Host of a VMware vCenter Server using Cohesity&#39;s Virtual Edition. &#39;kMicrosoftCloud&#39; indicates the Cohesity Cluster is hosted in a VM on Microsoft Azure using Cohesity&#39;s Cloud Edition. &#39;kAmazonCloud&#39; indicates the Cohesity Cluster is hosted in a VM on Amazon S3 using Cohesity&#39;s Cloud Edition. &#39;kGoogleCloud&#39; indicates the Cohesity Cluster is hosted in a VM on Google Cloud Platform using Cohesity&#39;s Cloud Edition. | [optional] 
**Domains** | Pointer to **[]string** | Array of Domains.  Specifies a list of domains joined to the Cohesity Cluster, including the default LOCAL Cohesity domain used to store the local Cohesity users. | [optional] 
**IdpConfigured** | Pointer to **NullableBool** | Specifies Idp is configured for the Cluster. | [optional] 
**IdpTenantExists** | Pointer to **NullableBool** | Specifies Idp is configured for a Tenant. | [optional] 
**LanguageLocale** | Pointer to **NullableString** | Specifies the language and locale for the Cohesity Cluster. | [optional] 
**McmMode** | Pointer to **NullableBool** | Specifies whether server is running in mcm-mode. If set to true, it is in mcm-mode. | [optional] 
**McmOnPremMode** | Pointer to **NullableBool** | Specifies whether server is running in mcm-on-prem-mode. If set to true, it is in mcm on prem mode. This need mcm-mode to be true. | [optional] 
**MultiTenancyEnabled** | Pointer to **NullableBool** | Specifies if multi-tenancy is enabled on the cluster. | [optional] 
**Name** | Pointer to **NullableString** | Specifies the name of the Cohesity Cluster. | [optional] 

## Methods

### NewBasicClusterInfo

`func NewBasicClusterInfo() *BasicClusterInfo`

NewBasicClusterInfo instantiates a new BasicClusterInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBasicClusterInfoWithDefaults

`func NewBasicClusterInfoWithDefaults() *BasicClusterInfo`

NewBasicClusterInfoWithDefaults instantiates a new BasicClusterInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuthenticationType

`func (o *BasicClusterInfo) GetAuthenticationType() string`

GetAuthenticationType returns the AuthenticationType field if non-nil, zero value otherwise.

### GetAuthenticationTypeOk

`func (o *BasicClusterInfo) GetAuthenticationTypeOk() (*string, bool)`

GetAuthenticationTypeOk returns a tuple with the AuthenticationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthenticationType

`func (o *BasicClusterInfo) SetAuthenticationType(v string)`

SetAuthenticationType sets AuthenticationType field to given value.

### HasAuthenticationType

`func (o *BasicClusterInfo) HasAuthenticationType() bool`

HasAuthenticationType returns a boolean if a field has been set.

### SetAuthenticationTypeNil

`func (o *BasicClusterInfo) SetAuthenticationTypeNil(b bool)`

 SetAuthenticationTypeNil sets the value for AuthenticationType to be an explicit nil

### UnsetAuthenticationType
`func (o *BasicClusterInfo) UnsetAuthenticationType()`

UnsetAuthenticationType ensures that no value is present for AuthenticationType, not even an explicit nil
### GetBannerEnabled

`func (o *BasicClusterInfo) GetBannerEnabled() bool`

GetBannerEnabled returns the BannerEnabled field if non-nil, zero value otherwise.

### GetBannerEnabledOk

`func (o *BasicClusterInfo) GetBannerEnabledOk() (*bool, bool)`

GetBannerEnabledOk returns a tuple with the BannerEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBannerEnabled

`func (o *BasicClusterInfo) SetBannerEnabled(v bool)`

SetBannerEnabled sets BannerEnabled field to given value.

### HasBannerEnabled

`func (o *BasicClusterInfo) HasBannerEnabled() bool`

HasBannerEnabled returns a boolean if a field has been set.

### SetBannerEnabledNil

`func (o *BasicClusterInfo) SetBannerEnabledNil(b bool)`

 SetBannerEnabledNil sets the value for BannerEnabled to be an explicit nil

### UnsetBannerEnabled
`func (o *BasicClusterInfo) UnsetBannerEnabled()`

UnsetBannerEnabled ensures that no value is present for BannerEnabled, not even an explicit nil
### GetClusterSoftwareVersion

`func (o *BasicClusterInfo) GetClusterSoftwareVersion() string`

GetClusterSoftwareVersion returns the ClusterSoftwareVersion field if non-nil, zero value otherwise.

### GetClusterSoftwareVersionOk

`func (o *BasicClusterInfo) GetClusterSoftwareVersionOk() (*string, bool)`

GetClusterSoftwareVersionOk returns a tuple with the ClusterSoftwareVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterSoftwareVersion

`func (o *BasicClusterInfo) SetClusterSoftwareVersion(v string)`

SetClusterSoftwareVersion sets ClusterSoftwareVersion field to given value.

### HasClusterSoftwareVersion

`func (o *BasicClusterInfo) HasClusterSoftwareVersion() bool`

HasClusterSoftwareVersion returns a boolean if a field has been set.

### SetClusterSoftwareVersionNil

`func (o *BasicClusterInfo) SetClusterSoftwareVersionNil(b bool)`

 SetClusterSoftwareVersionNil sets the value for ClusterSoftwareVersion to be an explicit nil

### UnsetClusterSoftwareVersion
`func (o *BasicClusterInfo) UnsetClusterSoftwareVersion()`

UnsetClusterSoftwareVersion ensures that no value is present for ClusterSoftwareVersion, not even an explicit nil
### GetClusterType

`func (o *BasicClusterInfo) GetClusterType() string`

GetClusterType returns the ClusterType field if non-nil, zero value otherwise.

### GetClusterTypeOk

`func (o *BasicClusterInfo) GetClusterTypeOk() (*string, bool)`

GetClusterTypeOk returns a tuple with the ClusterType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterType

`func (o *BasicClusterInfo) SetClusterType(v string)`

SetClusterType sets ClusterType field to given value.

### HasClusterType

`func (o *BasicClusterInfo) HasClusterType() bool`

HasClusterType returns a boolean if a field has been set.

### SetClusterTypeNil

`func (o *BasicClusterInfo) SetClusterTypeNil(b bool)`

 SetClusterTypeNil sets the value for ClusterType to be an explicit nil

### UnsetClusterType
`func (o *BasicClusterInfo) UnsetClusterType()`

UnsetClusterType ensures that no value is present for ClusterType, not even an explicit nil
### GetDomains

`func (o *BasicClusterInfo) GetDomains() []string`

GetDomains returns the Domains field if non-nil, zero value otherwise.

### GetDomainsOk

`func (o *BasicClusterInfo) GetDomainsOk() (*[]string, bool)`

GetDomainsOk returns a tuple with the Domains field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomains

`func (o *BasicClusterInfo) SetDomains(v []string)`

SetDomains sets Domains field to given value.

### HasDomains

`func (o *BasicClusterInfo) HasDomains() bool`

HasDomains returns a boolean if a field has been set.

### SetDomainsNil

`func (o *BasicClusterInfo) SetDomainsNil(b bool)`

 SetDomainsNil sets the value for Domains to be an explicit nil

### UnsetDomains
`func (o *BasicClusterInfo) UnsetDomains()`

UnsetDomains ensures that no value is present for Domains, not even an explicit nil
### GetIdpConfigured

`func (o *BasicClusterInfo) GetIdpConfigured() bool`

GetIdpConfigured returns the IdpConfigured field if non-nil, zero value otherwise.

### GetIdpConfiguredOk

`func (o *BasicClusterInfo) GetIdpConfiguredOk() (*bool, bool)`

GetIdpConfiguredOk returns a tuple with the IdpConfigured field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdpConfigured

`func (o *BasicClusterInfo) SetIdpConfigured(v bool)`

SetIdpConfigured sets IdpConfigured field to given value.

### HasIdpConfigured

`func (o *BasicClusterInfo) HasIdpConfigured() bool`

HasIdpConfigured returns a boolean if a field has been set.

### SetIdpConfiguredNil

`func (o *BasicClusterInfo) SetIdpConfiguredNil(b bool)`

 SetIdpConfiguredNil sets the value for IdpConfigured to be an explicit nil

### UnsetIdpConfigured
`func (o *BasicClusterInfo) UnsetIdpConfigured()`

UnsetIdpConfigured ensures that no value is present for IdpConfigured, not even an explicit nil
### GetIdpTenantExists

`func (o *BasicClusterInfo) GetIdpTenantExists() bool`

GetIdpTenantExists returns the IdpTenantExists field if non-nil, zero value otherwise.

### GetIdpTenantExistsOk

`func (o *BasicClusterInfo) GetIdpTenantExistsOk() (*bool, bool)`

GetIdpTenantExistsOk returns a tuple with the IdpTenantExists field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdpTenantExists

`func (o *BasicClusterInfo) SetIdpTenantExists(v bool)`

SetIdpTenantExists sets IdpTenantExists field to given value.

### HasIdpTenantExists

`func (o *BasicClusterInfo) HasIdpTenantExists() bool`

HasIdpTenantExists returns a boolean if a field has been set.

### SetIdpTenantExistsNil

`func (o *BasicClusterInfo) SetIdpTenantExistsNil(b bool)`

 SetIdpTenantExistsNil sets the value for IdpTenantExists to be an explicit nil

### UnsetIdpTenantExists
`func (o *BasicClusterInfo) UnsetIdpTenantExists()`

UnsetIdpTenantExists ensures that no value is present for IdpTenantExists, not even an explicit nil
### GetLanguageLocale

`func (o *BasicClusterInfo) GetLanguageLocale() string`

GetLanguageLocale returns the LanguageLocale field if non-nil, zero value otherwise.

### GetLanguageLocaleOk

`func (o *BasicClusterInfo) GetLanguageLocaleOk() (*string, bool)`

GetLanguageLocaleOk returns a tuple with the LanguageLocale field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanguageLocale

`func (o *BasicClusterInfo) SetLanguageLocale(v string)`

SetLanguageLocale sets LanguageLocale field to given value.

### HasLanguageLocale

`func (o *BasicClusterInfo) HasLanguageLocale() bool`

HasLanguageLocale returns a boolean if a field has been set.

### SetLanguageLocaleNil

`func (o *BasicClusterInfo) SetLanguageLocaleNil(b bool)`

 SetLanguageLocaleNil sets the value for LanguageLocale to be an explicit nil

### UnsetLanguageLocale
`func (o *BasicClusterInfo) UnsetLanguageLocale()`

UnsetLanguageLocale ensures that no value is present for LanguageLocale, not even an explicit nil
### GetMcmMode

`func (o *BasicClusterInfo) GetMcmMode() bool`

GetMcmMode returns the McmMode field if non-nil, zero value otherwise.

### GetMcmModeOk

`func (o *BasicClusterInfo) GetMcmModeOk() (*bool, bool)`

GetMcmModeOk returns a tuple with the McmMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMcmMode

`func (o *BasicClusterInfo) SetMcmMode(v bool)`

SetMcmMode sets McmMode field to given value.

### HasMcmMode

`func (o *BasicClusterInfo) HasMcmMode() bool`

HasMcmMode returns a boolean if a field has been set.

### SetMcmModeNil

`func (o *BasicClusterInfo) SetMcmModeNil(b bool)`

 SetMcmModeNil sets the value for McmMode to be an explicit nil

### UnsetMcmMode
`func (o *BasicClusterInfo) UnsetMcmMode()`

UnsetMcmMode ensures that no value is present for McmMode, not even an explicit nil
### GetMcmOnPremMode

`func (o *BasicClusterInfo) GetMcmOnPremMode() bool`

GetMcmOnPremMode returns the McmOnPremMode field if non-nil, zero value otherwise.

### GetMcmOnPremModeOk

`func (o *BasicClusterInfo) GetMcmOnPremModeOk() (*bool, bool)`

GetMcmOnPremModeOk returns a tuple with the McmOnPremMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMcmOnPremMode

`func (o *BasicClusterInfo) SetMcmOnPremMode(v bool)`

SetMcmOnPremMode sets McmOnPremMode field to given value.

### HasMcmOnPremMode

`func (o *BasicClusterInfo) HasMcmOnPremMode() bool`

HasMcmOnPremMode returns a boolean if a field has been set.

### SetMcmOnPremModeNil

`func (o *BasicClusterInfo) SetMcmOnPremModeNil(b bool)`

 SetMcmOnPremModeNil sets the value for McmOnPremMode to be an explicit nil

### UnsetMcmOnPremMode
`func (o *BasicClusterInfo) UnsetMcmOnPremMode()`

UnsetMcmOnPremMode ensures that no value is present for McmOnPremMode, not even an explicit nil
### GetMultiTenancyEnabled

`func (o *BasicClusterInfo) GetMultiTenancyEnabled() bool`

GetMultiTenancyEnabled returns the MultiTenancyEnabled field if non-nil, zero value otherwise.

### GetMultiTenancyEnabledOk

`func (o *BasicClusterInfo) GetMultiTenancyEnabledOk() (*bool, bool)`

GetMultiTenancyEnabledOk returns a tuple with the MultiTenancyEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMultiTenancyEnabled

`func (o *BasicClusterInfo) SetMultiTenancyEnabled(v bool)`

SetMultiTenancyEnabled sets MultiTenancyEnabled field to given value.

### HasMultiTenancyEnabled

`func (o *BasicClusterInfo) HasMultiTenancyEnabled() bool`

HasMultiTenancyEnabled returns a boolean if a field has been set.

### SetMultiTenancyEnabledNil

`func (o *BasicClusterInfo) SetMultiTenancyEnabledNil(b bool)`

 SetMultiTenancyEnabledNil sets the value for MultiTenancyEnabled to be an explicit nil

### UnsetMultiTenancyEnabled
`func (o *BasicClusterInfo) UnsetMultiTenancyEnabled()`

UnsetMultiTenancyEnabled ensures that no value is present for MultiTenancyEnabled, not even an explicit nil
### GetName

`func (o *BasicClusterInfo) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BasicClusterInfo) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BasicClusterInfo) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *BasicClusterInfo) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *BasicClusterInfo) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *BasicClusterInfo) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


