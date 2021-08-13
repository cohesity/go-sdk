# CancelObjectRunParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RunId** | **NullableString** | Specifies the id of the run to cancel. | 
**CancelLocalRun** | Pointer to **NullableBool** | Specifies whether to cancel the local backup run. Default is false. | [optional] 
**ArchivalTargetIds** | Pointer to **[]int64** | Specifies the archival target ids where the tasks run. If specified, the archival target ids must be present within the run specified by the runId above. | [optional] 
**ReplicationTargets** | Pointer to [**[]ClusterIdentifier**](ClusterIdentifier.md) | Specifies the cluster identifiers where the tasks run. If specified, the archival target ids must be present within the run specified by the runId above. | [optional] 
**CloudSpinTargetIds** | Pointer to **[]int64** | Specifies the cloud spin target ids where the tasks run. If specified, the archival target ids must be present within the run specified by the runId above. | [optional] 

## Methods

### NewCancelObjectRunParams

`func NewCancelObjectRunParams(runId NullableString, ) *CancelObjectRunParams`

NewCancelObjectRunParams instantiates a new CancelObjectRunParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCancelObjectRunParamsWithDefaults

`func NewCancelObjectRunParamsWithDefaults() *CancelObjectRunParams`

NewCancelObjectRunParamsWithDefaults instantiates a new CancelObjectRunParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRunId

`func (o *CancelObjectRunParams) GetRunId() string`

GetRunId returns the RunId field if non-nil, zero value otherwise.

### GetRunIdOk

`func (o *CancelObjectRunParams) GetRunIdOk() (*string, bool)`

GetRunIdOk returns a tuple with the RunId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRunId

`func (o *CancelObjectRunParams) SetRunId(v string)`

SetRunId sets RunId field to given value.


### SetRunIdNil

`func (o *CancelObjectRunParams) SetRunIdNil(b bool)`

 SetRunIdNil sets the value for RunId to be an explicit nil

### UnsetRunId
`func (o *CancelObjectRunParams) UnsetRunId()`

UnsetRunId ensures that no value is present for RunId, not even an explicit nil
### GetCancelLocalRun

`func (o *CancelObjectRunParams) GetCancelLocalRun() bool`

GetCancelLocalRun returns the CancelLocalRun field if non-nil, zero value otherwise.

### GetCancelLocalRunOk

`func (o *CancelObjectRunParams) GetCancelLocalRunOk() (*bool, bool)`

GetCancelLocalRunOk returns a tuple with the CancelLocalRun field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCancelLocalRun

`func (o *CancelObjectRunParams) SetCancelLocalRun(v bool)`

SetCancelLocalRun sets CancelLocalRun field to given value.

### HasCancelLocalRun

`func (o *CancelObjectRunParams) HasCancelLocalRun() bool`

HasCancelLocalRun returns a boolean if a field has been set.

### SetCancelLocalRunNil

`func (o *CancelObjectRunParams) SetCancelLocalRunNil(b bool)`

 SetCancelLocalRunNil sets the value for CancelLocalRun to be an explicit nil

### UnsetCancelLocalRun
`func (o *CancelObjectRunParams) UnsetCancelLocalRun()`

UnsetCancelLocalRun ensures that no value is present for CancelLocalRun, not even an explicit nil
### GetArchivalTargetIds

`func (o *CancelObjectRunParams) GetArchivalTargetIds() []int64`

GetArchivalTargetIds returns the ArchivalTargetIds field if non-nil, zero value otherwise.

### GetArchivalTargetIdsOk

`func (o *CancelObjectRunParams) GetArchivalTargetIdsOk() (*[]int64, bool)`

GetArchivalTargetIdsOk returns a tuple with the ArchivalTargetIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchivalTargetIds

`func (o *CancelObjectRunParams) SetArchivalTargetIds(v []int64)`

SetArchivalTargetIds sets ArchivalTargetIds field to given value.

### HasArchivalTargetIds

`func (o *CancelObjectRunParams) HasArchivalTargetIds() bool`

HasArchivalTargetIds returns a boolean if a field has been set.

### SetArchivalTargetIdsNil

`func (o *CancelObjectRunParams) SetArchivalTargetIdsNil(b bool)`

 SetArchivalTargetIdsNil sets the value for ArchivalTargetIds to be an explicit nil

### UnsetArchivalTargetIds
`func (o *CancelObjectRunParams) UnsetArchivalTargetIds()`

UnsetArchivalTargetIds ensures that no value is present for ArchivalTargetIds, not even an explicit nil
### GetReplicationTargets

`func (o *CancelObjectRunParams) GetReplicationTargets() []ClusterIdentifier`

GetReplicationTargets returns the ReplicationTargets field if non-nil, zero value otherwise.

### GetReplicationTargetsOk

`func (o *CancelObjectRunParams) GetReplicationTargetsOk() (*[]ClusterIdentifier, bool)`

GetReplicationTargetsOk returns a tuple with the ReplicationTargets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplicationTargets

`func (o *CancelObjectRunParams) SetReplicationTargets(v []ClusterIdentifier)`

SetReplicationTargets sets ReplicationTargets field to given value.

### HasReplicationTargets

`func (o *CancelObjectRunParams) HasReplicationTargets() bool`

HasReplicationTargets returns a boolean if a field has been set.

### SetReplicationTargetsNil

`func (o *CancelObjectRunParams) SetReplicationTargetsNil(b bool)`

 SetReplicationTargetsNil sets the value for ReplicationTargets to be an explicit nil

### UnsetReplicationTargets
`func (o *CancelObjectRunParams) UnsetReplicationTargets()`

UnsetReplicationTargets ensures that no value is present for ReplicationTargets, not even an explicit nil
### GetCloudSpinTargetIds

`func (o *CancelObjectRunParams) GetCloudSpinTargetIds() []int64`

GetCloudSpinTargetIds returns the CloudSpinTargetIds field if non-nil, zero value otherwise.

### GetCloudSpinTargetIdsOk

`func (o *CancelObjectRunParams) GetCloudSpinTargetIdsOk() (*[]int64, bool)`

GetCloudSpinTargetIdsOk returns a tuple with the CloudSpinTargetIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudSpinTargetIds

`func (o *CancelObjectRunParams) SetCloudSpinTargetIds(v []int64)`

SetCloudSpinTargetIds sets CloudSpinTargetIds field to given value.

### HasCloudSpinTargetIds

`func (o *CancelObjectRunParams) HasCloudSpinTargetIds() bool`

HasCloudSpinTargetIds returns a boolean if a field has been set.

### SetCloudSpinTargetIdsNil

`func (o *CancelObjectRunParams) SetCloudSpinTargetIdsNil(b bool)`

 SetCloudSpinTargetIdsNil sets the value for CloudSpinTargetIds to be an explicit nil

### UnsetCloudSpinTargetIds
`func (o *CancelObjectRunParams) UnsetCloudSpinTargetIds()`

UnsetCloudSpinTargetIds ensures that no value is present for CloudSpinTargetIds, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


