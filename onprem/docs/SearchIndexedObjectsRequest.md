# SearchIndexedObjectsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProtectionGroupIds** | Pointer to **[]string** | Specifies a list of Protection Group ids to filter the indexed objects. If specified, the objects indexed by specified Protection Group ids will be returned. | [optional] 
**StorageDomainIds** | Pointer to **[]int64** | Specifies the Storage Domain ids to filter indexed objects for which Protection Groups are writing data to Cohesity Views on the specified Storage Domains. | [optional] 
**TenantId** | Pointer to **NullableString** | TenantId contains id of the tenant for which objects are to be returned. | [optional] 
**IncludeTenants** | Pointer to **NullableBool** | If true, the response will include objects which belongs to all tenants which the current user has permission to see. Default value is false. | [optional] [default to false]
**Tags** | Pointer to **[]string** | \&quot;This field is deprecated. Please use mightHaveTagIds.\&quot; | [optional] 
**SnapshotTags** | Pointer to **[]string** | \&quot;This field is deprecated. Please use mightHaveSnapshotTagIds.\&quot; | [optional] 
**MustHaveTagIds** | Pointer to **[]string** | Specifies tags which must be all present in the document. | [optional] 
**MightHaveTagIds** | Pointer to **[]string** | Specifies list of tags, one or more of which might be present in the document. These are OR&#39;ed together and the resulting criteria AND&#39;ed with the rest of the query. | [optional] 
**MustHaveSnapshotTagIds** | Pointer to **[]string** | Specifies snapshot tags which must be all present in the document. | [optional] 
**MightHaveSnapshotTagIds** | Pointer to **[]string** | Specifies list of snapshot tags, one or more of which might be present in the document. These are OR&#39;ed together and the resulting criteria AND&#39;ed with the rest of the query. | [optional] 
**PaginationCookie** | Pointer to **NullableString** | Specifies the pagination cookie with which subsequent parts of the response can be fetched. | [optional] 
**Count** | Pointer to **NullableInt32** | Specifies the number of indexed objects to be fetched for the specified pagination cookie. | [optional] 
**ObjectType** | **string** | Specifies the object type to be searched for. | 
**EmailParams** | Pointer to [**SearchEmailRequestParams**](SearchEmailRequestParams.md) |  | [optional] 
**FileParams** | Pointer to [**SearchFileRequestParams**](SearchFileRequestParams.md) |  | [optional] 
**CassandraParams** | Pointer to [**CassandraOnPremSearchParams**](CassandraOnPremSearchParams.md) |  | [optional] 
**CouchbaseParams** | Pointer to [**CouchBaseOnPremSearchParams**](CouchBaseOnPremSearchParams.md) |  | [optional] 
**HbaseParams** | Pointer to [**HbaseOnPremSearchParams**](HbaseOnPremSearchParams.md) |  | [optional] 
**HiveParams** | Pointer to [**HiveOnPremSearchParams**](HiveOnPremSearchParams.md) |  | [optional] 
**MongodbParams** | Pointer to [**MongoDbOnPremSearchParams**](MongoDbOnPremSearchParams.md) |  | [optional] 
**HdfsParams** | Pointer to [**HDFSOnPremSearchParams**](HDFSOnPremSearchParams.md) |  | [optional] 
**ExchangeParams** | Pointer to [**SearchExchangeObjectsRequestParams**](SearchExchangeObjectsRequestParams.md) |  | [optional] 
**PublicFolderParams** | Pointer to [**SearchPublicFolderRequestParams**](SearchPublicFolderRequestParams.md) |  | [optional] 
**MsTeamsParams** | Pointer to [**SearchMsTeamsRequestParams**](SearchMsTeamsRequestParams.md) |  | [optional] 
**SharepointParams** | Pointer to [**SearchDocumentLibraryRequestParams**](SearchDocumentLibraryRequestParams.md) |  | [optional] 

## Methods

### NewSearchIndexedObjectsRequest

`func NewSearchIndexedObjectsRequest(objectType string, ) *SearchIndexedObjectsRequest`

NewSearchIndexedObjectsRequest instantiates a new SearchIndexedObjectsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSearchIndexedObjectsRequestWithDefaults

`func NewSearchIndexedObjectsRequestWithDefaults() *SearchIndexedObjectsRequest`

NewSearchIndexedObjectsRequestWithDefaults instantiates a new SearchIndexedObjectsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProtectionGroupIds

`func (o *SearchIndexedObjectsRequest) GetProtectionGroupIds() []string`

GetProtectionGroupIds returns the ProtectionGroupIds field if non-nil, zero value otherwise.

### GetProtectionGroupIdsOk

`func (o *SearchIndexedObjectsRequest) GetProtectionGroupIdsOk() (*[]string, bool)`

GetProtectionGroupIdsOk returns a tuple with the ProtectionGroupIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtectionGroupIds

`func (o *SearchIndexedObjectsRequest) SetProtectionGroupIds(v []string)`

SetProtectionGroupIds sets ProtectionGroupIds field to given value.

### HasProtectionGroupIds

`func (o *SearchIndexedObjectsRequest) HasProtectionGroupIds() bool`

HasProtectionGroupIds returns a boolean if a field has been set.

### SetProtectionGroupIdsNil

`func (o *SearchIndexedObjectsRequest) SetProtectionGroupIdsNil(b bool)`

 SetProtectionGroupIdsNil sets the value for ProtectionGroupIds to be an explicit nil

### UnsetProtectionGroupIds
`func (o *SearchIndexedObjectsRequest) UnsetProtectionGroupIds()`

UnsetProtectionGroupIds ensures that no value is present for ProtectionGroupIds, not even an explicit nil
### GetStorageDomainIds

`func (o *SearchIndexedObjectsRequest) GetStorageDomainIds() []int64`

GetStorageDomainIds returns the StorageDomainIds field if non-nil, zero value otherwise.

### GetStorageDomainIdsOk

`func (o *SearchIndexedObjectsRequest) GetStorageDomainIdsOk() (*[]int64, bool)`

GetStorageDomainIdsOk returns a tuple with the StorageDomainIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageDomainIds

`func (o *SearchIndexedObjectsRequest) SetStorageDomainIds(v []int64)`

SetStorageDomainIds sets StorageDomainIds field to given value.

### HasStorageDomainIds

`func (o *SearchIndexedObjectsRequest) HasStorageDomainIds() bool`

HasStorageDomainIds returns a boolean if a field has been set.

### SetStorageDomainIdsNil

`func (o *SearchIndexedObjectsRequest) SetStorageDomainIdsNil(b bool)`

 SetStorageDomainIdsNil sets the value for StorageDomainIds to be an explicit nil

### UnsetStorageDomainIds
`func (o *SearchIndexedObjectsRequest) UnsetStorageDomainIds()`

UnsetStorageDomainIds ensures that no value is present for StorageDomainIds, not even an explicit nil
### GetTenantId

`func (o *SearchIndexedObjectsRequest) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *SearchIndexedObjectsRequest) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *SearchIndexedObjectsRequest) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.

### HasTenantId

`func (o *SearchIndexedObjectsRequest) HasTenantId() bool`

HasTenantId returns a boolean if a field has been set.

### SetTenantIdNil

`func (o *SearchIndexedObjectsRequest) SetTenantIdNil(b bool)`

 SetTenantIdNil sets the value for TenantId to be an explicit nil

### UnsetTenantId
`func (o *SearchIndexedObjectsRequest) UnsetTenantId()`

UnsetTenantId ensures that no value is present for TenantId, not even an explicit nil
### GetIncludeTenants

`func (o *SearchIndexedObjectsRequest) GetIncludeTenants() bool`

GetIncludeTenants returns the IncludeTenants field if non-nil, zero value otherwise.

### GetIncludeTenantsOk

`func (o *SearchIndexedObjectsRequest) GetIncludeTenantsOk() (*bool, bool)`

GetIncludeTenantsOk returns a tuple with the IncludeTenants field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeTenants

`func (o *SearchIndexedObjectsRequest) SetIncludeTenants(v bool)`

SetIncludeTenants sets IncludeTenants field to given value.

### HasIncludeTenants

`func (o *SearchIndexedObjectsRequest) HasIncludeTenants() bool`

HasIncludeTenants returns a boolean if a field has been set.

### SetIncludeTenantsNil

`func (o *SearchIndexedObjectsRequest) SetIncludeTenantsNil(b bool)`

 SetIncludeTenantsNil sets the value for IncludeTenants to be an explicit nil

### UnsetIncludeTenants
`func (o *SearchIndexedObjectsRequest) UnsetIncludeTenants()`

UnsetIncludeTenants ensures that no value is present for IncludeTenants, not even an explicit nil
### GetTags

`func (o *SearchIndexedObjectsRequest) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *SearchIndexedObjectsRequest) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *SearchIndexedObjectsRequest) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *SearchIndexedObjectsRequest) HasTags() bool`

HasTags returns a boolean if a field has been set.

### SetTagsNil

`func (o *SearchIndexedObjectsRequest) SetTagsNil(b bool)`

 SetTagsNil sets the value for Tags to be an explicit nil

### UnsetTags
`func (o *SearchIndexedObjectsRequest) UnsetTags()`

UnsetTags ensures that no value is present for Tags, not even an explicit nil
### GetSnapshotTags

`func (o *SearchIndexedObjectsRequest) GetSnapshotTags() []string`

GetSnapshotTags returns the SnapshotTags field if non-nil, zero value otherwise.

### GetSnapshotTagsOk

`func (o *SearchIndexedObjectsRequest) GetSnapshotTagsOk() (*[]string, bool)`

GetSnapshotTagsOk returns a tuple with the SnapshotTags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnapshotTags

`func (o *SearchIndexedObjectsRequest) SetSnapshotTags(v []string)`

SetSnapshotTags sets SnapshotTags field to given value.

### HasSnapshotTags

`func (o *SearchIndexedObjectsRequest) HasSnapshotTags() bool`

HasSnapshotTags returns a boolean if a field has been set.

### GetMustHaveTagIds

`func (o *SearchIndexedObjectsRequest) GetMustHaveTagIds() []string`

GetMustHaveTagIds returns the MustHaveTagIds field if non-nil, zero value otherwise.

### GetMustHaveTagIdsOk

`func (o *SearchIndexedObjectsRequest) GetMustHaveTagIdsOk() (*[]string, bool)`

GetMustHaveTagIdsOk returns a tuple with the MustHaveTagIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMustHaveTagIds

`func (o *SearchIndexedObjectsRequest) SetMustHaveTagIds(v []string)`

SetMustHaveTagIds sets MustHaveTagIds field to given value.

### HasMustHaveTagIds

`func (o *SearchIndexedObjectsRequest) HasMustHaveTagIds() bool`

HasMustHaveTagIds returns a boolean if a field has been set.

### SetMustHaveTagIdsNil

`func (o *SearchIndexedObjectsRequest) SetMustHaveTagIdsNil(b bool)`

 SetMustHaveTagIdsNil sets the value for MustHaveTagIds to be an explicit nil

### UnsetMustHaveTagIds
`func (o *SearchIndexedObjectsRequest) UnsetMustHaveTagIds()`

UnsetMustHaveTagIds ensures that no value is present for MustHaveTagIds, not even an explicit nil
### GetMightHaveTagIds

`func (o *SearchIndexedObjectsRequest) GetMightHaveTagIds() []string`

GetMightHaveTagIds returns the MightHaveTagIds field if non-nil, zero value otherwise.

### GetMightHaveTagIdsOk

`func (o *SearchIndexedObjectsRequest) GetMightHaveTagIdsOk() (*[]string, bool)`

GetMightHaveTagIdsOk returns a tuple with the MightHaveTagIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMightHaveTagIds

`func (o *SearchIndexedObjectsRequest) SetMightHaveTagIds(v []string)`

SetMightHaveTagIds sets MightHaveTagIds field to given value.

### HasMightHaveTagIds

`func (o *SearchIndexedObjectsRequest) HasMightHaveTagIds() bool`

HasMightHaveTagIds returns a boolean if a field has been set.

### SetMightHaveTagIdsNil

`func (o *SearchIndexedObjectsRequest) SetMightHaveTagIdsNil(b bool)`

 SetMightHaveTagIdsNil sets the value for MightHaveTagIds to be an explicit nil

### UnsetMightHaveTagIds
`func (o *SearchIndexedObjectsRequest) UnsetMightHaveTagIds()`

UnsetMightHaveTagIds ensures that no value is present for MightHaveTagIds, not even an explicit nil
### GetMustHaveSnapshotTagIds

`func (o *SearchIndexedObjectsRequest) GetMustHaveSnapshotTagIds() []string`

GetMustHaveSnapshotTagIds returns the MustHaveSnapshotTagIds field if non-nil, zero value otherwise.

### GetMustHaveSnapshotTagIdsOk

`func (o *SearchIndexedObjectsRequest) GetMustHaveSnapshotTagIdsOk() (*[]string, bool)`

GetMustHaveSnapshotTagIdsOk returns a tuple with the MustHaveSnapshotTagIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMustHaveSnapshotTagIds

`func (o *SearchIndexedObjectsRequest) SetMustHaveSnapshotTagIds(v []string)`

SetMustHaveSnapshotTagIds sets MustHaveSnapshotTagIds field to given value.

### HasMustHaveSnapshotTagIds

`func (o *SearchIndexedObjectsRequest) HasMustHaveSnapshotTagIds() bool`

HasMustHaveSnapshotTagIds returns a boolean if a field has been set.

### SetMustHaveSnapshotTagIdsNil

`func (o *SearchIndexedObjectsRequest) SetMustHaveSnapshotTagIdsNil(b bool)`

 SetMustHaveSnapshotTagIdsNil sets the value for MustHaveSnapshotTagIds to be an explicit nil

### UnsetMustHaveSnapshotTagIds
`func (o *SearchIndexedObjectsRequest) UnsetMustHaveSnapshotTagIds()`

UnsetMustHaveSnapshotTagIds ensures that no value is present for MustHaveSnapshotTagIds, not even an explicit nil
### GetMightHaveSnapshotTagIds

`func (o *SearchIndexedObjectsRequest) GetMightHaveSnapshotTagIds() []string`

GetMightHaveSnapshotTagIds returns the MightHaveSnapshotTagIds field if non-nil, zero value otherwise.

### GetMightHaveSnapshotTagIdsOk

`func (o *SearchIndexedObjectsRequest) GetMightHaveSnapshotTagIdsOk() (*[]string, bool)`

GetMightHaveSnapshotTagIdsOk returns a tuple with the MightHaveSnapshotTagIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMightHaveSnapshotTagIds

`func (o *SearchIndexedObjectsRequest) SetMightHaveSnapshotTagIds(v []string)`

SetMightHaveSnapshotTagIds sets MightHaveSnapshotTagIds field to given value.

### HasMightHaveSnapshotTagIds

`func (o *SearchIndexedObjectsRequest) HasMightHaveSnapshotTagIds() bool`

HasMightHaveSnapshotTagIds returns a boolean if a field has been set.

### SetMightHaveSnapshotTagIdsNil

`func (o *SearchIndexedObjectsRequest) SetMightHaveSnapshotTagIdsNil(b bool)`

 SetMightHaveSnapshotTagIdsNil sets the value for MightHaveSnapshotTagIds to be an explicit nil

### UnsetMightHaveSnapshotTagIds
`func (o *SearchIndexedObjectsRequest) UnsetMightHaveSnapshotTagIds()`

UnsetMightHaveSnapshotTagIds ensures that no value is present for MightHaveSnapshotTagIds, not even an explicit nil
### GetPaginationCookie

`func (o *SearchIndexedObjectsRequest) GetPaginationCookie() string`

GetPaginationCookie returns the PaginationCookie field if non-nil, zero value otherwise.

### GetPaginationCookieOk

`func (o *SearchIndexedObjectsRequest) GetPaginationCookieOk() (*string, bool)`

GetPaginationCookieOk returns a tuple with the PaginationCookie field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaginationCookie

`func (o *SearchIndexedObjectsRequest) SetPaginationCookie(v string)`

SetPaginationCookie sets PaginationCookie field to given value.

### HasPaginationCookie

`func (o *SearchIndexedObjectsRequest) HasPaginationCookie() bool`

HasPaginationCookie returns a boolean if a field has been set.

### SetPaginationCookieNil

`func (o *SearchIndexedObjectsRequest) SetPaginationCookieNil(b bool)`

 SetPaginationCookieNil sets the value for PaginationCookie to be an explicit nil

### UnsetPaginationCookie
`func (o *SearchIndexedObjectsRequest) UnsetPaginationCookie()`

UnsetPaginationCookie ensures that no value is present for PaginationCookie, not even an explicit nil
### GetCount

`func (o *SearchIndexedObjectsRequest) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *SearchIndexedObjectsRequest) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *SearchIndexedObjectsRequest) SetCount(v int32)`

SetCount sets Count field to given value.

### HasCount

`func (o *SearchIndexedObjectsRequest) HasCount() bool`

HasCount returns a boolean if a field has been set.

### SetCountNil

`func (o *SearchIndexedObjectsRequest) SetCountNil(b bool)`

 SetCountNil sets the value for Count to be an explicit nil

### UnsetCount
`func (o *SearchIndexedObjectsRequest) UnsetCount()`

UnsetCount ensures that no value is present for Count, not even an explicit nil
### GetObjectType

`func (o *SearchIndexedObjectsRequest) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *SearchIndexedObjectsRequest) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *SearchIndexedObjectsRequest) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetEmailParams

`func (o *SearchIndexedObjectsRequest) GetEmailParams() SearchEmailRequestParams`

GetEmailParams returns the EmailParams field if non-nil, zero value otherwise.

### GetEmailParamsOk

`func (o *SearchIndexedObjectsRequest) GetEmailParamsOk() (*SearchEmailRequestParams, bool)`

GetEmailParamsOk returns a tuple with the EmailParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailParams

`func (o *SearchIndexedObjectsRequest) SetEmailParams(v SearchEmailRequestParams)`

SetEmailParams sets EmailParams field to given value.

### HasEmailParams

`func (o *SearchIndexedObjectsRequest) HasEmailParams() bool`

HasEmailParams returns a boolean if a field has been set.

### GetFileParams

`func (o *SearchIndexedObjectsRequest) GetFileParams() SearchFileRequestParams`

GetFileParams returns the FileParams field if non-nil, zero value otherwise.

### GetFileParamsOk

`func (o *SearchIndexedObjectsRequest) GetFileParamsOk() (*SearchFileRequestParams, bool)`

GetFileParamsOk returns a tuple with the FileParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileParams

`func (o *SearchIndexedObjectsRequest) SetFileParams(v SearchFileRequestParams)`

SetFileParams sets FileParams field to given value.

### HasFileParams

`func (o *SearchIndexedObjectsRequest) HasFileParams() bool`

HasFileParams returns a boolean if a field has been set.

### GetCassandraParams

`func (o *SearchIndexedObjectsRequest) GetCassandraParams() CassandraOnPremSearchParams`

GetCassandraParams returns the CassandraParams field if non-nil, zero value otherwise.

### GetCassandraParamsOk

`func (o *SearchIndexedObjectsRequest) GetCassandraParamsOk() (*CassandraOnPremSearchParams, bool)`

GetCassandraParamsOk returns a tuple with the CassandraParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCassandraParams

`func (o *SearchIndexedObjectsRequest) SetCassandraParams(v CassandraOnPremSearchParams)`

SetCassandraParams sets CassandraParams field to given value.

### HasCassandraParams

`func (o *SearchIndexedObjectsRequest) HasCassandraParams() bool`

HasCassandraParams returns a boolean if a field has been set.

### GetCouchbaseParams

`func (o *SearchIndexedObjectsRequest) GetCouchbaseParams() CouchBaseOnPremSearchParams`

GetCouchbaseParams returns the CouchbaseParams field if non-nil, zero value otherwise.

### GetCouchbaseParamsOk

`func (o *SearchIndexedObjectsRequest) GetCouchbaseParamsOk() (*CouchBaseOnPremSearchParams, bool)`

GetCouchbaseParamsOk returns a tuple with the CouchbaseParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCouchbaseParams

`func (o *SearchIndexedObjectsRequest) SetCouchbaseParams(v CouchBaseOnPremSearchParams)`

SetCouchbaseParams sets CouchbaseParams field to given value.

### HasCouchbaseParams

`func (o *SearchIndexedObjectsRequest) HasCouchbaseParams() bool`

HasCouchbaseParams returns a boolean if a field has been set.

### GetHbaseParams

`func (o *SearchIndexedObjectsRequest) GetHbaseParams() HbaseOnPremSearchParams`

GetHbaseParams returns the HbaseParams field if non-nil, zero value otherwise.

### GetHbaseParamsOk

`func (o *SearchIndexedObjectsRequest) GetHbaseParamsOk() (*HbaseOnPremSearchParams, bool)`

GetHbaseParamsOk returns a tuple with the HbaseParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHbaseParams

`func (o *SearchIndexedObjectsRequest) SetHbaseParams(v HbaseOnPremSearchParams)`

SetHbaseParams sets HbaseParams field to given value.

### HasHbaseParams

`func (o *SearchIndexedObjectsRequest) HasHbaseParams() bool`

HasHbaseParams returns a boolean if a field has been set.

### GetHiveParams

`func (o *SearchIndexedObjectsRequest) GetHiveParams() HiveOnPremSearchParams`

GetHiveParams returns the HiveParams field if non-nil, zero value otherwise.

### GetHiveParamsOk

`func (o *SearchIndexedObjectsRequest) GetHiveParamsOk() (*HiveOnPremSearchParams, bool)`

GetHiveParamsOk returns a tuple with the HiveParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHiveParams

`func (o *SearchIndexedObjectsRequest) SetHiveParams(v HiveOnPremSearchParams)`

SetHiveParams sets HiveParams field to given value.

### HasHiveParams

`func (o *SearchIndexedObjectsRequest) HasHiveParams() bool`

HasHiveParams returns a boolean if a field has been set.

### GetMongodbParams

`func (o *SearchIndexedObjectsRequest) GetMongodbParams() MongoDbOnPremSearchParams`

GetMongodbParams returns the MongodbParams field if non-nil, zero value otherwise.

### GetMongodbParamsOk

`func (o *SearchIndexedObjectsRequest) GetMongodbParamsOk() (*MongoDbOnPremSearchParams, bool)`

GetMongodbParamsOk returns a tuple with the MongodbParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMongodbParams

`func (o *SearchIndexedObjectsRequest) SetMongodbParams(v MongoDbOnPremSearchParams)`

SetMongodbParams sets MongodbParams field to given value.

### HasMongodbParams

`func (o *SearchIndexedObjectsRequest) HasMongodbParams() bool`

HasMongodbParams returns a boolean if a field has been set.

### GetHdfsParams

`func (o *SearchIndexedObjectsRequest) GetHdfsParams() HDFSOnPremSearchParams`

GetHdfsParams returns the HdfsParams field if non-nil, zero value otherwise.

### GetHdfsParamsOk

`func (o *SearchIndexedObjectsRequest) GetHdfsParamsOk() (*HDFSOnPremSearchParams, bool)`

GetHdfsParamsOk returns a tuple with the HdfsParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHdfsParams

`func (o *SearchIndexedObjectsRequest) SetHdfsParams(v HDFSOnPremSearchParams)`

SetHdfsParams sets HdfsParams field to given value.

### HasHdfsParams

`func (o *SearchIndexedObjectsRequest) HasHdfsParams() bool`

HasHdfsParams returns a boolean if a field has been set.

### GetExchangeParams

`func (o *SearchIndexedObjectsRequest) GetExchangeParams() SearchExchangeObjectsRequestParams`

GetExchangeParams returns the ExchangeParams field if non-nil, zero value otherwise.

### GetExchangeParamsOk

`func (o *SearchIndexedObjectsRequest) GetExchangeParamsOk() (*SearchExchangeObjectsRequestParams, bool)`

GetExchangeParamsOk returns a tuple with the ExchangeParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExchangeParams

`func (o *SearchIndexedObjectsRequest) SetExchangeParams(v SearchExchangeObjectsRequestParams)`

SetExchangeParams sets ExchangeParams field to given value.

### HasExchangeParams

`func (o *SearchIndexedObjectsRequest) HasExchangeParams() bool`

HasExchangeParams returns a boolean if a field has been set.

### GetPublicFolderParams

`func (o *SearchIndexedObjectsRequest) GetPublicFolderParams() SearchPublicFolderRequestParams`

GetPublicFolderParams returns the PublicFolderParams field if non-nil, zero value otherwise.

### GetPublicFolderParamsOk

`func (o *SearchIndexedObjectsRequest) GetPublicFolderParamsOk() (*SearchPublicFolderRequestParams, bool)`

GetPublicFolderParamsOk returns a tuple with the PublicFolderParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicFolderParams

`func (o *SearchIndexedObjectsRequest) SetPublicFolderParams(v SearchPublicFolderRequestParams)`

SetPublicFolderParams sets PublicFolderParams field to given value.

### HasPublicFolderParams

`func (o *SearchIndexedObjectsRequest) HasPublicFolderParams() bool`

HasPublicFolderParams returns a boolean if a field has been set.

### GetMsTeamsParams

`func (o *SearchIndexedObjectsRequest) GetMsTeamsParams() SearchMsTeamsRequestParams`

GetMsTeamsParams returns the MsTeamsParams field if non-nil, zero value otherwise.

### GetMsTeamsParamsOk

`func (o *SearchIndexedObjectsRequest) GetMsTeamsParamsOk() (*SearchMsTeamsRequestParams, bool)`

GetMsTeamsParamsOk returns a tuple with the MsTeamsParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsTeamsParams

`func (o *SearchIndexedObjectsRequest) SetMsTeamsParams(v SearchMsTeamsRequestParams)`

SetMsTeamsParams sets MsTeamsParams field to given value.

### HasMsTeamsParams

`func (o *SearchIndexedObjectsRequest) HasMsTeamsParams() bool`

HasMsTeamsParams returns a boolean if a field has been set.

### GetSharepointParams

`func (o *SearchIndexedObjectsRequest) GetSharepointParams() SearchDocumentLibraryRequestParams`

GetSharepointParams returns the SharepointParams field if non-nil, zero value otherwise.

### GetSharepointParamsOk

`func (o *SearchIndexedObjectsRequest) GetSharepointParamsOk() (*SearchDocumentLibraryRequestParams, bool)`

GetSharepointParamsOk returns a tuple with the SharepointParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharepointParams

`func (o *SearchIndexedObjectsRequest) SetSharepointParams(v SearchDocumentLibraryRequestParams)`

SetSharepointParams sets SharepointParams field to given value.

### HasSharepointParams

`func (o *SearchIndexedObjectsRequest) HasSharepointParams() bool`

HasSharepointParams returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


