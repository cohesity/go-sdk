# CassandraIndexedObject

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
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
**KeyspaceType** | Pointer to **NullableString** | Specifies type of Keyspace. | [optional] 
**Type** | Pointer to **NullableString** | Specifies the Cassandra Object type. | [optional] 

## Methods

### NewCassandraIndexedObject

`func NewCassandraIndexedObject() *CassandraIndexedObject`

NewCassandraIndexedObject instantiates a new CassandraIndexedObject object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCassandraIndexedObjectWithDefaults

`func NewCassandraIndexedObjectWithDefaults() *CassandraIndexedObject`

NewCassandraIndexedObjectWithDefaults instantiates a new CassandraIndexedObject object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CassandraIndexedObject) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CassandraIndexedObject) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CassandraIndexedObject) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CassandraIndexedObject) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *CassandraIndexedObject) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *CassandraIndexedObject) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetPath

`func (o *CassandraIndexedObject) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *CassandraIndexedObject) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *CassandraIndexedObject) SetPath(v string)`

SetPath sets Path field to given value.

### HasPath

`func (o *CassandraIndexedObject) HasPath() bool`

HasPath returns a boolean if a field has been set.

### SetPathNil

`func (o *CassandraIndexedObject) SetPathNil(b bool)`

 SetPathNil sets the value for Path to be an explicit nil

### UnsetPath
`func (o *CassandraIndexedObject) UnsetPath()`

UnsetPath ensures that no value is present for Path, not even an explicit nil
### GetPolicyId

`func (o *CassandraIndexedObject) GetPolicyId() string`

GetPolicyId returns the PolicyId field if non-nil, zero value otherwise.

### GetPolicyIdOk

`func (o *CassandraIndexedObject) GetPolicyIdOk() (*string, bool)`

GetPolicyIdOk returns a tuple with the PolicyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyId

`func (o *CassandraIndexedObject) SetPolicyId(v string)`

SetPolicyId sets PolicyId field to given value.

### HasPolicyId

`func (o *CassandraIndexedObject) HasPolicyId() bool`

HasPolicyId returns a boolean if a field has been set.

### SetPolicyIdNil

`func (o *CassandraIndexedObject) SetPolicyIdNil(b bool)`

 SetPolicyIdNil sets the value for PolicyId to be an explicit nil

### UnsetPolicyId
`func (o *CassandraIndexedObject) UnsetPolicyId()`

UnsetPolicyId ensures that no value is present for PolicyId, not even an explicit nil
### GetPolicyName

`func (o *CassandraIndexedObject) GetPolicyName() string`

GetPolicyName returns the PolicyName field if non-nil, zero value otherwise.

### GetPolicyNameOk

`func (o *CassandraIndexedObject) GetPolicyNameOk() (*string, bool)`

GetPolicyNameOk returns a tuple with the PolicyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyName

`func (o *CassandraIndexedObject) SetPolicyName(v string)`

SetPolicyName sets PolicyName field to given value.

### HasPolicyName

`func (o *CassandraIndexedObject) HasPolicyName() bool`

HasPolicyName returns a boolean if a field has been set.

### SetPolicyNameNil

`func (o *CassandraIndexedObject) SetPolicyNameNil(b bool)`

 SetPolicyNameNil sets the value for PolicyName to be an explicit nil

### UnsetPolicyName
`func (o *CassandraIndexedObject) UnsetPolicyName()`

UnsetPolicyName ensures that no value is present for PolicyName, not even an explicit nil
### GetProtectionGroupId

`func (o *CassandraIndexedObject) GetProtectionGroupId() string`

GetProtectionGroupId returns the ProtectionGroupId field if non-nil, zero value otherwise.

### GetProtectionGroupIdOk

`func (o *CassandraIndexedObject) GetProtectionGroupIdOk() (*string, bool)`

GetProtectionGroupIdOk returns a tuple with the ProtectionGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtectionGroupId

`func (o *CassandraIndexedObject) SetProtectionGroupId(v string)`

SetProtectionGroupId sets ProtectionGroupId field to given value.

### HasProtectionGroupId

`func (o *CassandraIndexedObject) HasProtectionGroupId() bool`

HasProtectionGroupId returns a boolean if a field has been set.

### SetProtectionGroupIdNil

`func (o *CassandraIndexedObject) SetProtectionGroupIdNil(b bool)`

 SetProtectionGroupIdNil sets the value for ProtectionGroupId to be an explicit nil

### UnsetProtectionGroupId
`func (o *CassandraIndexedObject) UnsetProtectionGroupId()`

UnsetProtectionGroupId ensures that no value is present for ProtectionGroupId, not even an explicit nil
### GetProtectionGroupName

`func (o *CassandraIndexedObject) GetProtectionGroupName() string`

GetProtectionGroupName returns the ProtectionGroupName field if non-nil, zero value otherwise.

### GetProtectionGroupNameOk

`func (o *CassandraIndexedObject) GetProtectionGroupNameOk() (*string, bool)`

GetProtectionGroupNameOk returns a tuple with the ProtectionGroupName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtectionGroupName

`func (o *CassandraIndexedObject) SetProtectionGroupName(v string)`

SetProtectionGroupName sets ProtectionGroupName field to given value.

### HasProtectionGroupName

`func (o *CassandraIndexedObject) HasProtectionGroupName() bool`

HasProtectionGroupName returns a boolean if a field has been set.

### SetProtectionGroupNameNil

`func (o *CassandraIndexedObject) SetProtectionGroupNameNil(b bool)`

 SetProtectionGroupNameNil sets the value for ProtectionGroupName to be an explicit nil

### UnsetProtectionGroupName
`func (o *CassandraIndexedObject) UnsetProtectionGroupName()`

UnsetProtectionGroupName ensures that no value is present for ProtectionGroupName, not even an explicit nil
### GetSourceInfo

`func (o *CassandraIndexedObject) GetSourceInfo() map[string]interface{}`

GetSourceInfo returns the SourceInfo field if non-nil, zero value otherwise.

### GetSourceInfoOk

`func (o *CassandraIndexedObject) GetSourceInfoOk() (*map[string]interface{}, bool)`

GetSourceInfoOk returns a tuple with the SourceInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceInfo

`func (o *CassandraIndexedObject) SetSourceInfo(v map[string]interface{})`

SetSourceInfo sets SourceInfo field to given value.

### HasSourceInfo

`func (o *CassandraIndexedObject) HasSourceInfo() bool`

HasSourceInfo returns a boolean if a field has been set.

### GetStorageDomainId

`func (o *CassandraIndexedObject) GetStorageDomainId() int64`

GetStorageDomainId returns the StorageDomainId field if non-nil, zero value otherwise.

### GetStorageDomainIdOk

`func (o *CassandraIndexedObject) GetStorageDomainIdOk() (*int64, bool)`

GetStorageDomainIdOk returns a tuple with the StorageDomainId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageDomainId

`func (o *CassandraIndexedObject) SetStorageDomainId(v int64)`

SetStorageDomainId sets StorageDomainId field to given value.

### HasStorageDomainId

`func (o *CassandraIndexedObject) HasStorageDomainId() bool`

HasStorageDomainId returns a boolean if a field has been set.

### SetStorageDomainIdNil

`func (o *CassandraIndexedObject) SetStorageDomainIdNil(b bool)`

 SetStorageDomainIdNil sets the value for StorageDomainId to be an explicit nil

### UnsetStorageDomainId
`func (o *CassandraIndexedObject) UnsetStorageDomainId()`

UnsetStorageDomainId ensures that no value is present for StorageDomainId, not even an explicit nil
### GetSnapshotTags

`func (o *CassandraIndexedObject) GetSnapshotTags() []SnapshotTagInfo`

GetSnapshotTags returns the SnapshotTags field if non-nil, zero value otherwise.

### GetSnapshotTagsOk

`func (o *CassandraIndexedObject) GetSnapshotTagsOk() (*[]SnapshotTagInfo, bool)`

GetSnapshotTagsOk returns a tuple with the SnapshotTags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnapshotTags

`func (o *CassandraIndexedObject) SetSnapshotTags(v []SnapshotTagInfo)`

SetSnapshotTags sets SnapshotTags field to given value.

### HasSnapshotTags

`func (o *CassandraIndexedObject) HasSnapshotTags() bool`

HasSnapshotTags returns a boolean if a field has been set.

### SetSnapshotTagsNil

`func (o *CassandraIndexedObject) SetSnapshotTagsNil(b bool)`

 SetSnapshotTagsNil sets the value for SnapshotTags to be an explicit nil

### UnsetSnapshotTags
`func (o *CassandraIndexedObject) UnsetSnapshotTags()`

UnsetSnapshotTags ensures that no value is present for SnapshotTags, not even an explicit nil
### GetTags

`func (o *CassandraIndexedObject) GetTags() []TagInfo`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *CassandraIndexedObject) GetTagsOk() (*[]TagInfo, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *CassandraIndexedObject) SetTags(v []TagInfo)`

SetTags sets Tags field to given value.

### HasTags

`func (o *CassandraIndexedObject) HasTags() bool`

HasTags returns a boolean if a field has been set.

### SetTagsNil

`func (o *CassandraIndexedObject) SetTagsNil(b bool)`

 SetTagsNil sets the value for Tags to be an explicit nil

### UnsetTags
`func (o *CassandraIndexedObject) UnsetTags()`

UnsetTags ensures that no value is present for Tags, not even an explicit nil
### GetId

`func (o *CassandraIndexedObject) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CassandraIndexedObject) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CassandraIndexedObject) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *CassandraIndexedObject) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *CassandraIndexedObject) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *CassandraIndexedObject) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetKeyspaceType

`func (o *CassandraIndexedObject) GetKeyspaceType() string`

GetKeyspaceType returns the KeyspaceType field if non-nil, zero value otherwise.

### GetKeyspaceTypeOk

`func (o *CassandraIndexedObject) GetKeyspaceTypeOk() (*string, bool)`

GetKeyspaceTypeOk returns a tuple with the KeyspaceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyspaceType

`func (o *CassandraIndexedObject) SetKeyspaceType(v string)`

SetKeyspaceType sets KeyspaceType field to given value.

### HasKeyspaceType

`func (o *CassandraIndexedObject) HasKeyspaceType() bool`

HasKeyspaceType returns a boolean if a field has been set.

### SetKeyspaceTypeNil

`func (o *CassandraIndexedObject) SetKeyspaceTypeNil(b bool)`

 SetKeyspaceTypeNil sets the value for KeyspaceType to be an explicit nil

### UnsetKeyspaceType
`func (o *CassandraIndexedObject) UnsetKeyspaceType()`

UnsetKeyspaceType ensures that no value is present for KeyspaceType, not even an explicit nil
### GetType

`func (o *CassandraIndexedObject) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CassandraIndexedObject) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CassandraIndexedObject) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *CassandraIndexedObject) HasType() bool`

HasType returns a boolean if a field has been set.

### SetTypeNil

`func (o *CassandraIndexedObject) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *CassandraIndexedObject) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


