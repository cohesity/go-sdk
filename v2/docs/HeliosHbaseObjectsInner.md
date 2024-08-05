# HeliosHbaseObjectsInner

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
**Type** | Pointer to **NullableString** | Specifies the Hbase Object Type. | [optional] 

## Methods

### NewHeliosHbaseObjectsInner

`func NewHeliosHbaseObjectsInner() *HeliosHbaseObjectsInner`

NewHeliosHbaseObjectsInner instantiates a new HeliosHbaseObjectsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHeliosHbaseObjectsInnerWithDefaults

`func NewHeliosHbaseObjectsInnerWithDefaults() *HeliosHbaseObjectsInner`

NewHeliosHbaseObjectsInnerWithDefaults instantiates a new HeliosHbaseObjectsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClusterIdentifier

`func (o *HeliosHbaseObjectsInner) GetClusterIdentifier() string`

GetClusterIdentifier returns the ClusterIdentifier field if non-nil, zero value otherwise.

### GetClusterIdentifierOk

`func (o *HeliosHbaseObjectsInner) GetClusterIdentifierOk() (*string, bool)`

GetClusterIdentifierOk returns a tuple with the ClusterIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterIdentifier

`func (o *HeliosHbaseObjectsInner) SetClusterIdentifier(v string)`

SetClusterIdentifier sets ClusterIdentifier field to given value.

### HasClusterIdentifier

`func (o *HeliosHbaseObjectsInner) HasClusterIdentifier() bool`

HasClusterIdentifier returns a boolean if a field has been set.

### SetClusterIdentifierNil

`func (o *HeliosHbaseObjectsInner) SetClusterIdentifierNil(b bool)`

 SetClusterIdentifierNil sets the value for ClusterIdentifier to be an explicit nil

### UnsetClusterIdentifier
`func (o *HeliosHbaseObjectsInner) UnsetClusterIdentifier()`

UnsetClusterIdentifier ensures that no value is present for ClusterIdentifier, not even an explicit nil
### GetRegionId

`func (o *HeliosHbaseObjectsInner) GetRegionId() string`

GetRegionId returns the RegionId field if non-nil, zero value otherwise.

### GetRegionIdOk

`func (o *HeliosHbaseObjectsInner) GetRegionIdOk() (*string, bool)`

GetRegionIdOk returns a tuple with the RegionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegionId

`func (o *HeliosHbaseObjectsInner) SetRegionId(v string)`

SetRegionId sets RegionId field to given value.

### HasRegionId

`func (o *HeliosHbaseObjectsInner) HasRegionId() bool`

HasRegionId returns a boolean if a field has been set.

### SetRegionIdNil

`func (o *HeliosHbaseObjectsInner) SetRegionIdNil(b bool)`

 SetRegionIdNil sets the value for RegionId to be an explicit nil

### UnsetRegionId
`func (o *HeliosHbaseObjectsInner) UnsetRegionId()`

UnsetRegionId ensures that no value is present for RegionId, not even an explicit nil
### GetName

`func (o *HeliosHbaseObjectsInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *HeliosHbaseObjectsInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *HeliosHbaseObjectsInner) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *HeliosHbaseObjectsInner) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *HeliosHbaseObjectsInner) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *HeliosHbaseObjectsInner) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetPath

`func (o *HeliosHbaseObjectsInner) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *HeliosHbaseObjectsInner) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *HeliosHbaseObjectsInner) SetPath(v string)`

SetPath sets Path field to given value.

### HasPath

`func (o *HeliosHbaseObjectsInner) HasPath() bool`

HasPath returns a boolean if a field has been set.

### SetPathNil

`func (o *HeliosHbaseObjectsInner) SetPathNil(b bool)`

 SetPathNil sets the value for Path to be an explicit nil

### UnsetPath
`func (o *HeliosHbaseObjectsInner) UnsetPath()`

UnsetPath ensures that no value is present for Path, not even an explicit nil
### GetPolicyId

`func (o *HeliosHbaseObjectsInner) GetPolicyId() string`

GetPolicyId returns the PolicyId field if non-nil, zero value otherwise.

### GetPolicyIdOk

`func (o *HeliosHbaseObjectsInner) GetPolicyIdOk() (*string, bool)`

GetPolicyIdOk returns a tuple with the PolicyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyId

`func (o *HeliosHbaseObjectsInner) SetPolicyId(v string)`

SetPolicyId sets PolicyId field to given value.

### HasPolicyId

`func (o *HeliosHbaseObjectsInner) HasPolicyId() bool`

HasPolicyId returns a boolean if a field has been set.

### SetPolicyIdNil

`func (o *HeliosHbaseObjectsInner) SetPolicyIdNil(b bool)`

 SetPolicyIdNil sets the value for PolicyId to be an explicit nil

### UnsetPolicyId
`func (o *HeliosHbaseObjectsInner) UnsetPolicyId()`

UnsetPolicyId ensures that no value is present for PolicyId, not even an explicit nil
### GetPolicyName

`func (o *HeliosHbaseObjectsInner) GetPolicyName() string`

GetPolicyName returns the PolicyName field if non-nil, zero value otherwise.

### GetPolicyNameOk

`func (o *HeliosHbaseObjectsInner) GetPolicyNameOk() (*string, bool)`

GetPolicyNameOk returns a tuple with the PolicyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyName

`func (o *HeliosHbaseObjectsInner) SetPolicyName(v string)`

SetPolicyName sets PolicyName field to given value.

### HasPolicyName

`func (o *HeliosHbaseObjectsInner) HasPolicyName() bool`

HasPolicyName returns a boolean if a field has been set.

### SetPolicyNameNil

`func (o *HeliosHbaseObjectsInner) SetPolicyNameNil(b bool)`

 SetPolicyNameNil sets the value for PolicyName to be an explicit nil

### UnsetPolicyName
`func (o *HeliosHbaseObjectsInner) UnsetPolicyName()`

UnsetPolicyName ensures that no value is present for PolicyName, not even an explicit nil
### GetProtectionGroupId

`func (o *HeliosHbaseObjectsInner) GetProtectionGroupId() string`

GetProtectionGroupId returns the ProtectionGroupId field if non-nil, zero value otherwise.

### GetProtectionGroupIdOk

`func (o *HeliosHbaseObjectsInner) GetProtectionGroupIdOk() (*string, bool)`

GetProtectionGroupIdOk returns a tuple with the ProtectionGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtectionGroupId

`func (o *HeliosHbaseObjectsInner) SetProtectionGroupId(v string)`

SetProtectionGroupId sets ProtectionGroupId field to given value.

### HasProtectionGroupId

`func (o *HeliosHbaseObjectsInner) HasProtectionGroupId() bool`

HasProtectionGroupId returns a boolean if a field has been set.

### SetProtectionGroupIdNil

`func (o *HeliosHbaseObjectsInner) SetProtectionGroupIdNil(b bool)`

 SetProtectionGroupIdNil sets the value for ProtectionGroupId to be an explicit nil

### UnsetProtectionGroupId
`func (o *HeliosHbaseObjectsInner) UnsetProtectionGroupId()`

UnsetProtectionGroupId ensures that no value is present for ProtectionGroupId, not even an explicit nil
### GetProtectionGroupName

`func (o *HeliosHbaseObjectsInner) GetProtectionGroupName() string`

GetProtectionGroupName returns the ProtectionGroupName field if non-nil, zero value otherwise.

### GetProtectionGroupNameOk

`func (o *HeliosHbaseObjectsInner) GetProtectionGroupNameOk() (*string, bool)`

GetProtectionGroupNameOk returns a tuple with the ProtectionGroupName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtectionGroupName

`func (o *HeliosHbaseObjectsInner) SetProtectionGroupName(v string)`

SetProtectionGroupName sets ProtectionGroupName field to given value.

### HasProtectionGroupName

`func (o *HeliosHbaseObjectsInner) HasProtectionGroupName() bool`

HasProtectionGroupName returns a boolean if a field has been set.

### SetProtectionGroupNameNil

`func (o *HeliosHbaseObjectsInner) SetProtectionGroupNameNil(b bool)`

 SetProtectionGroupNameNil sets the value for ProtectionGroupName to be an explicit nil

### UnsetProtectionGroupName
`func (o *HeliosHbaseObjectsInner) UnsetProtectionGroupName()`

UnsetProtectionGroupName ensures that no value is present for ProtectionGroupName, not even an explicit nil
### GetSourceInfo

`func (o *HeliosHbaseObjectsInner) GetSourceInfo() map[string]interface{}`

GetSourceInfo returns the SourceInfo field if non-nil, zero value otherwise.

### GetSourceInfoOk

`func (o *HeliosHbaseObjectsInner) GetSourceInfoOk() (*map[string]interface{}, bool)`

GetSourceInfoOk returns a tuple with the SourceInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceInfo

`func (o *HeliosHbaseObjectsInner) SetSourceInfo(v map[string]interface{})`

SetSourceInfo sets SourceInfo field to given value.

### HasSourceInfo

`func (o *HeliosHbaseObjectsInner) HasSourceInfo() bool`

HasSourceInfo returns a boolean if a field has been set.

### GetStorageDomainId

`func (o *HeliosHbaseObjectsInner) GetStorageDomainId() int64`

GetStorageDomainId returns the StorageDomainId field if non-nil, zero value otherwise.

### GetStorageDomainIdOk

`func (o *HeliosHbaseObjectsInner) GetStorageDomainIdOk() (*int64, bool)`

GetStorageDomainIdOk returns a tuple with the StorageDomainId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageDomainId

`func (o *HeliosHbaseObjectsInner) SetStorageDomainId(v int64)`

SetStorageDomainId sets StorageDomainId field to given value.

### HasStorageDomainId

`func (o *HeliosHbaseObjectsInner) HasStorageDomainId() bool`

HasStorageDomainId returns a boolean if a field has been set.

### SetStorageDomainIdNil

`func (o *HeliosHbaseObjectsInner) SetStorageDomainIdNil(b bool)`

 SetStorageDomainIdNil sets the value for StorageDomainId to be an explicit nil

### UnsetStorageDomainId
`func (o *HeliosHbaseObjectsInner) UnsetStorageDomainId()`

UnsetStorageDomainId ensures that no value is present for StorageDomainId, not even an explicit nil
### GetSnapshotTags

`func (o *HeliosHbaseObjectsInner) GetSnapshotTags() []SnapshotTagInfo`

GetSnapshotTags returns the SnapshotTags field if non-nil, zero value otherwise.

### GetSnapshotTagsOk

`func (o *HeliosHbaseObjectsInner) GetSnapshotTagsOk() (*[]SnapshotTagInfo, bool)`

GetSnapshotTagsOk returns a tuple with the SnapshotTags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnapshotTags

`func (o *HeliosHbaseObjectsInner) SetSnapshotTags(v []SnapshotTagInfo)`

SetSnapshotTags sets SnapshotTags field to given value.

### HasSnapshotTags

`func (o *HeliosHbaseObjectsInner) HasSnapshotTags() bool`

HasSnapshotTags returns a boolean if a field has been set.

### SetSnapshotTagsNil

`func (o *HeliosHbaseObjectsInner) SetSnapshotTagsNil(b bool)`

 SetSnapshotTagsNil sets the value for SnapshotTags to be an explicit nil

### UnsetSnapshotTags
`func (o *HeliosHbaseObjectsInner) UnsetSnapshotTags()`

UnsetSnapshotTags ensures that no value is present for SnapshotTags, not even an explicit nil
### GetTags

`func (o *HeliosHbaseObjectsInner) GetTags() []TagInfo`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *HeliosHbaseObjectsInner) GetTagsOk() (*[]TagInfo, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *HeliosHbaseObjectsInner) SetTags(v []TagInfo)`

SetTags sets Tags field to given value.

### HasTags

`func (o *HeliosHbaseObjectsInner) HasTags() bool`

HasTags returns a boolean if a field has been set.

### SetTagsNil

`func (o *HeliosHbaseObjectsInner) SetTagsNil(b bool)`

 SetTagsNil sets the value for Tags to be an explicit nil

### UnsetTags
`func (o *HeliosHbaseObjectsInner) UnsetTags()`

UnsetTags ensures that no value is present for Tags, not even an explicit nil
### GetId

`func (o *HeliosHbaseObjectsInner) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *HeliosHbaseObjectsInner) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *HeliosHbaseObjectsInner) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *HeliosHbaseObjectsInner) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *HeliosHbaseObjectsInner) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *HeliosHbaseObjectsInner) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetType

`func (o *HeliosHbaseObjectsInner) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *HeliosHbaseObjectsInner) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *HeliosHbaseObjectsInner) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *HeliosHbaseObjectsInner) HasType() bool`

HasType returns a boolean if a field has been set.

### SetTypeNil

`func (o *HeliosHbaseObjectsInner) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *HeliosHbaseObjectsInner) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


