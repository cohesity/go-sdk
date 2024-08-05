# HeliosPublicFolderItemsInner

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
**HasAttachments** | Pointer to **NullableBool** | Specifies whether the item has any attachments | [optional] 
**Id** | Pointer to **NullableString** | Specifies the id of the indexed item. | [optional] 
**ItemClass** | Pointer to **NullableString** | Specifies the item class of the indexed item. | [optional] 
**ItemSize** | Pointer to **NullableInt64** | Specifies the size in bytes for the indexed item. | [optional] 
**ParentFolderId** | Pointer to **NullableString** | Specifies the id of parent folder the indexed item. | [optional] 
**ReceivedTimeSecs** | Pointer to **NullableInt64** | Specifies the Unix timestamp epoch in seconds at which this item is received. | [optional] 
**Subject** | Pointer to **NullableString** | Specifies the subject of the indexed item. | [optional] 
**Type** | Pointer to **NullableString** | Specifies the Public folder item type. | [optional] 

## Methods

### NewHeliosPublicFolderItemsInner

`func NewHeliosPublicFolderItemsInner() *HeliosPublicFolderItemsInner`

NewHeliosPublicFolderItemsInner instantiates a new HeliosPublicFolderItemsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHeliosPublicFolderItemsInnerWithDefaults

`func NewHeliosPublicFolderItemsInnerWithDefaults() *HeliosPublicFolderItemsInner`

NewHeliosPublicFolderItemsInnerWithDefaults instantiates a new HeliosPublicFolderItemsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClusterIdentifier

`func (o *HeliosPublicFolderItemsInner) GetClusterIdentifier() string`

GetClusterIdentifier returns the ClusterIdentifier field if non-nil, zero value otherwise.

### GetClusterIdentifierOk

`func (o *HeliosPublicFolderItemsInner) GetClusterIdentifierOk() (*string, bool)`

GetClusterIdentifierOk returns a tuple with the ClusterIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterIdentifier

`func (o *HeliosPublicFolderItemsInner) SetClusterIdentifier(v string)`

SetClusterIdentifier sets ClusterIdentifier field to given value.

### HasClusterIdentifier

`func (o *HeliosPublicFolderItemsInner) HasClusterIdentifier() bool`

HasClusterIdentifier returns a boolean if a field has been set.

### SetClusterIdentifierNil

`func (o *HeliosPublicFolderItemsInner) SetClusterIdentifierNil(b bool)`

 SetClusterIdentifierNil sets the value for ClusterIdentifier to be an explicit nil

### UnsetClusterIdentifier
`func (o *HeliosPublicFolderItemsInner) UnsetClusterIdentifier()`

UnsetClusterIdentifier ensures that no value is present for ClusterIdentifier, not even an explicit nil
### GetRegionId

`func (o *HeliosPublicFolderItemsInner) GetRegionId() string`

GetRegionId returns the RegionId field if non-nil, zero value otherwise.

### GetRegionIdOk

`func (o *HeliosPublicFolderItemsInner) GetRegionIdOk() (*string, bool)`

GetRegionIdOk returns a tuple with the RegionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegionId

`func (o *HeliosPublicFolderItemsInner) SetRegionId(v string)`

SetRegionId sets RegionId field to given value.

### HasRegionId

`func (o *HeliosPublicFolderItemsInner) HasRegionId() bool`

HasRegionId returns a boolean if a field has been set.

### SetRegionIdNil

`func (o *HeliosPublicFolderItemsInner) SetRegionIdNil(b bool)`

 SetRegionIdNil sets the value for RegionId to be an explicit nil

### UnsetRegionId
`func (o *HeliosPublicFolderItemsInner) UnsetRegionId()`

UnsetRegionId ensures that no value is present for RegionId, not even an explicit nil
### GetName

`func (o *HeliosPublicFolderItemsInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *HeliosPublicFolderItemsInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *HeliosPublicFolderItemsInner) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *HeliosPublicFolderItemsInner) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *HeliosPublicFolderItemsInner) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *HeliosPublicFolderItemsInner) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetPath

`func (o *HeliosPublicFolderItemsInner) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *HeliosPublicFolderItemsInner) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *HeliosPublicFolderItemsInner) SetPath(v string)`

SetPath sets Path field to given value.

### HasPath

`func (o *HeliosPublicFolderItemsInner) HasPath() bool`

HasPath returns a boolean if a field has been set.

### SetPathNil

`func (o *HeliosPublicFolderItemsInner) SetPathNil(b bool)`

 SetPathNil sets the value for Path to be an explicit nil

### UnsetPath
`func (o *HeliosPublicFolderItemsInner) UnsetPath()`

UnsetPath ensures that no value is present for Path, not even an explicit nil
### GetPolicyId

`func (o *HeliosPublicFolderItemsInner) GetPolicyId() string`

GetPolicyId returns the PolicyId field if non-nil, zero value otherwise.

### GetPolicyIdOk

`func (o *HeliosPublicFolderItemsInner) GetPolicyIdOk() (*string, bool)`

GetPolicyIdOk returns a tuple with the PolicyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyId

`func (o *HeliosPublicFolderItemsInner) SetPolicyId(v string)`

SetPolicyId sets PolicyId field to given value.

### HasPolicyId

`func (o *HeliosPublicFolderItemsInner) HasPolicyId() bool`

HasPolicyId returns a boolean if a field has been set.

### SetPolicyIdNil

`func (o *HeliosPublicFolderItemsInner) SetPolicyIdNil(b bool)`

 SetPolicyIdNil sets the value for PolicyId to be an explicit nil

### UnsetPolicyId
`func (o *HeliosPublicFolderItemsInner) UnsetPolicyId()`

UnsetPolicyId ensures that no value is present for PolicyId, not even an explicit nil
### GetPolicyName

`func (o *HeliosPublicFolderItemsInner) GetPolicyName() string`

GetPolicyName returns the PolicyName field if non-nil, zero value otherwise.

### GetPolicyNameOk

`func (o *HeliosPublicFolderItemsInner) GetPolicyNameOk() (*string, bool)`

GetPolicyNameOk returns a tuple with the PolicyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyName

`func (o *HeliosPublicFolderItemsInner) SetPolicyName(v string)`

SetPolicyName sets PolicyName field to given value.

### HasPolicyName

`func (o *HeliosPublicFolderItemsInner) HasPolicyName() bool`

HasPolicyName returns a boolean if a field has been set.

### SetPolicyNameNil

`func (o *HeliosPublicFolderItemsInner) SetPolicyNameNil(b bool)`

 SetPolicyNameNil sets the value for PolicyName to be an explicit nil

### UnsetPolicyName
`func (o *HeliosPublicFolderItemsInner) UnsetPolicyName()`

UnsetPolicyName ensures that no value is present for PolicyName, not even an explicit nil
### GetProtectionGroupId

`func (o *HeliosPublicFolderItemsInner) GetProtectionGroupId() string`

GetProtectionGroupId returns the ProtectionGroupId field if non-nil, zero value otherwise.

### GetProtectionGroupIdOk

`func (o *HeliosPublicFolderItemsInner) GetProtectionGroupIdOk() (*string, bool)`

GetProtectionGroupIdOk returns a tuple with the ProtectionGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtectionGroupId

`func (o *HeliosPublicFolderItemsInner) SetProtectionGroupId(v string)`

SetProtectionGroupId sets ProtectionGroupId field to given value.

### HasProtectionGroupId

`func (o *HeliosPublicFolderItemsInner) HasProtectionGroupId() bool`

HasProtectionGroupId returns a boolean if a field has been set.

### SetProtectionGroupIdNil

`func (o *HeliosPublicFolderItemsInner) SetProtectionGroupIdNil(b bool)`

 SetProtectionGroupIdNil sets the value for ProtectionGroupId to be an explicit nil

### UnsetProtectionGroupId
`func (o *HeliosPublicFolderItemsInner) UnsetProtectionGroupId()`

UnsetProtectionGroupId ensures that no value is present for ProtectionGroupId, not even an explicit nil
### GetProtectionGroupName

`func (o *HeliosPublicFolderItemsInner) GetProtectionGroupName() string`

GetProtectionGroupName returns the ProtectionGroupName field if non-nil, zero value otherwise.

### GetProtectionGroupNameOk

`func (o *HeliosPublicFolderItemsInner) GetProtectionGroupNameOk() (*string, bool)`

GetProtectionGroupNameOk returns a tuple with the ProtectionGroupName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtectionGroupName

`func (o *HeliosPublicFolderItemsInner) SetProtectionGroupName(v string)`

SetProtectionGroupName sets ProtectionGroupName field to given value.

### HasProtectionGroupName

`func (o *HeliosPublicFolderItemsInner) HasProtectionGroupName() bool`

HasProtectionGroupName returns a boolean if a field has been set.

### SetProtectionGroupNameNil

`func (o *HeliosPublicFolderItemsInner) SetProtectionGroupNameNil(b bool)`

 SetProtectionGroupNameNil sets the value for ProtectionGroupName to be an explicit nil

### UnsetProtectionGroupName
`func (o *HeliosPublicFolderItemsInner) UnsetProtectionGroupName()`

UnsetProtectionGroupName ensures that no value is present for ProtectionGroupName, not even an explicit nil
### GetSourceInfo

`func (o *HeliosPublicFolderItemsInner) GetSourceInfo() map[string]interface{}`

GetSourceInfo returns the SourceInfo field if non-nil, zero value otherwise.

### GetSourceInfoOk

`func (o *HeliosPublicFolderItemsInner) GetSourceInfoOk() (*map[string]interface{}, bool)`

GetSourceInfoOk returns a tuple with the SourceInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceInfo

`func (o *HeliosPublicFolderItemsInner) SetSourceInfo(v map[string]interface{})`

SetSourceInfo sets SourceInfo field to given value.

### HasSourceInfo

`func (o *HeliosPublicFolderItemsInner) HasSourceInfo() bool`

HasSourceInfo returns a boolean if a field has been set.

### GetStorageDomainId

`func (o *HeliosPublicFolderItemsInner) GetStorageDomainId() int64`

GetStorageDomainId returns the StorageDomainId field if non-nil, zero value otherwise.

### GetStorageDomainIdOk

`func (o *HeliosPublicFolderItemsInner) GetStorageDomainIdOk() (*int64, bool)`

GetStorageDomainIdOk returns a tuple with the StorageDomainId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageDomainId

`func (o *HeliosPublicFolderItemsInner) SetStorageDomainId(v int64)`

SetStorageDomainId sets StorageDomainId field to given value.

### HasStorageDomainId

`func (o *HeliosPublicFolderItemsInner) HasStorageDomainId() bool`

HasStorageDomainId returns a boolean if a field has been set.

### SetStorageDomainIdNil

`func (o *HeliosPublicFolderItemsInner) SetStorageDomainIdNil(b bool)`

 SetStorageDomainIdNil sets the value for StorageDomainId to be an explicit nil

### UnsetStorageDomainId
`func (o *HeliosPublicFolderItemsInner) UnsetStorageDomainId()`

UnsetStorageDomainId ensures that no value is present for StorageDomainId, not even an explicit nil
### GetSnapshotTags

`func (o *HeliosPublicFolderItemsInner) GetSnapshotTags() []SnapshotTagInfo`

GetSnapshotTags returns the SnapshotTags field if non-nil, zero value otherwise.

### GetSnapshotTagsOk

`func (o *HeliosPublicFolderItemsInner) GetSnapshotTagsOk() (*[]SnapshotTagInfo, bool)`

GetSnapshotTagsOk returns a tuple with the SnapshotTags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnapshotTags

`func (o *HeliosPublicFolderItemsInner) SetSnapshotTags(v []SnapshotTagInfo)`

SetSnapshotTags sets SnapshotTags field to given value.

### HasSnapshotTags

`func (o *HeliosPublicFolderItemsInner) HasSnapshotTags() bool`

HasSnapshotTags returns a boolean if a field has been set.

### SetSnapshotTagsNil

`func (o *HeliosPublicFolderItemsInner) SetSnapshotTagsNil(b bool)`

 SetSnapshotTagsNil sets the value for SnapshotTags to be an explicit nil

### UnsetSnapshotTags
`func (o *HeliosPublicFolderItemsInner) UnsetSnapshotTags()`

UnsetSnapshotTags ensures that no value is present for SnapshotTags, not even an explicit nil
### GetTags

`func (o *HeliosPublicFolderItemsInner) GetTags() []TagInfo`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *HeliosPublicFolderItemsInner) GetTagsOk() (*[]TagInfo, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *HeliosPublicFolderItemsInner) SetTags(v []TagInfo)`

SetTags sets Tags field to given value.

### HasTags

`func (o *HeliosPublicFolderItemsInner) HasTags() bool`

HasTags returns a boolean if a field has been set.

### SetTagsNil

`func (o *HeliosPublicFolderItemsInner) SetTagsNil(b bool)`

 SetTagsNil sets the value for Tags to be an explicit nil

### UnsetTags
`func (o *HeliosPublicFolderItemsInner) UnsetTags()`

UnsetTags ensures that no value is present for Tags, not even an explicit nil
### GetHasAttachments

`func (o *HeliosPublicFolderItemsInner) GetHasAttachments() bool`

GetHasAttachments returns the HasAttachments field if non-nil, zero value otherwise.

### GetHasAttachmentsOk

`func (o *HeliosPublicFolderItemsInner) GetHasAttachmentsOk() (*bool, bool)`

GetHasAttachmentsOk returns a tuple with the HasAttachments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasAttachments

`func (o *HeliosPublicFolderItemsInner) SetHasAttachments(v bool)`

SetHasAttachments sets HasAttachments field to given value.

### HasHasAttachments

`func (o *HeliosPublicFolderItemsInner) HasHasAttachments() bool`

HasHasAttachments returns a boolean if a field has been set.

### SetHasAttachmentsNil

`func (o *HeliosPublicFolderItemsInner) SetHasAttachmentsNil(b bool)`

 SetHasAttachmentsNil sets the value for HasAttachments to be an explicit nil

### UnsetHasAttachments
`func (o *HeliosPublicFolderItemsInner) UnsetHasAttachments()`

UnsetHasAttachments ensures that no value is present for HasAttachments, not even an explicit nil
### GetId

`func (o *HeliosPublicFolderItemsInner) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *HeliosPublicFolderItemsInner) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *HeliosPublicFolderItemsInner) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *HeliosPublicFolderItemsInner) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *HeliosPublicFolderItemsInner) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *HeliosPublicFolderItemsInner) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetItemClass

`func (o *HeliosPublicFolderItemsInner) GetItemClass() string`

GetItemClass returns the ItemClass field if non-nil, zero value otherwise.

### GetItemClassOk

`func (o *HeliosPublicFolderItemsInner) GetItemClassOk() (*string, bool)`

GetItemClassOk returns a tuple with the ItemClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemClass

`func (o *HeliosPublicFolderItemsInner) SetItemClass(v string)`

SetItemClass sets ItemClass field to given value.

### HasItemClass

`func (o *HeliosPublicFolderItemsInner) HasItemClass() bool`

HasItemClass returns a boolean if a field has been set.

### SetItemClassNil

`func (o *HeliosPublicFolderItemsInner) SetItemClassNil(b bool)`

 SetItemClassNil sets the value for ItemClass to be an explicit nil

### UnsetItemClass
`func (o *HeliosPublicFolderItemsInner) UnsetItemClass()`

UnsetItemClass ensures that no value is present for ItemClass, not even an explicit nil
### GetItemSize

`func (o *HeliosPublicFolderItemsInner) GetItemSize() int64`

GetItemSize returns the ItemSize field if non-nil, zero value otherwise.

### GetItemSizeOk

`func (o *HeliosPublicFolderItemsInner) GetItemSizeOk() (*int64, bool)`

GetItemSizeOk returns a tuple with the ItemSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemSize

`func (o *HeliosPublicFolderItemsInner) SetItemSize(v int64)`

SetItemSize sets ItemSize field to given value.

### HasItemSize

`func (o *HeliosPublicFolderItemsInner) HasItemSize() bool`

HasItemSize returns a boolean if a field has been set.

### SetItemSizeNil

`func (o *HeliosPublicFolderItemsInner) SetItemSizeNil(b bool)`

 SetItemSizeNil sets the value for ItemSize to be an explicit nil

### UnsetItemSize
`func (o *HeliosPublicFolderItemsInner) UnsetItemSize()`

UnsetItemSize ensures that no value is present for ItemSize, not even an explicit nil
### GetParentFolderId

`func (o *HeliosPublicFolderItemsInner) GetParentFolderId() string`

GetParentFolderId returns the ParentFolderId field if non-nil, zero value otherwise.

### GetParentFolderIdOk

`func (o *HeliosPublicFolderItemsInner) GetParentFolderIdOk() (*string, bool)`

GetParentFolderIdOk returns a tuple with the ParentFolderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentFolderId

`func (o *HeliosPublicFolderItemsInner) SetParentFolderId(v string)`

SetParentFolderId sets ParentFolderId field to given value.

### HasParentFolderId

`func (o *HeliosPublicFolderItemsInner) HasParentFolderId() bool`

HasParentFolderId returns a boolean if a field has been set.

### SetParentFolderIdNil

`func (o *HeliosPublicFolderItemsInner) SetParentFolderIdNil(b bool)`

 SetParentFolderIdNil sets the value for ParentFolderId to be an explicit nil

### UnsetParentFolderId
`func (o *HeliosPublicFolderItemsInner) UnsetParentFolderId()`

UnsetParentFolderId ensures that no value is present for ParentFolderId, not even an explicit nil
### GetReceivedTimeSecs

`func (o *HeliosPublicFolderItemsInner) GetReceivedTimeSecs() int64`

GetReceivedTimeSecs returns the ReceivedTimeSecs field if non-nil, zero value otherwise.

### GetReceivedTimeSecsOk

`func (o *HeliosPublicFolderItemsInner) GetReceivedTimeSecsOk() (*int64, bool)`

GetReceivedTimeSecsOk returns a tuple with the ReceivedTimeSecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReceivedTimeSecs

`func (o *HeliosPublicFolderItemsInner) SetReceivedTimeSecs(v int64)`

SetReceivedTimeSecs sets ReceivedTimeSecs field to given value.

### HasReceivedTimeSecs

`func (o *HeliosPublicFolderItemsInner) HasReceivedTimeSecs() bool`

HasReceivedTimeSecs returns a boolean if a field has been set.

### SetReceivedTimeSecsNil

`func (o *HeliosPublicFolderItemsInner) SetReceivedTimeSecsNil(b bool)`

 SetReceivedTimeSecsNil sets the value for ReceivedTimeSecs to be an explicit nil

### UnsetReceivedTimeSecs
`func (o *HeliosPublicFolderItemsInner) UnsetReceivedTimeSecs()`

UnsetReceivedTimeSecs ensures that no value is present for ReceivedTimeSecs, not even an explicit nil
### GetSubject

`func (o *HeliosPublicFolderItemsInner) GetSubject() string`

GetSubject returns the Subject field if non-nil, zero value otherwise.

### GetSubjectOk

`func (o *HeliosPublicFolderItemsInner) GetSubjectOk() (*string, bool)`

GetSubjectOk returns a tuple with the Subject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubject

`func (o *HeliosPublicFolderItemsInner) SetSubject(v string)`

SetSubject sets Subject field to given value.

### HasSubject

`func (o *HeliosPublicFolderItemsInner) HasSubject() bool`

HasSubject returns a boolean if a field has been set.

### SetSubjectNil

`func (o *HeliosPublicFolderItemsInner) SetSubjectNil(b bool)`

 SetSubjectNil sets the value for Subject to be an explicit nil

### UnsetSubject
`func (o *HeliosPublicFolderItemsInner) UnsetSubject()`

UnsetSubject ensures that no value is present for Subject, not even an explicit nil
### GetType

`func (o *HeliosPublicFolderItemsInner) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *HeliosPublicFolderItemsInner) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *HeliosPublicFolderItemsInner) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *HeliosPublicFolderItemsInner) HasType() bool`

HasType returns a boolean if a field has been set.

### SetTypeNil

`func (o *HeliosPublicFolderItemsInner) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *HeliosPublicFolderItemsInner) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


