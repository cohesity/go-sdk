# CommonExternalTargetParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CloudDomains** | Pointer to [**[]CloudDomain**](CloudDomain.md) | Specifies the cloud domain information. | [optional] 
**Compression** | Pointer to **NullableString** | Specifies whether the type of compression of the External Target | [optional] 
**EnableObjectLock** | Pointer to **NullableBool** | Whether to enable object lock for this vault. If this field is set, all the objects written to the vault will be object locked until all the archives referring to them expire. | [optional] 
**ErrorMessage** | Pointer to **NullableString** | Specifies the error message if the event is in failed state. | [optional] [readonly] 
**GlobalId** | Pointer to **NullableString** | Specifies the global identifier of the External Target. | [optional] 
**Id** | Pointer to **NullableInt64** | Specifies the ID of the External Target. | [optional] [readonly] 
**IsWormCapable** | Pointer to **NullableBool** | Specifies whether this external target has been found to be capable of supporting WORM archives. | [optional] 
**Name** | **NullableString** | Specifies the name of the External Target. | 
**OwnershipContext** | Pointer to **NullableString** | Specifies whether how this external target is being consumed either Local or FortKnox. | [optional] 
**PurposeType** | **NullableString** | Specifies the purpose of the External Target. | 
**Status** | Pointer to **NullableString** | Specifies the registration status of the External Target | [optional] [readonly] 
**StorageDomainName** | Pointer to **NullableString** | Specifies the storage domain associated with the target. | [optional] 
**TenantIds** | Pointer to **[]string** | Specifies the list of tenantIds for the External Target | [optional] 
**UseForApolloMrStore** | Pointer to **NullableBool** | Specifies whether this external target is used to store apollo mr records. | [optional] 

## Methods

### NewCommonExternalTargetParams

`func NewCommonExternalTargetParams(name NullableString, purposeType NullableString, ) *CommonExternalTargetParams`

NewCommonExternalTargetParams instantiates a new CommonExternalTargetParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCommonExternalTargetParamsWithDefaults

`func NewCommonExternalTargetParamsWithDefaults() *CommonExternalTargetParams`

NewCommonExternalTargetParamsWithDefaults instantiates a new CommonExternalTargetParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCloudDomains

`func (o *CommonExternalTargetParams) GetCloudDomains() []CloudDomain`

GetCloudDomains returns the CloudDomains field if non-nil, zero value otherwise.

### GetCloudDomainsOk

`func (o *CommonExternalTargetParams) GetCloudDomainsOk() (*[]CloudDomain, bool)`

GetCloudDomainsOk returns a tuple with the CloudDomains field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudDomains

`func (o *CommonExternalTargetParams) SetCloudDomains(v []CloudDomain)`

SetCloudDomains sets CloudDomains field to given value.

### HasCloudDomains

`func (o *CommonExternalTargetParams) HasCloudDomains() bool`

HasCloudDomains returns a boolean if a field has been set.

### SetCloudDomainsNil

`func (o *CommonExternalTargetParams) SetCloudDomainsNil(b bool)`

 SetCloudDomainsNil sets the value for CloudDomains to be an explicit nil

### UnsetCloudDomains
`func (o *CommonExternalTargetParams) UnsetCloudDomains()`

UnsetCloudDomains ensures that no value is present for CloudDomains, not even an explicit nil
### GetCompression

`func (o *CommonExternalTargetParams) GetCompression() string`

GetCompression returns the Compression field if non-nil, zero value otherwise.

### GetCompressionOk

`func (o *CommonExternalTargetParams) GetCompressionOk() (*string, bool)`

GetCompressionOk returns a tuple with the Compression field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompression

`func (o *CommonExternalTargetParams) SetCompression(v string)`

SetCompression sets Compression field to given value.

### HasCompression

`func (o *CommonExternalTargetParams) HasCompression() bool`

HasCompression returns a boolean if a field has been set.

### SetCompressionNil

`func (o *CommonExternalTargetParams) SetCompressionNil(b bool)`

 SetCompressionNil sets the value for Compression to be an explicit nil

### UnsetCompression
`func (o *CommonExternalTargetParams) UnsetCompression()`

UnsetCompression ensures that no value is present for Compression, not even an explicit nil
### GetEnableObjectLock

`func (o *CommonExternalTargetParams) GetEnableObjectLock() bool`

GetEnableObjectLock returns the EnableObjectLock field if non-nil, zero value otherwise.

### GetEnableObjectLockOk

`func (o *CommonExternalTargetParams) GetEnableObjectLockOk() (*bool, bool)`

GetEnableObjectLockOk returns a tuple with the EnableObjectLock field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableObjectLock

`func (o *CommonExternalTargetParams) SetEnableObjectLock(v bool)`

SetEnableObjectLock sets EnableObjectLock field to given value.

### HasEnableObjectLock

`func (o *CommonExternalTargetParams) HasEnableObjectLock() bool`

HasEnableObjectLock returns a boolean if a field has been set.

### SetEnableObjectLockNil

`func (o *CommonExternalTargetParams) SetEnableObjectLockNil(b bool)`

 SetEnableObjectLockNil sets the value for EnableObjectLock to be an explicit nil

### UnsetEnableObjectLock
`func (o *CommonExternalTargetParams) UnsetEnableObjectLock()`

UnsetEnableObjectLock ensures that no value is present for EnableObjectLock, not even an explicit nil
### GetErrorMessage

`func (o *CommonExternalTargetParams) GetErrorMessage() string`

GetErrorMessage returns the ErrorMessage field if non-nil, zero value otherwise.

### GetErrorMessageOk

`func (o *CommonExternalTargetParams) GetErrorMessageOk() (*string, bool)`

GetErrorMessageOk returns a tuple with the ErrorMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorMessage

`func (o *CommonExternalTargetParams) SetErrorMessage(v string)`

SetErrorMessage sets ErrorMessage field to given value.

### HasErrorMessage

`func (o *CommonExternalTargetParams) HasErrorMessage() bool`

HasErrorMessage returns a boolean if a field has been set.

### SetErrorMessageNil

`func (o *CommonExternalTargetParams) SetErrorMessageNil(b bool)`

 SetErrorMessageNil sets the value for ErrorMessage to be an explicit nil

### UnsetErrorMessage
`func (o *CommonExternalTargetParams) UnsetErrorMessage()`

UnsetErrorMessage ensures that no value is present for ErrorMessage, not even an explicit nil
### GetGlobalId

`func (o *CommonExternalTargetParams) GetGlobalId() string`

GetGlobalId returns the GlobalId field if non-nil, zero value otherwise.

### GetGlobalIdOk

`func (o *CommonExternalTargetParams) GetGlobalIdOk() (*string, bool)`

GetGlobalIdOk returns a tuple with the GlobalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGlobalId

`func (o *CommonExternalTargetParams) SetGlobalId(v string)`

SetGlobalId sets GlobalId field to given value.

### HasGlobalId

`func (o *CommonExternalTargetParams) HasGlobalId() bool`

HasGlobalId returns a boolean if a field has been set.

### SetGlobalIdNil

`func (o *CommonExternalTargetParams) SetGlobalIdNil(b bool)`

 SetGlobalIdNil sets the value for GlobalId to be an explicit nil

### UnsetGlobalId
`func (o *CommonExternalTargetParams) UnsetGlobalId()`

UnsetGlobalId ensures that no value is present for GlobalId, not even an explicit nil
### GetId

`func (o *CommonExternalTargetParams) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CommonExternalTargetParams) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CommonExternalTargetParams) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *CommonExternalTargetParams) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *CommonExternalTargetParams) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *CommonExternalTargetParams) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetIsWormCapable

`func (o *CommonExternalTargetParams) GetIsWormCapable() bool`

GetIsWormCapable returns the IsWormCapable field if non-nil, zero value otherwise.

### GetIsWormCapableOk

`func (o *CommonExternalTargetParams) GetIsWormCapableOk() (*bool, bool)`

GetIsWormCapableOk returns a tuple with the IsWormCapable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsWormCapable

`func (o *CommonExternalTargetParams) SetIsWormCapable(v bool)`

SetIsWormCapable sets IsWormCapable field to given value.

### HasIsWormCapable

`func (o *CommonExternalTargetParams) HasIsWormCapable() bool`

HasIsWormCapable returns a boolean if a field has been set.

### SetIsWormCapableNil

`func (o *CommonExternalTargetParams) SetIsWormCapableNil(b bool)`

 SetIsWormCapableNil sets the value for IsWormCapable to be an explicit nil

### UnsetIsWormCapable
`func (o *CommonExternalTargetParams) UnsetIsWormCapable()`

UnsetIsWormCapable ensures that no value is present for IsWormCapable, not even an explicit nil
### GetName

`func (o *CommonExternalTargetParams) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CommonExternalTargetParams) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CommonExternalTargetParams) SetName(v string)`

SetName sets Name field to given value.


### SetNameNil

`func (o *CommonExternalTargetParams) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *CommonExternalTargetParams) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetOwnershipContext

`func (o *CommonExternalTargetParams) GetOwnershipContext() string`

GetOwnershipContext returns the OwnershipContext field if non-nil, zero value otherwise.

### GetOwnershipContextOk

`func (o *CommonExternalTargetParams) GetOwnershipContextOk() (*string, bool)`

GetOwnershipContextOk returns a tuple with the OwnershipContext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnershipContext

`func (o *CommonExternalTargetParams) SetOwnershipContext(v string)`

SetOwnershipContext sets OwnershipContext field to given value.

### HasOwnershipContext

`func (o *CommonExternalTargetParams) HasOwnershipContext() bool`

HasOwnershipContext returns a boolean if a field has been set.

### SetOwnershipContextNil

`func (o *CommonExternalTargetParams) SetOwnershipContextNil(b bool)`

 SetOwnershipContextNil sets the value for OwnershipContext to be an explicit nil

### UnsetOwnershipContext
`func (o *CommonExternalTargetParams) UnsetOwnershipContext()`

UnsetOwnershipContext ensures that no value is present for OwnershipContext, not even an explicit nil
### GetPurposeType

`func (o *CommonExternalTargetParams) GetPurposeType() string`

GetPurposeType returns the PurposeType field if non-nil, zero value otherwise.

### GetPurposeTypeOk

`func (o *CommonExternalTargetParams) GetPurposeTypeOk() (*string, bool)`

GetPurposeTypeOk returns a tuple with the PurposeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPurposeType

`func (o *CommonExternalTargetParams) SetPurposeType(v string)`

SetPurposeType sets PurposeType field to given value.


### SetPurposeTypeNil

`func (o *CommonExternalTargetParams) SetPurposeTypeNil(b bool)`

 SetPurposeTypeNil sets the value for PurposeType to be an explicit nil

### UnsetPurposeType
`func (o *CommonExternalTargetParams) UnsetPurposeType()`

UnsetPurposeType ensures that no value is present for PurposeType, not even an explicit nil
### GetStatus

`func (o *CommonExternalTargetParams) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *CommonExternalTargetParams) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *CommonExternalTargetParams) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *CommonExternalTargetParams) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### SetStatusNil

`func (o *CommonExternalTargetParams) SetStatusNil(b bool)`

 SetStatusNil sets the value for Status to be an explicit nil

### UnsetStatus
`func (o *CommonExternalTargetParams) UnsetStatus()`

UnsetStatus ensures that no value is present for Status, not even an explicit nil
### GetStorageDomainName

`func (o *CommonExternalTargetParams) GetStorageDomainName() string`

GetStorageDomainName returns the StorageDomainName field if non-nil, zero value otherwise.

### GetStorageDomainNameOk

`func (o *CommonExternalTargetParams) GetStorageDomainNameOk() (*string, bool)`

GetStorageDomainNameOk returns a tuple with the StorageDomainName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageDomainName

`func (o *CommonExternalTargetParams) SetStorageDomainName(v string)`

SetStorageDomainName sets StorageDomainName field to given value.

### HasStorageDomainName

`func (o *CommonExternalTargetParams) HasStorageDomainName() bool`

HasStorageDomainName returns a boolean if a field has been set.

### SetStorageDomainNameNil

`func (o *CommonExternalTargetParams) SetStorageDomainNameNil(b bool)`

 SetStorageDomainNameNil sets the value for StorageDomainName to be an explicit nil

### UnsetStorageDomainName
`func (o *CommonExternalTargetParams) UnsetStorageDomainName()`

UnsetStorageDomainName ensures that no value is present for StorageDomainName, not even an explicit nil
### GetTenantIds

`func (o *CommonExternalTargetParams) GetTenantIds() []string`

GetTenantIds returns the TenantIds field if non-nil, zero value otherwise.

### GetTenantIdsOk

`func (o *CommonExternalTargetParams) GetTenantIdsOk() (*[]string, bool)`

GetTenantIdsOk returns a tuple with the TenantIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantIds

`func (o *CommonExternalTargetParams) SetTenantIds(v []string)`

SetTenantIds sets TenantIds field to given value.

### HasTenantIds

`func (o *CommonExternalTargetParams) HasTenantIds() bool`

HasTenantIds returns a boolean if a field has been set.

### GetUseForApolloMrStore

`func (o *CommonExternalTargetParams) GetUseForApolloMrStore() bool`

GetUseForApolloMrStore returns the UseForApolloMrStore field if non-nil, zero value otherwise.

### GetUseForApolloMrStoreOk

`func (o *CommonExternalTargetParams) GetUseForApolloMrStoreOk() (*bool, bool)`

GetUseForApolloMrStoreOk returns a tuple with the UseForApolloMrStore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseForApolloMrStore

`func (o *CommonExternalTargetParams) SetUseForApolloMrStore(v bool)`

SetUseForApolloMrStore sets UseForApolloMrStore field to given value.

### HasUseForApolloMrStore

`func (o *CommonExternalTargetParams) HasUseForApolloMrStore() bool`

HasUseForApolloMrStore returns a boolean if a field has been set.

### SetUseForApolloMrStoreNil

`func (o *CommonExternalTargetParams) SetUseForApolloMrStoreNil(b bool)`

 SetUseForApolloMrStoreNil sets the value for UseForApolloMrStore to be an explicit nil

### UnsetUseForApolloMrStore
`func (o *CommonExternalTargetParams) UnsetUseForApolloMrStore()`

UnsetUseForApolloMrStore ensures that no value is present for UseForApolloMrStore, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


