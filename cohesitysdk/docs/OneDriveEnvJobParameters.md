# OneDriveEnvJobParameters

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FilePathFilter** | Pointer to [**FilePathFilter**](FilePathFilter.md) |  | [optional] 
**ShouldBackupOnedrive** | Pointer to **NullableBool** | Specifies whether OneDrive(s) of each Office365 Users/Groups within the job, should be backed up or not. | [optional] 

## Methods

### NewOneDriveEnvJobParameters

`func NewOneDriveEnvJobParameters() *OneDriveEnvJobParameters`

NewOneDriveEnvJobParameters instantiates a new OneDriveEnvJobParameters object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOneDriveEnvJobParametersWithDefaults

`func NewOneDriveEnvJobParametersWithDefaults() *OneDriveEnvJobParameters`

NewOneDriveEnvJobParametersWithDefaults instantiates a new OneDriveEnvJobParameters object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFilePathFilter

`func (o *OneDriveEnvJobParameters) GetFilePathFilter() FilePathFilter`

GetFilePathFilter returns the FilePathFilter field if non-nil, zero value otherwise.

### GetFilePathFilterOk

`func (o *OneDriveEnvJobParameters) GetFilePathFilterOk() (*FilePathFilter, bool)`

GetFilePathFilterOk returns a tuple with the FilePathFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilePathFilter

`func (o *OneDriveEnvJobParameters) SetFilePathFilter(v FilePathFilter)`

SetFilePathFilter sets FilePathFilter field to given value.

### HasFilePathFilter

`func (o *OneDriveEnvJobParameters) HasFilePathFilter() bool`

HasFilePathFilter returns a boolean if a field has been set.

### GetShouldBackupOnedrive

`func (o *OneDriveEnvJobParameters) GetShouldBackupOnedrive() bool`

GetShouldBackupOnedrive returns the ShouldBackupOnedrive field if non-nil, zero value otherwise.

### GetShouldBackupOnedriveOk

`func (o *OneDriveEnvJobParameters) GetShouldBackupOnedriveOk() (*bool, bool)`

GetShouldBackupOnedriveOk returns a tuple with the ShouldBackupOnedrive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShouldBackupOnedrive

`func (o *OneDriveEnvJobParameters) SetShouldBackupOnedrive(v bool)`

SetShouldBackupOnedrive sets ShouldBackupOnedrive field to given value.

### HasShouldBackupOnedrive

`func (o *OneDriveEnvJobParameters) HasShouldBackupOnedrive() bool`

HasShouldBackupOnedrive returns a boolean if a field has been set.

### SetShouldBackupOnedriveNil

`func (o *OneDriveEnvJobParameters) SetShouldBackupOnedriveNil(b bool)`

 SetShouldBackupOnedriveNil sets the value for ShouldBackupOnedrive to be an explicit nil

### UnsetShouldBackupOnedrive
`func (o *OneDriveEnvJobParameters) UnsetShouldBackupOnedrive()`

UnsetShouldBackupOnedrive ensures that no value is present for ShouldBackupOnedrive, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


