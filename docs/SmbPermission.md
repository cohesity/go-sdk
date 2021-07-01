# SmbPermission

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Access** | Pointer to **NullableString** | Specifies the read/write access to the SMB share. &#39;kReadyOnly&#39; indicates read only access to the SMB share. &#39;kReadWrite&#39; indicates read and write access to the SMB share. &#39;kFullControl&#39; indicates full administrative control of the SMB share. &#39;kSpecialAccess&#39; indicates custom permissions to the SMB share using access masks structures. &#39;kSuperUser&#39; indicates root permissions ignoring all SMB ACLs. | [optional] 
**Mode** | Pointer to **NullableString** | Specifies how the permission should be applied to folders and/or files. &#39;kFolderSubFoldersAndFiles&#39; indicates that permissions are applied to a Folder and it&#39;s sub folders and files. &#39;kFolderAndSubFolders&#39; indicates that permissions are applied to a Folder and it&#39;s sub folders. &#39;kFolderAndSubFiles&#39; indicates that permissions are applied to a Folder and it&#39;s sub files. &#39;kFolderOnly&#39; indicates that permsission are applied to folder only. &#39;kSubFoldersAndFilesOnly&#39; indicates that permissions are applied to sub folders and files only. &#39;kSubFoldersOnly&#39; indicates that permissiona are applied to sub folders only. &#39;kFilesOnly&#39; indicates that permissions are applied to files only. | [optional] 
**Sid** | Pointer to **NullableString** | Specifies the security identifier (SID) of the principal. | [optional] 
**SpecialAccessMask** | Pointer to **NullableInt32** | Specifies custom access permissions. When the access mask from the Access Control Entry (ACE) cannot be mapped to one of the enums in &#39;access&#39;, this field is populated with the custom mask derived from the ACE and &#39;access&#39; is set to kSpecialAccess. This is a placeholder for storing an unmapped access permission and should not be set when creating and editing a View. | [optional] 
**SpecialType** | Pointer to **NullableInt32** | Specifies a custom type. When the type from the Access Control Entry (ACE) cannot be mapped to one of the enums in &#39;type&#39;, this field is populated with the custom type derived from the ACE and &#39;type&#39; is set to kSpecialType. This is a placeholder for storing an unmapped type and should not be set when creating and editing a View. | [optional] 
**Type** | Pointer to **NullableString** | Specifies the type of permission. &#39;kAllow&#39; indicates access is allowed. &#39;kDeny&#39; indicates access is denied. &#39;kSpecialType&#39; indicates a type defined in the Access Control Entry (ACE) does not map to &#39;kAllow&#39; or &#39;kDeny&#39;. | [optional] 

## Methods

### NewSmbPermission

`func NewSmbPermission() *SmbPermission`

NewSmbPermission instantiates a new SmbPermission object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSmbPermissionWithDefaults

`func NewSmbPermissionWithDefaults() *SmbPermission`

NewSmbPermissionWithDefaults instantiates a new SmbPermission object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccess

`func (o *SmbPermission) GetAccess() string`

GetAccess returns the Access field if non-nil, zero value otherwise.

### GetAccessOk

`func (o *SmbPermission) GetAccessOk() (*string, bool)`

GetAccessOk returns a tuple with the Access field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccess

`func (o *SmbPermission) SetAccess(v string)`

SetAccess sets Access field to given value.

### HasAccess

`func (o *SmbPermission) HasAccess() bool`

HasAccess returns a boolean if a field has been set.

### SetAccessNil

`func (o *SmbPermission) SetAccessNil(b bool)`

 SetAccessNil sets the value for Access to be an explicit nil

### UnsetAccess
`func (o *SmbPermission) UnsetAccess()`

UnsetAccess ensures that no value is present for Access, not even an explicit nil
### GetMode

`func (o *SmbPermission) GetMode() string`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *SmbPermission) GetModeOk() (*string, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *SmbPermission) SetMode(v string)`

SetMode sets Mode field to given value.

### HasMode

`func (o *SmbPermission) HasMode() bool`

HasMode returns a boolean if a field has been set.

### SetModeNil

`func (o *SmbPermission) SetModeNil(b bool)`

 SetModeNil sets the value for Mode to be an explicit nil

### UnsetMode
`func (o *SmbPermission) UnsetMode()`

UnsetMode ensures that no value is present for Mode, not even an explicit nil
### GetSid

`func (o *SmbPermission) GetSid() string`

GetSid returns the Sid field if non-nil, zero value otherwise.

### GetSidOk

`func (o *SmbPermission) GetSidOk() (*string, bool)`

GetSidOk returns a tuple with the Sid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSid

`func (o *SmbPermission) SetSid(v string)`

SetSid sets Sid field to given value.

### HasSid

`func (o *SmbPermission) HasSid() bool`

HasSid returns a boolean if a field has been set.

### SetSidNil

`func (o *SmbPermission) SetSidNil(b bool)`

 SetSidNil sets the value for Sid to be an explicit nil

### UnsetSid
`func (o *SmbPermission) UnsetSid()`

UnsetSid ensures that no value is present for Sid, not even an explicit nil
### GetSpecialAccessMask

`func (o *SmbPermission) GetSpecialAccessMask() int32`

GetSpecialAccessMask returns the SpecialAccessMask field if non-nil, zero value otherwise.

### GetSpecialAccessMaskOk

`func (o *SmbPermission) GetSpecialAccessMaskOk() (*int32, bool)`

GetSpecialAccessMaskOk returns a tuple with the SpecialAccessMask field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpecialAccessMask

`func (o *SmbPermission) SetSpecialAccessMask(v int32)`

SetSpecialAccessMask sets SpecialAccessMask field to given value.

### HasSpecialAccessMask

`func (o *SmbPermission) HasSpecialAccessMask() bool`

HasSpecialAccessMask returns a boolean if a field has been set.

### SetSpecialAccessMaskNil

`func (o *SmbPermission) SetSpecialAccessMaskNil(b bool)`

 SetSpecialAccessMaskNil sets the value for SpecialAccessMask to be an explicit nil

### UnsetSpecialAccessMask
`func (o *SmbPermission) UnsetSpecialAccessMask()`

UnsetSpecialAccessMask ensures that no value is present for SpecialAccessMask, not even an explicit nil
### GetSpecialType

`func (o *SmbPermission) GetSpecialType() int32`

GetSpecialType returns the SpecialType field if non-nil, zero value otherwise.

### GetSpecialTypeOk

`func (o *SmbPermission) GetSpecialTypeOk() (*int32, bool)`

GetSpecialTypeOk returns a tuple with the SpecialType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpecialType

`func (o *SmbPermission) SetSpecialType(v int32)`

SetSpecialType sets SpecialType field to given value.

### HasSpecialType

`func (o *SmbPermission) HasSpecialType() bool`

HasSpecialType returns a boolean if a field has been set.

### SetSpecialTypeNil

`func (o *SmbPermission) SetSpecialTypeNil(b bool)`

 SetSpecialTypeNil sets the value for SpecialType to be an explicit nil

### UnsetSpecialType
`func (o *SmbPermission) UnsetSpecialType()`

UnsetSpecialType ensures that no value is present for SpecialType, not even an explicit nil
### GetType

`func (o *SmbPermission) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *SmbPermission) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *SmbPermission) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *SmbPermission) HasType() bool`

HasType returns a boolean if a field has been set.

### SetTypeNil

`func (o *SmbPermission) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *SmbPermission) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


