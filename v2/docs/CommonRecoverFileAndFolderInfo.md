# CommonRecoverFileAndFolderInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AbsolutePath** | **NullableString** | Specifies the absolute path to the file or folder. | 
**DestinationDir** | Pointer to **NullableString** | Specifies the destination directory where the file/directory was copied. | [optional] [readonly] 
**IsDirectory** | Pointer to **NullableBool** | Specifies whether this is a directory or not. | [optional] 
**IsViewFileRecovery** | Pointer to **NullableBool** | Specify if the recovery is of type view file/folder. | [optional] 
**Messages** | Pointer to **[]string** | Specify error messages about the file during recovery. | [optional] [readonly] 
**Status** | Pointer to **NullableString** | Specifies the recovery status for this file or folder. | [optional] [readonly] 

## Methods

### NewCommonRecoverFileAndFolderInfo

`func NewCommonRecoverFileAndFolderInfo(absolutePath NullableString, ) *CommonRecoverFileAndFolderInfo`

NewCommonRecoverFileAndFolderInfo instantiates a new CommonRecoverFileAndFolderInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCommonRecoverFileAndFolderInfoWithDefaults

`func NewCommonRecoverFileAndFolderInfoWithDefaults() *CommonRecoverFileAndFolderInfo`

NewCommonRecoverFileAndFolderInfoWithDefaults instantiates a new CommonRecoverFileAndFolderInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAbsolutePath

`func (o *CommonRecoverFileAndFolderInfo) GetAbsolutePath() string`

GetAbsolutePath returns the AbsolutePath field if non-nil, zero value otherwise.

### GetAbsolutePathOk

`func (o *CommonRecoverFileAndFolderInfo) GetAbsolutePathOk() (*string, bool)`

GetAbsolutePathOk returns a tuple with the AbsolutePath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAbsolutePath

`func (o *CommonRecoverFileAndFolderInfo) SetAbsolutePath(v string)`

SetAbsolutePath sets AbsolutePath field to given value.


### SetAbsolutePathNil

`func (o *CommonRecoverFileAndFolderInfo) SetAbsolutePathNil(b bool)`

 SetAbsolutePathNil sets the value for AbsolutePath to be an explicit nil

### UnsetAbsolutePath
`func (o *CommonRecoverFileAndFolderInfo) UnsetAbsolutePath()`

UnsetAbsolutePath ensures that no value is present for AbsolutePath, not even an explicit nil
### GetDestinationDir

`func (o *CommonRecoverFileAndFolderInfo) GetDestinationDir() string`

GetDestinationDir returns the DestinationDir field if non-nil, zero value otherwise.

### GetDestinationDirOk

`func (o *CommonRecoverFileAndFolderInfo) GetDestinationDirOk() (*string, bool)`

GetDestinationDirOk returns a tuple with the DestinationDir field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationDir

`func (o *CommonRecoverFileAndFolderInfo) SetDestinationDir(v string)`

SetDestinationDir sets DestinationDir field to given value.

### HasDestinationDir

`func (o *CommonRecoverFileAndFolderInfo) HasDestinationDir() bool`

HasDestinationDir returns a boolean if a field has been set.

### SetDestinationDirNil

`func (o *CommonRecoverFileAndFolderInfo) SetDestinationDirNil(b bool)`

 SetDestinationDirNil sets the value for DestinationDir to be an explicit nil

### UnsetDestinationDir
`func (o *CommonRecoverFileAndFolderInfo) UnsetDestinationDir()`

UnsetDestinationDir ensures that no value is present for DestinationDir, not even an explicit nil
### GetIsDirectory

`func (o *CommonRecoverFileAndFolderInfo) GetIsDirectory() bool`

GetIsDirectory returns the IsDirectory field if non-nil, zero value otherwise.

### GetIsDirectoryOk

`func (o *CommonRecoverFileAndFolderInfo) GetIsDirectoryOk() (*bool, bool)`

GetIsDirectoryOk returns a tuple with the IsDirectory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDirectory

`func (o *CommonRecoverFileAndFolderInfo) SetIsDirectory(v bool)`

SetIsDirectory sets IsDirectory field to given value.

### HasIsDirectory

`func (o *CommonRecoverFileAndFolderInfo) HasIsDirectory() bool`

HasIsDirectory returns a boolean if a field has been set.

### SetIsDirectoryNil

`func (o *CommonRecoverFileAndFolderInfo) SetIsDirectoryNil(b bool)`

 SetIsDirectoryNil sets the value for IsDirectory to be an explicit nil

### UnsetIsDirectory
`func (o *CommonRecoverFileAndFolderInfo) UnsetIsDirectory()`

UnsetIsDirectory ensures that no value is present for IsDirectory, not even an explicit nil
### GetIsViewFileRecovery

`func (o *CommonRecoverFileAndFolderInfo) GetIsViewFileRecovery() bool`

GetIsViewFileRecovery returns the IsViewFileRecovery field if non-nil, zero value otherwise.

### GetIsViewFileRecoveryOk

`func (o *CommonRecoverFileAndFolderInfo) GetIsViewFileRecoveryOk() (*bool, bool)`

GetIsViewFileRecoveryOk returns a tuple with the IsViewFileRecovery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsViewFileRecovery

`func (o *CommonRecoverFileAndFolderInfo) SetIsViewFileRecovery(v bool)`

SetIsViewFileRecovery sets IsViewFileRecovery field to given value.

### HasIsViewFileRecovery

`func (o *CommonRecoverFileAndFolderInfo) HasIsViewFileRecovery() bool`

HasIsViewFileRecovery returns a boolean if a field has been set.

### SetIsViewFileRecoveryNil

`func (o *CommonRecoverFileAndFolderInfo) SetIsViewFileRecoveryNil(b bool)`

 SetIsViewFileRecoveryNil sets the value for IsViewFileRecovery to be an explicit nil

### UnsetIsViewFileRecovery
`func (o *CommonRecoverFileAndFolderInfo) UnsetIsViewFileRecovery()`

UnsetIsViewFileRecovery ensures that no value is present for IsViewFileRecovery, not even an explicit nil
### GetMessages

`func (o *CommonRecoverFileAndFolderInfo) GetMessages() []string`

GetMessages returns the Messages field if non-nil, zero value otherwise.

### GetMessagesOk

`func (o *CommonRecoverFileAndFolderInfo) GetMessagesOk() (*[]string, bool)`

GetMessagesOk returns a tuple with the Messages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessages

`func (o *CommonRecoverFileAndFolderInfo) SetMessages(v []string)`

SetMessages sets Messages field to given value.

### HasMessages

`func (o *CommonRecoverFileAndFolderInfo) HasMessages() bool`

HasMessages returns a boolean if a field has been set.

### SetMessagesNil

`func (o *CommonRecoverFileAndFolderInfo) SetMessagesNil(b bool)`

 SetMessagesNil sets the value for Messages to be an explicit nil

### UnsetMessages
`func (o *CommonRecoverFileAndFolderInfo) UnsetMessages()`

UnsetMessages ensures that no value is present for Messages, not even an explicit nil
### GetStatus

`func (o *CommonRecoverFileAndFolderInfo) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *CommonRecoverFileAndFolderInfo) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *CommonRecoverFileAndFolderInfo) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *CommonRecoverFileAndFolderInfo) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### SetStatusNil

`func (o *CommonRecoverFileAndFolderInfo) SetStatusNil(b bool)`

 SetStatusNil sets the value for Status to be an explicit nil

### UnsetStatus
`func (o *CommonRecoverFileAndFolderInfo) UnsetStatus()`

UnsetStatus ensures that no value is present for Status, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


