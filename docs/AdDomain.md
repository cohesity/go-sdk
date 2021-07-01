# AdDomain

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DnsRoot** | Pointer to **NullableString** | Specifies DNS root. | [optional] 
**Forest** | Pointer to **NullableString** | Specifies AD forest name. | [optional] 
**Identity** | Pointer to [**AdDomainIdentity**](AdDomainIdentity.md) |  | [optional] 
**NetbiosName** | Pointer to **NullableString** | Specifies AD NetBIOS name. | [optional] 
**ParentDomain** | Pointer to **NullableString** | Specifies parent domain name. | [optional] 
**TombstoneDays** | Pointer to **NullableInt32** | Specifies tombstone time in days. | [optional] 

## Methods

### NewAdDomain

`func NewAdDomain() *AdDomain`

NewAdDomain instantiates a new AdDomain object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAdDomainWithDefaults

`func NewAdDomainWithDefaults() *AdDomain`

NewAdDomainWithDefaults instantiates a new AdDomain object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDnsRoot

`func (o *AdDomain) GetDnsRoot() string`

GetDnsRoot returns the DnsRoot field if non-nil, zero value otherwise.

### GetDnsRootOk

`func (o *AdDomain) GetDnsRootOk() (*string, bool)`

GetDnsRootOk returns a tuple with the DnsRoot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsRoot

`func (o *AdDomain) SetDnsRoot(v string)`

SetDnsRoot sets DnsRoot field to given value.

### HasDnsRoot

`func (o *AdDomain) HasDnsRoot() bool`

HasDnsRoot returns a boolean if a field has been set.

### SetDnsRootNil

`func (o *AdDomain) SetDnsRootNil(b bool)`

 SetDnsRootNil sets the value for DnsRoot to be an explicit nil

### UnsetDnsRoot
`func (o *AdDomain) UnsetDnsRoot()`

UnsetDnsRoot ensures that no value is present for DnsRoot, not even an explicit nil
### GetForest

`func (o *AdDomain) GetForest() string`

GetForest returns the Forest field if non-nil, zero value otherwise.

### GetForestOk

`func (o *AdDomain) GetForestOk() (*string, bool)`

GetForestOk returns a tuple with the Forest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForest

`func (o *AdDomain) SetForest(v string)`

SetForest sets Forest field to given value.

### HasForest

`func (o *AdDomain) HasForest() bool`

HasForest returns a boolean if a field has been set.

### SetForestNil

`func (o *AdDomain) SetForestNil(b bool)`

 SetForestNil sets the value for Forest to be an explicit nil

### UnsetForest
`func (o *AdDomain) UnsetForest()`

UnsetForest ensures that no value is present for Forest, not even an explicit nil
### GetIdentity

`func (o *AdDomain) GetIdentity() AdDomainIdentity`

GetIdentity returns the Identity field if non-nil, zero value otherwise.

### GetIdentityOk

`func (o *AdDomain) GetIdentityOk() (*AdDomainIdentity, bool)`

GetIdentityOk returns a tuple with the Identity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentity

`func (o *AdDomain) SetIdentity(v AdDomainIdentity)`

SetIdentity sets Identity field to given value.

### HasIdentity

`func (o *AdDomain) HasIdentity() bool`

HasIdentity returns a boolean if a field has been set.

### GetNetbiosName

`func (o *AdDomain) GetNetbiosName() string`

GetNetbiosName returns the NetbiosName field if non-nil, zero value otherwise.

### GetNetbiosNameOk

`func (o *AdDomain) GetNetbiosNameOk() (*string, bool)`

GetNetbiosNameOk returns a tuple with the NetbiosName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetbiosName

`func (o *AdDomain) SetNetbiosName(v string)`

SetNetbiosName sets NetbiosName field to given value.

### HasNetbiosName

`func (o *AdDomain) HasNetbiosName() bool`

HasNetbiosName returns a boolean if a field has been set.

### SetNetbiosNameNil

`func (o *AdDomain) SetNetbiosNameNil(b bool)`

 SetNetbiosNameNil sets the value for NetbiosName to be an explicit nil

### UnsetNetbiosName
`func (o *AdDomain) UnsetNetbiosName()`

UnsetNetbiosName ensures that no value is present for NetbiosName, not even an explicit nil
### GetParentDomain

`func (o *AdDomain) GetParentDomain() string`

GetParentDomain returns the ParentDomain field if non-nil, zero value otherwise.

### GetParentDomainOk

`func (o *AdDomain) GetParentDomainOk() (*string, bool)`

GetParentDomainOk returns a tuple with the ParentDomain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentDomain

`func (o *AdDomain) SetParentDomain(v string)`

SetParentDomain sets ParentDomain field to given value.

### HasParentDomain

`func (o *AdDomain) HasParentDomain() bool`

HasParentDomain returns a boolean if a field has been set.

### SetParentDomainNil

`func (o *AdDomain) SetParentDomainNil(b bool)`

 SetParentDomainNil sets the value for ParentDomain to be an explicit nil

### UnsetParentDomain
`func (o *AdDomain) UnsetParentDomain()`

UnsetParentDomain ensures that no value is present for ParentDomain, not even an explicit nil
### GetTombstoneDays

`func (o *AdDomain) GetTombstoneDays() int32`

GetTombstoneDays returns the TombstoneDays field if non-nil, zero value otherwise.

### GetTombstoneDaysOk

`func (o *AdDomain) GetTombstoneDaysOk() (*int32, bool)`

GetTombstoneDaysOk returns a tuple with the TombstoneDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTombstoneDays

`func (o *AdDomain) SetTombstoneDays(v int32)`

SetTombstoneDays sets TombstoneDays field to given value.

### HasTombstoneDays

`func (o *AdDomain) HasTombstoneDays() bool`

HasTombstoneDays returns a boolean if a field has been set.

### SetTombstoneDaysNil

`func (o *AdDomain) SetTombstoneDaysNil(b bool)`

 SetTombstoneDaysNil sets the value for TombstoneDays to be an explicit nil

### UnsetTombstoneDays
`func (o *AdDomain) UnsetTombstoneDays()`

UnsetTombstoneDays ensures that no value is present for TombstoneDays, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


