# ExternalTarget

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
**ArchivalParams** | Pointer to [**ArchivalExternalTargetParams**](ArchivalExternalTargetParams.md) |  | [optional] 
**TieringParams** | Pointer to [**TieringExternalTargetParams**](TieringExternalTargetParams.md) |  | [optional] 

## Methods

### NewExternalTarget

`func NewExternalTarget(name NullableString, purposeType NullableString, ) *ExternalTarget`

NewExternalTarget instantiates a new ExternalTarget object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExternalTargetWithDefaults

`func NewExternalTargetWithDefaults() *ExternalTarget`

NewExternalTargetWithDefaults instantiates a new ExternalTarget object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCloudDomains

`func (o *ExternalTarget) GetCloudDomains() []CloudDomain`

GetCloudDomains returns the CloudDomains field if non-nil, zero value otherwise.

### GetCloudDomainsOk

`func (o *ExternalTarget) GetCloudDomainsOk() (*[]CloudDomain, bool)`

GetCloudDomainsOk returns a tuple with the CloudDomains field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudDomains

`func (o *ExternalTarget) SetCloudDomains(v []CloudDomain)`

SetCloudDomains sets CloudDomains field to given value.

### HasCloudDomains

`func (o *ExternalTarget) HasCloudDomains() bool`

HasCloudDomains returns a boolean if a field has been set.

### SetCloudDomainsNil

`func (o *ExternalTarget) SetCloudDomainsNil(b bool)`

 SetCloudDomainsNil sets the value for CloudDomains to be an explicit nil

### UnsetCloudDomains
`func (o *ExternalTarget) UnsetCloudDomains()`

UnsetCloudDomains ensures that no value is present for CloudDomains, not even an explicit nil
### GetCompression

`func (o *ExternalTarget) GetCompression() string`

GetCompression returns the Compression field if non-nil, zero value otherwise.

### GetCompressionOk

`func (o *ExternalTarget) GetCompressionOk() (*string, bool)`

GetCompressionOk returns a tuple with the Compression field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompression

`func (o *ExternalTarget) SetCompression(v string)`

SetCompression sets Compression field to given value.

### HasCompression

`func (o *ExternalTarget) HasCompression() bool`

HasCompression returns a boolean if a field has been set.

### SetCompressionNil

`func (o *ExternalTarget) SetCompressionNil(b bool)`

 SetCompressionNil sets the value for Compression to be an explicit nil

### UnsetCompression
`func (o *ExternalTarget) UnsetCompression()`

UnsetCompression ensures that no value is present for Compression, not even an explicit nil
### GetEnableObjectLock

`func (o *ExternalTarget) GetEnableObjectLock() bool`

GetEnableObjectLock returns the EnableObjectLock field if non-nil, zero value otherwise.

### GetEnableObjectLockOk

`func (o *ExternalTarget) GetEnableObjectLockOk() (*bool, bool)`

GetEnableObjectLockOk returns a tuple with the EnableObjectLock field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableObjectLock

`func (o *ExternalTarget) SetEnableObjectLock(v bool)`

SetEnableObjectLock sets EnableObjectLock field to given value.

### HasEnableObjectLock

`func (o *ExternalTarget) HasEnableObjectLock() bool`

HasEnableObjectLock returns a boolean if a field has been set.

### SetEnableObjectLockNil

`func (o *ExternalTarget) SetEnableObjectLockNil(b bool)`

 SetEnableObjectLockNil sets the value for EnableObjectLock to be an explicit nil

### UnsetEnableObjectLock
`func (o *ExternalTarget) UnsetEnableObjectLock()`

UnsetEnableObjectLock ensures that no value is present for EnableObjectLock, not even an explicit nil
### GetErrorMessage

`func (o *ExternalTarget) GetErrorMessage() string`

GetErrorMessage returns the ErrorMessage field if non-nil, zero value otherwise.

### GetErrorMessageOk

`func (o *ExternalTarget) GetErrorMessageOk() (*string, bool)`

GetErrorMessageOk returns a tuple with the ErrorMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorMessage

`func (o *ExternalTarget) SetErrorMessage(v string)`

SetErrorMessage sets ErrorMessage field to given value.

### HasErrorMessage

`func (o *ExternalTarget) HasErrorMessage() bool`

HasErrorMessage returns a boolean if a field has been set.

### SetErrorMessageNil

`func (o *ExternalTarget) SetErrorMessageNil(b bool)`

 SetErrorMessageNil sets the value for ErrorMessage to be an explicit nil

### UnsetErrorMessage
`func (o *ExternalTarget) UnsetErrorMessage()`

UnsetErrorMessage ensures that no value is present for ErrorMessage, not even an explicit nil
### GetGlobalId

`func (o *ExternalTarget) GetGlobalId() string`

GetGlobalId returns the GlobalId field if non-nil, zero value otherwise.

### GetGlobalIdOk

`func (o *ExternalTarget) GetGlobalIdOk() (*string, bool)`

GetGlobalIdOk returns a tuple with the GlobalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGlobalId

`func (o *ExternalTarget) SetGlobalId(v string)`

SetGlobalId sets GlobalId field to given value.

### HasGlobalId

`func (o *ExternalTarget) HasGlobalId() bool`

HasGlobalId returns a boolean if a field has been set.

### SetGlobalIdNil

`func (o *ExternalTarget) SetGlobalIdNil(b bool)`

 SetGlobalIdNil sets the value for GlobalId to be an explicit nil

### UnsetGlobalId
`func (o *ExternalTarget) UnsetGlobalId()`

UnsetGlobalId ensures that no value is present for GlobalId, not even an explicit nil
### GetId

`func (o *ExternalTarget) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ExternalTarget) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ExternalTarget) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *ExternalTarget) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *ExternalTarget) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *ExternalTarget) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetIsWormCapable

`func (o *ExternalTarget) GetIsWormCapable() bool`

GetIsWormCapable returns the IsWormCapable field if non-nil, zero value otherwise.

### GetIsWormCapableOk

`func (o *ExternalTarget) GetIsWormCapableOk() (*bool, bool)`

GetIsWormCapableOk returns a tuple with the IsWormCapable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsWormCapable

`func (o *ExternalTarget) SetIsWormCapable(v bool)`

SetIsWormCapable sets IsWormCapable field to given value.

### HasIsWormCapable

`func (o *ExternalTarget) HasIsWormCapable() bool`

HasIsWormCapable returns a boolean if a field has been set.

### SetIsWormCapableNil

`func (o *ExternalTarget) SetIsWormCapableNil(b bool)`

 SetIsWormCapableNil sets the value for IsWormCapable to be an explicit nil

### UnsetIsWormCapable
`func (o *ExternalTarget) UnsetIsWormCapable()`

UnsetIsWormCapable ensures that no value is present for IsWormCapable, not even an explicit nil
### GetName

`func (o *ExternalTarget) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ExternalTarget) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ExternalTarget) SetName(v string)`

SetName sets Name field to given value.


### SetNameNil

`func (o *ExternalTarget) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *ExternalTarget) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetOwnershipContext

`func (o *ExternalTarget) GetOwnershipContext() string`

GetOwnershipContext returns the OwnershipContext field if non-nil, zero value otherwise.

### GetOwnershipContextOk

`func (o *ExternalTarget) GetOwnershipContextOk() (*string, bool)`

GetOwnershipContextOk returns a tuple with the OwnershipContext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnershipContext

`func (o *ExternalTarget) SetOwnershipContext(v string)`

SetOwnershipContext sets OwnershipContext field to given value.

### HasOwnershipContext

`func (o *ExternalTarget) HasOwnershipContext() bool`

HasOwnershipContext returns a boolean if a field has been set.

### SetOwnershipContextNil

`func (o *ExternalTarget) SetOwnershipContextNil(b bool)`

 SetOwnershipContextNil sets the value for OwnershipContext to be an explicit nil

### UnsetOwnershipContext
`func (o *ExternalTarget) UnsetOwnershipContext()`

UnsetOwnershipContext ensures that no value is present for OwnershipContext, not even an explicit nil
### GetPurposeType

`func (o *ExternalTarget) GetPurposeType() string`

GetPurposeType returns the PurposeType field if non-nil, zero value otherwise.

### GetPurposeTypeOk

`func (o *ExternalTarget) GetPurposeTypeOk() (*string, bool)`

GetPurposeTypeOk returns a tuple with the PurposeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPurposeType

`func (o *ExternalTarget) SetPurposeType(v string)`

SetPurposeType sets PurposeType field to given value.


### SetPurposeTypeNil

`func (o *ExternalTarget) SetPurposeTypeNil(b bool)`

 SetPurposeTypeNil sets the value for PurposeType to be an explicit nil

### UnsetPurposeType
`func (o *ExternalTarget) UnsetPurposeType()`

UnsetPurposeType ensures that no value is present for PurposeType, not even an explicit nil
### GetStatus

`func (o *ExternalTarget) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ExternalTarget) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ExternalTarget) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ExternalTarget) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### SetStatusNil

`func (o *ExternalTarget) SetStatusNil(b bool)`

 SetStatusNil sets the value for Status to be an explicit nil

### UnsetStatus
`func (o *ExternalTarget) UnsetStatus()`

UnsetStatus ensures that no value is present for Status, not even an explicit nil
### GetStorageDomainName

`func (o *ExternalTarget) GetStorageDomainName() string`

GetStorageDomainName returns the StorageDomainName field if non-nil, zero value otherwise.

### GetStorageDomainNameOk

`func (o *ExternalTarget) GetStorageDomainNameOk() (*string, bool)`

GetStorageDomainNameOk returns a tuple with the StorageDomainName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageDomainName

`func (o *ExternalTarget) SetStorageDomainName(v string)`

SetStorageDomainName sets StorageDomainName field to given value.

### HasStorageDomainName

`func (o *ExternalTarget) HasStorageDomainName() bool`

HasStorageDomainName returns a boolean if a field has been set.

### SetStorageDomainNameNil

`func (o *ExternalTarget) SetStorageDomainNameNil(b bool)`

 SetStorageDomainNameNil sets the value for StorageDomainName to be an explicit nil

### UnsetStorageDomainName
`func (o *ExternalTarget) UnsetStorageDomainName()`

UnsetStorageDomainName ensures that no value is present for StorageDomainName, not even an explicit nil
### GetTenantIds

`func (o *ExternalTarget) GetTenantIds() []string`

GetTenantIds returns the TenantIds field if non-nil, zero value otherwise.

### GetTenantIdsOk

`func (o *ExternalTarget) GetTenantIdsOk() (*[]string, bool)`

GetTenantIdsOk returns a tuple with the TenantIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantIds

`func (o *ExternalTarget) SetTenantIds(v []string)`

SetTenantIds sets TenantIds field to given value.

### HasTenantIds

`func (o *ExternalTarget) HasTenantIds() bool`

HasTenantIds returns a boolean if a field has been set.

### GetUseForApolloMrStore

`func (o *ExternalTarget) GetUseForApolloMrStore() bool`

GetUseForApolloMrStore returns the UseForApolloMrStore field if non-nil, zero value otherwise.

### GetUseForApolloMrStoreOk

`func (o *ExternalTarget) GetUseForApolloMrStoreOk() (*bool, bool)`

GetUseForApolloMrStoreOk returns a tuple with the UseForApolloMrStore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseForApolloMrStore

`func (o *ExternalTarget) SetUseForApolloMrStore(v bool)`

SetUseForApolloMrStore sets UseForApolloMrStore field to given value.

### HasUseForApolloMrStore

`func (o *ExternalTarget) HasUseForApolloMrStore() bool`

HasUseForApolloMrStore returns a boolean if a field has been set.

### SetUseForApolloMrStoreNil

`func (o *ExternalTarget) SetUseForApolloMrStoreNil(b bool)`

 SetUseForApolloMrStoreNil sets the value for UseForApolloMrStore to be an explicit nil

### UnsetUseForApolloMrStore
`func (o *ExternalTarget) UnsetUseForApolloMrStore()`

UnsetUseForApolloMrStore ensures that no value is present for UseForApolloMrStore, not even an explicit nil
### GetArchivalParams

`func (o *ExternalTarget) GetArchivalParams() ArchivalExternalTargetParams`

GetArchivalParams returns the ArchivalParams field if non-nil, zero value otherwise.

### GetArchivalParamsOk

`func (o *ExternalTarget) GetArchivalParamsOk() (*ArchivalExternalTargetParams, bool)`

GetArchivalParamsOk returns a tuple with the ArchivalParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchivalParams

`func (o *ExternalTarget) SetArchivalParams(v ArchivalExternalTargetParams)`

SetArchivalParams sets ArchivalParams field to given value.

### HasArchivalParams

`func (o *ExternalTarget) HasArchivalParams() bool`

HasArchivalParams returns a boolean if a field has been set.

### GetTieringParams

`func (o *ExternalTarget) GetTieringParams() TieringExternalTargetParams`

GetTieringParams returns the TieringParams field if non-nil, zero value otherwise.

### GetTieringParamsOk

`func (o *ExternalTarget) GetTieringParamsOk() (*TieringExternalTargetParams, bool)`

GetTieringParamsOk returns a tuple with the TieringParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTieringParams

`func (o *ExternalTarget) SetTieringParams(v TieringExternalTargetParams)`

SetTieringParams sets TieringParams field to given value.

### HasTieringParams

`func (o *ExternalTarget) HasTieringParams() bool`

HasTieringParams returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


