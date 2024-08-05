# MsGroupItem

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
**MailboxItem** | Pointer to [**Email**](Email.md) |  | [optional] 
**SiteItem** | Pointer to [**DocumentLibraryItem**](DocumentLibraryItem.md) |  | [optional] 
**Type** | Pointer to **NullableString** | Specifies the M365 Group item type. | [optional] 

## Methods

### NewMsGroupItem

`func NewMsGroupItem() *MsGroupItem`

NewMsGroupItem instantiates a new MsGroupItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMsGroupItemWithDefaults

`func NewMsGroupItemWithDefaults() *MsGroupItem`

NewMsGroupItemWithDefaults instantiates a new MsGroupItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *MsGroupItem) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *MsGroupItem) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *MsGroupItem) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *MsGroupItem) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *MsGroupItem) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *MsGroupItem) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetPath

`func (o *MsGroupItem) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *MsGroupItem) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *MsGroupItem) SetPath(v string)`

SetPath sets Path field to given value.

### HasPath

`func (o *MsGroupItem) HasPath() bool`

HasPath returns a boolean if a field has been set.

### SetPathNil

`func (o *MsGroupItem) SetPathNil(b bool)`

 SetPathNil sets the value for Path to be an explicit nil

### UnsetPath
`func (o *MsGroupItem) UnsetPath()`

UnsetPath ensures that no value is present for Path, not even an explicit nil
### GetPolicyId

`func (o *MsGroupItem) GetPolicyId() string`

GetPolicyId returns the PolicyId field if non-nil, zero value otherwise.

### GetPolicyIdOk

`func (o *MsGroupItem) GetPolicyIdOk() (*string, bool)`

GetPolicyIdOk returns a tuple with the PolicyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyId

`func (o *MsGroupItem) SetPolicyId(v string)`

SetPolicyId sets PolicyId field to given value.

### HasPolicyId

`func (o *MsGroupItem) HasPolicyId() bool`

HasPolicyId returns a boolean if a field has been set.

### SetPolicyIdNil

`func (o *MsGroupItem) SetPolicyIdNil(b bool)`

 SetPolicyIdNil sets the value for PolicyId to be an explicit nil

### UnsetPolicyId
`func (o *MsGroupItem) UnsetPolicyId()`

UnsetPolicyId ensures that no value is present for PolicyId, not even an explicit nil
### GetPolicyName

`func (o *MsGroupItem) GetPolicyName() string`

GetPolicyName returns the PolicyName field if non-nil, zero value otherwise.

### GetPolicyNameOk

`func (o *MsGroupItem) GetPolicyNameOk() (*string, bool)`

GetPolicyNameOk returns a tuple with the PolicyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyName

`func (o *MsGroupItem) SetPolicyName(v string)`

SetPolicyName sets PolicyName field to given value.

### HasPolicyName

`func (o *MsGroupItem) HasPolicyName() bool`

HasPolicyName returns a boolean if a field has been set.

### SetPolicyNameNil

`func (o *MsGroupItem) SetPolicyNameNil(b bool)`

 SetPolicyNameNil sets the value for PolicyName to be an explicit nil

### UnsetPolicyName
`func (o *MsGroupItem) UnsetPolicyName()`

UnsetPolicyName ensures that no value is present for PolicyName, not even an explicit nil
### GetProtectionGroupId

`func (o *MsGroupItem) GetProtectionGroupId() string`

GetProtectionGroupId returns the ProtectionGroupId field if non-nil, zero value otherwise.

### GetProtectionGroupIdOk

`func (o *MsGroupItem) GetProtectionGroupIdOk() (*string, bool)`

GetProtectionGroupIdOk returns a tuple with the ProtectionGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtectionGroupId

`func (o *MsGroupItem) SetProtectionGroupId(v string)`

SetProtectionGroupId sets ProtectionGroupId field to given value.

### HasProtectionGroupId

`func (o *MsGroupItem) HasProtectionGroupId() bool`

HasProtectionGroupId returns a boolean if a field has been set.

### SetProtectionGroupIdNil

`func (o *MsGroupItem) SetProtectionGroupIdNil(b bool)`

 SetProtectionGroupIdNil sets the value for ProtectionGroupId to be an explicit nil

### UnsetProtectionGroupId
`func (o *MsGroupItem) UnsetProtectionGroupId()`

UnsetProtectionGroupId ensures that no value is present for ProtectionGroupId, not even an explicit nil
### GetProtectionGroupName

`func (o *MsGroupItem) GetProtectionGroupName() string`

GetProtectionGroupName returns the ProtectionGroupName field if non-nil, zero value otherwise.

### GetProtectionGroupNameOk

`func (o *MsGroupItem) GetProtectionGroupNameOk() (*string, bool)`

GetProtectionGroupNameOk returns a tuple with the ProtectionGroupName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtectionGroupName

`func (o *MsGroupItem) SetProtectionGroupName(v string)`

SetProtectionGroupName sets ProtectionGroupName field to given value.

### HasProtectionGroupName

`func (o *MsGroupItem) HasProtectionGroupName() bool`

HasProtectionGroupName returns a boolean if a field has been set.

### SetProtectionGroupNameNil

`func (o *MsGroupItem) SetProtectionGroupNameNil(b bool)`

 SetProtectionGroupNameNil sets the value for ProtectionGroupName to be an explicit nil

### UnsetProtectionGroupName
`func (o *MsGroupItem) UnsetProtectionGroupName()`

UnsetProtectionGroupName ensures that no value is present for ProtectionGroupName, not even an explicit nil
### GetSourceInfo

`func (o *MsGroupItem) GetSourceInfo() map[string]interface{}`

GetSourceInfo returns the SourceInfo field if non-nil, zero value otherwise.

### GetSourceInfoOk

`func (o *MsGroupItem) GetSourceInfoOk() (*map[string]interface{}, bool)`

GetSourceInfoOk returns a tuple with the SourceInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceInfo

`func (o *MsGroupItem) SetSourceInfo(v map[string]interface{})`

SetSourceInfo sets SourceInfo field to given value.

### HasSourceInfo

`func (o *MsGroupItem) HasSourceInfo() bool`

HasSourceInfo returns a boolean if a field has been set.

### GetStorageDomainId

`func (o *MsGroupItem) GetStorageDomainId() int64`

GetStorageDomainId returns the StorageDomainId field if non-nil, zero value otherwise.

### GetStorageDomainIdOk

`func (o *MsGroupItem) GetStorageDomainIdOk() (*int64, bool)`

GetStorageDomainIdOk returns a tuple with the StorageDomainId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageDomainId

`func (o *MsGroupItem) SetStorageDomainId(v int64)`

SetStorageDomainId sets StorageDomainId field to given value.

### HasStorageDomainId

`func (o *MsGroupItem) HasStorageDomainId() bool`

HasStorageDomainId returns a boolean if a field has been set.

### SetStorageDomainIdNil

`func (o *MsGroupItem) SetStorageDomainIdNil(b bool)`

 SetStorageDomainIdNil sets the value for StorageDomainId to be an explicit nil

### UnsetStorageDomainId
`func (o *MsGroupItem) UnsetStorageDomainId()`

UnsetStorageDomainId ensures that no value is present for StorageDomainId, not even an explicit nil
### GetSnapshotTags

`func (o *MsGroupItem) GetSnapshotTags() []SnapshotTagInfo`

GetSnapshotTags returns the SnapshotTags field if non-nil, zero value otherwise.

### GetSnapshotTagsOk

`func (o *MsGroupItem) GetSnapshotTagsOk() (*[]SnapshotTagInfo, bool)`

GetSnapshotTagsOk returns a tuple with the SnapshotTags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnapshotTags

`func (o *MsGroupItem) SetSnapshotTags(v []SnapshotTagInfo)`

SetSnapshotTags sets SnapshotTags field to given value.

### HasSnapshotTags

`func (o *MsGroupItem) HasSnapshotTags() bool`

HasSnapshotTags returns a boolean if a field has been set.

### SetSnapshotTagsNil

`func (o *MsGroupItem) SetSnapshotTagsNil(b bool)`

 SetSnapshotTagsNil sets the value for SnapshotTags to be an explicit nil

### UnsetSnapshotTags
`func (o *MsGroupItem) UnsetSnapshotTags()`

UnsetSnapshotTags ensures that no value is present for SnapshotTags, not even an explicit nil
### GetTags

`func (o *MsGroupItem) GetTags() []TagInfo`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *MsGroupItem) GetTagsOk() (*[]TagInfo, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *MsGroupItem) SetTags(v []TagInfo)`

SetTags sets Tags field to given value.

### HasTags

`func (o *MsGroupItem) HasTags() bool`

HasTags returns a boolean if a field has been set.

### SetTagsNil

`func (o *MsGroupItem) SetTagsNil(b bool)`

 SetTagsNil sets the value for Tags to be an explicit nil

### UnsetTags
`func (o *MsGroupItem) UnsetTags()`

UnsetTags ensures that no value is present for Tags, not even an explicit nil
### GetMailboxItem

`func (o *MsGroupItem) GetMailboxItem() Email`

GetMailboxItem returns the MailboxItem field if non-nil, zero value otherwise.

### GetMailboxItemOk

`func (o *MsGroupItem) GetMailboxItemOk() (*Email, bool)`

GetMailboxItemOk returns a tuple with the MailboxItem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMailboxItem

`func (o *MsGroupItem) SetMailboxItem(v Email)`

SetMailboxItem sets MailboxItem field to given value.

### HasMailboxItem

`func (o *MsGroupItem) HasMailboxItem() bool`

HasMailboxItem returns a boolean if a field has been set.

### GetSiteItem

`func (o *MsGroupItem) GetSiteItem() DocumentLibraryItem`

GetSiteItem returns the SiteItem field if non-nil, zero value otherwise.

### GetSiteItemOk

`func (o *MsGroupItem) GetSiteItemOk() (*DocumentLibraryItem, bool)`

GetSiteItemOk returns a tuple with the SiteItem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteItem

`func (o *MsGroupItem) SetSiteItem(v DocumentLibraryItem)`

SetSiteItem sets SiteItem field to given value.

### HasSiteItem

`func (o *MsGroupItem) HasSiteItem() bool`

HasSiteItem returns a boolean if a field has been set.

### GetType

`func (o *MsGroupItem) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *MsGroupItem) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *MsGroupItem) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *MsGroupItem) HasType() bool`

HasType returns a boolean if a field has been set.

### SetTypeNil

`func (o *MsGroupItem) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *MsGroupItem) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


