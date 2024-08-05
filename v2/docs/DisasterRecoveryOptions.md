# DisasterRecoveryOptions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CleanupOriginalDbFiles** | Pointer to **NullableBool** | Specifies whether to cleanup the original database files or to do precheck to ensure no conflicting files exists. Recovery will fail if there are any conflicting files. | [optional] 
**IsDisasterRecovery** | Pointer to **NullableBool** | Specifies whether the recovery is of type Disaster Recovery. | [optional] 
**RenameDatabaseAsmDirectory** | Pointer to **NullableBool** | Whether to rename the database ASM directory. If false, the adapter will leave the database files and continue with clone and migration of datafiles. This might cause extra files left behind on the Oracle host from the existing database instance. | [optional] 

## Methods

### NewDisasterRecoveryOptions

`func NewDisasterRecoveryOptions() *DisasterRecoveryOptions`

NewDisasterRecoveryOptions instantiates a new DisasterRecoveryOptions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDisasterRecoveryOptionsWithDefaults

`func NewDisasterRecoveryOptionsWithDefaults() *DisasterRecoveryOptions`

NewDisasterRecoveryOptionsWithDefaults instantiates a new DisasterRecoveryOptions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCleanupOriginalDbFiles

`func (o *DisasterRecoveryOptions) GetCleanupOriginalDbFiles() bool`

GetCleanupOriginalDbFiles returns the CleanupOriginalDbFiles field if non-nil, zero value otherwise.

### GetCleanupOriginalDbFilesOk

`func (o *DisasterRecoveryOptions) GetCleanupOriginalDbFilesOk() (*bool, bool)`

GetCleanupOriginalDbFilesOk returns a tuple with the CleanupOriginalDbFiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCleanupOriginalDbFiles

`func (o *DisasterRecoveryOptions) SetCleanupOriginalDbFiles(v bool)`

SetCleanupOriginalDbFiles sets CleanupOriginalDbFiles field to given value.

### HasCleanupOriginalDbFiles

`func (o *DisasterRecoveryOptions) HasCleanupOriginalDbFiles() bool`

HasCleanupOriginalDbFiles returns a boolean if a field has been set.

### SetCleanupOriginalDbFilesNil

`func (o *DisasterRecoveryOptions) SetCleanupOriginalDbFilesNil(b bool)`

 SetCleanupOriginalDbFilesNil sets the value for CleanupOriginalDbFiles to be an explicit nil

### UnsetCleanupOriginalDbFiles
`func (o *DisasterRecoveryOptions) UnsetCleanupOriginalDbFiles()`

UnsetCleanupOriginalDbFiles ensures that no value is present for CleanupOriginalDbFiles, not even an explicit nil
### GetIsDisasterRecovery

`func (o *DisasterRecoveryOptions) GetIsDisasterRecovery() bool`

GetIsDisasterRecovery returns the IsDisasterRecovery field if non-nil, zero value otherwise.

### GetIsDisasterRecoveryOk

`func (o *DisasterRecoveryOptions) GetIsDisasterRecoveryOk() (*bool, bool)`

GetIsDisasterRecoveryOk returns a tuple with the IsDisasterRecovery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDisasterRecovery

`func (o *DisasterRecoveryOptions) SetIsDisasterRecovery(v bool)`

SetIsDisasterRecovery sets IsDisasterRecovery field to given value.

### HasIsDisasterRecovery

`func (o *DisasterRecoveryOptions) HasIsDisasterRecovery() bool`

HasIsDisasterRecovery returns a boolean if a field has been set.

### SetIsDisasterRecoveryNil

`func (o *DisasterRecoveryOptions) SetIsDisasterRecoveryNil(b bool)`

 SetIsDisasterRecoveryNil sets the value for IsDisasterRecovery to be an explicit nil

### UnsetIsDisasterRecovery
`func (o *DisasterRecoveryOptions) UnsetIsDisasterRecovery()`

UnsetIsDisasterRecovery ensures that no value is present for IsDisasterRecovery, not even an explicit nil
### GetRenameDatabaseAsmDirectory

`func (o *DisasterRecoveryOptions) GetRenameDatabaseAsmDirectory() bool`

GetRenameDatabaseAsmDirectory returns the RenameDatabaseAsmDirectory field if non-nil, zero value otherwise.

### GetRenameDatabaseAsmDirectoryOk

`func (o *DisasterRecoveryOptions) GetRenameDatabaseAsmDirectoryOk() (*bool, bool)`

GetRenameDatabaseAsmDirectoryOk returns a tuple with the RenameDatabaseAsmDirectory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRenameDatabaseAsmDirectory

`func (o *DisasterRecoveryOptions) SetRenameDatabaseAsmDirectory(v bool)`

SetRenameDatabaseAsmDirectory sets RenameDatabaseAsmDirectory field to given value.

### HasRenameDatabaseAsmDirectory

`func (o *DisasterRecoveryOptions) HasRenameDatabaseAsmDirectory() bool`

HasRenameDatabaseAsmDirectory returns a boolean if a field has been set.

### SetRenameDatabaseAsmDirectoryNil

`func (o *DisasterRecoveryOptions) SetRenameDatabaseAsmDirectoryNil(b bool)`

 SetRenameDatabaseAsmDirectoryNil sets the value for RenameDatabaseAsmDirectory to be an explicit nil

### UnsetRenameDatabaseAsmDirectory
`func (o *DisasterRecoveryOptions) UnsetRenameDatabaseAsmDirectory()`

UnsetRenameDatabaseAsmDirectory ensures that no value is present for RenameDatabaseAsmDirectory, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


