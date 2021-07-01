# FileStubbingParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ColdFileWindow** | Pointer to **NullableInt64** | Identifies the cold files in the NAS source. Files that haven&#39;t been accessed/modified in the last cold_file_window msecs or are older than cold_window_msecs are migrated. | [optional] 
**DeleteOrphanData** | Pointer to **NullableBool** | Delete migrated data if no symlink at source is pointing to it. | [optional] 
**FileSelectPolicy** | Pointer to **NullableInt32** | File migrate policy based on file access/modify time and age. | [optional] 
**FileSize** | Pointer to **NullableInt64** | Gives the size criteria to be used for selecting the files to be migrated. The cold files that are equal and greater than file_size or smaller than file_size are migrated. | [optional] 
**FileSizePolicy** | Pointer to **NullableInt32** | File size policy for selecting files to migrate. | [optional] 
**FilteringPolicy** | Pointer to [**FilteringPolicyProto**](FilteringPolicyProto.md) |  | [optional] 
**MigrateWithoutStub** | Pointer to **NullableBool** | Migrate data without stub. | [optional] 
**NfsMountPath** | Pointer to **NullableString** | Mount path where the Cohesity target view must be mounted on all NFS clients for accessing the migrated data. | [optional] 
**TargetViewName** | Pointer to **NullableString** | The target view name to which the data will be migrated. | [optional] 

## Methods

### NewFileStubbingParams

`func NewFileStubbingParams() *FileStubbingParams`

NewFileStubbingParams instantiates a new FileStubbingParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFileStubbingParamsWithDefaults

`func NewFileStubbingParamsWithDefaults() *FileStubbingParams`

NewFileStubbingParamsWithDefaults instantiates a new FileStubbingParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetColdFileWindow

`func (o *FileStubbingParams) GetColdFileWindow() int64`

GetColdFileWindow returns the ColdFileWindow field if non-nil, zero value otherwise.

### GetColdFileWindowOk

`func (o *FileStubbingParams) GetColdFileWindowOk() (*int64, bool)`

GetColdFileWindowOk returns a tuple with the ColdFileWindow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColdFileWindow

`func (o *FileStubbingParams) SetColdFileWindow(v int64)`

SetColdFileWindow sets ColdFileWindow field to given value.

### HasColdFileWindow

`func (o *FileStubbingParams) HasColdFileWindow() bool`

HasColdFileWindow returns a boolean if a field has been set.

### SetColdFileWindowNil

`func (o *FileStubbingParams) SetColdFileWindowNil(b bool)`

 SetColdFileWindowNil sets the value for ColdFileWindow to be an explicit nil

### UnsetColdFileWindow
`func (o *FileStubbingParams) UnsetColdFileWindow()`

UnsetColdFileWindow ensures that no value is present for ColdFileWindow, not even an explicit nil
### GetDeleteOrphanData

`func (o *FileStubbingParams) GetDeleteOrphanData() bool`

GetDeleteOrphanData returns the DeleteOrphanData field if non-nil, zero value otherwise.

### GetDeleteOrphanDataOk

`func (o *FileStubbingParams) GetDeleteOrphanDataOk() (*bool, bool)`

GetDeleteOrphanDataOk returns a tuple with the DeleteOrphanData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleteOrphanData

`func (o *FileStubbingParams) SetDeleteOrphanData(v bool)`

SetDeleteOrphanData sets DeleteOrphanData field to given value.

### HasDeleteOrphanData

`func (o *FileStubbingParams) HasDeleteOrphanData() bool`

HasDeleteOrphanData returns a boolean if a field has been set.

### SetDeleteOrphanDataNil

`func (o *FileStubbingParams) SetDeleteOrphanDataNil(b bool)`

 SetDeleteOrphanDataNil sets the value for DeleteOrphanData to be an explicit nil

### UnsetDeleteOrphanData
`func (o *FileStubbingParams) UnsetDeleteOrphanData()`

UnsetDeleteOrphanData ensures that no value is present for DeleteOrphanData, not even an explicit nil
### GetFileSelectPolicy

`func (o *FileStubbingParams) GetFileSelectPolicy() int32`

GetFileSelectPolicy returns the FileSelectPolicy field if non-nil, zero value otherwise.

### GetFileSelectPolicyOk

`func (o *FileStubbingParams) GetFileSelectPolicyOk() (*int32, bool)`

GetFileSelectPolicyOk returns a tuple with the FileSelectPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileSelectPolicy

`func (o *FileStubbingParams) SetFileSelectPolicy(v int32)`

SetFileSelectPolicy sets FileSelectPolicy field to given value.

### HasFileSelectPolicy

`func (o *FileStubbingParams) HasFileSelectPolicy() bool`

HasFileSelectPolicy returns a boolean if a field has been set.

### SetFileSelectPolicyNil

`func (o *FileStubbingParams) SetFileSelectPolicyNil(b bool)`

 SetFileSelectPolicyNil sets the value for FileSelectPolicy to be an explicit nil

### UnsetFileSelectPolicy
`func (o *FileStubbingParams) UnsetFileSelectPolicy()`

UnsetFileSelectPolicy ensures that no value is present for FileSelectPolicy, not even an explicit nil
### GetFileSize

`func (o *FileStubbingParams) GetFileSize() int64`

GetFileSize returns the FileSize field if non-nil, zero value otherwise.

### GetFileSizeOk

`func (o *FileStubbingParams) GetFileSizeOk() (*int64, bool)`

GetFileSizeOk returns a tuple with the FileSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileSize

`func (o *FileStubbingParams) SetFileSize(v int64)`

SetFileSize sets FileSize field to given value.

### HasFileSize

`func (o *FileStubbingParams) HasFileSize() bool`

HasFileSize returns a boolean if a field has been set.

### SetFileSizeNil

`func (o *FileStubbingParams) SetFileSizeNil(b bool)`

 SetFileSizeNil sets the value for FileSize to be an explicit nil

### UnsetFileSize
`func (o *FileStubbingParams) UnsetFileSize()`

UnsetFileSize ensures that no value is present for FileSize, not even an explicit nil
### GetFileSizePolicy

`func (o *FileStubbingParams) GetFileSizePolicy() int32`

GetFileSizePolicy returns the FileSizePolicy field if non-nil, zero value otherwise.

### GetFileSizePolicyOk

`func (o *FileStubbingParams) GetFileSizePolicyOk() (*int32, bool)`

GetFileSizePolicyOk returns a tuple with the FileSizePolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileSizePolicy

`func (o *FileStubbingParams) SetFileSizePolicy(v int32)`

SetFileSizePolicy sets FileSizePolicy field to given value.

### HasFileSizePolicy

`func (o *FileStubbingParams) HasFileSizePolicy() bool`

HasFileSizePolicy returns a boolean if a field has been set.

### SetFileSizePolicyNil

`func (o *FileStubbingParams) SetFileSizePolicyNil(b bool)`

 SetFileSizePolicyNil sets the value for FileSizePolicy to be an explicit nil

### UnsetFileSizePolicy
`func (o *FileStubbingParams) UnsetFileSizePolicy()`

UnsetFileSizePolicy ensures that no value is present for FileSizePolicy, not even an explicit nil
### GetFilteringPolicy

`func (o *FileStubbingParams) GetFilteringPolicy() FilteringPolicyProto`

GetFilteringPolicy returns the FilteringPolicy field if non-nil, zero value otherwise.

### GetFilteringPolicyOk

`func (o *FileStubbingParams) GetFilteringPolicyOk() (*FilteringPolicyProto, bool)`

GetFilteringPolicyOk returns a tuple with the FilteringPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilteringPolicy

`func (o *FileStubbingParams) SetFilteringPolicy(v FilteringPolicyProto)`

SetFilteringPolicy sets FilteringPolicy field to given value.

### HasFilteringPolicy

`func (o *FileStubbingParams) HasFilteringPolicy() bool`

HasFilteringPolicy returns a boolean if a field has been set.

### GetMigrateWithoutStub

`func (o *FileStubbingParams) GetMigrateWithoutStub() bool`

GetMigrateWithoutStub returns the MigrateWithoutStub field if non-nil, zero value otherwise.

### GetMigrateWithoutStubOk

`func (o *FileStubbingParams) GetMigrateWithoutStubOk() (*bool, bool)`

GetMigrateWithoutStubOk returns a tuple with the MigrateWithoutStub field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMigrateWithoutStub

`func (o *FileStubbingParams) SetMigrateWithoutStub(v bool)`

SetMigrateWithoutStub sets MigrateWithoutStub field to given value.

### HasMigrateWithoutStub

`func (o *FileStubbingParams) HasMigrateWithoutStub() bool`

HasMigrateWithoutStub returns a boolean if a field has been set.

### SetMigrateWithoutStubNil

`func (o *FileStubbingParams) SetMigrateWithoutStubNil(b bool)`

 SetMigrateWithoutStubNil sets the value for MigrateWithoutStub to be an explicit nil

### UnsetMigrateWithoutStub
`func (o *FileStubbingParams) UnsetMigrateWithoutStub()`

UnsetMigrateWithoutStub ensures that no value is present for MigrateWithoutStub, not even an explicit nil
### GetNfsMountPath

`func (o *FileStubbingParams) GetNfsMountPath() string`

GetNfsMountPath returns the NfsMountPath field if non-nil, zero value otherwise.

### GetNfsMountPathOk

`func (o *FileStubbingParams) GetNfsMountPathOk() (*string, bool)`

GetNfsMountPathOk returns a tuple with the NfsMountPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNfsMountPath

`func (o *FileStubbingParams) SetNfsMountPath(v string)`

SetNfsMountPath sets NfsMountPath field to given value.

### HasNfsMountPath

`func (o *FileStubbingParams) HasNfsMountPath() bool`

HasNfsMountPath returns a boolean if a field has been set.

### SetNfsMountPathNil

`func (o *FileStubbingParams) SetNfsMountPathNil(b bool)`

 SetNfsMountPathNil sets the value for NfsMountPath to be an explicit nil

### UnsetNfsMountPath
`func (o *FileStubbingParams) UnsetNfsMountPath()`

UnsetNfsMountPath ensures that no value is present for NfsMountPath, not even an explicit nil
### GetTargetViewName

`func (o *FileStubbingParams) GetTargetViewName() string`

GetTargetViewName returns the TargetViewName field if non-nil, zero value otherwise.

### GetTargetViewNameOk

`func (o *FileStubbingParams) GetTargetViewNameOk() (*string, bool)`

GetTargetViewNameOk returns a tuple with the TargetViewName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetViewName

`func (o *FileStubbingParams) SetTargetViewName(v string)`

SetTargetViewName sets TargetViewName field to given value.

### HasTargetViewName

`func (o *FileStubbingParams) HasTargetViewName() bool`

HasTargetViewName returns a boolean if a field has been set.

### SetTargetViewNameNil

`func (o *FileStubbingParams) SetTargetViewNameNil(b bool)`

 SetTargetViewNameNil sets the value for TargetViewName to be an explicit nil

### UnsetTargetViewName
`func (o *FileStubbingParams) UnsetTargetViewName()`

UnsetTargetViewName ensures that no value is present for TargetViewName, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


