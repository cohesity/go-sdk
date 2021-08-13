# ReplicationDataStats

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LogicalSizeBytes** | Pointer to **NullableInt64** | Specifies the total logical size in bytes. | [optional] 
**LogicalBytesTransferred** | Pointer to **NullableInt64** | Specifies the total logical bytes transferred. | [optional] 
**PhysicalBytesTransferred** | Pointer to **NullableInt64** | Specifies the total physical bytes transferred. | [optional] 

## Methods

### NewReplicationDataStats

`func NewReplicationDataStats() *ReplicationDataStats`

NewReplicationDataStats instantiates a new ReplicationDataStats object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReplicationDataStatsWithDefaults

`func NewReplicationDataStatsWithDefaults() *ReplicationDataStats`

NewReplicationDataStatsWithDefaults instantiates a new ReplicationDataStats object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLogicalSizeBytes

`func (o *ReplicationDataStats) GetLogicalSizeBytes() int64`

GetLogicalSizeBytes returns the LogicalSizeBytes field if non-nil, zero value otherwise.

### GetLogicalSizeBytesOk

`func (o *ReplicationDataStats) GetLogicalSizeBytesOk() (*int64, bool)`

GetLogicalSizeBytesOk returns a tuple with the LogicalSizeBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogicalSizeBytes

`func (o *ReplicationDataStats) SetLogicalSizeBytes(v int64)`

SetLogicalSizeBytes sets LogicalSizeBytes field to given value.

### HasLogicalSizeBytes

`func (o *ReplicationDataStats) HasLogicalSizeBytes() bool`

HasLogicalSizeBytes returns a boolean if a field has been set.

### SetLogicalSizeBytesNil

`func (o *ReplicationDataStats) SetLogicalSizeBytesNil(b bool)`

 SetLogicalSizeBytesNil sets the value for LogicalSizeBytes to be an explicit nil

### UnsetLogicalSizeBytes
`func (o *ReplicationDataStats) UnsetLogicalSizeBytes()`

UnsetLogicalSizeBytes ensures that no value is present for LogicalSizeBytes, not even an explicit nil
### GetLogicalBytesTransferred

`func (o *ReplicationDataStats) GetLogicalBytesTransferred() int64`

GetLogicalBytesTransferred returns the LogicalBytesTransferred field if non-nil, zero value otherwise.

### GetLogicalBytesTransferredOk

`func (o *ReplicationDataStats) GetLogicalBytesTransferredOk() (*int64, bool)`

GetLogicalBytesTransferredOk returns a tuple with the LogicalBytesTransferred field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogicalBytesTransferred

`func (o *ReplicationDataStats) SetLogicalBytesTransferred(v int64)`

SetLogicalBytesTransferred sets LogicalBytesTransferred field to given value.

### HasLogicalBytesTransferred

`func (o *ReplicationDataStats) HasLogicalBytesTransferred() bool`

HasLogicalBytesTransferred returns a boolean if a field has been set.

### SetLogicalBytesTransferredNil

`func (o *ReplicationDataStats) SetLogicalBytesTransferredNil(b bool)`

 SetLogicalBytesTransferredNil sets the value for LogicalBytesTransferred to be an explicit nil

### UnsetLogicalBytesTransferred
`func (o *ReplicationDataStats) UnsetLogicalBytesTransferred()`

UnsetLogicalBytesTransferred ensures that no value is present for LogicalBytesTransferred, not even an explicit nil
### GetPhysicalBytesTransferred

`func (o *ReplicationDataStats) GetPhysicalBytesTransferred() int64`

GetPhysicalBytesTransferred returns the PhysicalBytesTransferred field if non-nil, zero value otherwise.

### GetPhysicalBytesTransferredOk

`func (o *ReplicationDataStats) GetPhysicalBytesTransferredOk() (*int64, bool)`

GetPhysicalBytesTransferredOk returns a tuple with the PhysicalBytesTransferred field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhysicalBytesTransferred

`func (o *ReplicationDataStats) SetPhysicalBytesTransferred(v int64)`

SetPhysicalBytesTransferred sets PhysicalBytesTransferred field to given value.

### HasPhysicalBytesTransferred

`func (o *ReplicationDataStats) HasPhysicalBytesTransferred() bool`

HasPhysicalBytesTransferred returns a boolean if a field has been set.

### SetPhysicalBytesTransferredNil

`func (o *ReplicationDataStats) SetPhysicalBytesTransferredNil(b bool)`

 SetPhysicalBytesTransferredNil sets the value for PhysicalBytesTransferred to be an explicit nil

### UnsetPhysicalBytesTransferred
`func (o *ReplicationDataStats) UnsetPhysicalBytesTransferred()`

UnsetPhysicalBytesTransferred ensures that no value is present for PhysicalBytesTransferred, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


