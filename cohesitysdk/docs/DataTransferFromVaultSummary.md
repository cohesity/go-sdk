# DataTransferFromVaultSummary

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DataTransferPerTask** | Pointer to [**[]DataTransferFromVaultPerTask**](DataTransferFromVaultPerTask.md) | Array of Data Transferred Per Task.  Specifies the transfer of data from this Vault to this Cohesity Cluster for each clone or recover task. | [optional] 
**NumLogicalBytesTransferred** | Pointer to **NullableInt64** | Specifies the total number of logical bytes that have been transferred from this Vault (External Target) to this Cohesity Cluster. The logical size is when the data is fully hydrated or expanded. | [optional] 
**NumPhysicalBytesTransferred** | Pointer to **NullableInt64** | Specifies the total number of physical bytes that have been transferred from this Vault (External Target) to the Cohesity Cluster. | [optional] 
**NumTasks** | Pointer to **NullableInt64** | Specifies the number of recover or clone tasks that have transferred data from this Vault (External Target) to this Cohesity Cluster. | [optional] 
**PhysicalDataTransferredBytesDuringTimeRange** | Pointer to **[]int64** | Array of Physical Data Transferred Per Day.  Specifies the physical data transferred from this Vault to the Cohesity Cluster during the time period specified using the startTimeMsecs and endTimeMsecs parameters. For each day in the time period, an array element is returned, for example if 7 days are specified, 7 array elements are returned. | [optional] 
**VaultName** | Pointer to **NullableString** | Specifies the name of the Vault (External Target). | [optional] 

## Methods

### NewDataTransferFromVaultSummary

`func NewDataTransferFromVaultSummary() *DataTransferFromVaultSummary`

NewDataTransferFromVaultSummary instantiates a new DataTransferFromVaultSummary object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDataTransferFromVaultSummaryWithDefaults

`func NewDataTransferFromVaultSummaryWithDefaults() *DataTransferFromVaultSummary`

NewDataTransferFromVaultSummaryWithDefaults instantiates a new DataTransferFromVaultSummary object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDataTransferPerTask

`func (o *DataTransferFromVaultSummary) GetDataTransferPerTask() []DataTransferFromVaultPerTask`

GetDataTransferPerTask returns the DataTransferPerTask field if non-nil, zero value otherwise.

### GetDataTransferPerTaskOk

`func (o *DataTransferFromVaultSummary) GetDataTransferPerTaskOk() (*[]DataTransferFromVaultPerTask, bool)`

GetDataTransferPerTaskOk returns a tuple with the DataTransferPerTask field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataTransferPerTask

`func (o *DataTransferFromVaultSummary) SetDataTransferPerTask(v []DataTransferFromVaultPerTask)`

SetDataTransferPerTask sets DataTransferPerTask field to given value.

### HasDataTransferPerTask

`func (o *DataTransferFromVaultSummary) HasDataTransferPerTask() bool`

HasDataTransferPerTask returns a boolean if a field has been set.

### SetDataTransferPerTaskNil

`func (o *DataTransferFromVaultSummary) SetDataTransferPerTaskNil(b bool)`

 SetDataTransferPerTaskNil sets the value for DataTransferPerTask to be an explicit nil

### UnsetDataTransferPerTask
`func (o *DataTransferFromVaultSummary) UnsetDataTransferPerTask()`

UnsetDataTransferPerTask ensures that no value is present for DataTransferPerTask, not even an explicit nil
### GetNumLogicalBytesTransferred

`func (o *DataTransferFromVaultSummary) GetNumLogicalBytesTransferred() int64`

GetNumLogicalBytesTransferred returns the NumLogicalBytesTransferred field if non-nil, zero value otherwise.

### GetNumLogicalBytesTransferredOk

`func (o *DataTransferFromVaultSummary) GetNumLogicalBytesTransferredOk() (*int64, bool)`

GetNumLogicalBytesTransferredOk returns a tuple with the NumLogicalBytesTransferred field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumLogicalBytesTransferred

`func (o *DataTransferFromVaultSummary) SetNumLogicalBytesTransferred(v int64)`

SetNumLogicalBytesTransferred sets NumLogicalBytesTransferred field to given value.

### HasNumLogicalBytesTransferred

`func (o *DataTransferFromVaultSummary) HasNumLogicalBytesTransferred() bool`

HasNumLogicalBytesTransferred returns a boolean if a field has been set.

### SetNumLogicalBytesTransferredNil

`func (o *DataTransferFromVaultSummary) SetNumLogicalBytesTransferredNil(b bool)`

 SetNumLogicalBytesTransferredNil sets the value for NumLogicalBytesTransferred to be an explicit nil

### UnsetNumLogicalBytesTransferred
`func (o *DataTransferFromVaultSummary) UnsetNumLogicalBytesTransferred()`

UnsetNumLogicalBytesTransferred ensures that no value is present for NumLogicalBytesTransferred, not even an explicit nil
### GetNumPhysicalBytesTransferred

`func (o *DataTransferFromVaultSummary) GetNumPhysicalBytesTransferred() int64`

GetNumPhysicalBytesTransferred returns the NumPhysicalBytesTransferred field if non-nil, zero value otherwise.

### GetNumPhysicalBytesTransferredOk

`func (o *DataTransferFromVaultSummary) GetNumPhysicalBytesTransferredOk() (*int64, bool)`

GetNumPhysicalBytesTransferredOk returns a tuple with the NumPhysicalBytesTransferred field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumPhysicalBytesTransferred

`func (o *DataTransferFromVaultSummary) SetNumPhysicalBytesTransferred(v int64)`

SetNumPhysicalBytesTransferred sets NumPhysicalBytesTransferred field to given value.

### HasNumPhysicalBytesTransferred

`func (o *DataTransferFromVaultSummary) HasNumPhysicalBytesTransferred() bool`

HasNumPhysicalBytesTransferred returns a boolean if a field has been set.

### SetNumPhysicalBytesTransferredNil

`func (o *DataTransferFromVaultSummary) SetNumPhysicalBytesTransferredNil(b bool)`

 SetNumPhysicalBytesTransferredNil sets the value for NumPhysicalBytesTransferred to be an explicit nil

### UnsetNumPhysicalBytesTransferred
`func (o *DataTransferFromVaultSummary) UnsetNumPhysicalBytesTransferred()`

UnsetNumPhysicalBytesTransferred ensures that no value is present for NumPhysicalBytesTransferred, not even an explicit nil
### GetNumTasks

`func (o *DataTransferFromVaultSummary) GetNumTasks() int64`

GetNumTasks returns the NumTasks field if non-nil, zero value otherwise.

### GetNumTasksOk

`func (o *DataTransferFromVaultSummary) GetNumTasksOk() (*int64, bool)`

GetNumTasksOk returns a tuple with the NumTasks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumTasks

`func (o *DataTransferFromVaultSummary) SetNumTasks(v int64)`

SetNumTasks sets NumTasks field to given value.

### HasNumTasks

`func (o *DataTransferFromVaultSummary) HasNumTasks() bool`

HasNumTasks returns a boolean if a field has been set.

### SetNumTasksNil

`func (o *DataTransferFromVaultSummary) SetNumTasksNil(b bool)`

 SetNumTasksNil sets the value for NumTasks to be an explicit nil

### UnsetNumTasks
`func (o *DataTransferFromVaultSummary) UnsetNumTasks()`

UnsetNumTasks ensures that no value is present for NumTasks, not even an explicit nil
### GetPhysicalDataTransferredBytesDuringTimeRange

`func (o *DataTransferFromVaultSummary) GetPhysicalDataTransferredBytesDuringTimeRange() []int64`

GetPhysicalDataTransferredBytesDuringTimeRange returns the PhysicalDataTransferredBytesDuringTimeRange field if non-nil, zero value otherwise.

### GetPhysicalDataTransferredBytesDuringTimeRangeOk

`func (o *DataTransferFromVaultSummary) GetPhysicalDataTransferredBytesDuringTimeRangeOk() (*[]int64, bool)`

GetPhysicalDataTransferredBytesDuringTimeRangeOk returns a tuple with the PhysicalDataTransferredBytesDuringTimeRange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhysicalDataTransferredBytesDuringTimeRange

`func (o *DataTransferFromVaultSummary) SetPhysicalDataTransferredBytesDuringTimeRange(v []int64)`

SetPhysicalDataTransferredBytesDuringTimeRange sets PhysicalDataTransferredBytesDuringTimeRange field to given value.

### HasPhysicalDataTransferredBytesDuringTimeRange

`func (o *DataTransferFromVaultSummary) HasPhysicalDataTransferredBytesDuringTimeRange() bool`

HasPhysicalDataTransferredBytesDuringTimeRange returns a boolean if a field has been set.

### SetPhysicalDataTransferredBytesDuringTimeRangeNil

`func (o *DataTransferFromVaultSummary) SetPhysicalDataTransferredBytesDuringTimeRangeNil(b bool)`

 SetPhysicalDataTransferredBytesDuringTimeRangeNil sets the value for PhysicalDataTransferredBytesDuringTimeRange to be an explicit nil

### UnsetPhysicalDataTransferredBytesDuringTimeRange
`func (o *DataTransferFromVaultSummary) UnsetPhysicalDataTransferredBytesDuringTimeRange()`

UnsetPhysicalDataTransferredBytesDuringTimeRange ensures that no value is present for PhysicalDataTransferredBytesDuringTimeRange, not even an explicit nil
### GetVaultName

`func (o *DataTransferFromVaultSummary) GetVaultName() string`

GetVaultName returns the VaultName field if non-nil, zero value otherwise.

### GetVaultNameOk

`func (o *DataTransferFromVaultSummary) GetVaultNameOk() (*string, bool)`

GetVaultNameOk returns a tuple with the VaultName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVaultName

`func (o *DataTransferFromVaultSummary) SetVaultName(v string)`

SetVaultName sets VaultName field to given value.

### HasVaultName

`func (o *DataTransferFromVaultSummary) HasVaultName() bool`

HasVaultName returns a boolean if a field has been set.

### SetVaultNameNil

`func (o *DataTransferFromVaultSummary) SetVaultNameNil(b bool)`

 SetVaultNameNil sets the value for VaultName to be an explicit nil

### UnsetVaultName
`func (o *DataTransferFromVaultSummary) UnsetVaultName()`

UnsetVaultName ensures that no value is present for VaultName, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


