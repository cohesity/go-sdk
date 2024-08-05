# ObjectProtectionInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClusterId** | Pointer to **NullableInt64** | Specifies the cluster id where this object belongs to. | [optional] 
**ClusterIncarnationId** | Pointer to **NullableInt64** | Specifies the cluster incarnation id where this object belongs to. | [optional] 
**IsDeleted** | Pointer to **NullableBool** | Specifies whether the object is deleted. Deleted objects can&#39;t be protected but can be recovered or unprotected. | [optional] 
**LastRunStatus** | Pointer to **NullableString** | Specifies the status of the object&#39;s last protection run. | [optional] 
**ObjectBackupConfiguration** | Pointer to [**[]ProtectionSummary**](ProtectionSummary.md) | Specifies a list of object protections. | [optional] 
**ObjectId** | Pointer to **NullableInt64** | Specifies the object id. | [optional] 
**ProtectionGroups** | Pointer to [**[]ObjectProtectionGroupSummary**](ObjectProtectionGroupSummary.md) | Specifies a list of protection groups protecting this object. | [optional] 
**RegionId** | Pointer to **NullableString** | Specifies the region id where this object belongs to. | [optional] 
**SourceId** | Pointer to **NullableInt64** | Specifies the source id. | [optional] 
**TenantIds** | Pointer to **[]string** | List of Tenants the object belongs to. | [optional] 
**ViewId** | Pointer to **NullableInt64** | Specifies the view id for the object. | [optional] 

## Methods

### NewObjectProtectionInfo

`func NewObjectProtectionInfo() *ObjectProtectionInfo`

NewObjectProtectionInfo instantiates a new ObjectProtectionInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewObjectProtectionInfoWithDefaults

`func NewObjectProtectionInfoWithDefaults() *ObjectProtectionInfo`

NewObjectProtectionInfoWithDefaults instantiates a new ObjectProtectionInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClusterId

`func (o *ObjectProtectionInfo) GetClusterId() int64`

GetClusterId returns the ClusterId field if non-nil, zero value otherwise.

### GetClusterIdOk

`func (o *ObjectProtectionInfo) GetClusterIdOk() (*int64, bool)`

GetClusterIdOk returns a tuple with the ClusterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterId

`func (o *ObjectProtectionInfo) SetClusterId(v int64)`

SetClusterId sets ClusterId field to given value.

### HasClusterId

`func (o *ObjectProtectionInfo) HasClusterId() bool`

HasClusterId returns a boolean if a field has been set.

### SetClusterIdNil

`func (o *ObjectProtectionInfo) SetClusterIdNil(b bool)`

 SetClusterIdNil sets the value for ClusterId to be an explicit nil

### UnsetClusterId
`func (o *ObjectProtectionInfo) UnsetClusterId()`

UnsetClusterId ensures that no value is present for ClusterId, not even an explicit nil
### GetClusterIncarnationId

`func (o *ObjectProtectionInfo) GetClusterIncarnationId() int64`

GetClusterIncarnationId returns the ClusterIncarnationId field if non-nil, zero value otherwise.

### GetClusterIncarnationIdOk

`func (o *ObjectProtectionInfo) GetClusterIncarnationIdOk() (*int64, bool)`

GetClusterIncarnationIdOk returns a tuple with the ClusterIncarnationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterIncarnationId

`func (o *ObjectProtectionInfo) SetClusterIncarnationId(v int64)`

SetClusterIncarnationId sets ClusterIncarnationId field to given value.

### HasClusterIncarnationId

`func (o *ObjectProtectionInfo) HasClusterIncarnationId() bool`

HasClusterIncarnationId returns a boolean if a field has been set.

### SetClusterIncarnationIdNil

`func (o *ObjectProtectionInfo) SetClusterIncarnationIdNil(b bool)`

 SetClusterIncarnationIdNil sets the value for ClusterIncarnationId to be an explicit nil

### UnsetClusterIncarnationId
`func (o *ObjectProtectionInfo) UnsetClusterIncarnationId()`

UnsetClusterIncarnationId ensures that no value is present for ClusterIncarnationId, not even an explicit nil
### GetIsDeleted

`func (o *ObjectProtectionInfo) GetIsDeleted() bool`

GetIsDeleted returns the IsDeleted field if non-nil, zero value otherwise.

### GetIsDeletedOk

`func (o *ObjectProtectionInfo) GetIsDeletedOk() (*bool, bool)`

GetIsDeletedOk returns a tuple with the IsDeleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDeleted

`func (o *ObjectProtectionInfo) SetIsDeleted(v bool)`

SetIsDeleted sets IsDeleted field to given value.

### HasIsDeleted

`func (o *ObjectProtectionInfo) HasIsDeleted() bool`

HasIsDeleted returns a boolean if a field has been set.

### SetIsDeletedNil

`func (o *ObjectProtectionInfo) SetIsDeletedNil(b bool)`

 SetIsDeletedNil sets the value for IsDeleted to be an explicit nil

### UnsetIsDeleted
`func (o *ObjectProtectionInfo) UnsetIsDeleted()`

UnsetIsDeleted ensures that no value is present for IsDeleted, not even an explicit nil
### GetLastRunStatus

`func (o *ObjectProtectionInfo) GetLastRunStatus() string`

GetLastRunStatus returns the LastRunStatus field if non-nil, zero value otherwise.

### GetLastRunStatusOk

`func (o *ObjectProtectionInfo) GetLastRunStatusOk() (*string, bool)`

GetLastRunStatusOk returns a tuple with the LastRunStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastRunStatus

`func (o *ObjectProtectionInfo) SetLastRunStatus(v string)`

SetLastRunStatus sets LastRunStatus field to given value.

### HasLastRunStatus

`func (o *ObjectProtectionInfo) HasLastRunStatus() bool`

HasLastRunStatus returns a boolean if a field has been set.

### SetLastRunStatusNil

`func (o *ObjectProtectionInfo) SetLastRunStatusNil(b bool)`

 SetLastRunStatusNil sets the value for LastRunStatus to be an explicit nil

### UnsetLastRunStatus
`func (o *ObjectProtectionInfo) UnsetLastRunStatus()`

UnsetLastRunStatus ensures that no value is present for LastRunStatus, not even an explicit nil
### GetObjectBackupConfiguration

`func (o *ObjectProtectionInfo) GetObjectBackupConfiguration() []ProtectionSummary`

GetObjectBackupConfiguration returns the ObjectBackupConfiguration field if non-nil, zero value otherwise.

### GetObjectBackupConfigurationOk

`func (o *ObjectProtectionInfo) GetObjectBackupConfigurationOk() (*[]ProtectionSummary, bool)`

GetObjectBackupConfigurationOk returns a tuple with the ObjectBackupConfiguration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectBackupConfiguration

`func (o *ObjectProtectionInfo) SetObjectBackupConfiguration(v []ProtectionSummary)`

SetObjectBackupConfiguration sets ObjectBackupConfiguration field to given value.

### HasObjectBackupConfiguration

`func (o *ObjectProtectionInfo) HasObjectBackupConfiguration() bool`

HasObjectBackupConfiguration returns a boolean if a field has been set.

### SetObjectBackupConfigurationNil

`func (o *ObjectProtectionInfo) SetObjectBackupConfigurationNil(b bool)`

 SetObjectBackupConfigurationNil sets the value for ObjectBackupConfiguration to be an explicit nil

### UnsetObjectBackupConfiguration
`func (o *ObjectProtectionInfo) UnsetObjectBackupConfiguration()`

UnsetObjectBackupConfiguration ensures that no value is present for ObjectBackupConfiguration, not even an explicit nil
### GetObjectId

`func (o *ObjectProtectionInfo) GetObjectId() int64`

GetObjectId returns the ObjectId field if non-nil, zero value otherwise.

### GetObjectIdOk

`func (o *ObjectProtectionInfo) GetObjectIdOk() (*int64, bool)`

GetObjectIdOk returns a tuple with the ObjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectId

`func (o *ObjectProtectionInfo) SetObjectId(v int64)`

SetObjectId sets ObjectId field to given value.

### HasObjectId

`func (o *ObjectProtectionInfo) HasObjectId() bool`

HasObjectId returns a boolean if a field has been set.

### SetObjectIdNil

`func (o *ObjectProtectionInfo) SetObjectIdNil(b bool)`

 SetObjectIdNil sets the value for ObjectId to be an explicit nil

### UnsetObjectId
`func (o *ObjectProtectionInfo) UnsetObjectId()`

UnsetObjectId ensures that no value is present for ObjectId, not even an explicit nil
### GetProtectionGroups

`func (o *ObjectProtectionInfo) GetProtectionGroups() []ObjectProtectionGroupSummary`

GetProtectionGroups returns the ProtectionGroups field if non-nil, zero value otherwise.

### GetProtectionGroupsOk

`func (o *ObjectProtectionInfo) GetProtectionGroupsOk() (*[]ObjectProtectionGroupSummary, bool)`

GetProtectionGroupsOk returns a tuple with the ProtectionGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtectionGroups

`func (o *ObjectProtectionInfo) SetProtectionGroups(v []ObjectProtectionGroupSummary)`

SetProtectionGroups sets ProtectionGroups field to given value.

### HasProtectionGroups

`func (o *ObjectProtectionInfo) HasProtectionGroups() bool`

HasProtectionGroups returns a boolean if a field has been set.

### SetProtectionGroupsNil

`func (o *ObjectProtectionInfo) SetProtectionGroupsNil(b bool)`

 SetProtectionGroupsNil sets the value for ProtectionGroups to be an explicit nil

### UnsetProtectionGroups
`func (o *ObjectProtectionInfo) UnsetProtectionGroups()`

UnsetProtectionGroups ensures that no value is present for ProtectionGroups, not even an explicit nil
### GetRegionId

`func (o *ObjectProtectionInfo) GetRegionId() string`

GetRegionId returns the RegionId field if non-nil, zero value otherwise.

### GetRegionIdOk

`func (o *ObjectProtectionInfo) GetRegionIdOk() (*string, bool)`

GetRegionIdOk returns a tuple with the RegionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegionId

`func (o *ObjectProtectionInfo) SetRegionId(v string)`

SetRegionId sets RegionId field to given value.

### HasRegionId

`func (o *ObjectProtectionInfo) HasRegionId() bool`

HasRegionId returns a boolean if a field has been set.

### SetRegionIdNil

`func (o *ObjectProtectionInfo) SetRegionIdNil(b bool)`

 SetRegionIdNil sets the value for RegionId to be an explicit nil

### UnsetRegionId
`func (o *ObjectProtectionInfo) UnsetRegionId()`

UnsetRegionId ensures that no value is present for RegionId, not even an explicit nil
### GetSourceId

`func (o *ObjectProtectionInfo) GetSourceId() int64`

GetSourceId returns the SourceId field if non-nil, zero value otherwise.

### GetSourceIdOk

`func (o *ObjectProtectionInfo) GetSourceIdOk() (*int64, bool)`

GetSourceIdOk returns a tuple with the SourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceId

`func (o *ObjectProtectionInfo) SetSourceId(v int64)`

SetSourceId sets SourceId field to given value.

### HasSourceId

`func (o *ObjectProtectionInfo) HasSourceId() bool`

HasSourceId returns a boolean if a field has been set.

### SetSourceIdNil

`func (o *ObjectProtectionInfo) SetSourceIdNil(b bool)`

 SetSourceIdNil sets the value for SourceId to be an explicit nil

### UnsetSourceId
`func (o *ObjectProtectionInfo) UnsetSourceId()`

UnsetSourceId ensures that no value is present for SourceId, not even an explicit nil
### GetTenantIds

`func (o *ObjectProtectionInfo) GetTenantIds() []string`

GetTenantIds returns the TenantIds field if non-nil, zero value otherwise.

### GetTenantIdsOk

`func (o *ObjectProtectionInfo) GetTenantIdsOk() (*[]string, bool)`

GetTenantIdsOk returns a tuple with the TenantIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantIds

`func (o *ObjectProtectionInfo) SetTenantIds(v []string)`

SetTenantIds sets TenantIds field to given value.

### HasTenantIds

`func (o *ObjectProtectionInfo) HasTenantIds() bool`

HasTenantIds returns a boolean if a field has been set.

### GetViewId

`func (o *ObjectProtectionInfo) GetViewId() int64`

GetViewId returns the ViewId field if non-nil, zero value otherwise.

### GetViewIdOk

`func (o *ObjectProtectionInfo) GetViewIdOk() (*int64, bool)`

GetViewIdOk returns a tuple with the ViewId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViewId

`func (o *ObjectProtectionInfo) SetViewId(v int64)`

SetViewId sets ViewId field to given value.

### HasViewId

`func (o *ObjectProtectionInfo) HasViewId() bool`

HasViewId returns a boolean if a field has been set.

### SetViewIdNil

`func (o *ObjectProtectionInfo) SetViewIdNil(b bool)`

 SetViewIdNil sets the value for ViewId to be an explicit nil

### UnsetViewId
`func (o *ObjectProtectionInfo) UnsetViewId()`

UnsetViewId ensures that no value is present for ViewId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


