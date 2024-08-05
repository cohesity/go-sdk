# CommonSearchIndexedObjectsRequestParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Count** | Pointer to **NullableInt32** | Specifies the number of indexed objects to be fetched for the specified pagination cookie. | [optional] 
**IncludeTenants** | Pointer to **NullableBool** | If true, the response will include objects which belongs to all tenants which the current user has permission to see. Default value is false. | [optional] [default to false]
**MightHaveSnapshotTagIds** | Pointer to **[]string** | Specifies list of snapshot tags, one or more of which might be present in the document. These are OR&#39;ed together and the resulting criteria AND&#39;ed with the rest of the query. | [optional] 
**MightHaveTagIds** | Pointer to **[]string** | Specifies list of tags, one or more of which might be present in the document. These are OR&#39;ed together and the resulting criteria AND&#39;ed with the rest of the query. | [optional] 
**MustHaveSnapshotTagIds** | Pointer to **[]string** | Specifies snapshot tags which must be all present in the document. | [optional] 
**MustHaveTagIds** | Pointer to **[]string** | Specifies tags which must be all present in the document. | [optional] 
**ObjectType** | **string** | Specifies the object type to be searched for. | 
**PaginationCookie** | Pointer to **NullableString** | Specifies the pagination cookie with which subsequent parts of the response can be fetched. | [optional] 
**ProtectionGroupIds** | Pointer to **[]string** | Specifies a list of Protection Group ids to filter the indexed objects. If specified, the objects indexed by specified Protection Group ids will be returned. | [optional] 
**SnapshotTags** | Pointer to **[]string** | \&quot;This field is deprecated. Please use mightHaveSnapshotTagIds.\&quot; | [optional] 
**StorageDomainIds** | Pointer to **[]int64** | Specifies the Storage Domain ids to filter indexed objects for which Protection Groups are writing data to Cohesity Views on the specified Storage Domains. | [optional] 
**Tags** | Pointer to **[]string** | \&quot;This field is deprecated. Please use mightHaveTagIds.\&quot; | [optional] 
**TenantId** | Pointer to **NullableString** | TenantId contains id of the tenant for which objects are to be returned. | [optional] 
**UseCachedData** | Pointer to **NullableBool** | Specifies whether we can serve the GET request from the read replica cache. There is a lag of 15 seconds between the read replica and primary data source. | [optional] 

## Methods

### NewCommonSearchIndexedObjectsRequestParams

`func NewCommonSearchIndexedObjectsRequestParams(objectType string, ) *CommonSearchIndexedObjectsRequestParams`

NewCommonSearchIndexedObjectsRequestParams instantiates a new CommonSearchIndexedObjectsRequestParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCommonSearchIndexedObjectsRequestParamsWithDefaults

`func NewCommonSearchIndexedObjectsRequestParamsWithDefaults() *CommonSearchIndexedObjectsRequestParams`

NewCommonSearchIndexedObjectsRequestParamsWithDefaults instantiates a new CommonSearchIndexedObjectsRequestParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCount

`func (o *CommonSearchIndexedObjectsRequestParams) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *CommonSearchIndexedObjectsRequestParams) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *CommonSearchIndexedObjectsRequestParams) SetCount(v int32)`

SetCount sets Count field to given value.

### HasCount

`func (o *CommonSearchIndexedObjectsRequestParams) HasCount() bool`

HasCount returns a boolean if a field has been set.

### SetCountNil

`func (o *CommonSearchIndexedObjectsRequestParams) SetCountNil(b bool)`

 SetCountNil sets the value for Count to be an explicit nil

### UnsetCount
`func (o *CommonSearchIndexedObjectsRequestParams) UnsetCount()`

UnsetCount ensures that no value is present for Count, not even an explicit nil
### GetIncludeTenants

`func (o *CommonSearchIndexedObjectsRequestParams) GetIncludeTenants() bool`

GetIncludeTenants returns the IncludeTenants field if non-nil, zero value otherwise.

### GetIncludeTenantsOk

`func (o *CommonSearchIndexedObjectsRequestParams) GetIncludeTenantsOk() (*bool, bool)`

GetIncludeTenantsOk returns a tuple with the IncludeTenants field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeTenants

`func (o *CommonSearchIndexedObjectsRequestParams) SetIncludeTenants(v bool)`

SetIncludeTenants sets IncludeTenants field to given value.

### HasIncludeTenants

`func (o *CommonSearchIndexedObjectsRequestParams) HasIncludeTenants() bool`

HasIncludeTenants returns a boolean if a field has been set.

### SetIncludeTenantsNil

`func (o *CommonSearchIndexedObjectsRequestParams) SetIncludeTenantsNil(b bool)`

 SetIncludeTenantsNil sets the value for IncludeTenants to be an explicit nil

### UnsetIncludeTenants
`func (o *CommonSearchIndexedObjectsRequestParams) UnsetIncludeTenants()`

UnsetIncludeTenants ensures that no value is present for IncludeTenants, not even an explicit nil
### GetMightHaveSnapshotTagIds

`func (o *CommonSearchIndexedObjectsRequestParams) GetMightHaveSnapshotTagIds() []string`

GetMightHaveSnapshotTagIds returns the MightHaveSnapshotTagIds field if non-nil, zero value otherwise.

### GetMightHaveSnapshotTagIdsOk

`func (o *CommonSearchIndexedObjectsRequestParams) GetMightHaveSnapshotTagIdsOk() (*[]string, bool)`

GetMightHaveSnapshotTagIdsOk returns a tuple with the MightHaveSnapshotTagIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMightHaveSnapshotTagIds

`func (o *CommonSearchIndexedObjectsRequestParams) SetMightHaveSnapshotTagIds(v []string)`

SetMightHaveSnapshotTagIds sets MightHaveSnapshotTagIds field to given value.

### HasMightHaveSnapshotTagIds

`func (o *CommonSearchIndexedObjectsRequestParams) HasMightHaveSnapshotTagIds() bool`

HasMightHaveSnapshotTagIds returns a boolean if a field has been set.

### SetMightHaveSnapshotTagIdsNil

`func (o *CommonSearchIndexedObjectsRequestParams) SetMightHaveSnapshotTagIdsNil(b bool)`

 SetMightHaveSnapshotTagIdsNil sets the value for MightHaveSnapshotTagIds to be an explicit nil

### UnsetMightHaveSnapshotTagIds
`func (o *CommonSearchIndexedObjectsRequestParams) UnsetMightHaveSnapshotTagIds()`

UnsetMightHaveSnapshotTagIds ensures that no value is present for MightHaveSnapshotTagIds, not even an explicit nil
### GetMightHaveTagIds

`func (o *CommonSearchIndexedObjectsRequestParams) GetMightHaveTagIds() []string`

GetMightHaveTagIds returns the MightHaveTagIds field if non-nil, zero value otherwise.

### GetMightHaveTagIdsOk

`func (o *CommonSearchIndexedObjectsRequestParams) GetMightHaveTagIdsOk() (*[]string, bool)`

GetMightHaveTagIdsOk returns a tuple with the MightHaveTagIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMightHaveTagIds

`func (o *CommonSearchIndexedObjectsRequestParams) SetMightHaveTagIds(v []string)`

SetMightHaveTagIds sets MightHaveTagIds field to given value.

### HasMightHaveTagIds

`func (o *CommonSearchIndexedObjectsRequestParams) HasMightHaveTagIds() bool`

HasMightHaveTagIds returns a boolean if a field has been set.

### SetMightHaveTagIdsNil

`func (o *CommonSearchIndexedObjectsRequestParams) SetMightHaveTagIdsNil(b bool)`

 SetMightHaveTagIdsNil sets the value for MightHaveTagIds to be an explicit nil

### UnsetMightHaveTagIds
`func (o *CommonSearchIndexedObjectsRequestParams) UnsetMightHaveTagIds()`

UnsetMightHaveTagIds ensures that no value is present for MightHaveTagIds, not even an explicit nil
### GetMustHaveSnapshotTagIds

`func (o *CommonSearchIndexedObjectsRequestParams) GetMustHaveSnapshotTagIds() []string`

GetMustHaveSnapshotTagIds returns the MustHaveSnapshotTagIds field if non-nil, zero value otherwise.

### GetMustHaveSnapshotTagIdsOk

`func (o *CommonSearchIndexedObjectsRequestParams) GetMustHaveSnapshotTagIdsOk() (*[]string, bool)`

GetMustHaveSnapshotTagIdsOk returns a tuple with the MustHaveSnapshotTagIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMustHaveSnapshotTagIds

`func (o *CommonSearchIndexedObjectsRequestParams) SetMustHaveSnapshotTagIds(v []string)`

SetMustHaveSnapshotTagIds sets MustHaveSnapshotTagIds field to given value.

### HasMustHaveSnapshotTagIds

`func (o *CommonSearchIndexedObjectsRequestParams) HasMustHaveSnapshotTagIds() bool`

HasMustHaveSnapshotTagIds returns a boolean if a field has been set.

### SetMustHaveSnapshotTagIdsNil

`func (o *CommonSearchIndexedObjectsRequestParams) SetMustHaveSnapshotTagIdsNil(b bool)`

 SetMustHaveSnapshotTagIdsNil sets the value for MustHaveSnapshotTagIds to be an explicit nil

### UnsetMustHaveSnapshotTagIds
`func (o *CommonSearchIndexedObjectsRequestParams) UnsetMustHaveSnapshotTagIds()`

UnsetMustHaveSnapshotTagIds ensures that no value is present for MustHaveSnapshotTagIds, not even an explicit nil
### GetMustHaveTagIds

`func (o *CommonSearchIndexedObjectsRequestParams) GetMustHaveTagIds() []string`

GetMustHaveTagIds returns the MustHaveTagIds field if non-nil, zero value otherwise.

### GetMustHaveTagIdsOk

`func (o *CommonSearchIndexedObjectsRequestParams) GetMustHaveTagIdsOk() (*[]string, bool)`

GetMustHaveTagIdsOk returns a tuple with the MustHaveTagIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMustHaveTagIds

`func (o *CommonSearchIndexedObjectsRequestParams) SetMustHaveTagIds(v []string)`

SetMustHaveTagIds sets MustHaveTagIds field to given value.

### HasMustHaveTagIds

`func (o *CommonSearchIndexedObjectsRequestParams) HasMustHaveTagIds() bool`

HasMustHaveTagIds returns a boolean if a field has been set.

### SetMustHaveTagIdsNil

`func (o *CommonSearchIndexedObjectsRequestParams) SetMustHaveTagIdsNil(b bool)`

 SetMustHaveTagIdsNil sets the value for MustHaveTagIds to be an explicit nil

### UnsetMustHaveTagIds
`func (o *CommonSearchIndexedObjectsRequestParams) UnsetMustHaveTagIds()`

UnsetMustHaveTagIds ensures that no value is present for MustHaveTagIds, not even an explicit nil
### GetObjectType

`func (o *CommonSearchIndexedObjectsRequestParams) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *CommonSearchIndexedObjectsRequestParams) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *CommonSearchIndexedObjectsRequestParams) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetPaginationCookie

`func (o *CommonSearchIndexedObjectsRequestParams) GetPaginationCookie() string`

GetPaginationCookie returns the PaginationCookie field if non-nil, zero value otherwise.

### GetPaginationCookieOk

`func (o *CommonSearchIndexedObjectsRequestParams) GetPaginationCookieOk() (*string, bool)`

GetPaginationCookieOk returns a tuple with the PaginationCookie field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaginationCookie

`func (o *CommonSearchIndexedObjectsRequestParams) SetPaginationCookie(v string)`

SetPaginationCookie sets PaginationCookie field to given value.

### HasPaginationCookie

`func (o *CommonSearchIndexedObjectsRequestParams) HasPaginationCookie() bool`

HasPaginationCookie returns a boolean if a field has been set.

### SetPaginationCookieNil

`func (o *CommonSearchIndexedObjectsRequestParams) SetPaginationCookieNil(b bool)`

 SetPaginationCookieNil sets the value for PaginationCookie to be an explicit nil

### UnsetPaginationCookie
`func (o *CommonSearchIndexedObjectsRequestParams) UnsetPaginationCookie()`

UnsetPaginationCookie ensures that no value is present for PaginationCookie, not even an explicit nil
### GetProtectionGroupIds

`func (o *CommonSearchIndexedObjectsRequestParams) GetProtectionGroupIds() []string`

GetProtectionGroupIds returns the ProtectionGroupIds field if non-nil, zero value otherwise.

### GetProtectionGroupIdsOk

`func (o *CommonSearchIndexedObjectsRequestParams) GetProtectionGroupIdsOk() (*[]string, bool)`

GetProtectionGroupIdsOk returns a tuple with the ProtectionGroupIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtectionGroupIds

`func (o *CommonSearchIndexedObjectsRequestParams) SetProtectionGroupIds(v []string)`

SetProtectionGroupIds sets ProtectionGroupIds field to given value.

### HasProtectionGroupIds

`func (o *CommonSearchIndexedObjectsRequestParams) HasProtectionGroupIds() bool`

HasProtectionGroupIds returns a boolean if a field has been set.

### SetProtectionGroupIdsNil

`func (o *CommonSearchIndexedObjectsRequestParams) SetProtectionGroupIdsNil(b bool)`

 SetProtectionGroupIdsNil sets the value for ProtectionGroupIds to be an explicit nil

### UnsetProtectionGroupIds
`func (o *CommonSearchIndexedObjectsRequestParams) UnsetProtectionGroupIds()`

UnsetProtectionGroupIds ensures that no value is present for ProtectionGroupIds, not even an explicit nil
### GetSnapshotTags

`func (o *CommonSearchIndexedObjectsRequestParams) GetSnapshotTags() []string`

GetSnapshotTags returns the SnapshotTags field if non-nil, zero value otherwise.

### GetSnapshotTagsOk

`func (o *CommonSearchIndexedObjectsRequestParams) GetSnapshotTagsOk() (*[]string, bool)`

GetSnapshotTagsOk returns a tuple with the SnapshotTags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnapshotTags

`func (o *CommonSearchIndexedObjectsRequestParams) SetSnapshotTags(v []string)`

SetSnapshotTags sets SnapshotTags field to given value.

### HasSnapshotTags

`func (o *CommonSearchIndexedObjectsRequestParams) HasSnapshotTags() bool`

HasSnapshotTags returns a boolean if a field has been set.

### GetStorageDomainIds

`func (o *CommonSearchIndexedObjectsRequestParams) GetStorageDomainIds() []int64`

GetStorageDomainIds returns the StorageDomainIds field if non-nil, zero value otherwise.

### GetStorageDomainIdsOk

`func (o *CommonSearchIndexedObjectsRequestParams) GetStorageDomainIdsOk() (*[]int64, bool)`

GetStorageDomainIdsOk returns a tuple with the StorageDomainIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageDomainIds

`func (o *CommonSearchIndexedObjectsRequestParams) SetStorageDomainIds(v []int64)`

SetStorageDomainIds sets StorageDomainIds field to given value.

### HasStorageDomainIds

`func (o *CommonSearchIndexedObjectsRequestParams) HasStorageDomainIds() bool`

HasStorageDomainIds returns a boolean if a field has been set.

### SetStorageDomainIdsNil

`func (o *CommonSearchIndexedObjectsRequestParams) SetStorageDomainIdsNil(b bool)`

 SetStorageDomainIdsNil sets the value for StorageDomainIds to be an explicit nil

### UnsetStorageDomainIds
`func (o *CommonSearchIndexedObjectsRequestParams) UnsetStorageDomainIds()`

UnsetStorageDomainIds ensures that no value is present for StorageDomainIds, not even an explicit nil
### GetTags

`func (o *CommonSearchIndexedObjectsRequestParams) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *CommonSearchIndexedObjectsRequestParams) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *CommonSearchIndexedObjectsRequestParams) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *CommonSearchIndexedObjectsRequestParams) HasTags() bool`

HasTags returns a boolean if a field has been set.

### SetTagsNil

`func (o *CommonSearchIndexedObjectsRequestParams) SetTagsNil(b bool)`

 SetTagsNil sets the value for Tags to be an explicit nil

### UnsetTags
`func (o *CommonSearchIndexedObjectsRequestParams) UnsetTags()`

UnsetTags ensures that no value is present for Tags, not even an explicit nil
### GetTenantId

`func (o *CommonSearchIndexedObjectsRequestParams) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *CommonSearchIndexedObjectsRequestParams) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *CommonSearchIndexedObjectsRequestParams) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.

### HasTenantId

`func (o *CommonSearchIndexedObjectsRequestParams) HasTenantId() bool`

HasTenantId returns a boolean if a field has been set.

### SetTenantIdNil

`func (o *CommonSearchIndexedObjectsRequestParams) SetTenantIdNil(b bool)`

 SetTenantIdNil sets the value for TenantId to be an explicit nil

### UnsetTenantId
`func (o *CommonSearchIndexedObjectsRequestParams) UnsetTenantId()`

UnsetTenantId ensures that no value is present for TenantId, not even an explicit nil
### GetUseCachedData

`func (o *CommonSearchIndexedObjectsRequestParams) GetUseCachedData() bool`

GetUseCachedData returns the UseCachedData field if non-nil, zero value otherwise.

### GetUseCachedDataOk

`func (o *CommonSearchIndexedObjectsRequestParams) GetUseCachedDataOk() (*bool, bool)`

GetUseCachedDataOk returns a tuple with the UseCachedData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseCachedData

`func (o *CommonSearchIndexedObjectsRequestParams) SetUseCachedData(v bool)`

SetUseCachedData sets UseCachedData field to given value.

### HasUseCachedData

`func (o *CommonSearchIndexedObjectsRequestParams) HasUseCachedData() bool`

HasUseCachedData returns a boolean if a field has been set.

### SetUseCachedDataNil

`func (o *CommonSearchIndexedObjectsRequestParams) SetUseCachedDataNil(b bool)`

 SetUseCachedDataNil sets the value for UseCachedData to be an explicit nil

### UnsetUseCachedData
`func (o *CommonSearchIndexedObjectsRequestParams) UnsetUseCachedData()`

UnsetUseCachedData ensures that no value is present for UseCachedData, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


