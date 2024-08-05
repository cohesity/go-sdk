# PhysicalEnvJobParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CobmrBackup** | Pointer to **NullableBool** | Specifies whether to enable CoBMR backup. | [optional] 
**FilePathFilters** | Pointer to [**FileFilteringPolicy**](FileFilteringPolicy.md) |  | [optional] 
**IncrementalSnapshotUponRestart** | Pointer to **NullableBool** | If true, performs an incremental backup after server restarts. Otherwise a full backup is done. NOTE: This is applicable only to Windows servers. If not set, default value is false. | [optional] 

## Methods

### NewPhysicalEnvJobParams

`func NewPhysicalEnvJobParams() *PhysicalEnvJobParams`

NewPhysicalEnvJobParams instantiates a new PhysicalEnvJobParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPhysicalEnvJobParamsWithDefaults

`func NewPhysicalEnvJobParamsWithDefaults() *PhysicalEnvJobParams`

NewPhysicalEnvJobParamsWithDefaults instantiates a new PhysicalEnvJobParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCobmrBackup

`func (o *PhysicalEnvJobParams) GetCobmrBackup() bool`

GetCobmrBackup returns the CobmrBackup field if non-nil, zero value otherwise.

### GetCobmrBackupOk

`func (o *PhysicalEnvJobParams) GetCobmrBackupOk() (*bool, bool)`

GetCobmrBackupOk returns a tuple with the CobmrBackup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCobmrBackup

`func (o *PhysicalEnvJobParams) SetCobmrBackup(v bool)`

SetCobmrBackup sets CobmrBackup field to given value.

### HasCobmrBackup

`func (o *PhysicalEnvJobParams) HasCobmrBackup() bool`

HasCobmrBackup returns a boolean if a field has been set.

### SetCobmrBackupNil

`func (o *PhysicalEnvJobParams) SetCobmrBackupNil(b bool)`

 SetCobmrBackupNil sets the value for CobmrBackup to be an explicit nil

### UnsetCobmrBackup
`func (o *PhysicalEnvJobParams) UnsetCobmrBackup()`

UnsetCobmrBackup ensures that no value is present for CobmrBackup, not even an explicit nil
### GetFilePathFilters

`func (o *PhysicalEnvJobParams) GetFilePathFilters() FileFilteringPolicy`

GetFilePathFilters returns the FilePathFilters field if non-nil, zero value otherwise.

### GetFilePathFiltersOk

`func (o *PhysicalEnvJobParams) GetFilePathFiltersOk() (*FileFilteringPolicy, bool)`

GetFilePathFiltersOk returns a tuple with the FilePathFilters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilePathFilters

`func (o *PhysicalEnvJobParams) SetFilePathFilters(v FileFilteringPolicy)`

SetFilePathFilters sets FilePathFilters field to given value.

### HasFilePathFilters

`func (o *PhysicalEnvJobParams) HasFilePathFilters() bool`

HasFilePathFilters returns a boolean if a field has been set.

### GetIncrementalSnapshotUponRestart

`func (o *PhysicalEnvJobParams) GetIncrementalSnapshotUponRestart() bool`

GetIncrementalSnapshotUponRestart returns the IncrementalSnapshotUponRestart field if non-nil, zero value otherwise.

### GetIncrementalSnapshotUponRestartOk

`func (o *PhysicalEnvJobParams) GetIncrementalSnapshotUponRestartOk() (*bool, bool)`

GetIncrementalSnapshotUponRestartOk returns a tuple with the IncrementalSnapshotUponRestart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncrementalSnapshotUponRestart

`func (o *PhysicalEnvJobParams) SetIncrementalSnapshotUponRestart(v bool)`

SetIncrementalSnapshotUponRestart sets IncrementalSnapshotUponRestart field to given value.

### HasIncrementalSnapshotUponRestart

`func (o *PhysicalEnvJobParams) HasIncrementalSnapshotUponRestart() bool`

HasIncrementalSnapshotUponRestart returns a boolean if a field has been set.

### SetIncrementalSnapshotUponRestartNil

`func (o *PhysicalEnvJobParams) SetIncrementalSnapshotUponRestartNil(b bool)`

 SetIncrementalSnapshotUponRestartNil sets the value for IncrementalSnapshotUponRestart to be an explicit nil

### UnsetIncrementalSnapshotUponRestart
`func (o *PhysicalEnvJobParams) UnsetIncrementalSnapshotUponRestart()`

UnsetIncrementalSnapshotUponRestart ensures that no value is present for IncrementalSnapshotUponRestart, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


