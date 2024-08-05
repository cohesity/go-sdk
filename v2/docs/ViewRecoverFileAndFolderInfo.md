# ViewRecoverFileAndFolderInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AbsolutePath** | **NullableString** | Specifies the absolute path to the file or folder. | 
**DestinationDir** | Pointer to **NullableString** | Specifies the destination directory where the file/directory was copied. | [optional] [readonly] 
**IsDirectory** | Pointer to **NullableBool** | Specifies whether this is a directory or not. | [optional] 
**IsViewFileRecovery** | Pointer to **NullableBool** | Specify if the recovery is of type view file/folder. | [optional] 
**Messages** | Pointer to **[]string** | Specify error messages about the file during recovery. | [optional] [readonly] 
**Status** | Pointer to **NullableString** | Specifies the recovery status for this file or folder. | [optional] [readonly] 
**InodeId** | Pointer to **NullableInt64** | Specifies the source inode number of the file being recovered. | [optional] 

## Methods

### NewViewRecoverFileAndFolderInfo

`func NewViewRecoverFileAndFolderInfo(absolutePath NullableString, ) *ViewRecoverFileAndFolderInfo`

NewViewRecoverFileAndFolderInfo instantiates a new ViewRecoverFileAndFolderInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewViewRecoverFileAndFolderInfoWithDefaults

`func NewViewRecoverFileAndFolderInfoWithDefaults() *ViewRecoverFileAndFolderInfo`

NewViewRecoverFileAndFolderInfoWithDefaults instantiates a new ViewRecoverFileAndFolderInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAbsolutePath

`func (o *ViewRecoverFileAndFolderInfo) GetAbsolutePath() string`

GetAbsolutePath returns the AbsolutePath field if non-nil, zero value otherwise.

### GetAbsolutePathOk

`func (o *ViewRecoverFileAndFolderInfo) GetAbsolutePathOk() (*string, bool)`

GetAbsolutePathOk returns a tuple with the AbsolutePath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAbsolutePath

`func (o *ViewRecoverFileAndFolderInfo) SetAbsolutePath(v string)`

SetAbsolutePath sets AbsolutePath field to given value.


### SetAbsolutePathNil

`func (o *ViewRecoverFileAndFolderInfo) SetAbsolutePathNil(b bool)`

 SetAbsolutePathNil sets the value for AbsolutePath to be an explicit nil

### UnsetAbsolutePath
`func (o *ViewRecoverFileAndFolderInfo) UnsetAbsolutePath()`

UnsetAbsolutePath ensures that no value is present for AbsolutePath, not even an explicit nil
### GetDestinationDir

`func (o *ViewRecoverFileAndFolderInfo) GetDestinationDir() string`

GetDestinationDir returns the DestinationDir field if non-nil, zero value otherwise.

### GetDestinationDirOk

`func (o *ViewRecoverFileAndFolderInfo) GetDestinationDirOk() (*string, bool)`

GetDestinationDirOk returns a tuple with the DestinationDir field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationDir

`func (o *ViewRecoverFileAndFolderInfo) SetDestinationDir(v string)`

SetDestinationDir sets DestinationDir field to given value.

### HasDestinationDir

`func (o *ViewRecoverFileAndFolderInfo) HasDestinationDir() bool`

HasDestinationDir returns a boolean if a field has been set.

### SetDestinationDirNil

`func (o *ViewRecoverFileAndFolderInfo) SetDestinationDirNil(b bool)`

 SetDestinationDirNil sets the value for DestinationDir to be an explicit nil

### UnsetDestinationDir
`func (o *ViewRecoverFileAndFolderInfo) UnsetDestinationDir()`

UnsetDestinationDir ensures that no value is present for DestinationDir, not even an explicit nil
### GetIsDirectory

`func (o *ViewRecoverFileAndFolderInfo) GetIsDirectory() bool`

GetIsDirectory returns the IsDirectory field if non-nil, zero value otherwise.

### GetIsDirectoryOk

`func (o *ViewRecoverFileAndFolderInfo) GetIsDirectoryOk() (*bool, bool)`

GetIsDirectoryOk returns a tuple with the IsDirectory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDirectory

`func (o *ViewRecoverFileAndFolderInfo) SetIsDirectory(v bool)`

SetIsDirectory sets IsDirectory field to given value.

### HasIsDirectory

`func (o *ViewRecoverFileAndFolderInfo) HasIsDirectory() bool`

HasIsDirectory returns a boolean if a field has been set.

### SetIsDirectoryNil

`func (o *ViewRecoverFileAndFolderInfo) SetIsDirectoryNil(b bool)`

 SetIsDirectoryNil sets the value for IsDirectory to be an explicit nil

### UnsetIsDirectory
`func (o *ViewRecoverFileAndFolderInfo) UnsetIsDirectory()`

UnsetIsDirectory ensures that no value is present for IsDirectory, not even an explicit nil
### GetIsViewFileRecovery

`func (o *ViewRecoverFileAndFolderInfo) GetIsViewFileRecovery() bool`

GetIsViewFileRecovery returns the IsViewFileRecovery field if non-nil, zero value otherwise.

### GetIsViewFileRecoveryOk

`func (o *ViewRecoverFileAndFolderInfo) GetIsViewFileRecoveryOk() (*bool, bool)`

GetIsViewFileRecoveryOk returns a tuple with the IsViewFileRecovery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsViewFileRecovery

`func (o *ViewRecoverFileAndFolderInfo) SetIsViewFileRecovery(v bool)`

SetIsViewFileRecovery sets IsViewFileRecovery field to given value.

### HasIsViewFileRecovery

`func (o *ViewRecoverFileAndFolderInfo) HasIsViewFileRecovery() bool`

HasIsViewFileRecovery returns a boolean if a field has been set.

### SetIsViewFileRecoveryNil

`func (o *ViewRecoverFileAndFolderInfo) SetIsViewFileRecoveryNil(b bool)`

 SetIsViewFileRecoveryNil sets the value for IsViewFileRecovery to be an explicit nil

### UnsetIsViewFileRecovery
`func (o *ViewRecoverFileAndFolderInfo) UnsetIsViewFileRecovery()`

UnsetIsViewFileRecovery ensures that no value is present for IsViewFileRecovery, not even an explicit nil
### GetMessages

`func (o *ViewRecoverFileAndFolderInfo) GetMessages() []string`

GetMessages returns the Messages field if non-nil, zero value otherwise.

### GetMessagesOk

`func (o *ViewRecoverFileAndFolderInfo) GetMessagesOk() (*[]string, bool)`

GetMessagesOk returns a tuple with the Messages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessages

`func (o *ViewRecoverFileAndFolderInfo) SetMessages(v []string)`

SetMessages sets Messages field to given value.

### HasMessages

`func (o *ViewRecoverFileAndFolderInfo) HasMessages() bool`

HasMessages returns a boolean if a field has been set.

### SetMessagesNil

`func (o *ViewRecoverFileAndFolderInfo) SetMessagesNil(b bool)`

 SetMessagesNil sets the value for Messages to be an explicit nil

### UnsetMessages
`func (o *ViewRecoverFileAndFolderInfo) UnsetMessages()`

UnsetMessages ensures that no value is present for Messages, not even an explicit nil
### GetStatus

`func (o *ViewRecoverFileAndFolderInfo) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ViewRecoverFileAndFolderInfo) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ViewRecoverFileAndFolderInfo) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ViewRecoverFileAndFolderInfo) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### SetStatusNil

`func (o *ViewRecoverFileAndFolderInfo) SetStatusNil(b bool)`

 SetStatusNil sets the value for Status to be an explicit nil

### UnsetStatus
`func (o *ViewRecoverFileAndFolderInfo) UnsetStatus()`

UnsetStatus ensures that no value is present for Status, not even an explicit nil
### GetInodeId

`func (o *ViewRecoverFileAndFolderInfo) GetInodeId() int64`

GetInodeId returns the InodeId field if non-nil, zero value otherwise.

### GetInodeIdOk

`func (o *ViewRecoverFileAndFolderInfo) GetInodeIdOk() (*int64, bool)`

GetInodeIdOk returns a tuple with the InodeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInodeId

`func (o *ViewRecoverFileAndFolderInfo) SetInodeId(v int64)`

SetInodeId sets InodeId field to given value.

### HasInodeId

`func (o *ViewRecoverFileAndFolderInfo) HasInodeId() bool`

HasInodeId returns a boolean if a field has been set.

### SetInodeIdNil

`func (o *ViewRecoverFileAndFolderInfo) SetInodeIdNil(b bool)`

 SetInodeIdNil sets the value for InodeId to be an explicit nil

### UnsetInodeId
`func (o *ViewRecoverFileAndFolderInfo) UnsetInodeId()`

UnsetInodeId ensures that no value is present for InodeId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


