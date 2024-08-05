# TrustedDomainParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BlacklistedDomains** | Pointer to **[]string** | Specifies a list of domains to add to blacklist. These domains will be blacklisted in trusted domain discovery. | [optional] 
**DiscoveryStatus** | Pointer to **NullableString** | Specifies the discovery status of trusted domains. | [optional] [readonly] 
**Enabled** | **NullableBool** | Specifies if trusted domain discovery is enabled. | 
**OnlyUseWhitelistedDomains** | Pointer to **NullableBool** | Specifies whether to use &#39;whitelistedDomains&#39; only for authentication. | [optional] 
**TaskIdentifier** | Pointer to **NullableString** | Specifies the identifier for the task running discovery. | [optional] [readonly] 
**TrustedDomains** | Pointer to [**[]TrustedDomain**](TrustedDomain.md) | Specifies a list of trusted domains. | [optional] 
**WhitelistedDomains** | Pointer to **[]string** | Specifies a list of domains to add to whitelist. Only these domains will be used for authentication if &#39;onlyUseWhitelistedDomains&#39; is set. | [optional] 

## Methods

### NewTrustedDomainParams

`func NewTrustedDomainParams(enabled NullableBool, ) *TrustedDomainParams`

NewTrustedDomainParams instantiates a new TrustedDomainParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTrustedDomainParamsWithDefaults

`func NewTrustedDomainParamsWithDefaults() *TrustedDomainParams`

NewTrustedDomainParamsWithDefaults instantiates a new TrustedDomainParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBlacklistedDomains

`func (o *TrustedDomainParams) GetBlacklistedDomains() []string`

GetBlacklistedDomains returns the BlacklistedDomains field if non-nil, zero value otherwise.

### GetBlacklistedDomainsOk

`func (o *TrustedDomainParams) GetBlacklistedDomainsOk() (*[]string, bool)`

GetBlacklistedDomainsOk returns a tuple with the BlacklistedDomains field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlacklistedDomains

`func (o *TrustedDomainParams) SetBlacklistedDomains(v []string)`

SetBlacklistedDomains sets BlacklistedDomains field to given value.

### HasBlacklistedDomains

`func (o *TrustedDomainParams) HasBlacklistedDomains() bool`

HasBlacklistedDomains returns a boolean if a field has been set.

### SetBlacklistedDomainsNil

`func (o *TrustedDomainParams) SetBlacklistedDomainsNil(b bool)`

 SetBlacklistedDomainsNil sets the value for BlacklistedDomains to be an explicit nil

### UnsetBlacklistedDomains
`func (o *TrustedDomainParams) UnsetBlacklistedDomains()`

UnsetBlacklistedDomains ensures that no value is present for BlacklistedDomains, not even an explicit nil
### GetDiscoveryStatus

`func (o *TrustedDomainParams) GetDiscoveryStatus() string`

GetDiscoveryStatus returns the DiscoveryStatus field if non-nil, zero value otherwise.

### GetDiscoveryStatusOk

`func (o *TrustedDomainParams) GetDiscoveryStatusOk() (*string, bool)`

GetDiscoveryStatusOk returns a tuple with the DiscoveryStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscoveryStatus

`func (o *TrustedDomainParams) SetDiscoveryStatus(v string)`

SetDiscoveryStatus sets DiscoveryStatus field to given value.

### HasDiscoveryStatus

`func (o *TrustedDomainParams) HasDiscoveryStatus() bool`

HasDiscoveryStatus returns a boolean if a field has been set.

### SetDiscoveryStatusNil

`func (o *TrustedDomainParams) SetDiscoveryStatusNil(b bool)`

 SetDiscoveryStatusNil sets the value for DiscoveryStatus to be an explicit nil

### UnsetDiscoveryStatus
`func (o *TrustedDomainParams) UnsetDiscoveryStatus()`

UnsetDiscoveryStatus ensures that no value is present for DiscoveryStatus, not even an explicit nil
### GetEnabled

`func (o *TrustedDomainParams) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *TrustedDomainParams) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *TrustedDomainParams) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### SetEnabledNil

`func (o *TrustedDomainParams) SetEnabledNil(b bool)`

 SetEnabledNil sets the value for Enabled to be an explicit nil

### UnsetEnabled
`func (o *TrustedDomainParams) UnsetEnabled()`

UnsetEnabled ensures that no value is present for Enabled, not even an explicit nil
### GetOnlyUseWhitelistedDomains

`func (o *TrustedDomainParams) GetOnlyUseWhitelistedDomains() bool`

GetOnlyUseWhitelistedDomains returns the OnlyUseWhitelistedDomains field if non-nil, zero value otherwise.

### GetOnlyUseWhitelistedDomainsOk

`func (o *TrustedDomainParams) GetOnlyUseWhitelistedDomainsOk() (*bool, bool)`

GetOnlyUseWhitelistedDomainsOk returns a tuple with the OnlyUseWhitelistedDomains field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnlyUseWhitelistedDomains

`func (o *TrustedDomainParams) SetOnlyUseWhitelistedDomains(v bool)`

SetOnlyUseWhitelistedDomains sets OnlyUseWhitelistedDomains field to given value.

### HasOnlyUseWhitelistedDomains

`func (o *TrustedDomainParams) HasOnlyUseWhitelistedDomains() bool`

HasOnlyUseWhitelistedDomains returns a boolean if a field has been set.

### SetOnlyUseWhitelistedDomainsNil

`func (o *TrustedDomainParams) SetOnlyUseWhitelistedDomainsNil(b bool)`

 SetOnlyUseWhitelistedDomainsNil sets the value for OnlyUseWhitelistedDomains to be an explicit nil

### UnsetOnlyUseWhitelistedDomains
`func (o *TrustedDomainParams) UnsetOnlyUseWhitelistedDomains()`

UnsetOnlyUseWhitelistedDomains ensures that no value is present for OnlyUseWhitelistedDomains, not even an explicit nil
### GetTaskIdentifier

`func (o *TrustedDomainParams) GetTaskIdentifier() string`

GetTaskIdentifier returns the TaskIdentifier field if non-nil, zero value otherwise.

### GetTaskIdentifierOk

`func (o *TrustedDomainParams) GetTaskIdentifierOk() (*string, bool)`

GetTaskIdentifierOk returns a tuple with the TaskIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskIdentifier

`func (o *TrustedDomainParams) SetTaskIdentifier(v string)`

SetTaskIdentifier sets TaskIdentifier field to given value.

### HasTaskIdentifier

`func (o *TrustedDomainParams) HasTaskIdentifier() bool`

HasTaskIdentifier returns a boolean if a field has been set.

### SetTaskIdentifierNil

`func (o *TrustedDomainParams) SetTaskIdentifierNil(b bool)`

 SetTaskIdentifierNil sets the value for TaskIdentifier to be an explicit nil

### UnsetTaskIdentifier
`func (o *TrustedDomainParams) UnsetTaskIdentifier()`

UnsetTaskIdentifier ensures that no value is present for TaskIdentifier, not even an explicit nil
### GetTrustedDomains

`func (o *TrustedDomainParams) GetTrustedDomains() []TrustedDomain`

GetTrustedDomains returns the TrustedDomains field if non-nil, zero value otherwise.

### GetTrustedDomainsOk

`func (o *TrustedDomainParams) GetTrustedDomainsOk() (*[]TrustedDomain, bool)`

GetTrustedDomainsOk returns a tuple with the TrustedDomains field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrustedDomains

`func (o *TrustedDomainParams) SetTrustedDomains(v []TrustedDomain)`

SetTrustedDomains sets TrustedDomains field to given value.

### HasTrustedDomains

`func (o *TrustedDomainParams) HasTrustedDomains() bool`

HasTrustedDomains returns a boolean if a field has been set.

### SetTrustedDomainsNil

`func (o *TrustedDomainParams) SetTrustedDomainsNil(b bool)`

 SetTrustedDomainsNil sets the value for TrustedDomains to be an explicit nil

### UnsetTrustedDomains
`func (o *TrustedDomainParams) UnsetTrustedDomains()`

UnsetTrustedDomains ensures that no value is present for TrustedDomains, not even an explicit nil
### GetWhitelistedDomains

`func (o *TrustedDomainParams) GetWhitelistedDomains() []string`

GetWhitelistedDomains returns the WhitelistedDomains field if non-nil, zero value otherwise.

### GetWhitelistedDomainsOk

`func (o *TrustedDomainParams) GetWhitelistedDomainsOk() (*[]string, bool)`

GetWhitelistedDomainsOk returns a tuple with the WhitelistedDomains field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWhitelistedDomains

`func (o *TrustedDomainParams) SetWhitelistedDomains(v []string)`

SetWhitelistedDomains sets WhitelistedDomains field to given value.

### HasWhitelistedDomains

`func (o *TrustedDomainParams) HasWhitelistedDomains() bool`

HasWhitelistedDomains returns a boolean if a field has been set.

### SetWhitelistedDomainsNil

`func (o *TrustedDomainParams) SetWhitelistedDomainsNil(b bool)`

 SetWhitelistedDomainsNil sets the value for WhitelistedDomains to be an explicit nil

### UnsetWhitelistedDomains
`func (o *TrustedDomainParams) UnsetWhitelistedDomains()`

UnsetWhitelistedDomains ensures that no value is present for WhitelistedDomains, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


