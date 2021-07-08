# RestoreFilesPreferences

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AlternateRestoreBaseDirectory** | Pointer to **NullableString** | This must be set to a directory path if restore_to_original_paths is false. All the files and directories restored will be restored under this location. | [optional] 
**ContinueOnError** | Pointer to **NullableBool** | Whether to continue with the copy in case of encountering an error. | [optional] 
**EncryptionEnabled** | Pointer to **NullableBool** | Whether to enable encryption for NFS and SMB restores. | [optional] 
**GenerateSshKeys** | Pointer to **NullableBool** | In case of GCP Linux restores, whether to generate ssh keys to connect to the customer&#39;s instance. | [optional] 
**OverrideOriginals** | Pointer to **NullableBool** | This is relevant only if restore_to_original_paths is true. If this is true, then already existing files will be overridden, otherwise new files will be skipped. | [optional] 
**PreserveAcls** | Pointer to **NullableBool** | Whether to preserve the ACLs of the original file. | [optional] 
**PreserveAttributes** | Pointer to **NullableBool** | Whether to preserve the original attributes. | [optional] 
**PreserveTimestamps** | Pointer to **NullableBool** | Whether to preserve the original time stamps. | [optional] 
**RestoreToOriginalPaths** | Pointer to **NullableBool** | If this is true, then files will be restored to original paths. | [optional] 
**SkipEstimation** | Pointer to **NullableBool** | Whether to skip the estimation step. | [optional] 

## Methods

### NewRestoreFilesPreferences

`func NewRestoreFilesPreferences() *RestoreFilesPreferences`

NewRestoreFilesPreferences instantiates a new RestoreFilesPreferences object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRestoreFilesPreferencesWithDefaults

`func NewRestoreFilesPreferencesWithDefaults() *RestoreFilesPreferences`

NewRestoreFilesPreferencesWithDefaults instantiates a new RestoreFilesPreferences object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAlternateRestoreBaseDirectory

`func (o *RestoreFilesPreferences) GetAlternateRestoreBaseDirectory() string`

GetAlternateRestoreBaseDirectory returns the AlternateRestoreBaseDirectory field if non-nil, zero value otherwise.

### GetAlternateRestoreBaseDirectoryOk

`func (o *RestoreFilesPreferences) GetAlternateRestoreBaseDirectoryOk() (*string, bool)`

GetAlternateRestoreBaseDirectoryOk returns a tuple with the AlternateRestoreBaseDirectory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlternateRestoreBaseDirectory

`func (o *RestoreFilesPreferences) SetAlternateRestoreBaseDirectory(v string)`

SetAlternateRestoreBaseDirectory sets AlternateRestoreBaseDirectory field to given value.

### HasAlternateRestoreBaseDirectory

`func (o *RestoreFilesPreferences) HasAlternateRestoreBaseDirectory() bool`

HasAlternateRestoreBaseDirectory returns a boolean if a field has been set.

### SetAlternateRestoreBaseDirectoryNil

`func (o *RestoreFilesPreferences) SetAlternateRestoreBaseDirectoryNil(b bool)`

 SetAlternateRestoreBaseDirectoryNil sets the value for AlternateRestoreBaseDirectory to be an explicit nil

### UnsetAlternateRestoreBaseDirectory
`func (o *RestoreFilesPreferences) UnsetAlternateRestoreBaseDirectory()`

UnsetAlternateRestoreBaseDirectory ensures that no value is present for AlternateRestoreBaseDirectory, not even an explicit nil
### GetContinueOnError

`func (o *RestoreFilesPreferences) GetContinueOnError() bool`

GetContinueOnError returns the ContinueOnError field if non-nil, zero value otherwise.

### GetContinueOnErrorOk

`func (o *RestoreFilesPreferences) GetContinueOnErrorOk() (*bool, bool)`

GetContinueOnErrorOk returns a tuple with the ContinueOnError field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContinueOnError

`func (o *RestoreFilesPreferences) SetContinueOnError(v bool)`

SetContinueOnError sets ContinueOnError field to given value.

### HasContinueOnError

`func (o *RestoreFilesPreferences) HasContinueOnError() bool`

HasContinueOnError returns a boolean if a field has been set.

### SetContinueOnErrorNil

`func (o *RestoreFilesPreferences) SetContinueOnErrorNil(b bool)`

 SetContinueOnErrorNil sets the value for ContinueOnError to be an explicit nil

### UnsetContinueOnError
`func (o *RestoreFilesPreferences) UnsetContinueOnError()`

UnsetContinueOnError ensures that no value is present for ContinueOnError, not even an explicit nil
### GetEncryptionEnabled

`func (o *RestoreFilesPreferences) GetEncryptionEnabled() bool`

GetEncryptionEnabled returns the EncryptionEnabled field if non-nil, zero value otherwise.

### GetEncryptionEnabledOk

`func (o *RestoreFilesPreferences) GetEncryptionEnabledOk() (*bool, bool)`

GetEncryptionEnabledOk returns a tuple with the EncryptionEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncryptionEnabled

`func (o *RestoreFilesPreferences) SetEncryptionEnabled(v bool)`

SetEncryptionEnabled sets EncryptionEnabled field to given value.

### HasEncryptionEnabled

`func (o *RestoreFilesPreferences) HasEncryptionEnabled() bool`

HasEncryptionEnabled returns a boolean if a field has been set.

### SetEncryptionEnabledNil

`func (o *RestoreFilesPreferences) SetEncryptionEnabledNil(b bool)`

 SetEncryptionEnabledNil sets the value for EncryptionEnabled to be an explicit nil

### UnsetEncryptionEnabled
`func (o *RestoreFilesPreferences) UnsetEncryptionEnabled()`

UnsetEncryptionEnabled ensures that no value is present for EncryptionEnabled, not even an explicit nil
### GetGenerateSshKeys

`func (o *RestoreFilesPreferences) GetGenerateSshKeys() bool`

GetGenerateSshKeys returns the GenerateSshKeys field if non-nil, zero value otherwise.

### GetGenerateSshKeysOk

`func (o *RestoreFilesPreferences) GetGenerateSshKeysOk() (*bool, bool)`

GetGenerateSshKeysOk returns a tuple with the GenerateSshKeys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGenerateSshKeys

`func (o *RestoreFilesPreferences) SetGenerateSshKeys(v bool)`

SetGenerateSshKeys sets GenerateSshKeys field to given value.

### HasGenerateSshKeys

`func (o *RestoreFilesPreferences) HasGenerateSshKeys() bool`

HasGenerateSshKeys returns a boolean if a field has been set.

### SetGenerateSshKeysNil

`func (o *RestoreFilesPreferences) SetGenerateSshKeysNil(b bool)`

 SetGenerateSshKeysNil sets the value for GenerateSshKeys to be an explicit nil

### UnsetGenerateSshKeys
`func (o *RestoreFilesPreferences) UnsetGenerateSshKeys()`

UnsetGenerateSshKeys ensures that no value is present for GenerateSshKeys, not even an explicit nil
### GetOverrideOriginals

`func (o *RestoreFilesPreferences) GetOverrideOriginals() bool`

GetOverrideOriginals returns the OverrideOriginals field if non-nil, zero value otherwise.

### GetOverrideOriginalsOk

`func (o *RestoreFilesPreferences) GetOverrideOriginalsOk() (*bool, bool)`

GetOverrideOriginalsOk returns a tuple with the OverrideOriginals field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverrideOriginals

`func (o *RestoreFilesPreferences) SetOverrideOriginals(v bool)`

SetOverrideOriginals sets OverrideOriginals field to given value.

### HasOverrideOriginals

`func (o *RestoreFilesPreferences) HasOverrideOriginals() bool`

HasOverrideOriginals returns a boolean if a field has been set.

### SetOverrideOriginalsNil

`func (o *RestoreFilesPreferences) SetOverrideOriginalsNil(b bool)`

 SetOverrideOriginalsNil sets the value for OverrideOriginals to be an explicit nil

### UnsetOverrideOriginals
`func (o *RestoreFilesPreferences) UnsetOverrideOriginals()`

UnsetOverrideOriginals ensures that no value is present for OverrideOriginals, not even an explicit nil
### GetPreserveAcls

`func (o *RestoreFilesPreferences) GetPreserveAcls() bool`

GetPreserveAcls returns the PreserveAcls field if non-nil, zero value otherwise.

### GetPreserveAclsOk

`func (o *RestoreFilesPreferences) GetPreserveAclsOk() (*bool, bool)`

GetPreserveAclsOk returns a tuple with the PreserveAcls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreserveAcls

`func (o *RestoreFilesPreferences) SetPreserveAcls(v bool)`

SetPreserveAcls sets PreserveAcls field to given value.

### HasPreserveAcls

`func (o *RestoreFilesPreferences) HasPreserveAcls() bool`

HasPreserveAcls returns a boolean if a field has been set.

### SetPreserveAclsNil

`func (o *RestoreFilesPreferences) SetPreserveAclsNil(b bool)`

 SetPreserveAclsNil sets the value for PreserveAcls to be an explicit nil

### UnsetPreserveAcls
`func (o *RestoreFilesPreferences) UnsetPreserveAcls()`

UnsetPreserveAcls ensures that no value is present for PreserveAcls, not even an explicit nil
### GetPreserveAttributes

`func (o *RestoreFilesPreferences) GetPreserveAttributes() bool`

GetPreserveAttributes returns the PreserveAttributes field if non-nil, zero value otherwise.

### GetPreserveAttributesOk

`func (o *RestoreFilesPreferences) GetPreserveAttributesOk() (*bool, bool)`

GetPreserveAttributesOk returns a tuple with the PreserveAttributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreserveAttributes

`func (o *RestoreFilesPreferences) SetPreserveAttributes(v bool)`

SetPreserveAttributes sets PreserveAttributes field to given value.

### HasPreserveAttributes

`func (o *RestoreFilesPreferences) HasPreserveAttributes() bool`

HasPreserveAttributes returns a boolean if a field has been set.

### SetPreserveAttributesNil

`func (o *RestoreFilesPreferences) SetPreserveAttributesNil(b bool)`

 SetPreserveAttributesNil sets the value for PreserveAttributes to be an explicit nil

### UnsetPreserveAttributes
`func (o *RestoreFilesPreferences) UnsetPreserveAttributes()`

UnsetPreserveAttributes ensures that no value is present for PreserveAttributes, not even an explicit nil
### GetPreserveTimestamps

`func (o *RestoreFilesPreferences) GetPreserveTimestamps() bool`

GetPreserveTimestamps returns the PreserveTimestamps field if non-nil, zero value otherwise.

### GetPreserveTimestampsOk

`func (o *RestoreFilesPreferences) GetPreserveTimestampsOk() (*bool, bool)`

GetPreserveTimestampsOk returns a tuple with the PreserveTimestamps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreserveTimestamps

`func (o *RestoreFilesPreferences) SetPreserveTimestamps(v bool)`

SetPreserveTimestamps sets PreserveTimestamps field to given value.

### HasPreserveTimestamps

`func (o *RestoreFilesPreferences) HasPreserveTimestamps() bool`

HasPreserveTimestamps returns a boolean if a field has been set.

### SetPreserveTimestampsNil

`func (o *RestoreFilesPreferences) SetPreserveTimestampsNil(b bool)`

 SetPreserveTimestampsNil sets the value for PreserveTimestamps to be an explicit nil

### UnsetPreserveTimestamps
`func (o *RestoreFilesPreferences) UnsetPreserveTimestamps()`

UnsetPreserveTimestamps ensures that no value is present for PreserveTimestamps, not even an explicit nil
### GetRestoreToOriginalPaths

`func (o *RestoreFilesPreferences) GetRestoreToOriginalPaths() bool`

GetRestoreToOriginalPaths returns the RestoreToOriginalPaths field if non-nil, zero value otherwise.

### GetRestoreToOriginalPathsOk

`func (o *RestoreFilesPreferences) GetRestoreToOriginalPathsOk() (*bool, bool)`

GetRestoreToOriginalPathsOk returns a tuple with the RestoreToOriginalPaths field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestoreToOriginalPaths

`func (o *RestoreFilesPreferences) SetRestoreToOriginalPaths(v bool)`

SetRestoreToOriginalPaths sets RestoreToOriginalPaths field to given value.

### HasRestoreToOriginalPaths

`func (o *RestoreFilesPreferences) HasRestoreToOriginalPaths() bool`

HasRestoreToOriginalPaths returns a boolean if a field has been set.

### SetRestoreToOriginalPathsNil

`func (o *RestoreFilesPreferences) SetRestoreToOriginalPathsNil(b bool)`

 SetRestoreToOriginalPathsNil sets the value for RestoreToOriginalPaths to be an explicit nil

### UnsetRestoreToOriginalPaths
`func (o *RestoreFilesPreferences) UnsetRestoreToOriginalPaths()`

UnsetRestoreToOriginalPaths ensures that no value is present for RestoreToOriginalPaths, not even an explicit nil
### GetSkipEstimation

`func (o *RestoreFilesPreferences) GetSkipEstimation() bool`

GetSkipEstimation returns the SkipEstimation field if non-nil, zero value otherwise.

### GetSkipEstimationOk

`func (o *RestoreFilesPreferences) GetSkipEstimationOk() (*bool, bool)`

GetSkipEstimationOk returns a tuple with the SkipEstimation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkipEstimation

`func (o *RestoreFilesPreferences) SetSkipEstimation(v bool)`

SetSkipEstimation sets SkipEstimation field to given value.

### HasSkipEstimation

`func (o *RestoreFilesPreferences) HasSkipEstimation() bool`

HasSkipEstimation returns a boolean if a field has been set.

### SetSkipEstimationNil

`func (o *RestoreFilesPreferences) SetSkipEstimationNil(b bool)`

 SetSkipEstimationNil sets the value for SkipEstimation to be an explicit nil

### UnsetSkipEstimation
`func (o *RestoreFilesPreferences) UnsetSkipEstimation()`

UnsetSkipEstimation ensures that no value is present for SkipEstimation, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


