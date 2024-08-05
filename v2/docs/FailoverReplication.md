# FailoverReplication

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EndTimeUsecs** | Pointer to **NullableInt64** | Specifies the replication complete time in micro seconds. | [optional] 
**ErrorMessage** | Pointer to **NullableString** | Specifies the error details if replication status is &#39;Failed&#39;. | [optional] 
**Id** | Pointer to **NullableString** | Specifies the replication id. | [optional] 
**LogicalBytesTransferred** | Pointer to **NullableInt64** | Specifies the number of logical bytes transferred for this replication so far. This value can never exceed the total logical size of the replicated view. | [optional] 
**LogicalSizeBytes** | Pointer to **NullableInt64** | Specifies the total amount of logical data to be transferred for this replication. | [optional] 
**PercentageCompleted** | Pointer to **NullableInt32** | Specifies the percentage completed in the replication. | [optional] 
**PhysicalBytesTransferred** | Pointer to **NullableInt64** | Specifies the number of bytes sent over the wire for this replication so far. | [optional] 
**StartTimeUsecs** | Pointer to **NullableInt64** | Specifies the replication start time in micro seconds. | [optional] 
**Status** | Pointer to **NullableString** | Specifies the replication status. | [optional] 
**TargetClusterId** | Pointer to **NullableInt64** | Specifies the failover target cluster id. | [optional] 
**TargetClusterIncarnationId** | Pointer to **NullableInt64** | Specifies the failover target cluster incarnation id. | [optional] 
**TargetClusterName** | Pointer to **NullableString** | Specifies the failover target cluster name. | [optional] 

## Methods

### NewFailoverReplication

`func NewFailoverReplication() *FailoverReplication`

NewFailoverReplication instantiates a new FailoverReplication object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFailoverReplicationWithDefaults

`func NewFailoverReplicationWithDefaults() *FailoverReplication`

NewFailoverReplicationWithDefaults instantiates a new FailoverReplication object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEndTimeUsecs

`func (o *FailoverReplication) GetEndTimeUsecs() int64`

GetEndTimeUsecs returns the EndTimeUsecs field if non-nil, zero value otherwise.

### GetEndTimeUsecsOk

`func (o *FailoverReplication) GetEndTimeUsecsOk() (*int64, bool)`

GetEndTimeUsecsOk returns a tuple with the EndTimeUsecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndTimeUsecs

`func (o *FailoverReplication) SetEndTimeUsecs(v int64)`

SetEndTimeUsecs sets EndTimeUsecs field to given value.

### HasEndTimeUsecs

`func (o *FailoverReplication) HasEndTimeUsecs() bool`

HasEndTimeUsecs returns a boolean if a field has been set.

### SetEndTimeUsecsNil

`func (o *FailoverReplication) SetEndTimeUsecsNil(b bool)`

 SetEndTimeUsecsNil sets the value for EndTimeUsecs to be an explicit nil

### UnsetEndTimeUsecs
`func (o *FailoverReplication) UnsetEndTimeUsecs()`

UnsetEndTimeUsecs ensures that no value is present for EndTimeUsecs, not even an explicit nil
### GetErrorMessage

`func (o *FailoverReplication) GetErrorMessage() string`

GetErrorMessage returns the ErrorMessage field if non-nil, zero value otherwise.

### GetErrorMessageOk

`func (o *FailoverReplication) GetErrorMessageOk() (*string, bool)`

GetErrorMessageOk returns a tuple with the ErrorMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorMessage

`func (o *FailoverReplication) SetErrorMessage(v string)`

SetErrorMessage sets ErrorMessage field to given value.

### HasErrorMessage

`func (o *FailoverReplication) HasErrorMessage() bool`

HasErrorMessage returns a boolean if a field has been set.

### SetErrorMessageNil

`func (o *FailoverReplication) SetErrorMessageNil(b bool)`

 SetErrorMessageNil sets the value for ErrorMessage to be an explicit nil

### UnsetErrorMessage
`func (o *FailoverReplication) UnsetErrorMessage()`

UnsetErrorMessage ensures that no value is present for ErrorMessage, not even an explicit nil
### GetId

`func (o *FailoverReplication) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *FailoverReplication) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *FailoverReplication) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *FailoverReplication) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *FailoverReplication) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *FailoverReplication) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetLogicalBytesTransferred

`func (o *FailoverReplication) GetLogicalBytesTransferred() int64`

GetLogicalBytesTransferred returns the LogicalBytesTransferred field if non-nil, zero value otherwise.

### GetLogicalBytesTransferredOk

`func (o *FailoverReplication) GetLogicalBytesTransferredOk() (*int64, bool)`

GetLogicalBytesTransferredOk returns a tuple with the LogicalBytesTransferred field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogicalBytesTransferred

`func (o *FailoverReplication) SetLogicalBytesTransferred(v int64)`

SetLogicalBytesTransferred sets LogicalBytesTransferred field to given value.

### HasLogicalBytesTransferred

`func (o *FailoverReplication) HasLogicalBytesTransferred() bool`

HasLogicalBytesTransferred returns a boolean if a field has been set.

### SetLogicalBytesTransferredNil

`func (o *FailoverReplication) SetLogicalBytesTransferredNil(b bool)`

 SetLogicalBytesTransferredNil sets the value for LogicalBytesTransferred to be an explicit nil

### UnsetLogicalBytesTransferred
`func (o *FailoverReplication) UnsetLogicalBytesTransferred()`

UnsetLogicalBytesTransferred ensures that no value is present for LogicalBytesTransferred, not even an explicit nil
### GetLogicalSizeBytes

`func (o *FailoverReplication) GetLogicalSizeBytes() int64`

GetLogicalSizeBytes returns the LogicalSizeBytes field if non-nil, zero value otherwise.

### GetLogicalSizeBytesOk

`func (o *FailoverReplication) GetLogicalSizeBytesOk() (*int64, bool)`

GetLogicalSizeBytesOk returns a tuple with the LogicalSizeBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogicalSizeBytes

`func (o *FailoverReplication) SetLogicalSizeBytes(v int64)`

SetLogicalSizeBytes sets LogicalSizeBytes field to given value.

### HasLogicalSizeBytes

`func (o *FailoverReplication) HasLogicalSizeBytes() bool`

HasLogicalSizeBytes returns a boolean if a field has been set.

### SetLogicalSizeBytesNil

`func (o *FailoverReplication) SetLogicalSizeBytesNil(b bool)`

 SetLogicalSizeBytesNil sets the value for LogicalSizeBytes to be an explicit nil

### UnsetLogicalSizeBytes
`func (o *FailoverReplication) UnsetLogicalSizeBytes()`

UnsetLogicalSizeBytes ensures that no value is present for LogicalSizeBytes, not even an explicit nil
### GetPercentageCompleted

`func (o *FailoverReplication) GetPercentageCompleted() int32`

GetPercentageCompleted returns the PercentageCompleted field if non-nil, zero value otherwise.

### GetPercentageCompletedOk

`func (o *FailoverReplication) GetPercentageCompletedOk() (*int32, bool)`

GetPercentageCompletedOk returns a tuple with the PercentageCompleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPercentageCompleted

`func (o *FailoverReplication) SetPercentageCompleted(v int32)`

SetPercentageCompleted sets PercentageCompleted field to given value.

### HasPercentageCompleted

`func (o *FailoverReplication) HasPercentageCompleted() bool`

HasPercentageCompleted returns a boolean if a field has been set.

### SetPercentageCompletedNil

`func (o *FailoverReplication) SetPercentageCompletedNil(b bool)`

 SetPercentageCompletedNil sets the value for PercentageCompleted to be an explicit nil

### UnsetPercentageCompleted
`func (o *FailoverReplication) UnsetPercentageCompleted()`

UnsetPercentageCompleted ensures that no value is present for PercentageCompleted, not even an explicit nil
### GetPhysicalBytesTransferred

`func (o *FailoverReplication) GetPhysicalBytesTransferred() int64`

GetPhysicalBytesTransferred returns the PhysicalBytesTransferred field if non-nil, zero value otherwise.

### GetPhysicalBytesTransferredOk

`func (o *FailoverReplication) GetPhysicalBytesTransferredOk() (*int64, bool)`

GetPhysicalBytesTransferredOk returns a tuple with the PhysicalBytesTransferred field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhysicalBytesTransferred

`func (o *FailoverReplication) SetPhysicalBytesTransferred(v int64)`

SetPhysicalBytesTransferred sets PhysicalBytesTransferred field to given value.

### HasPhysicalBytesTransferred

`func (o *FailoverReplication) HasPhysicalBytesTransferred() bool`

HasPhysicalBytesTransferred returns a boolean if a field has been set.

### SetPhysicalBytesTransferredNil

`func (o *FailoverReplication) SetPhysicalBytesTransferredNil(b bool)`

 SetPhysicalBytesTransferredNil sets the value for PhysicalBytesTransferred to be an explicit nil

### UnsetPhysicalBytesTransferred
`func (o *FailoverReplication) UnsetPhysicalBytesTransferred()`

UnsetPhysicalBytesTransferred ensures that no value is present for PhysicalBytesTransferred, not even an explicit nil
### GetStartTimeUsecs

`func (o *FailoverReplication) GetStartTimeUsecs() int64`

GetStartTimeUsecs returns the StartTimeUsecs field if non-nil, zero value otherwise.

### GetStartTimeUsecsOk

`func (o *FailoverReplication) GetStartTimeUsecsOk() (*int64, bool)`

GetStartTimeUsecsOk returns a tuple with the StartTimeUsecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTimeUsecs

`func (o *FailoverReplication) SetStartTimeUsecs(v int64)`

SetStartTimeUsecs sets StartTimeUsecs field to given value.

### HasStartTimeUsecs

`func (o *FailoverReplication) HasStartTimeUsecs() bool`

HasStartTimeUsecs returns a boolean if a field has been set.

### SetStartTimeUsecsNil

`func (o *FailoverReplication) SetStartTimeUsecsNil(b bool)`

 SetStartTimeUsecsNil sets the value for StartTimeUsecs to be an explicit nil

### UnsetStartTimeUsecs
`func (o *FailoverReplication) UnsetStartTimeUsecs()`

UnsetStartTimeUsecs ensures that no value is present for StartTimeUsecs, not even an explicit nil
### GetStatus

`func (o *FailoverReplication) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *FailoverReplication) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *FailoverReplication) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *FailoverReplication) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### SetStatusNil

`func (o *FailoverReplication) SetStatusNil(b bool)`

 SetStatusNil sets the value for Status to be an explicit nil

### UnsetStatus
`func (o *FailoverReplication) UnsetStatus()`

UnsetStatus ensures that no value is present for Status, not even an explicit nil
### GetTargetClusterId

`func (o *FailoverReplication) GetTargetClusterId() int64`

GetTargetClusterId returns the TargetClusterId field if non-nil, zero value otherwise.

### GetTargetClusterIdOk

`func (o *FailoverReplication) GetTargetClusterIdOk() (*int64, bool)`

GetTargetClusterIdOk returns a tuple with the TargetClusterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetClusterId

`func (o *FailoverReplication) SetTargetClusterId(v int64)`

SetTargetClusterId sets TargetClusterId field to given value.

### HasTargetClusterId

`func (o *FailoverReplication) HasTargetClusterId() bool`

HasTargetClusterId returns a boolean if a field has been set.

### SetTargetClusterIdNil

`func (o *FailoverReplication) SetTargetClusterIdNil(b bool)`

 SetTargetClusterIdNil sets the value for TargetClusterId to be an explicit nil

### UnsetTargetClusterId
`func (o *FailoverReplication) UnsetTargetClusterId()`

UnsetTargetClusterId ensures that no value is present for TargetClusterId, not even an explicit nil
### GetTargetClusterIncarnationId

`func (o *FailoverReplication) GetTargetClusterIncarnationId() int64`

GetTargetClusterIncarnationId returns the TargetClusterIncarnationId field if non-nil, zero value otherwise.

### GetTargetClusterIncarnationIdOk

`func (o *FailoverReplication) GetTargetClusterIncarnationIdOk() (*int64, bool)`

GetTargetClusterIncarnationIdOk returns a tuple with the TargetClusterIncarnationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetClusterIncarnationId

`func (o *FailoverReplication) SetTargetClusterIncarnationId(v int64)`

SetTargetClusterIncarnationId sets TargetClusterIncarnationId field to given value.

### HasTargetClusterIncarnationId

`func (o *FailoverReplication) HasTargetClusterIncarnationId() bool`

HasTargetClusterIncarnationId returns a boolean if a field has been set.

### SetTargetClusterIncarnationIdNil

`func (o *FailoverReplication) SetTargetClusterIncarnationIdNil(b bool)`

 SetTargetClusterIncarnationIdNil sets the value for TargetClusterIncarnationId to be an explicit nil

### UnsetTargetClusterIncarnationId
`func (o *FailoverReplication) UnsetTargetClusterIncarnationId()`

UnsetTargetClusterIncarnationId ensures that no value is present for TargetClusterIncarnationId, not even an explicit nil
### GetTargetClusterName

`func (o *FailoverReplication) GetTargetClusterName() string`

GetTargetClusterName returns the TargetClusterName field if non-nil, zero value otherwise.

### GetTargetClusterNameOk

`func (o *FailoverReplication) GetTargetClusterNameOk() (*string, bool)`

GetTargetClusterNameOk returns a tuple with the TargetClusterName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetClusterName

`func (o *FailoverReplication) SetTargetClusterName(v string)`

SetTargetClusterName sets TargetClusterName field to given value.

### HasTargetClusterName

`func (o *FailoverReplication) HasTargetClusterName() bool`

HasTargetClusterName returns a boolean if a field has been set.

### SetTargetClusterNameNil

`func (o *FailoverReplication) SetTargetClusterNameNil(b bool)`

 SetTargetClusterNameNil sets the value for TargetClusterName to be an explicit nil

### UnsetTargetClusterName
`func (o *FailoverReplication) UnsetTargetClusterName()`

UnsetTargetClusterName ensures that no value is present for TargetClusterName, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


