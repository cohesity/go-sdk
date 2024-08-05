# GetProtectionRunProgressBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ArchivalRun** | Pointer to [**[]ArchivalTargetProgressInfo**](ArchivalTargetProgressInfo.md) | Progress for the archival run. | [optional] 
**LocalRun** | Pointer to [**BackupRunProgressInfo**](BackupRunProgressInfo.md) |  | [optional] 
**ReplicationRun** | Pointer to [**[]ReplicationTargetProgressInfo**](ReplicationTargetProgressInfo.md) | Progress for the replication run. | [optional] 

## Methods

### NewGetProtectionRunProgressBody

`func NewGetProtectionRunProgressBody() *GetProtectionRunProgressBody`

NewGetProtectionRunProgressBody instantiates a new GetProtectionRunProgressBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetProtectionRunProgressBodyWithDefaults

`func NewGetProtectionRunProgressBodyWithDefaults() *GetProtectionRunProgressBody`

NewGetProtectionRunProgressBodyWithDefaults instantiates a new GetProtectionRunProgressBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetArchivalRun

`func (o *GetProtectionRunProgressBody) GetArchivalRun() []ArchivalTargetProgressInfo`

GetArchivalRun returns the ArchivalRun field if non-nil, zero value otherwise.

### GetArchivalRunOk

`func (o *GetProtectionRunProgressBody) GetArchivalRunOk() (*[]ArchivalTargetProgressInfo, bool)`

GetArchivalRunOk returns a tuple with the ArchivalRun field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchivalRun

`func (o *GetProtectionRunProgressBody) SetArchivalRun(v []ArchivalTargetProgressInfo)`

SetArchivalRun sets ArchivalRun field to given value.

### HasArchivalRun

`func (o *GetProtectionRunProgressBody) HasArchivalRun() bool`

HasArchivalRun returns a boolean if a field has been set.

### SetArchivalRunNil

`func (o *GetProtectionRunProgressBody) SetArchivalRunNil(b bool)`

 SetArchivalRunNil sets the value for ArchivalRun to be an explicit nil

### UnsetArchivalRun
`func (o *GetProtectionRunProgressBody) UnsetArchivalRun()`

UnsetArchivalRun ensures that no value is present for ArchivalRun, not even an explicit nil
### GetLocalRun

`func (o *GetProtectionRunProgressBody) GetLocalRun() BackupRunProgressInfo`

GetLocalRun returns the LocalRun field if non-nil, zero value otherwise.

### GetLocalRunOk

`func (o *GetProtectionRunProgressBody) GetLocalRunOk() (*BackupRunProgressInfo, bool)`

GetLocalRunOk returns a tuple with the LocalRun field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalRun

`func (o *GetProtectionRunProgressBody) SetLocalRun(v BackupRunProgressInfo)`

SetLocalRun sets LocalRun field to given value.

### HasLocalRun

`func (o *GetProtectionRunProgressBody) HasLocalRun() bool`

HasLocalRun returns a boolean if a field has been set.

### GetReplicationRun

`func (o *GetProtectionRunProgressBody) GetReplicationRun() []ReplicationTargetProgressInfo`

GetReplicationRun returns the ReplicationRun field if non-nil, zero value otherwise.

### GetReplicationRunOk

`func (o *GetProtectionRunProgressBody) GetReplicationRunOk() (*[]ReplicationTargetProgressInfo, bool)`

GetReplicationRunOk returns a tuple with the ReplicationRun field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplicationRun

`func (o *GetProtectionRunProgressBody) SetReplicationRun(v []ReplicationTargetProgressInfo)`

SetReplicationRun sets ReplicationRun field to given value.

### HasReplicationRun

`func (o *GetProtectionRunProgressBody) HasReplicationRun() bool`

HasReplicationRun returns a boolean if a field has been set.

### SetReplicationRunNil

`func (o *GetProtectionRunProgressBody) SetReplicationRunNil(b bool)`

 SetReplicationRunNil sets the value for ReplicationRun to be an explicit nil

### UnsetReplicationRun
`func (o *GetProtectionRunProgressBody) UnsetReplicationRun()`

UnsetReplicationRun ensures that no value is present for ReplicationRun, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


