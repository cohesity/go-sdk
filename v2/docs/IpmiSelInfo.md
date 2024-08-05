# IpmiSelInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AllocationUnitSize** | Pointer to **NullableString** | Specifies the size of each allocation unit in bytes. | [optional] 
**AllocationUnits** | Pointer to **NullableString** | Specifies the number of allocation units available in SEL. | [optional] 
**Entries** | Pointer to **NullableString** | Specifies the number of log entries stores in SEL. | [optional] 
**FreeSpace** | Pointer to **NullableString** | Specifies the number of free bytes in SEL. | [optional] 
**FreeUnits** | Pointer to **NullableString** | Specifies the number of free allocation units present in SEL. | [optional] 
**LargestFreeBlk** | Pointer to **NullableString** | Specifies the size of the largest contiguous block of free space available in the SEL. | [optional] 
**LargestRecordSize** | Pointer to **NullableString** | Specifies the maximum size of a single log record that can be stored in the SEL measured in bytes. | [optional] 
**LastAdditionTime** | Pointer to **NullableString** | Specifies the latest time stamp at which a log entry was added to SEL. | [optional] 
**LastDeletionTime** | Pointer to **NullableString** | Specifies the latest time stamp at which a log entry was deleted from SEL. | [optional] 
**Overflow** | Pointer to **NullableString** | Specifies whether an overflow has occured in SEL. | [optional] 
**PercentUsed** | Pointer to **NullableString** | Specifies the percentage of SEL used by log entries. | [optional] 
**SupportedCommands** | Pointer to **NullableString** | Specifies a space seperated list of commands that are supported for managing the SEL. | [optional] 
**Version** | Pointer to **NullableString** | Specifies the SEL(System Event Log) version for given node. | [optional] 

## Methods

### NewIpmiSelInfo

`func NewIpmiSelInfo() *IpmiSelInfo`

NewIpmiSelInfo instantiates a new IpmiSelInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIpmiSelInfoWithDefaults

`func NewIpmiSelInfoWithDefaults() *IpmiSelInfo`

NewIpmiSelInfoWithDefaults instantiates a new IpmiSelInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAllocationUnitSize

`func (o *IpmiSelInfo) GetAllocationUnitSize() string`

GetAllocationUnitSize returns the AllocationUnitSize field if non-nil, zero value otherwise.

### GetAllocationUnitSizeOk

`func (o *IpmiSelInfo) GetAllocationUnitSizeOk() (*string, bool)`

GetAllocationUnitSizeOk returns a tuple with the AllocationUnitSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllocationUnitSize

`func (o *IpmiSelInfo) SetAllocationUnitSize(v string)`

SetAllocationUnitSize sets AllocationUnitSize field to given value.

### HasAllocationUnitSize

`func (o *IpmiSelInfo) HasAllocationUnitSize() bool`

HasAllocationUnitSize returns a boolean if a field has been set.

### SetAllocationUnitSizeNil

`func (o *IpmiSelInfo) SetAllocationUnitSizeNil(b bool)`

 SetAllocationUnitSizeNil sets the value for AllocationUnitSize to be an explicit nil

### UnsetAllocationUnitSize
`func (o *IpmiSelInfo) UnsetAllocationUnitSize()`

UnsetAllocationUnitSize ensures that no value is present for AllocationUnitSize, not even an explicit nil
### GetAllocationUnits

`func (o *IpmiSelInfo) GetAllocationUnits() string`

GetAllocationUnits returns the AllocationUnits field if non-nil, zero value otherwise.

### GetAllocationUnitsOk

`func (o *IpmiSelInfo) GetAllocationUnitsOk() (*string, bool)`

GetAllocationUnitsOk returns a tuple with the AllocationUnits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllocationUnits

`func (o *IpmiSelInfo) SetAllocationUnits(v string)`

SetAllocationUnits sets AllocationUnits field to given value.

### HasAllocationUnits

`func (o *IpmiSelInfo) HasAllocationUnits() bool`

HasAllocationUnits returns a boolean if a field has been set.

### SetAllocationUnitsNil

`func (o *IpmiSelInfo) SetAllocationUnitsNil(b bool)`

 SetAllocationUnitsNil sets the value for AllocationUnits to be an explicit nil

### UnsetAllocationUnits
`func (o *IpmiSelInfo) UnsetAllocationUnits()`

UnsetAllocationUnits ensures that no value is present for AllocationUnits, not even an explicit nil
### GetEntries

`func (o *IpmiSelInfo) GetEntries() string`

GetEntries returns the Entries field if non-nil, zero value otherwise.

### GetEntriesOk

`func (o *IpmiSelInfo) GetEntriesOk() (*string, bool)`

GetEntriesOk returns a tuple with the Entries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntries

`func (o *IpmiSelInfo) SetEntries(v string)`

SetEntries sets Entries field to given value.

### HasEntries

`func (o *IpmiSelInfo) HasEntries() bool`

HasEntries returns a boolean if a field has been set.

### SetEntriesNil

`func (o *IpmiSelInfo) SetEntriesNil(b bool)`

 SetEntriesNil sets the value for Entries to be an explicit nil

### UnsetEntries
`func (o *IpmiSelInfo) UnsetEntries()`

UnsetEntries ensures that no value is present for Entries, not even an explicit nil
### GetFreeSpace

`func (o *IpmiSelInfo) GetFreeSpace() string`

GetFreeSpace returns the FreeSpace field if non-nil, zero value otherwise.

### GetFreeSpaceOk

`func (o *IpmiSelInfo) GetFreeSpaceOk() (*string, bool)`

GetFreeSpaceOk returns a tuple with the FreeSpace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFreeSpace

`func (o *IpmiSelInfo) SetFreeSpace(v string)`

SetFreeSpace sets FreeSpace field to given value.

### HasFreeSpace

`func (o *IpmiSelInfo) HasFreeSpace() bool`

HasFreeSpace returns a boolean if a field has been set.

### SetFreeSpaceNil

`func (o *IpmiSelInfo) SetFreeSpaceNil(b bool)`

 SetFreeSpaceNil sets the value for FreeSpace to be an explicit nil

### UnsetFreeSpace
`func (o *IpmiSelInfo) UnsetFreeSpace()`

UnsetFreeSpace ensures that no value is present for FreeSpace, not even an explicit nil
### GetFreeUnits

`func (o *IpmiSelInfo) GetFreeUnits() string`

GetFreeUnits returns the FreeUnits field if non-nil, zero value otherwise.

### GetFreeUnitsOk

`func (o *IpmiSelInfo) GetFreeUnitsOk() (*string, bool)`

GetFreeUnitsOk returns a tuple with the FreeUnits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFreeUnits

`func (o *IpmiSelInfo) SetFreeUnits(v string)`

SetFreeUnits sets FreeUnits field to given value.

### HasFreeUnits

`func (o *IpmiSelInfo) HasFreeUnits() bool`

HasFreeUnits returns a boolean if a field has been set.

### SetFreeUnitsNil

`func (o *IpmiSelInfo) SetFreeUnitsNil(b bool)`

 SetFreeUnitsNil sets the value for FreeUnits to be an explicit nil

### UnsetFreeUnits
`func (o *IpmiSelInfo) UnsetFreeUnits()`

UnsetFreeUnits ensures that no value is present for FreeUnits, not even an explicit nil
### GetLargestFreeBlk

`func (o *IpmiSelInfo) GetLargestFreeBlk() string`

GetLargestFreeBlk returns the LargestFreeBlk field if non-nil, zero value otherwise.

### GetLargestFreeBlkOk

`func (o *IpmiSelInfo) GetLargestFreeBlkOk() (*string, bool)`

GetLargestFreeBlkOk returns a tuple with the LargestFreeBlk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLargestFreeBlk

`func (o *IpmiSelInfo) SetLargestFreeBlk(v string)`

SetLargestFreeBlk sets LargestFreeBlk field to given value.

### HasLargestFreeBlk

`func (o *IpmiSelInfo) HasLargestFreeBlk() bool`

HasLargestFreeBlk returns a boolean if a field has been set.

### SetLargestFreeBlkNil

`func (o *IpmiSelInfo) SetLargestFreeBlkNil(b bool)`

 SetLargestFreeBlkNil sets the value for LargestFreeBlk to be an explicit nil

### UnsetLargestFreeBlk
`func (o *IpmiSelInfo) UnsetLargestFreeBlk()`

UnsetLargestFreeBlk ensures that no value is present for LargestFreeBlk, not even an explicit nil
### GetLargestRecordSize

`func (o *IpmiSelInfo) GetLargestRecordSize() string`

GetLargestRecordSize returns the LargestRecordSize field if non-nil, zero value otherwise.

### GetLargestRecordSizeOk

`func (o *IpmiSelInfo) GetLargestRecordSizeOk() (*string, bool)`

GetLargestRecordSizeOk returns a tuple with the LargestRecordSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLargestRecordSize

`func (o *IpmiSelInfo) SetLargestRecordSize(v string)`

SetLargestRecordSize sets LargestRecordSize field to given value.

### HasLargestRecordSize

`func (o *IpmiSelInfo) HasLargestRecordSize() bool`

HasLargestRecordSize returns a boolean if a field has been set.

### SetLargestRecordSizeNil

`func (o *IpmiSelInfo) SetLargestRecordSizeNil(b bool)`

 SetLargestRecordSizeNil sets the value for LargestRecordSize to be an explicit nil

### UnsetLargestRecordSize
`func (o *IpmiSelInfo) UnsetLargestRecordSize()`

UnsetLargestRecordSize ensures that no value is present for LargestRecordSize, not even an explicit nil
### GetLastAdditionTime

`func (o *IpmiSelInfo) GetLastAdditionTime() string`

GetLastAdditionTime returns the LastAdditionTime field if non-nil, zero value otherwise.

### GetLastAdditionTimeOk

`func (o *IpmiSelInfo) GetLastAdditionTimeOk() (*string, bool)`

GetLastAdditionTimeOk returns a tuple with the LastAdditionTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastAdditionTime

`func (o *IpmiSelInfo) SetLastAdditionTime(v string)`

SetLastAdditionTime sets LastAdditionTime field to given value.

### HasLastAdditionTime

`func (o *IpmiSelInfo) HasLastAdditionTime() bool`

HasLastAdditionTime returns a boolean if a field has been set.

### SetLastAdditionTimeNil

`func (o *IpmiSelInfo) SetLastAdditionTimeNil(b bool)`

 SetLastAdditionTimeNil sets the value for LastAdditionTime to be an explicit nil

### UnsetLastAdditionTime
`func (o *IpmiSelInfo) UnsetLastAdditionTime()`

UnsetLastAdditionTime ensures that no value is present for LastAdditionTime, not even an explicit nil
### GetLastDeletionTime

`func (o *IpmiSelInfo) GetLastDeletionTime() string`

GetLastDeletionTime returns the LastDeletionTime field if non-nil, zero value otherwise.

### GetLastDeletionTimeOk

`func (o *IpmiSelInfo) GetLastDeletionTimeOk() (*string, bool)`

GetLastDeletionTimeOk returns a tuple with the LastDeletionTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastDeletionTime

`func (o *IpmiSelInfo) SetLastDeletionTime(v string)`

SetLastDeletionTime sets LastDeletionTime field to given value.

### HasLastDeletionTime

`func (o *IpmiSelInfo) HasLastDeletionTime() bool`

HasLastDeletionTime returns a boolean if a field has been set.

### SetLastDeletionTimeNil

`func (o *IpmiSelInfo) SetLastDeletionTimeNil(b bool)`

 SetLastDeletionTimeNil sets the value for LastDeletionTime to be an explicit nil

### UnsetLastDeletionTime
`func (o *IpmiSelInfo) UnsetLastDeletionTime()`

UnsetLastDeletionTime ensures that no value is present for LastDeletionTime, not even an explicit nil
### GetOverflow

`func (o *IpmiSelInfo) GetOverflow() string`

GetOverflow returns the Overflow field if non-nil, zero value otherwise.

### GetOverflowOk

`func (o *IpmiSelInfo) GetOverflowOk() (*string, bool)`

GetOverflowOk returns a tuple with the Overflow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverflow

`func (o *IpmiSelInfo) SetOverflow(v string)`

SetOverflow sets Overflow field to given value.

### HasOverflow

`func (o *IpmiSelInfo) HasOverflow() bool`

HasOverflow returns a boolean if a field has been set.

### SetOverflowNil

`func (o *IpmiSelInfo) SetOverflowNil(b bool)`

 SetOverflowNil sets the value for Overflow to be an explicit nil

### UnsetOverflow
`func (o *IpmiSelInfo) UnsetOverflow()`

UnsetOverflow ensures that no value is present for Overflow, not even an explicit nil
### GetPercentUsed

`func (o *IpmiSelInfo) GetPercentUsed() string`

GetPercentUsed returns the PercentUsed field if non-nil, zero value otherwise.

### GetPercentUsedOk

`func (o *IpmiSelInfo) GetPercentUsedOk() (*string, bool)`

GetPercentUsedOk returns a tuple with the PercentUsed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPercentUsed

`func (o *IpmiSelInfo) SetPercentUsed(v string)`

SetPercentUsed sets PercentUsed field to given value.

### HasPercentUsed

`func (o *IpmiSelInfo) HasPercentUsed() bool`

HasPercentUsed returns a boolean if a field has been set.

### SetPercentUsedNil

`func (o *IpmiSelInfo) SetPercentUsedNil(b bool)`

 SetPercentUsedNil sets the value for PercentUsed to be an explicit nil

### UnsetPercentUsed
`func (o *IpmiSelInfo) UnsetPercentUsed()`

UnsetPercentUsed ensures that no value is present for PercentUsed, not even an explicit nil
### GetSupportedCommands

`func (o *IpmiSelInfo) GetSupportedCommands() string`

GetSupportedCommands returns the SupportedCommands field if non-nil, zero value otherwise.

### GetSupportedCommandsOk

`func (o *IpmiSelInfo) GetSupportedCommandsOk() (*string, bool)`

GetSupportedCommandsOk returns a tuple with the SupportedCommands field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportedCommands

`func (o *IpmiSelInfo) SetSupportedCommands(v string)`

SetSupportedCommands sets SupportedCommands field to given value.

### HasSupportedCommands

`func (o *IpmiSelInfo) HasSupportedCommands() bool`

HasSupportedCommands returns a boolean if a field has been set.

### SetSupportedCommandsNil

`func (o *IpmiSelInfo) SetSupportedCommandsNil(b bool)`

 SetSupportedCommandsNil sets the value for SupportedCommands to be an explicit nil

### UnsetSupportedCommands
`func (o *IpmiSelInfo) UnsetSupportedCommands()`

UnsetSupportedCommands ensures that no value is present for SupportedCommands, not even an explicit nil
### GetVersion

`func (o *IpmiSelInfo) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *IpmiSelInfo) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *IpmiSelInfo) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *IpmiSelInfo) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### SetVersionNil

`func (o *IpmiSelInfo) SetVersionNil(b bool)`

 SetVersionNil sets the value for Version to be an explicit nil

### UnsetVersion
`func (o *IpmiSelInfo) UnsetVersion()`

UnsetVersion ensures that no value is present for Version, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


