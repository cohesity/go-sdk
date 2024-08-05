# DocumentLibraryItem

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
**CreationTimeSecs** | Pointer to **NullableInt64** | Specifies the Unix timestamp epoch in seconds at which this item is created. | [optional] 
**FileType** | Pointer to **NullableString** | Specifies the file type. | [optional] 
**ItemId** | Pointer to **NullableString** | Specifies the id of the document library item. | [optional] 
**ItemSize** | Pointer to **NullableInt64** | Specifies the size in bytes for the indexed item. | [optional] 
**OwnerEmail** | Pointer to **NullableString** | Specifies the email of the owner of the document library item. | [optional] 
**OwnerName** | Pointer to **NullableString** | Specifies the name of the owner of the document library item. | [optional] 

## Methods

### NewDocumentLibraryItem

`func NewDocumentLibraryItem() *DocumentLibraryItem`

NewDocumentLibraryItem instantiates a new DocumentLibraryItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDocumentLibraryItemWithDefaults

`func NewDocumentLibraryItemWithDefaults() *DocumentLibraryItem`

NewDocumentLibraryItemWithDefaults instantiates a new DocumentLibraryItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *DocumentLibraryItem) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DocumentLibraryItem) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DocumentLibraryItem) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *DocumentLibraryItem) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *DocumentLibraryItem) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *DocumentLibraryItem) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetPath

`func (o *DocumentLibraryItem) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *DocumentLibraryItem) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *DocumentLibraryItem) SetPath(v string)`

SetPath sets Path field to given value.

### HasPath

`func (o *DocumentLibraryItem) HasPath() bool`

HasPath returns a boolean if a field has been set.

### SetPathNil

`func (o *DocumentLibraryItem) SetPathNil(b bool)`

 SetPathNil sets the value for Path to be an explicit nil

### UnsetPath
`func (o *DocumentLibraryItem) UnsetPath()`

UnsetPath ensures that no value is present for Path, not even an explicit nil
### GetPolicyId

`func (o *DocumentLibraryItem) GetPolicyId() string`

GetPolicyId returns the PolicyId field if non-nil, zero value otherwise.

### GetPolicyIdOk

`func (o *DocumentLibraryItem) GetPolicyIdOk() (*string, bool)`

GetPolicyIdOk returns a tuple with the PolicyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyId

`func (o *DocumentLibraryItem) SetPolicyId(v string)`

SetPolicyId sets PolicyId field to given value.

### HasPolicyId

`func (o *DocumentLibraryItem) HasPolicyId() bool`

HasPolicyId returns a boolean if a field has been set.

### SetPolicyIdNil

`func (o *DocumentLibraryItem) SetPolicyIdNil(b bool)`

 SetPolicyIdNil sets the value for PolicyId to be an explicit nil

### UnsetPolicyId
`func (o *DocumentLibraryItem) UnsetPolicyId()`

UnsetPolicyId ensures that no value is present for PolicyId, not even an explicit nil
### GetPolicyName

`func (o *DocumentLibraryItem) GetPolicyName() string`

GetPolicyName returns the PolicyName field if non-nil, zero value otherwise.

### GetPolicyNameOk

`func (o *DocumentLibraryItem) GetPolicyNameOk() (*string, bool)`

GetPolicyNameOk returns a tuple with the PolicyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyName

`func (o *DocumentLibraryItem) SetPolicyName(v string)`

SetPolicyName sets PolicyName field to given value.

### HasPolicyName

`func (o *DocumentLibraryItem) HasPolicyName() bool`

HasPolicyName returns a boolean if a field has been set.

### SetPolicyNameNil

`func (o *DocumentLibraryItem) SetPolicyNameNil(b bool)`

 SetPolicyNameNil sets the value for PolicyName to be an explicit nil

### UnsetPolicyName
`func (o *DocumentLibraryItem) UnsetPolicyName()`

UnsetPolicyName ensures that no value is present for PolicyName, not even an explicit nil
### GetProtectionGroupId

`func (o *DocumentLibraryItem) GetProtectionGroupId() string`

GetProtectionGroupId returns the ProtectionGroupId field if non-nil, zero value otherwise.

### GetProtectionGroupIdOk

`func (o *DocumentLibraryItem) GetProtectionGroupIdOk() (*string, bool)`

GetProtectionGroupIdOk returns a tuple with the ProtectionGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtectionGroupId

`func (o *DocumentLibraryItem) SetProtectionGroupId(v string)`

SetProtectionGroupId sets ProtectionGroupId field to given value.

### HasProtectionGroupId

`func (o *DocumentLibraryItem) HasProtectionGroupId() bool`

HasProtectionGroupId returns a boolean if a field has been set.

### SetProtectionGroupIdNil

`func (o *DocumentLibraryItem) SetProtectionGroupIdNil(b bool)`

 SetProtectionGroupIdNil sets the value for ProtectionGroupId to be an explicit nil

### UnsetProtectionGroupId
`func (o *DocumentLibraryItem) UnsetProtectionGroupId()`

UnsetProtectionGroupId ensures that no value is present for ProtectionGroupId, not even an explicit nil
### GetProtectionGroupName

`func (o *DocumentLibraryItem) GetProtectionGroupName() string`

GetProtectionGroupName returns the ProtectionGroupName field if non-nil, zero value otherwise.

### GetProtectionGroupNameOk

`func (o *DocumentLibraryItem) GetProtectionGroupNameOk() (*string, bool)`

GetProtectionGroupNameOk returns a tuple with the ProtectionGroupName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtectionGroupName

`func (o *DocumentLibraryItem) SetProtectionGroupName(v string)`

SetProtectionGroupName sets ProtectionGroupName field to given value.

### HasProtectionGroupName

`func (o *DocumentLibraryItem) HasProtectionGroupName() bool`

HasProtectionGroupName returns a boolean if a field has been set.

### SetProtectionGroupNameNil

`func (o *DocumentLibraryItem) SetProtectionGroupNameNil(b bool)`

 SetProtectionGroupNameNil sets the value for ProtectionGroupName to be an explicit nil

### UnsetProtectionGroupName
`func (o *DocumentLibraryItem) UnsetProtectionGroupName()`

UnsetProtectionGroupName ensures that no value is present for ProtectionGroupName, not even an explicit nil
### GetSourceInfo

`func (o *DocumentLibraryItem) GetSourceInfo() map[string]interface{}`

GetSourceInfo returns the SourceInfo field if non-nil, zero value otherwise.

### GetSourceInfoOk

`func (o *DocumentLibraryItem) GetSourceInfoOk() (*map[string]interface{}, bool)`

GetSourceInfoOk returns a tuple with the SourceInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceInfo

`func (o *DocumentLibraryItem) SetSourceInfo(v map[string]interface{})`

SetSourceInfo sets SourceInfo field to given value.

### HasSourceInfo

`func (o *DocumentLibraryItem) HasSourceInfo() bool`

HasSourceInfo returns a boolean if a field has been set.

### GetStorageDomainId

`func (o *DocumentLibraryItem) GetStorageDomainId() int64`

GetStorageDomainId returns the StorageDomainId field if non-nil, zero value otherwise.

### GetStorageDomainIdOk

`func (o *DocumentLibraryItem) GetStorageDomainIdOk() (*int64, bool)`

GetStorageDomainIdOk returns a tuple with the StorageDomainId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageDomainId

`func (o *DocumentLibraryItem) SetStorageDomainId(v int64)`

SetStorageDomainId sets StorageDomainId field to given value.

### HasStorageDomainId

`func (o *DocumentLibraryItem) HasStorageDomainId() bool`

HasStorageDomainId returns a boolean if a field has been set.

### SetStorageDomainIdNil

`func (o *DocumentLibraryItem) SetStorageDomainIdNil(b bool)`

 SetStorageDomainIdNil sets the value for StorageDomainId to be an explicit nil

### UnsetStorageDomainId
`func (o *DocumentLibraryItem) UnsetStorageDomainId()`

UnsetStorageDomainId ensures that no value is present for StorageDomainId, not even an explicit nil
### GetSnapshotTags

`func (o *DocumentLibraryItem) GetSnapshotTags() []SnapshotTagInfo`

GetSnapshotTags returns the SnapshotTags field if non-nil, zero value otherwise.

### GetSnapshotTagsOk

`func (o *DocumentLibraryItem) GetSnapshotTagsOk() (*[]SnapshotTagInfo, bool)`

GetSnapshotTagsOk returns a tuple with the SnapshotTags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnapshotTags

`func (o *DocumentLibraryItem) SetSnapshotTags(v []SnapshotTagInfo)`

SetSnapshotTags sets SnapshotTags field to given value.

### HasSnapshotTags

`func (o *DocumentLibraryItem) HasSnapshotTags() bool`

HasSnapshotTags returns a boolean if a field has been set.

### SetSnapshotTagsNil

`func (o *DocumentLibraryItem) SetSnapshotTagsNil(b bool)`

 SetSnapshotTagsNil sets the value for SnapshotTags to be an explicit nil

### UnsetSnapshotTags
`func (o *DocumentLibraryItem) UnsetSnapshotTags()`

UnsetSnapshotTags ensures that no value is present for SnapshotTags, not even an explicit nil
### GetTags

`func (o *DocumentLibraryItem) GetTags() []TagInfo`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *DocumentLibraryItem) GetTagsOk() (*[]TagInfo, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *DocumentLibraryItem) SetTags(v []TagInfo)`

SetTags sets Tags field to given value.

### HasTags

`func (o *DocumentLibraryItem) HasTags() bool`

HasTags returns a boolean if a field has been set.

### SetTagsNil

`func (o *DocumentLibraryItem) SetTagsNil(b bool)`

 SetTagsNil sets the value for Tags to be an explicit nil

### UnsetTags
`func (o *DocumentLibraryItem) UnsetTags()`

UnsetTags ensures that no value is present for Tags, not even an explicit nil
### GetCreationTimeSecs

`func (o *DocumentLibraryItem) GetCreationTimeSecs() int64`

GetCreationTimeSecs returns the CreationTimeSecs field if non-nil, zero value otherwise.

### GetCreationTimeSecsOk

`func (o *DocumentLibraryItem) GetCreationTimeSecsOk() (*int64, bool)`

GetCreationTimeSecsOk returns a tuple with the CreationTimeSecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreationTimeSecs

`func (o *DocumentLibraryItem) SetCreationTimeSecs(v int64)`

SetCreationTimeSecs sets CreationTimeSecs field to given value.

### HasCreationTimeSecs

`func (o *DocumentLibraryItem) HasCreationTimeSecs() bool`

HasCreationTimeSecs returns a boolean if a field has been set.

### SetCreationTimeSecsNil

`func (o *DocumentLibraryItem) SetCreationTimeSecsNil(b bool)`

 SetCreationTimeSecsNil sets the value for CreationTimeSecs to be an explicit nil

### UnsetCreationTimeSecs
`func (o *DocumentLibraryItem) UnsetCreationTimeSecs()`

UnsetCreationTimeSecs ensures that no value is present for CreationTimeSecs, not even an explicit nil
### GetFileType

`func (o *DocumentLibraryItem) GetFileType() string`

GetFileType returns the FileType field if non-nil, zero value otherwise.

### GetFileTypeOk

`func (o *DocumentLibraryItem) GetFileTypeOk() (*string, bool)`

GetFileTypeOk returns a tuple with the FileType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileType

`func (o *DocumentLibraryItem) SetFileType(v string)`

SetFileType sets FileType field to given value.

### HasFileType

`func (o *DocumentLibraryItem) HasFileType() bool`

HasFileType returns a boolean if a field has been set.

### SetFileTypeNil

`func (o *DocumentLibraryItem) SetFileTypeNil(b bool)`

 SetFileTypeNil sets the value for FileType to be an explicit nil

### UnsetFileType
`func (o *DocumentLibraryItem) UnsetFileType()`

UnsetFileType ensures that no value is present for FileType, not even an explicit nil
### GetItemId

`func (o *DocumentLibraryItem) GetItemId() string`

GetItemId returns the ItemId field if non-nil, zero value otherwise.

### GetItemIdOk

`func (o *DocumentLibraryItem) GetItemIdOk() (*string, bool)`

GetItemIdOk returns a tuple with the ItemId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemId

`func (o *DocumentLibraryItem) SetItemId(v string)`

SetItemId sets ItemId field to given value.

### HasItemId

`func (o *DocumentLibraryItem) HasItemId() bool`

HasItemId returns a boolean if a field has been set.

### SetItemIdNil

`func (o *DocumentLibraryItem) SetItemIdNil(b bool)`

 SetItemIdNil sets the value for ItemId to be an explicit nil

### UnsetItemId
`func (o *DocumentLibraryItem) UnsetItemId()`

UnsetItemId ensures that no value is present for ItemId, not even an explicit nil
### GetItemSize

`func (o *DocumentLibraryItem) GetItemSize() int64`

GetItemSize returns the ItemSize field if non-nil, zero value otherwise.

### GetItemSizeOk

`func (o *DocumentLibraryItem) GetItemSizeOk() (*int64, bool)`

GetItemSizeOk returns a tuple with the ItemSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemSize

`func (o *DocumentLibraryItem) SetItemSize(v int64)`

SetItemSize sets ItemSize field to given value.

### HasItemSize

`func (o *DocumentLibraryItem) HasItemSize() bool`

HasItemSize returns a boolean if a field has been set.

### SetItemSizeNil

`func (o *DocumentLibraryItem) SetItemSizeNil(b bool)`

 SetItemSizeNil sets the value for ItemSize to be an explicit nil

### UnsetItemSize
`func (o *DocumentLibraryItem) UnsetItemSize()`

UnsetItemSize ensures that no value is present for ItemSize, not even an explicit nil
### GetOwnerEmail

`func (o *DocumentLibraryItem) GetOwnerEmail() string`

GetOwnerEmail returns the OwnerEmail field if non-nil, zero value otherwise.

### GetOwnerEmailOk

`func (o *DocumentLibraryItem) GetOwnerEmailOk() (*string, bool)`

GetOwnerEmailOk returns a tuple with the OwnerEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerEmail

`func (o *DocumentLibraryItem) SetOwnerEmail(v string)`

SetOwnerEmail sets OwnerEmail field to given value.

### HasOwnerEmail

`func (o *DocumentLibraryItem) HasOwnerEmail() bool`

HasOwnerEmail returns a boolean if a field has been set.

### SetOwnerEmailNil

`func (o *DocumentLibraryItem) SetOwnerEmailNil(b bool)`

 SetOwnerEmailNil sets the value for OwnerEmail to be an explicit nil

### UnsetOwnerEmail
`func (o *DocumentLibraryItem) UnsetOwnerEmail()`

UnsetOwnerEmail ensures that no value is present for OwnerEmail, not even an explicit nil
### GetOwnerName

`func (o *DocumentLibraryItem) GetOwnerName() string`

GetOwnerName returns the OwnerName field if non-nil, zero value otherwise.

### GetOwnerNameOk

`func (o *DocumentLibraryItem) GetOwnerNameOk() (*string, bool)`

GetOwnerNameOk returns a tuple with the OwnerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerName

`func (o *DocumentLibraryItem) SetOwnerName(v string)`

SetOwnerName sets OwnerName field to given value.

### HasOwnerName

`func (o *DocumentLibraryItem) HasOwnerName() bool`

HasOwnerName returns a boolean if a field has been set.

### SetOwnerNameNil

`func (o *DocumentLibraryItem) SetOwnerNameNil(b bool)`

 SetOwnerNameNil sets the value for OwnerName to be an explicit nil

### UnsetOwnerName
`func (o *DocumentLibraryItem) UnsetOwnerName()`

UnsetOwnerName ensures that no value is present for OwnerName, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


