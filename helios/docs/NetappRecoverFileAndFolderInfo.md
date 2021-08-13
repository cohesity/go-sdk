# NetappRecoverFileAndFolderInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AbsolutePath** | **NullableString** | Specifies the absolute path to the file or folder. | 
**DestinationDir** | Pointer to **NullableString** | Specifies the destination directory where the file/directory was copied. | [optional] [readonly] 
**IsDirectory** | Pointer to **NullableBool** | Specifies whether this is a directory or not. | [optional] 
**Status** | Pointer to **NullableString** | Specifies the recovery status for this file or folder. | [optional] [readonly] 
**Messages** | Pointer to **[]string** | Specify error messages about the file during recovery. | [optional] [readonly] 
**InodeId** | Pointer to **NullableInt64** | Specifies the source inode number of the file being recovered. | [optional] 

## Methods

### NewNetappRecoverFileAndFolderInfo

`func NewNetappRecoverFileAndFolderInfo(absolutePath NullableString, ) *NetappRecoverFileAndFolderInfo`

NewNetappRecoverFileAndFolderInfo instantiates a new NetappRecoverFileAndFolderInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetappRecoverFileAndFolderInfoWithDefaults

`func NewNetappRecoverFileAndFolderInfoWithDefaults() *NetappRecoverFileAndFolderInfo`

NewNetappRecoverFileAndFolderInfoWithDefaults instantiates a new NetappRecoverFileAndFolderInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAbsolutePath

`func (o *NetappRecoverFileAndFolderInfo) GetAbsolutePath() string`

GetAbsolutePath returns the AbsolutePath field if non-nil, zero value otherwise.

### GetAbsolutePathOk

`func (o *NetappRecoverFileAndFolderInfo) GetAbsolutePathOk() (*string, bool)`

GetAbsolutePathOk returns a tuple with the AbsolutePath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAbsolutePath

`func (o *NetappRecoverFileAndFolderInfo) SetAbsolutePath(v string)`

SetAbsolutePath sets AbsolutePath field to given value.


### SetAbsolutePathNil

`func (o *NetappRecoverFileAndFolderInfo) SetAbsolutePathNil(b bool)`

 SetAbsolutePathNil sets the value for AbsolutePath to be an explicit nil

### UnsetAbsolutePath
`func (o *NetappRecoverFileAndFolderInfo) UnsetAbsolutePath()`

UnsetAbsolutePath ensures that no value is present for AbsolutePath, not even an explicit nil
### GetDestinationDir

`func (o *NetappRecoverFileAndFolderInfo) GetDestinationDir() string`

GetDestinationDir returns the DestinationDir field if non-nil, zero value otherwise.

### GetDestinationDirOk

`func (o *NetappRecoverFileAndFolderInfo) GetDestinationDirOk() (*string, bool)`

GetDestinationDirOk returns a tuple with the DestinationDir field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationDir

`func (o *NetappRecoverFileAndFolderInfo) SetDestinationDir(v string)`

SetDestinationDir sets DestinationDir field to given value.

### HasDestinationDir

`func (o *NetappRecoverFileAndFolderInfo) HasDestinationDir() bool`

HasDestinationDir returns a boolean if a field has been set.

### SetDestinationDirNil

`func (o *NetappRecoverFileAndFolderInfo) SetDestinationDirNil(b bool)`

 SetDestinationDirNil sets the value for DestinationDir to be an explicit nil

### UnsetDestinationDir
`func (o *NetappRecoverFileAndFolderInfo) UnsetDestinationDir()`

UnsetDestinationDir ensures that no value is present for DestinationDir, not even an explicit nil
### GetIsDirectory

`func (o *NetappRecoverFileAndFolderInfo) GetIsDirectory() bool`

GetIsDirectory returns the IsDirectory field if non-nil, zero value otherwise.

### GetIsDirectoryOk

`func (o *NetappRecoverFileAndFolderInfo) GetIsDirectoryOk() (*bool, bool)`

GetIsDirectoryOk returns a tuple with the IsDirectory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDirectory

`func (o *NetappRecoverFileAndFolderInfo) SetIsDirectory(v bool)`

SetIsDirectory sets IsDirectory field to given value.

### HasIsDirectory

`func (o *NetappRecoverFileAndFolderInfo) HasIsDirectory() bool`

HasIsDirectory returns a boolean if a field has been set.

### SetIsDirectoryNil

`func (o *NetappRecoverFileAndFolderInfo) SetIsDirectoryNil(b bool)`

 SetIsDirectoryNil sets the value for IsDirectory to be an explicit nil

### UnsetIsDirectory
`func (o *NetappRecoverFileAndFolderInfo) UnsetIsDirectory()`

UnsetIsDirectory ensures that no value is present for IsDirectory, not even an explicit nil
### GetStatus

`func (o *NetappRecoverFileAndFolderInfo) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *NetappRecoverFileAndFolderInfo) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *NetappRecoverFileAndFolderInfo) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *NetappRecoverFileAndFolderInfo) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### SetStatusNil

`func (o *NetappRecoverFileAndFolderInfo) SetStatusNil(b bool)`

 SetStatusNil sets the value for Status to be an explicit nil

### UnsetStatus
`func (o *NetappRecoverFileAndFolderInfo) UnsetStatus()`

UnsetStatus ensures that no value is present for Status, not even an explicit nil
### GetMessages

`func (o *NetappRecoverFileAndFolderInfo) GetMessages() []string`

GetMessages returns the Messages field if non-nil, zero value otherwise.

### GetMessagesOk

`func (o *NetappRecoverFileAndFolderInfo) GetMessagesOk() (*[]string, bool)`

GetMessagesOk returns a tuple with the Messages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessages

`func (o *NetappRecoverFileAndFolderInfo) SetMessages(v []string)`

SetMessages sets Messages field to given value.

### HasMessages

`func (o *NetappRecoverFileAndFolderInfo) HasMessages() bool`

HasMessages returns a boolean if a field has been set.

### SetMessagesNil

`func (o *NetappRecoverFileAndFolderInfo) SetMessagesNil(b bool)`

 SetMessagesNil sets the value for Messages to be an explicit nil

### UnsetMessages
`func (o *NetappRecoverFileAndFolderInfo) UnsetMessages()`

UnsetMessages ensures that no value is present for Messages, not even an explicit nil
### GetInodeId

`func (o *NetappRecoverFileAndFolderInfo) GetInodeId() int64`

GetInodeId returns the InodeId field if non-nil, zero value otherwise.

### GetInodeIdOk

`func (o *NetappRecoverFileAndFolderInfo) GetInodeIdOk() (*int64, bool)`

GetInodeIdOk returns a tuple with the InodeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInodeId

`func (o *NetappRecoverFileAndFolderInfo) SetInodeId(v int64)`

SetInodeId sets InodeId field to given value.

### HasInodeId

`func (o *NetappRecoverFileAndFolderInfo) HasInodeId() bool`

HasInodeId returns a boolean if a field has been set.

### SetInodeIdNil

`func (o *NetappRecoverFileAndFolderInfo) SetInodeIdNil(b bool)`

 SetInodeIdNil sets the value for InodeId to be an explicit nil

### UnsetInodeId
`func (o *NetappRecoverFileAndFolderInfo) UnsetInodeId()`

UnsetInodeId ensures that no value is present for InodeId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


