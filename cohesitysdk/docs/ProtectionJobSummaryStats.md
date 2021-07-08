# ProtectionJobSummaryStats

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AverageRunTimeUsecs** | Pointer to **NullableInt64** | Specifies the average run time of all successful Protection Runs. It is specified as a Unix epoch Timestamp (in microseconds). | [optional] 
**FastestRunTimeUsecs** | Pointer to **NullableInt64** | Specifies the time taken for a fastest successful Protection Run so far. It is specified as a Unix epoch Timestamp (in microseconds). | [optional] 
**NumCanceledRuns** | Pointer to **NullableInt64** | Specifies the number of runs that were cancelled. | [optional] 
**NumFailedRuns** | Pointer to **NullableInt64** | Specifies the number of runs that failed to finish. | [optional] 
**NumSlaViolations** | Pointer to **NullableInt64** | Specifies the number of runs having SLA violations. | [optional] 
**NumSuccessfulRuns** | Pointer to **NullableInt64** | Specifies the number of runs that finished successfully. | [optional] 
**SlowestRunTimeUsecs** | Pointer to **NullableInt64** | Specifies the time taken for a slowest successful Protection Run so far. It is specified as a Unix epoch Timestamp (in microseconds). | [optional] 
**TotalBytesReadFromSource** | Pointer to **NullableInt64** | Specifies the total amount of data read from the source (so far). | [optional] 
**TotalLogicalBackupSizeBytes** | Pointer to **NullableInt64** | Specifies the size of the source object (such as a VM) protected by this task on the primary storage after the snapshot is taken. The logical size of the data on the source if the data is fully hydrated or expanded and not reduced by change-block tracking, compression and deduplication. | [optional] 
**TotalPhysicalBackupSizeBytes** | Pointer to **NullableInt64** | Specifies the total amount of physical space used on the Cohesity Cluster to store the protected object after being reduced by change-block tracking, compression and deduplication. For example, if the logical backup size is 1GB, but only 1MB was used on the Cohesity Cluster to store it, this field be equal to 1MB. | [optional] 

## Methods

### NewProtectionJobSummaryStats

`func NewProtectionJobSummaryStats() *ProtectionJobSummaryStats`

NewProtectionJobSummaryStats instantiates a new ProtectionJobSummaryStats object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProtectionJobSummaryStatsWithDefaults

`func NewProtectionJobSummaryStatsWithDefaults() *ProtectionJobSummaryStats`

NewProtectionJobSummaryStatsWithDefaults instantiates a new ProtectionJobSummaryStats object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAverageRunTimeUsecs

`func (o *ProtectionJobSummaryStats) GetAverageRunTimeUsecs() int64`

GetAverageRunTimeUsecs returns the AverageRunTimeUsecs field if non-nil, zero value otherwise.

### GetAverageRunTimeUsecsOk

`func (o *ProtectionJobSummaryStats) GetAverageRunTimeUsecsOk() (*int64, bool)`

GetAverageRunTimeUsecsOk returns a tuple with the AverageRunTimeUsecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAverageRunTimeUsecs

`func (o *ProtectionJobSummaryStats) SetAverageRunTimeUsecs(v int64)`

SetAverageRunTimeUsecs sets AverageRunTimeUsecs field to given value.

### HasAverageRunTimeUsecs

`func (o *ProtectionJobSummaryStats) HasAverageRunTimeUsecs() bool`

HasAverageRunTimeUsecs returns a boolean if a field has been set.

### SetAverageRunTimeUsecsNil

`func (o *ProtectionJobSummaryStats) SetAverageRunTimeUsecsNil(b bool)`

 SetAverageRunTimeUsecsNil sets the value for AverageRunTimeUsecs to be an explicit nil

### UnsetAverageRunTimeUsecs
`func (o *ProtectionJobSummaryStats) UnsetAverageRunTimeUsecs()`

UnsetAverageRunTimeUsecs ensures that no value is present for AverageRunTimeUsecs, not even an explicit nil
### GetFastestRunTimeUsecs

`func (o *ProtectionJobSummaryStats) GetFastestRunTimeUsecs() int64`

GetFastestRunTimeUsecs returns the FastestRunTimeUsecs field if non-nil, zero value otherwise.

### GetFastestRunTimeUsecsOk

`func (o *ProtectionJobSummaryStats) GetFastestRunTimeUsecsOk() (*int64, bool)`

GetFastestRunTimeUsecsOk returns a tuple with the FastestRunTimeUsecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFastestRunTimeUsecs

`func (o *ProtectionJobSummaryStats) SetFastestRunTimeUsecs(v int64)`

SetFastestRunTimeUsecs sets FastestRunTimeUsecs field to given value.

### HasFastestRunTimeUsecs

`func (o *ProtectionJobSummaryStats) HasFastestRunTimeUsecs() bool`

HasFastestRunTimeUsecs returns a boolean if a field has been set.

### SetFastestRunTimeUsecsNil

`func (o *ProtectionJobSummaryStats) SetFastestRunTimeUsecsNil(b bool)`

 SetFastestRunTimeUsecsNil sets the value for FastestRunTimeUsecs to be an explicit nil

### UnsetFastestRunTimeUsecs
`func (o *ProtectionJobSummaryStats) UnsetFastestRunTimeUsecs()`

UnsetFastestRunTimeUsecs ensures that no value is present for FastestRunTimeUsecs, not even an explicit nil
### GetNumCanceledRuns

`func (o *ProtectionJobSummaryStats) GetNumCanceledRuns() int64`

GetNumCanceledRuns returns the NumCanceledRuns field if non-nil, zero value otherwise.

### GetNumCanceledRunsOk

`func (o *ProtectionJobSummaryStats) GetNumCanceledRunsOk() (*int64, bool)`

GetNumCanceledRunsOk returns a tuple with the NumCanceledRuns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumCanceledRuns

`func (o *ProtectionJobSummaryStats) SetNumCanceledRuns(v int64)`

SetNumCanceledRuns sets NumCanceledRuns field to given value.

### HasNumCanceledRuns

`func (o *ProtectionJobSummaryStats) HasNumCanceledRuns() bool`

HasNumCanceledRuns returns a boolean if a field has been set.

### SetNumCanceledRunsNil

`func (o *ProtectionJobSummaryStats) SetNumCanceledRunsNil(b bool)`

 SetNumCanceledRunsNil sets the value for NumCanceledRuns to be an explicit nil

### UnsetNumCanceledRuns
`func (o *ProtectionJobSummaryStats) UnsetNumCanceledRuns()`

UnsetNumCanceledRuns ensures that no value is present for NumCanceledRuns, not even an explicit nil
### GetNumFailedRuns

`func (o *ProtectionJobSummaryStats) GetNumFailedRuns() int64`

GetNumFailedRuns returns the NumFailedRuns field if non-nil, zero value otherwise.

### GetNumFailedRunsOk

`func (o *ProtectionJobSummaryStats) GetNumFailedRunsOk() (*int64, bool)`

GetNumFailedRunsOk returns a tuple with the NumFailedRuns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumFailedRuns

`func (o *ProtectionJobSummaryStats) SetNumFailedRuns(v int64)`

SetNumFailedRuns sets NumFailedRuns field to given value.

### HasNumFailedRuns

`func (o *ProtectionJobSummaryStats) HasNumFailedRuns() bool`

HasNumFailedRuns returns a boolean if a field has been set.

### SetNumFailedRunsNil

`func (o *ProtectionJobSummaryStats) SetNumFailedRunsNil(b bool)`

 SetNumFailedRunsNil sets the value for NumFailedRuns to be an explicit nil

### UnsetNumFailedRuns
`func (o *ProtectionJobSummaryStats) UnsetNumFailedRuns()`

UnsetNumFailedRuns ensures that no value is present for NumFailedRuns, not even an explicit nil
### GetNumSlaViolations

`func (o *ProtectionJobSummaryStats) GetNumSlaViolations() int64`

GetNumSlaViolations returns the NumSlaViolations field if non-nil, zero value otherwise.

### GetNumSlaViolationsOk

`func (o *ProtectionJobSummaryStats) GetNumSlaViolationsOk() (*int64, bool)`

GetNumSlaViolationsOk returns a tuple with the NumSlaViolations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumSlaViolations

`func (o *ProtectionJobSummaryStats) SetNumSlaViolations(v int64)`

SetNumSlaViolations sets NumSlaViolations field to given value.

### HasNumSlaViolations

`func (o *ProtectionJobSummaryStats) HasNumSlaViolations() bool`

HasNumSlaViolations returns a boolean if a field has been set.

### SetNumSlaViolationsNil

`func (o *ProtectionJobSummaryStats) SetNumSlaViolationsNil(b bool)`

 SetNumSlaViolationsNil sets the value for NumSlaViolations to be an explicit nil

### UnsetNumSlaViolations
`func (o *ProtectionJobSummaryStats) UnsetNumSlaViolations()`

UnsetNumSlaViolations ensures that no value is present for NumSlaViolations, not even an explicit nil
### GetNumSuccessfulRuns

`func (o *ProtectionJobSummaryStats) GetNumSuccessfulRuns() int64`

GetNumSuccessfulRuns returns the NumSuccessfulRuns field if non-nil, zero value otherwise.

### GetNumSuccessfulRunsOk

`func (o *ProtectionJobSummaryStats) GetNumSuccessfulRunsOk() (*int64, bool)`

GetNumSuccessfulRunsOk returns a tuple with the NumSuccessfulRuns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumSuccessfulRuns

`func (o *ProtectionJobSummaryStats) SetNumSuccessfulRuns(v int64)`

SetNumSuccessfulRuns sets NumSuccessfulRuns field to given value.

### HasNumSuccessfulRuns

`func (o *ProtectionJobSummaryStats) HasNumSuccessfulRuns() bool`

HasNumSuccessfulRuns returns a boolean if a field has been set.

### SetNumSuccessfulRunsNil

`func (o *ProtectionJobSummaryStats) SetNumSuccessfulRunsNil(b bool)`

 SetNumSuccessfulRunsNil sets the value for NumSuccessfulRuns to be an explicit nil

### UnsetNumSuccessfulRuns
`func (o *ProtectionJobSummaryStats) UnsetNumSuccessfulRuns()`

UnsetNumSuccessfulRuns ensures that no value is present for NumSuccessfulRuns, not even an explicit nil
### GetSlowestRunTimeUsecs

`func (o *ProtectionJobSummaryStats) GetSlowestRunTimeUsecs() int64`

GetSlowestRunTimeUsecs returns the SlowestRunTimeUsecs field if non-nil, zero value otherwise.

### GetSlowestRunTimeUsecsOk

`func (o *ProtectionJobSummaryStats) GetSlowestRunTimeUsecsOk() (*int64, bool)`

GetSlowestRunTimeUsecsOk returns a tuple with the SlowestRunTimeUsecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlowestRunTimeUsecs

`func (o *ProtectionJobSummaryStats) SetSlowestRunTimeUsecs(v int64)`

SetSlowestRunTimeUsecs sets SlowestRunTimeUsecs field to given value.

### HasSlowestRunTimeUsecs

`func (o *ProtectionJobSummaryStats) HasSlowestRunTimeUsecs() bool`

HasSlowestRunTimeUsecs returns a boolean if a field has been set.

### SetSlowestRunTimeUsecsNil

`func (o *ProtectionJobSummaryStats) SetSlowestRunTimeUsecsNil(b bool)`

 SetSlowestRunTimeUsecsNil sets the value for SlowestRunTimeUsecs to be an explicit nil

### UnsetSlowestRunTimeUsecs
`func (o *ProtectionJobSummaryStats) UnsetSlowestRunTimeUsecs()`

UnsetSlowestRunTimeUsecs ensures that no value is present for SlowestRunTimeUsecs, not even an explicit nil
### GetTotalBytesReadFromSource

`func (o *ProtectionJobSummaryStats) GetTotalBytesReadFromSource() int64`

GetTotalBytesReadFromSource returns the TotalBytesReadFromSource field if non-nil, zero value otherwise.

### GetTotalBytesReadFromSourceOk

`func (o *ProtectionJobSummaryStats) GetTotalBytesReadFromSourceOk() (*int64, bool)`

GetTotalBytesReadFromSourceOk returns a tuple with the TotalBytesReadFromSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalBytesReadFromSource

`func (o *ProtectionJobSummaryStats) SetTotalBytesReadFromSource(v int64)`

SetTotalBytesReadFromSource sets TotalBytesReadFromSource field to given value.

### HasTotalBytesReadFromSource

`func (o *ProtectionJobSummaryStats) HasTotalBytesReadFromSource() bool`

HasTotalBytesReadFromSource returns a boolean if a field has been set.

### SetTotalBytesReadFromSourceNil

`func (o *ProtectionJobSummaryStats) SetTotalBytesReadFromSourceNil(b bool)`

 SetTotalBytesReadFromSourceNil sets the value for TotalBytesReadFromSource to be an explicit nil

### UnsetTotalBytesReadFromSource
`func (o *ProtectionJobSummaryStats) UnsetTotalBytesReadFromSource()`

UnsetTotalBytesReadFromSource ensures that no value is present for TotalBytesReadFromSource, not even an explicit nil
### GetTotalLogicalBackupSizeBytes

`func (o *ProtectionJobSummaryStats) GetTotalLogicalBackupSizeBytes() int64`

GetTotalLogicalBackupSizeBytes returns the TotalLogicalBackupSizeBytes field if non-nil, zero value otherwise.

### GetTotalLogicalBackupSizeBytesOk

`func (o *ProtectionJobSummaryStats) GetTotalLogicalBackupSizeBytesOk() (*int64, bool)`

GetTotalLogicalBackupSizeBytesOk returns a tuple with the TotalLogicalBackupSizeBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalLogicalBackupSizeBytes

`func (o *ProtectionJobSummaryStats) SetTotalLogicalBackupSizeBytes(v int64)`

SetTotalLogicalBackupSizeBytes sets TotalLogicalBackupSizeBytes field to given value.

### HasTotalLogicalBackupSizeBytes

`func (o *ProtectionJobSummaryStats) HasTotalLogicalBackupSizeBytes() bool`

HasTotalLogicalBackupSizeBytes returns a boolean if a field has been set.

### SetTotalLogicalBackupSizeBytesNil

`func (o *ProtectionJobSummaryStats) SetTotalLogicalBackupSizeBytesNil(b bool)`

 SetTotalLogicalBackupSizeBytesNil sets the value for TotalLogicalBackupSizeBytes to be an explicit nil

### UnsetTotalLogicalBackupSizeBytes
`func (o *ProtectionJobSummaryStats) UnsetTotalLogicalBackupSizeBytes()`

UnsetTotalLogicalBackupSizeBytes ensures that no value is present for TotalLogicalBackupSizeBytes, not even an explicit nil
### GetTotalPhysicalBackupSizeBytes

`func (o *ProtectionJobSummaryStats) GetTotalPhysicalBackupSizeBytes() int64`

GetTotalPhysicalBackupSizeBytes returns the TotalPhysicalBackupSizeBytes field if non-nil, zero value otherwise.

### GetTotalPhysicalBackupSizeBytesOk

`func (o *ProtectionJobSummaryStats) GetTotalPhysicalBackupSizeBytesOk() (*int64, bool)`

GetTotalPhysicalBackupSizeBytesOk returns a tuple with the TotalPhysicalBackupSizeBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalPhysicalBackupSizeBytes

`func (o *ProtectionJobSummaryStats) SetTotalPhysicalBackupSizeBytes(v int64)`

SetTotalPhysicalBackupSizeBytes sets TotalPhysicalBackupSizeBytes field to given value.

### HasTotalPhysicalBackupSizeBytes

`func (o *ProtectionJobSummaryStats) HasTotalPhysicalBackupSizeBytes() bool`

HasTotalPhysicalBackupSizeBytes returns a boolean if a field has been set.

### SetTotalPhysicalBackupSizeBytesNil

`func (o *ProtectionJobSummaryStats) SetTotalPhysicalBackupSizeBytesNil(b bool)`

 SetTotalPhysicalBackupSizeBytesNil sets the value for TotalPhysicalBackupSizeBytes to be an explicit nil

### UnsetTotalPhysicalBackupSizeBytes
`func (o *ProtectionJobSummaryStats) UnsetTotalPhysicalBackupSizeBytes()`

UnsetTotalPhysicalBackupSizeBytes ensures that no value is present for TotalPhysicalBackupSizeBytes, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


