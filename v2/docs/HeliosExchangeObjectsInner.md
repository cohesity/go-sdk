# HeliosExchangeObjectsInner

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
**DatabaseName** | Pointer to **NullableString** | Specifies the name of the Exchange database corresponding to the mailbox. | [optional] 
**Email** | Pointer to **NullableString** | Specifies the email corresponding to the mailbox. | [optional] 
**ObjectName** | Pointer to **NullableString** | Specifies the name of the Exchange mailbox. | [optional] 

## Methods

### NewHeliosExchangeObjectsInner

`func NewHeliosExchangeObjectsInner() *HeliosExchangeObjectsInner`

NewHeliosExchangeObjectsInner instantiates a new HeliosExchangeObjectsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHeliosExchangeObjectsInnerWithDefaults

`func NewHeliosExchangeObjectsInnerWithDefaults() *HeliosExchangeObjectsInner`

NewHeliosExchangeObjectsInnerWithDefaults instantiates a new HeliosExchangeObjectsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClusterIdentifier

`func (o *HeliosExchangeObjectsInner) GetClusterIdentifier() string`

GetClusterIdentifier returns the ClusterIdentifier field if non-nil, zero value otherwise.

### GetClusterIdentifierOk

`func (o *HeliosExchangeObjectsInner) GetClusterIdentifierOk() (*string, bool)`

GetClusterIdentifierOk returns a tuple with the ClusterIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterIdentifier

`func (o *HeliosExchangeObjectsInner) SetClusterIdentifier(v string)`

SetClusterIdentifier sets ClusterIdentifier field to given value.

### HasClusterIdentifier

`func (o *HeliosExchangeObjectsInner) HasClusterIdentifier() bool`

HasClusterIdentifier returns a boolean if a field has been set.

### SetClusterIdentifierNil

`func (o *HeliosExchangeObjectsInner) SetClusterIdentifierNil(b bool)`

 SetClusterIdentifierNil sets the value for ClusterIdentifier to be an explicit nil

### UnsetClusterIdentifier
`func (o *HeliosExchangeObjectsInner) UnsetClusterIdentifier()`

UnsetClusterIdentifier ensures that no value is present for ClusterIdentifier, not even an explicit nil
### GetRegionId

`func (o *HeliosExchangeObjectsInner) GetRegionId() string`

GetRegionId returns the RegionId field if non-nil, zero value otherwise.

### GetRegionIdOk

`func (o *HeliosExchangeObjectsInner) GetRegionIdOk() (*string, bool)`

GetRegionIdOk returns a tuple with the RegionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegionId

`func (o *HeliosExchangeObjectsInner) SetRegionId(v string)`

SetRegionId sets RegionId field to given value.

### HasRegionId

`func (o *HeliosExchangeObjectsInner) HasRegionId() bool`

HasRegionId returns a boolean if a field has been set.

### SetRegionIdNil

`func (o *HeliosExchangeObjectsInner) SetRegionIdNil(b bool)`

 SetRegionIdNil sets the value for RegionId to be an explicit nil

### UnsetRegionId
`func (o *HeliosExchangeObjectsInner) UnsetRegionId()`

UnsetRegionId ensures that no value is present for RegionId, not even an explicit nil
### GetName

`func (o *HeliosExchangeObjectsInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *HeliosExchangeObjectsInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *HeliosExchangeObjectsInner) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *HeliosExchangeObjectsInner) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *HeliosExchangeObjectsInner) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *HeliosExchangeObjectsInner) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetPath

`func (o *HeliosExchangeObjectsInner) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *HeliosExchangeObjectsInner) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *HeliosExchangeObjectsInner) SetPath(v string)`

SetPath sets Path field to given value.

### HasPath

`func (o *HeliosExchangeObjectsInner) HasPath() bool`

HasPath returns a boolean if a field has been set.

### SetPathNil

`func (o *HeliosExchangeObjectsInner) SetPathNil(b bool)`

 SetPathNil sets the value for Path to be an explicit nil

### UnsetPath
`func (o *HeliosExchangeObjectsInner) UnsetPath()`

UnsetPath ensures that no value is present for Path, not even an explicit nil
### GetPolicyId

`func (o *HeliosExchangeObjectsInner) GetPolicyId() string`

GetPolicyId returns the PolicyId field if non-nil, zero value otherwise.

### GetPolicyIdOk

`func (o *HeliosExchangeObjectsInner) GetPolicyIdOk() (*string, bool)`

GetPolicyIdOk returns a tuple with the PolicyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyId

`func (o *HeliosExchangeObjectsInner) SetPolicyId(v string)`

SetPolicyId sets PolicyId field to given value.

### HasPolicyId

`func (o *HeliosExchangeObjectsInner) HasPolicyId() bool`

HasPolicyId returns a boolean if a field has been set.

### SetPolicyIdNil

`func (o *HeliosExchangeObjectsInner) SetPolicyIdNil(b bool)`

 SetPolicyIdNil sets the value for PolicyId to be an explicit nil

### UnsetPolicyId
`func (o *HeliosExchangeObjectsInner) UnsetPolicyId()`

UnsetPolicyId ensures that no value is present for PolicyId, not even an explicit nil
### GetPolicyName

`func (o *HeliosExchangeObjectsInner) GetPolicyName() string`

GetPolicyName returns the PolicyName field if non-nil, zero value otherwise.

### GetPolicyNameOk

`func (o *HeliosExchangeObjectsInner) GetPolicyNameOk() (*string, bool)`

GetPolicyNameOk returns a tuple with the PolicyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyName

`func (o *HeliosExchangeObjectsInner) SetPolicyName(v string)`

SetPolicyName sets PolicyName field to given value.

### HasPolicyName

`func (o *HeliosExchangeObjectsInner) HasPolicyName() bool`

HasPolicyName returns a boolean if a field has been set.

### SetPolicyNameNil

`func (o *HeliosExchangeObjectsInner) SetPolicyNameNil(b bool)`

 SetPolicyNameNil sets the value for PolicyName to be an explicit nil

### UnsetPolicyName
`func (o *HeliosExchangeObjectsInner) UnsetPolicyName()`

UnsetPolicyName ensures that no value is present for PolicyName, not even an explicit nil
### GetProtectionGroupId

`func (o *HeliosExchangeObjectsInner) GetProtectionGroupId() string`

GetProtectionGroupId returns the ProtectionGroupId field if non-nil, zero value otherwise.

### GetProtectionGroupIdOk

`func (o *HeliosExchangeObjectsInner) GetProtectionGroupIdOk() (*string, bool)`

GetProtectionGroupIdOk returns a tuple with the ProtectionGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtectionGroupId

`func (o *HeliosExchangeObjectsInner) SetProtectionGroupId(v string)`

SetProtectionGroupId sets ProtectionGroupId field to given value.

### HasProtectionGroupId

`func (o *HeliosExchangeObjectsInner) HasProtectionGroupId() bool`

HasProtectionGroupId returns a boolean if a field has been set.

### SetProtectionGroupIdNil

`func (o *HeliosExchangeObjectsInner) SetProtectionGroupIdNil(b bool)`

 SetProtectionGroupIdNil sets the value for ProtectionGroupId to be an explicit nil

### UnsetProtectionGroupId
`func (o *HeliosExchangeObjectsInner) UnsetProtectionGroupId()`

UnsetProtectionGroupId ensures that no value is present for ProtectionGroupId, not even an explicit nil
### GetProtectionGroupName

`func (o *HeliosExchangeObjectsInner) GetProtectionGroupName() string`

GetProtectionGroupName returns the ProtectionGroupName field if non-nil, zero value otherwise.

### GetProtectionGroupNameOk

`func (o *HeliosExchangeObjectsInner) GetProtectionGroupNameOk() (*string, bool)`

GetProtectionGroupNameOk returns a tuple with the ProtectionGroupName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtectionGroupName

`func (o *HeliosExchangeObjectsInner) SetProtectionGroupName(v string)`

SetProtectionGroupName sets ProtectionGroupName field to given value.

### HasProtectionGroupName

`func (o *HeliosExchangeObjectsInner) HasProtectionGroupName() bool`

HasProtectionGroupName returns a boolean if a field has been set.

### SetProtectionGroupNameNil

`func (o *HeliosExchangeObjectsInner) SetProtectionGroupNameNil(b bool)`

 SetProtectionGroupNameNil sets the value for ProtectionGroupName to be an explicit nil

### UnsetProtectionGroupName
`func (o *HeliosExchangeObjectsInner) UnsetProtectionGroupName()`

UnsetProtectionGroupName ensures that no value is present for ProtectionGroupName, not even an explicit nil
### GetSourceInfo

`func (o *HeliosExchangeObjectsInner) GetSourceInfo() map[string]interface{}`

GetSourceInfo returns the SourceInfo field if non-nil, zero value otherwise.

### GetSourceInfoOk

`func (o *HeliosExchangeObjectsInner) GetSourceInfoOk() (*map[string]interface{}, bool)`

GetSourceInfoOk returns a tuple with the SourceInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceInfo

`func (o *HeliosExchangeObjectsInner) SetSourceInfo(v map[string]interface{})`

SetSourceInfo sets SourceInfo field to given value.

### HasSourceInfo

`func (o *HeliosExchangeObjectsInner) HasSourceInfo() bool`

HasSourceInfo returns a boolean if a field has been set.

### GetStorageDomainId

`func (o *HeliosExchangeObjectsInner) GetStorageDomainId() int64`

GetStorageDomainId returns the StorageDomainId field if non-nil, zero value otherwise.

### GetStorageDomainIdOk

`func (o *HeliosExchangeObjectsInner) GetStorageDomainIdOk() (*int64, bool)`

GetStorageDomainIdOk returns a tuple with the StorageDomainId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageDomainId

`func (o *HeliosExchangeObjectsInner) SetStorageDomainId(v int64)`

SetStorageDomainId sets StorageDomainId field to given value.

### HasStorageDomainId

`func (o *HeliosExchangeObjectsInner) HasStorageDomainId() bool`

HasStorageDomainId returns a boolean if a field has been set.

### SetStorageDomainIdNil

`func (o *HeliosExchangeObjectsInner) SetStorageDomainIdNil(b bool)`

 SetStorageDomainIdNil sets the value for StorageDomainId to be an explicit nil

### UnsetStorageDomainId
`func (o *HeliosExchangeObjectsInner) UnsetStorageDomainId()`

UnsetStorageDomainId ensures that no value is present for StorageDomainId, not even an explicit nil
### GetSnapshotTags

`func (o *HeliosExchangeObjectsInner) GetSnapshotTags() []SnapshotTagInfo`

GetSnapshotTags returns the SnapshotTags field if non-nil, zero value otherwise.

### GetSnapshotTagsOk

`func (o *HeliosExchangeObjectsInner) GetSnapshotTagsOk() (*[]SnapshotTagInfo, bool)`

GetSnapshotTagsOk returns a tuple with the SnapshotTags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnapshotTags

`func (o *HeliosExchangeObjectsInner) SetSnapshotTags(v []SnapshotTagInfo)`

SetSnapshotTags sets SnapshotTags field to given value.

### HasSnapshotTags

`func (o *HeliosExchangeObjectsInner) HasSnapshotTags() bool`

HasSnapshotTags returns a boolean if a field has been set.

### SetSnapshotTagsNil

`func (o *HeliosExchangeObjectsInner) SetSnapshotTagsNil(b bool)`

 SetSnapshotTagsNil sets the value for SnapshotTags to be an explicit nil

### UnsetSnapshotTags
`func (o *HeliosExchangeObjectsInner) UnsetSnapshotTags()`

UnsetSnapshotTags ensures that no value is present for SnapshotTags, not even an explicit nil
### GetTags

`func (o *HeliosExchangeObjectsInner) GetTags() []TagInfo`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *HeliosExchangeObjectsInner) GetTagsOk() (*[]TagInfo, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *HeliosExchangeObjectsInner) SetTags(v []TagInfo)`

SetTags sets Tags field to given value.

### HasTags

`func (o *HeliosExchangeObjectsInner) HasTags() bool`

HasTags returns a boolean if a field has been set.

### SetTagsNil

`func (o *HeliosExchangeObjectsInner) SetTagsNil(b bool)`

 SetTagsNil sets the value for Tags to be an explicit nil

### UnsetTags
`func (o *HeliosExchangeObjectsInner) UnsetTags()`

UnsetTags ensures that no value is present for Tags, not even an explicit nil
### GetDatabaseName

`func (o *HeliosExchangeObjectsInner) GetDatabaseName() string`

GetDatabaseName returns the DatabaseName field if non-nil, zero value otherwise.

### GetDatabaseNameOk

`func (o *HeliosExchangeObjectsInner) GetDatabaseNameOk() (*string, bool)`

GetDatabaseNameOk returns a tuple with the DatabaseName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatabaseName

`func (o *HeliosExchangeObjectsInner) SetDatabaseName(v string)`

SetDatabaseName sets DatabaseName field to given value.

### HasDatabaseName

`func (o *HeliosExchangeObjectsInner) HasDatabaseName() bool`

HasDatabaseName returns a boolean if a field has been set.

### SetDatabaseNameNil

`func (o *HeliosExchangeObjectsInner) SetDatabaseNameNil(b bool)`

 SetDatabaseNameNil sets the value for DatabaseName to be an explicit nil

### UnsetDatabaseName
`func (o *HeliosExchangeObjectsInner) UnsetDatabaseName()`

UnsetDatabaseName ensures that no value is present for DatabaseName, not even an explicit nil
### GetEmail

`func (o *HeliosExchangeObjectsInner) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *HeliosExchangeObjectsInner) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *HeliosExchangeObjectsInner) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *HeliosExchangeObjectsInner) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### SetEmailNil

`func (o *HeliosExchangeObjectsInner) SetEmailNil(b bool)`

 SetEmailNil sets the value for Email to be an explicit nil

### UnsetEmail
`func (o *HeliosExchangeObjectsInner) UnsetEmail()`

UnsetEmail ensures that no value is present for Email, not even an explicit nil
### GetObjectName

`func (o *HeliosExchangeObjectsInner) GetObjectName() string`

GetObjectName returns the ObjectName field if non-nil, zero value otherwise.

### GetObjectNameOk

`func (o *HeliosExchangeObjectsInner) GetObjectNameOk() (*string, bool)`

GetObjectNameOk returns a tuple with the ObjectName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectName

`func (o *HeliosExchangeObjectsInner) SetObjectName(v string)`

SetObjectName sets ObjectName field to given value.

### HasObjectName

`func (o *HeliosExchangeObjectsInner) HasObjectName() bool`

HasObjectName returns a boolean if a field has been set.

### SetObjectNameNil

`func (o *HeliosExchangeObjectsInner) SetObjectNameNil(b bool)`

 SetObjectNameNil sets the value for ObjectName to be an explicit nil

### UnsetObjectName
`func (o *HeliosExchangeObjectsInner) UnsetObjectName()`

UnsetObjectName ensures that no value is present for ObjectName, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


