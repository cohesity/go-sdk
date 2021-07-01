# DataMigrationJobParameters

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ColdFileWindow** | Pointer to **NullableInt64** | Identifies the cold files in the NAS source. Files that haven&#39;t been accessed/modified in the last cold_file_window are migrated. | [optional] 
**DeleteOrphanData** | Pointer to **NullableBool** | Delete migrated data if no symlink at source is pointing to it. | [optional] 
**FilePathFilter** | Pointer to [**FilePathFilter**](FilePathFilter.md) |  | [optional] 
**FileSelectionPolicy** | Pointer to **NullableString** | Specifies policy to select a file to migrate based on its creation, last access or modification time. eg. A file can be selected to migrate if it has not been accessed/modified in the ColdFileWindow. enum: kOlderThan, kLastAccessed, kLastModified. Specifies policy for file selection in data migration jobs based on time. &#39;kOlderThan&#39;: Migrate the files that are older than cold file window. &#39;kLastAccessed&#39;: Migrate the files that are not accessed in cold file window. &#39;kLastModified&#39;: Migrate the files that have not been modified in cold file window. | [optional] 
**FileSizeBytes** | Pointer to **NullableInt64** | Gives the size criteria to be used for selecting the files to be migrated in bytes. The cold files that are equal and greater than this size are migrated. | [optional] 
**FileSizePolicy** | Pointer to **NullableString** | Specifies policy to select a file to migrate based on its size. eg. A file can be selected to migrate if its size is greater than or smaller than the FileSizeBytes. enum: kGreaterThan, kSmallerThan. Specifies policy for file selection in data migration jobs based on file size. &#39;kGreaterThan&#39;: Migrate the files whose size are greater than specified file size. &#39;kSmallerThan&#39;: Migrate the files whose size are smaller than specified file size. | [optional] 
**MigrateWithoutStub** | Pointer to **NullableBool** | Specifies if data is to be migrated without stub. | [optional] 
**NfsMountPath** | Pointer to **NullableString** | Mount path where the target view must be mounted on all NFS clients for accessing the migrated data. | [optional] 
**TargetViewName** | Pointer to **NullableString** | The target view name to which the data will be migrated. | [optional] 

## Methods

### NewDataMigrationJobParameters

`func NewDataMigrationJobParameters() *DataMigrationJobParameters`

NewDataMigrationJobParameters instantiates a new DataMigrationJobParameters object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDataMigrationJobParametersWithDefaults

`func NewDataMigrationJobParametersWithDefaults() *DataMigrationJobParameters`

NewDataMigrationJobParametersWithDefaults instantiates a new DataMigrationJobParameters object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetColdFileWindow

`func (o *DataMigrationJobParameters) GetColdFileWindow() int64`

GetColdFileWindow returns the ColdFileWindow field if non-nil, zero value otherwise.

### GetColdFileWindowOk

`func (o *DataMigrationJobParameters) GetColdFileWindowOk() (*int64, bool)`

GetColdFileWindowOk returns a tuple with the ColdFileWindow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColdFileWindow

`func (o *DataMigrationJobParameters) SetColdFileWindow(v int64)`

SetColdFileWindow sets ColdFileWindow field to given value.

### HasColdFileWindow

`func (o *DataMigrationJobParameters) HasColdFileWindow() bool`

HasColdFileWindow returns a boolean if a field has been set.

### SetColdFileWindowNil

`func (o *DataMigrationJobParameters) SetColdFileWindowNil(b bool)`

 SetColdFileWindowNil sets the value for ColdFileWindow to be an explicit nil

### UnsetColdFileWindow
`func (o *DataMigrationJobParameters) UnsetColdFileWindow()`

UnsetColdFileWindow ensures that no value is present for ColdFileWindow, not even an explicit nil
### GetDeleteOrphanData

`func (o *DataMigrationJobParameters) GetDeleteOrphanData() bool`

GetDeleteOrphanData returns the DeleteOrphanData field if non-nil, zero value otherwise.

### GetDeleteOrphanDataOk

`func (o *DataMigrationJobParameters) GetDeleteOrphanDataOk() (*bool, bool)`

GetDeleteOrphanDataOk returns a tuple with the DeleteOrphanData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleteOrphanData

`func (o *DataMigrationJobParameters) SetDeleteOrphanData(v bool)`

SetDeleteOrphanData sets DeleteOrphanData field to given value.

### HasDeleteOrphanData

`func (o *DataMigrationJobParameters) HasDeleteOrphanData() bool`

HasDeleteOrphanData returns a boolean if a field has been set.

### SetDeleteOrphanDataNil

`func (o *DataMigrationJobParameters) SetDeleteOrphanDataNil(b bool)`

 SetDeleteOrphanDataNil sets the value for DeleteOrphanData to be an explicit nil

### UnsetDeleteOrphanData
`func (o *DataMigrationJobParameters) UnsetDeleteOrphanData()`

UnsetDeleteOrphanData ensures that no value is present for DeleteOrphanData, not even an explicit nil
### GetFilePathFilter

`func (o *DataMigrationJobParameters) GetFilePathFilter() FilePathFilter`

GetFilePathFilter returns the FilePathFilter field if non-nil, zero value otherwise.

### GetFilePathFilterOk

`func (o *DataMigrationJobParameters) GetFilePathFilterOk() (*FilePathFilter, bool)`

GetFilePathFilterOk returns a tuple with the FilePathFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilePathFilter

`func (o *DataMigrationJobParameters) SetFilePathFilter(v FilePathFilter)`

SetFilePathFilter sets FilePathFilter field to given value.

### HasFilePathFilter

`func (o *DataMigrationJobParameters) HasFilePathFilter() bool`

HasFilePathFilter returns a boolean if a field has been set.

### GetFileSelectionPolicy

`func (o *DataMigrationJobParameters) GetFileSelectionPolicy() string`

GetFileSelectionPolicy returns the FileSelectionPolicy field if non-nil, zero value otherwise.

### GetFileSelectionPolicyOk

`func (o *DataMigrationJobParameters) GetFileSelectionPolicyOk() (*string, bool)`

GetFileSelectionPolicyOk returns a tuple with the FileSelectionPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileSelectionPolicy

`func (o *DataMigrationJobParameters) SetFileSelectionPolicy(v string)`

SetFileSelectionPolicy sets FileSelectionPolicy field to given value.

### HasFileSelectionPolicy

`func (o *DataMigrationJobParameters) HasFileSelectionPolicy() bool`

HasFileSelectionPolicy returns a boolean if a field has been set.

### SetFileSelectionPolicyNil

`func (o *DataMigrationJobParameters) SetFileSelectionPolicyNil(b bool)`

 SetFileSelectionPolicyNil sets the value for FileSelectionPolicy to be an explicit nil

### UnsetFileSelectionPolicy
`func (o *DataMigrationJobParameters) UnsetFileSelectionPolicy()`

UnsetFileSelectionPolicy ensures that no value is present for FileSelectionPolicy, not even an explicit nil
### GetFileSizeBytes

`func (o *DataMigrationJobParameters) GetFileSizeBytes() int64`

GetFileSizeBytes returns the FileSizeBytes field if non-nil, zero value otherwise.

### GetFileSizeBytesOk

`func (o *DataMigrationJobParameters) GetFileSizeBytesOk() (*int64, bool)`

GetFileSizeBytesOk returns a tuple with the FileSizeBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileSizeBytes

`func (o *DataMigrationJobParameters) SetFileSizeBytes(v int64)`

SetFileSizeBytes sets FileSizeBytes field to given value.

### HasFileSizeBytes

`func (o *DataMigrationJobParameters) HasFileSizeBytes() bool`

HasFileSizeBytes returns a boolean if a field has been set.

### SetFileSizeBytesNil

`func (o *DataMigrationJobParameters) SetFileSizeBytesNil(b bool)`

 SetFileSizeBytesNil sets the value for FileSizeBytes to be an explicit nil

### UnsetFileSizeBytes
`func (o *DataMigrationJobParameters) UnsetFileSizeBytes()`

UnsetFileSizeBytes ensures that no value is present for FileSizeBytes, not even an explicit nil
### GetFileSizePolicy

`func (o *DataMigrationJobParameters) GetFileSizePolicy() string`

GetFileSizePolicy returns the FileSizePolicy field if non-nil, zero value otherwise.

### GetFileSizePolicyOk

`func (o *DataMigrationJobParameters) GetFileSizePolicyOk() (*string, bool)`

GetFileSizePolicyOk returns a tuple with the FileSizePolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileSizePolicy

`func (o *DataMigrationJobParameters) SetFileSizePolicy(v string)`

SetFileSizePolicy sets FileSizePolicy field to given value.

### HasFileSizePolicy

`func (o *DataMigrationJobParameters) HasFileSizePolicy() bool`

HasFileSizePolicy returns a boolean if a field has been set.

### SetFileSizePolicyNil

`func (o *DataMigrationJobParameters) SetFileSizePolicyNil(b bool)`

 SetFileSizePolicyNil sets the value for FileSizePolicy to be an explicit nil

### UnsetFileSizePolicy
`func (o *DataMigrationJobParameters) UnsetFileSizePolicy()`

UnsetFileSizePolicy ensures that no value is present for FileSizePolicy, not even an explicit nil
### GetMigrateWithoutStub

`func (o *DataMigrationJobParameters) GetMigrateWithoutStub() bool`

GetMigrateWithoutStub returns the MigrateWithoutStub field if non-nil, zero value otherwise.

### GetMigrateWithoutStubOk

`func (o *DataMigrationJobParameters) GetMigrateWithoutStubOk() (*bool, bool)`

GetMigrateWithoutStubOk returns a tuple with the MigrateWithoutStub field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMigrateWithoutStub

`func (o *DataMigrationJobParameters) SetMigrateWithoutStub(v bool)`

SetMigrateWithoutStub sets MigrateWithoutStub field to given value.

### HasMigrateWithoutStub

`func (o *DataMigrationJobParameters) HasMigrateWithoutStub() bool`

HasMigrateWithoutStub returns a boolean if a field has been set.

### SetMigrateWithoutStubNil

`func (o *DataMigrationJobParameters) SetMigrateWithoutStubNil(b bool)`

 SetMigrateWithoutStubNil sets the value for MigrateWithoutStub to be an explicit nil

### UnsetMigrateWithoutStub
`func (o *DataMigrationJobParameters) UnsetMigrateWithoutStub()`

UnsetMigrateWithoutStub ensures that no value is present for MigrateWithoutStub, not even an explicit nil
### GetNfsMountPath

`func (o *DataMigrationJobParameters) GetNfsMountPath() string`

GetNfsMountPath returns the NfsMountPath field if non-nil, zero value otherwise.

### GetNfsMountPathOk

`func (o *DataMigrationJobParameters) GetNfsMountPathOk() (*string, bool)`

GetNfsMountPathOk returns a tuple with the NfsMountPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNfsMountPath

`func (o *DataMigrationJobParameters) SetNfsMountPath(v string)`

SetNfsMountPath sets NfsMountPath field to given value.

### HasNfsMountPath

`func (o *DataMigrationJobParameters) HasNfsMountPath() bool`

HasNfsMountPath returns a boolean if a field has been set.

### SetNfsMountPathNil

`func (o *DataMigrationJobParameters) SetNfsMountPathNil(b bool)`

 SetNfsMountPathNil sets the value for NfsMountPath to be an explicit nil

### UnsetNfsMountPath
`func (o *DataMigrationJobParameters) UnsetNfsMountPath()`

UnsetNfsMountPath ensures that no value is present for NfsMountPath, not even an explicit nil
### GetTargetViewName

`func (o *DataMigrationJobParameters) GetTargetViewName() string`

GetTargetViewName returns the TargetViewName field if non-nil, zero value otherwise.

### GetTargetViewNameOk

`func (o *DataMigrationJobParameters) GetTargetViewNameOk() (*string, bool)`

GetTargetViewNameOk returns a tuple with the TargetViewName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetViewName

`func (o *DataMigrationJobParameters) SetTargetViewName(v string)`

SetTargetViewName sets TargetViewName field to given value.

### HasTargetViewName

`func (o *DataMigrationJobParameters) HasTargetViewName() bool`

HasTargetViewName returns a boolean if a field has been set.

### SetTargetViewNameNil

`func (o *DataMigrationJobParameters) SetTargetViewNameNil(b bool)`

 SetTargetViewNameNil sets the value for TargetViewName to be an explicit nil

### UnsetTargetViewName
`func (o *DataMigrationJobParameters) UnsetTargetViewName()`

UnsetTargetViewName ensures that no value is present for TargetViewName, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


