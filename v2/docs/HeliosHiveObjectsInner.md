# HeliosHiveObjectsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClusterIdentifier** | Pointer to **NullableString** | List of Clusters Identifiers to filter from. The format is clusterId:clusterIncarnationId. | [optional] 
**RegionId** | Pointer to **NullableString** | Specifies the region id of the cluster. Only valid for DMaaS clusters. | [optional] 
**Name** | Pointer to **NullableString** | Specifies the name of the object. | [optional] 
**Path** | Pointer to **NullableString** | Specifies the path of the object. | [optional] 
**PolicyId** | Pointer to **NullableString** | Specifies the protection policy id for this file. | [optional] 
**PolicyName** | Pointer to **NullableString** | Specifies the protection policy name for this file. | [optional] 
**ProtectionGroupId** | Pointer to **NullableString** | \&quot;Specifies the protection group id which contains this object.\&quot; | [optional] 
**ProtectionGroupName** | Pointer to **NullableString** | \&quot;Specifies the protection group name which contains this object.\&quot; | [optional] 
**SourceInfo** | Pointer to **map[string]interface{}** | Specifies the Source Object information. | [optional] 
**StorageDomainId** | Pointer to **NullableInt64** | \&quot;Specifies the Storage Domain id where the backup data of Object is present.\&quot; | [optional] 
**SnapshotTags** | Pointer to [**[]SnapshotTagInfo**](SnapshotTagInfo.md) | Specifies snapshot tags applied to the object. | [optional] 
**Tags** | Pointer to [**[]TagInfo**](TagInfo.md) | Specifies tag applied to the object. | [optional] 
**Id** | Pointer to **NullableString** | Specifies the id of the indexed object. | [optional] 
**Type** | Pointer to **NullableString** | Specifies the Hive Object Type. | [optional] 

## Methods

### NewHeliosHiveObjectsInner

`func NewHeliosHiveObjectsInner() *HeliosHiveObjectsInner`

NewHeliosHiveObjectsInner instantiates a new HeliosHiveObjectsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHeliosHiveObjectsInnerWithDefaults

`func NewHeliosHiveObjectsInnerWithDefaults() *HeliosHiveObjectsInner`

NewHeliosHiveObjectsInnerWithDefaults instantiates a new HeliosHiveObjectsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClusterIdentifier

`func (o *HeliosHiveObjectsInner) GetClusterIdentifier() string`

GetClusterIdentifier returns the ClusterIdentifier field if non-nil, zero value otherwise.

### GetClusterIdentifierOk

`func (o *HeliosHiveObjectsInner) GetClusterIdentifierOk() (*string, bool)`

GetClusterIdentifierOk returns a tuple with the ClusterIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterIdentifier

`func (o *HeliosHiveObjectsInner) SetClusterIdentifier(v string)`

SetClusterIdentifier sets ClusterIdentifier field to given value.

### HasClusterIdentifier

`func (o *HeliosHiveObjectsInner) HasClusterIdentifier() bool`

HasClusterIdentifier returns a boolean if a field has been set.

### SetClusterIdentifierNil

`func (o *HeliosHiveObjectsInner) SetClusterIdentifierNil(b bool)`

 SetClusterIdentifierNil sets the value for ClusterIdentifier to be an explicit nil

### UnsetClusterIdentifier
`func (o *HeliosHiveObjectsInner) UnsetClusterIdentifier()`

UnsetClusterIdentifier ensures that no value is present for ClusterIdentifier, not even an explicit nil
### GetRegionId

`func (o *HeliosHiveObjectsInner) GetRegionId() string`

GetRegionId returns the RegionId field if non-nil, zero value otherwise.

### GetRegionIdOk

`func (o *HeliosHiveObjectsInner) GetRegionIdOk() (*string, bool)`

GetRegionIdOk returns a tuple with the RegionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegionId

`func (o *HeliosHiveObjectsInner) SetRegionId(v string)`

SetRegionId sets RegionId field to given value.

### HasRegionId

`func (o *HeliosHiveObjectsInner) HasRegionId() bool`

HasRegionId returns a boolean if a field has been set.

### SetRegionIdNil

`func (o *HeliosHiveObjectsInner) SetRegionIdNil(b bool)`

 SetRegionIdNil sets the value for RegionId to be an explicit nil

### UnsetRegionId
`func (o *HeliosHiveObjectsInner) UnsetRegionId()`

UnsetRegionId ensures that no value is present for RegionId, not even an explicit nil
### GetName

`func (o *HeliosHiveObjectsInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *HeliosHiveObjectsInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *HeliosHiveObjectsInner) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *HeliosHiveObjectsInner) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *HeliosHiveObjectsInner) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *HeliosHiveObjectsInner) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetPath

`func (o *HeliosHiveObjectsInner) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *HeliosHiveObjectsInner) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *HeliosHiveObjectsInner) SetPath(v string)`

SetPath sets Path field to given value.

### HasPath

`func (o *HeliosHiveObjectsInner) HasPath() bool`

HasPath returns a boolean if a field has been set.

### SetPathNil

`func (o *HeliosHiveObjectsInner) SetPathNil(b bool)`

 SetPathNil sets the value for Path to be an explicit nil

### UnsetPath
`func (o *HeliosHiveObjectsInner) UnsetPath()`

UnsetPath ensures that no value is present for Path, not even an explicit nil
### GetPolicyId

`func (o *HeliosHiveObjectsInner) GetPolicyId() string`

GetPolicyId returns the PolicyId field if non-nil, zero value otherwise.

### GetPolicyIdOk

`func (o *HeliosHiveObjectsInner) GetPolicyIdOk() (*string, bool)`

GetPolicyIdOk returns a tuple with the PolicyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyId

`func (o *HeliosHiveObjectsInner) SetPolicyId(v string)`

SetPolicyId sets PolicyId field to given value.

### HasPolicyId

`func (o *HeliosHiveObjectsInner) HasPolicyId() bool`

HasPolicyId returns a boolean if a field has been set.

### SetPolicyIdNil

`func (o *HeliosHiveObjectsInner) SetPolicyIdNil(b bool)`

 SetPolicyIdNil sets the value for PolicyId to be an explicit nil

### UnsetPolicyId
`func (o *HeliosHiveObjectsInner) UnsetPolicyId()`

UnsetPolicyId ensures that no value is present for PolicyId, not even an explicit nil
### GetPolicyName

`func (o *HeliosHiveObjectsInner) GetPolicyName() string`

GetPolicyName returns the PolicyName field if non-nil, zero value otherwise.

### GetPolicyNameOk

`func (o *HeliosHiveObjectsInner) GetPolicyNameOk() (*string, bool)`

GetPolicyNameOk returns a tuple with the PolicyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyName

`func (o *HeliosHiveObjectsInner) SetPolicyName(v string)`

SetPolicyName sets PolicyName field to given value.

### HasPolicyName

`func (o *HeliosHiveObjectsInner) HasPolicyName() bool`

HasPolicyName returns a boolean if a field has been set.

### SetPolicyNameNil

`func (o *HeliosHiveObjectsInner) SetPolicyNameNil(b bool)`

 SetPolicyNameNil sets the value for PolicyName to be an explicit nil

### UnsetPolicyName
`func (o *HeliosHiveObjectsInner) UnsetPolicyName()`

UnsetPolicyName ensures that no value is present for PolicyName, not even an explicit nil
### GetProtectionGroupId

`func (o *HeliosHiveObjectsInner) GetProtectionGroupId() string`

GetProtectionGroupId returns the ProtectionGroupId field if non-nil, zero value otherwise.

### GetProtectionGroupIdOk

`func (o *HeliosHiveObjectsInner) GetProtectionGroupIdOk() (*string, bool)`

GetProtectionGroupIdOk returns a tuple with the ProtectionGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtectionGroupId

`func (o *HeliosHiveObjectsInner) SetProtectionGroupId(v string)`

SetProtectionGroupId sets ProtectionGroupId field to given value.

### HasProtectionGroupId

`func (o *HeliosHiveObjectsInner) HasProtectionGroupId() bool`

HasProtectionGroupId returns a boolean if a field has been set.

### SetProtectionGroupIdNil

`func (o *HeliosHiveObjectsInner) SetProtectionGroupIdNil(b bool)`

 SetProtectionGroupIdNil sets the value for ProtectionGroupId to be an explicit nil

### UnsetProtectionGroupId
`func (o *HeliosHiveObjectsInner) UnsetProtectionGroupId()`

UnsetProtectionGroupId ensures that no value is present for ProtectionGroupId, not even an explicit nil
### GetProtectionGroupName

`func (o *HeliosHiveObjectsInner) GetProtectionGroupName() string`

GetProtectionGroupName returns the ProtectionGroupName field if non-nil, zero value otherwise.

### GetProtectionGroupNameOk

`func (o *HeliosHiveObjectsInner) GetProtectionGroupNameOk() (*string, bool)`

GetProtectionGroupNameOk returns a tuple with the ProtectionGroupName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtectionGroupName

`func (o *HeliosHiveObjectsInner) SetProtectionGroupName(v string)`

SetProtectionGroupName sets ProtectionGroupName field to given value.

### HasProtectionGroupName

`func (o *HeliosHiveObjectsInner) HasProtectionGroupName() bool`

HasProtectionGroupName returns a boolean if a field has been set.

### SetProtectionGroupNameNil

`func (o *HeliosHiveObjectsInner) SetProtectionGroupNameNil(b bool)`

 SetProtectionGroupNameNil sets the value for ProtectionGroupName to be an explicit nil

### UnsetProtectionGroupName
`func (o *HeliosHiveObjectsInner) UnsetProtectionGroupName()`

UnsetProtectionGroupName ensures that no value is present for ProtectionGroupName, not even an explicit nil
### GetSourceInfo

`func (o *HeliosHiveObjectsInner) GetSourceInfo() map[string]interface{}`

GetSourceInfo returns the SourceInfo field if non-nil, zero value otherwise.

### GetSourceInfoOk

`func (o *HeliosHiveObjectsInner) GetSourceInfoOk() (*map[string]interface{}, bool)`

GetSourceInfoOk returns a tuple with the SourceInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceInfo

`func (o *HeliosHiveObjectsInner) SetSourceInfo(v map[string]interface{})`

SetSourceInfo sets SourceInfo field to given value.

### HasSourceInfo

`func (o *HeliosHiveObjectsInner) HasSourceInfo() bool`

HasSourceInfo returns a boolean if a field has been set.

### GetStorageDomainId

`func (o *HeliosHiveObjectsInner) GetStorageDomainId() int64`

GetStorageDomainId returns the StorageDomainId field if non-nil, zero value otherwise.

### GetStorageDomainIdOk

`func (o *HeliosHiveObjectsInner) GetStorageDomainIdOk() (*int64, bool)`

GetStorageDomainIdOk returns a tuple with the StorageDomainId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageDomainId

`func (o *HeliosHiveObjectsInner) SetStorageDomainId(v int64)`

SetStorageDomainId sets StorageDomainId field to given value.

### HasStorageDomainId

`func (o *HeliosHiveObjectsInner) HasStorageDomainId() bool`

HasStorageDomainId returns a boolean if a field has been set.

### SetStorageDomainIdNil

`func (o *HeliosHiveObjectsInner) SetStorageDomainIdNil(b bool)`

 SetStorageDomainIdNil sets the value for StorageDomainId to be an explicit nil

### UnsetStorageDomainId
`func (o *HeliosHiveObjectsInner) UnsetStorageDomainId()`

UnsetStorageDomainId ensures that no value is present for StorageDomainId, not even an explicit nil
### GetSnapshotTags

`func (o *HeliosHiveObjectsInner) GetSnapshotTags() []SnapshotTagInfo`

GetSnapshotTags returns the SnapshotTags field if non-nil, zero value otherwise.

### GetSnapshotTagsOk

`func (o *HeliosHiveObjectsInner) GetSnapshotTagsOk() (*[]SnapshotTagInfo, bool)`

GetSnapshotTagsOk returns a tuple with the SnapshotTags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnapshotTags

`func (o *HeliosHiveObjectsInner) SetSnapshotTags(v []SnapshotTagInfo)`

SetSnapshotTags sets SnapshotTags field to given value.

### HasSnapshotTags

`func (o *HeliosHiveObjectsInner) HasSnapshotTags() bool`

HasSnapshotTags returns a boolean if a field has been set.

### SetSnapshotTagsNil

`func (o *HeliosHiveObjectsInner) SetSnapshotTagsNil(b bool)`

 SetSnapshotTagsNil sets the value for SnapshotTags to be an explicit nil

### UnsetSnapshotTags
`func (o *HeliosHiveObjectsInner) UnsetSnapshotTags()`

UnsetSnapshotTags ensures that no value is present for SnapshotTags, not even an explicit nil
### GetTags

`func (o *HeliosHiveObjectsInner) GetTags() []TagInfo`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *HeliosHiveObjectsInner) GetTagsOk() (*[]TagInfo, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *HeliosHiveObjectsInner) SetTags(v []TagInfo)`

SetTags sets Tags field to given value.

### HasTags

`func (o *HeliosHiveObjectsInner) HasTags() bool`

HasTags returns a boolean if a field has been set.

### SetTagsNil

`func (o *HeliosHiveObjectsInner) SetTagsNil(b bool)`

 SetTagsNil sets the value for Tags to be an explicit nil

### UnsetTags
`func (o *HeliosHiveObjectsInner) UnsetTags()`

UnsetTags ensures that no value is present for Tags, not even an explicit nil
### GetId

`func (o *HeliosHiveObjectsInner) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *HeliosHiveObjectsInner) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *HeliosHiveObjectsInner) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *HeliosHiveObjectsInner) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *HeliosHiveObjectsInner) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *HeliosHiveObjectsInner) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetType

`func (o *HeliosHiveObjectsInner) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *HeliosHiveObjectsInner) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *HeliosHiveObjectsInner) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *HeliosHiveObjectsInner) HasType() bool`

HasType returns a boolean if a field has been set.

### SetTypeNil

`func (o *HeliosHiveObjectsInner) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *HeliosHiveObjectsInner) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


