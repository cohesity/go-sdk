# PublicFolderItem

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
**HasAttachments** | Pointer to **NullableBool** | Specifies whether the item has any attachments | [optional] 
**Id** | Pointer to **NullableString** | Specifies the id of the indexed item. | [optional] 
**ItemClass** | Pointer to **NullableString** | Specifies the item class of the indexed item. | [optional] 
**ItemSize** | Pointer to **NullableInt64** | Specifies the size in bytes for the indexed item. | [optional] 
**ParentFolderId** | Pointer to **NullableString** | Specifies the id of parent folder the indexed item. | [optional] 
**ReceivedTimeSecs** | Pointer to **NullableInt64** | Specifies the Unix timestamp epoch in seconds at which this item is received. | [optional] 
**Subject** | Pointer to **NullableString** | Specifies the subject of the indexed item. | [optional] 
**Type** | Pointer to **NullableString** | Specifies the Public folder item type. | [optional] 

## Methods

### NewPublicFolderItem

`func NewPublicFolderItem() *PublicFolderItem`

NewPublicFolderItem instantiates a new PublicFolderItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPublicFolderItemWithDefaults

`func NewPublicFolderItemWithDefaults() *PublicFolderItem`

NewPublicFolderItemWithDefaults instantiates a new PublicFolderItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *PublicFolderItem) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PublicFolderItem) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PublicFolderItem) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PublicFolderItem) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *PublicFolderItem) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *PublicFolderItem) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetPath

`func (o *PublicFolderItem) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *PublicFolderItem) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *PublicFolderItem) SetPath(v string)`

SetPath sets Path field to given value.

### HasPath

`func (o *PublicFolderItem) HasPath() bool`

HasPath returns a boolean if a field has been set.

### SetPathNil

`func (o *PublicFolderItem) SetPathNil(b bool)`

 SetPathNil sets the value for Path to be an explicit nil

### UnsetPath
`func (o *PublicFolderItem) UnsetPath()`

UnsetPath ensures that no value is present for Path, not even an explicit nil
### GetPolicyId

`func (o *PublicFolderItem) GetPolicyId() string`

GetPolicyId returns the PolicyId field if non-nil, zero value otherwise.

### GetPolicyIdOk

`func (o *PublicFolderItem) GetPolicyIdOk() (*string, bool)`

GetPolicyIdOk returns a tuple with the PolicyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyId

`func (o *PublicFolderItem) SetPolicyId(v string)`

SetPolicyId sets PolicyId field to given value.

### HasPolicyId

`func (o *PublicFolderItem) HasPolicyId() bool`

HasPolicyId returns a boolean if a field has been set.

### SetPolicyIdNil

`func (o *PublicFolderItem) SetPolicyIdNil(b bool)`

 SetPolicyIdNil sets the value for PolicyId to be an explicit nil

### UnsetPolicyId
`func (o *PublicFolderItem) UnsetPolicyId()`

UnsetPolicyId ensures that no value is present for PolicyId, not even an explicit nil
### GetPolicyName

`func (o *PublicFolderItem) GetPolicyName() string`

GetPolicyName returns the PolicyName field if non-nil, zero value otherwise.

### GetPolicyNameOk

`func (o *PublicFolderItem) GetPolicyNameOk() (*string, bool)`

GetPolicyNameOk returns a tuple with the PolicyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyName

`func (o *PublicFolderItem) SetPolicyName(v string)`

SetPolicyName sets PolicyName field to given value.

### HasPolicyName

`func (o *PublicFolderItem) HasPolicyName() bool`

HasPolicyName returns a boolean if a field has been set.

### SetPolicyNameNil

`func (o *PublicFolderItem) SetPolicyNameNil(b bool)`

 SetPolicyNameNil sets the value for PolicyName to be an explicit nil

### UnsetPolicyName
`func (o *PublicFolderItem) UnsetPolicyName()`

UnsetPolicyName ensures that no value is present for PolicyName, not even an explicit nil
### GetProtectionGroupId

`func (o *PublicFolderItem) GetProtectionGroupId() string`

GetProtectionGroupId returns the ProtectionGroupId field if non-nil, zero value otherwise.

### GetProtectionGroupIdOk

`func (o *PublicFolderItem) GetProtectionGroupIdOk() (*string, bool)`

GetProtectionGroupIdOk returns a tuple with the ProtectionGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtectionGroupId

`func (o *PublicFolderItem) SetProtectionGroupId(v string)`

SetProtectionGroupId sets ProtectionGroupId field to given value.

### HasProtectionGroupId

`func (o *PublicFolderItem) HasProtectionGroupId() bool`

HasProtectionGroupId returns a boolean if a field has been set.

### SetProtectionGroupIdNil

`func (o *PublicFolderItem) SetProtectionGroupIdNil(b bool)`

 SetProtectionGroupIdNil sets the value for ProtectionGroupId to be an explicit nil

### UnsetProtectionGroupId
`func (o *PublicFolderItem) UnsetProtectionGroupId()`

UnsetProtectionGroupId ensures that no value is present for ProtectionGroupId, not even an explicit nil
### GetProtectionGroupName

`func (o *PublicFolderItem) GetProtectionGroupName() string`

GetProtectionGroupName returns the ProtectionGroupName field if non-nil, zero value otherwise.

### GetProtectionGroupNameOk

`func (o *PublicFolderItem) GetProtectionGroupNameOk() (*string, bool)`

GetProtectionGroupNameOk returns a tuple with the ProtectionGroupName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtectionGroupName

`func (o *PublicFolderItem) SetProtectionGroupName(v string)`

SetProtectionGroupName sets ProtectionGroupName field to given value.

### HasProtectionGroupName

`func (o *PublicFolderItem) HasProtectionGroupName() bool`

HasProtectionGroupName returns a boolean if a field has been set.

### SetProtectionGroupNameNil

`func (o *PublicFolderItem) SetProtectionGroupNameNil(b bool)`

 SetProtectionGroupNameNil sets the value for ProtectionGroupName to be an explicit nil

### UnsetProtectionGroupName
`func (o *PublicFolderItem) UnsetProtectionGroupName()`

UnsetProtectionGroupName ensures that no value is present for ProtectionGroupName, not even an explicit nil
### GetSourceInfo

`func (o *PublicFolderItem) GetSourceInfo() map[string]interface{}`

GetSourceInfo returns the SourceInfo field if non-nil, zero value otherwise.

### GetSourceInfoOk

`func (o *PublicFolderItem) GetSourceInfoOk() (*map[string]interface{}, bool)`

GetSourceInfoOk returns a tuple with the SourceInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceInfo

`func (o *PublicFolderItem) SetSourceInfo(v map[string]interface{})`

SetSourceInfo sets SourceInfo field to given value.

### HasSourceInfo

`func (o *PublicFolderItem) HasSourceInfo() bool`

HasSourceInfo returns a boolean if a field has been set.

### GetStorageDomainId

`func (o *PublicFolderItem) GetStorageDomainId() int64`

GetStorageDomainId returns the StorageDomainId field if non-nil, zero value otherwise.

### GetStorageDomainIdOk

`func (o *PublicFolderItem) GetStorageDomainIdOk() (*int64, bool)`

GetStorageDomainIdOk returns a tuple with the StorageDomainId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageDomainId

`func (o *PublicFolderItem) SetStorageDomainId(v int64)`

SetStorageDomainId sets StorageDomainId field to given value.

### HasStorageDomainId

`func (o *PublicFolderItem) HasStorageDomainId() bool`

HasStorageDomainId returns a boolean if a field has been set.

### SetStorageDomainIdNil

`func (o *PublicFolderItem) SetStorageDomainIdNil(b bool)`

 SetStorageDomainIdNil sets the value for StorageDomainId to be an explicit nil

### UnsetStorageDomainId
`func (o *PublicFolderItem) UnsetStorageDomainId()`

UnsetStorageDomainId ensures that no value is present for StorageDomainId, not even an explicit nil
### GetSnapshotTags

`func (o *PublicFolderItem) GetSnapshotTags() []SnapshotTagInfo`

GetSnapshotTags returns the SnapshotTags field if non-nil, zero value otherwise.

### GetSnapshotTagsOk

`func (o *PublicFolderItem) GetSnapshotTagsOk() (*[]SnapshotTagInfo, bool)`

GetSnapshotTagsOk returns a tuple with the SnapshotTags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnapshotTags

`func (o *PublicFolderItem) SetSnapshotTags(v []SnapshotTagInfo)`

SetSnapshotTags sets SnapshotTags field to given value.

### HasSnapshotTags

`func (o *PublicFolderItem) HasSnapshotTags() bool`

HasSnapshotTags returns a boolean if a field has been set.

### SetSnapshotTagsNil

`func (o *PublicFolderItem) SetSnapshotTagsNil(b bool)`

 SetSnapshotTagsNil sets the value for SnapshotTags to be an explicit nil

### UnsetSnapshotTags
`func (o *PublicFolderItem) UnsetSnapshotTags()`

UnsetSnapshotTags ensures that no value is present for SnapshotTags, not even an explicit nil
### GetTags

`func (o *PublicFolderItem) GetTags() []TagInfo`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *PublicFolderItem) GetTagsOk() (*[]TagInfo, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *PublicFolderItem) SetTags(v []TagInfo)`

SetTags sets Tags field to given value.

### HasTags

`func (o *PublicFolderItem) HasTags() bool`

HasTags returns a boolean if a field has been set.

### SetTagsNil

`func (o *PublicFolderItem) SetTagsNil(b bool)`

 SetTagsNil sets the value for Tags to be an explicit nil

### UnsetTags
`func (o *PublicFolderItem) UnsetTags()`

UnsetTags ensures that no value is present for Tags, not even an explicit nil
### GetHasAttachments

`func (o *PublicFolderItem) GetHasAttachments() bool`

GetHasAttachments returns the HasAttachments field if non-nil, zero value otherwise.

### GetHasAttachmentsOk

`func (o *PublicFolderItem) GetHasAttachmentsOk() (*bool, bool)`

GetHasAttachmentsOk returns a tuple with the HasAttachments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasAttachments

`func (o *PublicFolderItem) SetHasAttachments(v bool)`

SetHasAttachments sets HasAttachments field to given value.

### HasHasAttachments

`func (o *PublicFolderItem) HasHasAttachments() bool`

HasHasAttachments returns a boolean if a field has been set.

### SetHasAttachmentsNil

`func (o *PublicFolderItem) SetHasAttachmentsNil(b bool)`

 SetHasAttachmentsNil sets the value for HasAttachments to be an explicit nil

### UnsetHasAttachments
`func (o *PublicFolderItem) UnsetHasAttachments()`

UnsetHasAttachments ensures that no value is present for HasAttachments, not even an explicit nil
### GetId

`func (o *PublicFolderItem) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PublicFolderItem) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PublicFolderItem) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *PublicFolderItem) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *PublicFolderItem) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *PublicFolderItem) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetItemClass

`func (o *PublicFolderItem) GetItemClass() string`

GetItemClass returns the ItemClass field if non-nil, zero value otherwise.

### GetItemClassOk

`func (o *PublicFolderItem) GetItemClassOk() (*string, bool)`

GetItemClassOk returns a tuple with the ItemClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemClass

`func (o *PublicFolderItem) SetItemClass(v string)`

SetItemClass sets ItemClass field to given value.

### HasItemClass

`func (o *PublicFolderItem) HasItemClass() bool`

HasItemClass returns a boolean if a field has been set.

### SetItemClassNil

`func (o *PublicFolderItem) SetItemClassNil(b bool)`

 SetItemClassNil sets the value for ItemClass to be an explicit nil

### UnsetItemClass
`func (o *PublicFolderItem) UnsetItemClass()`

UnsetItemClass ensures that no value is present for ItemClass, not even an explicit nil
### GetItemSize

`func (o *PublicFolderItem) GetItemSize() int64`

GetItemSize returns the ItemSize field if non-nil, zero value otherwise.

### GetItemSizeOk

`func (o *PublicFolderItem) GetItemSizeOk() (*int64, bool)`

GetItemSizeOk returns a tuple with the ItemSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemSize

`func (o *PublicFolderItem) SetItemSize(v int64)`

SetItemSize sets ItemSize field to given value.

### HasItemSize

`func (o *PublicFolderItem) HasItemSize() bool`

HasItemSize returns a boolean if a field has been set.

### SetItemSizeNil

`func (o *PublicFolderItem) SetItemSizeNil(b bool)`

 SetItemSizeNil sets the value for ItemSize to be an explicit nil

### UnsetItemSize
`func (o *PublicFolderItem) UnsetItemSize()`

UnsetItemSize ensures that no value is present for ItemSize, not even an explicit nil
### GetParentFolderId

`func (o *PublicFolderItem) GetParentFolderId() string`

GetParentFolderId returns the ParentFolderId field if non-nil, zero value otherwise.

### GetParentFolderIdOk

`func (o *PublicFolderItem) GetParentFolderIdOk() (*string, bool)`

GetParentFolderIdOk returns a tuple with the ParentFolderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentFolderId

`func (o *PublicFolderItem) SetParentFolderId(v string)`

SetParentFolderId sets ParentFolderId field to given value.

### HasParentFolderId

`func (o *PublicFolderItem) HasParentFolderId() bool`

HasParentFolderId returns a boolean if a field has been set.

### SetParentFolderIdNil

`func (o *PublicFolderItem) SetParentFolderIdNil(b bool)`

 SetParentFolderIdNil sets the value for ParentFolderId to be an explicit nil

### UnsetParentFolderId
`func (o *PublicFolderItem) UnsetParentFolderId()`

UnsetParentFolderId ensures that no value is present for ParentFolderId, not even an explicit nil
### GetReceivedTimeSecs

`func (o *PublicFolderItem) GetReceivedTimeSecs() int64`

GetReceivedTimeSecs returns the ReceivedTimeSecs field if non-nil, zero value otherwise.

### GetReceivedTimeSecsOk

`func (o *PublicFolderItem) GetReceivedTimeSecsOk() (*int64, bool)`

GetReceivedTimeSecsOk returns a tuple with the ReceivedTimeSecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReceivedTimeSecs

`func (o *PublicFolderItem) SetReceivedTimeSecs(v int64)`

SetReceivedTimeSecs sets ReceivedTimeSecs field to given value.

### HasReceivedTimeSecs

`func (o *PublicFolderItem) HasReceivedTimeSecs() bool`

HasReceivedTimeSecs returns a boolean if a field has been set.

### SetReceivedTimeSecsNil

`func (o *PublicFolderItem) SetReceivedTimeSecsNil(b bool)`

 SetReceivedTimeSecsNil sets the value for ReceivedTimeSecs to be an explicit nil

### UnsetReceivedTimeSecs
`func (o *PublicFolderItem) UnsetReceivedTimeSecs()`

UnsetReceivedTimeSecs ensures that no value is present for ReceivedTimeSecs, not even an explicit nil
### GetSubject

`func (o *PublicFolderItem) GetSubject() string`

GetSubject returns the Subject field if non-nil, zero value otherwise.

### GetSubjectOk

`func (o *PublicFolderItem) GetSubjectOk() (*string, bool)`

GetSubjectOk returns a tuple with the Subject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubject

`func (o *PublicFolderItem) SetSubject(v string)`

SetSubject sets Subject field to given value.

### HasSubject

`func (o *PublicFolderItem) HasSubject() bool`

HasSubject returns a boolean if a field has been set.

### SetSubjectNil

`func (o *PublicFolderItem) SetSubjectNil(b bool)`

 SetSubjectNil sets the value for Subject to be an explicit nil

### UnsetSubject
`func (o *PublicFolderItem) UnsetSubject()`

UnsetSubject ensures that no value is present for Subject, not even an explicit nil
### GetType

`func (o *PublicFolderItem) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PublicFolderItem) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PublicFolderItem) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *PublicFolderItem) HasType() bool`

HasType returns a boolean if a field has been set.

### SetTypeNil

`func (o *PublicFolderItem) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *PublicFolderItem) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


