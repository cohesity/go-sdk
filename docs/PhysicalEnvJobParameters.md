# PhysicalEnvJobParameters

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FilePathFilters** | Pointer to [**FilePathFilter**](FilePathFilter.md) |  | [optional] 
**IncrementalSnapshotUponRestart** | Pointer to **NullableBool** | If true, performs an incremental backup after server restarts. Otherwise a full backup is done. NOTE: This is applicable only to Windows servers. If not set, default value is false. | [optional] 

## Methods

### NewPhysicalEnvJobParameters

`func NewPhysicalEnvJobParameters() *PhysicalEnvJobParameters`

NewPhysicalEnvJobParameters instantiates a new PhysicalEnvJobParameters object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPhysicalEnvJobParametersWithDefaults

`func NewPhysicalEnvJobParametersWithDefaults() *PhysicalEnvJobParameters`

NewPhysicalEnvJobParametersWithDefaults instantiates a new PhysicalEnvJobParameters object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFilePathFilters

`func (o *PhysicalEnvJobParameters) GetFilePathFilters() FilePathFilter`

GetFilePathFilters returns the FilePathFilters field if non-nil, zero value otherwise.

### GetFilePathFiltersOk

`func (o *PhysicalEnvJobParameters) GetFilePathFiltersOk() (*FilePathFilter, bool)`

GetFilePathFiltersOk returns a tuple with the FilePathFilters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilePathFilters

`func (o *PhysicalEnvJobParameters) SetFilePathFilters(v FilePathFilter)`

SetFilePathFilters sets FilePathFilters field to given value.

### HasFilePathFilters

`func (o *PhysicalEnvJobParameters) HasFilePathFilters() bool`

HasFilePathFilters returns a boolean if a field has been set.

### GetIncrementalSnapshotUponRestart

`func (o *PhysicalEnvJobParameters) GetIncrementalSnapshotUponRestart() bool`

GetIncrementalSnapshotUponRestart returns the IncrementalSnapshotUponRestart field if non-nil, zero value otherwise.

### GetIncrementalSnapshotUponRestartOk

`func (o *PhysicalEnvJobParameters) GetIncrementalSnapshotUponRestartOk() (*bool, bool)`

GetIncrementalSnapshotUponRestartOk returns a tuple with the IncrementalSnapshotUponRestart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncrementalSnapshotUponRestart

`func (o *PhysicalEnvJobParameters) SetIncrementalSnapshotUponRestart(v bool)`

SetIncrementalSnapshotUponRestart sets IncrementalSnapshotUponRestart field to given value.

### HasIncrementalSnapshotUponRestart

`func (o *PhysicalEnvJobParameters) HasIncrementalSnapshotUponRestart() bool`

HasIncrementalSnapshotUponRestart returns a boolean if a field has been set.

### SetIncrementalSnapshotUponRestartNil

`func (o *PhysicalEnvJobParameters) SetIncrementalSnapshotUponRestartNil(b bool)`

 SetIncrementalSnapshotUponRestartNil sets the value for IncrementalSnapshotUponRestart to be an explicit nil

### UnsetIncrementalSnapshotUponRestart
`func (o *PhysicalEnvJobParameters) UnsetIncrementalSnapshotUponRestart()`

UnsetIncrementalSnapshotUponRestart ensures that no value is present for IncrementalSnapshotUponRestart, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


